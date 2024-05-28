package validation

import (
	"fmt"
	"strings"
)

type Result struct {
	FieldErrors []FieldError
}

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func (result *Result) AddFieldError(field string, error string) {
	fieldError := FieldError{
		Field: field,
		Error: error,
	}
	result.FieldErrors = append(result.FieldErrors, fieldError)
}

func (result *Result) ToError() error {
	if len(result.FieldErrors) == 0 {
		return nil
	}

	return result
}

func (result Result) Error() string {
	fieldErrorMsgs := make([]string, len(result.FieldErrors))
	for i, fieldError := range result.FieldErrors {
		fieldErrorMsgs[i] = fmt.Sprintf("field='%s' error='%s'", fieldError.Field, fieldError.Error)
	}

	return strings.Join(fieldErrorMsgs, ",")
}

var IsRequiredMsg = "is required"
