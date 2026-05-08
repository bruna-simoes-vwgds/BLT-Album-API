package routes

import (
	"album-app/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadSongRoutes(router *gin.RouterGroup, handler handlers.SongHandler) {
	album := router.Group("/song")
	{
		album.GET("/id/:id", handler.GetSongById)
	}
}
