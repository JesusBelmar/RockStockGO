package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"username,omitempty" bson:"username,omitempty"`
	DNI       string             `json:"dni,omitempty" bson:"dni,omitempty"`
	Pass      string             `json:"credential,omitempty" bson:"credential,omitempty"`
	Avatar    WebImage           `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Rol       Rol                `json:"rol,omitempty" bson:"rol,omitempty"`
	Login     string             `json:"login,omitempty" bson:"login,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	LastLogin string             `json:"last_login,omitempty" bson:"last_login,omitempty"`
	Created   time.Time          `json:"created,omitempty" bson:"created,omitempty"`
	Updated   time.Time          `json:"updated,omitempty" bson:"updated,omitempty"`
	Link      string             `json:"link,omitempty" bson:"link,omitempty"`
	Status    string             `json:"status,omitempty" bson:"status,omitempty"`
}
