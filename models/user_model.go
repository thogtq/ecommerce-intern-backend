package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID   primitive.ObjectID `bson:"_id,omitempty" json:"userID,omitempty"`
	Fullname string             `bson:"fullname,omitempty" json:"fullname"`
	Email    string             `bson:"email,omitempty" json:"email"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Role     string             `bson:"role,omitempty" json:"role,omitempty"`
}
type UserToken struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	ExpiredAt    int64  `json:"expiredAt"`
}
type UserLogin struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
