package models

import "time"

type Song struct {
	SongID       string    `json:"song_id"`
	SongName     string    `json:"song_name"`
	ArtistName   string    `json:"artist_name"`
	DateLaunched time.Time `json:"date_launched"`
	DateAltered  time.Time `json:"date_altered"`
	SongLength   []string  `json:"song_length"`
}
