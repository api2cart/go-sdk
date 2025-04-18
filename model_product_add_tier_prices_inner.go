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

// checks if the ProductAddTierPricesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddTierPricesInner{}

// ProductAddTierPricesInner struct for ProductAddTierPricesInner
type ProductAddTierPricesInner struct {
	Quantity *float32 `json:"quantity,omitempty"`
	Price *float32 `json:"price,omitempty"`
}

// NewProductAddTierPricesInner instantiates a new ProductAddTierPricesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddTierPricesInner() *ProductAddTierPricesInner {
	this := ProductAddTierPricesInner{}
	return &this
}

// NewProductAddTierPricesInnerWithDefaults instantiates a new ProductAddTierPricesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddTierPricesInnerWithDefaults() *ProductAddTierPricesInner {
	this := ProductAddTierPricesInner{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ProductAddTierPricesInner) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddTierPricesInner) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ProductAddTierPricesInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *ProductAddTierPricesInner) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProductAddTierPricesInner) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddTierPricesInner) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProductAddTierPricesInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ProductAddTierPricesInner) SetPrice(v float32) {
	o.Price = &v
}

func (o ProductAddTierPricesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddTierPricesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableProductAddTierPricesInner struct {
	value *ProductAddTierPricesInner
	isSet bool
}

func (v NullableProductAddTierPricesInner) Get() *ProductAddTierPricesInner {
	return v.value
}

func (v *NullableProductAddTierPricesInner) Set(val *ProductAddTierPricesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddTierPricesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddTierPricesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddTierPricesInner(val *ProductAddTierPricesInner) *NullableProductAddTierPricesInner {
	return &NullableProductAddTierPricesInner{value: val, isSet: true}
}

func (v NullableProductAddTierPricesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddTierPricesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


