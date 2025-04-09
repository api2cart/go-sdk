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

// checks if the ProductAddSizeChart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddSizeChart{}

// ProductAddSizeChart A size chart for the product. Only one property is supported.
type ProductAddSizeChart struct {
	// Defines a pre-generated size chart template
	Id *string `json:"id,omitempty"`
	// Defines a size chart image URL
	Url *string `json:"url,omitempty"`
}

// NewProductAddSizeChart instantiates a new ProductAddSizeChart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddSizeChart() *ProductAddSizeChart {
	this := ProductAddSizeChart{}
	return &this
}

// NewProductAddSizeChartWithDefaults instantiates a new ProductAddSizeChart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddSizeChartWithDefaults() *ProductAddSizeChart {
	this := ProductAddSizeChart{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductAddSizeChart) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddSizeChart) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductAddSizeChart) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductAddSizeChart) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ProductAddSizeChart) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddSizeChart) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ProductAddSizeChart) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ProductAddSizeChart) SetUrl(v string) {
	o.Url = &v
}

func (o ProductAddSizeChart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddSizeChart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableProductAddSizeChart struct {
	value *ProductAddSizeChart
	isSet bool
}

func (v NullableProductAddSizeChart) Get() *ProductAddSizeChart {
	return v.value
}

func (v *NullableProductAddSizeChart) Set(val *ProductAddSizeChart) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddSizeChart) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddSizeChart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddSizeChart(val *ProductAddSizeChart) *NullableProductAddSizeChart {
	return &NullableProductAddSizeChart{value: val, isSet: true}
}

func (v NullableProductAddSizeChart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddSizeChart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


