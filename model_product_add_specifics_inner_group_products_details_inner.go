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

// checks if the ProductAddSpecificsInnerGroupProductsDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddSpecificsInnerGroupProductsDetailsInner{}

// ProductAddSpecificsInnerGroupProductsDetailsInner struct for ProductAddSpecificsInnerGroupProductsDetailsInner
type ProductAddSpecificsInnerGroupProductsDetailsInner struct {
	Id string `json:"id"`
	Quantity int32 `json:"quantity"`
}

type _ProductAddSpecificsInnerGroupProductsDetailsInner ProductAddSpecificsInnerGroupProductsDetailsInner

// NewProductAddSpecificsInnerGroupProductsDetailsInner instantiates a new ProductAddSpecificsInnerGroupProductsDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddSpecificsInnerGroupProductsDetailsInner(id string, quantity int32) *ProductAddSpecificsInnerGroupProductsDetailsInner {
	this := ProductAddSpecificsInnerGroupProductsDetailsInner{}
	this.Id = id
	this.Quantity = quantity
	return &this
}

// NewProductAddSpecificsInnerGroupProductsDetailsInnerWithDefaults instantiates a new ProductAddSpecificsInnerGroupProductsDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddSpecificsInnerGroupProductsDetailsInnerWithDefaults() *ProductAddSpecificsInnerGroupProductsDetailsInner {
	this := ProductAddSpecificsInnerGroupProductsDetailsInner{}
	return &this
}

// GetId returns the Id field value
func (o *ProductAddSpecificsInnerGroupProductsDetailsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInnerGroupProductsDetailsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProductAddSpecificsInnerGroupProductsDetailsInner) SetId(v string) {
	o.Id = v
}

// GetQuantity returns the Quantity field value
func (o *ProductAddSpecificsInnerGroupProductsDetailsInner) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInnerGroupProductsDetailsInner) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ProductAddSpecificsInnerGroupProductsDetailsInner) SetQuantity(v int32) {
	o.Quantity = v
}

func (o ProductAddSpecificsInnerGroupProductsDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddSpecificsInnerGroupProductsDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

func (o *ProductAddSpecificsInnerGroupProductsDetailsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"quantity",
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

	varProductAddSpecificsInnerGroupProductsDetailsInner := _ProductAddSpecificsInnerGroupProductsDetailsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductAddSpecificsInnerGroupProductsDetailsInner)

	if err != nil {
		return err
	}

	*o = ProductAddSpecificsInnerGroupProductsDetailsInner(varProductAddSpecificsInnerGroupProductsDetailsInner)

	return err
}

type NullableProductAddSpecificsInnerGroupProductsDetailsInner struct {
	value *ProductAddSpecificsInnerGroupProductsDetailsInner
	isSet bool
}

func (v NullableProductAddSpecificsInnerGroupProductsDetailsInner) Get() *ProductAddSpecificsInnerGroupProductsDetailsInner {
	return v.value
}

func (v *NullableProductAddSpecificsInnerGroupProductsDetailsInner) Set(val *ProductAddSpecificsInnerGroupProductsDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddSpecificsInnerGroupProductsDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddSpecificsInnerGroupProductsDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddSpecificsInnerGroupProductsDetailsInner(val *ProductAddSpecificsInnerGroupProductsDetailsInner) *NullableProductAddSpecificsInnerGroupProductsDetailsInner {
	return &NullableProductAddSpecificsInnerGroupProductsDetailsInner{value: val, isSet: true}
}

func (v NullableProductAddSpecificsInnerGroupProductsDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddSpecificsInnerGroupProductsDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


