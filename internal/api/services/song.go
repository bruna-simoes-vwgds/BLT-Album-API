package services

import (
	"album-app/internal/api/models"
	"album-app/internal/api/repositories"

	"github.com/gin-gonic/gin"
)

type SongService struct {
	Repository *repositories.SongRepository
}

func NewSongService(repository *repositories.SongRepository) *SongService {
	return &SongService{Repository: repository}
}

func (songService *SongService) GetSongById(c *gin.Context, id string) (*models.Song, error) {

	song, err := songService.Repository.FindSongById(c, id)
	if err != nil {
		return &models.Song{}, err
	}

	return song, err
}
