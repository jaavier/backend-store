package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Nickname string             `json:"nickname" bson:"nickname"`
	Password string             `json:"password" bson:"password"`
	Role     string             `json:"role" bson:"role"`
}
