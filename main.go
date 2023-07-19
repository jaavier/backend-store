package main

import (
	"bgelato/app/handlers"
	"bgelato/app/middlewares"
	"bgelato/db"

	"github.com/labstack/echo/v4"
)

func main() {
	LoadEnv()
	db.Connect()

	var e = echo.New()

	e.GET("/", handlers.PrivateHandler, middlewares.JwtMiddleware)

	// Routes for products
	e.GET("/products", handlers.LoadProducts)
	e.POST("/products", handlers.InsertProduct, middlewares.JwtMiddleware)

	// Routes for orders
	e.GET("/orders/:orderId", handlers.ViewOrder, middlewares.JwtMiddleware)
	e.GET("/orders", handlers.ViewAllOrders, middlewares.JwtMiddleware)
	e.POST("/orders", handlers.CreateOrder, middlewares.JwtMiddleware)

	// Routes for users
	e.POST("/register", handlers.RegisterUser)
	e.POST("/login", handlers.LoginUser)

	e.Logger.Fatal(e.Start(":3000"))
}
