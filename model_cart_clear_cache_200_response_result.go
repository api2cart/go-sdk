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

// checks if the CartClearCache200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartClearCache200ResponseResult{}

// CartClearCache200ResponseResult struct for CartClearCache200ResponseResult
type CartClearCache200ResponseResult struct {
	CacheCleared *string `json:"cache_cleared,omitempty"`
}

// NewCartClearCache200ResponseResult instantiates a new CartClearCache200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartClearCache200ResponseResult() *CartClearCache200ResponseResult {
	this := CartClearCache200ResponseResult{}
	return &this
}

// NewCartClearCache200ResponseResultWithDefaults instantiates a new CartClearCache200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartClearCache200ResponseResultWithDefaults() *CartClearCache200ResponseResult {
	this := CartClearCache200ResponseResult{}
	return &this
}

// GetCacheCleared returns the CacheCleared field value if set, zero value otherwise.
func (o *CartClearCache200ResponseResult) GetCacheCleared() string {
	if o == nil || IsNil(o.CacheCleared) {
		var ret string
		return ret
	}
	return *o.CacheCleared
}

// GetCacheClearedOk returns a tuple with the CacheCleared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartClearCache200ResponseResult) GetCacheClearedOk() (*string, bool) {
	if o == nil || IsNil(o.CacheCleared) {
		return nil, false
	}
	return o.CacheCleared, true
}

// HasCacheCleared returns a boolean if a field has been set.
func (o *CartClearCache200ResponseResult) HasCacheCleared() bool {
	if o != nil && !IsNil(o.CacheCleared) {
		return true
	}

	return false
}

// SetCacheCleared gets a reference to the given string and assigns it to the CacheCleared field.
func (o *CartClearCache200ResponseResult) SetCacheCleared(v string) {
	o.CacheCleared = &v
}

func (o CartClearCache200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartClearCache200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CacheCleared) {
		toSerialize["cache_cleared"] = o.CacheCleared
	}
	return toSerialize, nil
}

type NullableCartClearCache200ResponseResult struct {
	value *CartClearCache200ResponseResult
	isSet bool
}

func (v NullableCartClearCache200ResponseResult) Get() *CartClearCache200ResponseResult {
	return v.value
}

func (v *NullableCartClearCache200ResponseResult) Set(val *CartClearCache200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCartClearCache200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCartClearCache200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartClearCache200ResponseResult(val *CartClearCache200ResponseResult) *NullableCartClearCache200ResponseResult {
	return &NullableCartClearCache200ResponseResult{value: val, isSet: true}
}

func (v NullableCartClearCache200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartClearCache200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


