package mongo

import (
	"context"
	"fmt"
	"log"

	envUtil "github.com/MohitVachhani/go-learn/pkg/utils/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// CreateMongoClient is
func createMongoClient(ctx context.Context) *mongo.Client {

	mongoURI := envUtil.Get("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongo!!")

	return client
}

// GetCollection is
func GetCollection(mongoClient *mongo.Client, collectionName string) *mongo.Collection {

	database := envUtil.Get("DATABASE")

	collectionModel := mongoClient.Database(database).Collection(collectionName)

	return collectionModel
}

func MongoConnection(ctx context.Context) *mongo.Client {

	MongoClient = createMongoClient(ctx)

	return MongoClient
}
