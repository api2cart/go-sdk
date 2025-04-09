# OrderPreestimateShippingListOrderItemInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderItemId** | **string** | Defines orders specified by order item id | 
**OrderItemModel** | Pointer to **string** | Defines orders specified by order item model | [optional] 
**OrderItemQuantity** | **int32** | Defines orders specified by order item quantity | 
**OrderItemWeight** | Pointer to **float32** | Defines orders specified by order item weight | [optional] 
**OrderItemVariantId** | Pointer to **string** | Ordered product variant. Where x is order item ID | [optional] 
**OrderItemOption** | Pointer to [**[]OrderPreestimateShippingListOrderItemInnerOrderItemOptionInner**](OrderPreestimateShippingListOrderItemInnerOrderItemOptionInner.md) |  | [optional] 

## Methods

### NewOrderPreestimateShippingListOrderItemInner

`func NewOrderPreestimateShippingListOrderItemInner(orderItemId string, orderItemQuantity int32, ) *OrderPreestimateShippingListOrderItemInner`

NewOrderPreestimateShippingListOrderItemInner instantiates a new OrderPreestimateShippingListOrderItemInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPreestimateShippingListOrderItemInnerWithDefaults

`func NewOrderPreestimateShippingListOrderItemInnerWithDefaults() *OrderPreestimateShippingListOrderItemInner`

NewOrderPreestimateShippingListOrderItemInnerWithDefaults instantiates a new OrderPreestimateShippingListOrderItemInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderItemId

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemId() string`

GetOrderItemId returns the OrderItemId field if non-nil, zero value otherwise.

### GetOrderItemIdOk

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemIdOk() (*string, bool)`

GetOrderItemIdOk returns a tuple with the OrderItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemId

`func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemId(v string)`

SetOrderItemId sets OrderItemId field to given value.


### GetOrderItemModel

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemModel() string`

GetOrderItemModel returns the OrderItemModel field if non-nil, zero value otherwise.

### GetOrderItemModelOk

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemModelOk() (*string, bool)`

GetOrderItemModelOk returns a tuple with the OrderItemModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemModel

`func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemModel(v string)`

SetOrderItemModel sets OrderItemModel field to given value.

### HasOrderItemModel

`func (o *OrderPreestimateShippingListOrderItemInner) HasOrderItemModel() bool`

HasOrderItemModel returns a boolean if a field has been set.

### GetOrderItemQuantity

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemQuantity() int32`

GetOrderItemQuantity returns the OrderItemQuantity field if non-nil, zero value otherwise.

### GetOrderItemQuantityOk

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemQuantityOk() (*int32, bool)`

GetOrderItemQuantityOk returns a tuple with the OrderItemQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemQuantity

`func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemQuantity(v int32)`

SetOrderItemQuantity sets OrderItemQuantity field to given value.


### GetOrderItemWeight

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemWeight() float32`

GetOrderItemWeight returns the OrderItemWeight field if non-nil, zero value otherwise.

### GetOrderItemWeightOk

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemWeightOk() (*float32, bool)`

GetOrderItemWeightOk returns a tuple with the OrderItemWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemWeight

`func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemWeight(v float32)`

SetOrderItemWeight sets OrderItemWeight field to given value.

### HasOrderItemWeight

`func (o *OrderPreestimateShippingListOrderItemInner) HasOrderItemWeight() bool`

HasOrderItemWeight returns a boolean if a field has been set.

### GetOrderItemVariantId

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemVariantId() string`

GetOrderItemVariantId returns the OrderItemVariantId field if non-nil, zero value otherwise.

### GetOrderItemVariantIdOk

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemVariantIdOk() (*string, bool)`

GetOrderItemVariantIdOk returns a tuple with the OrderItemVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemVariantId

`func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemVariantId(v string)`

SetOrderItemVariantId sets OrderItemVariantId field to given value.

### HasOrderItemVariantId

`func (o *OrderPreestimateShippingListOrderItemInner) HasOrderItemVariantId() bool`

HasOrderItemVariantId returns a boolean if a field has been set.

### GetOrderItemOption

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemOption() []OrderPreestimateShippingListOrderItemInnerOrderItemOptionInner`

GetOrderItemOption returns the OrderItemOption field if non-nil, zero value otherwise.

### GetOrderItemOptionOk

`func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemOptionOk() (*[]OrderPreestimateShippingListOrderItemInnerOrderItemOptionInner, bool)`

GetOrderItemOptionOk returns a tuple with the OrderItemOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemOption

`func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemOption(v []OrderPreestimateShippingListOrderItemInnerOrderItemOptionInner)`

SetOrderItemOption sets OrderItemOption field to given value.

### HasOrderItemOption

`func (o *OrderPreestimateShippingListOrderItemInner) HasOrderItemOption() bool`

HasOrderItemOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


