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

// checks if the CustomerCount200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerCount200ResponseResult{}

// CustomerCount200ResponseResult struct for CustomerCount200ResponseResult
type CustomerCount200ResponseResult struct {
	CustomersCount *int32 `json:"customers_count,omitempty"`
}

// NewCustomerCount200ResponseResult instantiates a new CustomerCount200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCount200ResponseResult() *CustomerCount200ResponseResult {
	this := CustomerCount200ResponseResult{}
	return &this
}

// NewCustomerCount200ResponseResultWithDefaults instantiates a new CustomerCount200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCount200ResponseResultWithDefaults() *CustomerCount200ResponseResult {
	this := CustomerCount200ResponseResult{}
	return &this
}

// GetCustomersCount returns the CustomersCount field value if set, zero value otherwise.
func (o *CustomerCount200ResponseResult) GetCustomersCount() int32 {
	if o == nil || IsNil(o.CustomersCount) {
		var ret int32
		return ret
	}
	return *o.CustomersCount
}

// GetCustomersCountOk returns a tuple with the CustomersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCount200ResponseResult) GetCustomersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomersCount) {
		return nil, false
	}
	return o.CustomersCount, true
}

// HasCustomersCount returns a boolean if a field has been set.
func (o *CustomerCount200ResponseResult) HasCustomersCount() bool {
	if o != nil && !IsNil(o.CustomersCount) {
		return true
	}

	return false
}

// SetCustomersCount gets a reference to the given int32 and assigns it to the CustomersCount field.
func (o *CustomerCount200ResponseResult) SetCustomersCount(v int32) {
	o.CustomersCount = &v
}

func (o CustomerCount200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerCount200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomersCount) {
		toSerialize["customers_count"] = o.CustomersCount
	}
	return toSerialize, nil
}

type NullableCustomerCount200ResponseResult struct {
	value *CustomerCount200ResponseResult
	isSet bool
}

func (v NullableCustomerCount200ResponseResult) Get() *CustomerCount200ResponseResult {
	return v.value
}

func (v *NullableCustomerCount200ResponseResult) Set(val *CustomerCount200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCount200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCount200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCount200ResponseResult(val *CustomerCount200ResponseResult) *NullableCustomerCount200ResponseResult {
	return &NullableCustomerCount200ResponseResult{value: val, isSet: true}
}

func (v NullableCustomerCount200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCount200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


