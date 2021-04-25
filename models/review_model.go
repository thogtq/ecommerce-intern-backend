package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ReviewID   primitive.ObjectID `bson:"_id,omitempty" json:"reviewID"`
	ProductID  string             `bson:"productID" json:"productID"`
	UserID     string             `bson:"userID" json:"userID"`
	Fullname   string             `bson:"fullname" json:"fullname"`
	ReviewDate time.Time          `bson:"reviewDate" json:"reviewDate"`
	Title      string             `bson:"title" json:"title"`
	Content    string             `bson:"content" json:"content"`
	Star       int                `bson:"star" json:"star"`
}
type ReviewFilters struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}
