package errors

import "errors"

var (
	ErrEmailExisted      = errors.New("EMAIL_EXISTED")
	ErrEmailNotExisted   = errors.New("EMAIL_NOT_EXISTED")
	ErrInvalidPassword   = errors.New("INVALID_PASSWORD")
	ErrNotInserted       = errors.New("NOT_INSERTED")
	ErrTokenNotGenerated = errors.New("TOKEN_NOT_GENERATED")
)

type H map[string]interface{}

func ErrorResponse(ErrCode error) H {
	return H{
		"status":  "error",
		"message": ErrCode.Error(),
	}
}
