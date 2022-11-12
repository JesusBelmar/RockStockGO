package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	. "github.com/rockstock-go-api/config"
	. "github.com/rockstock-go-api/entities"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProduct(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	fmt.Println("Obtenido producto: ", id)
	var product Product
	collection, ctx := GetCollection("product")
	err := collection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&product)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting product": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(product)
}

func GetProducts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var products []Product

	collection, ctx := GetCollection("product")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting products ": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product Product
		cursor.Decode(&product)
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting products": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(products)
}

func CreateProduct(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var product Product
	json.NewDecoder(request.Body).Decode(&product)

	product.Created = time.Now()
	product.Updated = time.Now()

	collection, ctx := GetCollection("product")

	result, _ := collection.InsertOne(ctx, product)
	json.NewEncoder(response).Encode(result)
}

func UpdateProduct(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	var product Product
	_ = json.NewDecoder(request.Body).Decode(&product)

	collection, ctx := GetCollection("product")

	oid, _ := primitive.ObjectIDFromHex(product.ID.Hex())
	filter := bson.D{primitive.E{Key: "_id", Value: oid}}
	update := bson.M{"$set": bson.M{
		"name":        product.Name,
		"description": product.Description,
		"category":    product.Category,
	}}
	fmt.Print("Actualizando producto: ", product.ID)
	result, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error updating product": "` + err.Error() + `" }`))
	}
	if result.ModifiedCount > 0 {
		fmt.Println(", actualizado")
	} else {
		fmt.Println(", NO actualizado")
	}

	json.NewEncoder(response).Encode(result)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	opts := options.Delete().SetCollation(&options.Collation{})
	collection, _ := GetCollection("product")
	res, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v document\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted
}
