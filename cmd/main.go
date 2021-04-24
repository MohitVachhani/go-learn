package main

import (
	"context"
	"log"
	"net/http"
	"time"

	userController "github.com/MohitVachhani/go-learn/cmd/controller/user"
	envUtil "github.com/MohitVachhani/go-learn/pkg/utils/env"
	mongoUtils "github.com/MohitVachhani/go-learn/pkg/utils/mongo"

	"github.com/gorilla/mux"
)

func initializeRoutes() {

	// init router
	var router = mux.NewRouter()

	// route handlers
	router.HandleFunc("/auth/email/signUp", userController.RegisterUser).Methods("POST")

	router.HandleFunc("/user/get", userController.GetUser).Methods("GET")

	// start server and throw error if anything goes wrong.
	port := ":" + envUtil.Get("PORT")
	log.Fatal(http.ListenAndServe(port, router))

}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	mongoClient := mongoUtils.MongoConnection(ctx)
	defer mongoClient.Disconnect(ctx)

	initializeRoutes()

}

/*

1. Auth apis. (Normal)
2. CRUD Categories api.
3. CRUD State api.
4. create a contact api in that category api.
5. upVote a contact api.

*/
