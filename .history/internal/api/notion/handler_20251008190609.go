package notion

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	db *pgxpool.Pool
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetNotes(c *gin.Context) {
	// Читаем JSON файл с данными страниц
	jsonPath := filepath.Join("internal", "api", "notion", "notion-pages.json")
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read pages data"})
		return
	}

	// Парсим JSON данные
	var pages []map[string]interface{}
	if err := json.Unmarshal(data, &pages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse pages data"})
		return
	}

	// Возвращаем данные страниц
	c.JSON(http.StatusOK, gin.H{
		"pages": pages,
		"count": len(pages),
	})
}

func (h *Handler) GetPages(c *gin.Context) {
	// Чтение из БД
	type Page struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		IconSrc   string `json:"iconSrc"`
		IconClass string `json:"iconClass"`
		CoverSrc  string `json:"coverSrc"`
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx, `SELECT id, title FROM public.pages ORDER BY id`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query pages"})
		return
	}
	defer rows.Close()

	var pages []Page
	for rows.Next() {
		var p Page
		if err := rows.Scan(&p.ID, &p.Title); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan page"})
			return
		}
		pages = append(pages, p)
	}
	if rows.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Rows error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pages": pages, "count": len(pages)})
}
