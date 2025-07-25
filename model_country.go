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

// checks if the Country type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Country{}

// Country struct for Country
type Country struct {
	Code2 *string `json:"code2,omitempty"`
	Code3 *string `json:"code3,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewCountry instantiates a new Country object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountry() *Country {
	this := Country{}
	return &this
}

// NewCountryWithDefaults instantiates a new Country object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryWithDefaults() *Country {
	this := Country{}
	return &this
}

// GetCode2 returns the Code2 field value if set, zero value otherwise.
func (o *Country) GetCode2() string {
	if o == nil || IsNil(o.Code2) {
		var ret string
		return ret
	}
	return *o.Code2
}

// GetCode2Ok returns a tuple with the Code2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetCode2Ok() (*string, bool) {
	if o == nil || IsNil(o.Code2) {
		return nil, false
	}
	return o.Code2, true
}

// HasCode2 returns a boolean if a field has been set.
func (o *Country) HasCode2() bool {
	if o != nil && !IsNil(o.Code2) {
		return true
	}

	return false
}

// SetCode2 gets a reference to the given string and assigns it to the Code2 field.
func (o *Country) SetCode2(v string) {
	o.Code2 = &v
}

// GetCode3 returns the Code3 field value if set, zero value otherwise.
func (o *Country) GetCode3() string {
	if o == nil || IsNil(o.Code3) {
		var ret string
		return ret
	}
	return *o.Code3
}

// GetCode3Ok returns a tuple with the Code3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetCode3Ok() (*string, bool) {
	if o == nil || IsNil(o.Code3) {
		return nil, false
	}
	return o.Code3, true
}

// HasCode3 returns a boolean if a field has been set.
func (o *Country) HasCode3() bool {
	if o != nil && !IsNil(o.Code3) {
		return true
	}

	return false
}

// SetCode3 gets a reference to the given string and assigns it to the Code3 field.
func (o *Country) SetCode3(v string) {
	o.Code3 = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Country) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Country) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Country) SetName(v string) {
	o.Name = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Country) GetAdditionalFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Country) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *Country) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *Country) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Country) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Country) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Country) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Country) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Country) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Country) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code2) {
		toSerialize["code2"] = o.Code2
	}
	if !IsNil(o.Code3) {
		toSerialize["code3"] = o.Code3
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.AdditionalFields != nil {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableCountry struct {
	value *Country
	isSet bool
}

func (v NullableCountry) Get() *Country {
	return v.value
}

func (v *NullableCountry) Set(val *Country) {
	v.value = val
	v.isSet = true
}

func (v NullableCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountry(val *Country) *NullableCountry {
	return &NullableCountry{value: val, isSet: true}
}

func (v NullableCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


