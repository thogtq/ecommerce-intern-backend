package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/models"
)

var userModel models.User

func Regiser(c *gin.Context) {
	newUser := &models.User{}
	c.BindJSON(newUser)
	res, err := userModel.Register(newUser)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"data":    gin.H{"userID": res},
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
