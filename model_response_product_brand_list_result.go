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

// checks if the ResponseProductBrandListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseProductBrandListResult{}

// ResponseProductBrandListResult struct for ResponseProductBrandListResult
type ResponseProductBrandListResult struct {
	TotalCount *int32 `json:"total_count,omitempty"`
	Brands []Brand `json:"brands,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewResponseProductBrandListResult instantiates a new ResponseProductBrandListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseProductBrandListResult() *ResponseProductBrandListResult {
	this := ResponseProductBrandListResult{}
	return &this
}

// NewResponseProductBrandListResultWithDefaults instantiates a new ResponseProductBrandListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseProductBrandListResultWithDefaults() *ResponseProductBrandListResult {
	this := ResponseProductBrandListResult{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ResponseProductBrandListResult) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseProductBrandListResult) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ResponseProductBrandListResult) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ResponseProductBrandListResult) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *ResponseProductBrandListResult) GetBrands() []Brand {
	if o == nil || IsNil(o.Brands) {
		var ret []Brand
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseProductBrandListResult) GetBrandsOk() ([]Brand, bool) {
	if o == nil || IsNil(o.Brands) {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *ResponseProductBrandListResult) HasBrands() bool {
	if o != nil && !IsNil(o.Brands) {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []Brand and assigns it to the Brands field.
func (o *ResponseProductBrandListResult) SetBrands(v []Brand) {
	o.Brands = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *ResponseProductBrandListResult) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseProductBrandListResult) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ResponseProductBrandListResult) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ResponseProductBrandListResult) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ResponseProductBrandListResult) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseProductBrandListResult) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ResponseProductBrandListResult) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ResponseProductBrandListResult) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ResponseProductBrandListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseProductBrandListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	if !IsNil(o.Brands) {
		toSerialize["brands"] = o.Brands
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableResponseProductBrandListResult struct {
	value *ResponseProductBrandListResult
	isSet bool
}

func (v NullableResponseProductBrandListResult) Get() *ResponseProductBrandListResult {
	return v.value
}

func (v *NullableResponseProductBrandListResult) Set(val *ResponseProductBrandListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseProductBrandListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseProductBrandListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseProductBrandListResult(val *ResponseProductBrandListResult) *NullableResponseProductBrandListResult {
	return &NullableResponseProductBrandListResult{value: val, isSet: true}
}

func (v NullableResponseProductBrandListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseProductBrandListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


