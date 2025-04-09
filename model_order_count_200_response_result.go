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

// checks if the OrderCount200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCount200ResponseResult{}

// OrderCount200ResponseResult struct for OrderCount200ResponseResult
type OrderCount200ResponseResult struct {
	OrdersCount *int32 `json:"orders_count,omitempty"`
}

// NewOrderCount200ResponseResult instantiates a new OrderCount200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCount200ResponseResult() *OrderCount200ResponseResult {
	this := OrderCount200ResponseResult{}
	return &this
}

// NewOrderCount200ResponseResultWithDefaults instantiates a new OrderCount200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCount200ResponseResultWithDefaults() *OrderCount200ResponseResult {
	this := OrderCount200ResponseResult{}
	return &this
}

// GetOrdersCount returns the OrdersCount field value if set, zero value otherwise.
func (o *OrderCount200ResponseResult) GetOrdersCount() int32 {
	if o == nil || IsNil(o.OrdersCount) {
		var ret int32
		return ret
	}
	return *o.OrdersCount
}

// GetOrdersCountOk returns a tuple with the OrdersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCount200ResponseResult) GetOrdersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OrdersCount) {
		return nil, false
	}
	return o.OrdersCount, true
}

// HasOrdersCount returns a boolean if a field has been set.
func (o *OrderCount200ResponseResult) HasOrdersCount() bool {
	if o != nil && !IsNil(o.OrdersCount) {
		return true
	}

	return false
}

// SetOrdersCount gets a reference to the given int32 and assigns it to the OrdersCount field.
func (o *OrderCount200ResponseResult) SetOrdersCount(v int32) {
	o.OrdersCount = &v
}

func (o OrderCount200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCount200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrdersCount) {
		toSerialize["orders_count"] = o.OrdersCount
	}
	return toSerialize, nil
}

type NullableOrderCount200ResponseResult struct {
	value *OrderCount200ResponseResult
	isSet bool
}

func (v NullableOrderCount200ResponseResult) Get() *OrderCount200ResponseResult {
	return v.value
}

func (v *NullableOrderCount200ResponseResult) Set(val *OrderCount200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCount200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCount200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCount200ResponseResult(val *OrderCount200ResponseResult) *NullableOrderCount200ResponseResult {
	return &NullableOrderCount200ResponseResult{value: val, isSet: true}
}

func (v NullableOrderCount200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCount200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


