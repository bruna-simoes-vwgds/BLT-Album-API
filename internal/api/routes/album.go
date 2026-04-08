package routes

import (
	"album-app/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadAlbumRoutes(router *gin.RouterGroup, handler handlers.AlbumHandler) {
	router.Group("/album")
	{
		router.GET("/id/:id", handler.GetAlbumById)
	}
}
