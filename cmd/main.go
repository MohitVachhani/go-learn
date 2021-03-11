package main

import (
	"context"
	"fmt"
	"log"
	"time"

	envUtil "github.com/MohitVachhani/go-learn/pkg/utils/env"
	mongoUtils "github.com/MohitVachhani/go-learn/pkg/utils/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectMongoAndQueryDatabase() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	mongoURI := envUtil.Get("MONGO_URI")

	mongoClient := mongoUtils.CreateMongoClient(ctx, mongoURI)

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
4. Make a util function for creating mongo client.

To Do:
5. Fetch all the data from a collection.
6. Make a repo layer for the same collection.
7. Create endpoints regarding CRUD-user.
8. Deploy the changes on heroku.
9. Automate the deployment method.
10. Best project structure.
*/
