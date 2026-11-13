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
	albumService := services.NewAlbumService(albumRepository, songRepository)
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
		return
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

	songsNotFound, err := albumHandler.Service.CreateNewAlbum(c, album)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage("Error occurred while creating the album: "+err.Error()))
		return
	}

	response := models.CreateResponseMessage("Album was created successfully")
	if len(songsNotFound) != 0 {
		response.Warning = "Some songs not found in the database were not added to the album: "
		for i, song := range songsNotFound {
			if i == len(songsNotFound)-1 {
				response.Warning += song
			} else {
				response.Warning += " ,"
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

func (albumHandler *AlbumHandler) DeleteAlbumById(c *gin.Context) {
	id := c.Param("id")
	hasDeletedDocuments, err := albumHandler.Service.DeleteAlbum(c, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage("Couldn't delete album: "+err.Error()))
		return
	}

	if !hasDeletedDocuments {
		c.JSON(http.StatusNotFound, models.CreateResponseMessage("No albums found id "+id))
		return
	}

	c.JSON(http.StatusOK, models.CreateResponseMessage("Album was deleted successfully"))
}
