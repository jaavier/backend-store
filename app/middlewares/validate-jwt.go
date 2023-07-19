package middlewares

import (
	"bgelato/app/secret"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		// Verifica si el token está presente
		if tokenString == "" {
			return echo.ErrUnauthorized
		}

		// Parsea y verifica el token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verifica el algoritmo de firma
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("algoritmo de firma inválido")
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return echo.ErrUnauthorized, nil
			}

			if claims["userId"] != nil {
				userId := claims["userId"].(string)
				c.Set("userId", userId)
				return secret.JwtKey, nil
			} else {
				return echo.ErrUnauthorized, nil
			}

		})

		if err != nil || !token.Valid {
			return echo.ErrUnauthorized
		}

		// Si el token es válido, continua con el siguiente manejador
		return next(c)
	}
}
