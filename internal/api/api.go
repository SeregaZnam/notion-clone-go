package api

import (
	middleware "github.com/SeregaZnam/notion-clone-go/internal/api/middleware"
	"github.com/SeregaZnam/notion-clone-go/internal/api/notion"
	"github.com/SeregaZnam/notion-clone-go/internal/env"

	"github.com/gin-gonic/gin"
)

func NewAPI(e *env.Env) *gin.Engine {
	r := gin.New()

	r.Use(middleware.CorsMiddleware())

	r.GET("/", Health)

	registerNotionRoutes(r, e.NotionHandler)

	return r
}

func registerNotionRoutes(r *gin.Engine, h *notion.Handler) {
	r.GET("/pages", h.GetPages)
	r.POST("/pages", h.AddPage)

	r.GET("/text-blocks", h.GetTextBlocks)
	r.POST("/text-blocks", h.AddTextBlocks)
}
