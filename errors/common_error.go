package errors

var(
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