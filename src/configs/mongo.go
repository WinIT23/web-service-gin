package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var singletonInstance *mongo.Client
var lock = &sync.Mutex{}

const dbName = "golang-api"

var uri = os.Getenv("MONGO_URL")

func GetMongoClient() *mongo.Client {
	if singletonInstance == nil {
		lock.Lock()
		singletonInstance = connectDB(context.Background())
		lock.Unlock()
	}
	return singletonInstance
}

func connectDB(ctx context.Context) *mongo.Client {
	if uri == "" {
		return nil
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("[ConnectDB] Connected to MongoDB Successfully.")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
