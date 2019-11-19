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
)

func OpenConnection(dbName string) *mongo.Database {
	logFlags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stdout, "MONGO - ", logFlags)

	if MongoHost == "" && MongoPort == "" {
		logger.Fatalln("Mongodb host and port need to be present")
	}

	uri := fmt.Sprintf("mongodb://root:root@%v:%v", MongoHost, MongoPort)
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(opts)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Println("Connecting to MongoDB")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		logger.Fatalf("Error while connecting to MongoDB: %v\n", err)
	}

	return client.Database("task")
}
