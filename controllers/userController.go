package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userModel models.User

func Regiser(c *gin.Context) {
	newUser := &models.User{UserID: primitive.NilObjectID}
	res, err := userModel.Register(newUser)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "new user not added",
			"debug":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "new user added",
		"userID":  res,
	})
}
func Login(c *gin.Context) {
	userLogin := &models.Login{}
	c.BindJSON(userLogin)
	c.JSON(200, gin.H{
		"status": "success",
		"data":   userLogin.Email + ` ` + userLogin.Password,
	})
}
