package handlers

import (
	"album-app/internal/api/database"
	"album-app/internal/api/models"
	"album-app/internal/api/repositories"
	"album-app/internal/api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	Service *services.SongService
}

func NewSongHandler(db *database.DBHandler) *SongHandler {
	albumRepository := repositories.NewAlbumRepository(db.MongoDb)
	songRepository := repositories.NewSongRepository(db.MongoDb)
	songService := services.NewSongService(songRepository, albumRepository)
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

func (songHandler *SongHandler) AddSong(c *gin.Context) {
	var song models.Song

	err := c.ShouldBindJSON(&song)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CreateResponseMessage("Request body malformed"))
		return
	}

	duplicatedSong, errSong := songHandler.Service.GetSongById(c, song.SongID)

	if errSong != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage("Error in database search: "+errSong.Error()))
		return
	}

	if duplicatedSong != nil {
		c.JSON(http.StatusConflict, models.CreateResponseMessage("A song entry already exists with that id"))
		return
	}

	err = songHandler.Service.CreateNewSong(c, song)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage("Error occurred while creating the song: "+err.Error()))
		return
	}

	c.JSON(http.StatusCreated, models.CreateResponseMessage("Song was created successfully"))
}

func (songHandler *SongHandler) DeleteSongById(c *gin.Context) {
	id := c.Param("id")

	hasDeletedDocuments, removedFromAlbumsCount, err := songHandler.Service.DeleteSong(c, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.CreateResponseMessage("Couldn't delete song: "+err.Error()))
		return
	}

	if !hasDeletedDocuments {
		c.JSON(http.StatusNotFound, models.CreateResponseMessage("No songs found with id "+id))
		return
	}

	c.JSON(http.StatusOK, models.CreateResponseMessage("Song was deleted successfully. Removed from "+
		strconv.FormatInt(removedFromAlbumsCount, 10)+"albums."))
}
