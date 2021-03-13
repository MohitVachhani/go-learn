package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	userRepo "github.com/MohitVachhani/go-learn/cmd/repo/user"
	"github.com/MohitVachhani/go-learn/pkg/utils/env"
	envUtil "github.com/MohitVachhani/go-learn/pkg/utils/env"
	mongoUtils "github.com/MohitVachhani/go-learn/pkg/utils/mongo"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
)

func connectMongoAndQueryDatabase() bson.M {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	mongoURI := envUtil.Get("MONGO_URI")

	mongoClient := mongoUtils.CreateMongoClient(ctx, mongoURI)

	defer mongoClient.Disconnect(ctx)

	userCollection := mongoUtils.GetCollection(mongoClient, "users")

	var user bson.M = userRepo.GetUserByEmailID(ctx, userCollection, "mohitvachhanispam@gmail.com")

	fmt.Println("user name is", user["firstName"], user["lastName"])

	return user
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hitting test endpoint")

	// GET params
	params := mux.Vars(r)

	// Body params
	// _ := json.NewDecoder(w).Decode(user)

	fmt.Println("hello from server", params)

	user := connectMongoAndQueryDatabase()

	// returns the client with json.
	json.NewEncoder(w).Encode(user)
}

func server() {

	// init router
	var r = mux.NewRouter()

	// route handlers
	r.HandleFunc("/user/get", getUser).Methods("GET")

	// start server and throw error if anything goes wrong.
	port := ":" + env.Get("PORT")
	log.Fatal(http.ListenAndServe(port, r))
}

func main() {
	server()
}

/**
Done:
1. Find how to call functions which are in another file
2. Find how to call functions which are in another package
3. Learn how to use env variables.
4. Make a util function for creating mongo client.
5. Fetch all the data from a collection.
7. setup mux for endpoint.

To Do:
6. Make a repo layer for the same collection.
8. Create endpoints regarding CRUD-user.
9. Deploy the changes on heroku.
10. Automate the deployment method.
*/
