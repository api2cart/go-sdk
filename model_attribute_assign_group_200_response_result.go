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

// checks if the AttributeAssignGroup200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeAssignGroup200ResponseResult{}

// AttributeAssignGroup200ResponseResult struct for AttributeAssignGroup200ResponseResult
type AttributeAssignGroup200ResponseResult struct {
	Assigned *string `json:"assigned,omitempty"`
}

// NewAttributeAssignGroup200ResponseResult instantiates a new AttributeAssignGroup200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeAssignGroup200ResponseResult() *AttributeAssignGroup200ResponseResult {
	this := AttributeAssignGroup200ResponseResult{}
	return &this
}

// NewAttributeAssignGroup200ResponseResultWithDefaults instantiates a new AttributeAssignGroup200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeAssignGroup200ResponseResultWithDefaults() *AttributeAssignGroup200ResponseResult {
	this := AttributeAssignGroup200ResponseResult{}
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *AttributeAssignGroup200ResponseResult) GetAssigned() string {
	if o == nil || IsNil(o.Assigned) {
		var ret string
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeAssignGroup200ResponseResult) GetAssignedOk() (*string, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *AttributeAssignGroup200ResponseResult) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given string and assigns it to the Assigned field.
func (o *AttributeAssignGroup200ResponseResult) SetAssigned(v string) {
	o.Assigned = &v
}

func (o AttributeAssignGroup200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeAssignGroup200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	return toSerialize, nil
}

type NullableAttributeAssignGroup200ResponseResult struct {
	value *AttributeAssignGroup200ResponseResult
	isSet bool
}

func (v NullableAttributeAssignGroup200ResponseResult) Get() *AttributeAssignGroup200ResponseResult {
	return v.value
}

func (v *NullableAttributeAssignGroup200ResponseResult) Set(val *AttributeAssignGroup200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeAssignGroup200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeAssignGroup200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeAssignGroup200ResponseResult(val *AttributeAssignGroup200ResponseResult) *NullableAttributeAssignGroup200ResponseResult {
	return &NullableAttributeAssignGroup200ResponseResult{value: val, isSet: true}
}

func (v NullableAttributeAssignGroup200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeAssignGroup200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


