package errors

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
	ErrNoUserID = &AppError{
		HttpCode: 400,
		Code:     "MISSING_USERID",
		Message:  "UserID is missing",
	}
	ErrInvalidUserID = &AppError{
		HttpCode: 400,
		Code:     "INVALID_USERID",
		Message:  "UserID is invalid",
	}
	ErrUserNotFound = &AppError{
		HttpCode: 400,
		Code:     "USER_NOT_FOUND",
		Message:  "User not found",
	}
)
