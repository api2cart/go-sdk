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
	"bytes"
	"fmt"
)

// checks if the ProductOptionAddValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductOptionAddValuesInner{}

// ProductOptionAddValuesInner struct for ProductOptionAddValuesInner
type ProductOptionAddValuesInner struct {
	Value string `json:"value"`
	DisplayValue *string `json:"display_value,omitempty"`
	IsDefault *bool `json:"is_default,omitempty"`
}

type _ProductOptionAddValuesInner ProductOptionAddValuesInner

// NewProductOptionAddValuesInner instantiates a new ProductOptionAddValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductOptionAddValuesInner(value string) *ProductOptionAddValuesInner {
	this := ProductOptionAddValuesInner{}
	this.Value = value
	var isDefault bool = false
	this.IsDefault = &isDefault
	return &this
}

// NewProductOptionAddValuesInnerWithDefaults instantiates a new ProductOptionAddValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductOptionAddValuesInnerWithDefaults() *ProductOptionAddValuesInner {
	this := ProductOptionAddValuesInner{}
	var isDefault bool = false
	this.IsDefault = &isDefault
	return &this
}

// GetValue returns the Value field value
func (o *ProductOptionAddValuesInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProductOptionAddValuesInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ProductOptionAddValuesInner) SetValue(v string) {
	o.Value = v
}

// GetDisplayValue returns the DisplayValue field value if set, zero value otherwise.
func (o *ProductOptionAddValuesInner) GetDisplayValue() string {
	if o == nil || IsNil(o.DisplayValue) {
		var ret string
		return ret
	}
	return *o.DisplayValue
}

// GetDisplayValueOk returns a tuple with the DisplayValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductOptionAddValuesInner) GetDisplayValueOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayValue) {
		return nil, false
	}
	return o.DisplayValue, true
}

// HasDisplayValue returns a boolean if a field has been set.
func (o *ProductOptionAddValuesInner) HasDisplayValue() bool {
	if o != nil && !IsNil(o.DisplayValue) {
		return true
	}

	return false
}

// SetDisplayValue gets a reference to the given string and assigns it to the DisplayValue field.
func (o *ProductOptionAddValuesInner) SetDisplayValue(v string) {
	o.DisplayValue = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ProductOptionAddValuesInner) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductOptionAddValuesInner) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ProductOptionAddValuesInner) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ProductOptionAddValuesInner) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o ProductOptionAddValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductOptionAddValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.DisplayValue) {
		toSerialize["display_value"] = o.DisplayValue
	}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	return toSerialize, nil
}

func (o *ProductOptionAddValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProductOptionAddValuesInner := _ProductOptionAddValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductOptionAddValuesInner)

	if err != nil {
		return err
	}

	*o = ProductOptionAddValuesInner(varProductOptionAddValuesInner)

	return err
}

type NullableProductOptionAddValuesInner struct {
	value *ProductOptionAddValuesInner
	isSet bool
}

func (v NullableProductOptionAddValuesInner) Get() *ProductOptionAddValuesInner {
	return v.value
}

func (v *NullableProductOptionAddValuesInner) Set(val *ProductOptionAddValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductOptionAddValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductOptionAddValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductOptionAddValuesInner(val *ProductOptionAddValuesInner) *NullableProductOptionAddValuesInner {
	return &NullableProductOptionAddValuesInner{value: val, isSet: true}
}

func (v NullableProductOptionAddValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductOptionAddValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


