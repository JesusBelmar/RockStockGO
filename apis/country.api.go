package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "github.com/rockstock-go-api/config"
	. "github.com/rockstock-go-api/entities"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCountry(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var country Country
	collection, ctx := GetCollection("countries")
	err := collection.FindOne(ctx, Country{ID: id}).Decode(&country)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting country": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(country)
}

func GetCountries(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var countries []Country

	collection, ctx := GetCollection("countries")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents ": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var country Country
		cursor.Decode(&country)
		countries = append(countries, country)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(countries)
}

func CreateCountry(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var country Country
	json.NewDecoder(request.Body).Decode(&country)

	collection, ctx := GetCollection("countries")

	result, _ := collection.InsertOne(ctx, country)
	json.NewEncoder(response).Encode(result)
}

func DeleteCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{})
	collection, _ := GetCollection("countries")
	res, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v document\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted
}
