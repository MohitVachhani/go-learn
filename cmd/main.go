package main

import (
	"context"
	"log"
	"net/http"
	"time"

	userRouter "github.com/MohitVachhani/go-learn/cmd/router/user"
	envUtil "github.com/MohitVachhani/go-learn/pkg/utils/env"
	mongoUtils "github.com/MohitVachhani/go-learn/pkg/utils/mongo"

	emailAuthRouter "github.com/MohitVachhani/go-learn/cmd/router/auth/email"

	"github.com/gorilla/mux"
)

func initializeRoutes() {

	// init router
	var router = mux.NewRouter()

	// user router
	userR := router.PathPrefix("/user").Subrouter()
	userRouter.InitUserRouter(userR)

	// auth router
	authR := router.PathPrefix("/auth").Subrouter()
	emailAuthRouter.InitalizeEmailAuthRouter(authR)

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
