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

// checks if the ModelResponseCategoryList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelResponseCategoryList{}

// ModelResponseCategoryList struct for ModelResponseCategoryList
type ModelResponseCategoryList struct {
	ReturnCode NullableInt32 `json:"return_code,omitempty"`
	ReturnMessage NullableString `json:"return_message,omitempty"`
	Pagination NullablePagination `json:"pagination,omitempty"`
	Result NullableResponseCategoryListResult `json:"result,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewModelResponseCategoryList instantiates a new ModelResponseCategoryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelResponseCategoryList() *ModelResponseCategoryList {
	this := ModelResponseCategoryList{}
	return &this
}

// NewModelResponseCategoryListWithDefaults instantiates a new ModelResponseCategoryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelResponseCategoryListWithDefaults() *ModelResponseCategoryList {
	this := ModelResponseCategoryList{}
	return &this
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelResponseCategoryList) GetReturnCode() int32 {
	if o == nil || IsNil(o.ReturnCode.Get()) {
		var ret int32
		return ret
	}
	return *o.ReturnCode.Get()
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelResponseCategoryList) GetReturnCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnCode.Get(), o.ReturnCode.IsSet()
}

// HasReturnCode returns a boolean if a field has been set.
func (o *ModelResponseCategoryList) HasReturnCode() bool {
	if o != nil && o.ReturnCode.IsSet() {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given NullableInt32 and assigns it to the ReturnCode field.
func (o *ModelResponseCategoryList) SetReturnCode(v int32) {
	o.ReturnCode.Set(&v)
}
// SetReturnCodeNil sets the value for ReturnCode to be an explicit nil
func (o *ModelResponseCategoryList) SetReturnCodeNil() {
	o.ReturnCode.Set(nil)
}

// UnsetReturnCode ensures that no value is present for ReturnCode, not even an explicit nil
func (o *ModelResponseCategoryList) UnsetReturnCode() {
	o.ReturnCode.Unset()
}

// GetReturnMessage returns the ReturnMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelResponseCategoryList) GetReturnMessage() string {
	if o == nil || IsNil(o.ReturnMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ReturnMessage.Get()
}

// GetReturnMessageOk returns a tuple with the ReturnMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelResponseCategoryList) GetReturnMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnMessage.Get(), o.ReturnMessage.IsSet()
}

// HasReturnMessage returns a boolean if a field has been set.
func (o *ModelResponseCategoryList) HasReturnMessage() bool {
	if o != nil && o.ReturnMessage.IsSet() {
		return true
	}

	return false
}

// SetReturnMessage gets a reference to the given NullableString and assigns it to the ReturnMessage field.
func (o *ModelResponseCategoryList) SetReturnMessage(v string) {
	o.ReturnMessage.Set(&v)
}
// SetReturnMessageNil sets the value for ReturnMessage to be an explicit nil
func (o *ModelResponseCategoryList) SetReturnMessageNil() {
	o.ReturnMessage.Set(nil)
}

// UnsetReturnMessage ensures that no value is present for ReturnMessage, not even an explicit nil
func (o *ModelResponseCategoryList) UnsetReturnMessage() {
	o.ReturnMessage.Unset()
}

// GetPagination returns the Pagination field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelResponseCategoryList) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination.Get()) {
		var ret Pagination
		return ret
	}
	return *o.Pagination.Get()
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelResponseCategoryList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pagination.Get(), o.Pagination.IsSet()
}

// HasPagination returns a boolean if a field has been set.
func (o *ModelResponseCategoryList) HasPagination() bool {
	if o != nil && o.Pagination.IsSet() {
		return true
	}

	return false
}

// SetPagination gets a reference to the given NullablePagination and assigns it to the Pagination field.
func (o *ModelResponseCategoryList) SetPagination(v Pagination) {
	o.Pagination.Set(&v)
}
// SetPaginationNil sets the value for Pagination to be an explicit nil
func (o *ModelResponseCategoryList) SetPaginationNil() {
	o.Pagination.Set(nil)
}

// UnsetPagination ensures that no value is present for Pagination, not even an explicit nil
func (o *ModelResponseCategoryList) UnsetPagination() {
	o.Pagination.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelResponseCategoryList) GetResult() ResponseCategoryListResult {
	if o == nil || IsNil(o.Result.Get()) {
		var ret ResponseCategoryListResult
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelResponseCategoryList) GetResultOk() (*ResponseCategoryListResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *ModelResponseCategoryList) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableResponseCategoryListResult and assigns it to the Result field.
func (o *ModelResponseCategoryList) SetResult(v ResponseCategoryListResult) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *ModelResponseCategoryList) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *ModelResponseCategoryList) UnsetResult() {
	o.Result.Unset()
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelResponseCategoryList) GetAdditionalFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelResponseCategoryList) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ModelResponseCategoryList) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ModelResponseCategoryList) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelResponseCategoryList) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelResponseCategoryList) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ModelResponseCategoryList) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ModelResponseCategoryList) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ModelResponseCategoryList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelResponseCategoryList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReturnCode.IsSet() {
		toSerialize["return_code"] = o.ReturnCode.Get()
	}
	if o.ReturnMessage.IsSet() {
		toSerialize["return_message"] = o.ReturnMessage.Get()
	}
	if o.Pagination.IsSet() {
		toSerialize["pagination"] = o.Pagination.Get()
	}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}
	if o.AdditionalFields != nil {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableModelResponseCategoryList struct {
	value *ModelResponseCategoryList
	isSet bool
}

func (v NullableModelResponseCategoryList) Get() *ModelResponseCategoryList {
	return v.value
}

func (v *NullableModelResponseCategoryList) Set(val *ModelResponseCategoryList) {
	v.value = val
	v.isSet = true
}

func (v NullableModelResponseCategoryList) IsSet() bool {
	return v.isSet
}

func (v *NullableModelResponseCategoryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelResponseCategoryList(val *ModelResponseCategoryList) *NullableModelResponseCategoryList {
	return &NullableModelResponseCategoryList{value: val, isSet: true}
}

func (v NullableModelResponseCategoryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelResponseCategoryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


