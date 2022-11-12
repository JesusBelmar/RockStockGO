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

func GetOrder(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var order Order
	collection, ctx := GetCollection("order")
	err := collection.FindOne(ctx, Order{ID: id}).Decode(&order)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting order": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(order)
}

func GetOrders(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var orders []Order

	collection, ctx := GetCollection("order")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents ": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var order Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(orders)
}

func CreateOrder(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var order Order
	json.NewDecoder(request.Body).Decode(&order)

	collection, ctx := GetCollection("order")

	result, _ := collection.InsertOne(ctx, order)
	json.NewEncoder(response).Encode(result)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{})
	collection, _ := GetCollection("order")
	res, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v document\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted
}
