package models

import "time"

type Album struct {
	AlbumID      string    `json:"album_id"`
	AlbumName    string    `json:"album_name"`
	ArtistName   string    `json:"artist_name"`
	DateLaunched time.Time `json:"date_launched"`
	DateAltered  time.Time `json:"date_altered"`
	Songs        []string  `json:"songs"`
	//TODO: change this so it receives slice of Song objects, not string
}
