package user

import (
	"context"
	"fmt"
	"log"
	"time"

	userInterface "github.com/MohitVachhani/go-learn/pkg/structs/user"
	mongoUtils "github.com/MohitVachhani/go-learn/pkg/utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

//GetUserByEmailID is
func GetUserByEmailID(ctx context.Context, emailID string) userInterface.User {

	userCollection := mongoUtils.GetCollection(mongoUtils.MongoClient, "users")

	var user userInterface.User

	err := userCollection.FindOne(ctx, bson.M{"emailId": emailID}).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	return user
}

func GetUser(userFilters userInterface.UserFilters) userInterface.User {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	userCollection := mongoUtils.GetCollection(mongoUtils.MongoClient, "users")

	var user userInterface.User

	var emailID = userFilters.EmailID

	userCollection.FindOne(ctx, bson.M{"emailId": emailID}).Decode(&user)

	return user
}

func CreateUser(input userInterface.CreateUserInput) bson.M {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	userCollection := mongoUtils.GetCollection(mongoUtils.MongoClient, "users")

	emailId := input.EmailID

	var createUserInput userInterface.User = userInterface.User{
		EmailID:    input.EmailID,
		Password:   input.Password,
		SignUpType: input.SignUpType,
		CreatedAt:  time.Now().UTC(),
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Status:     input.Status,
	}

	insertOneUserResult, err := userCollection.InsertOne(ctx, createUserInput)

	if err != nil {
		log.Fatal("Error occurred while creating a new user in mongo", err)
	}

	fmt.Println("createdUser", insertOneUserResult)

	return bson.M{"emailID": emailId}
}
