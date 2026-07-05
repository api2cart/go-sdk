# ResponseReturnActionListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnActions** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseReturnActionListResult

`func NewResponseReturnActionListResult() *ResponseReturnActionListResult`

NewResponseReturnActionListResult instantiates a new ResponseReturnActionListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseReturnActionListResultWithDefaults

`func NewResponseReturnActionListResultWithDefaults() *ResponseReturnActionListResult`

NewResponseReturnActionListResultWithDefaults instantiates a new ResponseReturnActionListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnActions

`func (o *ResponseReturnActionListResult) GetReturnActions() []Status`

GetReturnActions returns the ReturnActions field if non-nil, zero value otherwise.

### GetReturnActionsOk

`func (o *ResponseReturnActionListResult) GetReturnActionsOk() (*[]Status, bool)`

GetReturnActionsOk returns a tuple with the ReturnActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnActions

`func (o *ResponseReturnActionListResult) SetReturnActions(v []Status)`

SetReturnActions sets ReturnActions field to given value.

### HasReturnActions

`func (o *ResponseReturnActionListResult) HasReturnActions() bool`

HasReturnActions returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseReturnActionListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseReturnActionListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseReturnActionListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseReturnActionListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseReturnActionListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseReturnActionListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseReturnActionListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseReturnActionListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseReturnActionListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseReturnActionListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseReturnActionListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseReturnActionListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


