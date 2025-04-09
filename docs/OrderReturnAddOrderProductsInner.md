# OrderReturnAddOrderProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderProductId** | **string** | Defines which products from the order should be returned | 
**OrderProductQuantity** | **int32** | Defines how many product units from the order should be returned | 
**OrderProductReasonId** | **string** | Defines the ID of the return reason | 
**OrderProductActionId** | **string** | Defines the ID of the return action | 
**OrderProductCustomerComment** | Pointer to **string** | Defines the customer&#39;s comment for return | [optional] 
**OrderProductHandlingStatus** | Pointer to **string** | Defines handling status | [optional] 
**OrderProductCondition** | Pointer to **string** | Defines the product condition | [optional] 
**OrderProductReason** | Pointer to **string** | Defines return reason | [optional] 
**OrderProductStatus** | Pointer to **string** | Defines product return status | [optional] 

## Methods

### NewOrderReturnAddOrderProductsInner

`func NewOrderReturnAddOrderProductsInner(orderProductId string, orderProductQuantity int32, orderProductReasonId string, orderProductActionId string, ) *OrderReturnAddOrderProductsInner`

NewOrderReturnAddOrderProductsInner instantiates a new OrderReturnAddOrderProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderReturnAddOrderProductsInnerWithDefaults

`func NewOrderReturnAddOrderProductsInnerWithDefaults() *OrderReturnAddOrderProductsInner`

NewOrderReturnAddOrderProductsInnerWithDefaults instantiates a new OrderReturnAddOrderProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderProductId

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductId() string`

GetOrderProductId returns the OrderProductId field if non-nil, zero value otherwise.

### GetOrderProductIdOk

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductIdOk() (*string, bool)`

GetOrderProductIdOk returns a tuple with the OrderProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductId

`func (o *OrderReturnAddOrderProductsInner) SetOrderProductId(v string)`

SetOrderProductId sets OrderProductId field to given value.


### GetOrderProductQuantity

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductQuantity() int32`

GetOrderProductQuantity returns the OrderProductQuantity field if non-nil, zero value otherwise.

### GetOrderProductQuantityOk

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductQuantityOk() (*int32, bool)`

GetOrderProductQuantityOk returns a tuple with the OrderProductQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductQuantity

`func (o *OrderReturnAddOrderProductsInner) SetOrderProductQuantity(v int32)`

SetOrderProductQuantity sets OrderProductQuantity field to given value.


### GetOrderProductReasonId

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductReasonId() string`

GetOrderProductReasonId returns the OrderProductReasonId field if non-nil, zero value otherwise.

### GetOrderProductReasonIdOk

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductReasonIdOk() (*string, bool)`

GetOrderProductReasonIdOk returns a tuple with the OrderProductReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductReasonId

`func (o *OrderReturnAddOrderProductsInner) SetOrderProductReasonId(v string)`

SetOrderProductReasonId sets OrderProductReasonId field to given value.


### GetOrderProductActionId

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductActionId() string`

GetOrderProductActionId returns the OrderProductActionId field if non-nil, zero value otherwise.

### GetOrderProductActionIdOk

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductActionIdOk() (*string, bool)`

GetOrderProductActionIdOk returns a tuple with the OrderProductActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductActionId

`func (o *OrderReturnAddOrderProductsInner) SetOrderProductActionId(v string)`

SetOrderProductActionId sets OrderProductActionId field to given value.


### GetOrderProductCustomerComment

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductCustomerComment() string`

GetOrderProductCustomerComment returns the OrderProductCustomerComment field if non-nil, zero value otherwise.

### GetOrderProductCustomerCommentOk

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductCustomerCommentOk() (*string, bool)`

GetOrderProductCustomerCommentOk returns a tuple with the OrderProductCustomerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductCustomerComment

`func (o *OrderReturnAddOrderProductsInner) SetOrderProductCustomerComment(v string)`

SetOrderProductCustomerComment sets OrderProductCustomerComment field to given value.

### HasOrderProductCustomerComment

`func (o *OrderReturnAddOrderProductsInner) HasOrderProductCustomerComment() bool`

HasOrderProductCustomerComment returns a boolean if a field has been set.

### GetOrderProductHandlingStatus

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductHandlingStatus() string`

GetOrderProductHandlingStatus returns the OrderProductHandlingStatus field if non-nil, zero value otherwise.

### GetOrderProductHandlingStatusOk

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductHandlingStatusOk() (*string, bool)`

GetOrderProductHandlingStatusOk returns a tuple with the OrderProductHandlingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductHandlingStatus

`func (o *OrderReturnAddOrderProductsInner) SetOrderProductHandlingStatus(v string)`

SetOrderProductHandlingStatus sets OrderProductHandlingStatus field to given value.

### HasOrderProductHandlingStatus

`func (o *OrderReturnAddOrderProductsInner) HasOrderProductHandlingStatus() bool`

HasOrderProductHandlingStatus returns a boolean if a field has been set.

### GetOrderProductCondition

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductCondition() string`

GetOrderProductCondition returns the OrderProductCondition field if non-nil, zero value otherwise.

### GetOrderProductConditionOk

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductConditionOk() (*string, bool)`

GetOrderProductConditionOk returns a tuple with the OrderProductCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductCondition

`func (o *OrderReturnAddOrderProductsInner) SetOrderProductCondition(v string)`

SetOrderProductCondition sets OrderProductCondition field to given value.

### HasOrderProductCondition

`func (o *OrderReturnAddOrderProductsInner) HasOrderProductCondition() bool`

HasOrderProductCondition returns a boolean if a field has been set.

### GetOrderProductReason

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductReason() string`

GetOrderProductReason returns the OrderProductReason field if non-nil, zero value otherwise.

### GetOrderProductReasonOk

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductReasonOk() (*string, bool)`

GetOrderProductReasonOk returns a tuple with the OrderProductReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductReason

`func (o *OrderReturnAddOrderProductsInner) SetOrderProductReason(v string)`

SetOrderProductReason sets OrderProductReason field to given value.

### HasOrderProductReason

`func (o *OrderReturnAddOrderProductsInner) HasOrderProductReason() bool`

HasOrderProductReason returns a boolean if a field has been set.

### GetOrderProductStatus

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductStatus() string`

GetOrderProductStatus returns the OrderProductStatus field if non-nil, zero value otherwise.

### GetOrderProductStatusOk

`func (o *OrderReturnAddOrderProductsInner) GetOrderProductStatusOk() (*string, bool)`

GetOrderProductStatusOk returns a tuple with the OrderProductStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductStatus

`func (o *OrderReturnAddOrderProductsInner) SetOrderProductStatus(v string)`

SetOrderProductStatus sets OrderProductStatus field to given value.

### HasOrderProductStatus

`func (o *OrderReturnAddOrderProductsInner) HasOrderProductStatus() bool`

HasOrderProductStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


