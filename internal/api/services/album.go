package services

import (
	"album-app/internal/api/models"
	"album-app/internal/api/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

type AlbumService struct {
	Repository *repositories.AlbumRepository
}

func NewAlbumService(repository *repositories.AlbumRepository) *AlbumService {
	return &AlbumService{Repository: repository}
}

func (albumService *AlbumService) GetAlbumById(c *gin.Context, id string) (models.Album, bool) {

	album, err := albumService.Repository.FindAlbumById(c, id)
	if err != nil {
		return models.Album{}, false
	}

	if album == nil {
		return models.Album{}, false
	}

	return *album, true
}

func getAlbumById(id string, albums []models.Album) (models.Album, bool) {
	for i, album := range albums {
		if album.AlbumID == id {
			return albums[i], true
		}
	}
	return models.Album{}, false
}

// this will become obsolete once DB is implemented
func populateAlbums() []models.Album {
	var modelsList []models.Album

	model1 := models.Album{
		AlbumID:      "123",
		AlbumName:    "Steal This Song!",
		ArtistName:   "System of a Down",
		DateLaunched: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		DateAltered:  time.Now(),
		Songs:        []string{},
	}
	model2 := models.Album{
		AlbumID:      "456",
		AlbumName:    "Queen",
		ArtistName:   "Queen",
		DateLaunched: time.Date(1980, time.September, 10, 23, 0, 0, 0, time.UTC),
		DateAltered:  time.Now(),
		Songs:        []string{},
	}

	modelsList = append(modelsList, model1, model2)

	return modelsList
}
