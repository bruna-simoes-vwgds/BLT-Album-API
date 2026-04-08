package repositories

import (
	"album-app/internal/api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AlbumRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewAlbumRepository(db *mongo.Database) *AlbumRepository {
	return &AlbumRepository{
		db:         db,
		collection: db.Collection("albums"),
	}
}

func (repo *AlbumRepository) FindAlbumById(c *gin.Context, id string) (*models.Album, error) {
	cursor, err := repo.collection.Find(c, bson.M{"id": id})

	if err != nil {
		return nil, err
	}

	var album models.Album
	if err := cursor.All(c, &album); err != nil {
		return nil, err
	}

	return &album, nil
}
