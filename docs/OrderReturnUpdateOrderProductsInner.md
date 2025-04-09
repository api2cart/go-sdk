# OrderReturnUpdateOrderProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderProductId** | **string** | Defines which products from the order should be returned | 
**OrderProductQuantity** | **int32** | Defines how many product units from the order should be returned | 
**OrderProductStatus** | Pointer to **string** | Defines product return status | [optional] 
**OrderProductActionId** | **string** | Defines the ID of the return action | 

## Methods

### NewOrderReturnUpdateOrderProductsInner

`func NewOrderReturnUpdateOrderProductsInner(orderProductId string, orderProductQuantity int32, orderProductActionId string, ) *OrderReturnUpdateOrderProductsInner`

NewOrderReturnUpdateOrderProductsInner instantiates a new OrderReturnUpdateOrderProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderReturnUpdateOrderProductsInnerWithDefaults

`func NewOrderReturnUpdateOrderProductsInnerWithDefaults() *OrderReturnUpdateOrderProductsInner`

NewOrderReturnUpdateOrderProductsInnerWithDefaults instantiates a new OrderReturnUpdateOrderProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderProductId

`func (o *OrderReturnUpdateOrderProductsInner) GetOrderProductId() string`

GetOrderProductId returns the OrderProductId field if non-nil, zero value otherwise.

### GetOrderProductIdOk

`func (o *OrderReturnUpdateOrderProductsInner) GetOrderProductIdOk() (*string, bool)`

GetOrderProductIdOk returns a tuple with the OrderProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductId

`func (o *OrderReturnUpdateOrderProductsInner) SetOrderProductId(v string)`

SetOrderProductId sets OrderProductId field to given value.


### GetOrderProductQuantity

`func (o *OrderReturnUpdateOrderProductsInner) GetOrderProductQuantity() int32`

GetOrderProductQuantity returns the OrderProductQuantity field if non-nil, zero value otherwise.

### GetOrderProductQuantityOk

`func (o *OrderReturnUpdateOrderProductsInner) GetOrderProductQuantityOk() (*int32, bool)`

GetOrderProductQuantityOk returns a tuple with the OrderProductQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductQuantity

`func (o *OrderReturnUpdateOrderProductsInner) SetOrderProductQuantity(v int32)`

SetOrderProductQuantity sets OrderProductQuantity field to given value.


### GetOrderProductStatus

`func (o *OrderReturnUpdateOrderProductsInner) GetOrderProductStatus() string`

GetOrderProductStatus returns the OrderProductStatus field if non-nil, zero value otherwise.

### GetOrderProductStatusOk

`func (o *OrderReturnUpdateOrderProductsInner) GetOrderProductStatusOk() (*string, bool)`

GetOrderProductStatusOk returns a tuple with the OrderProductStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductStatus

`func (o *OrderReturnUpdateOrderProductsInner) SetOrderProductStatus(v string)`

SetOrderProductStatus sets OrderProductStatus field to given value.

### HasOrderProductStatus

`func (o *OrderReturnUpdateOrderProductsInner) HasOrderProductStatus() bool`

HasOrderProductStatus returns a boolean if a field has been set.

### GetOrderProductActionId

`func (o *OrderReturnUpdateOrderProductsInner) GetOrderProductActionId() string`

GetOrderProductActionId returns the OrderProductActionId field if non-nil, zero value otherwise.

### GetOrderProductActionIdOk

`func (o *OrderReturnUpdateOrderProductsInner) GetOrderProductActionIdOk() (*string, bool)`

GetOrderProductActionIdOk returns a tuple with the OrderProductActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductActionId

`func (o *OrderReturnUpdateOrderProductsInner) SetOrderProductActionId(v string)`

SetOrderProductActionId sets OrderProductActionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


