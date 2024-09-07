package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	userRouter "github.com/MohitVachhani/go-learn/cmd/router/user"
	envUtil "github.com/MohitVachhani/go-learn/pkg/utils/env"
	mongoUtils "github.com/MohitVachhani/go-learn/pkg/utils/mongo"

	emailAuthRouter "github.com/MohitVachhani/go-learn/cmd/router/auth/email"
	googleAuthRouter "github.com/MohitVachhani/go-learn/cmd/router/auth/google"
	resourceRouter "github.com/MohitVachhani/go-learn/cmd/router/resource"

	"github.com/gorilla/mux"
)

func initializeRoutes() {
	fmt.Println("Initializing routes")

	// init router
	var router = mux.NewRouter()

	// user router
	userR := router.PathPrefix("/user").Subrouter()
	userRouter.InitUserRouter(userR)

	// auth router
	authR := router.PathPrefix("/auth").Subrouter()
	emailAuthRouter.InitalizeEmailAuthRouter(authR)
	googleAuthRouter.InitializeGoogleAuthRouter(authR)

	// resource route
	resourceR := router.PathPrefix("/resource").Subrouter()
	resourceRouter.InitializeResourceRouter(resourceR)

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
