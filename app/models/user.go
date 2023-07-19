package models

type User struct {
	Nickname string `json:"nickname" bson:"nickname"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}
