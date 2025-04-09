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

// checks if the ResponseOrderPreestimateShippingListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseOrderPreestimateShippingListResult{}

// ResponseOrderPreestimateShippingListResult struct for ResponseOrderPreestimateShippingListResult
type ResponseOrderPreestimateShippingListResult struct {
	PreestimateShippingsCount *int32 `json:"preestimate_shippings_count,omitempty"`
	PreestimateShippings []OrderPreestimateShipping `json:"preestimate_shippings,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewResponseOrderPreestimateShippingListResult instantiates a new ResponseOrderPreestimateShippingListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseOrderPreestimateShippingListResult() *ResponseOrderPreestimateShippingListResult {
	this := ResponseOrderPreestimateShippingListResult{}
	return &this
}

// NewResponseOrderPreestimateShippingListResultWithDefaults instantiates a new ResponseOrderPreestimateShippingListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseOrderPreestimateShippingListResultWithDefaults() *ResponseOrderPreestimateShippingListResult {
	this := ResponseOrderPreestimateShippingListResult{}
	return &this
}

// GetPreestimateShippingsCount returns the PreestimateShippingsCount field value if set, zero value otherwise.
func (o *ResponseOrderPreestimateShippingListResult) GetPreestimateShippingsCount() int32 {
	if o == nil || IsNil(o.PreestimateShippingsCount) {
		var ret int32
		return ret
	}
	return *o.PreestimateShippingsCount
}

// GetPreestimateShippingsCountOk returns a tuple with the PreestimateShippingsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseOrderPreestimateShippingListResult) GetPreestimateShippingsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PreestimateShippingsCount) {
		return nil, false
	}
	return o.PreestimateShippingsCount, true
}

// HasPreestimateShippingsCount returns a boolean if a field has been set.
func (o *ResponseOrderPreestimateShippingListResult) HasPreestimateShippingsCount() bool {
	if o != nil && !IsNil(o.PreestimateShippingsCount) {
		return true
	}

	return false
}

// SetPreestimateShippingsCount gets a reference to the given int32 and assigns it to the PreestimateShippingsCount field.
func (o *ResponseOrderPreestimateShippingListResult) SetPreestimateShippingsCount(v int32) {
	o.PreestimateShippingsCount = &v
}

// GetPreestimateShippings returns the PreestimateShippings field value if set, zero value otherwise.
func (o *ResponseOrderPreestimateShippingListResult) GetPreestimateShippings() []OrderPreestimateShipping {
	if o == nil || IsNil(o.PreestimateShippings) {
		var ret []OrderPreestimateShipping
		return ret
	}
	return o.PreestimateShippings
}

// GetPreestimateShippingsOk returns a tuple with the PreestimateShippings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseOrderPreestimateShippingListResult) GetPreestimateShippingsOk() ([]OrderPreestimateShipping, bool) {
	if o == nil || IsNil(o.PreestimateShippings) {
		return nil, false
	}
	return o.PreestimateShippings, true
}

// HasPreestimateShippings returns a boolean if a field has been set.
func (o *ResponseOrderPreestimateShippingListResult) HasPreestimateShippings() bool {
	if o != nil && !IsNil(o.PreestimateShippings) {
		return true
	}

	return false
}

// SetPreestimateShippings gets a reference to the given []OrderPreestimateShipping and assigns it to the PreestimateShippings field.
func (o *ResponseOrderPreestimateShippingListResult) SetPreestimateShippings(v []OrderPreestimateShipping) {
	o.PreestimateShippings = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *ResponseOrderPreestimateShippingListResult) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseOrderPreestimateShippingListResult) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ResponseOrderPreestimateShippingListResult) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ResponseOrderPreestimateShippingListResult) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ResponseOrderPreestimateShippingListResult) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseOrderPreestimateShippingListResult) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ResponseOrderPreestimateShippingListResult) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ResponseOrderPreestimateShippingListResult) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ResponseOrderPreestimateShippingListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseOrderPreestimateShippingListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreestimateShippingsCount) {
		toSerialize["preestimate_shippings_count"] = o.PreestimateShippingsCount
	}
	if !IsNil(o.PreestimateShippings) {
		toSerialize["preestimate_shippings"] = o.PreestimateShippings
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableResponseOrderPreestimateShippingListResult struct {
	value *ResponseOrderPreestimateShippingListResult
	isSet bool
}

func (v NullableResponseOrderPreestimateShippingListResult) Get() *ResponseOrderPreestimateShippingListResult {
	return v.value
}

func (v *NullableResponseOrderPreestimateShippingListResult) Set(val *ResponseOrderPreestimateShippingListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseOrderPreestimateShippingListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseOrderPreestimateShippingListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseOrderPreestimateShippingListResult(val *ResponseOrderPreestimateShippingListResult) *NullableResponseOrderPreestimateShippingListResult {
	return &NullableResponseOrderPreestimateShippingListResult{value: val, isSet: true}
}

func (v NullableResponseOrderPreestimateShippingListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseOrderPreestimateShippingListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


