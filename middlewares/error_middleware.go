package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}
		switch t := err.Err.(type) {
		case *errors.AppError:
			c.JSON(t.HttpCode, errors.ErrorResponse(t))
		default:
			c.JSON(500, errors.ErrInternal(err.Error()))
		}
	}
}
