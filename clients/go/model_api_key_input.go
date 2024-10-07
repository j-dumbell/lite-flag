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

// checks if the ApiKeyInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyInput{}

// ApiKeyInput struct for ApiKeyInput
type ApiKeyInput struct {
	// Unique identifier for the API key
	Name string `json:"name"`
	// The associated permissions of the key
	Role string `json:"role"`
}

type _ApiKeyInput ApiKeyInput

// NewApiKeyInput instantiates a new ApiKeyInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyInput(name string, role string) *ApiKeyInput {
	this := ApiKeyInput{}
	this.Name = name
	this.Role = role
	return &this
}

// NewApiKeyInputWithDefaults instantiates a new ApiKeyInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyInputWithDefaults() *ApiKeyInput {
	this := ApiKeyInput{}
	return &this
}

// GetName returns the Name field value
func (o *ApiKeyInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiKeyInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiKeyInput) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value
func (o *ApiKeyInput) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ApiKeyInput) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ApiKeyInput) SetRole(v string) {
	o.Role = v
}

func (o ApiKeyInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *ApiKeyInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"role",
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

	varApiKeyInput := _ApiKeyInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiKeyInput)

	if err != nil {
		return err
	}

	*o = ApiKeyInput(varApiKeyInput)

	return err
}

type NullableApiKeyInput struct {
	value *ApiKeyInput
	isSet bool
}

func (v NullableApiKeyInput) Get() *ApiKeyInput {
	return v.value
}

func (v *NullableApiKeyInput) Set(val *ApiKeyInput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyInput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyInput(val *ApiKeyInput) *NullableApiKeyInput {
	return &NullableApiKeyInput{value: val, isSet: true}
}

func (v NullableApiKeyInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

