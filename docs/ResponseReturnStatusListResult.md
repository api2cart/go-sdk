# ResponseReturnStatusListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnStatuses** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseReturnStatusListResult

`func NewResponseReturnStatusListResult() *ResponseReturnStatusListResult`

NewResponseReturnStatusListResult instantiates a new ResponseReturnStatusListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseReturnStatusListResultWithDefaults

`func NewResponseReturnStatusListResultWithDefaults() *ResponseReturnStatusListResult`

NewResponseReturnStatusListResultWithDefaults instantiates a new ResponseReturnStatusListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnStatuses

`func (o *ResponseReturnStatusListResult) GetReturnStatuses() []Status`

GetReturnStatuses returns the ReturnStatuses field if non-nil, zero value otherwise.

### GetReturnStatusesOk

`func (o *ResponseReturnStatusListResult) GetReturnStatusesOk() (*[]Status, bool)`

GetReturnStatusesOk returns a tuple with the ReturnStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnStatuses

`func (o *ResponseReturnStatusListResult) SetReturnStatuses(v []Status)`

SetReturnStatuses sets ReturnStatuses field to given value.

### HasReturnStatuses

`func (o *ResponseReturnStatusListResult) HasReturnStatuses() bool`

HasReturnStatuses returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseReturnStatusListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseReturnStatusListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseReturnStatusListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseReturnStatusListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseReturnStatusListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseReturnStatusListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseReturnStatusListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseReturnStatusListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseReturnStatusListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseReturnStatusListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseReturnStatusListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseReturnStatusListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


