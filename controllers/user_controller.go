package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/dao"
	"github.com/thogtq/ecommerce-server/helpers"
	"github.com/thogtq/ecommerce-server/models"
)

var userDAO dao.UserDAO

func Regiser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	newUser := &models.User{}
	c.BindJSON(newUser)
	res, err := userDAO.CreateUser(ctx, newUser)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"userID": res}))
}
func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	userLogin := &models.UserLogin{}
	userToken := &models.UserToken{}
	c.BindJSON(userLogin)
	userObj, err := userDAO.Login(ctx, userLogin)
	if err != nil {
		c.Error(err)
		return
	}
	userToken.AccessToken, userToken.RefreshToken, err = helpers.GenerateTokens(userObj.UserID.Hex(), userObj.Email)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{
		"user":         userObj,
		"token":        userToken.AccessToken,
		"refreshToken": userToken.RefreshToken,
	}))
}
func UpdateUser(c *gin.Context) {
	userID, _ := GetContextUserID(c)
	userDetails, err := userDAO.GetUserByUserID(c, userID)
	_, _ = userDetails, err
	c.JSON(200, gin.H{"status": userDetails, "id": userID})
}
