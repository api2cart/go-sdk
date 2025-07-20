# OrderStatusRefundItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **NullableString** |  | [optional] 
**VariantId** | Pointer to **NullableString** |  | [optional] 
**OrderProductId** | Pointer to **NullableString** |  | [optional] 
**Qty** | Pointer to **NullableFloat32** |  | [optional] 
**Refund** | Pointer to **NullableFloat32** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderStatusRefundItem

`func NewOrderStatusRefundItem() *OrderStatusRefundItem`

NewOrderStatusRefundItem instantiates a new OrderStatusRefundItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusRefundItemWithDefaults

`func NewOrderStatusRefundItemWithDefaults() *OrderStatusRefundItem`

NewOrderStatusRefundItemWithDefaults instantiates a new OrderStatusRefundItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *OrderStatusRefundItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *OrderStatusRefundItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *OrderStatusRefundItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *OrderStatusRefundItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *OrderStatusRefundItem) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *OrderStatusRefundItem) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetVariantId

`func (o *OrderStatusRefundItem) GetVariantId() string`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *OrderStatusRefundItem) GetVariantIdOk() (*string, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *OrderStatusRefundItem) SetVariantId(v string)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *OrderStatusRefundItem) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.

### SetVariantIdNil

`func (o *OrderStatusRefundItem) SetVariantIdNil(b bool)`

 SetVariantIdNil sets the value for VariantId to be an explicit nil

### UnsetVariantId
`func (o *OrderStatusRefundItem) UnsetVariantId()`

UnsetVariantId ensures that no value is present for VariantId, not even an explicit nil
### GetOrderProductId

`func (o *OrderStatusRefundItem) GetOrderProductId() string`

GetOrderProductId returns the OrderProductId field if non-nil, zero value otherwise.

### GetOrderProductIdOk

`func (o *OrderStatusRefundItem) GetOrderProductIdOk() (*string, bool)`

GetOrderProductIdOk returns a tuple with the OrderProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductId

`func (o *OrderStatusRefundItem) SetOrderProductId(v string)`

SetOrderProductId sets OrderProductId field to given value.

### HasOrderProductId

`func (o *OrderStatusRefundItem) HasOrderProductId() bool`

HasOrderProductId returns a boolean if a field has been set.

### SetOrderProductIdNil

`func (o *OrderStatusRefundItem) SetOrderProductIdNil(b bool)`

 SetOrderProductIdNil sets the value for OrderProductId to be an explicit nil

### UnsetOrderProductId
`func (o *OrderStatusRefundItem) UnsetOrderProductId()`

UnsetOrderProductId ensures that no value is present for OrderProductId, not even an explicit nil
### GetQty

`func (o *OrderStatusRefundItem) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *OrderStatusRefundItem) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *OrderStatusRefundItem) SetQty(v float32)`

SetQty sets Qty field to given value.

### HasQty

`func (o *OrderStatusRefundItem) HasQty() bool`

HasQty returns a boolean if a field has been set.

### SetQtyNil

`func (o *OrderStatusRefundItem) SetQtyNil(b bool)`

 SetQtyNil sets the value for Qty to be an explicit nil

### UnsetQty
`func (o *OrderStatusRefundItem) UnsetQty()`

UnsetQty ensures that no value is present for Qty, not even an explicit nil
### GetRefund

`func (o *OrderStatusRefundItem) GetRefund() float32`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *OrderStatusRefundItem) GetRefundOk() (*float32, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *OrderStatusRefundItem) SetRefund(v float32)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *OrderStatusRefundItem) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### SetRefundNil

`func (o *OrderStatusRefundItem) SetRefundNil(b bool)`

 SetRefundNil sets the value for Refund to be an explicit nil

### UnsetRefund
`func (o *OrderStatusRefundItem) UnsetRefund()`

UnsetRefund ensures that no value is present for Refund, not even an explicit nil
### GetAdditionalFields

`func (o *OrderStatusRefundItem) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderStatusRefundItem) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderStatusRefundItem) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderStatusRefundItem) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *OrderStatusRefundItem) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *OrderStatusRefundItem) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *OrderStatusRefundItem) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderStatusRefundItem) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderStatusRefundItem) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderStatusRefundItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *OrderStatusRefundItem) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *OrderStatusRefundItem) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


