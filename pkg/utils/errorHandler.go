package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (r Response) ErrorHandler(err error) {
	var errData map[string]any
	var errMsg string

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		errData = map[string]any{}
		for _, err := range castedObject {
			field := PascalToSnake(err.Field())
			switch err.Tag() {
			case "required":
				errData[field] = fmt.Sprintf("The %s field is required.", err.Field())
			case "email":
				errData[field] = fmt.Sprintf("The %s field must contain a valid email address.", err.Field())
			case "gte":
				errData[field] = fmt.Sprintf("The %s field must contain a number greater than or equal to %s.", err.Field(), err.Param())
			case "lte":
				errData[field] = fmt.Sprintf("The %s field must contain a number less than or equal to %s.", err.Field(), err.Param())
			case "len":
				errData[field] = fmt.Sprintf("The %s field must be exactly %s characters in length.", err.Field(), err.Param())
			case "min":
				errData[field] = fmt.Sprintf("The %s field must be at least %s characters in length.", err.Field(), err.Param())
			case "max":
				errData[field] = fmt.Sprintf("The %s field must be at most %s characters in length.", err.Field(), err.Param())
			case "oneof":
				errData[field] = fmt.Sprintf("The %s field must be %s.", err.Field(), err.Param())
			case "datetime":
				errData[field] = fmt.Sprintf("The %s field must be format %s.", err.Field(), err.Param())
			case "number":
				errData[field] = fmt.Sprintf("The %s field must contain a number.", err.Field())
			case "required_if":
				errData[field] = fmt.Sprintf("The %s field is required.", err.Field())
			case "decimal":
				parts := strings.Split(err.Param(), " ")
				errData[field] = fmt.Sprintf("The %s field must contain a decimal number with format %s after comma.", err.Field(), parts[1])
			}
		}
	}
	errMsg = "invalid payloads"
	response := map[string]any{
		"data":    errData,
		"message": errMsg,
		"status":  false,
	}

	r.Context.JSON(r.StatusCode, response)
}
