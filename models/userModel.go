package models

import (
	"context"
	"fmt"
	"time"

	"github.com/thogtq/ecommerce-server/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type User struct {
	UserID   primitive.ObjectID `bson:"_id,omitempty" json:"userID,omitempty"`
	Fullname string             `bson:"fullname" json:"fullname"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Token    string             `bson:"token" json:"token,omitempty"`
}

func (*User) Register(userData *User) (string, error) {
	userCollection = database.DBClient.Database("ecommerce").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := userCollection.InsertOne(ctx, userData)
	if err != nil {
		return "", fmt.Errorf("can not insert new user %v", err)
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
