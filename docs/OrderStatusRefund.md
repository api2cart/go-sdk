# OrderStatusRefund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipping** | Pointer to **float32** |  | [optional] 
**Fee** | Pointer to **float32** |  | [optional] 
**Tax** | Pointer to **float32** |  | [optional] 
**TotalRefunded** | Pointer to **float32** |  | [optional] 
**Time** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**RefundedItems** | Pointer to [**[]OrderStatusRefundItem**](OrderStatusRefundItem.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderStatusRefund

`func NewOrderStatusRefund() *OrderStatusRefund`

NewOrderStatusRefund instantiates a new OrderStatusRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusRefundWithDefaults

`func NewOrderStatusRefundWithDefaults() *OrderStatusRefund`

NewOrderStatusRefundWithDefaults instantiates a new OrderStatusRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipping

`func (o *OrderStatusRefund) GetShipping() float32`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *OrderStatusRefund) GetShippingOk() (*float32, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *OrderStatusRefund) SetShipping(v float32)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *OrderStatusRefund) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetFee

`func (o *OrderStatusRefund) GetFee() float32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *OrderStatusRefund) GetFeeOk() (*float32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *OrderStatusRefund) SetFee(v float32)`

SetFee sets Fee field to given value.

### HasFee

`func (o *OrderStatusRefund) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetTax

`func (o *OrderStatusRefund) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *OrderStatusRefund) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *OrderStatusRefund) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *OrderStatusRefund) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTotalRefunded

`func (o *OrderStatusRefund) GetTotalRefunded() float32`

GetTotalRefunded returns the TotalRefunded field if non-nil, zero value otherwise.

### GetTotalRefundedOk

`func (o *OrderStatusRefund) GetTotalRefundedOk() (*float32, bool)`

GetTotalRefundedOk returns a tuple with the TotalRefunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRefunded

`func (o *OrderStatusRefund) SetTotalRefunded(v float32)`

SetTotalRefunded sets TotalRefunded field to given value.

### HasTotalRefunded

`func (o *OrderStatusRefund) HasTotalRefunded() bool`

HasTotalRefunded returns a boolean if a field has been set.

### GetTime

`func (o *OrderStatusRefund) GetTime() A2CDateTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OrderStatusRefund) GetTimeOk() (*A2CDateTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OrderStatusRefund) SetTime(v A2CDateTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *OrderStatusRefund) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetComment

`func (o *OrderStatusRefund) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *OrderStatusRefund) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *OrderStatusRefund) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *OrderStatusRefund) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRefundedItems

`func (o *OrderStatusRefund) GetRefundedItems() []OrderStatusRefundItem`

GetRefundedItems returns the RefundedItems field if non-nil, zero value otherwise.

### GetRefundedItemsOk

`func (o *OrderStatusRefund) GetRefundedItemsOk() (*[]OrderStatusRefundItem, bool)`

GetRefundedItemsOk returns a tuple with the RefundedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedItems

`func (o *OrderStatusRefund) SetRefundedItems(v []OrderStatusRefundItem)`

SetRefundedItems sets RefundedItems field to given value.

### HasRefundedItems

`func (o *OrderStatusRefund) HasRefundedItems() bool`

HasRefundedItems returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *OrderStatusRefund) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderStatusRefund) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderStatusRefund) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderStatusRefund) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *OrderStatusRefund) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderStatusRefund) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderStatusRefund) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderStatusRefund) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


