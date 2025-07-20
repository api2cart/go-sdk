# ResponseOrderStatusListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartOrderStatuses** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseOrderStatusListResult

`func NewResponseOrderStatusListResult() *ResponseOrderStatusListResult`

NewResponseOrderStatusListResult instantiates a new ResponseOrderStatusListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrderStatusListResultWithDefaults

`func NewResponseOrderStatusListResultWithDefaults() *ResponseOrderStatusListResult`

NewResponseOrderStatusListResultWithDefaults instantiates a new ResponseOrderStatusListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartOrderStatuses

`func (o *ResponseOrderStatusListResult) GetCartOrderStatuses() []Status`

GetCartOrderStatuses returns the CartOrderStatuses field if non-nil, zero value otherwise.

### GetCartOrderStatusesOk

`func (o *ResponseOrderStatusListResult) GetCartOrderStatusesOk() (*[]Status, bool)`

GetCartOrderStatusesOk returns a tuple with the CartOrderStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartOrderStatuses

`func (o *ResponseOrderStatusListResult) SetCartOrderStatuses(v []Status)`

SetCartOrderStatuses sets CartOrderStatuses field to given value.

### HasCartOrderStatuses

`func (o *ResponseOrderStatusListResult) HasCartOrderStatuses() bool`

HasCartOrderStatuses returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseOrderStatusListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseOrderStatusListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseOrderStatusListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseOrderStatusListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseOrderStatusListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseOrderStatusListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseOrderStatusListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseOrderStatusListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseOrderStatusListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseOrderStatusListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseOrderStatusListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseOrderStatusListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


