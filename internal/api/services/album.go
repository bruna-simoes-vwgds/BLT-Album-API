package services

import (
	"album-app/internal/api/models"
	"album-app/internal/api/repositories"

	"github.com/gin-gonic/gin"
)

type AlbumService struct {
	AlbumRepository *repositories.AlbumRepository
	SongRepository  *repositories.SongRepository
}

func NewAlbumService(repository *repositories.AlbumRepository, songRepository *repositories.SongRepository) *AlbumService {
	return &AlbumService{
		AlbumRepository: repository,
		SongRepository:  songRepository,
	}
}

func (albumService *AlbumService) GetAlbumById(c *gin.Context, id string) (*models.Album, error) {
	album, err := albumService.AlbumRepository.FindAlbumById(c, id)
	if err != nil {
		return &models.Album{}, err
	}

	return album, nil
}

func (albumService *AlbumService) GetAlbums(c *gin.Context) ([]models.Album, error) {
	albums, err := albumService.AlbumRepository.FindAlbums(c)

	return albums, err
}

func (albumService *AlbumService) CreateNewAlbum(c *gin.Context, album models.Album) ([]string, error) {
	var songs []string
	for _, songID := range album.Songs {
		song, err := albumService.SongRepository.FindSongById(c, songID)
		if err != nil {
			return nil, err
		}
		if song == nil {
			songs = append(songs, songID)
		}
	}

	err := albumService.AlbumRepository.CreateAlbum(c, album)

	return songs, err
}

func (albumService *AlbumService) GetAllSongsByAlbumId(c *gin.Context, id string) ([]models.Song, error) {
	album, err := albumService.AlbumRepository.FindAlbumById(c, id)

	if err != nil {
		return nil, err
	}
	if album == nil {
		return []models.Song{}, nil
	}

	songs := make([]models.Song, 0, len(album.Songs))
	for _, songID := range album.Songs {
		song, err := albumService.SongRepository.FindSongById(c, songID)
		if err != nil {
			return nil, err
		}
		if song != nil {
			songs = append(songs, *song)
		}
	}

	return songs, nil
}

func (albumService *AlbumService) DeleteAlbum(c *gin.Context, id string) (bool, error) {
	return albumService.AlbumRepository.DeleteAlbumById(c, id)
}
