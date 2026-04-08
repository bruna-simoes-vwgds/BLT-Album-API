package database

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Handler struct {
	MongoDb *mongo.Database
}
