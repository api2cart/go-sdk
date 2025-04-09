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

// checks if the ProductAttributeValueUnset200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAttributeValueUnset200ResponseResult{}

// ProductAttributeValueUnset200ResponseResult struct for ProductAttributeValueUnset200ResponseResult
type ProductAttributeValueUnset200ResponseResult struct {
	Success *bool `json:"success,omitempty"`
}

// NewProductAttributeValueUnset200ResponseResult instantiates a new ProductAttributeValueUnset200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAttributeValueUnset200ResponseResult() *ProductAttributeValueUnset200ResponseResult {
	this := ProductAttributeValueUnset200ResponseResult{}
	return &this
}

// NewProductAttributeValueUnset200ResponseResultWithDefaults instantiates a new ProductAttributeValueUnset200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAttributeValueUnset200ResponseResultWithDefaults() *ProductAttributeValueUnset200ResponseResult {
	this := ProductAttributeValueUnset200ResponseResult{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ProductAttributeValueUnset200ResponseResult) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAttributeValueUnset200ResponseResult) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ProductAttributeValueUnset200ResponseResult) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ProductAttributeValueUnset200ResponseResult) SetSuccess(v bool) {
	o.Success = &v
}

func (o ProductAttributeValueUnset200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAttributeValueUnset200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableProductAttributeValueUnset200ResponseResult struct {
	value *ProductAttributeValueUnset200ResponseResult
	isSet bool
}

func (v NullableProductAttributeValueUnset200ResponseResult) Get() *ProductAttributeValueUnset200ResponseResult {
	return v.value
}

func (v *NullableProductAttributeValueUnset200ResponseResult) Set(val *ProductAttributeValueUnset200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAttributeValueUnset200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAttributeValueUnset200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAttributeValueUnset200ResponseResult(val *ProductAttributeValueUnset200ResponseResult) *NullableProductAttributeValueUnset200ResponseResult {
	return &NullableProductAttributeValueUnset200ResponseResult{value: val, isSet: true}
}

func (v NullableProductAttributeValueUnset200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAttributeValueUnset200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


