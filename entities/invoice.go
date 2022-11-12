package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LineItem struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ProductId   int                `json:"product_id,omitempty" bson:"product_idid,omitempty"`
	ProductName string             `json:"product_name,omitempty" bson:"product_name,omitempty"`
	Discount    string             `json:"discount,omitempty" bson:"discount,omitempty"`
	Value       string             `json:"value,omitempty" bson:"value,omitempty"`
	Quantity    string             `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

type Invoice struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ClientId     int32              `json:"client_id,omitempty" bson:"client_id,omitempty"`
	ClientName   string             `json:"client_name,omitempty" bson:"client_name,omitempty"`
	Quantity     int16              `json:"quantity,omitempty" bson:"quantity,omitempty"`
	SubTotal     float32            `json:"subtotal,omitempty" bson:"subtotal,omitempty"`
	StartDate    time.Time          `json:"startdate,omitempty" bson:"startdate,omitempty"`
	EndDate      time.Time          `json:"endate,omitempty" bson:"endate,omitempty"`
	AmountBilled float32            `json:"amount,omitempty" bson:"amount,omitempty"`
	AmountPaid   float32            `json:"paid,omitempty" bson:"paid,omitempty"`
	SalesTax     float32            `json:"tax,omitempty" bson:"tax,omitempty"`
	LineItems    []LineItem         `json:"items,omitempty" bson:"items,omitempty"`
	Links        []Link             `json:"links,omitempty" bson:"links,omitempty"`
	Created      time.Time          `json:"created,omitempty" bson:"created,omitempty"`
	Updated      time.Time          `json:"updated,omitempty" bson:"updated,omitempty"`
	StatusName   string             `json:"status,omitempty" bson:"status,omitempty"`
}
