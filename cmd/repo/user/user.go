package user

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//GetUserByEmailID is
func GetUserByEmailID(ctx context.Context, userCollection *mongo.Collection, emailID string) bson.M {

	fmt.Println("Fetching user with emailId:", emailID)

	var user bson.M

	err := userCollection.FindOne(ctx, bson.M{"emailId": emailID}).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	return user
}
