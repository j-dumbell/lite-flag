# Flag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Unique identifier for the feature flag | 
**Type** | **string** | Type of the flag value | 
**IsPublic** | **bool** | Whether or not the flag is public. | 
**Value** | [**FlagValue**](FlagValue.md) |  | 

## Methods

### NewFlag

`func NewFlag(key string, type_ string, isPublic bool, value FlagValue, ) *Flag`

NewFlag instantiates a new Flag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagWithDefaults

`func NewFlagWithDefaults() *Flag`

NewFlagWithDefaults instantiates a new Flag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Flag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Flag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Flag) SetKey(v string)`

SetKey sets Key field to given value.


### GetType

`func (o *Flag) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Flag) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Flag) SetType(v string)`

SetType sets Type field to given value.


### GetIsPublic

`func (o *Flag) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Flag) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Flag) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetValue

`func (o *Flag) GetValue() FlagValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Flag) GetValueOk() (*FlagValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Flag) SetValue(v FlagValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


