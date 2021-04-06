package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/dao"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/helpers"
	"github.com/thogtq/ecommerce-server/models"
)

var userDAO dao.UserDAO

func Regiser(c *gin.Context) {
	newUser := &models.User{}
	c.BindJSON(newUser)
	res, err := userDAO.CreateUser(c.Request.Context(), newUser)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"userID": res}))
}
func Login(c *gin.Context) {
	userLogin := &models.UserLogin{}
	userToken := &models.UserToken{}
	c.BindJSON(userLogin)
	userObj, err := userDAO.Login(c.Request.Context(), userLogin)
	if err != nil {
		c.Error(err)
		return
	}
	userToken.AccessToken, userToken.RefreshToken, err = helpers.GenerateTokens(userObj.UserID.Hex(), userObj.Role)
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

func AdminLogin(c *gin.Context) {
	userLogin := &models.UserLogin{}
	userToken := &models.UserToken{}
	c.BindJSON(userLogin)
	userObj, err := userDAO.Login(c.Request.Context(), userLogin)
	if err != nil {
		c.Error(err)
		return
	}
	if userObj.Role != "admin" {
		c.Error(errors.ErrUnauthorized)
		return
	}
	userToken.AccessToken, _, err = helpers.GenerateTokens(userObj.UserID.Hex(), userObj.Role)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{
		"user":  userObj,
		"token": userToken.AccessToken,
	}))
}
