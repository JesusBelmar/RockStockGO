package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/rockstock-go-api/config"
	. "github.com/rockstock-go-api/entities"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetState(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["country_id"])
	var state CountryState
	collection, ctx := GetCollection("country_states")
	err := collection.FindOne(ctx, CountryState{ID: id}).Decode(&state)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting state": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(state)
}

func GetStates(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var states []CountryState
	params := mux.Vars(request)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	collection, ctx := GetCollection("country_states")

	cursor, err := collection.Find(ctx, bson.D{primitive.E{Key: "_id", Value: _id}})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents ": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var state CountryState
		cursor.Decode(&state)
		states = append(states, state)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(states)
}
