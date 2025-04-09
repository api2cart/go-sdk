# OrderReturnAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Defines the order id | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**ReturnStatusId** | **string** | Defines return request status | 
**ReturnActionId** | **string** | Defines return request action | 
**ReturnReasonId** | **string** | Defines return request reason | 
**ReturnReason** | Pointer to **string** | Defines return request reason | [optional] 
**ItemRestock** | Pointer to **bool** | Boolean, whether or not to add the line items back to the store inventory. | [optional] [default to false]
**StaffNote** | Pointer to **string** | Specifies staff note | [optional] 
**Comment** | Pointer to **string** | Specifies return comment | [optional] 
**SendNotifications** | Pointer to **bool** | Send notifications to customer after order was created | [optional] [default to false]
**RejectReason** | Pointer to **string** | Defines return reject reason | [optional] 
**OrderProducts** | [**[]OrderReturnAddOrderProductsInner**](OrderReturnAddOrderProductsInner.md) |  | 

## Methods

### NewOrderReturnAdd

`func NewOrderReturnAdd(returnStatusId string, returnActionId string, returnReasonId string, orderProducts []OrderReturnAddOrderProductsInner, ) *OrderReturnAdd`

NewOrderReturnAdd instantiates a new OrderReturnAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderReturnAddWithDefaults

`func NewOrderReturnAddWithDefaults() *OrderReturnAdd`

NewOrderReturnAddWithDefaults instantiates a new OrderReturnAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *OrderReturnAdd) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderReturnAdd) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderReturnAdd) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderReturnAdd) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStoreId

`func (o *OrderReturnAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderReturnAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderReturnAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderReturnAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetReturnStatusId

`func (o *OrderReturnAdd) GetReturnStatusId() string`

GetReturnStatusId returns the ReturnStatusId field if non-nil, zero value otherwise.

### GetReturnStatusIdOk

`func (o *OrderReturnAdd) GetReturnStatusIdOk() (*string, bool)`

GetReturnStatusIdOk returns a tuple with the ReturnStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnStatusId

`func (o *OrderReturnAdd) SetReturnStatusId(v string)`

SetReturnStatusId sets ReturnStatusId field to given value.


### GetReturnActionId

`func (o *OrderReturnAdd) GetReturnActionId() string`

GetReturnActionId returns the ReturnActionId field if non-nil, zero value otherwise.

### GetReturnActionIdOk

`func (o *OrderReturnAdd) GetReturnActionIdOk() (*string, bool)`

GetReturnActionIdOk returns a tuple with the ReturnActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnActionId

`func (o *OrderReturnAdd) SetReturnActionId(v string)`

SetReturnActionId sets ReturnActionId field to given value.


### GetReturnReasonId

`func (o *OrderReturnAdd) GetReturnReasonId() string`

GetReturnReasonId returns the ReturnReasonId field if non-nil, zero value otherwise.

### GetReturnReasonIdOk

`func (o *OrderReturnAdd) GetReturnReasonIdOk() (*string, bool)`

GetReturnReasonIdOk returns a tuple with the ReturnReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReasonId

`func (o *OrderReturnAdd) SetReturnReasonId(v string)`

SetReturnReasonId sets ReturnReasonId field to given value.


### GetReturnReason

`func (o *OrderReturnAdd) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *OrderReturnAdd) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *OrderReturnAdd) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *OrderReturnAdd) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetItemRestock

`func (o *OrderReturnAdd) GetItemRestock() bool`

GetItemRestock returns the ItemRestock field if non-nil, zero value otherwise.

### GetItemRestockOk

`func (o *OrderReturnAdd) GetItemRestockOk() (*bool, bool)`

GetItemRestockOk returns a tuple with the ItemRestock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRestock

`func (o *OrderReturnAdd) SetItemRestock(v bool)`

SetItemRestock sets ItemRestock field to given value.

### HasItemRestock

`func (o *OrderReturnAdd) HasItemRestock() bool`

HasItemRestock returns a boolean if a field has been set.

### GetStaffNote

`func (o *OrderReturnAdd) GetStaffNote() string`

GetStaffNote returns the StaffNote field if non-nil, zero value otherwise.

### GetStaffNoteOk

`func (o *OrderReturnAdd) GetStaffNoteOk() (*string, bool)`

GetStaffNoteOk returns a tuple with the StaffNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaffNote

`func (o *OrderReturnAdd) SetStaffNote(v string)`

SetStaffNote sets StaffNote field to given value.

### HasStaffNote

`func (o *OrderReturnAdd) HasStaffNote() bool`

HasStaffNote returns a boolean if a field has been set.

### GetComment

`func (o *OrderReturnAdd) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *OrderReturnAdd) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *OrderReturnAdd) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *OrderReturnAdd) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetSendNotifications

`func (o *OrderReturnAdd) GetSendNotifications() bool`

GetSendNotifications returns the SendNotifications field if non-nil, zero value otherwise.

### GetSendNotificationsOk

`func (o *OrderReturnAdd) GetSendNotificationsOk() (*bool, bool)`

GetSendNotificationsOk returns a tuple with the SendNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotifications

`func (o *OrderReturnAdd) SetSendNotifications(v bool)`

SetSendNotifications sets SendNotifications field to given value.

### HasSendNotifications

`func (o *OrderReturnAdd) HasSendNotifications() bool`

HasSendNotifications returns a boolean if a field has been set.

### GetRejectReason

`func (o *OrderReturnAdd) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *OrderReturnAdd) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *OrderReturnAdd) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *OrderReturnAdd) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetOrderProducts

`func (o *OrderReturnAdd) GetOrderProducts() []OrderReturnAddOrderProductsInner`

GetOrderProducts returns the OrderProducts field if non-nil, zero value otherwise.

### GetOrderProductsOk

`func (o *OrderReturnAdd) GetOrderProductsOk() (*[]OrderReturnAddOrderProductsInner, bool)`

GetOrderProductsOk returns a tuple with the OrderProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProducts

`func (o *OrderReturnAdd) SetOrderProducts(v []OrderReturnAddOrderProductsInner)`

SetOrderProducts sets OrderProducts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


