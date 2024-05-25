package validation

import "encoding/json"

type ValidationError struct {
	FieldErrors map[string]string `json:"errors"`
}

func (validationError ValidationError) Error() string {
	jsonError, err := json.Marshal(validationError.FieldErrors)
	if err != nil {
		return ""
	}

	return string(jsonError)
}

func (validationError *ValidationError) AddFieldError(field string, error string) {
	validationError.FieldErrors[field] = error
}

func NewValidationError() ValidationError {
	return ValidationError{
		FieldErrors: map[string]string{},
	}
}
