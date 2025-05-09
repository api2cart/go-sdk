/*
API2Cart OpenAPI

API2Cart

API version: 1.1
Contact: contact@api2cart.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OrderTotals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderTotals{}

// OrderTotals struct for OrderTotals
type OrderTotals struct {
	Total *float32 `json:"total,omitempty"`
	Subtotal *float32 `json:"subtotal,omitempty"`
	Shipping *float32 `json:"shipping,omitempty"`
	Tax *float32 `json:"tax,omitempty"`
	Discount *float32 `json:"discount,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewOrderTotals instantiates a new OrderTotals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTotals() *OrderTotals {
	this := OrderTotals{}
	return &this
}

// NewOrderTotalsWithDefaults instantiates a new OrderTotals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTotalsWithDefaults() *OrderTotals {
	this := OrderTotals{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrderTotals) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotals) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrderTotals) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *OrderTotals) SetTotal(v float32) {
	o.Total = &v
}

// GetSubtotal returns the Subtotal field value if set, zero value otherwise.
func (o *OrderTotals) GetSubtotal() float32 {
	if o == nil || IsNil(o.Subtotal) {
		var ret float32
		return ret
	}
	return *o.Subtotal
}

// GetSubtotalOk returns a tuple with the Subtotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotals) GetSubtotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Subtotal) {
		return nil, false
	}
	return o.Subtotal, true
}

// HasSubtotal returns a boolean if a field has been set.
func (o *OrderTotals) HasSubtotal() bool {
	if o != nil && !IsNil(o.Subtotal) {
		return true
	}

	return false
}

// SetSubtotal gets a reference to the given float32 and assigns it to the Subtotal field.
func (o *OrderTotals) SetSubtotal(v float32) {
	o.Subtotal = &v
}

// GetShipping returns the Shipping field value if set, zero value otherwise.
func (o *OrderTotals) GetShipping() float32 {
	if o == nil || IsNil(o.Shipping) {
		var ret float32
		return ret
	}
	return *o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotals) GetShippingOk() (*float32, bool) {
	if o == nil || IsNil(o.Shipping) {
		return nil, false
	}
	return o.Shipping, true
}

// HasShipping returns a boolean if a field has been set.
func (o *OrderTotals) HasShipping() bool {
	if o != nil && !IsNil(o.Shipping) {
		return true
	}

	return false
}

// SetShipping gets a reference to the given float32 and assigns it to the Shipping field.
func (o *OrderTotals) SetShipping(v float32) {
	o.Shipping = &v
}

// GetTax returns the Tax field value if set, zero value otherwise.
func (o *OrderTotals) GetTax() float32 {
	if o == nil || IsNil(o.Tax) {
		var ret float32
		return ret
	}
	return *o.Tax
}

// GetTaxOk returns a tuple with the Tax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotals) GetTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Tax) {
		return nil, false
	}
	return o.Tax, true
}

// HasTax returns a boolean if a field has been set.
func (o *OrderTotals) HasTax() bool {
	if o != nil && !IsNil(o.Tax) {
		return true
	}

	return false
}

// SetTax gets a reference to the given float32 and assigns it to the Tax field.
func (o *OrderTotals) SetTax(v float32) {
	o.Tax = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *OrderTotals) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotals) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *OrderTotals) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *OrderTotals) SetDiscount(v float32) {
	o.Discount = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *OrderTotals) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotals) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *OrderTotals) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *OrderTotals) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *OrderTotals) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotals) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *OrderTotals) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *OrderTotals) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o OrderTotals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTotals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Subtotal) {
		toSerialize["subtotal"] = o.Subtotal
	}
	if !IsNil(o.Shipping) {
		toSerialize["shipping"] = o.Shipping
	}
	if !IsNil(o.Tax) {
		toSerialize["tax"] = o.Tax
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableOrderTotals struct {
	value *OrderTotals
	isSet bool
}

func (v NullableOrderTotals) Get() *OrderTotals {
	return v.value
}

func (v *NullableOrderTotals) Set(val *OrderTotals) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTotals) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTotals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTotals(val *OrderTotals) *NullableOrderTotals {
	return &NullableOrderTotals{value: val, isSet: true}
}

func (v NullableOrderTotals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTotals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


