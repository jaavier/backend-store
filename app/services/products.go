package services

import (
	"bgelato/app/models"
	"bgelato/db"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertProduct(product models.Product) (models.Product, error) {
	_, err := db.Collections.Products.InsertOne(context.TODO(), product)
	if err != nil {
		return models.Product{}, err
	}
	fmt.Println("[SUCCESS] Product inserted successfully")
	return product, nil
}

func LoadProducts() ([]models.Product, error) {
	var productsList []models.Product
	cursor, err := db.Collections.Products.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			fmt.Println(err)
			return []models.Product{}, err
		}
		productsList = append(productsList, product)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}

	fmt.Println(productsList)
	return productsList, nil
}
