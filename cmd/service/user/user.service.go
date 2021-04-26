package userservice

import (
	"fmt"

	userRepo "github.com/MohitVachhani/go-learn/cmd/repo/user"
	registerUserInterface "github.com/MohitVachhani/go-learn/pkg/structs/auth"

	userInterface "github.com/MohitVachhani/go-learn/pkg/structs/user"
	passwordUtil "github.com/MohitVachhani/go-learn/pkg/utils/auth/password"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterUser(registerUserInput registerUserInterface.RegisterUserInput) {

	emailID := registerUserInput.EmailID

	var userFiters userInterface.UserFilters

	userFiters.EmailID = emailID

	user := GetUser(userFiters)

	if (userInterface.User{} == user) {
		fmt.Println("no user found with the following emailId", emailID)

		var createUserInput userInterface.CreateUserInput
		var encryptedPassword = passwordUtil.ConvertToEncryptedString(registerUserInput.Password)

		createUserInput.EmailID = emailID
		createUserInput.Password = encryptedPassword
		createUserInput.SignUpType = "email"

		createUser(createUserInput)

	} else {
		fmt.Println("User already exists with this emailId:", emailID)
	}
}

// User service layer for getting user
func GetUser(userFilters userInterface.UserFilters) userInterface.User {

	user := userRepo.GetUser(userFilters)

	return user
}

func createUser(createUserInput userInterface.CreateUserInput) bson.M {

	user := userRepo.CreateUser(createUserInput)

	return user
}
