package helpers

import (
	"bgelato/app/secret"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expira en 24 horas

	tokenString, err := token.SignedString(secret.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
