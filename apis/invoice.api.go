package apis

import (
	"encoding/json"
	"net/http"

	. "github.com/rockstock-go-api/config"
	. "github.com/rockstock-go-api/entities"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetInvoice(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var invoice Invoice
	collection, ctx := GetCollection("invoice")
	err := collection.FindOne(ctx, Invoice{ID: id}).Decode(&invoice)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting invoice": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(invoice)
}

func GetInvoices(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var invoices []Invoice
	collection, ctx := GetCollection("invoice")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting invoices": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var invoice Invoice
		cursor.Decode(&invoice)
		invoices = append(invoices, invoice)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "Error getting invoices": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(invoices)
}

func CreateInvoice(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var invoice Invoice
	json.NewDecoder(request.Body).Decode(&invoice)
	collection, ctx := GetCollection("invoice")
	result, _ := collection.InsertOne(ctx, invoice)
	json.NewEncoder(response).Encode(result)
}
