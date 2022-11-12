package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rol struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Status string             `json:"status,omitempty" bson:"status,omitempty"`
}
