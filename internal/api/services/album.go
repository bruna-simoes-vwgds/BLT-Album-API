package services

import (
	"album-app/internal/api/models"
	"album-app/internal/api/repositories"

	"github.com/gin-gonic/gin"
)

type AlbumService struct {
	Repository *repositories.AlbumRepository
}

func NewAlbumService(repository *repositories.AlbumRepository) *AlbumService {
	return &AlbumService{Repository: repository}
}

func (albumService *AlbumService) GetAlbumById(c *gin.Context, id string) (*models.Album, error) {
	album, err := albumService.Repository.FindAlbumById(c, id)
	if err != nil {
		return &models.Album{}, err
	}

	return album, nil
}
