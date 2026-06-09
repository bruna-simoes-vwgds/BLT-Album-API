package repositories

import (
	"album-app/internal/api/models"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SongRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewSongRepository(db *mongo.Database) *SongRepository {
	return &SongRepository{
		db:         db,
		collection: db.Collection("songs"),
	}
}

func (repo *SongRepository) FindSongById(c *gin.Context, id string) (*models.Song, error) {
	cursor, err := repo.collection.Find(c, bson.M{"song_id": id})

	if err != nil {
		logger.Error("Error finding song by id in database: ", err)
		return nil, err
	}

	var song []models.Song
	if err = cursor.All(c, &song); err != nil {
		logger.Error("Error iterating through song with id in database: ", err)
		return nil, err
	}

	if len(song) == 0 {
		return nil, nil
	}

	return &song[0], nil
}
