package googleauth

import (
	"encoding/json"
	"fmt"
	"net/http"

	authInterface "github.com/MohitVachhani/go-learn/pkg/structs/auth"
	userinterface "github.com/MohitVachhani/go-learn/pkg/structs/user"

	userService "github.com/MohitVachhani/go-learn/cmd/service/user"
	"github.com/markbates/goth/gothic"
)

func HandleGoogleAuth(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}

func HandleGoogleAuthCallback(res http.ResponseWriter, req *http.Request) {

	var state = req.URL.Query().Get("state")

	var stateJson authInterface.GoogleAuthInput

	json.Unmarshal([]byte(state), &stateJson)

	providerProfile, err := gothic.CompleteUserAuth(res, req)

	// profilePicture is in providerProfile.rawData

	if stateJson.AuthType == "login" {
		user := userService.GetUser(userinterface.UserFilters{EmailID: providerProfile.Email})
		fmt.Println("user", user)
	} else {
		userService.CreateUser(userinterface.CreateUserInput{
			EmailID:    providerProfile.Email,
			SignUpType: "google",
			FirstName:  providerProfile.FirstName,
			LastName:   providerProfile.LastName,
			Status:     "active",
		})
	}

	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

}
