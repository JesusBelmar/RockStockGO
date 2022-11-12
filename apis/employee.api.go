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

func GetEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var employee Employee
	collection, ctx := GetCollection("employee")
	err := collection.FindOne(ctx, Employee{ID: id}).Decode(&employee)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting employee": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(employee)
}

func GetEmployees(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var employees []Employee

	collection, ctx := GetCollection("employee")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents ": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var employee Employee
		cursor.Decode(&employee)
		employees = append(employees, employee)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting documents": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(employees)
}

func CreateEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var employee Employee
	json.NewDecoder(request.Body).Decode(&employee)

	collection, ctx := GetCollection("employee")

	result, _ := collection.InsertOne(ctx, employee)
	json.NewEncoder(response).Encode(result)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{})
	collection, _ := GetCollection("employee")
	res, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v document\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted
}
