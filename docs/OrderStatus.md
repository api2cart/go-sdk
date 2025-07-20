# OrderStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**History** | Pointer to [**[]OrderStatusHistoryItem**](OrderStatusHistoryItem.md) |  | [optional] 
**RefundInfo** | Pointer to [**NullableOrderStatusRefund**](OrderStatusRefund.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderStatus

`func NewOrderStatus() *OrderStatus`

NewOrderStatus instantiates a new OrderStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusWithDefaults

`func NewOrderStatusWithDefaults() *OrderStatus`

NewOrderStatusWithDefaults instantiates a new OrderStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrderStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHistory

`func (o *OrderStatus) GetHistory() []OrderStatusHistoryItem`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *OrderStatus) GetHistoryOk() (*[]OrderStatusHistoryItem, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *OrderStatus) SetHistory(v []OrderStatusHistoryItem)`

SetHistory sets History field to given value.

### HasHistory

`func (o *OrderStatus) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetRefundInfo

`func (o *OrderStatus) GetRefundInfo() OrderStatusRefund`

GetRefundInfo returns the RefundInfo field if non-nil, zero value otherwise.

### GetRefundInfoOk

`func (o *OrderStatus) GetRefundInfoOk() (*OrderStatusRefund, bool)`

GetRefundInfoOk returns a tuple with the RefundInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundInfo

`func (o *OrderStatus) SetRefundInfo(v OrderStatusRefund)`

SetRefundInfo sets RefundInfo field to given value.

### HasRefundInfo

`func (o *OrderStatus) HasRefundInfo() bool`

HasRefundInfo returns a boolean if a field has been set.

### SetRefundInfoNil

`func (o *OrderStatus) SetRefundInfoNil(b bool)`

 SetRefundInfoNil sets the value for RefundInfo to be an explicit nil

### UnsetRefundInfo
`func (o *OrderStatus) UnsetRefundInfo()`

UnsetRefundInfo ensures that no value is present for RefundInfo, not even an explicit nil
### GetAdditionalFields

`func (o *OrderStatus) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderStatus) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderStatus) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderStatus) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *OrderStatus) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *OrderStatus) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *OrderStatus) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderStatus) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderStatus) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderStatus) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *OrderStatus) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *OrderStatus) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


