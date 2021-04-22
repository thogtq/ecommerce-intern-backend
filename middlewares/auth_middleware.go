package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/helpers"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get(("token"))
		clientRefreshToken := c.Request.Header.Get(("refreshToken"))
		claims, err := helpers.ValidateToken(clientToken)
		if err != nil {
			//Try to handle expired access token
			if err == errors.ErrExpiredToken && clientRefreshToken != "" {
				claims, err = helpers.ValidateToken(clientRefreshToken)
				if err != nil {
					fmt.Print("trigger here!")
					c.Error(errors.ErrUnauthorized)
					c.Abort()
					return
				}
			} else { //Another token error
				c.Error(err)
				c.Abort()
				return
			}
		}
		controllers.SetContextUserID(c, claims.UserID)
		c.Next()
	}
}
func AdminAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get(("token"))
		claims, err := helpers.ValidateToken(clientToken)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		if claims.Role != "admin" {
			c.Error(errors.ErrUnauthorized)
			c.Abort()
			return
		}
		controllers.SetContextUserID(c, claims.UserID)
		c.Next()
	}
}
