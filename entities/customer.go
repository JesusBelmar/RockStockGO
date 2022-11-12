package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	DNI         string             `json:"dni,omitempty" bson:"dni,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Avatar      WebImage           `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Rol         Rol                `json:"rol,omitempty" bson:"rol,omitempty"`
	Login       string             `json:"login,omitempty" bson:"login,omitempty"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Address     Address            `json:"address,omitempty" bson:"address,omitempty"`
	BillingInfo []BillingInfo      `json:"billing_info,omitempty" bson:"billing_info,omitempty"`
	Contact     Contact            `json:"contact,omitempty" bson:"contact,omitempty"`
	Contacts    []Contact          `json:"contacts,omitempty" bson:"contacts,omitempty"`
	LastLogin   string             `json:"last_login,omitempty" bson:"last_login,omitempty"`
	Created     time.Time          `json:"created,omitempty" bson:"created,omitempty"`
	Updated     time.Time          `json:"updated,omitempty" bson:"updated,omitempty"`
	Link        string             `json:"link,omitempty" bson:"link,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
}
