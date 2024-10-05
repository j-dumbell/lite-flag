# FlagInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPublic** | **bool** | Whether the feature flag is public or not | 
**Type** | **string** | Type of the flag value | 
**Value** | [**FlagValue**](FlagValue.md) |  | 

## Methods

### NewFlagInput

`func NewFlagInput(isPublic bool, type_ string, value FlagValue, ) *FlagInput`

NewFlagInput instantiates a new FlagInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagInputWithDefaults

`func NewFlagInputWithDefaults() *FlagInput`

NewFlagInputWithDefaults instantiates a new FlagInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPublic

`func (o *FlagInput) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *FlagInput) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *FlagInput) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetType

`func (o *FlagInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlagInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlagInput) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *FlagInput) GetValue() FlagValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FlagInput) GetValueOk() (*FlagValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FlagInput) SetValue(v FlagValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


