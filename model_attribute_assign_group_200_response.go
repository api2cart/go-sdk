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

// checks if the AttributeAssignGroup200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeAssignGroup200Response{}

// AttributeAssignGroup200Response struct for AttributeAssignGroup200Response
type AttributeAssignGroup200Response struct {
	ReturnCode *int32 `json:"return_code,omitempty"`
	ReturnMessage *string `json:"return_message,omitempty"`
	Result *AttributeAssignGroup200ResponseResult `json:"result,omitempty"`
}

// NewAttributeAssignGroup200Response instantiates a new AttributeAssignGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeAssignGroup200Response() *AttributeAssignGroup200Response {
	this := AttributeAssignGroup200Response{}
	return &this
}

// NewAttributeAssignGroup200ResponseWithDefaults instantiates a new AttributeAssignGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeAssignGroup200ResponseWithDefaults() *AttributeAssignGroup200Response {
	this := AttributeAssignGroup200Response{}
	return &this
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *AttributeAssignGroup200Response) GetReturnCode() int32 {
	if o == nil || IsNil(o.ReturnCode) {
		var ret int32
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeAssignGroup200Response) GetReturnCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnCode) {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *AttributeAssignGroup200Response) HasReturnCode() bool {
	if o != nil && !IsNil(o.ReturnCode) {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given int32 and assigns it to the ReturnCode field.
func (o *AttributeAssignGroup200Response) SetReturnCode(v int32) {
	o.ReturnCode = &v
}

// GetReturnMessage returns the ReturnMessage field value if set, zero value otherwise.
func (o *AttributeAssignGroup200Response) GetReturnMessage() string {
	if o == nil || IsNil(o.ReturnMessage) {
		var ret string
		return ret
	}
	return *o.ReturnMessage
}

// GetReturnMessageOk returns a tuple with the ReturnMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeAssignGroup200Response) GetReturnMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnMessage) {
		return nil, false
	}
	return o.ReturnMessage, true
}

// HasReturnMessage returns a boolean if a field has been set.
func (o *AttributeAssignGroup200Response) HasReturnMessage() bool {
	if o != nil && !IsNil(o.ReturnMessage) {
		return true
	}

	return false
}

// SetReturnMessage gets a reference to the given string and assigns it to the ReturnMessage field.
func (o *AttributeAssignGroup200Response) SetReturnMessage(v string) {
	o.ReturnMessage = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AttributeAssignGroup200Response) GetResult() AttributeAssignGroup200ResponseResult {
	if o == nil || IsNil(o.Result) {
		var ret AttributeAssignGroup200ResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeAssignGroup200Response) GetResultOk() (*AttributeAssignGroup200ResponseResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AttributeAssignGroup200Response) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given AttributeAssignGroup200ResponseResult and assigns it to the Result field.
func (o *AttributeAssignGroup200Response) SetResult(v AttributeAssignGroup200ResponseResult) {
	o.Result = &v
}

func (o AttributeAssignGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeAssignGroup200Response) ToMap() (map[string]interface{}, error) {
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

type NullableAttributeAssignGroup200Response struct {
	value *AttributeAssignGroup200Response
	isSet bool
}

func (v NullableAttributeAssignGroup200Response) Get() *AttributeAssignGroup200Response {
	return v.value
}

func (v *NullableAttributeAssignGroup200Response) Set(val *AttributeAssignGroup200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeAssignGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeAssignGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeAssignGroup200Response(val *AttributeAssignGroup200Response) *NullableAttributeAssignGroup200Response {
	return &NullableAttributeAssignGroup200Response{value: val, isSet: true}
}

func (v NullableAttributeAssignGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeAssignGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


