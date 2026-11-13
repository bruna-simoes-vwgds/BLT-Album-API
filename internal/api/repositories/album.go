package repositories

import (
	"album-app/internal/api/models"

	"github.com/bytedance/gopkg/util/logger"
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
	cursor, err := repo.collection.Find(c, bson.M{"album_id": id})

	if err != nil {
		logger.Error("Error finding album by id in database: ", err)
		return nil, err
	}

	var album []models.Album
	if err = cursor.All(c, &album); err != nil {
		logger.Error("Error iterating through album with id in database: ", err)
		return nil, err
	}

	if len(album) == 0 {
		return nil, nil
	}

	return &album[0], nil
}

func (repo *AlbumRepository) FindAlbums(c *gin.Context) ([]models.Album, error) {
	cursor, err := repo.collection.Find(c, bson.M{})

	if err != nil {
		logger.Error("Error finding albums in database: ", err)
		return nil, err
	}

	var albums []models.Album

	if err = cursor.All(c, &albums); err != nil {
		logger.Error("Error iterating albums in collection: ", err)
		return nil, err
	}

	return albums, nil
}

func (repo *AlbumRepository) CreateAlbum(c *gin.Context, album models.Album) error {
	_, err := repo.collection.InsertOne(c, album)

	if err != nil {
		return err
	}

	return nil
}

func (repo *AlbumRepository) DeleteAlbumById(c *gin.Context, id string) (bool, error) {
	deleteResult, err := repo.collection.DeleteOne(c, bson.M{"id": id})

	if err != nil {
		return false, err
	}

	return deleteResult.DeletedCount > 0, nil
}

func (repo *AlbumRepository) DeleteSongByIdInAlbum(c *gin.Context, id string) (int64, error) {
	filter := bson.M{"$pull": bson.M{"songs": bson.M{"$in": id}}}
	updateResults, err := repo.collection.UpdateMany(c, bson.M{}, filter)

	if err != nil {
		return 0, err
	}

	return updateResults.UpsertedCount, nil
}
