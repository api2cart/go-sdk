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

// checks if the ResponseCategoryListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCategoryListResult{}

// ResponseCategoryListResult struct for ResponseCategoryListResult
type ResponseCategoryListResult struct {
	CategoriesCount *int32 `json:"categories_count,omitempty"`
	Category []Category `json:"category,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewResponseCategoryListResult instantiates a new ResponseCategoryListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCategoryListResult() *ResponseCategoryListResult {
	this := ResponseCategoryListResult{}
	return &this
}

// NewResponseCategoryListResultWithDefaults instantiates a new ResponseCategoryListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCategoryListResultWithDefaults() *ResponseCategoryListResult {
	this := ResponseCategoryListResult{}
	return &this
}

// GetCategoriesCount returns the CategoriesCount field value if set, zero value otherwise.
func (o *ResponseCategoryListResult) GetCategoriesCount() int32 {
	if o == nil || IsNil(o.CategoriesCount) {
		var ret int32
		return ret
	}
	return *o.CategoriesCount
}

// GetCategoriesCountOk returns a tuple with the CategoriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCategoryListResult) GetCategoriesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CategoriesCount) {
		return nil, false
	}
	return o.CategoriesCount, true
}

// HasCategoriesCount returns a boolean if a field has been set.
func (o *ResponseCategoryListResult) HasCategoriesCount() bool {
	if o != nil && !IsNil(o.CategoriesCount) {
		return true
	}

	return false
}

// SetCategoriesCount gets a reference to the given int32 and assigns it to the CategoriesCount field.
func (o *ResponseCategoryListResult) SetCategoriesCount(v int32) {
	o.CategoriesCount = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ResponseCategoryListResult) GetCategory() []Category {
	if o == nil || IsNil(o.Category) {
		var ret []Category
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCategoryListResult) GetCategoryOk() ([]Category, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ResponseCategoryListResult) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given []Category and assigns it to the Category field.
func (o *ResponseCategoryListResult) SetCategory(v []Category) {
	o.Category = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *ResponseCategoryListResult) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCategoryListResult) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ResponseCategoryListResult) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ResponseCategoryListResult) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ResponseCategoryListResult) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCategoryListResult) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ResponseCategoryListResult) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ResponseCategoryListResult) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ResponseCategoryListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCategoryListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoriesCount) {
		toSerialize["categories_count"] = o.CategoriesCount
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableResponseCategoryListResult struct {
	value *ResponseCategoryListResult
	isSet bool
}

func (v NullableResponseCategoryListResult) Get() *ResponseCategoryListResult {
	return v.value
}

func (v *NullableResponseCategoryListResult) Set(val *ResponseCategoryListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCategoryListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCategoryListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCategoryListResult(val *ResponseCategoryListResult) *NullableResponseCategoryListResult {
	return &NullableResponseCategoryListResult{value: val, isSet: true}
}

func (v NullableResponseCategoryListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCategoryListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


