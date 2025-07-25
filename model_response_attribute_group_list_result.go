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

// checks if the ResponseAttributeGroupListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseAttributeGroupListResult{}

// ResponseAttributeGroupListResult struct for ResponseAttributeGroupListResult
type ResponseAttributeGroupListResult struct {
	Group []StoreAttributeGroup `json:"group,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewResponseAttributeGroupListResult instantiates a new ResponseAttributeGroupListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAttributeGroupListResult() *ResponseAttributeGroupListResult {
	this := ResponseAttributeGroupListResult{}
	return &this
}

// NewResponseAttributeGroupListResultWithDefaults instantiates a new ResponseAttributeGroupListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAttributeGroupListResultWithDefaults() *ResponseAttributeGroupListResult {
	this := ResponseAttributeGroupListResult{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ResponseAttributeGroupListResult) GetGroup() []StoreAttributeGroup {
	if o == nil || IsNil(o.Group) {
		var ret []StoreAttributeGroup
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAttributeGroupListResult) GetGroupOk() ([]StoreAttributeGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ResponseAttributeGroupListResult) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given []StoreAttributeGroup and assigns it to the Group field.
func (o *ResponseAttributeGroupListResult) SetGroup(v []StoreAttributeGroup) {
	o.Group = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseAttributeGroupListResult) GetAdditionalFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseAttributeGroupListResult) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ResponseAttributeGroupListResult) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ResponseAttributeGroupListResult) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseAttributeGroupListResult) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseAttributeGroupListResult) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ResponseAttributeGroupListResult) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ResponseAttributeGroupListResult) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ResponseAttributeGroupListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAttributeGroupListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if o.AdditionalFields != nil {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableResponseAttributeGroupListResult struct {
	value *ResponseAttributeGroupListResult
	isSet bool
}

func (v NullableResponseAttributeGroupListResult) Get() *ResponseAttributeGroupListResult {
	return v.value
}

func (v *NullableResponseAttributeGroupListResult) Set(val *ResponseAttributeGroupListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAttributeGroupListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAttributeGroupListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAttributeGroupListResult(val *ResponseAttributeGroupListResult) *NullableResponseAttributeGroupListResult {
	return &NullableResponseAttributeGroupListResult{value: val, isSet: true}
}

func (v NullableResponseAttributeGroupListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAttributeGroupListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


