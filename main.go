package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SeregaZnam/notion-clone-go/internal/api"
	"github.com/SeregaZnam/notion-clone-go/internal/env"
)

const serverReadTimeout = 60 * time.Second

func main() {
	// Создаем контекст, который будет отменен при получении SIGINT (Ctrl+C) или SIGTERM
	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	e, err := env.CreateAndInit(rootCtx)
	if err != nil {
		slog.Error("can't run server", "err", err)
		os.Exit(1)
	}
	handler := api.NewAPI(&e)

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", "0.0.0.0", 8080),
		Handler:           handler,
		ReadHeaderTimeout: serverReadTimeout,
	}

	// Запускаем сервер в отдельной горутине
	go func() {
		slog.Info("Starting server", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", "err", err)
		}
	}()

	// Ждем сигнал завершения
	<-rootCtx.Done()
	slog.Info("Shutting down server...")

	// Создаем контекст с таймаутом для graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Graceful shutdown сервера
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("Server forced to shutdown", "err", err)
	} else {
		slog.Info("Server exited gracefully")
	}
}
