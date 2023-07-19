package handlers

import (
	"bgelato/app/models"
	"bgelato/app/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(c echo.Context) error {
	var newOrder models.Order
	err := json.NewDecoder(c.Request().Body).Decode(&newOrder)
	if err != nil {
		fmt.Println("[Decode Payload] Error creating order", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error creating order",
		})
	}
	userId := fmt.Sprint(c.Get("userId"))
	objectId, _ := primitive.ObjectIDFromHex(userId)
	if orderId, err := services.CreateOrder(newOrder.Products, objectId); err != nil {
		fmt.Println("[Insert Order] Error creating order", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error creating order",
		})
	} else {
		return c.JSON(http.StatusOK, map[string]string{
			"orderId": orderId,
		})
	}
}

func ViewOrder(c echo.Context) error {
	orderId := c.Param("orderId")
	result, err := services.ViewOrder(orderId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func ViewAllOrders(c echo.Context) error {
	userId := c.Get("userId").(string)
	fmt.Println("UserId:", userId)
	objectId, _ := primitive.ObjectIDFromHex(userId)
	result, err := services.ViewAllOrders(objectId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func PayOrder(c echo.Context) error {
	orderId := c.Param("orderId")
	userId := c.Get("userId")
	if len(orderId) < 5 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid orderId",
		})
	}
	var objectId, _ = primitive.ObjectIDFromHex(fmt.Sprint(userId))
	if err := services.PayOrder(orderId, objectId); err != nil {
		fmt.Printf("[CRITICAL] Error paying order %s: %s\n", orderId, err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	fmt.Printf("[SUCCESS] Order %s payed successfully\n", orderId)
	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Order %s payed successfully", orderId),
	})
}
