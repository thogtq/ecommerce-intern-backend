package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(userPassword string) (hashedPassword string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(userPassword), 14)
	return string(bytes)
}
func VerifyPassword(hashedPassword string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
	return err == nil
}
func CheckIfEmailExist(ctx context.Context, userCollection *mongo.Collection, userEmail string) bool {
	result := userCollection.FindOne(ctx, bson.M{"email": userEmail})
	return result.Err() != mongo.ErrNoDocuments
}
