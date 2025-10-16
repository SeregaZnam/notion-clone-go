package notion

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
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
