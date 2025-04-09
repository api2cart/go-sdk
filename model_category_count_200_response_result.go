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

// checks if the CategoryCount200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryCount200ResponseResult{}

// CategoryCount200ResponseResult struct for CategoryCount200ResponseResult
type CategoryCount200ResponseResult struct {
	CategoriesCount *int32 `json:"categories_count,omitempty"`
}

// NewCategoryCount200ResponseResult instantiates a new CategoryCount200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryCount200ResponseResult() *CategoryCount200ResponseResult {
	this := CategoryCount200ResponseResult{}
	return &this
}

// NewCategoryCount200ResponseResultWithDefaults instantiates a new CategoryCount200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryCount200ResponseResultWithDefaults() *CategoryCount200ResponseResult {
	this := CategoryCount200ResponseResult{}
	return &this
}

// GetCategoriesCount returns the CategoriesCount field value if set, zero value otherwise.
func (o *CategoryCount200ResponseResult) GetCategoriesCount() int32 {
	if o == nil || IsNil(o.CategoriesCount) {
		var ret int32
		return ret
	}
	return *o.CategoriesCount
}

// GetCategoriesCountOk returns a tuple with the CategoriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryCount200ResponseResult) GetCategoriesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CategoriesCount) {
		return nil, false
	}
	return o.CategoriesCount, true
}

// HasCategoriesCount returns a boolean if a field has been set.
func (o *CategoryCount200ResponseResult) HasCategoriesCount() bool {
	if o != nil && !IsNil(o.CategoriesCount) {
		return true
	}

	return false
}

// SetCategoriesCount gets a reference to the given int32 and assigns it to the CategoriesCount field.
func (o *CategoryCount200ResponseResult) SetCategoriesCount(v int32) {
	o.CategoriesCount = &v
}

func (o CategoryCount200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryCount200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoriesCount) {
		toSerialize["categories_count"] = o.CategoriesCount
	}
	return toSerialize, nil
}

type NullableCategoryCount200ResponseResult struct {
	value *CategoryCount200ResponseResult
	isSet bool
}

func (v NullableCategoryCount200ResponseResult) Get() *CategoryCount200ResponseResult {
	return v.value
}

func (v *NullableCategoryCount200ResponseResult) Set(val *CategoryCount200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryCount200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryCount200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryCount200ResponseResult(val *CategoryCount200ResponseResult) *NullableCategoryCount200ResponseResult {
	return &NullableCategoryCount200ResponseResult{value: val, isSet: true}
}

func (v NullableCategoryCount200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryCount200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


