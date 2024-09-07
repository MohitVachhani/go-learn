package resourcerouter

import (
	ResourceController "github.com/MohitVachhani/go-learn/cmd/controller/resource"
	"github.com/gorilla/mux"
)

func InitializeResourceRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/resource/get", ResourceController.GetResource).Methods("GET")
	// router.HandleFunc("/resource/create", resourceController.createResource).Methods("POST")

	return router
}
