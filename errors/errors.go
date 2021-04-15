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
func ErrInternal(msg string) *AppError {
	return &AppError{
		HttpCode: 500,
		Code:     INTERNAL_ERROR_CODE,
		Message:  msg,
	}
}
