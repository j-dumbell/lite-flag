/*
Liteflag API

API for managing feature flags.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FlagInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlagInput{}

// FlagInput struct for FlagInput
type FlagInput struct {
	// Whether the feature flag is public or not
	IsPublic bool `json:"isPublic"`
	// Type of the flag value
	Type string `json:"type"`
	Value FlagValue `json:"value"`
}

type _FlagInput FlagInput

// NewFlagInput instantiates a new FlagInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlagInput(isPublic bool, type_ string, value FlagValue) *FlagInput {
	this := FlagInput{}
	this.IsPublic = isPublic
	this.Type = type_
	this.Value = value
	return &this
}

// NewFlagInputWithDefaults instantiates a new FlagInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlagInputWithDefaults() *FlagInput {
	this := FlagInput{}
	return &this
}

// GetIsPublic returns the IsPublic field value
func (o *FlagInput) GetIsPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value
// and a boolean to check if the value has been set.
func (o *FlagInput) GetIsPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublic, true
}

// SetIsPublic sets field value
func (o *FlagInput) SetIsPublic(v bool) {
	o.IsPublic = v
}

// GetType returns the Type field value
func (o *FlagInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FlagInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FlagInput) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *FlagInput) GetValue() FlagValue {
	if o == nil {
		var ret FlagValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FlagInput) GetValueOk() (*FlagValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FlagInput) SetValue(v FlagValue) {
	o.Value = v
}

func (o FlagInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlagInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isPublic"] = o.IsPublic
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *FlagInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isPublic",
		"type",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFlagInput := _FlagInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlagInput)

	if err != nil {
		return err
	}

	*o = FlagInput(varFlagInput)

	return err
}

type NullableFlagInput struct {
	value *FlagInput
	isSet bool
}

func (v NullableFlagInput) Get() *FlagInput {
	return v.value
}

func (v *NullableFlagInput) Set(val *FlagInput) {
	v.value = val
	v.isSet = true
}

func (v NullableFlagInput) IsSet() bool {
	return v.isSet
}

func (v *NullableFlagInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlagInput(val *FlagInput) *NullableFlagInput {
	return &NullableFlagInput{value: val, isSet: true}
}

func (v NullableFlagInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlagInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


