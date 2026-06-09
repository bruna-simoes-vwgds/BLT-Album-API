package models

type Song struct {
	SongID       string `json:"song_id" bson:"song_id"`
	SongName     string `json:"song_name" bson:"song_name"`
	ArtistName   string `json:"artist_name" bson:"artist_name"`
	DateLaunched string `json:"date_launched" bson:"date_launched"`
	DateAltered  string `json:"date_altered" bson:"date_altered"`
	SongLength   string `json:"song_length" bson:"song_length"`
}
