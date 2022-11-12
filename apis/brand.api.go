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

func GetBrand(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var rol Brand
	collection, ctx := GetCollection("rol")
	err := collection.FindOne(ctx, Brand{ID: id}).Decode(&rol)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting rol": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(rol)
}

func CreateBrand(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var rol Brand
	json.NewDecoder(request.Body).Decode(&rol)
	collection, ctx := GetCollection("rol")
	result, _ := collection.InsertOne(ctx, rol)
	json.NewEncoder(response).Encode(result)
}

func GetBrands(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var roles []Brand

	collection, ctx := GetCollection("rol")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting roles": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var rol Brand
		cursor.Decode(&rol)
		roles = append(roles, rol)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting roles": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(roles)
}

func updateBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		Name string `json:"name"` //value that has to be matched
		City string `json:"city"` // value that has to be modified
	}
	var brand Brand
	e := json.NewDecoder(r.Body).Decode(&brand)
	if e != nil {
		fmt.Print(e)
	}
	//filter := bson.D{{"id", brand.ID}} // converting value to BSON type
	filter := bson.D{primitive.E{Key: "name", Value: brand.Name}}
	after := options.After // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update := bson.D{primitive.E{Key: "name", Value: brand.Name}}

	collection, ctx := GetCollection("rol")
	updateResult := collection.FindOneAndUpdate(ctx, filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

//Delete Profile of Brand

func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}

	opts := options.Delete().SetCollation(&options.Collation{})
	collection, _ := GetCollection("rol")
	res, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}
