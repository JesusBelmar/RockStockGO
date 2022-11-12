package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AttrubuteItem struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Key   string             `json:"key,omitempty" bson:"key,omitempty"`
	Value string             `json:"value,omitempty" bson:"value,omitempty"`
}

type Seller struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Address  string             `json:"address,omitempty" bson:"address,omitempty"`
	Phone    string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Contacts []Contact          `json:"contacts,omitempty" bson:"contacts,omitempty"`
}

type Product struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Seller        Seller             `json:"provider,omitempty" bson:"provider,omitempty"`
	Category      Category           `json:"category,omitempty" bson:"category,omitempty"`
	SKU           string             `json:"sku,omitempty" bson:"sku,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Description   string             `json:"description,omitempty" bson:"description,omitempty"`
	Image         WebImage           `json:"image,omitempty" bson:"image,omitempty"`
	Images        []WebImage         `json:"images,omitempty" bson:"images,omitempty"`
	Quantity      int16              `json:"quantity,omitempty" bson:"quantity,omitempty"`
	PurchasePrice float32            `json:"purchase_price,omitempty" bson:"purchase_price,omitempty"`
	SalePrice     float32            `json:"sale_price,omitempty" bson:"sale_price,omitempty"`
	Tax           float32            `json:"tax,omitempty" bson:"tax,omitempty"`
	Discount      float32            `json:"discount,omitempty" bson:"discount,omitempty"`
	Link          Link               `json:"link,omitempty" bson:"link,omitempty"`
	Attributes    []AttrubuteItem    `json:"attributes,omitempty" bson:"attributes,omitempty"`
	Created       time.Time          `json:"created,omitempty" bson:"created,omitempty"`
	Updated       time.Time          `json:"updated,omitempty" bson:"updated,omitempty"`
	LastUser      string             `json:"last_user,omitempty" bson:"last_user,omitempty"`
	StatusName    Status             `json:"status,omitempty" bson:"status,omitempty"`
}
