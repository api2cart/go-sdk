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

// checks if the CartGiftcardAdd200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartGiftcardAdd200ResponseResult{}

// CartGiftcardAdd200ResponseResult struct for CartGiftcardAdd200ResponseResult
type CartGiftcardAdd200ResponseResult struct {
	Id *string `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewCartGiftcardAdd200ResponseResult instantiates a new CartGiftcardAdd200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartGiftcardAdd200ResponseResult() *CartGiftcardAdd200ResponseResult {
	this := CartGiftcardAdd200ResponseResult{}
	return &this
}

// NewCartGiftcardAdd200ResponseResultWithDefaults instantiates a new CartGiftcardAdd200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartGiftcardAdd200ResponseResultWithDefaults() *CartGiftcardAdd200ResponseResult {
	this := CartGiftcardAdd200ResponseResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CartGiftcardAdd200ResponseResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartGiftcardAdd200ResponseResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CartGiftcardAdd200ResponseResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CartGiftcardAdd200ResponseResult) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CartGiftcardAdd200ResponseResult) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartGiftcardAdd200ResponseResult) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CartGiftcardAdd200ResponseResult) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CartGiftcardAdd200ResponseResult) SetCode(v string) {
	o.Code = &v
}

func (o CartGiftcardAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartGiftcardAdd200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableCartGiftcardAdd200ResponseResult struct {
	value *CartGiftcardAdd200ResponseResult
	isSet bool
}

func (v NullableCartGiftcardAdd200ResponseResult) Get() *CartGiftcardAdd200ResponseResult {
	return v.value
}

func (v *NullableCartGiftcardAdd200ResponseResult) Set(val *CartGiftcardAdd200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCartGiftcardAdd200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCartGiftcardAdd200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartGiftcardAdd200ResponseResult(val *CartGiftcardAdd200ResponseResult) *NullableCartGiftcardAdd200ResponseResult {
	return &NullableCartGiftcardAdd200ResponseResult{value: val, isSet: true}
}

func (v NullableCartGiftcardAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartGiftcardAdd200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


