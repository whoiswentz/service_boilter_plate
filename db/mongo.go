package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var (
	MongoHost = os.Getenv("MONGO_HOST")
	MongoPort = os.Getenv("MONGO_PORT")
	MongoUser = os.Getenv("MONGO_USER")
	MongoPass = os.Getenv("MONGO_PASS")
)

func NewMongoConnection(dbName string) *mongo.Client {
	logFlags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stdout, "MONGO - ", logFlags)

	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v", MongoUser, MongoPass, MongoHost, MongoPort)
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(opts)
	if err != nil {
		logger.Fatalln(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		logger.Fatalf("Error while connecting to MongoDB: %v\n", err)
	}

	return client
}
