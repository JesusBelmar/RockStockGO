package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/rockstock-go-api/apis"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	fmt.Println("Starting the application...")

	router := mux.NewRouter()

	router.HandleFunc("/rockstock/auth/login", Login).Methods("POST")

	router.HandleFunc("/rockstock/user/create", CreateUser).Methods("POST")
	router.HandleFunc("/rockstock/user/list", GetUsers).Methods("GET")
	router.HandleFunc("/rockstock/user/get", GetUser).Methods("GET")
	router.HandleFunc("/rockstock/user/delete/{id}", DeleteUser).Methods("DELETE")

	router.HandleFunc("/rockstock/invoice/create", CreateInvoice).Methods("POST")
	router.HandleFunc("/rockstock/invoice/get", GetInvoice).Methods("GET")
	router.HandleFunc("/rockstock/invoice/list", GetInvoices).Methods("GET")

	router.HandleFunc("/rockstock/product/create", CreateProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/rockstock/product/get/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/rockstock/product/list", GetProducts).Methods("GET")
	router.HandleFunc("/rockstock/product/update", UpdateProduct).Methods("PATCH")
	router.HandleFunc("/rockstock/product/delete/{id}", DeleteProduct).Methods("DELETE")

	router.HandleFunc("/rockstock/category/create", CreateCategory).Methods("POST", "OPTIONS")
	router.HandleFunc("/rockstock/category/get", GetCategory).Methods("GET")
	router.HandleFunc("/rockstock/category/list", GetCategories).Methods("GET")
	router.HandleFunc("/rockstock/category/delete/{id}", DeleteCategory).Methods("DELETE")

	router.HandleFunc("/rockstock/provider/create", CreateProvider).Methods("POST", "OPTIONS")
	router.HandleFunc("/rockstock/provider/get", GetProvider).Methods("GET")
	router.HandleFunc("/rockstock/provider/list", GetProviders).Methods("GET")
	router.HandleFunc("/rockstock/provider/delete/{id}", DeleteProvider).Methods("DELETE")

	router.HandleFunc("/rockstock/customer/create", CreateCustomer).Methods("POST")
	router.HandleFunc("/rockstock/customer/list", GetCustomers).Methods("GET")
	router.HandleFunc("/rockstock/customer/get", GetCustomer).Methods("GET")
	router.HandleFunc("/rockstock/customer/delete/{id}", DeleteCustomer).Methods("DELETE")

	router.HandleFunc("/rockstock/employee/create", CreateEmployee).Methods("POST")
	router.HandleFunc("/rockstock/employee/list", GetEmployees).Methods("GET")
	router.HandleFunc("/rockstock/employee/get", GetEmployee).Methods("GET")
	router.HandleFunc("/rockstock/employee/delete/{id}", DeleteEmployee).Methods("DELETE")

	router.HandleFunc("/rockstock/rol/create", CreateRol).Methods("POST")
	router.HandleFunc("/rockstock/rol/list", GetRoles).Methods("GET")
	router.HandleFunc("/rockstock/rol/get", GetRol).Methods("GET")
	router.HandleFunc("/rockstock/rol/delete/{id}", DeleteRol).Methods("DELETE")

	router.HandleFunc("/rockstock/brand/create", CreateBrand).Methods("POST")
	router.HandleFunc("/rockstock/brand/list", GetBrand).Methods("GET")

	router.HandleFunc("/rockstock/country/list", GetCountries).Methods("GET")
	router.HandleFunc("/rockstock/country/get", GetCountry).Methods("GET")

	router.HandleFunc("/rockstock/state/list/{id}", GetStates).Methods("GET")
	router.HandleFunc("/rockstock/state/get", GetState).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router)) // Run Server
}
