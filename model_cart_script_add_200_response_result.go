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

// checks if the CartScriptAdd200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartScriptAdd200ResponseResult{}

// CartScriptAdd200ResponseResult struct for CartScriptAdd200ResponseResult
type CartScriptAdd200ResponseResult struct {
	ScriptId *string `json:"script_id,omitempty"`
}

// NewCartScriptAdd200ResponseResult instantiates a new CartScriptAdd200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartScriptAdd200ResponseResult() *CartScriptAdd200ResponseResult {
	this := CartScriptAdd200ResponseResult{}
	return &this
}

// NewCartScriptAdd200ResponseResultWithDefaults instantiates a new CartScriptAdd200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartScriptAdd200ResponseResultWithDefaults() *CartScriptAdd200ResponseResult {
	this := CartScriptAdd200ResponseResult{}
	return &this
}

// GetScriptId returns the ScriptId field value if set, zero value otherwise.
func (o *CartScriptAdd200ResponseResult) GetScriptId() string {
	if o == nil || IsNil(o.ScriptId) {
		var ret string
		return ret
	}
	return *o.ScriptId
}

// GetScriptIdOk returns a tuple with the ScriptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartScriptAdd200ResponseResult) GetScriptIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScriptId) {
		return nil, false
	}
	return o.ScriptId, true
}

// HasScriptId returns a boolean if a field has been set.
func (o *CartScriptAdd200ResponseResult) HasScriptId() bool {
	if o != nil && !IsNil(o.ScriptId) {
		return true
	}

	return false
}

// SetScriptId gets a reference to the given string and assigns it to the ScriptId field.
func (o *CartScriptAdd200ResponseResult) SetScriptId(v string) {
	o.ScriptId = &v
}

func (o CartScriptAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartScriptAdd200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScriptId) {
		toSerialize["script_id"] = o.ScriptId
	}
	return toSerialize, nil
}

type NullableCartScriptAdd200ResponseResult struct {
	value *CartScriptAdd200ResponseResult
	isSet bool
}

func (v NullableCartScriptAdd200ResponseResult) Get() *CartScriptAdd200ResponseResult {
	return v.value
}

func (v *NullableCartScriptAdd200ResponseResult) Set(val *CartScriptAdd200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCartScriptAdd200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCartScriptAdd200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartScriptAdd200ResponseResult(val *CartScriptAdd200ResponseResult) *NullableCartScriptAdd200ResponseResult {
	return &NullableCartScriptAdd200ResponseResult{value: val, isSet: true}
}

func (v NullableCartScriptAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartScriptAdd200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


