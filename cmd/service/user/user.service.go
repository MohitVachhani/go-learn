package userservice

import (
	"context"
	"fmt"
	"time"

	userRepo "github.com/MohitVachhani/go-learn/cmd/repo/user"
	registerUserInterface "github.com/MohitVachhani/go-learn/pkg/structs/auth"
	userInterface "github.com/MohitVachhani/go-learn/pkg/structs/user"
)

func RegisterUser(registerUserInput registerUserInterface.RegisterUserInput) {

	emailID := registerUserInput.EmailID

	var userFiters userInterface.UserFilters

	userFiters.EmailID = emailID

	user := GetUser(userFiters)

	fmt.Println(user)

}

// User service layer for getting user
func GetUser(userFilters userInterface.UserFilters) userInterface.User {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	user := userRepo.GetUser(ctx, userFilters)

	return user
}
