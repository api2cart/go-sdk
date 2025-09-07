# OrderCalculateOrderItemInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderItemId** | **string** | Defines orders specified by order item id | 
**OrderItemQuantity** | **int32** | Defines orders specified by order item quantity | 
**OrderItemVariantId** | Pointer to **string** | Ordered product variant. Where x is order item ID | [optional] 
**OrderItemOption** | Pointer to [**[]OrderCalculateOrderItemInnerOrderItemOptionInner**](OrderCalculateOrderItemInnerOrderItemOptionInner.md) |  | [optional] 

## Methods

### NewOrderCalculateOrderItemInner

`func NewOrderCalculateOrderItemInner(orderItemId string, orderItemQuantity int32, ) *OrderCalculateOrderItemInner`

NewOrderCalculateOrderItemInner instantiates a new OrderCalculateOrderItemInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCalculateOrderItemInnerWithDefaults

`func NewOrderCalculateOrderItemInnerWithDefaults() *OrderCalculateOrderItemInner`

NewOrderCalculateOrderItemInnerWithDefaults instantiates a new OrderCalculateOrderItemInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderItemId

`func (o *OrderCalculateOrderItemInner) GetOrderItemId() string`

GetOrderItemId returns the OrderItemId field if non-nil, zero value otherwise.

### GetOrderItemIdOk

`func (o *OrderCalculateOrderItemInner) GetOrderItemIdOk() (*string, bool)`

GetOrderItemIdOk returns a tuple with the OrderItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemId

`func (o *OrderCalculateOrderItemInner) SetOrderItemId(v string)`

SetOrderItemId sets OrderItemId field to given value.


### GetOrderItemQuantity

`func (o *OrderCalculateOrderItemInner) GetOrderItemQuantity() int32`

GetOrderItemQuantity returns the OrderItemQuantity field if non-nil, zero value otherwise.

### GetOrderItemQuantityOk

`func (o *OrderCalculateOrderItemInner) GetOrderItemQuantityOk() (*int32, bool)`

GetOrderItemQuantityOk returns a tuple with the OrderItemQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemQuantity

`func (o *OrderCalculateOrderItemInner) SetOrderItemQuantity(v int32)`

SetOrderItemQuantity sets OrderItemQuantity field to given value.


### GetOrderItemVariantId

`func (o *OrderCalculateOrderItemInner) GetOrderItemVariantId() string`

GetOrderItemVariantId returns the OrderItemVariantId field if non-nil, zero value otherwise.

### GetOrderItemVariantIdOk

`func (o *OrderCalculateOrderItemInner) GetOrderItemVariantIdOk() (*string, bool)`

GetOrderItemVariantIdOk returns a tuple with the OrderItemVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemVariantId

`func (o *OrderCalculateOrderItemInner) SetOrderItemVariantId(v string)`

SetOrderItemVariantId sets OrderItemVariantId field to given value.

### HasOrderItemVariantId

`func (o *OrderCalculateOrderItemInner) HasOrderItemVariantId() bool`

HasOrderItemVariantId returns a boolean if a field has been set.

### GetOrderItemOption

`func (o *OrderCalculateOrderItemInner) GetOrderItemOption() []OrderCalculateOrderItemInnerOrderItemOptionInner`

GetOrderItemOption returns the OrderItemOption field if non-nil, zero value otherwise.

### GetOrderItemOptionOk

`func (o *OrderCalculateOrderItemInner) GetOrderItemOptionOk() (*[]OrderCalculateOrderItemInnerOrderItemOptionInner, bool)`

GetOrderItemOptionOk returns a tuple with the OrderItemOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemOption

`func (o *OrderCalculateOrderItemInner) SetOrderItemOption(v []OrderCalculateOrderItemInnerOrderItemOptionInner)`

SetOrderItemOption sets OrderItemOption field to given value.

### HasOrderItemOption

`func (o *OrderCalculateOrderItemInner) HasOrderItemOption() bool`

HasOrderItemOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


