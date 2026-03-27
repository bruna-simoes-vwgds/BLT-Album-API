package datamodel

import "time"

type Album struct {
	AlbumID      string
	AlbumName    string
	ArtistName   string
	DateLaunched time.Time
	DateAltered  time.Time
	Songs        []string //TODO: change this so it receives slice of Song objects, not string
}
