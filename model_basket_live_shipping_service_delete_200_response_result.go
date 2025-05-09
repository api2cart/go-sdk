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

// checks if the BasketLiveShippingServiceDelete200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasketLiveShippingServiceDelete200ResponseResult{}

// BasketLiveShippingServiceDelete200ResponseResult struct for BasketLiveShippingServiceDelete200ResponseResult
type BasketLiveShippingServiceDelete200ResponseResult struct {
	Status *bool `json:"status,omitempty"`
}

// NewBasketLiveShippingServiceDelete200ResponseResult instantiates a new BasketLiveShippingServiceDelete200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasketLiveShippingServiceDelete200ResponseResult() *BasketLiveShippingServiceDelete200ResponseResult {
	this := BasketLiveShippingServiceDelete200ResponseResult{}
	return &this
}

// NewBasketLiveShippingServiceDelete200ResponseResultWithDefaults instantiates a new BasketLiveShippingServiceDelete200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasketLiveShippingServiceDelete200ResponseResultWithDefaults() *BasketLiveShippingServiceDelete200ResponseResult {
	this := BasketLiveShippingServiceDelete200ResponseResult{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BasketLiveShippingServiceDelete200ResponseResult) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasketLiveShippingServiceDelete200ResponseResult) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BasketLiveShippingServiceDelete200ResponseResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *BasketLiveShippingServiceDelete200ResponseResult) SetStatus(v bool) {
	o.Status = &v
}

func (o BasketLiveShippingServiceDelete200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasketLiveShippingServiceDelete200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableBasketLiveShippingServiceDelete200ResponseResult struct {
	value *BasketLiveShippingServiceDelete200ResponseResult
	isSet bool
}

func (v NullableBasketLiveShippingServiceDelete200ResponseResult) Get() *BasketLiveShippingServiceDelete200ResponseResult {
	return v.value
}

func (v *NullableBasketLiveShippingServiceDelete200ResponseResult) Set(val *BasketLiveShippingServiceDelete200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBasketLiveShippingServiceDelete200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBasketLiveShippingServiceDelete200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasketLiveShippingServiceDelete200ResponseResult(val *BasketLiveShippingServiceDelete200ResponseResult) *NullableBasketLiveShippingServiceDelete200ResponseResult {
	return &NullableBasketLiveShippingServiceDelete200ResponseResult{value: val, isSet: true}
}

func (v NullableBasketLiveShippingServiceDelete200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasketLiveShippingServiceDelete200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


