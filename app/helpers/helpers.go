package helpers

import (
	"bgelato/app/secret"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(userId primitive.ObjectID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	role, _ := FindRole(userId)
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expira en 24 horas

	tokenString, err := token.SignedString(secret.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func FindRole(userId primitive.ObjectID) (string, error) {
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
