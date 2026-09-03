package services

import (
	"album-app/internal/api/models"
	"album-app/internal/api/repositories"

	"github.com/gin-gonic/gin"
)

type AlbumService struct {
	Repository  *repositories.AlbumRepository
	SongService *SongService
}

func NewAlbumService(repository *repositories.AlbumRepository, songService *SongService) *AlbumService {
	return &AlbumService{
		Repository:  repository,
		SongService: songService,
	}
}

func (albumService *AlbumService) GetAlbumById(c *gin.Context, id string) (*models.Album, error) {
	album, err := albumService.Repository.FindAlbumById(c, id)
	if err != nil {
		return &models.Album{}, err
	}

	return album, nil
}

func (albumService *AlbumService) GetAlbums(c *gin.Context) ([]models.Album, error) {
	albums, err := albumService.Repository.FindAlbums(c)

	return albums, err
}

func (albumService *AlbumService) GetAllSongsByAlbumId(c *gin.Context, id string) ([]models.Song, error) {
	album, err := albumService.Repository.FindAlbumById(c, id)

	if err != nil {
		return nil, err
	}
	if album == nil {
		return []models.Song{}, nil
	}

	songs := make([]models.Song, 0, len(album.Songs))
	for _, songID := range album.Songs {
		song, err := albumService.SongService.GetSongById(c, songID)
		if err != nil {
			return nil, err
		}
		if song != nil {
			songs = append(songs, *song)
		}
	}

	return songs, nil
}
