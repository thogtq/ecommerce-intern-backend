package middlewares

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
