package routes

import (
	"album-app/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadSongRoutes(router *gin.RouterGroup, handler handlers.SongHandler) {
	song := router.Group("/song")
	{
		song.GET("/id/:id", handler.GetSongById)
	}

	router.GET("/songs", handler.GetAllSongs)
}
