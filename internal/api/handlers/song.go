package handlers

import (
	"album-app/internal/api/database"
	"album-app/internal/api/models"
	"album-app/internal/api/repositories"
	"album-app/internal/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	Service *services.SongService
}

func NewSongHandler(db *database.DBHandler) *SongHandler {
	songRepository := repositories.NewSongRepository(db.MongoDb)
	songService := services.NewSongService(songRepository)
	return &SongHandler{Service: songService}
}

func (songHandler *SongHandler) GetSongById(c *gin.Context) {
	id := c.Param("id")
	song, err := songHandler.Service.GetSongById(c, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage("Error getting song: "))
		return
	}

	if song == nil {
		c.JSON(http.StatusNotFound, models.CreateResponseMessage("Song not found"))
		return
	}

	c.JSON(http.StatusOK, song)
}

func (songHandler *SongHandler) GetAllSongs(c *gin.Context) {
	songs, err := songHandler.Service.GetAllSongs(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage("Error getting all songs: "+err.Error()))
		return
	}

	if len(songs) == 0 {
		c.JSON(http.StatusNoContent, models.CreateResponseMessage("No songs found."))
		return
	}

	c.JSON(http.StatusOK, songs)
}
