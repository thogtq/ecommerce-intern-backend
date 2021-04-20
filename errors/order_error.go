package errors

var (
	ErrInvalidOrderStatus = &AppError{
		HttpCode: 400,
		Code:     "INVALID_ORDER_STATUS",
		Message:  "Status is invalid",
	}
	ErrNoOrderID = &AppError{
		HttpCode: 400,
		Code:     "MISSING_ORDERID",
		Message:  "Your orderID is missing",
	}
	ErrOrderIDNotFound = &AppError{
		HttpCode: 400,
		Code:     "OrderID_NOT_FOUND",
		Message:  "OrderID not found",
	}
	ErrEmptyCart = &AppError{
		HttpCode: 400,
		Code:     "EMPTY_CART",
		Message:  "Your cart is empty",
	}
)
