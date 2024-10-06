package exceptions

import "errors"

var (
	InvalidInput                = errors.New("invalid input")
	InvalidToken                = errors.New("invalid token")
	InvalidUsernameOrPassword   = errors.New("username or password invalid")
	RequestBodyValidationFailed = errors.New("request body validation failed")
	ResourceNotFound            = errors.New("resource not found")
	UserAlreadyExists           = errors.New("user already exists in our database")
	UserNotFoundInOurDB         = errors.New("user does not exists in our database")
)
