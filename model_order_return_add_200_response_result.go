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

// checks if the OrderReturnAdd200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderReturnAdd200ResponseResult{}

// OrderReturnAdd200ResponseResult struct for OrderReturnAdd200ResponseResult
type OrderReturnAdd200ResponseResult struct {
	ReturnId *string `json:"return_id,omitempty"`
}

// NewOrderReturnAdd200ResponseResult instantiates a new OrderReturnAdd200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderReturnAdd200ResponseResult() *OrderReturnAdd200ResponseResult {
	this := OrderReturnAdd200ResponseResult{}
	return &this
}

// NewOrderReturnAdd200ResponseResultWithDefaults instantiates a new OrderReturnAdd200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderReturnAdd200ResponseResultWithDefaults() *OrderReturnAdd200ResponseResult {
	this := OrderReturnAdd200ResponseResult{}
	return &this
}

// GetReturnId returns the ReturnId field value if set, zero value otherwise.
func (o *OrderReturnAdd200ResponseResult) GetReturnId() string {
	if o == nil || IsNil(o.ReturnId) {
		var ret string
		return ret
	}
	return *o.ReturnId
}

// GetReturnIdOk returns a tuple with the ReturnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderReturnAdd200ResponseResult) GetReturnIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnId) {
		return nil, false
	}
	return o.ReturnId, true
}

// HasReturnId returns a boolean if a field has been set.
func (o *OrderReturnAdd200ResponseResult) HasReturnId() bool {
	if o != nil && !IsNil(o.ReturnId) {
		return true
	}

	return false
}

// SetReturnId gets a reference to the given string and assigns it to the ReturnId field.
func (o *OrderReturnAdd200ResponseResult) SetReturnId(v string) {
	o.ReturnId = &v
}

func (o OrderReturnAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderReturnAdd200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnId) {
		toSerialize["return_id"] = o.ReturnId
	}
	return toSerialize, nil
}

type NullableOrderReturnAdd200ResponseResult struct {
	value *OrderReturnAdd200ResponseResult
	isSet bool
}

func (v NullableOrderReturnAdd200ResponseResult) Get() *OrderReturnAdd200ResponseResult {
	return v.value
}

func (v *NullableOrderReturnAdd200ResponseResult) Set(val *OrderReturnAdd200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderReturnAdd200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderReturnAdd200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderReturnAdd200ResponseResult(val *OrderReturnAdd200ResponseResult) *NullableOrderReturnAdd200ResponseResult {
	return &NullableOrderReturnAdd200ResponseResult{value: val, isSet: true}
}

func (v NullableOrderReturnAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderReturnAdd200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


