package errors

var (
	ErrProductNotFound = &AppError{
		HttpCode: 400,
		Code:     "PRODUCT_NOT_FOUND",
		Message:  "Product not found",
	}
	ErrInvalidProductFilter = &AppError{
		HttpCode: 400,
		Code:     "INVALID_PRODUCT_FILTER",
		Message:  "Invalid product product filter parameters",
	}
)
