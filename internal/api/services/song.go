package services

import (
	"album-app/internal/api/models"
	"album-app/internal/api/repositories"

	"github.com/gin-gonic/gin"
)

type SongService struct {
	SongRepository  *repositories.SongRepository
	AlbumRepository *repositories.AlbumRepository
}

func NewSongService(songRepository *repositories.SongRepository, albumRepository *repositories.AlbumRepository) *SongService {
	return &SongService{SongRepository: songRepository, AlbumRepository: albumRepository}
}

func (songService *SongService) GetSongById(c *gin.Context, id string) (*models.Song, error) {
	song, err := songService.SongRepository.FindSongById(c, id)
	if err != nil {
		return nil, err
	}

	return song, err
}

func (songService *SongService) GetAllSongs(c *gin.Context) ([]models.Song, error) {
	songs, err := songService.SongRepository.GetAllSongs(c)

	if err != nil {
		return []models.Song{}, err
	}

	return songs, nil
}

func (songService *SongService) CreateNewSong(c *gin.Context, song models.Song) error {
	return songService.SongRepository.CreateSong(c, song)
}

func (songService *SongService) DeleteSong(c *gin.Context, id string) (bool, int64, error) {
	hasDeletedDocuments, err := songService.SongRepository.DeleteSongById(c, id)
	if err != nil {
		return false, 0, err
	}

	//not only do we have to remove the song from the service, but also remove it from the album list
	upsertedCount, err := songService.AlbumRepository.DeleteSongByIdInAlbum(c, id)

	if err != nil {
		return hasDeletedDocuments, 0, err
	}

	return hasDeletedDocuments, upsertedCount, err
}
