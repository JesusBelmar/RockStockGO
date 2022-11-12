package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	No          primitive.ObjectID `json:"no,omitempty" bson:"no,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Desc        string             `json:"desc,omitempty" bson:"desc,omitempty"`
	ETA         time.Time          `json:"eta,omitempty" bson:"eta,omitempty"`
	AutorizedBy string             `json:"autorizedby,omitempty" bson:"autorizedby,omitempty"`
	Autorized   string             `json:"autorized,omitempty" bson:"autorized,omitempty"`
	Created     time.Time          `json:"created,omitempty" bson:"created,omitempty"`
	Updated     time.Time          `json:"updated,omitempty" bson:"updated,omitempty"`
	Link        string             `json:"link,omitempty" bson:"link,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	Employee    []Employee         `json:"employee,omitempty" bson:"employee,omitempty"`
	Items       []Item             `json:"item,omitempty" bson:"item,omitempty"`
	Materials   []Material         `json:"materials,omitempty" bson:"materials,omitempty"`
}

type Item struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Desc     string             `json:"desc,omitempty" bson:"desc,omitempty"`
	Time     float32            `json:"time,omitempty" bson:"time,omitempty"`
	Quantity float32            `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Budget   float32            `json:"budget,omitempty" bson:"budget,omitempty"`
	Created  time.Time          `json:"created,omitempty" bson:"created,omitempty"`
	Updated  time.Time          `json:"updated,omitempty" bson:"updated,omitempty"`
	Link     string             `json:"link,omitempty" bson:"link,omitempty"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
}

type Material struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Desc     string             `json:"desc,omitempty" bson:"desc,omitempty"`
	Time     float32            `json:"time,omitempty" bson:"time,omitempty"`
	Quantity float32            `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Budget   float32            `json:"budget,omitempty" bson:"budget,omitempty"`
	Created  time.Time          `json:"created,omitempty" bson:"created,omitempty"`
	Updated  time.Time          `json:"updated,omitempty" bson:"updated,omitempty"`
	Link     string             `json:"link,omitempty" bson:"link,omitempty"`
	Status   string             `json:"status,omitempty" bson:"status,omitempty"`
}
