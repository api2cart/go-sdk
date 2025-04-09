# OrderReturnUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnId** | **string** | Return ID | 
**OrderId** | Pointer to **string** | Defines the order id | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**ItemRestock** | Pointer to **bool** | Boolean, whether or not to add the line items back to the store inventory. | [optional] [default to false]
**ReturnStatusId** | Pointer to **string** | Defines return request status | [optional] 
**StaffNote** | Pointer to **string** | Specifies staff note | [optional] 
**Comment** | Pointer to **string** | Specifies return comment | [optional] 
**SendNotifications** | Pointer to **bool** | Send notifications to customer after order was created | [optional] [default to false]
**RejectReason** | Pointer to **string** | Defines return reject reason | [optional] 
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


