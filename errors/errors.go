package errors

type H map[string]interface{}

type AppError struct {
	HttpCode int
	Code     string
	Message  string
}

const INTERNAL_ERROR_CODE = "INTERNAL_ERROR"

func (se AppError) Error() string {
	return se.Message
}
func ErrorResponse(err error) H {
	switch t := err.(type) {
	case *AppError:
		return H{
			"status": "error",
			"error": H{
				"httpCode": t.HttpCode,
				"code":     t.Code,
				"message":  t.Message,
			},
		}
	default:
		return H{
			"status": "error",
			"error": H{
				"httpCode": 500,
				"code":     INTERNAL_ERROR_CODE,
				"message":  err.Error(),
			},
		}
	}

}

var (
	ErrEmailExisted = &AppError{
		HttpCode: 400,
		Code:     "EMAIL_EXISTED",
		Message:  "Email address have been existed",
	}
	ErrInvalidPassword = &AppError{
		HttpCode: 400,
		Code:     "INVALID_PASSWORD",
		Message:  "Invalid password",
	}
	ErrEmailNotFound = &AppError{
		HttpCode: 400,
		Code:     "EMAIL_NOT_FOUND",
		Message:  "Email address not found",
	}
	ErrUnauthorized = &AppError{
		HttpCode: 400,
		Code:     "UNAUTHORIZED",
		Message:  "Unauthorized access",
	}
	ErrInvalidToken = &AppError{
		HttpCode: 400,
		Code:     "INVALID_TOKEN",
		Message:  "Your token is invalid",
	}
	ErrExpiredToken = &AppError{
		HttpCode: 400,
		Code:     "EXPIRED_TOKEN",
		Message:  "Your token is expired",
	}
	ErrInvalidExtension = &AppError{
		HttpCode: 400,
		Code:     "INVALID_FILE_EXTENSION",
		Message:  "Invalid file extension",
	}
	ErrNoFile = &AppError{
		HttpCode: 400,
		Code:     "NO_FILE",
		Message:  "No file received",
	}
)

func ErrInternal(msg string) *AppError {
	return &AppError{
		HttpCode: 500,
		Code:     INTERNAL_ERROR_CODE,
		Message:  msg,
	}
}
