package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/dao"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/models"
)

var userDAO dao.UserDAO

func Regiser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	newUser := &models.User{}

	c.BindJSON(newUser)
	res, err := userDAO.Register(ctx, newUser)
	if err != nil {
		c.JSON(400, errors.ErrorResponse(err))
		return
	}
	c.JSON(200, SuccessResponse(gin.H{"userID": res}))
}
func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	userLogin := &models.UserLogin{}

	c.BindJSON(userLogin)
	userObj, userToken, err := userDAO.Login(ctx, userLogin)
	if err != nil {
		c.JSON(400, errors.ErrorResponse(err))
		return
	}
	c.JSON(200, SuccessResponse(gin.H{
		"user":         userObj,
		"token":        userToken.AccessToken,
		"refreshToken": userToken.RefreshToken,
	}))
}
