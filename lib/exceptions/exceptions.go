package exceptions

import "errors"

var (
	InvalidUsernameOrPassword   = errors.New("username or password invalid")
	RequestBodyValidationFailed = errors.New("request body validation failed")
	ResourceNotFound            = errors.New("resource not found")
	UserNotFoundInOurDB         = errors.New("user does not exists in our database")
)
