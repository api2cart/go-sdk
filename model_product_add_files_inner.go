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

// checks if the ProductAddFilesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddFilesInner{}

// ProductAddFilesInner struct for ProductAddFilesInner
type ProductAddFilesInner struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type _ProductAddFilesInner ProductAddFilesInner

// NewProductAddFilesInner instantiates a new ProductAddFilesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddFilesInner(name string, url string) *ProductAddFilesInner {
	this := ProductAddFilesInner{}
	this.Name = name
	this.Url = url
	return &this
}

// NewProductAddFilesInnerWithDefaults instantiates a new ProductAddFilesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddFilesInnerWithDefaults() *ProductAddFilesInner {
	this := ProductAddFilesInner{}
	return &this
}

// GetName returns the Name field value
func (o *ProductAddFilesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductAddFilesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductAddFilesInner) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *ProductAddFilesInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ProductAddFilesInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ProductAddFilesInner) SetUrl(v string) {
	o.Url = v
}

func (o ProductAddFilesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddFilesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *ProductAddFilesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"url",
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

	varProductAddFilesInner := _ProductAddFilesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductAddFilesInner)

	if err != nil {
		return err
	}

	*o = ProductAddFilesInner(varProductAddFilesInner)

	return err
}

type NullableProductAddFilesInner struct {
	value *ProductAddFilesInner
	isSet bool
}

func (v NullableProductAddFilesInner) Get() *ProductAddFilesInner {
	return v.value
}

func (v *NullableProductAddFilesInner) Set(val *ProductAddFilesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddFilesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddFilesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddFilesInner(val *ProductAddFilesInner) *NullableProductAddFilesInner {
	return &NullableProductAddFilesInner{value: val, isSet: true}
}

func (v NullableProductAddFilesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddFilesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


