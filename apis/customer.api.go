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

func GetCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var customer Customer
	collection, ctx := GetCollection("product")
	err := collection.FindOne(ctx, Customer{ID: id}).Decode(&customer)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting customer": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(customer)
}

func CreateCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var customer Customer
	json.NewDecoder(request.Body).Decode(&customer)
	collection, ctx := GetCollection("product")
	result, _ := collection.InsertOne(ctx, customer)
	json.NewEncoder(response).Encode(result)
}

func GetCustomers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var customers []Customer

	collection, ctx := GetCollection("customer")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting customers": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var customer Customer
		cursor.Decode(&customer)
		customers = append(customers, customer)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting customers": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(customers)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		Name string `json:"name"` //value that has to be matched
		City string `json:"city"` // value that has to be modified
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {
		fmt.Print(e)
	}
	filter := bson.D{{"name", body.Name}} // converting value to BSON type
	after := options.After                // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{{"set", bson.D{{"city", body.City}}}}

	collection, ctx := GetCollection("customer")
	updateResult := collection.FindOneAndUpdate(ctx, filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

//Delete Profile of Customer

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}

	opts := options.Delete().SetCollation(&options.Collation{})
	collection, _ := GetCollection("customer")
	res, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}
