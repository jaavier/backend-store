package services

import (
	"bgelato/app/db"
	"bgelato/app/helpers"
	"bgelato/app/models"
	"fmt"

	"github.com/google/uuid"
)

func ValidateLogin(user models.User) (string, error) {
	if user.Nickname == "root" && user.Passwd == "toor" {
		token, err := helpers.GenerateToken(user.Nickname)
		if err != nil {
			fmt.Println("Error while generating JWT token:", err)
			return "", fmt.Errorf("Not valid")
		}
		return token, nil
	}
	return "", fmt.Errorf("Invalid Credentials")
}

func CreateUser(nickname, passwd string) (models.User, error) {
	if len(nickname) < 3 {
		return models.User{}, fmt.Errorf("invalid nickname")
	}

	if len(passwd) < 6 {
		return models.User{}, fmt.Errorf("password length must be > 6")
	}
	var uniqueId = uuid.NewString()
	db.Users[uniqueId] = models.User{Id: uniqueId, Nickname: nickname, Passwd: passwd}
	return db.Users[uniqueId], nil
}
