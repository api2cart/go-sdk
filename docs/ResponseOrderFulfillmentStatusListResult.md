# ResponseOrderFulfillmentStatusListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**OrderFulfillmentStatuses** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseOrderFulfillmentStatusListResult

`func NewResponseOrderFulfillmentStatusListResult() *ResponseOrderFulfillmentStatusListResult`

NewResponseOrderFulfillmentStatusListResult instantiates a new ResponseOrderFulfillmentStatusListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrderFulfillmentStatusListResultWithDefaults

`func NewResponseOrderFulfillmentStatusListResultWithDefaults() *ResponseOrderFulfillmentStatusListResult`

NewResponseOrderFulfillmentStatusListResultWithDefaults instantiates a new ResponseOrderFulfillmentStatusListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ResponseOrderFulfillmentStatusListResult) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResponseOrderFulfillmentStatusListResult) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResponseOrderFulfillmentStatusListResult) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ResponseOrderFulfillmentStatusListResult) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ResponseOrderFulfillmentStatusListResult) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ResponseOrderFulfillmentStatusListResult) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetOrderFulfillmentStatuses

`func (o *ResponseOrderFulfillmentStatusListResult) GetOrderFulfillmentStatuses() []Status`

GetOrderFulfillmentStatuses returns the OrderFulfillmentStatuses field if non-nil, zero value otherwise.

### GetOrderFulfillmentStatusesOk

`func (o *ResponseOrderFulfillmentStatusListResult) GetOrderFulfillmentStatusesOk() (*[]Status, bool)`

GetOrderFulfillmentStatusesOk returns a tuple with the OrderFulfillmentStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFulfillmentStatuses

`func (o *ResponseOrderFulfillmentStatusListResult) SetOrderFulfillmentStatuses(v []Status)`

SetOrderFulfillmentStatuses sets OrderFulfillmentStatuses field to given value.

### HasOrderFulfillmentStatuses

`func (o *ResponseOrderFulfillmentStatusListResult) HasOrderFulfillmentStatuses() bool`

HasOrderFulfillmentStatuses returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseOrderFulfillmentStatusListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseOrderFulfillmentStatusListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseOrderFulfillmentStatusListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseOrderFulfillmentStatusListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseOrderFulfillmentStatusListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseOrderFulfillmentStatusListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseOrderFulfillmentStatusListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseOrderFulfillmentStatusListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseOrderFulfillmentStatusListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseOrderFulfillmentStatusListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseOrderFulfillmentStatusListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseOrderFulfillmentStatusListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


