package services

import (
	"bgelato/app/helpers"
	"bgelato/app/models"
	"bgelato/db"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateLogin(user models.User) (string, error) {
	var userFound models.User
	var findUser = db.Users.FindOne(context.TODO(), bson.M{
		"nickname": user.Nickname,
	})
	err := findUser.Decode(&userFound)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	if helpers.ValidatePassword(userFound.Password, user.Password) {
		token, err := helpers.GenerateToken(user.Nickname)
		if err != nil {
			fmt.Println("Error while generating JWT token:", err)
			return "", fmt.Errorf("not valid")
		}
		return token, nil
	}
	return "", fmt.Errorf("invalid credentials")
}

func UserExists(nickname string) bool {
	var user models.User
	result := db.Users.FindOne(context.TODO(), bson.M{
		"nickname": nickname,
	})
	if err := result.Decode(&user); err != nil {
		return false
	}
	return user.Nickname == nickname
}

func CreateUser(nickname, password string) (primitive.ObjectID, error) {
	if len(nickname) < 3 {
		return primitive.ObjectID{}, fmt.Errorf("invalid nickname")
	}

	if len(password) < 6 {
		return primitive.ObjectID{}, fmt.Errorf("password length must be > 6")
	}

	if UserExists(nickname) {
		return primitive.ObjectID{}, fmt.Errorf("user already exists")
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		fmt.Println("[Hashing password] Error:", err)
		return primitive.ObjectID{}, err
	}
	resultInsert, err := db.Users.InsertOne(context.TODO(), models.User{
		Nickname: nickname,
		Password: hashedPassword,
	})

	if err != nil {
		fmt.Println("[Error Users.Insert] Error", err.Error())
		return primitive.ObjectID{}, err
	}

	return resultInsert.InsertedID.(primitive.ObjectID), nil
}
