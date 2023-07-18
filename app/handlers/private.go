package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func PrivateHandler(c echo.Context) error {
	username := c.Get("username")
	return c.JSON(200, map[string]string{"nickname": fmt.Sprint(username)})
}
