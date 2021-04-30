package emailauth

import (
	"fmt"

	userservice "github.com/MohitVachhani/go-learn/cmd/service/user"
	authInterface "github.com/MohitVachhani/go-learn/pkg/structs/auth"
	userInterface "github.com/MohitVachhani/go-learn/pkg/structs/user"
	accessTokenUtil "github.com/MohitVachhani/go-learn/pkg/utils/auth/accessToken"
	passwordUtil "github.com/MohitVachhani/go-learn/pkg/utils/auth/password"
)

func EmailLogin(emailLoginInput authInterface.EmailLoginInput) authInterface.EmailLoginOutput {

	var userFilters userInterface.UserFilters = userInterface.UserFilters{
		EmailID: emailLoginInput.EmailID,
	}

	user := userservice.GetUser(userFilters)

	if (userInterface.User{} == user) {
		fmt.Println("User do not exists")

		return authInterface.EmailLoginOutput{
			Success:   false,
			ErrorCode: "USER_NOT_FOUND",
		}
	}

	var userPassword = user.Password

	var inputPassword = passwordUtil.ConvertToEncryptedString(emailLoginInput.Password)

	if userPassword == inputPassword {
		accessToken := accessTokenUtil.CreateAccessToken(authInterface.AccessTokenPayload{
			TokenUse: "accessToken",
			EmailId:  user.EmailID,
			UserId:   user.ID.Hex(),
		})

		refreshToken := accessTokenUtil.CreateRefreshToken(authInterface.RefreshTokenPayload{
			TokenUse: "refreshToken",
			UserId:   user.ID.Hex(),
		})

		return authInterface.EmailLoginOutput{
			Success: true,
			LoginToken: &authInterface.LoginToken{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		}
	}

	return authInterface.EmailLoginOutput{
		Success:   false,
		ErrorCode: "PASSWORD_NOT_CORRECT",
	}
}
