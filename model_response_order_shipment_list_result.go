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

// checks if the ResponseOrderShipmentListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseOrderShipmentListResult{}

// ResponseOrderShipmentListResult struct for ResponseOrderShipmentListResult
type ResponseOrderShipmentListResult struct {
	ShipmentCount *int32 `json:"shipment_count,omitempty"`
	Shipment []Shipment `json:"shipment,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewResponseOrderShipmentListResult instantiates a new ResponseOrderShipmentListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseOrderShipmentListResult() *ResponseOrderShipmentListResult {
	this := ResponseOrderShipmentListResult{}
	return &this
}

// NewResponseOrderShipmentListResultWithDefaults instantiates a new ResponseOrderShipmentListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseOrderShipmentListResultWithDefaults() *ResponseOrderShipmentListResult {
	this := ResponseOrderShipmentListResult{}
	return &this
}

// GetShipmentCount returns the ShipmentCount field value if set, zero value otherwise.
func (o *ResponseOrderShipmentListResult) GetShipmentCount() int32 {
	if o == nil || IsNil(o.ShipmentCount) {
		var ret int32
		return ret
	}
	return *o.ShipmentCount
}

// GetShipmentCountOk returns a tuple with the ShipmentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseOrderShipmentListResult) GetShipmentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ShipmentCount) {
		return nil, false
	}
	return o.ShipmentCount, true
}

// HasShipmentCount returns a boolean if a field has been set.
func (o *ResponseOrderShipmentListResult) HasShipmentCount() bool {
	if o != nil && !IsNil(o.ShipmentCount) {
		return true
	}

	return false
}

// SetShipmentCount gets a reference to the given int32 and assigns it to the ShipmentCount field.
func (o *ResponseOrderShipmentListResult) SetShipmentCount(v int32) {
	o.ShipmentCount = &v
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *ResponseOrderShipmentListResult) GetShipment() []Shipment {
	if o == nil || IsNil(o.Shipment) {
		var ret []Shipment
		return ret
	}
	return o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseOrderShipmentListResult) GetShipmentOk() ([]Shipment, bool) {
	if o == nil || IsNil(o.Shipment) {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *ResponseOrderShipmentListResult) HasShipment() bool {
	if o != nil && !IsNil(o.Shipment) {
		return true
	}

	return false
}

// SetShipment gets a reference to the given []Shipment and assigns it to the Shipment field.
func (o *ResponseOrderShipmentListResult) SetShipment(v []Shipment) {
	o.Shipment = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *ResponseOrderShipmentListResult) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseOrderShipmentListResult) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ResponseOrderShipmentListResult) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ResponseOrderShipmentListResult) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ResponseOrderShipmentListResult) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseOrderShipmentListResult) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ResponseOrderShipmentListResult) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ResponseOrderShipmentListResult) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ResponseOrderShipmentListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseOrderShipmentListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShipmentCount) {
		toSerialize["shipment_count"] = o.ShipmentCount
	}
	if !IsNil(o.Shipment) {
		toSerialize["shipment"] = o.Shipment
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableResponseOrderShipmentListResult struct {
	value *ResponseOrderShipmentListResult
	isSet bool
}

func (v NullableResponseOrderShipmentListResult) Get() *ResponseOrderShipmentListResult {
	return v.value
}

func (v *NullableResponseOrderShipmentListResult) Set(val *ResponseOrderShipmentListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseOrderShipmentListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseOrderShipmentListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseOrderShipmentListResult(val *ResponseOrderShipmentListResult) *NullableResponseOrderShipmentListResult {
	return &NullableResponseOrderShipmentListResult{value: val, isSet: true}
}

func (v NullableResponseOrderShipmentListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseOrderShipmentListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


