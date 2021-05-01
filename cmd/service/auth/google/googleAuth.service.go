package googleauth

import (
	"encoding/json"
	"fmt"
	"net/http"

	authInterface "github.com/MohitVachhani/go-learn/pkg/structs/auth"
	userInterface "github.com/MohitVachhani/go-learn/pkg/structs/user"
	accessTokenUtil "github.com/MohitVachhani/go-learn/pkg/utils/auth/accessToken"

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
	user := userService.GetUser(userInterface.UserFilters{EmailID: providerProfile.Email})

	if stateJson.AuthType == "login" {
		if (user != userInterface.User{}) {
			accessToken := accessTokenUtil.CreateAccessToken(authInterface.AccessTokenPayload{
				TokenUse: "accessToken",
				EmailId:  user.EmailID,
				UserId:   user.ID.Hex(),
			})

			refreshToken := accessTokenUtil.CreateRefreshToken(authInterface.RefreshTokenPayload{
				TokenUse: "refreshToken",
				UserId:   user.ID.Hex(),
			})

			json.NewEncoder(res).Encode(authInterface.GoogleAuthPayload{
				Success: true,
				LoginToken: &authInterface.LoginToken{
					AccessToken:  accessToken,
					RefreshToken: refreshToken,
				}})
		} else {
			json.NewEncoder(res).Encode(authInterface.GoogleAuthPayload{Success: false, ErrorCode: "ACCOUNT_DO_NOT_EXISTS"})
		}
	} else {
		if (user == userInterface.User{}) {
			userService.CreateUser(userInterface.CreateUserInput{
				EmailID:    providerProfile.Email,
				SignUpType: "google",
				FirstName:  providerProfile.FirstName,
				LastName:   providerProfile.LastName,
				Status:     "active",
			})
			json.NewEncoder(res).Encode(authInterface.GoogleAuthPayload{Success: true})
		} else {
			json.NewEncoder(res).Encode(authInterface.GoogleAuthPayload{Success: false, ErrorCode: "ACCOUNT_ALREADY_EXISTS"})
		}
	}

	if err != nil {
		fmt.Fprintln(res, err)
		return
	}

}
