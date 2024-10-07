/*
Liteflag API

API for managing feature flags.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// FlagValue - Value of the flag, must match the type
type FlagValue struct {
	Bool *bool
	String *string
}

// boolAsFlagValue is a convenience function that returns bool wrapped in FlagValue
func BoolAsFlagValue(v *bool) FlagValue {
	return FlagValue{
		Bool: v,
	}
}

// stringAsFlagValue is a convenience function that returns string wrapped in FlagValue
func StringAsFlagValue(v *string) FlagValue {
	return FlagValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *FlagValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			if err = validator.Validate(dst.Bool); err != nil {
				dst.Bool = nil
			} else {
				match++
			}
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(FlagValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(FlagValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FlagValue) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FlagValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableFlagValue struct {
	value *FlagValue
	isSet bool
}

func (v NullableFlagValue) Get() *FlagValue {
	return v.value
}

func (v *NullableFlagValue) Set(val *FlagValue) {
	v.value = val
	v.isSet = true
}

func (v NullableFlagValue) IsSet() bool {
	return v.isSet
}

func (v *NullableFlagValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlagValue(val *FlagValue) *NullableFlagValue {
	return &NullableFlagValue{value: val, isSet: true}
}

func (v NullableFlagValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlagValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

