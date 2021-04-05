package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
		return "", fmt.Errorf("can not get user id from context, userID field was not setted")
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