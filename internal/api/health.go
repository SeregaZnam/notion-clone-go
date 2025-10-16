package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Добро пожаловать в Notion Clone API!",
		"version": "1.0.0",
	})
}
