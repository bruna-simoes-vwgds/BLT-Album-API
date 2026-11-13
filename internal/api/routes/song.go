package routes

import (
	"album-app/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadSongRoutes(router *gin.RouterGroup, handler handlers.SongHandler) {
	song := router.Group("/song")
	{
		song.GET("/id/:id", handler.GetSongById)
		song.DELETE("/id/:id", handler.DeleteSongById)
		song.POST("", handler.AddSong)
	}

	router.GET("/songs", handler.GetAllSongs)
}
