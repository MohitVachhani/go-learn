package resourceinterface

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetResourceInputInterface struct {
	ResourceId string `json:"resourceId"`
}

type GetResourcePayloadInterface struct {
	Resource *ResourceSchema `json:"resource"`
}

type ResourceSchema struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"firstName,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time         `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Status    string             `json:"status" bson:"status,omitempty"`
}
