# ResponseAccountSupportedPlatformsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedPlatformsCount** | Pointer to **NullableInt32** |  | [optional] 
**SupportedPlatforms** | Pointer to [**[]AccountSupportedPlatform**](AccountSupportedPlatform.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseAccountSupportedPlatformsResult

`func NewResponseAccountSupportedPlatformsResult() *ResponseAccountSupportedPlatformsResult`

NewResponseAccountSupportedPlatformsResult instantiates a new ResponseAccountSupportedPlatformsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAccountSupportedPlatformsResultWithDefaults

`func NewResponseAccountSupportedPlatformsResultWithDefaults() *ResponseAccountSupportedPlatformsResult`

NewResponseAccountSupportedPlatformsResultWithDefaults instantiates a new ResponseAccountSupportedPlatformsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedPlatformsCount

`func (o *ResponseAccountSupportedPlatformsResult) GetSupportedPlatformsCount() int32`

GetSupportedPlatformsCount returns the SupportedPlatformsCount field if non-nil, zero value otherwise.

### GetSupportedPlatformsCountOk

`func (o *ResponseAccountSupportedPlatformsResult) GetSupportedPlatformsCountOk() (*int32, bool)`

GetSupportedPlatformsCountOk returns a tuple with the SupportedPlatformsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPlatformsCount

`func (o *ResponseAccountSupportedPlatformsResult) SetSupportedPlatformsCount(v int32)`

SetSupportedPlatformsCount sets SupportedPlatformsCount field to given value.

### HasSupportedPlatformsCount

`func (o *ResponseAccountSupportedPlatformsResult) HasSupportedPlatformsCount() bool`

HasSupportedPlatformsCount returns a boolean if a field has been set.

### SetSupportedPlatformsCountNil

`func (o *ResponseAccountSupportedPlatformsResult) SetSupportedPlatformsCountNil(b bool)`

 SetSupportedPlatformsCountNil sets the value for SupportedPlatformsCount to be an explicit nil

### UnsetSupportedPlatformsCount
`func (o *ResponseAccountSupportedPlatformsResult) UnsetSupportedPlatformsCount()`

UnsetSupportedPlatformsCount ensures that no value is present for SupportedPlatformsCount, not even an explicit nil
### GetSupportedPlatforms

`func (o *ResponseAccountSupportedPlatformsResult) GetSupportedPlatforms() []AccountSupportedPlatform`

GetSupportedPlatforms returns the SupportedPlatforms field if non-nil, zero value otherwise.

### GetSupportedPlatformsOk

`func (o *ResponseAccountSupportedPlatformsResult) GetSupportedPlatformsOk() (*[]AccountSupportedPlatform, bool)`

GetSupportedPlatformsOk returns a tuple with the SupportedPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPlatforms

`func (o *ResponseAccountSupportedPlatformsResult) SetSupportedPlatforms(v []AccountSupportedPlatform)`

SetSupportedPlatforms sets SupportedPlatforms field to given value.

### HasSupportedPlatforms

`func (o *ResponseAccountSupportedPlatformsResult) HasSupportedPlatforms() bool`

HasSupportedPlatforms returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseAccountSupportedPlatformsResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseAccountSupportedPlatformsResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseAccountSupportedPlatformsResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseAccountSupportedPlatformsResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseAccountSupportedPlatformsResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseAccountSupportedPlatformsResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseAccountSupportedPlatformsResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseAccountSupportedPlatformsResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseAccountSupportedPlatformsResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseAccountSupportedPlatformsResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseAccountSupportedPlatformsResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseAccountSupportedPlatformsResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


