package errors

import "errors"

var (
	Err                = errors.New("")
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

func ErrorResponse(ErrCode error) H {
	return H{
		"status":  "error",
		"message": ErrCode.Error(),
	}
}
