# OrderRefundAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Defines the order for which the refund will be created. | [optional] 
**Items** | Pointer to [**[]OrderRefundAddItemsInner**](OrderRefundAddItemsInner.md) | Defines items in the order that will be refunded | [optional] 
**TotalPrice** | Pointer to **float32** | Defines order refund amount. | [optional] 
**ShippingPrice** | Pointer to **float32** | Defines refund shipping amount. | [optional] 
**FeePrice** | Pointer to **float32** | Specifies refund&#39;s fee price | [optional] 
**Message** | Pointer to **string** | Refund reason, or some else message which assigned to refund. | [optional] 
**ItemRestock** | Pointer to **bool** | Boolean, whether or not to add the line items back to the store inventory. | [optional] [default to false]
**SendNotifications** | Pointer to **bool** | Send notifications to customer after refund was created | [optional] [default to false]
**Date** | Pointer to **string** | Specifies an order creation date in format Y-m-d H:i:s | [optional] 
**IsOnline** | Pointer to **bool** | Indicates whether refund type is online | [optional] [default to false]

## Methods

### NewOrderRefundAdd

`func NewOrderRefundAdd() *OrderRefundAdd`

NewOrderRefundAdd instantiates a new OrderRefundAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRefundAddWithDefaults

`func NewOrderRefundAddWithDefaults() *OrderRefundAdd`

NewOrderRefundAddWithDefaults instantiates a new OrderRefundAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *OrderRefundAdd) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderRefundAdd) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderRefundAdd) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderRefundAdd) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetItems

`func (o *OrderRefundAdd) GetItems() []OrderRefundAddItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrderRefundAdd) GetItemsOk() (*[]OrderRefundAddItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrderRefundAdd) SetItems(v []OrderRefundAddItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *OrderRefundAdd) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalPrice

`func (o *OrderRefundAdd) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *OrderRefundAdd) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *OrderRefundAdd) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *OrderRefundAdd) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetShippingPrice

`func (o *OrderRefundAdd) GetShippingPrice() float32`

GetShippingPrice returns the ShippingPrice field if non-nil, zero value otherwise.

### GetShippingPriceOk

`func (o *OrderRefundAdd) GetShippingPriceOk() (*float32, bool)`

GetShippingPriceOk returns a tuple with the ShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPrice

`func (o *OrderRefundAdd) SetShippingPrice(v float32)`

SetShippingPrice sets ShippingPrice field to given value.

### HasShippingPrice

`func (o *OrderRefundAdd) HasShippingPrice() bool`

HasShippingPrice returns a boolean if a field has been set.

### GetFeePrice

`func (o *OrderRefundAdd) GetFeePrice() float32`

GetFeePrice returns the FeePrice field if non-nil, zero value otherwise.

### GetFeePriceOk

`func (o *OrderRefundAdd) GetFeePriceOk() (*float32, bool)`

GetFeePriceOk returns a tuple with the FeePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePrice

`func (o *OrderRefundAdd) SetFeePrice(v float32)`

SetFeePrice sets FeePrice field to given value.

### HasFeePrice

`func (o *OrderRefundAdd) HasFeePrice() bool`

HasFeePrice returns a boolean if a field has been set.

### GetMessage

`func (o *OrderRefundAdd) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OrderRefundAdd) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OrderRefundAdd) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OrderRefundAdd) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetItemRestock

`func (o *OrderRefundAdd) GetItemRestock() bool`

GetItemRestock returns the ItemRestock field if non-nil, zero value otherwise.

### GetItemRestockOk

`func (o *OrderRefundAdd) GetItemRestockOk() (*bool, bool)`

GetItemRestockOk returns a tuple with the ItemRestock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRestock

`func (o *OrderRefundAdd) SetItemRestock(v bool)`

SetItemRestock sets ItemRestock field to given value.

### HasItemRestock

`func (o *OrderRefundAdd) HasItemRestock() bool`

HasItemRestock returns a boolean if a field has been set.

### GetSendNotifications

`func (o *OrderRefundAdd) GetSendNotifications() bool`

GetSendNotifications returns the SendNotifications field if non-nil, zero value otherwise.

### GetSendNotificationsOk

`func (o *OrderRefundAdd) GetSendNotificationsOk() (*bool, bool)`

GetSendNotificationsOk returns a tuple with the SendNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotifications

`func (o *OrderRefundAdd) SetSendNotifications(v bool)`

SetSendNotifications sets SendNotifications field to given value.

### HasSendNotifications

`func (o *OrderRefundAdd) HasSendNotifications() bool`

HasSendNotifications returns a boolean if a field has been set.

### GetDate

`func (o *OrderRefundAdd) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *OrderRefundAdd) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *OrderRefundAdd) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *OrderRefundAdd) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetIsOnline

`func (o *OrderRefundAdd) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *OrderRefundAdd) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *OrderRefundAdd) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.

### HasIsOnline

`func (o *OrderRefundAdd) HasIsOnline() bool`

HasIsOnline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


