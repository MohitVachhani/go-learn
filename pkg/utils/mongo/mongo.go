package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/MohitVachhani/go-learn/pkg/utils/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateMongoClient is
func CreateMongoClient(ctx context.Context, mongoURI string) *mongo.Client {

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

	database := env.Get("DATABASE")

	collectionModel := mongoClient.Database(database).Collection(collectionName)

	return collectionModel
}
