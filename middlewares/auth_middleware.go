package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/controllers"
	"github.com/thogtq/ecommerce-server/errors"
	"github.com/thogtq/ecommerce-server/helpers"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get(("token"))
		claims, err := helpers.ValidateToken(clientToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.ErrorResponse(errors.ErrUnauthorized))
			return
		}
		_ = claims
		controllers.SetContextUserID(c, claims.UserID)
		c.Next()
	}
}
