package database

import (
	"Magazine_backend/pkg/utils"
	"context"
	"io"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbMongo *mongo.Client
var err error
var DBMongoName = "products_db"

func ConnectMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	DbMongo, err = mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://localhost:27017",
	))
	if err != nil {
		log.Fatal("connect error:", err)
		return err
	}

	err = DbMongo.Ping(ctx, nil)
	if err != nil {
		utils.Logger.Fatal("ping error:", err)
		return err
	}

	utils.Logger.Printf("MongoDB connected")

	db := DbMongo.Database(DBMongoName)
	utils.Logger.Printf("MongoDB connected, name: %v", db.Name())
	return nil
}

func UploadImage(db *mongo.Database, filename string, r io.Reader) (string, error) {
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		return "", err
	}

	id, err := bucket.UploadFromStream(filename, r)
	if err != nil {
		return "", err
	}

	return id.Hex(), nil
}
