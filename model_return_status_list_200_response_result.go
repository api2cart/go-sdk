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

// checks if the ReturnStatusList200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnStatusList200ResponseResult{}

// ReturnStatusList200ResponseResult struct for ReturnStatusList200ResponseResult
type ReturnStatusList200ResponseResult struct {
	ReturnStatuses []OrderFinancialStatusList200ResponseResultOrderFinancialStatusesInner `json:"return_statuses,omitempty"`
}

// NewReturnStatusList200ResponseResult instantiates a new ReturnStatusList200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnStatusList200ResponseResult() *ReturnStatusList200ResponseResult {
	this := ReturnStatusList200ResponseResult{}
	return &this
}

// NewReturnStatusList200ResponseResultWithDefaults instantiates a new ReturnStatusList200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnStatusList200ResponseResultWithDefaults() *ReturnStatusList200ResponseResult {
	this := ReturnStatusList200ResponseResult{}
	return &this
}

// GetReturnStatuses returns the ReturnStatuses field value if set, zero value otherwise.
func (o *ReturnStatusList200ResponseResult) GetReturnStatuses() []OrderFinancialStatusList200ResponseResultOrderFinancialStatusesInner {
	if o == nil || IsNil(o.ReturnStatuses) {
		var ret []OrderFinancialStatusList200ResponseResultOrderFinancialStatusesInner
		return ret
	}
	return o.ReturnStatuses
}

// GetReturnStatusesOk returns a tuple with the ReturnStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnStatusList200ResponseResult) GetReturnStatusesOk() ([]OrderFinancialStatusList200ResponseResultOrderFinancialStatusesInner, bool) {
	if o == nil || IsNil(o.ReturnStatuses) {
		return nil, false
	}
	return o.ReturnStatuses, true
}

// HasReturnStatuses returns a boolean if a field has been set.
func (o *ReturnStatusList200ResponseResult) HasReturnStatuses() bool {
	if o != nil && !IsNil(o.ReturnStatuses) {
		return true
	}

	return false
}

// SetReturnStatuses gets a reference to the given []OrderFinancialStatusList200ResponseResultOrderFinancialStatusesInner and assigns it to the ReturnStatuses field.
func (o *ReturnStatusList200ResponseResult) SetReturnStatuses(v []OrderFinancialStatusList200ResponseResultOrderFinancialStatusesInner) {
	o.ReturnStatuses = v
}

func (o ReturnStatusList200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnStatusList200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnStatuses) {
		toSerialize["return_statuses"] = o.ReturnStatuses
	}
	return toSerialize, nil
}

type NullableReturnStatusList200ResponseResult struct {
	value *ReturnStatusList200ResponseResult
	isSet bool
}

func (v NullableReturnStatusList200ResponseResult) Get() *ReturnStatusList200ResponseResult {
	return v.value
}

func (v *NullableReturnStatusList200ResponseResult) Set(val *ReturnStatusList200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnStatusList200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnStatusList200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnStatusList200ResponseResult(val *ReturnStatusList200ResponseResult) *NullableReturnStatusList200ResponseResult {
	return &NullableReturnStatusList200ResponseResult{value: val, isSet: true}
}

func (v NullableReturnStatusList200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnStatusList200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


