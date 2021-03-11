package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MohitVachhani/go-learn/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func createMongoClient(ctx context.Context) *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(utils.Get("MONGO_URI")))

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

func connectMongoAndQueryDatabase() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	mongoClient := createMongoClient(ctx)

	defer mongoClient.Disconnect(ctx)

	err := mongoClient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	// BSON is the binary encoding of JSON-like documents that MongoDB uses when storing documents in collections
	databaseNames, err := mongoClient.ListDatabaseNames(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("databaseNames:", databaseNames)
}

func main() {
	connectMongoAndQueryDatabase()
}

/**
Done:
1. Find how to call functions which are in another file
2. Find how to call functions which are in another package
3. Learn how to use env variables.

To Do:
4. Make a function for creating mongo client.
5. Fetch all the data from a collection.
6. Make a repo layer for the same collection.
7. Create endpoints regarding CRUD-user.
8. Deploy the changes on heroku.
9. Automate the deployment method.
10. Best project structure.
*/
