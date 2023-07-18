package main

import (
	"bgelato/app/handlers"
	"bgelato/app/middlewares"

	"github.com/labstack/echo/v4"
)

func main() {
	var e = echo.New()

	e.GET("/", handlers.PrivateHandler, middlewares.JwtMiddleware)

	e.POST("/register", handlers.RegisterUser)
	e.POST("/login", handlers.LoginUser)

	e.Logger.Fatal(e.Start(":3000"))
}
