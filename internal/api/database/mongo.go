package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBHandler struct {
	MongoDb *mongo.Database
}

func (DBHandler *DBHandler) InitializeMongoDB() {
	val, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		log.Fatal("MONGO_URI environment variable not set")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(val))

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("public")

	DBHandler.MongoDb = db
}

func (DBHandler *DBHandler) DisconnectMongoDB() {
	err := DBHandler.MongoDb.Client().Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
