package handlers

import (
	"bgelato/app/models"
	"bgelato/app/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
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
	if orderId, err := services.CreateOrder(newOrder.Products); err != nil {
		fmt.Println("[Insert Order] Error creating order", err)
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
	result, err := services.ViewAllOrders()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}
