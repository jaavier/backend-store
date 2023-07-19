package handlers

import (
	"bgelato/app/models"
	"bgelato/app/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	var user models.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return c.JSON(400, map[string]string{
			"message": "Cannot create user",
		})
	}
	newUser, err := services.CreateUser(user.Nickname, user.Password)
	if err != nil {
		return c.JSON(400, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(200, newUser)
}

func LoginUser(c echo.Context) error {
	var body models.User

	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		fmt.Println("Error while decoding body /login")
		return c.JSON(400, map[string]string{"message": "Not valid"})
	}

	result, err := services.ValidateLogin(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": result})
}
