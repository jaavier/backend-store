package models

type Product struct {
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Stock       int    `bson:"stock"`
	StoreId     int    `bson:"storeId"`
}
