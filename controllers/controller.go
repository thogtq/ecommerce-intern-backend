package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thogtq/ecommerce-server/errors"
)

type H map[string]interface{}

func SuccessResponse(data interface{}) H {
	return H{
		"status": "success",
		"data":   data,
	}
}
func SetContextUserID(ctx *gin.Context, userID string) {
	ctx.Set("userID", userID)
}
func GetContextUserID(ctx *gin.Context) (string, error) {
	userID, err := ctx.Get("userID")
	if !err {
		return "", errors.ErrInternal("field userID in context not found")
	}
	return userID.(string), nil
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
