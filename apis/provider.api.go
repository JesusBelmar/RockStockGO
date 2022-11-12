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

func GetProvider(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var provider Seller
	collection, ctx := GetCollection("provider")
	err := collection.FindOne(ctx, Seller{ID: id}).Decode(&provider)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting provider": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(provider)
}

func GetProviders(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var providers []Seller

	collection, ctx := GetCollection("provider")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents ": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var provider Seller
		cursor.Decode(&provider)
		providers = append(providers, provider)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(providers)
}

func CreateProvider(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var provider Seller
	json.NewDecoder(request.Body).Decode(&provider)

	collection, ctx := GetCollection("provider")

	result, _ := collection.InsertOne(ctx, provider)
	json.NewEncoder(response).Encode(result)
}

func DeleteProvider(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{})
	collection, _ := GetCollection("provider")
	res, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v document\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted
}
