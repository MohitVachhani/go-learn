package emailauth

import (
	"fmt"

	userservice "github.com/MohitVachhani/go-learn/cmd/service/user"
	authInterface "github.com/MohitVachhani/go-learn/pkg/structs/auth"
	userInterface "github.com/MohitVachhani/go-learn/pkg/structs/user"
	passwordutil "github.com/MohitVachhani/go-learn/pkg/utils/auth"
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

	var inputPassword = passwordutil.ConvertToEncryptedString(emailLoginInput.Password)

	if userPassword == inputPassword {
		return authInterface.EmailLoginOutput{
			Success: true,
		}
	}

	return authInterface.EmailLoginOutput{
		Success:   true,
		ErrorCode: "PASSWORD_NOT_CORRECT",
	}

}
