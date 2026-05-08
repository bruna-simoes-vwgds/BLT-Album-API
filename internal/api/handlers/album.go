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

func NewAlbumHandler(db *database.DBHandler) *AlbumHandler {
	albumRepository := repositories.NewAlbumRepository(db.MongoDb)
	albumService := services.NewAlbumService(albumRepository)
	return &AlbumHandler{Service: albumService}
}

func (albumHandler *AlbumHandler) GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, errAlbum := albumHandler.Service.GetAlbumById(c, id)

	if errAlbum != nil {
		c.JSON(http.StatusInternalServerError, models.Error{Message: errAlbum.Error(), Timestamp: time.Now().String()})
		return
	}

	if album == nil {
		c.JSON(http.StatusNotFound, models.Error{Message: "Album not found", Timestamp: time.Now().String()})
		return
	}

	c.JSON(http.StatusOK, album)
}
