package helpers

import (
	"bgelato/app/secret"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	role, _ := FindRole(username)
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expira en 24 horas

	tokenString, err := token.SignedString(secret.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func FindRole(username string) (string, error) {
	return "user", nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidatePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
