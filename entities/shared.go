package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
}

type Link struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
	Rel  string             `json:"rel,omitempty" bson:"rel,omitempty"`
	Href int                `json:"href,omitempty" bson:"href,omitempty"`
}

type WebImage struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"image_name,omitempty" bson:"image_name,omitempty"`
	Extension string             `json:"extension,omitempty" bson:"extension,omitempty"`
	Src       string             `json:"src,omitempty" bson:"src,omitempty"`
	Rel       string             `json:"rel,omitempty" bson:"rel,omitempty"`
	Created   time.Time          `json:"created,omitempty" bson:"created,omitempty"`
}

type Contact struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Lastname string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone    string             `json:"phone,omitempty" bson:"phone,omitempty"`
}

type Address struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Address   string             `json:"address,omitempty" bson:"address,omitempty"`
	Aditional string             `json:"aditional,omitempty" bson:"aditional,omitempty"`
	City      string             `json:"city,omitempty" bson:"city,omitempty"`
	Region    string             `json:"region,omitempty" bson:"region,omitempty"`
	Country   string             `json:"country,omitempty" bson:"country,omitempty"`
	Postal    string             `json:"postal,omitempty" bson:"postal,omitempty"`
}
type BillingInfo struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	LastName string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Company  string             `json:"company,omitempty" bson:"company,omitempty"`
	Contact  Contact            `json:"contact,omitempty" bson:"contact,omitempty"`
	Address  Address            `json:"address,omitempty" bson:"address,omitempty"`
}
