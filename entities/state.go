package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type CountryState struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CountryId  int32              `json:"country_id,omitempty" bson:"country_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Iso3       string             `json:"iso3,omitempty" bson:"iso3,omitempty"`
	Iso2       string             `json:"iso2,omitempty" bson:"iso2,omitempty"`
	Phone_code string             `json:"phone_code,omitempty" bson:"phone_code,omitempty"`
	Capital    string             `json:"capital,omitempty" bson:"capital,omitempty"`
	Currency   string             `json:"currency,omitempty" bson:"currency,omitempty"`
	Native     string             `json:"native,omitempty" bson:"native,omitempty"`
	Region     string             `json:"region,omitempty" bson:"region,omitempty"`
	Subregion  string             `json:"subregion,omitempty" bson:"subregion,omitempty"`
	Emoji      string             `json:"emoji,omitempty" bson:"emoji,omitempty"`
	EmojiU     string             `json:"emojiu,omitempty" bson:"emojiu,omitempty"`
	State      []State            `json:"states,omitempty" bson:"states,omitempty"`
}

type State struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	StateId   int32              `json:"state_id,omitempty" bson:"state_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	StateCode string             `json:"state_code,omitempty" bson:"state_code,omitempty"`
}
