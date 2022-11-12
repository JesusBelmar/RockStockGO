package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	DNI       string             `json:"dni,omitempty" bson:"dni,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Gender    string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Birthday  string             `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Login     string             `json:"login,omitempty" bson:"login,omitempty"`
	Pass      string             `json:"credential,omitempty" bson:"credential,omitempty"`
	Avatar    string             `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Rol       Rol                `json:"rol,omitempty" bson:"rol,omitempty"`
	Contacts  Contact            `json:"contact,omitempty" bson:"contact,omitempty"`
	Addresses Address            `json:"address,omitempty" bson:"address,omitempty"`
	LastLogin string             `json:"last_login,omitempty" bson:"last_login,omitempty"`
	Created   string             `json:"created,omitempty" bson:"created,omitempty"`
	Updated   string             `json:"updated,omitempty" bson:"updated,omitempty"`
	Link      string             `json:"link,omitempty" bson:"link,omitempty"`
	Status    string             `json:"status,omitempty" bson:"status,omitempty"`
}
