package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDB() *mongo.Client {
	host := "mongodb://rockstock:rockstock.123@rockstock-mongo:27017"
	clientOptions := options.Client().ApplyURI(host)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetCollection(name string) (*mongo.Collection, context.Context) {
	var dbName = "RockStock"
	var collection = GetMongoDB().Database(dbName).Collection(name)
	return collection, context.TODO()
}

func GetCollectionWithoutContext(name string) *mongo.Collection {
	var dbName = "RockStock"
	var collection = GetMongoDB().Database(dbName).Collection(name)
	return collection
}
