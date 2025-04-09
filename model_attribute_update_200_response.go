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

// checks if the AttributeUpdate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeUpdate200Response{}

// AttributeUpdate200Response struct for AttributeUpdate200Response
type AttributeUpdate200Response struct {
	ReturnCode *int32 `json:"return_code,omitempty"`
	ReturnMessage *string `json:"return_message,omitempty"`
	Result *AttributeUpdate200ResponseResult `json:"result,omitempty"`
}

// NewAttributeUpdate200Response instantiates a new AttributeUpdate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeUpdate200Response() *AttributeUpdate200Response {
	this := AttributeUpdate200Response{}
	return &this
}

// NewAttributeUpdate200ResponseWithDefaults instantiates a new AttributeUpdate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeUpdate200ResponseWithDefaults() *AttributeUpdate200Response {
	this := AttributeUpdate200Response{}
	return &this
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *AttributeUpdate200Response) GetReturnCode() int32 {
	if o == nil || IsNil(o.ReturnCode) {
		var ret int32
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeUpdate200Response) GetReturnCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnCode) {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *AttributeUpdate200Response) HasReturnCode() bool {
	if o != nil && !IsNil(o.ReturnCode) {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given int32 and assigns it to the ReturnCode field.
func (o *AttributeUpdate200Response) SetReturnCode(v int32) {
	o.ReturnCode = &v
}

// GetReturnMessage returns the ReturnMessage field value if set, zero value otherwise.
func (o *AttributeUpdate200Response) GetReturnMessage() string {
	if o == nil || IsNil(o.ReturnMessage) {
		var ret string
		return ret
	}
	return *o.ReturnMessage
}

// GetReturnMessageOk returns a tuple with the ReturnMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeUpdate200Response) GetReturnMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnMessage) {
		return nil, false
	}
	return o.ReturnMessage, true
}

// HasReturnMessage returns a boolean if a field has been set.
func (o *AttributeUpdate200Response) HasReturnMessage() bool {
	if o != nil && !IsNil(o.ReturnMessage) {
		return true
	}

	return false
}

// SetReturnMessage gets a reference to the given string and assigns it to the ReturnMessage field.
func (o *AttributeUpdate200Response) SetReturnMessage(v string) {
	o.ReturnMessage = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AttributeUpdate200Response) GetResult() AttributeUpdate200ResponseResult {
	if o == nil || IsNil(o.Result) {
		var ret AttributeUpdate200ResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeUpdate200Response) GetResultOk() (*AttributeUpdate200ResponseResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AttributeUpdate200Response) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given AttributeUpdate200ResponseResult and assigns it to the Result field.
func (o *AttributeUpdate200Response) SetResult(v AttributeUpdate200ResponseResult) {
	o.Result = &v
}

func (o AttributeUpdate200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeUpdate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnCode) {
		toSerialize["return_code"] = o.ReturnCode
	}
	if !IsNil(o.ReturnMessage) {
		toSerialize["return_message"] = o.ReturnMessage
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableAttributeUpdate200Response struct {
	value *AttributeUpdate200Response
	isSet bool
}

func (v NullableAttributeUpdate200Response) Get() *AttributeUpdate200Response {
	return v.value
}

func (v *NullableAttributeUpdate200Response) Set(val *AttributeUpdate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeUpdate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeUpdate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeUpdate200Response(val *AttributeUpdate200Response) *NullableAttributeUpdate200Response {
	return &NullableAttributeUpdate200Response{value: val, isSet: true}
}

func (v NullableAttributeUpdate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeUpdate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


