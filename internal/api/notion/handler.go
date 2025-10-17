package notion

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	domainPage "github.com/SeregaZnam/notion-clone-go/internal/domain/page"
	domainTextBlock "github.com/SeregaZnam/notion-clone-go/internal/domain/text_block"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetPages(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx, `SELECT id, title FROM public.pages ORDER BY id`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query pages"})
		return
	}
	defer rows.Close()

	var pages []domainPage.Page
	for rows.Next() {
		var p domainPage.Page
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

func (h *Handler) GetTextBlocks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx, `SELECT id, title FROM public.pages ORDER BY id`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query pages"})
		return
	}
	defer rows.Close()

	var pages []domainPage.Page
	for rows.Next() {
		var p domainPage.Page
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

func (h *Handler) AddTextBlocks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var textBlock domainTextBlock.TextBlock
	if err := c.ShouldBindJSON(&textBlock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := h.db.Exec(ctx, `INSERT INTO public.text_blocks (text, "page_id", "order", "type") VALUES ($1, $2, $3, $4)`, textBlock.Text, textBlock.PageId, textBlock.Order, textBlock.Type)
	if err != nil {
		slog.Error("Failed to add page", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Text block added successfully"})
}

func (h *Handler) AddPage(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var page domainPage.Page
	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := h.db.Exec(ctx, `INSERT INTO public.pages (title, "icon_src", "icon_class", "cover_src") VALUES ($1, $2, $3, $4)`, page.Title, page.IconSrc, page.IconClass, page.CoverSrc)
	if err != nil {
		slog.Error("Failed to add page", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page added successfully"})
}
