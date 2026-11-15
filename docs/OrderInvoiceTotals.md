# OrderInvoiceTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subtotal** | Pointer to **NullableFloat32** |  | [optional] 
**Discount** | Pointer to **NullableFloat32** |  | [optional] 
**Shipping** | Pointer to **NullableFloat32** |  | [optional] 
**Tax** | Pointer to **NullableFloat32** |  | [optional] 
**Total** | Pointer to **NullableFloat32** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderInvoiceTotals

`func NewOrderInvoiceTotals() *OrderInvoiceTotals`

NewOrderInvoiceTotals instantiates a new OrderInvoiceTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderInvoiceTotalsWithDefaults

`func NewOrderInvoiceTotalsWithDefaults() *OrderInvoiceTotals`

NewOrderInvoiceTotalsWithDefaults instantiates a new OrderInvoiceTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtotal

`func (o *OrderInvoiceTotals) GetSubtotal() float32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *OrderInvoiceTotals) GetSubtotalOk() (*float32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *OrderInvoiceTotals) SetSubtotal(v float32)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *OrderInvoiceTotals) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### SetSubtotalNil

`func (o *OrderInvoiceTotals) SetSubtotalNil(b bool)`

 SetSubtotalNil sets the value for Subtotal to be an explicit nil

### UnsetSubtotal
`func (o *OrderInvoiceTotals) UnsetSubtotal()`

UnsetSubtotal ensures that no value is present for Subtotal, not even an explicit nil
### GetDiscount

`func (o *OrderInvoiceTotals) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *OrderInvoiceTotals) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *OrderInvoiceTotals) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *OrderInvoiceTotals) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *OrderInvoiceTotals) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *OrderInvoiceTotals) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
### GetShipping

`func (o *OrderInvoiceTotals) GetShipping() float32`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *OrderInvoiceTotals) GetShippingOk() (*float32, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *OrderInvoiceTotals) SetShipping(v float32)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *OrderInvoiceTotals) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### SetShippingNil

`func (o *OrderInvoiceTotals) SetShippingNil(b bool)`

 SetShippingNil sets the value for Shipping to be an explicit nil

### UnsetShipping
`func (o *OrderInvoiceTotals) UnsetShipping()`

UnsetShipping ensures that no value is present for Shipping, not even an explicit nil
### GetTax

`func (o *OrderInvoiceTotals) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *OrderInvoiceTotals) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *OrderInvoiceTotals) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *OrderInvoiceTotals) HasTax() bool`

HasTax returns a boolean if a field has been set.

### SetTaxNil

`func (o *OrderInvoiceTotals) SetTaxNil(b bool)`

 SetTaxNil sets the value for Tax to be an explicit nil

### UnsetTax
`func (o *OrderInvoiceTotals) UnsetTax()`

UnsetTax ensures that no value is present for Tax, not even an explicit nil
### GetTotal

`func (o *OrderInvoiceTotals) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrderInvoiceTotals) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrderInvoiceTotals) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OrderInvoiceTotals) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *OrderInvoiceTotals) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *OrderInvoiceTotals) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetAdditionalFields

`func (o *OrderInvoiceTotals) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderInvoiceTotals) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderInvoiceTotals) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderInvoiceTotals) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *OrderInvoiceTotals) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *OrderInvoiceTotals) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *OrderInvoiceTotals) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderInvoiceTotals) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderInvoiceTotals) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderInvoiceTotals) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *OrderInvoiceTotals) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *OrderInvoiceTotals) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


