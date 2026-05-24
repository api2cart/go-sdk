# OrderReturnUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnId** | **string** | Return ID | 
**OrderId** | Pointer to **string** | Defines the order id | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**ItemRestock** | Pointer to **bool** | Boolean, whether or not to add the line items back to the store inventory. | [optional] [default to false]
**ReturnStatusId** | Pointer to **string** | Defines return request status | [optional] 
**ReturnReasonId** | Pointer to **string** | Defines return request reason | [optional] 
**ReturnActionId** | Pointer to **string** | Defines return request action | [optional] 
**StaffNote** | Pointer to **string** | Specifies staff note | [optional] 
**Comment** | Pointer to **string** | Specifies return comment | [optional] 
**Message** | Pointer to **string** | Customer-visible message attached to the return request (updated). | [optional] 
**SendNotifications** | Pointer to **bool** | Send notifications to customer after order was created | [optional] [default to false]
**RejectReason** | Pointer to **string** | Defines return reject reason | [optional] 
**ReturnAction** | Pointer to **string** | Defines return request action | [optional] 
**ReturnReason** | Pointer to **string** | Defines return request reason | [optional] 
**IsOnline** | Pointer to **bool** | Indicates whether refund type is online | [optional] [default to false]
**FeePrice** | Pointer to **float32** | Specifies refund&#39;s fee price | [optional] 
**ShippingPrice** | Pointer to **float32** | Specifies order&#39;s shipping price | [optional] [default to 0]
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 
**OrderProducts** | [**[]OrderReturnUpdateOrderProductsInner**](OrderReturnUpdateOrderProductsInner.md) |  | 

## Methods

### NewOrderReturnUpdate

`func NewOrderReturnUpdate(returnId string, orderProducts []OrderReturnUpdateOrderProductsInner, ) *OrderReturnUpdate`

NewOrderReturnUpdate instantiates a new OrderReturnUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderReturnUpdateWithDefaults

`func NewOrderReturnUpdateWithDefaults() *OrderReturnUpdate`

NewOrderReturnUpdateWithDefaults instantiates a new OrderReturnUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnId

`func (o *OrderReturnUpdate) GetReturnId() string`

GetReturnId returns the ReturnId field if non-nil, zero value otherwise.

### GetReturnIdOk

`func (o *OrderReturnUpdate) GetReturnIdOk() (*string, bool)`

GetReturnIdOk returns a tuple with the ReturnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnId

`func (o *OrderReturnUpdate) SetReturnId(v string)`

SetReturnId sets ReturnId field to given value.


### GetOrderId

`func (o *OrderReturnUpdate) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderReturnUpdate) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderReturnUpdate) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderReturnUpdate) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStoreId

`func (o *OrderReturnUpdate) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderReturnUpdate) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderReturnUpdate) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderReturnUpdate) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetItemRestock

`func (o *OrderReturnUpdate) GetItemRestock() bool`

GetItemRestock returns the ItemRestock field if non-nil, zero value otherwise.

### GetItemRestockOk

`func (o *OrderReturnUpdate) GetItemRestockOk() (*bool, bool)`

GetItemRestockOk returns a tuple with the ItemRestock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRestock

`func (o *OrderReturnUpdate) SetItemRestock(v bool)`

SetItemRestock sets ItemRestock field to given value.

### HasItemRestock

`func (o *OrderReturnUpdate) HasItemRestock() bool`

HasItemRestock returns a boolean if a field has been set.

### GetReturnStatusId

`func (o *OrderReturnUpdate) GetReturnStatusId() string`

GetReturnStatusId returns the ReturnStatusId field if non-nil, zero value otherwise.

### GetReturnStatusIdOk

`func (o *OrderReturnUpdate) GetReturnStatusIdOk() (*string, bool)`

GetReturnStatusIdOk returns a tuple with the ReturnStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnStatusId

`func (o *OrderReturnUpdate) SetReturnStatusId(v string)`

SetReturnStatusId sets ReturnStatusId field to given value.

### HasReturnStatusId

`func (o *OrderReturnUpdate) HasReturnStatusId() bool`

HasReturnStatusId returns a boolean if a field has been set.

### GetReturnReasonId

`func (o *OrderReturnUpdate) GetReturnReasonId() string`

GetReturnReasonId returns the ReturnReasonId field if non-nil, zero value otherwise.

### GetReturnReasonIdOk

`func (o *OrderReturnUpdate) GetReturnReasonIdOk() (*string, bool)`

GetReturnReasonIdOk returns a tuple with the ReturnReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReasonId

`func (o *OrderReturnUpdate) SetReturnReasonId(v string)`

SetReturnReasonId sets ReturnReasonId field to given value.

### HasReturnReasonId

`func (o *OrderReturnUpdate) HasReturnReasonId() bool`

HasReturnReasonId returns a boolean if a field has been set.

### GetReturnActionId

`func (o *OrderReturnUpdate) GetReturnActionId() string`

GetReturnActionId returns the ReturnActionId field if non-nil, zero value otherwise.

### GetReturnActionIdOk

`func (o *OrderReturnUpdate) GetReturnActionIdOk() (*string, bool)`

GetReturnActionIdOk returns a tuple with the ReturnActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnActionId

`func (o *OrderReturnUpdate) SetReturnActionId(v string)`

SetReturnActionId sets ReturnActionId field to given value.

### HasReturnActionId

`func (o *OrderReturnUpdate) HasReturnActionId() bool`

HasReturnActionId returns a boolean if a field has been set.

### GetStaffNote

`func (o *OrderReturnUpdate) GetStaffNote() string`

GetStaffNote returns the StaffNote field if non-nil, zero value otherwise.

### GetStaffNoteOk

`func (o *OrderReturnUpdate) GetStaffNoteOk() (*string, bool)`

GetStaffNoteOk returns a tuple with the StaffNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaffNote

`func (o *OrderReturnUpdate) SetStaffNote(v string)`

SetStaffNote sets StaffNote field to given value.

### HasStaffNote

`func (o *OrderReturnUpdate) HasStaffNote() bool`

HasStaffNote returns a boolean if a field has been set.

### GetComment

`func (o *OrderReturnUpdate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *OrderReturnUpdate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *OrderReturnUpdate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *OrderReturnUpdate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetMessage

`func (o *OrderReturnUpdate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OrderReturnUpdate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OrderReturnUpdate) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OrderReturnUpdate) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSendNotifications

`func (o *OrderReturnUpdate) GetSendNotifications() bool`

GetSendNotifications returns the SendNotifications field if non-nil, zero value otherwise.

### GetSendNotificationsOk

`func (o *OrderReturnUpdate) GetSendNotificationsOk() (*bool, bool)`

GetSendNotificationsOk returns a tuple with the SendNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotifications

`func (o *OrderReturnUpdate) SetSendNotifications(v bool)`

SetSendNotifications sets SendNotifications field to given value.

### HasSendNotifications

`func (o *OrderReturnUpdate) HasSendNotifications() bool`

HasSendNotifications returns a boolean if a field has been set.

### GetRejectReason

`func (o *OrderReturnUpdate) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *OrderReturnUpdate) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *OrderReturnUpdate) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *OrderReturnUpdate) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetReturnAction

`func (o *OrderReturnUpdate) GetReturnAction() string`

GetReturnAction returns the ReturnAction field if non-nil, zero value otherwise.

### GetReturnActionOk

`func (o *OrderReturnUpdate) GetReturnActionOk() (*string, bool)`

GetReturnActionOk returns a tuple with the ReturnAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAction

`func (o *OrderReturnUpdate) SetReturnAction(v string)`

SetReturnAction sets ReturnAction field to given value.

### HasReturnAction

`func (o *OrderReturnUpdate) HasReturnAction() bool`

HasReturnAction returns a boolean if a field has been set.

### GetReturnReason

`func (o *OrderReturnUpdate) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *OrderReturnUpdate) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *OrderReturnUpdate) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *OrderReturnUpdate) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetIsOnline

`func (o *OrderReturnUpdate) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *OrderReturnUpdate) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *OrderReturnUpdate) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.

### HasIsOnline

`func (o *OrderReturnUpdate) HasIsOnline() bool`

HasIsOnline returns a boolean if a field has been set.

### GetFeePrice

`func (o *OrderReturnUpdate) GetFeePrice() float32`

GetFeePrice returns the FeePrice field if non-nil, zero value otherwise.

### GetFeePriceOk

`func (o *OrderReturnUpdate) GetFeePriceOk() (*float32, bool)`

GetFeePriceOk returns a tuple with the FeePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePrice

`func (o *OrderReturnUpdate) SetFeePrice(v float32)`

SetFeePrice sets FeePrice field to given value.

### HasFeePrice

`func (o *OrderReturnUpdate) HasFeePrice() bool`

HasFeePrice returns a boolean if a field has been set.

### GetShippingPrice

`func (o *OrderReturnUpdate) GetShippingPrice() float32`

GetShippingPrice returns the ShippingPrice field if non-nil, zero value otherwise.

### GetShippingPriceOk

`func (o *OrderReturnUpdate) GetShippingPriceOk() (*float32, bool)`

GetShippingPriceOk returns a tuple with the ShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPrice

`func (o *OrderReturnUpdate) SetShippingPrice(v float32)`

SetShippingPrice sets ShippingPrice field to given value.

### HasShippingPrice

`func (o *OrderReturnUpdate) HasShippingPrice() bool`

HasShippingPrice returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *OrderReturnUpdate) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *OrderReturnUpdate) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *OrderReturnUpdate) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *OrderReturnUpdate) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetOrderProducts

`func (o *OrderReturnUpdate) GetOrderProducts() []OrderReturnUpdateOrderProductsInner`

GetOrderProducts returns the OrderProducts field if non-nil, zero value otherwise.

### GetOrderProductsOk

`func (o *OrderReturnUpdate) GetOrderProductsOk() (*[]OrderReturnUpdateOrderProductsInner, bool)`

GetOrderProductsOk returns a tuple with the OrderProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProducts

`func (o *OrderReturnUpdate) SetOrderProducts(v []OrderReturnUpdateOrderProductsInner)`

SetOrderProducts sets OrderProducts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


