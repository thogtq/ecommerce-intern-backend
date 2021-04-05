package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID   primitive.ObjectID `bson:"_id,omitempty" json:"userID,omitempty"` //Fix me
	Fullname string             `bson:"fullname" json:"fullname"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Role     string             `bson:"role" json:"role,omitempty"`
}
type UserToken struct {
	AccessToken  string
	RefreshToken string
}
type UserLogin struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (*User) HashPassword(userPassword string) (hashedPassword string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(userPassword), 14)
	return string(bytes)
}
func (*User) VerifyPassword(hashedPassword string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
	return err == nil
}
func (*User) CheckIfEmailExist(ctx context.Context, userColl *mongo.Collection, userEmail string) bool {
	result := userColl.FindOne(ctx, bson.M{"email": userEmail})
	err := result.Decode(&User{})
	return err != mongo.ErrNoDocuments
}
