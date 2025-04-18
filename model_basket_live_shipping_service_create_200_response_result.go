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

// checks if the BasketLiveShippingServiceCreate200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasketLiveShippingServiceCreate200ResponseResult{}

// BasketLiveShippingServiceCreate200ResponseResult struct for BasketLiveShippingServiceCreate200ResponseResult
type BasketLiveShippingServiceCreate200ResponseResult struct {
	Id *int32 `json:"id,omitempty"`
}

// NewBasketLiveShippingServiceCreate200ResponseResult instantiates a new BasketLiveShippingServiceCreate200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasketLiveShippingServiceCreate200ResponseResult() *BasketLiveShippingServiceCreate200ResponseResult {
	this := BasketLiveShippingServiceCreate200ResponseResult{}
	return &this
}

// NewBasketLiveShippingServiceCreate200ResponseResultWithDefaults instantiates a new BasketLiveShippingServiceCreate200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasketLiveShippingServiceCreate200ResponseResultWithDefaults() *BasketLiveShippingServiceCreate200ResponseResult {
	this := BasketLiveShippingServiceCreate200ResponseResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BasketLiveShippingServiceCreate200ResponseResult) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasketLiveShippingServiceCreate200ResponseResult) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BasketLiveShippingServiceCreate200ResponseResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BasketLiveShippingServiceCreate200ResponseResult) SetId(v int32) {
	o.Id = &v
}

func (o BasketLiveShippingServiceCreate200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasketLiveShippingServiceCreate200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableBasketLiveShippingServiceCreate200ResponseResult struct {
	value *BasketLiveShippingServiceCreate200ResponseResult
	isSet bool
}

func (v NullableBasketLiveShippingServiceCreate200ResponseResult) Get() *BasketLiveShippingServiceCreate200ResponseResult {
	return v.value
}

func (v *NullableBasketLiveShippingServiceCreate200ResponseResult) Set(val *BasketLiveShippingServiceCreate200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBasketLiveShippingServiceCreate200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBasketLiveShippingServiceCreate200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasketLiveShippingServiceCreate200ResponseResult(val *BasketLiveShippingServiceCreate200ResponseResult) *NullableBasketLiveShippingServiceCreate200ResponseResult {
	return &NullableBasketLiveShippingServiceCreate200ResponseResult{value: val, isSet: true}
}

func (v NullableBasketLiveShippingServiceCreate200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasketLiveShippingServiceCreate200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


