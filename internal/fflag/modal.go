package fflag

import (
	"slices"
	"unicode"

	"github.com/j-dumbell/lite-flag/pkg/validation"
)

type FlagType string

const (
	FlagTypeString  FlagType = "string"
	FlagTypeBoolean FlagType = "boolean"
	FlagTypeJSON    FlagType = "json"
)

type Flag struct {
	Key          string                 `json:"key"`
	Type         FlagType               `json:"type"`
	BooleanValue *bool                  `json:"booleanValue,omitempty"`
	StringValue  *string                `json:"stringValue,omitempty"`
	JSONValue    map[string]interface{} `json:"jsonValue,omitempty"`
}

func isValidKey(key string) bool {
	if key == "" {
		return false
	}

	for _, r := range key {
		if !(unicode.IsLetter(r) || unicode.IsNumber(r) || r == '-' || r == '_') {
			return false
		}
	}
	return true
}

func (flag *Flag) Validate() error {
	validationResult := validation.Result{}

	if !isValidKey(flag.Key) {
		validationResult.AddFieldError("key", "must be non-empty, and can only contain letters, numbers, hyphens and underscores")
	}

	if !slices.Contains([]FlagType{FlagTypeString, FlagTypeBoolean, FlagTypeJSON}, flag.Type) {
		validationResult.AddFieldError("type", "must be one of 'string' | 'boolean' | 'json'")
		return validationResult
	}

	switch flag.Type {
	case FlagTypeString:
		if flag.JSONValue != nil {
			validationResult.AddFieldError("jsonValue", "must be nil when flag is string type")
		}
		if flag.BooleanValue != nil {
			validationResult.AddFieldError("booleanValue", "must be nil when flag is string type")
		}
	case FlagTypeBoolean:
		if flag.JSONValue != nil {
			validationResult.AddFieldError("jsonValue", "must be nil when flag is boolean type")
		}
		if flag.StringValue != nil {
			validationResult.AddFieldError("stringValue", "must be nil when flag is boolean type")
		}
	case FlagTypeJSON:
		if flag.StringValue != nil {
			validationResult.AddFieldError("stringValue", "must be nil when flag is JSON type")
		}
		if flag.BooleanValue != nil {
			validationResult.AddFieldError("booleanValue", "must be nil when flag is JSON type")
		}
	}

	return validationResult.ToError()
}
