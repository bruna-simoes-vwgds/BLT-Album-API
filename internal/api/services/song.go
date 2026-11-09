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
		return nil, err
	}

	return song, err
}

func (songService *SongService) GetAllSongs(c *gin.Context) ([]models.Song, error) {
	songs, err := songService.Repository.GetAllSongs(c)

	if err != nil {
		return []models.Song{}, err
	}

	return songs, nil
}

func (songService *SongService) CreateNewSong(c *gin.Context, song models.Song) (*models.DefaultResponse, error) {
	var response models.DefaultResponse

	err := songService.Repository.CreateSong(c, song)

	if err == nil {
		response = models.CreateResponseMessage("Album was created successfully")
	} else {
		response = models.CreateResponseMessage("Couldn't create album: " + err.Error())
	}

	return &response, err
}
