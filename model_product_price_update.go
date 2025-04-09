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

// checks if the ProductPriceUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductPriceUpdate{}

// ProductPriceUpdate struct for ProductPriceUpdate
type ProductPriceUpdate struct {
	// Defines the product where the price has to be updated
	ProductId *string `json:"product_id,omitempty"`
	// Defines product's group prices
	GroupPrices []ProductPriceUpdateGroupPricesInner `json:"group_prices,omitempty"`
}

// NewProductPriceUpdate instantiates a new ProductPriceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductPriceUpdate() *ProductPriceUpdate {
	this := ProductPriceUpdate{}
	return &this
}

// NewProductPriceUpdateWithDefaults instantiates a new ProductPriceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductPriceUpdateWithDefaults() *ProductPriceUpdate {
	this := ProductPriceUpdate{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ProductPriceUpdate) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductPriceUpdate) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ProductPriceUpdate) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ProductPriceUpdate) SetProductId(v string) {
	o.ProductId = &v
}

// GetGroupPrices returns the GroupPrices field value if set, zero value otherwise.
func (o *ProductPriceUpdate) GetGroupPrices() []ProductPriceUpdateGroupPricesInner {
	if o == nil || IsNil(o.GroupPrices) {
		var ret []ProductPriceUpdateGroupPricesInner
		return ret
	}
	return o.GroupPrices
}

// GetGroupPricesOk returns a tuple with the GroupPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductPriceUpdate) GetGroupPricesOk() ([]ProductPriceUpdateGroupPricesInner, bool) {
	if o == nil || IsNil(o.GroupPrices) {
		return nil, false
	}
	return o.GroupPrices, true
}

// HasGroupPrices returns a boolean if a field has been set.
func (o *ProductPriceUpdate) HasGroupPrices() bool {
	if o != nil && !IsNil(o.GroupPrices) {
		return true
	}

	return false
}

// SetGroupPrices gets a reference to the given []ProductPriceUpdateGroupPricesInner and assigns it to the GroupPrices field.
func (o *ProductPriceUpdate) SetGroupPrices(v []ProductPriceUpdateGroupPricesInner) {
	o.GroupPrices = v
}

func (o ProductPriceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductPriceUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.GroupPrices) {
		toSerialize["group_prices"] = o.GroupPrices
	}
	return toSerialize, nil
}

type NullableProductPriceUpdate struct {
	value *ProductPriceUpdate
	isSet bool
}

func (v NullableProductPriceUpdate) Get() *ProductPriceUpdate {
	return v.value
}

func (v *NullableProductPriceUpdate) Set(val *ProductPriceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableProductPriceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableProductPriceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductPriceUpdate(val *ProductPriceUpdate) *NullableProductPriceUpdate {
	return &NullableProductPriceUpdate{value: val, isSet: true}
}

func (v NullableProductPriceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductPriceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


