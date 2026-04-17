package routes

import (
	"album-app/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func HealthCheck(router *gin.RouterGroup, handler handlers.GeneralHandler) {
	router.GET("/health", handler.HealthCheck)
}

func LoadAlbumRoutes(router *gin.RouterGroup, handler handlers.AlbumHandler) {
	album := router.Group("/album")
	{
		album.GET("/id/:id", handler.GetAlbumById)
	}
}
