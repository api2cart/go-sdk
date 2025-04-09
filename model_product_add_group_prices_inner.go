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

// checks if the ProductAddGroupPricesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddGroupPricesInner{}

// ProductAddGroupPricesInner struct for ProductAddGroupPricesInner
type ProductAddGroupPricesInner struct {
	GroupId *string `json:"group_id,omitempty"`
	Price *float32 `json:"price,omitempty"`
}

// NewProductAddGroupPricesInner instantiates a new ProductAddGroupPricesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddGroupPricesInner() *ProductAddGroupPricesInner {
	this := ProductAddGroupPricesInner{}
	return &this
}

// NewProductAddGroupPricesInnerWithDefaults instantiates a new ProductAddGroupPricesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddGroupPricesInnerWithDefaults() *ProductAddGroupPricesInner {
	this := ProductAddGroupPricesInner{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ProductAddGroupPricesInner) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddGroupPricesInner) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ProductAddGroupPricesInner) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ProductAddGroupPricesInner) SetGroupId(v string) {
	o.GroupId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProductAddGroupPricesInner) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddGroupPricesInner) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProductAddGroupPricesInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ProductAddGroupPricesInner) SetPrice(v float32) {
	o.Price = &v
}

func (o ProductAddGroupPricesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddGroupPricesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableProductAddGroupPricesInner struct {
	value *ProductAddGroupPricesInner
	isSet bool
}

func (v NullableProductAddGroupPricesInner) Get() *ProductAddGroupPricesInner {
	return v.value
}

func (v *NullableProductAddGroupPricesInner) Set(val *ProductAddGroupPricesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddGroupPricesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddGroupPricesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddGroupPricesInner(val *ProductAddGroupPricesInner) *NullableProductAddGroupPricesInner {
	return &NullableProductAddGroupPricesInner{value: val, isSet: true}
}

func (v NullableProductAddGroupPricesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddGroupPricesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


