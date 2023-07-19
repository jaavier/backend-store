package handlers

import (
	"bgelato/app/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetOrdersKitchen(c echo.Context) error {
	if result, err := services.GetOrdersKitchen(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func MarkOrderAs(c echo.Context) error {
	orderId := c.Param("orderId")
	var body = make(map[string]string)
	json.NewDecoder(c.Request().Body).Decode(&body)
	err := services.MarkOrderAs(orderId, body["status"])
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Order %s marked as ready successfully", orderId),
	})
}
