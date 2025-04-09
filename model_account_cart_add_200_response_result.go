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

// checks if the AccountCartAdd200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountCartAdd200ResponseResult{}

// AccountCartAdd200ResponseResult struct for AccountCartAdd200ResponseResult
type AccountCartAdd200ResponseResult struct {
	StoreKey *string `json:"store_key,omitempty"`
}

// NewAccountCartAdd200ResponseResult instantiates a new AccountCartAdd200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCartAdd200ResponseResult() *AccountCartAdd200ResponseResult {
	this := AccountCartAdd200ResponseResult{}
	return &this
}

// NewAccountCartAdd200ResponseResultWithDefaults instantiates a new AccountCartAdd200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCartAdd200ResponseResultWithDefaults() *AccountCartAdd200ResponseResult {
	this := AccountCartAdd200ResponseResult{}
	return &this
}

// GetStoreKey returns the StoreKey field value if set, zero value otherwise.
func (o *AccountCartAdd200ResponseResult) GetStoreKey() string {
	if o == nil || IsNil(o.StoreKey) {
		var ret string
		return ret
	}
	return *o.StoreKey
}

// GetStoreKeyOk returns a tuple with the StoreKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCartAdd200ResponseResult) GetStoreKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StoreKey) {
		return nil, false
	}
	return o.StoreKey, true
}

// HasStoreKey returns a boolean if a field has been set.
func (o *AccountCartAdd200ResponseResult) HasStoreKey() bool {
	if o != nil && !IsNil(o.StoreKey) {
		return true
	}

	return false
}

// SetStoreKey gets a reference to the given string and assigns it to the StoreKey field.
func (o *AccountCartAdd200ResponseResult) SetStoreKey(v string) {
	o.StoreKey = &v
}

func (o AccountCartAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCartAdd200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StoreKey) {
		toSerialize["store_key"] = o.StoreKey
	}
	return toSerialize, nil
}

type NullableAccountCartAdd200ResponseResult struct {
	value *AccountCartAdd200ResponseResult
	isSet bool
}

func (v NullableAccountCartAdd200ResponseResult) Get() *AccountCartAdd200ResponseResult {
	return v.value
}

func (v *NullableAccountCartAdd200ResponseResult) Set(val *AccountCartAdd200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCartAdd200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCartAdd200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCartAdd200ResponseResult(val *AccountCartAdd200ResponseResult) *NullableAccountCartAdd200ResponseResult {
	return &NullableAccountCartAdd200ResponseResult{value: val, isSet: true}
}

func (v NullableAccountCartAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCartAdd200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


