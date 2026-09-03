package routes

import (
	"album-app/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadAlbumRoutes(router *gin.RouterGroup, handler handlers.AlbumHandler) {
	album := router.Group("/album")
	{
		album.GET("/id/:id", handler.GetAlbumById)
		album.GET("/id/:id/songs", handler.GetSongsByAlbumId)
	}

	router.GET("/albums", handler.GetAlbums)
}
