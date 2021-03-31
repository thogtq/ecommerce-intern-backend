package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductOrder struct {
	ProductID primitive.ObjectID
	Color     string
	Size      string
	Quantity  int
	Amount    float32
}
type Order struct {
	OrderID       primitive.ObjectID
	UseID         primitive.ObjectID
	Status        string
	OrderDate     time.Time
	SubTotal      float32
	ProductsOrder []ProductOrder
}
