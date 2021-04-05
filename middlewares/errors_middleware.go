package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/errors"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}
		switch err.Err.(type) {
		case *errors.ClientError:
			c.JSON(400, errors.ErrorResponse(err.Err.(*errors.ClientError)))
		case *errors.ServerError:
			c.JSON(500, errors.ErrorResponse(err.Err.(*errors.ServerError)))
		}
	}
}
