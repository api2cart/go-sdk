# ResponseOrderListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrdersCount** | Pointer to **NullableInt32** |  | [optional] 
**Order** | Pointer to [**[]Order**](Order.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseOrderListResult

`func NewResponseOrderListResult() *ResponseOrderListResult`

NewResponseOrderListResult instantiates a new ResponseOrderListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrderListResultWithDefaults

`func NewResponseOrderListResultWithDefaults() *ResponseOrderListResult`

NewResponseOrderListResultWithDefaults instantiates a new ResponseOrderListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrdersCount

`func (o *ResponseOrderListResult) GetOrdersCount() int32`

GetOrdersCount returns the OrdersCount field if non-nil, zero value otherwise.

### GetOrdersCountOk

`func (o *ResponseOrderListResult) GetOrdersCountOk() (*int32, bool)`

GetOrdersCountOk returns a tuple with the OrdersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersCount

`func (o *ResponseOrderListResult) SetOrdersCount(v int32)`

SetOrdersCount sets OrdersCount field to given value.

### HasOrdersCount

`func (o *ResponseOrderListResult) HasOrdersCount() bool`

HasOrdersCount returns a boolean if a field has been set.

### SetOrdersCountNil

`func (o *ResponseOrderListResult) SetOrdersCountNil(b bool)`

 SetOrdersCountNil sets the value for OrdersCount to be an explicit nil

### UnsetOrdersCount
`func (o *ResponseOrderListResult) UnsetOrdersCount()`

UnsetOrdersCount ensures that no value is present for OrdersCount, not even an explicit nil
### GetOrder

`func (o *ResponseOrderListResult) GetOrder() []Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ResponseOrderListResult) GetOrderOk() (*[]Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ResponseOrderListResult) SetOrder(v []Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ResponseOrderListResult) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseOrderListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseOrderListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseOrderListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseOrderListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseOrderListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseOrderListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseOrderListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseOrderListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseOrderListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseOrderListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseOrderListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseOrderListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


