package handlers

import (
	"bgelato/app/models"
	"bgelato/app/services"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Filter struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Date        int    `json:"date,omitempty" bson:"date,omitempty"`
	Flavor      string `json:"flavor,omitempty" bson:"flavor,omitempty"`
}

func createFilter(values url.Values) bson.M {
	filters := []string{"name", "description", "date", "flavor", "storeId", "stock"}
	var results = bson.M{}

	for _, filter := range filters {
		value := values.Get(filter)
		if len(value) > 0 {
			if _, ok := results[filter]; !ok {
				var finalValue interface{}
				if filter == "storeId" || filter == "stock" {
					toInt, err := strconv.Atoi(value)
					if err != nil {
						fmt.Println("error converting storeId to int")
						break
					}
					finalValue = toInt
				}
				results[filter] = finalValue
			}
		}
	}

	return results
}

func LoadProducts(c echo.Context) error {
	results := createFilter(c.Request().URL.Query())
	result, err := services.LoadProducts(results)
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
