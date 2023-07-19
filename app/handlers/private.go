package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func PrivateHandler(c echo.Context) error {
	userId := c.Get("userId")
	return c.JSON(200, map[string]string{"userId": fmt.Sprint(userId)})
}
