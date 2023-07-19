package models

type Product struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Stock       int     `json:"stock" bson:"stock"`
	StoreId     int     `json:"storeId" bson:"storeId"`
	Price       float64 `json:"price" bson:"price"`
	Id          string  `json:"productId,omitempty" bson:"productId,omitempty"`
}
