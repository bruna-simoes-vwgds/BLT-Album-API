package handlers

import (
	"album-app/internal/api/database"
	"album-app/internal/api/models"
	"album-app/internal/api/repositories"
	"album-app/internal/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	Service *services.AlbumService
}

func NewAlbumHandler(db *database.DBHandler) *AlbumHandler {
	albumRepository := repositories.NewAlbumRepository(db.MongoDb)
	songRepository := repositories.NewSongRepository(db.MongoDb)
	songService := services.NewSongService(songRepository)
	albumService := services.NewAlbumService(albumRepository, songService)
	return &AlbumHandler{Service: albumService}
}

func (albumHandler *AlbumHandler) GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, errAlbum := albumHandler.Service.GetAlbumById(c, id)

	if errAlbum != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage(errAlbum.Error()))
		return
	}

	if album == nil {
		c.JSON(http.StatusNotFound, models.CreateResponseMessage("Album not found"))
		return
	}

	c.JSON(http.StatusOK, album)
}

func (albumHandler *AlbumHandler) GetAlbums(c *gin.Context) {
	albums, errAlbums := albumHandler.Service.GetAlbums(c)

	if errAlbums != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage(errAlbums.Error()))
		return
	}

	c.JSON(http.StatusOK, albums)
}

func (albumHandler *AlbumHandler) GetSongsByAlbumId(c *gin.Context) {
	id := c.Param("id")
	songs, errSongs := albumHandler.Service.GetAllSongsByAlbumId(c, id)

	if errSongs != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage(errSongs.Error()))
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (albumHandler *AlbumHandler) AddAlbum(c *gin.Context) {
	var album models.Album

	err := c.ShouldBindJSON(&album)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResponseMessage("Request body malformed"))
	}

	duplicatedAlbum, errAlbum := albumHandler.Service.GetAlbumById(c, album.AlbumID)

	if errAlbum != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage("Error in database search: "+errAlbum.Error()))
		return
	}

	if duplicatedAlbum != nil {
		c.JSON(http.StatusConflict, models.CreateResponseMessage("An album entry already exists with that id"))
		return
	}

	response, err := albumHandler.Service.CreateNewAlbum(c, album)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage("Error occurred while creating the album: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}
