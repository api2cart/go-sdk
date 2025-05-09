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

// checks if the ModelResponseProductChildItemList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelResponseProductChildItemList{}

// ModelResponseProductChildItemList struct for ModelResponseProductChildItemList
type ModelResponseProductChildItemList struct {
	ReturnCode *int32 `json:"return_code,omitempty"`
	ReturnMessage *string `json:"return_message,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Result *ResponseProductChildItemListResult `json:"result,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewModelResponseProductChildItemList instantiates a new ModelResponseProductChildItemList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelResponseProductChildItemList() *ModelResponseProductChildItemList {
	this := ModelResponseProductChildItemList{}
	return &this
}

// NewModelResponseProductChildItemListWithDefaults instantiates a new ModelResponseProductChildItemList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelResponseProductChildItemListWithDefaults() *ModelResponseProductChildItemList {
	this := ModelResponseProductChildItemList{}
	return &this
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *ModelResponseProductChildItemList) GetReturnCode() int32 {
	if o == nil || IsNil(o.ReturnCode) {
		var ret int32
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProductChildItemList) GetReturnCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnCode) {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *ModelResponseProductChildItemList) HasReturnCode() bool {
	if o != nil && !IsNil(o.ReturnCode) {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given int32 and assigns it to the ReturnCode field.
func (o *ModelResponseProductChildItemList) SetReturnCode(v int32) {
	o.ReturnCode = &v
}

// GetReturnMessage returns the ReturnMessage field value if set, zero value otherwise.
func (o *ModelResponseProductChildItemList) GetReturnMessage() string {
	if o == nil || IsNil(o.ReturnMessage) {
		var ret string
		return ret
	}
	return *o.ReturnMessage
}

// GetReturnMessageOk returns a tuple with the ReturnMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProductChildItemList) GetReturnMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnMessage) {
		return nil, false
	}
	return o.ReturnMessage, true
}

// HasReturnMessage returns a boolean if a field has been set.
func (o *ModelResponseProductChildItemList) HasReturnMessage() bool {
	if o != nil && !IsNil(o.ReturnMessage) {
		return true
	}

	return false
}

// SetReturnMessage gets a reference to the given string and assigns it to the ReturnMessage field.
func (o *ModelResponseProductChildItemList) SetReturnMessage(v string) {
	o.ReturnMessage = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ModelResponseProductChildItemList) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProductChildItemList) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ModelResponseProductChildItemList) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ModelResponseProductChildItemList) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ModelResponseProductChildItemList) GetResult() ResponseProductChildItemListResult {
	if o == nil || IsNil(o.Result) {
		var ret ResponseProductChildItemListResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProductChildItemList) GetResultOk() (*ResponseProductChildItemListResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ModelResponseProductChildItemList) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ResponseProductChildItemListResult and assigns it to the Result field.
func (o *ModelResponseProductChildItemList) SetResult(v ResponseProductChildItemListResult) {
	o.Result = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *ModelResponseProductChildItemList) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProductChildItemList) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ModelResponseProductChildItemList) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ModelResponseProductChildItemList) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ModelResponseProductChildItemList) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProductChildItemList) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ModelResponseProductChildItemList) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ModelResponseProductChildItemList) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ModelResponseProductChildItemList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelResponseProductChildItemList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnCode) {
		toSerialize["return_code"] = o.ReturnCode
	}
	if !IsNil(o.ReturnMessage) {
		toSerialize["return_message"] = o.ReturnMessage
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableModelResponseProductChildItemList struct {
	value *ModelResponseProductChildItemList
	isSet bool
}

func (v NullableModelResponseProductChildItemList) Get() *ModelResponseProductChildItemList {
	return v.value
}

func (v *NullableModelResponseProductChildItemList) Set(val *ModelResponseProductChildItemList) {
	v.value = val
	v.isSet = true
}

func (v NullableModelResponseProductChildItemList) IsSet() bool {
	return v.isSet
}

func (v *NullableModelResponseProductChildItemList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelResponseProductChildItemList(val *ModelResponseProductChildItemList) *NullableModelResponseProductChildItemList {
	return &NullableModelResponseProductChildItemList{value: val, isSet: true}
}

func (v NullableModelResponseProductChildItemList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelResponseProductChildItemList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


