package validators

import (
	"fmt"
	"html"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	"github.com/montinger-com/montinger-server/app/shared/models/response_model"
	"github.com/montinger-com/montinger-server/lib/exceptions"
)

func ValidateJsonBody[BodyType any](c *gin.Context) {

	var body BodyType

	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response_model.Result{
			Message: exceptions.RequestBodyValidationFailed.Error(),
			Errors:  []string{err.Error()},
		})
		c.Abort()
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		var validationErrors []string
		for _, errData := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, formatValidationError(errData))
		}

		c.JSON(http.StatusBadRequest, response_model.Result{
			Message: exceptions.RequestBodyValidationFailed.Error(),
			Errors:  validationErrors,
		})
		c.Abort()
		return
	}

	sanitize(&body)

	c.Set("body", body)
	c.Next()
}

func formatValidationError(err validator.FieldError) string {
	var builder strings.Builder

	field := strings.ToLower(err.Field())
	switch err.Tag() {
	case "required":
		fmt.Fprintf(&builder, "The %s field is required.", field)
	case "email":
		fmt.Fprintf(&builder, "The %s field must be a valid email address.", field)
	case "max":
		max := err.Param()
		fmt.Fprintf(&builder, "The %s field must be at most %s characters long.", field, max)
	case "min":
		min := err.Param()
		fmt.Fprintf(&builder, "The %s field must be at least %s characters long.", field, min)
	default:
		fmt.Fprintf(&builder, "Invalid value for the %s field.", field)
	}

	return builder.String()
}

// Sanitizes the fields in the parsed JSON
func sanitize(v interface{}) {
	value := reflect.ValueOf(v).Elem()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		switch field.Kind() {
		case reflect.String:
			value := field.String()
			sanitized := html.EscapeString(value)
			field.SetString(sanitized)
		case reflect.Struct:
			sanitize(field.Addr().Interface())
		}
	}
}

func ValidatePathParams[PathParamType any](c *gin.Context) {

	var pathParams PathParamType

	err := c.ShouldBindUri(&pathParams)
	if err != nil {

		c.JSON(http.StatusBadRequest, response_model.Result{
			Message: exceptions.RequestParamsValidationFailed.Error(),
			Errors:  []string{err.Error()},
		})
		c.Abort()
		return
	}

	validate := validator.New()
	if err := validate.Struct(pathParams); err != nil {
		var validationErrors []string
		for _, errData := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, formatValidationError(errData))
		}

		c.JSON(http.StatusBadRequest, response_model.Result{
			Message: exceptions.RequestParamsValidationFailed.Error(),
			Errors:  validationErrors,
		})
		c.Abort()
		return
	}

	// Sanitize the parsed path parameters
	sanitize(&pathParams)

	c.Set("params", pathParams)

	c.Next()

}
