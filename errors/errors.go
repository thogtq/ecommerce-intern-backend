package errors

type H map[string]interface{}

//Fix me
type error interface {
	Error() string
	GetMsg() string
}
type ClientError struct {
	ErrCode string
	ErrMsg  string
}

func (ce ClientError) Error() string {
	return ce.ErrCode
}
func (ce ClientError) GetMsg() string {
	return ce.ErrMsg
}

type ServerError struct {
	ErrCode string
	ErrMsg  string
}

func (se ServerError) Error() string {
	return se.ErrCode
}
func (se ServerError) GetMsg() string {
	return se.ErrMsg
}

func ErrorResponse(err error) H {
	return H{
		"status": "error",
		"error": H{
			"code":    err.Error(),
			"message": err.GetMsg(),
		},
	}
}

var (
	ErrEmailExisted = &ClientError{
		ErrCode: "EMAIL_EXISTED",
		ErrMsg:  "Email address have been existed",
	}
	ErrInvalidPassword = &ClientError{
		ErrCode: "INVALID_PASSWORD",
		ErrMsg:  "Invalid password",
	}
	ErrEmailNotFound = &ClientError{
		ErrCode: "EMAIL_NOT_FOUND",
		ErrMsg:  "Email address not found",
	}
	ErrUnauthorized = &ClientError{
		ErrCode: "UNAUTHORIZED",
		ErrMsg:  "Unauthorized access",
	}
	ErrInvalidToken = &ClientError{
		ErrCode: "INVALID_TOKEN",
		ErrMsg:  "Your token is invalid",
	}
	ErrExpiredToken = &ClientError{
		ErrCode: "EXPIRED_TOKEN",
		ErrMsg:  "Your token is expired",
	}
	ErrInvalidExtension = &ClientError{
		ErrCode: "INVALID_FILE_EXTENSION",
		ErrMsg:  "Invalid file extension",
	}
	ErrNoFile = &ClientError{
		ErrCode: "NO_FILE",
		ErrMsg:  "No file received",
	}
)

func ErrInternal(msg string) *ServerError {
	return &ServerError{
		ErrCode: "INTERNAL_ERROR",
		ErrMsg:  msg,
	}
}
