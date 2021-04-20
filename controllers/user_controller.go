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
	userData := &models.User{}
	userID, err := GetContextUserID(c)
	if err != nil {
		c.Error(errors.ErrUnauthorized)
		return
	}
	err = c.BindJSON(userData)
	if err != nil {
		c.Error(errors.ErrInternal(err.Error()))
		return
	}
	userDAO.New().UpdateUser(c, userData, userID)
	c.JSON(200, SuccessResponse(gin.H{"result": "updated"}))
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
func GetUser(c *gin.Context) {
	userID, err := GetContextUserID(c)
	if err != nil {
		c.Error(errors.ErrUnauthorized)
		return
	}
	user, err := userDAO.New().GetUserByUserID(c.Request.Context(), userID)
	if err != nil {
		c.Error(err)
		return
	}
	user.Password = ""
	c.JSON(200, SuccessResponse(gin.H{"user": user}))
}
func UpdateUserPassword(c *gin.Context) {
	var (
		params struct {
			OldPassword string `json:"oldPassword" binding:"required"`
			NewPassword string `json:"newPassword" binding:"required"`
		}
	)
	userID, err := GetContextUserID(c)
	if err != nil {
		c.Error(errors.ErrUnauthorized)
		return
	}
	err = c.BindJSON(params)
	if err != nil {
		//test
		c.Error(err)
	}
	err = userDAO.New().UpdateUserPassword(c, params.OldPassword, params.NewPassword, userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"result": "updated"}))
}
