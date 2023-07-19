package services

import (
	"bgelato/app/models"
	"bgelato/db"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CalculatePrice(productsIds []string) float64 {
	var total float64
	for _, productId := range productsIds {
		var product models.Product
		var id, _ = primitive.ObjectIDFromHex(productId)
		result := db.Products.FindOne(context.TODO(), bson.M{
			"_id": id,
		})
		if err := result.Decode(&product); err != nil {
			fmt.Println("Error getting price for productId", productId)
		} else {
			total += product.Price
		}
	}
	return total
}

func CreateOrder(productsIds []string) (string, error) {
	generateId := uuid.NewString()
	newOrder := models.Order{
		Products: productsIds,
		Date:     time.Now(),
		Id:       generateId,
		Total:    CalculatePrice(productsIds),
	}
	_, err := db.Orders.InsertOne(context.TODO(), newOrder)
	if err != nil {
		fmt.Println("Error creating order:", err)
		return "", err
	}
	return generateId, nil
}

func ViewOrder(orderId string) (models.Order, error) {
	var order models.Order
	result := db.Orders.FindOne(context.TODO(), bson.M{
		"orderId": orderId,
	})
	result.Decode(&order)
	if len(order.Products) == 0 {
		fmt.Println("[Find Orders] Order not found")
		return models.Order{}, fmt.Errorf("order not found")
	}
	return order, nil
}

func ViewAllOrders() ([]models.Order, error) {
	var order []models.Order
	results, err := db.Orders.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("[Find Orders] Error while trying to find order")
		return []models.Order{}, fmt.Errorf("error while trying to find order")
	}
	results.All(context.TODO(), &order)
	return order, nil
}
