package server

import (
	"commerce-platform/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	healthHandler := handler.NewHealthHandler()

	r.GET("/health", healthHandler.Check)

	api := r.Group("/api")
	{
		api.GET("/health", healthHandler.Check)
	}

	return r
}
