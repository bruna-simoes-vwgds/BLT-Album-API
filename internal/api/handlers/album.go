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

type AlbumHandler struct {
	Service *services.AlbumService
}

func NewAlbumHandler(db *database.Handler) *AlbumHandler {
	albumRepository := repositories.NewAlbumRepository(db.MongoDb)
	albumService := services.NewAlbumService(albumRepository)
	return &AlbumHandler{Service: albumService}
}

func (albumHandler *AlbumHandler) GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, foundAlbum := albumHandler.Service.GetAlbumById(c, id)

	if foundAlbum {
		c.JSON(http.StatusOK, album)
	} else {
		c.JSON(http.StatusNotFound, models.Error{Message: "Album not found", Timestamp: time.Now().String()})
	}
}
