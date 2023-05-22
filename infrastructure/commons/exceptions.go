package commons

import "errors"

var (
	// ErrInvalidCredentials is thrown when the user credentials are invalid
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrUnauthorized is thrown when the user is unauthorized
	ErrUnauthorized = errors.New("unauthorized")
	// ErrInternalServerError is thrown when the server encounters an error
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound is thrown when the object is not found
	ErrNotFound = errors.New("not found")
	// ErrUserAlreadyExists is thrown when the user already exists
	ErrUserAlreadyExists = errors.New("user already exists")
	// ErrCategoryAlreadyExists is thrown when the category already exists
	ErrCategoryAlreadyExists = errors.New("category already exists")
	// ErrPaymentMethodAlreadyExists is thrown when the payment method already exists
	ErrPaymentMethodAlreadyExists = errors.New("payment method already exists")
	// ErrTransactionStatusAlreadyExists is thrown when the transaction status already exists
	ErrTransactionStatusAlreadyExists = errors.New("transaction status already exists")
	// ErrCategoryNotFound is thrown when the category is not found
	ErrCategoryNotFound = errors.New("category not found")
	// ErrUserNotFound is thrown when the user is not found
	ErrUserNotFound = errors.New("user not found")
	// ErrProductNotFound is thrown when the product is not found
	ErrProductNotFound = errors.New("product not found")
	// ErrTransactionNotFound is thrown when the transaction is not found
	ErrTransactionNotFound = errors.New("transaction not found")
	// ErrProductOutOfStock is thrown when the product is out of stock
	ErrProductOutOfStock = errors.New("product out of stock")
	// ErrCartIsEmpty is thrown when the cart is empty
	ErrCartIsEmpty = errors.New("cart is empty")
	// ErrEmptyInput is thrown when the input is empty
	ErrEmptyInput = errors.New("empty input")
	// ErrValidationFailed is thrown when the input validation is failed
	ErrValidationFailed = errors.New("validation failed")
	// ErrMissingId is thrown when id is missing
	ErrMissingId = errors.New("missing require id")
	// ErrBadRequest is thrown when request message is invalid
	ErrBadRequest = errors.New("invalid request message")
)
