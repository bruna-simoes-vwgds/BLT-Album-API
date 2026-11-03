package models

type Album struct {
	AlbumID      string   `json:"album_id" bson:"album_id" binding:"required"`
	AlbumName    string   `json:"album_name" bson:"album_name" binding:"required"`
	ArtistName   string   `json:"artist_name" bson:"artist_name" binding:"required"`
	DateLaunched string   `json:"date_launched" bson:"date_launched" binding:"required"`
	DateAltered  string   `json:"date_altered" bson:"date_altered"`
	Songs        []string `json:"-" bson:"songs" binding:"required"`
}
