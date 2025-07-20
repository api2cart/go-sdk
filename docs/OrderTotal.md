# OrderTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubtotalExTax** | Pointer to **float32** |  | [optional] 
**WrappingExTax** | Pointer to **NullableFloat32** |  | [optional] 
**ShippingExTax** | Pointer to **float32** |  | [optional] 
**TotalDiscount** | Pointer to **float32** |  | [optional] 
**TotalTax** | Pointer to **float32** |  | [optional] 
**Total** | Pointer to **float32** |  | [optional] 
**TotalPaid** | Pointer to **NullableFloat32** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderTotal

`func NewOrderTotal() *OrderTotal`

NewOrderTotal instantiates a new OrderTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTotalWithDefaults

`func NewOrderTotalWithDefaults() *OrderTotal`

NewOrderTotalWithDefaults instantiates a new OrderTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtotalExTax

`func (o *OrderTotal) GetSubtotalExTax() float32`

GetSubtotalExTax returns the SubtotalExTax field if non-nil, zero value otherwise.

### GetSubtotalExTaxOk

`func (o *OrderTotal) GetSubtotalExTaxOk() (*float32, bool)`

GetSubtotalExTaxOk returns a tuple with the SubtotalExTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalExTax

`func (o *OrderTotal) SetSubtotalExTax(v float32)`

SetSubtotalExTax sets SubtotalExTax field to given value.

### HasSubtotalExTax

`func (o *OrderTotal) HasSubtotalExTax() bool`

HasSubtotalExTax returns a boolean if a field has been set.

### GetWrappingExTax

`func (o *OrderTotal) GetWrappingExTax() float32`

GetWrappingExTax returns the WrappingExTax field if non-nil, zero value otherwise.

### GetWrappingExTaxOk

`func (o *OrderTotal) GetWrappingExTaxOk() (*float32, bool)`

GetWrappingExTaxOk returns a tuple with the WrappingExTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrappingExTax

`func (o *OrderTotal) SetWrappingExTax(v float32)`

SetWrappingExTax sets WrappingExTax field to given value.

### HasWrappingExTax

`func (o *OrderTotal) HasWrappingExTax() bool`

HasWrappingExTax returns a boolean if a field has been set.

### SetWrappingExTaxNil

`func (o *OrderTotal) SetWrappingExTaxNil(b bool)`

 SetWrappingExTaxNil sets the value for WrappingExTax to be an explicit nil

### UnsetWrappingExTax
`func (o *OrderTotal) UnsetWrappingExTax()`

UnsetWrappingExTax ensures that no value is present for WrappingExTax, not even an explicit nil
### GetShippingExTax

`func (o *OrderTotal) GetShippingExTax() float32`

GetShippingExTax returns the ShippingExTax field if non-nil, zero value otherwise.

### GetShippingExTaxOk

`func (o *OrderTotal) GetShippingExTaxOk() (*float32, bool)`

GetShippingExTaxOk returns a tuple with the ShippingExTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingExTax

`func (o *OrderTotal) SetShippingExTax(v float32)`

SetShippingExTax sets ShippingExTax field to given value.

### HasShippingExTax

`func (o *OrderTotal) HasShippingExTax() bool`

HasShippingExTax returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *OrderTotal) GetTotalDiscount() float32`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *OrderTotal) GetTotalDiscountOk() (*float32, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *OrderTotal) SetTotalDiscount(v float32)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *OrderTotal) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### GetTotalTax

`func (o *OrderTotal) GetTotalTax() float32`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *OrderTotal) GetTotalTaxOk() (*float32, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *OrderTotal) SetTotalTax(v float32)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTax

`func (o *OrderTotal) HasTotalTax() bool`

HasTotalTax returns a boolean if a field has been set.

### GetTotal

`func (o *OrderTotal) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrderTotal) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrderTotal) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OrderTotal) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalPaid

`func (o *OrderTotal) GetTotalPaid() float32`

GetTotalPaid returns the TotalPaid field if non-nil, zero value otherwise.

### GetTotalPaidOk

`func (o *OrderTotal) GetTotalPaidOk() (*float32, bool)`

GetTotalPaidOk returns a tuple with the TotalPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaid

`func (o *OrderTotal) SetTotalPaid(v float32)`

SetTotalPaid sets TotalPaid field to given value.

### HasTotalPaid

`func (o *OrderTotal) HasTotalPaid() bool`

HasTotalPaid returns a boolean if a field has been set.

### SetTotalPaidNil

`func (o *OrderTotal) SetTotalPaidNil(b bool)`

 SetTotalPaidNil sets the value for TotalPaid to be an explicit nil

### UnsetTotalPaid
`func (o *OrderTotal) UnsetTotalPaid()`

UnsetTotalPaid ensures that no value is present for TotalPaid, not even an explicit nil
### GetAdditionalFields

`func (o *OrderTotal) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderTotal) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderTotal) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderTotal) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *OrderTotal) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *OrderTotal) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *OrderTotal) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderTotal) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderTotal) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderTotal) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *OrderTotal) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *OrderTotal) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


