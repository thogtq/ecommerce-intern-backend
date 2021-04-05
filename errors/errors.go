package errors

import (
	"errors"
)

//Mapping error code, http code , message<-client

var (
	// ErrInternal = &ServerError{
	// 	ErrCode: "INTERNAL_ERROR",
	// 	ErrMsg:  "Internal server error",
	// }
	// ErrEmailExisted_ = &ClientError{
	// 	ErrCode: "EMAIL_EXISTED",
	// 	ErrMsg:  "Email address have been existed!",
	// }
	ErrEmailExisted    = errors.New("EMAIL_EXISTED")
	ErrInvalidPassword = errors.New("INVALID_PASSWORD")
	ErrEmailNotFound   = errors.New("EMAIL_NOT_FOUND")
)
var (
	ErrUnauthorized = errors.New("UNAUTHORIZED")
	ErrInvalidToken = errors.New("INVALID_TOKEN")
	ErrExpiredToken = errors.New("EXPIRED_TOKEN")
)

type H map[string]interface{}

//Status error
// func ErrorResponse_(err interface{},) H {
// 	return H{
// 		"status": "error",
// 		"error": H{
// 			"code":    err.Err(),
// 			"message": err.Msg,
// 		},
// 	}
// }

func ErrorResponse(ErrCode error) H {
	return H{
		"status":  "error",
		"message": ErrCode.Error(),
	}
}
