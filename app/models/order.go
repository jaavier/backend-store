package models

import "time"

type Order struct {
	Products []string  `json:"productsIds" bson:"productsIds"`
	Date     time.Time `json:"date" bson:"date"`
	Id       string    `json:"orderId" bson:"orderId"`
}
