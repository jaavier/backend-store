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

func calculatePrice(productsIds []primitive.ObjectID) float64 {
	var total float64
	var products []models.Product
	cursor, _ := db.Products.Find(context.TODO(), bson.M{
		"_id": bson.M{
			"$in": productsIds,
		},
	})
	if err := cursor.All(context.TODO(), &products); err != nil {
		fmt.Println("Error getting price for productId")
	} else {
		for _, product := range products {
			total += product.Price
		}
	}
	return total
}

func CreateOrder(productsIds []primitive.ObjectID, userId primitive.ObjectID) (string, error) {
	generateId := uuid.NewString()
	newOrder := models.Order{
		Products: productsIds,
		Date:     time.Now(),
		Id:       generateId,
		Total:    calculatePrice(productsIds),
		Step:     "pending",
		UserId:   userId,
	}
	_, err := db.Orders.InsertOne(context.TODO(), newOrder)
	if err != nil {
		fmt.Println("Error creating order:", err)
		return "", err
	}
	return generateId, nil
}

func PayOrder(orderId string, userId primitive.ObjectID) error {
	order, err := ViewOrder(orderId)
	if err != nil {
		return err
	}
	if order.UserId != userId {
		return fmt.Errorf("you don't own this order")
	}
	if order.Step != "pending" {
		return fmt.Errorf("order already paid or canceled")
	}
	updateOrder, err := db.Orders.UpdateOne(context.TODO(), bson.M{
		"orderId": orderId,
	}, bson.M{
		"$set": bson.M{
			"step": "preparing",
		},
	})
	if err != nil {
		return err
	}
	if updateOrder.ModifiedCount != 1 {
		fmt.Printf("[CRITICAL] Payed %d invoices at once\n", updateOrder.ModifiedCount)
		return fmt.Errorf("")
	}
	return nil
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

func ViewAllOrders(userId primitive.ObjectID) ([]models.Order, error) {
	var order []models.Order
	results, err := db.Orders.Find(context.TODO(), bson.M{
		"userId": userId,
	})
	if err != nil {
		fmt.Println("[Find Orders] Error while trying to find order")
		return []models.Order{}, fmt.Errorf("error while trying to find order")
	}
	results.All(context.TODO(), &order)
	return order, nil
}
