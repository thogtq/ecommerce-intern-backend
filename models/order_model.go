package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductOrder struct {
	ProductID string `bson:"productID" json:"productID"`
	Color     string `bson:"color,omitempty" json:"color"`
	Size      string `bson:"size,omitempty" json:"size"`
	Quantity  int    `bson:"quantity,omitempty" json:"quantity"`
	Amount    int    `bson:"amount" json:"amount"`
}
type Order struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	OrderID       string             `bson:"orderID,omitempty" json:"orderID"`
	UserID        string             `bson:"userID,omitempty" json:"userID"`
	Status        string             `bson:"status,omitempty" json:"status"`
	OrderDate     time.Time          `bson:"orderDate,omitempty" json:"orderDate"`
	Subtotal      int                `bson:"subtotal,omitempty" json:"subtotal"`
	Note          string             `bson:"note,omitempty" json:"note"`
	ProductsOrder []ProductOrder     `bson:"products,omitempty" json:"products"`
}
type OrderFilter struct {
	Limit  int       `form:"limit"`
	Date   time.Time `form:"date"`
	Search string    `form:"search"`
	Page   int       `form:"page"`
}
