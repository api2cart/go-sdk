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

// checks if the ResponseProductAttributeListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseProductAttributeListResult{}

// ResponseProductAttributeListResult struct for ResponseProductAttributeListResult
type ResponseProductAttributeListResult struct {
	Attribute []ProductAttribute `json:"attribute,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewResponseProductAttributeListResult instantiates a new ResponseProductAttributeListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseProductAttributeListResult() *ResponseProductAttributeListResult {
	this := ResponseProductAttributeListResult{}
	return &this
}

// NewResponseProductAttributeListResultWithDefaults instantiates a new ResponseProductAttributeListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseProductAttributeListResultWithDefaults() *ResponseProductAttributeListResult {
	this := ResponseProductAttributeListResult{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *ResponseProductAttributeListResult) GetAttribute() []ProductAttribute {
	if o == nil || IsNil(o.Attribute) {
		var ret []ProductAttribute
		return ret
	}
	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseProductAttributeListResult) GetAttributeOk() ([]ProductAttribute, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *ResponseProductAttributeListResult) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given []ProductAttribute and assigns it to the Attribute field.
func (o *ResponseProductAttributeListResult) SetAttribute(v []ProductAttribute) {
	o.Attribute = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *ResponseProductAttributeListResult) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseProductAttributeListResult) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ResponseProductAttributeListResult) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ResponseProductAttributeListResult) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ResponseProductAttributeListResult) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseProductAttributeListResult) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ResponseProductAttributeListResult) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ResponseProductAttributeListResult) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ResponseProductAttributeListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseProductAttributeListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableResponseProductAttributeListResult struct {
	value *ResponseProductAttributeListResult
	isSet bool
}

func (v NullableResponseProductAttributeListResult) Get() *ResponseProductAttributeListResult {
	return v.value
}

func (v *NullableResponseProductAttributeListResult) Set(val *ResponseProductAttributeListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseProductAttributeListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseProductAttributeListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseProductAttributeListResult(val *ResponseProductAttributeListResult) *NullableResponseProductAttributeListResult {
	return &NullableResponseProductAttributeListResult{value: val, isSet: true}
}

func (v NullableResponseProductAttributeListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseProductAttributeListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


