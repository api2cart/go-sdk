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

// checks if the ProductAdd200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAdd200ResponseResult{}

// ProductAdd200ResponseResult struct for ProductAdd200ResponseResult
type ProductAdd200ResponseResult struct {
	ProductId *string `json:"product_id,omitempty"`
}

// NewProductAdd200ResponseResult instantiates a new ProductAdd200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAdd200ResponseResult() *ProductAdd200ResponseResult {
	this := ProductAdd200ResponseResult{}
	return &this
}

// NewProductAdd200ResponseResultWithDefaults instantiates a new ProductAdd200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAdd200ResponseResultWithDefaults() *ProductAdd200ResponseResult {
	this := ProductAdd200ResponseResult{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ProductAdd200ResponseResult) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdd200ResponseResult) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ProductAdd200ResponseResult) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ProductAdd200ResponseResult) SetProductId(v string) {
	o.ProductId = &v
}

func (o ProductAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAdd200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	return toSerialize, nil
}

type NullableProductAdd200ResponseResult struct {
	value *ProductAdd200ResponseResult
	isSet bool
}

func (v NullableProductAdd200ResponseResult) Get() *ProductAdd200ResponseResult {
	return v.value
}

func (v *NullableProductAdd200ResponseResult) Set(val *ProductAdd200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAdd200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAdd200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAdd200ResponseResult(val *ProductAdd200ResponseResult) *NullableProductAdd200ResponseResult {
	return &NullableProductAdd200ResponseResult{value: val, isSet: true}
}

func (v NullableProductAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAdd200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


