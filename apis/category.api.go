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

func GetCategory(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var category Category
	collection, ctx := GetCollection("category")
	err := collection.FindOne(ctx, Category{ID: id}).Decode(&category)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting category": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(category)
}

func GetCategories(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var categories []Category

	collection, ctx := GetCollection("category")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents ": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var category Category
		cursor.Decode(&category)
		categories = append(categories, category)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(categories)
}

func CreateCategory(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var category Category
	json.NewDecoder(request.Body).Decode(&category)

	collection, ctx := GetCollection("category")

	result, _ := collection.InsertOne(ctx, category)
	json.NewEncoder(response).Encode(result)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{})
	collection, _ := GetCollection("category")
	res, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v document\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted
}
