package resourcerepo

import (
	"context"
	"log"
	"time"

	ResourceInterface "github.com/MohitVachhani/go-learn/pkg/structs/resource"
	mongoUtils "github.com/MohitVachhani/go-learn/pkg/utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func GetResourceById(id string) ResourceInterface.ResourceSchema {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	resourceCollection := mongoUtils.GetCollection(mongoUtils.MongoClient, "resources")

	var resource ResourceInterface.ResourceSchema

	err := resourceCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&resource)

	if err != nil {
		log.Fatal(err)
	}

	return resource
}
