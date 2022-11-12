package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Brand struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CountryId string             `json:"brand_id,omitempty" bson:"brand_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Log       Audit              `json:"log,omitempty" bson:"log,omitempty"`
}
