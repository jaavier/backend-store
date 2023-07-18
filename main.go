package main

import (
	"bgelato/middlewares"
	"bgelato/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	var e = echo.New()

	e.GET("/", func(c echo.Context) error {
		username := c.Get("username")
		return c.JSON(200, map[string]string{"nickname": fmt.Sprint(username)})
	}, middlewares.JwtMiddleware)

	e.POST("/register", func(c echo.Context) error {
		var user models.User
		if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
			return c.JSON(400, map[string]string{
				"message": "Cannot create user",
			})
		}
		if userCreated, err := models.CreateUser(user.Nickname, user.Passwd); err != nil {
			fmt.Println("Error creating user:", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		} else {
			fmt.Println(userCreated)
			return c.JSON(http.StatusOK, map[string]string{"userId": userCreated.Id})
		}
	})

	e.POST("/login", func(c echo.Context) error {
		var body models.User
		err := json.NewDecoder(c.Request().Body).Decode(&body)
		if err != nil {
			fmt.Println("Error while decoding body /login")
			return c.JSON(400, map[string]string{"message": "Not valid"})
		}
		if body.Nickname == "root" && body.Passwd == "toor" {
			token, err := generateToken(body.Nickname)
			if err != nil {
				fmt.Println("Error while generating JWT token:", err)
				return c.JSON(400, map[string]string{"message": "Not valid"})
			}
			return c.JSON(200, map[string]string{"token": token})
		}
		return c.JSON(200, body)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
