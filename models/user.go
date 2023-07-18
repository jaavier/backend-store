package models

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Passwd   string `json:"passwd"`
}

var users = make(map[string]User)

func CreateUser(nickname, passwd string) (User, error) {
	if len(nickname) < 3 {
		return User{}, fmt.Errorf("invalid nickname")
	}

	if len(passwd) < 6 {
		return User{}, fmt.Errorf("password length must be > 6")
	}
	var uniqueId = uuid.NewString()
	fmt.Println(uniqueId)
	// if err != nil {
	// 	fmt.Printf("Error creating unique id: %x", err)
	// } else {
	// 	fmt.Printf("User id: %x", uniqueId.String())
	// }
	users[uniqueId] = User{Id: uniqueId, Nickname: nickname, Passwd: passwd}
	return users[uniqueId], nil
}
