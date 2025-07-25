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

// checks if the TaxClassZipCodesRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxClassZipCodesRange{}

// TaxClassZipCodesRange struct for TaxClassZipCodesRange
type TaxClassZipCodesRange struct {
	From NullableString `json:"from,omitempty"`
	To NullableString `json:"to,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewTaxClassZipCodesRange instantiates a new TaxClassZipCodesRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxClassZipCodesRange() *TaxClassZipCodesRange {
	this := TaxClassZipCodesRange{}
	return &this
}

// NewTaxClassZipCodesRangeWithDefaults instantiates a new TaxClassZipCodesRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxClassZipCodesRangeWithDefaults() *TaxClassZipCodesRange {
	this := TaxClassZipCodesRange{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxClassZipCodesRange) GetFrom() string {
	if o == nil || IsNil(o.From.Get()) {
		var ret string
		return ret
	}
	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxClassZipCodesRange) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// HasFrom returns a boolean if a field has been set.
func (o *TaxClassZipCodesRange) HasFrom() bool {
	if o != nil && o.From.IsSet() {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NullableString and assigns it to the From field.
func (o *TaxClassZipCodesRange) SetFrom(v string) {
	o.From.Set(&v)
}
// SetFromNil sets the value for From to be an explicit nil
func (o *TaxClassZipCodesRange) SetFromNil() {
	o.From.Set(nil)
}

// UnsetFrom ensures that no value is present for From, not even an explicit nil
func (o *TaxClassZipCodesRange) UnsetFrom() {
	o.From.Unset()
}

// GetTo returns the To field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxClassZipCodesRange) GetTo() string {
	if o == nil || IsNil(o.To.Get()) {
		var ret string
		return ret
	}
	return *o.To.Get()
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxClassZipCodesRange) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.To.Get(), o.To.IsSet()
}

// HasTo returns a boolean if a field has been set.
func (o *TaxClassZipCodesRange) HasTo() bool {
	if o != nil && o.To.IsSet() {
		return true
	}

	return false
}

// SetTo gets a reference to the given NullableString and assigns it to the To field.
func (o *TaxClassZipCodesRange) SetTo(v string) {
	o.To.Set(&v)
}
// SetToNil sets the value for To to be an explicit nil
func (o *TaxClassZipCodesRange) SetToNil() {
	o.To.Set(nil)
}

// UnsetTo ensures that no value is present for To, not even an explicit nil
func (o *TaxClassZipCodesRange) UnsetTo() {
	o.To.Unset()
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxClassZipCodesRange) GetAdditionalFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxClassZipCodesRange) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *TaxClassZipCodesRange) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *TaxClassZipCodesRange) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxClassZipCodesRange) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxClassZipCodesRange) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *TaxClassZipCodesRange) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *TaxClassZipCodesRange) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o TaxClassZipCodesRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxClassZipCodesRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.From.IsSet() {
		toSerialize["from"] = o.From.Get()
	}
	if o.To.IsSet() {
		toSerialize["to"] = o.To.Get()
	}
	if o.AdditionalFields != nil {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableTaxClassZipCodesRange struct {
	value *TaxClassZipCodesRange
	isSet bool
}

func (v NullableTaxClassZipCodesRange) Get() *TaxClassZipCodesRange {
	return v.value
}

func (v *NullableTaxClassZipCodesRange) Set(val *TaxClassZipCodesRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxClassZipCodesRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxClassZipCodesRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxClassZipCodesRange(val *TaxClassZipCodesRange) *NullableTaxClassZipCodesRange {
	return &NullableTaxClassZipCodesRange{value: val, isSet: true}
}

func (v NullableTaxClassZipCodesRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxClassZipCodesRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


