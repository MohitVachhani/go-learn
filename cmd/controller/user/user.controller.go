package usercontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	userService "github.com/MohitVachhani/go-learn/cmd/service/user"
	registerUserInterface "github.com/MohitVachhani/go-learn/pkg/structs/auth"
	userinterface "github.com/MohitVachhani/go-learn/pkg/structs/user"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	// body parameters
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var registerUserInput registerUserInterface.RegisterUserInput

	json.Unmarshal(body, &registerUserInput)

	userService.RegisterUser(registerUserInput)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	// Query parameters
	var emailID = r.URL.Query().Get("emailId")

	var userFilters userinterface.UserFilters

	userFilters.EmailID = emailID

	user := userService.GetUser(userFilters)

	// returns the client with json.
	json.NewEncoder(w).Encode(user)
}
