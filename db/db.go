package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionsStruct struct {
	Products *mongo.Collection
	Users    *mongo.Collection
	Orders   *mongo.Collection
}

var MongoClient *mongo.Client
var clientOnce sync.Once
var Collections CollectionsStruct

func Connect() {
	// clientOnce.Do(func() {
	host := os.Getenv("HOST")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")
	connectionStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s", username, password, host, database)
	clientOptions := options.Client().ApplyURI(connectionStr)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	MongoClient = client
	Collections.Products = client.Database("slider").Collection("products")
	Collections.Users = client.Database("slider").Collection("users")
	Collections.Orders = client.Database("slider").Collection("orders")
	fmt.Printf("[SUCCESS] Connected successfully to MongoDB %s\n", host)
	// defer client.Disconnect(context.Background())
	// })
}
