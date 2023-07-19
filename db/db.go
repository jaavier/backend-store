package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Products *mongo.Collection
var Users *mongo.Collection
var Orders *mongo.Collection

var MongoClient *mongo.Client

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
	Products = client.Database("slider").Collection("products")
	Users = client.Database("slider").Collection("users")
	Orders = client.Database("slider").Collection("orders")
	fmt.Printf("[SUCCESS] Connected successfully to MongoDB %s\n", host)
	// defer client.Disconnect(context.Background())
	// })
}
