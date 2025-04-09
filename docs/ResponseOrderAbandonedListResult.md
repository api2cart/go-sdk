# ResponseOrderAbandonedListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**[]OrderAbandoned**](OrderAbandoned.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseOrderAbandonedListResult

`func NewResponseOrderAbandonedListResult() *ResponseOrderAbandonedListResult`

NewResponseOrderAbandonedListResult instantiates a new ResponseOrderAbandonedListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrderAbandonedListResultWithDefaults

`func NewResponseOrderAbandonedListResultWithDefaults() *ResponseOrderAbandonedListResult`

NewResponseOrderAbandonedListResultWithDefaults instantiates a new ResponseOrderAbandonedListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *ResponseOrderAbandonedListResult) GetOrder() []OrderAbandoned`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ResponseOrderAbandonedListResult) GetOrderOk() (*[]OrderAbandoned, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ResponseOrderAbandonedListResult) SetOrder(v []OrderAbandoned)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ResponseOrderAbandonedListResult) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseOrderAbandonedListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseOrderAbandonedListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseOrderAbandonedListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseOrderAbandonedListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseOrderAbandonedListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseOrderAbandonedListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseOrderAbandonedListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseOrderAbandonedListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


