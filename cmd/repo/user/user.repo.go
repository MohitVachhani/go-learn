package user

import (
	"context"
	"log"

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

func GetUser(ctx context.Context, userFilters userInterface.UserFilters) userInterface.User {
	userCollection := mongoUtils.GetCollection(mongoUtils.MongoClient, "users")

	var user userInterface.User

	var emailID = userFilters.EmailID

	err := userCollection.FindOne(ctx, bson.M{"emailId": emailID}).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	return user
}
