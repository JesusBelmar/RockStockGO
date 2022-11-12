package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Audit struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Created      time.Time          `json:"created,omitempty" bson:"created,omitempty"`
	CreationUser string             `json:"creation_user,omitempty" bson:"creation_user,omitempty"`
	Modified     string             `json:"modified,omitempty" bson:"modified,omitempty"`
	ModifiedUser time.Time          `json:"modified_user,omitempty" bson:"modified_user,omitempty"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty"`
}
