# ResponseOrderFinancialStatusListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**OrderFinancialStatuses** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseOrderFinancialStatusListResult

`func NewResponseOrderFinancialStatusListResult() *ResponseOrderFinancialStatusListResult`

NewResponseOrderFinancialStatusListResult instantiates a new ResponseOrderFinancialStatusListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrderFinancialStatusListResultWithDefaults

`func NewResponseOrderFinancialStatusListResultWithDefaults() *ResponseOrderFinancialStatusListResult`

NewResponseOrderFinancialStatusListResultWithDefaults instantiates a new ResponseOrderFinancialStatusListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ResponseOrderFinancialStatusListResult) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResponseOrderFinancialStatusListResult) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResponseOrderFinancialStatusListResult) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ResponseOrderFinancialStatusListResult) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ResponseOrderFinancialStatusListResult) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ResponseOrderFinancialStatusListResult) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetOrderFinancialStatuses

`func (o *ResponseOrderFinancialStatusListResult) GetOrderFinancialStatuses() []Status`

GetOrderFinancialStatuses returns the OrderFinancialStatuses field if non-nil, zero value otherwise.

### GetOrderFinancialStatusesOk

`func (o *ResponseOrderFinancialStatusListResult) GetOrderFinancialStatusesOk() (*[]Status, bool)`

GetOrderFinancialStatusesOk returns a tuple with the OrderFinancialStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFinancialStatuses

`func (o *ResponseOrderFinancialStatusListResult) SetOrderFinancialStatuses(v []Status)`

SetOrderFinancialStatuses sets OrderFinancialStatuses field to given value.

### HasOrderFinancialStatuses

`func (o *ResponseOrderFinancialStatusListResult) HasOrderFinancialStatuses() bool`

HasOrderFinancialStatuses returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseOrderFinancialStatusListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseOrderFinancialStatusListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseOrderFinancialStatusListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseOrderFinancialStatusListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseOrderFinancialStatusListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseOrderFinancialStatusListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseOrderFinancialStatusListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseOrderFinancialStatusListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseOrderFinancialStatusListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseOrderFinancialStatusListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseOrderFinancialStatusListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseOrderFinancialStatusListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


