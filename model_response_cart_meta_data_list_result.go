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

// checks if the ResponseCartMetaDataListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCartMetaDataListResult{}

// ResponseCartMetaDataListResult struct for ResponseCartMetaDataListResult
type ResponseCartMetaDataListResult struct {
	TotalCount NullableInt32 `json:"total_count,omitempty"`
	Items []CartMetaData `json:"items,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewResponseCartMetaDataListResult instantiates a new ResponseCartMetaDataListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCartMetaDataListResult() *ResponseCartMetaDataListResult {
	this := ResponseCartMetaDataListResult{}
	return &this
}

// NewResponseCartMetaDataListResultWithDefaults instantiates a new ResponseCartMetaDataListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCartMetaDataListResultWithDefaults() *ResponseCartMetaDataListResult {
	this := ResponseCartMetaDataListResult{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseCartMetaDataListResult) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalCount.Get()
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseCartMetaDataListResult) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCount.Get(), o.TotalCount.IsSet()
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ResponseCartMetaDataListResult) HasTotalCount() bool {
	if o != nil && o.TotalCount.IsSet() {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given NullableInt32 and assigns it to the TotalCount field.
func (o *ResponseCartMetaDataListResult) SetTotalCount(v int32) {
	o.TotalCount.Set(&v)
}
// SetTotalCountNil sets the value for TotalCount to be an explicit nil
func (o *ResponseCartMetaDataListResult) SetTotalCountNil() {
	o.TotalCount.Set(nil)
}

// UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
func (o *ResponseCartMetaDataListResult) UnsetTotalCount() {
	o.TotalCount.Unset()
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ResponseCartMetaDataListResult) GetItems() []CartMetaData {
	if o == nil || IsNil(o.Items) {
		var ret []CartMetaData
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCartMetaDataListResult) GetItemsOk() ([]CartMetaData, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ResponseCartMetaDataListResult) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []CartMetaData and assigns it to the Items field.
func (o *ResponseCartMetaDataListResult) SetItems(v []CartMetaData) {
	o.Items = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseCartMetaDataListResult) GetAdditionalFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseCartMetaDataListResult) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ResponseCartMetaDataListResult) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ResponseCartMetaDataListResult) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseCartMetaDataListResult) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseCartMetaDataListResult) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ResponseCartMetaDataListResult) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ResponseCartMetaDataListResult) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ResponseCartMetaDataListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCartMetaDataListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount.IsSet() {
		toSerialize["total_count"] = o.TotalCount.Get()
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if o.AdditionalFields != nil {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableResponseCartMetaDataListResult struct {
	value *ResponseCartMetaDataListResult
	isSet bool
}

func (v NullableResponseCartMetaDataListResult) Get() *ResponseCartMetaDataListResult {
	return v.value
}

func (v *NullableResponseCartMetaDataListResult) Set(val *ResponseCartMetaDataListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCartMetaDataListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCartMetaDataListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCartMetaDataListResult(val *ResponseCartMetaDataListResult) *NullableResponseCartMetaDataListResult {
	return &NullableResponseCartMetaDataListResult{value: val, isSet: true}
}

func (v NullableResponseCartMetaDataListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCartMetaDataListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


