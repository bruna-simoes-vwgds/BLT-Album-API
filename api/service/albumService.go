package service

import (
	"album-app/api/datamodel"
	"time"
)

func GetAlbumById(id string) (datamodel.Album, bool) {
	// TODO: connection to MongoDB repo to store and load data
	albums := populateAlbums()
	return getAlbumById(id, albums)
}

func getAlbumById(id string, albums []datamodel.Album) (datamodel.Album, bool) {
	for i, album := range albums {
		if album.AlbumID == id {
			return albums[i], true
		}
	}
	return datamodel.Album{}, false
}

// this will become obsolete once DB is implemented
func populateAlbums() []datamodel.Album {
	var models []datamodel.Album

	model1 := datamodel.Album{
		AlbumID:      "123",
		AlbumName:    "Steal This Song!",
		ArtistName:   "System of a Down",
		DateLaunched: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		DateAltered:  time.Now(),
		Songs:        []string{},
	}
	model2 := datamodel.Album{
		AlbumID:      "456",
		AlbumName:    "Queen",
		ArtistName:   "Queen",
		DateLaunched: time.Date(1980, time.September, 10, 23, 0, 0, 0, time.UTC),
		DateAltered:  time.Now(),
		Songs:        []string{},
	}

	models = append(models, model1, model2)

	return models
}
