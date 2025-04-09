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

// checks if the ProductAddShippingDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddShippingDetailsInner{}

// ProductAddShippingDetailsInner struct for ProductAddShippingDetailsInner
type ProductAddShippingDetailsInner struct {
	ShippingType *string `json:"shipping_type,omitempty"`
	ShippingService *string `json:"shipping_service,omitempty"`
	ShippingCost *float32 `json:"shipping_cost,omitempty"`
}

// NewProductAddShippingDetailsInner instantiates a new ProductAddShippingDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddShippingDetailsInner() *ProductAddShippingDetailsInner {
	this := ProductAddShippingDetailsInner{}
	return &this
}

// NewProductAddShippingDetailsInnerWithDefaults instantiates a new ProductAddShippingDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddShippingDetailsInnerWithDefaults() *ProductAddShippingDetailsInner {
	this := ProductAddShippingDetailsInner{}
	return &this
}

// GetShippingType returns the ShippingType field value if set, zero value otherwise.
func (o *ProductAddShippingDetailsInner) GetShippingType() string {
	if o == nil || IsNil(o.ShippingType) {
		var ret string
		return ret
	}
	return *o.ShippingType
}

// GetShippingTypeOk returns a tuple with the ShippingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddShippingDetailsInner) GetShippingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingType) {
		return nil, false
	}
	return o.ShippingType, true
}

// HasShippingType returns a boolean if a field has been set.
func (o *ProductAddShippingDetailsInner) HasShippingType() bool {
	if o != nil && !IsNil(o.ShippingType) {
		return true
	}

	return false
}

// SetShippingType gets a reference to the given string and assigns it to the ShippingType field.
func (o *ProductAddShippingDetailsInner) SetShippingType(v string) {
	o.ShippingType = &v
}

// GetShippingService returns the ShippingService field value if set, zero value otherwise.
func (o *ProductAddShippingDetailsInner) GetShippingService() string {
	if o == nil || IsNil(o.ShippingService) {
		var ret string
		return ret
	}
	return *o.ShippingService
}

// GetShippingServiceOk returns a tuple with the ShippingService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddShippingDetailsInner) GetShippingServiceOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingService) {
		return nil, false
	}
	return o.ShippingService, true
}

// HasShippingService returns a boolean if a field has been set.
func (o *ProductAddShippingDetailsInner) HasShippingService() bool {
	if o != nil && !IsNil(o.ShippingService) {
		return true
	}

	return false
}

// SetShippingService gets a reference to the given string and assigns it to the ShippingService field.
func (o *ProductAddShippingDetailsInner) SetShippingService(v string) {
	o.ShippingService = &v
}

// GetShippingCost returns the ShippingCost field value if set, zero value otherwise.
func (o *ProductAddShippingDetailsInner) GetShippingCost() float32 {
	if o == nil || IsNil(o.ShippingCost) {
		var ret float32
		return ret
	}
	return *o.ShippingCost
}

// GetShippingCostOk returns a tuple with the ShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddShippingDetailsInner) GetShippingCostOk() (*float32, bool) {
	if o == nil || IsNil(o.ShippingCost) {
		return nil, false
	}
	return o.ShippingCost, true
}

// HasShippingCost returns a boolean if a field has been set.
func (o *ProductAddShippingDetailsInner) HasShippingCost() bool {
	if o != nil && !IsNil(o.ShippingCost) {
		return true
	}

	return false
}

// SetShippingCost gets a reference to the given float32 and assigns it to the ShippingCost field.
func (o *ProductAddShippingDetailsInner) SetShippingCost(v float32) {
	o.ShippingCost = &v
}

func (o ProductAddShippingDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddShippingDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShippingType) {
		toSerialize["shipping_type"] = o.ShippingType
	}
	if !IsNil(o.ShippingService) {
		toSerialize["shipping_service"] = o.ShippingService
	}
	if !IsNil(o.ShippingCost) {
		toSerialize["shipping_cost"] = o.ShippingCost
	}
	return toSerialize, nil
}

type NullableProductAddShippingDetailsInner struct {
	value *ProductAddShippingDetailsInner
	isSet bool
}

func (v NullableProductAddShippingDetailsInner) Get() *ProductAddShippingDetailsInner {
	return v.value
}

func (v *NullableProductAddShippingDetailsInner) Set(val *ProductAddShippingDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddShippingDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddShippingDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddShippingDetailsInner(val *ProductAddShippingDetailsInner) *NullableProductAddShippingDetailsInner {
	return &NullableProductAddShippingDetailsInner{value: val, isSet: true}
}

func (v NullableProductAddShippingDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddShippingDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


