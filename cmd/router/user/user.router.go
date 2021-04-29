package userrouter

import (
	userController "github.com/MohitVachhani/go-learn/cmd/controller/user"

	"github.com/gorilla/mux"
)

func InitUserR(router *mux.Router) *mux.Router {

	router.HandleFunc("/get", userController.GetUser).Methods("GET")

	return router
}
