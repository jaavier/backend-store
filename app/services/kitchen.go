package services

import (
	"bgelato/app/models"
	"bgelato/db"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func ValidateStep(currentStep string) bool {
	var validSteps = []string{"pay", "preparing", "ready", "delivery", "received"}
	var found bool = false
	for _, step := range validSteps {
		if step == currentStep {
			found = true
		}
	}
	return found
}

func GetOrdersKitchen() ([]models.Order, error) {
	var orders []models.Order
	pendingOrders, err := db.Orders.Find(context.TODO(), bson.M{
		"step": "preparing",
	})
	if err != nil {
		return []models.Order{}, fmt.Errorf("cannot get all orders pending")
	}
	if err := pendingOrders.All(context.TODO(), &orders); err != nil {
		fmt.Println("[CRITICAL] Check GetOrdersKitchen")
		return []models.Order{}, fmt.Errorf("cannot get all orders pending")
	}
	return orders, nil
}

func MarkOrderAs(orderId, step string) error {
	if !ValidateStep(step) {
		return fmt.Errorf("%s is not a valid step", step)
	}
	_, err := db.Orders.UpdateOne(context.TODO(), bson.M{
		"orderId": orderId,
	}, bson.M{
		"$set": bson.M{
			"step": step,
		},
	})
	if err != nil {
		fmt.Printf("[CRITICAL] Error while marking order %s as ready\n", orderId)
		return err
	}

	return nil
}
