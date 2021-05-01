package userinterface

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserFilters interface
type UserFilters struct {
	EmailID string `json:"emailId"`
}

type CreateUserInput struct {
	EmailID        string `json:"emailId"`
	Password       string `json:"password" bson:"password,omitempty"`
	SignUpType     string `json:"signUpType"`
	FirstName      string `json:"firstName" bson:"firstName,omitempty"`
	LastName       string `json:"lastName" bson:"lastName,omitempty"`
	ProfilePicture string `json:"profilePicture" bson:"profilePicture,omitempty"`
	Status         string `json:"status" bson:"status,omitempty"`
}

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	EmailID        string             `json:"emailId" bson:"emailId"`
	FirstName      string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName       string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Password       string             `json:"password" bson:"password,omitempty"`
	SignUpType     string             `json:"signUpType" bson:"signUpType"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	LastLogin      *time.Time         `json:"lastLogin,omitempty" bson:"lastLogin,omitempty"`
	UpdatedAt      *time.Time         `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	ProfilePicture string             `json:"profilePicture,omitempty" bson:"profilePicture,omitempty"`
	Status         string             `json:"status" bson:"status,omitempty"`
}
