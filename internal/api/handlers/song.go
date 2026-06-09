package handlers

import (
	"album-app/internal/api/database"
	"album-app/internal/api/models"
	"album-app/internal/api/repositories"
	"album-app/internal/api/services"
	"net/http"
	"time"

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
		c.JSON(http.StatusInternalServerError, models.Error{Message: "Error getting song: " + err.Error(), Timestamp: time.Now().String()})
		return
	}

	if song == nil {
		c.JSON(http.StatusNotFound, models.Error{Message: "Song not found", Timestamp: time.Now().String()})
		return
	}

	c.JSON(http.StatusOK, song)
}
