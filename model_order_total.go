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

// checks if the OrderTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderTotal{}

// OrderTotal struct for OrderTotal
type OrderTotal struct {
	SubtotalExTax *float32 `json:"subtotal_ex_tax,omitempty"`
	WrappingExTax NullableFloat32 `json:"wrapping_ex_tax,omitempty"`
	ShippingExTax *float32 `json:"shipping_ex_tax,omitempty"`
	TotalDiscount *float32 `json:"total_discount,omitempty"`
	TotalTax *float32 `json:"total_tax,omitempty"`
	Total *float32 `json:"total,omitempty"`
	TotalPaid NullableFloat32 `json:"total_paid,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewOrderTotal instantiates a new OrderTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTotal() *OrderTotal {
	this := OrderTotal{}
	return &this
}

// NewOrderTotalWithDefaults instantiates a new OrderTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTotalWithDefaults() *OrderTotal {
	this := OrderTotal{}
	return &this
}

// GetSubtotalExTax returns the SubtotalExTax field value if set, zero value otherwise.
func (o *OrderTotal) GetSubtotalExTax() float32 {
	if o == nil || IsNil(o.SubtotalExTax) {
		var ret float32
		return ret
	}
	return *o.SubtotalExTax
}

// GetSubtotalExTaxOk returns a tuple with the SubtotalExTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotal) GetSubtotalExTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.SubtotalExTax) {
		return nil, false
	}
	return o.SubtotalExTax, true
}

// HasSubtotalExTax returns a boolean if a field has been set.
func (o *OrderTotal) HasSubtotalExTax() bool {
	if o != nil && !IsNil(o.SubtotalExTax) {
		return true
	}

	return false
}

// SetSubtotalExTax gets a reference to the given float32 and assigns it to the SubtotalExTax field.
func (o *OrderTotal) SetSubtotalExTax(v float32) {
	o.SubtotalExTax = &v
}

// GetWrappingExTax returns the WrappingExTax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTotal) GetWrappingExTax() float32 {
	if o == nil || IsNil(o.WrappingExTax.Get()) {
		var ret float32
		return ret
	}
	return *o.WrappingExTax.Get()
}

// GetWrappingExTaxOk returns a tuple with the WrappingExTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTotal) GetWrappingExTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WrappingExTax.Get(), o.WrappingExTax.IsSet()
}

// HasWrappingExTax returns a boolean if a field has been set.
func (o *OrderTotal) HasWrappingExTax() bool {
	if o != nil && o.WrappingExTax.IsSet() {
		return true
	}

	return false
}

// SetWrappingExTax gets a reference to the given NullableFloat32 and assigns it to the WrappingExTax field.
func (o *OrderTotal) SetWrappingExTax(v float32) {
	o.WrappingExTax.Set(&v)
}
// SetWrappingExTaxNil sets the value for WrappingExTax to be an explicit nil
func (o *OrderTotal) SetWrappingExTaxNil() {
	o.WrappingExTax.Set(nil)
}

// UnsetWrappingExTax ensures that no value is present for WrappingExTax, not even an explicit nil
func (o *OrderTotal) UnsetWrappingExTax() {
	o.WrappingExTax.Unset()
}

// GetShippingExTax returns the ShippingExTax field value if set, zero value otherwise.
func (o *OrderTotal) GetShippingExTax() float32 {
	if o == nil || IsNil(o.ShippingExTax) {
		var ret float32
		return ret
	}
	return *o.ShippingExTax
}

// GetShippingExTaxOk returns a tuple with the ShippingExTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotal) GetShippingExTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.ShippingExTax) {
		return nil, false
	}
	return o.ShippingExTax, true
}

// HasShippingExTax returns a boolean if a field has been set.
func (o *OrderTotal) HasShippingExTax() bool {
	if o != nil && !IsNil(o.ShippingExTax) {
		return true
	}

	return false
}

// SetShippingExTax gets a reference to the given float32 and assigns it to the ShippingExTax field.
func (o *OrderTotal) SetShippingExTax(v float32) {
	o.ShippingExTax = &v
}

// GetTotalDiscount returns the TotalDiscount field value if set, zero value otherwise.
func (o *OrderTotal) GetTotalDiscount() float32 {
	if o == nil || IsNil(o.TotalDiscount) {
		var ret float32
		return ret
	}
	return *o.TotalDiscount
}

// GetTotalDiscountOk returns a tuple with the TotalDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotal) GetTotalDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalDiscount) {
		return nil, false
	}
	return o.TotalDiscount, true
}

// HasTotalDiscount returns a boolean if a field has been set.
func (o *OrderTotal) HasTotalDiscount() bool {
	if o != nil && !IsNil(o.TotalDiscount) {
		return true
	}

	return false
}

// SetTotalDiscount gets a reference to the given float32 and assigns it to the TotalDiscount field.
func (o *OrderTotal) SetTotalDiscount(v float32) {
	o.TotalDiscount = &v
}

// GetTotalTax returns the TotalTax field value if set, zero value otherwise.
func (o *OrderTotal) GetTotalTax() float32 {
	if o == nil || IsNil(o.TotalTax) {
		var ret float32
		return ret
	}
	return *o.TotalTax
}

// GetTotalTaxOk returns a tuple with the TotalTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotal) GetTotalTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalTax) {
		return nil, false
	}
	return o.TotalTax, true
}

// HasTotalTax returns a boolean if a field has been set.
func (o *OrderTotal) HasTotalTax() bool {
	if o != nil && !IsNil(o.TotalTax) {
		return true
	}

	return false
}

// SetTotalTax gets a reference to the given float32 and assigns it to the TotalTax field.
func (o *OrderTotal) SetTotalTax(v float32) {
	o.TotalTax = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrderTotal) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTotal) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrderTotal) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *OrderTotal) SetTotal(v float32) {
	o.Total = &v
}

// GetTotalPaid returns the TotalPaid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTotal) GetTotalPaid() float32 {
	if o == nil || IsNil(o.TotalPaid.Get()) {
		var ret float32
		return ret
	}
	return *o.TotalPaid.Get()
}

// GetTotalPaidOk returns a tuple with the TotalPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTotal) GetTotalPaidOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalPaid.Get(), o.TotalPaid.IsSet()
}

// HasTotalPaid returns a boolean if a field has been set.
func (o *OrderTotal) HasTotalPaid() bool {
	if o != nil && o.TotalPaid.IsSet() {
		return true
	}

	return false
}

// SetTotalPaid gets a reference to the given NullableFloat32 and assigns it to the TotalPaid field.
func (o *OrderTotal) SetTotalPaid(v float32) {
	o.TotalPaid.Set(&v)
}
// SetTotalPaidNil sets the value for TotalPaid to be an explicit nil
func (o *OrderTotal) SetTotalPaidNil() {
	o.TotalPaid.Set(nil)
}

// UnsetTotalPaid ensures that no value is present for TotalPaid, not even an explicit nil
func (o *OrderTotal) UnsetTotalPaid() {
	o.TotalPaid.Unset()
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTotal) GetAdditionalFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTotal) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *OrderTotal) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *OrderTotal) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTotal) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTotal) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *OrderTotal) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *OrderTotal) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o OrderTotal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubtotalExTax) {
		toSerialize["subtotal_ex_tax"] = o.SubtotalExTax
	}
	if o.WrappingExTax.IsSet() {
		toSerialize["wrapping_ex_tax"] = o.WrappingExTax.Get()
	}
	if !IsNil(o.ShippingExTax) {
		toSerialize["shipping_ex_tax"] = o.ShippingExTax
	}
	if !IsNil(o.TotalDiscount) {
		toSerialize["total_discount"] = o.TotalDiscount
	}
	if !IsNil(o.TotalTax) {
		toSerialize["total_tax"] = o.TotalTax
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if o.TotalPaid.IsSet() {
		toSerialize["total_paid"] = o.TotalPaid.Get()
	}
	if o.AdditionalFields != nil {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableOrderTotal struct {
	value *OrderTotal
	isSet bool
}

func (v NullableOrderTotal) Get() *OrderTotal {
	return v.value
}

func (v *NullableOrderTotal) Set(val *OrderTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTotal(val *OrderTotal) *NullableOrderTotal {
	return &NullableOrderTotal{value: val, isSet: true}
}

func (v NullableOrderTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


