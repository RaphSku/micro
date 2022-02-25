package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Price    float32            `json:"price" bson:"price"`
	Category string             `json:"category" bson:"category"`
	Storage  map[string]int     `json:"storage" bson:"storage"`
}

type ProductList struct {
	Products []Product `json:"products" bson:"products"`
}
