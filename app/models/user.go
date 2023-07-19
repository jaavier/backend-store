package models

type User struct {
	Id       string `json:"userId" bson:"userId"`
	Nickname string `json:"nickname" bson:"nickname"`
	Password string `json:"password" bson:"password"`
}
