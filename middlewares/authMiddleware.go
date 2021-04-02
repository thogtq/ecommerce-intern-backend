package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/helpers"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get(("token"))
		if clientToken != "" {
			if claims, err := helpers.ValidateToken(clientToken); err == nil {
				c.Set("userID", claims.UserID)
				c.Set("email", claims.Email)
			}
		}
	}
}
