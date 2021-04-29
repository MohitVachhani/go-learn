package emailrouter

import (
	"github.com/gorilla/mux"

	userController "github.com/MohitVachhani/go-learn/cmd/controller/user"
)

func InitalizeEmailAuthRouter(router *mux.Router) *mux.Router {

	router.HandleFunc("/email/signUp", userController.RegisterUser).Methods("POST")
	router.HandleFunc("/email/login", userController.EmailLogin).Methods("POST")

	return router
}
