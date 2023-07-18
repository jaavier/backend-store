package handlers

import (
	"bgelato/app/models"
	"bgelato/app/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoadProducts(c echo.Context) error {
	result, err := services.LoadProducts()
	if err != nil {
		fmt.Println("ERROR:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error loading products",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"products": result,
	})
}

func InsertProduct(c echo.Context) error {
	var product models.Product
	if err := json.NewDecoder(c.Request().Body).Decode(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error in payload",
		})
	}
	result, err := services.InsertProduct(product)
	if err != nil {
		fmt.Println("ERROR:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Error loading products",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"products": result,
	})
}
