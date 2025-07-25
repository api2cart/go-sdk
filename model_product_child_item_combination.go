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

// checks if the ProductChildItemCombination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductChildItemCombination{}

// ProductChildItemCombination struct for ProductChildItemCombination
type ProductChildItemCombination struct {
	OptionId *string `json:"option_id,omitempty"`
	OptionValueId *string `json:"option_value_id,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewProductChildItemCombination instantiates a new ProductChildItemCombination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductChildItemCombination() *ProductChildItemCombination {
	this := ProductChildItemCombination{}
	return &this
}

// NewProductChildItemCombinationWithDefaults instantiates a new ProductChildItemCombination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductChildItemCombinationWithDefaults() *ProductChildItemCombination {
	this := ProductChildItemCombination{}
	return &this
}

// GetOptionId returns the OptionId field value if set, zero value otherwise.
func (o *ProductChildItemCombination) GetOptionId() string {
	if o == nil || IsNil(o.OptionId) {
		var ret string
		return ret
	}
	return *o.OptionId
}

// GetOptionIdOk returns a tuple with the OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductChildItemCombination) GetOptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptionId) {
		return nil, false
	}
	return o.OptionId, true
}

// HasOptionId returns a boolean if a field has been set.
func (o *ProductChildItemCombination) HasOptionId() bool {
	if o != nil && !IsNil(o.OptionId) {
		return true
	}

	return false
}

// SetOptionId gets a reference to the given string and assigns it to the OptionId field.
func (o *ProductChildItemCombination) SetOptionId(v string) {
	o.OptionId = &v
}

// GetOptionValueId returns the OptionValueId field value if set, zero value otherwise.
func (o *ProductChildItemCombination) GetOptionValueId() string {
	if o == nil || IsNil(o.OptionValueId) {
		var ret string
		return ret
	}
	return *o.OptionValueId
}

// GetOptionValueIdOk returns a tuple with the OptionValueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductChildItemCombination) GetOptionValueIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptionValueId) {
		return nil, false
	}
	return o.OptionValueId, true
}

// HasOptionValueId returns a boolean if a field has been set.
func (o *ProductChildItemCombination) HasOptionValueId() bool {
	if o != nil && !IsNil(o.OptionValueId) {
		return true
	}

	return false
}

// SetOptionValueId gets a reference to the given string and assigns it to the OptionValueId field.
func (o *ProductChildItemCombination) SetOptionValueId(v string) {
	o.OptionValueId = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductChildItemCombination) GetAdditionalFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductChildItemCombination) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ProductChildItemCombination) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ProductChildItemCombination) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductChildItemCombination) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductChildItemCombination) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ProductChildItemCombination) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ProductChildItemCombination) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ProductChildItemCombination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductChildItemCombination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptionId) {
		toSerialize["option_id"] = o.OptionId
	}
	if !IsNil(o.OptionValueId) {
		toSerialize["option_value_id"] = o.OptionValueId
	}
	if o.AdditionalFields != nil {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableProductChildItemCombination struct {
	value *ProductChildItemCombination
	isSet bool
}

func (v NullableProductChildItemCombination) Get() *ProductChildItemCombination {
	return v.value
}

func (v *NullableProductChildItemCombination) Set(val *ProductChildItemCombination) {
	v.value = val
	v.isSet = true
}

func (v NullableProductChildItemCombination) IsSet() bool {
	return v.isSet
}

func (v *NullableProductChildItemCombination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductChildItemCombination(val *ProductChildItemCombination) *NullableProductChildItemCombination {
	return &NullableProductChildItemCombination{value: val, isSet: true}
}

func (v NullableProductChildItemCombination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductChildItemCombination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


