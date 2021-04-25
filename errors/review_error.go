package errors

var (
	ErrReviewNotFound = &AppError{
		HttpCode: 400,
		Code:     "REVIEW_NOT_FOUND",
		Message:  "Review not found",
	}
)
