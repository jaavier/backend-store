package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Products []primitive.ObjectID `json:"productsIds" bson:"productsIds"`
	Date     time.Time            `json:"date" bson:"date"`
	Id       string               `json:"orderId" bson:"orderId"`
	Total    float64              `json:"total" bson:"total"`
	Step     string               `json:"step" bson:"step"`
	UserId   primitive.ObjectID   `json:"userId" bson:"userId"`
}
