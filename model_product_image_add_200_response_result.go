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

// checks if the ProductImageAdd200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductImageAdd200ResponseResult{}

// ProductImageAdd200ResponseResult struct for ProductImageAdd200ResponseResult
type ProductImageAdd200ResponseResult struct {
	Id *string `json:"id,omitempty"`
	ImagePath *string `json:"image_path,omitempty"`
}

// NewProductImageAdd200ResponseResult instantiates a new ProductImageAdd200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductImageAdd200ResponseResult() *ProductImageAdd200ResponseResult {
	this := ProductImageAdd200ResponseResult{}
	return &this
}

// NewProductImageAdd200ResponseResultWithDefaults instantiates a new ProductImageAdd200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductImageAdd200ResponseResultWithDefaults() *ProductImageAdd200ResponseResult {
	this := ProductImageAdd200ResponseResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductImageAdd200ResponseResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductImageAdd200ResponseResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductImageAdd200ResponseResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductImageAdd200ResponseResult) SetId(v string) {
	o.Id = &v
}

// GetImagePath returns the ImagePath field value if set, zero value otherwise.
func (o *ProductImageAdd200ResponseResult) GetImagePath() string {
	if o == nil || IsNil(o.ImagePath) {
		var ret string
		return ret
	}
	return *o.ImagePath
}

// GetImagePathOk returns a tuple with the ImagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductImageAdd200ResponseResult) GetImagePathOk() (*string, bool) {
	if o == nil || IsNil(o.ImagePath) {
		return nil, false
	}
	return o.ImagePath, true
}

// HasImagePath returns a boolean if a field has been set.
func (o *ProductImageAdd200ResponseResult) HasImagePath() bool {
	if o != nil && !IsNil(o.ImagePath) {
		return true
	}

	return false
}

// SetImagePath gets a reference to the given string and assigns it to the ImagePath field.
func (o *ProductImageAdd200ResponseResult) SetImagePath(v string) {
	o.ImagePath = &v
}

func (o ProductImageAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductImageAdd200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImagePath) {
		toSerialize["image_path"] = o.ImagePath
	}
	return toSerialize, nil
}

type NullableProductImageAdd200ResponseResult struct {
	value *ProductImageAdd200ResponseResult
	isSet bool
}

func (v NullableProductImageAdd200ResponseResult) Get() *ProductImageAdd200ResponseResult {
	return v.value
}

func (v *NullableProductImageAdd200ResponseResult) Set(val *ProductImageAdd200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableProductImageAdd200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableProductImageAdd200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductImageAdd200ResponseResult(val *ProductImageAdd200ResponseResult) *NullableProductImageAdd200ResponseResult {
	return &NullableProductImageAdd200ResponseResult{value: val, isSet: true}
}

func (v NullableProductImageAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductImageAdd200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


