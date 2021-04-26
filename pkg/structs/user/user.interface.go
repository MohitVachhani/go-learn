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
	EmailID    string `json:"emailId"`
	Password   string `json:"password"`
	SignUpType string `json:"signUpType"`
}

type User struct {
	ID         primitive.ObjectID  `json:"_id" bson:"_id,omitempty"`
	EmailID    string              `json:"emailId" bson:"emailId"`
	Password   string              `json:"password"`
	SignUpType string              `json:"signUpType" bson:"signUpType"`
	CreatedAt  time.Time           `json:"createdAt" bson:"createdAt"`
	LastLogin  primitive.Timestamp `json:"lastLogin" bson:"lastLogin,omitempty"`
	UpdatedAt  primitive.Timestamp `json:"updatedAt" bson:"updatedAt,omitempty"`
}
