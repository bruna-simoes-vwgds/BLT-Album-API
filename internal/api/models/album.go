package models

type Album struct {
	AlbumID      string   `json:"album_id" bson:"album_id"`
	AlbumName    string   `json:"album_name" bson:"album_name"`
	ArtistName   string   `json:"artist_name" bson:"artist_name"`
	DateLaunched string   `json:"date_launched" bson:"date_launched"`
	DateAltered  string   `json:"date_altered" bson:"date_altered"`
	Songs        []string `json:"songs" bson:"songs"`
}
