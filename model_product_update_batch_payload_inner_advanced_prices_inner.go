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

// checks if the ProductUpdateBatchPayloadInnerAdvancedPricesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductUpdateBatchPayloadInnerAdvancedPricesInner{}

// ProductUpdateBatchPayloadInnerAdvancedPricesInner struct for ProductUpdateBatchPayloadInnerAdvancedPricesInner
type ProductUpdateBatchPayloadInnerAdvancedPricesInner struct {
	Value float32 `json:"value"`
	GroupId *int32 `json:"group_id,omitempty"`
	Quantity float32 `json:"quantity"`
}

type _ProductUpdateBatchPayloadInnerAdvancedPricesInner ProductUpdateBatchPayloadInnerAdvancedPricesInner

// NewProductUpdateBatchPayloadInnerAdvancedPricesInner instantiates a new ProductUpdateBatchPayloadInnerAdvancedPricesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductUpdateBatchPayloadInnerAdvancedPricesInner(value float32, quantity float32) *ProductUpdateBatchPayloadInnerAdvancedPricesInner {
	this := ProductUpdateBatchPayloadInnerAdvancedPricesInner{}
	this.Value = value
	this.Quantity = quantity
	return &this
}

// NewProductUpdateBatchPayloadInnerAdvancedPricesInnerWithDefaults instantiates a new ProductUpdateBatchPayloadInnerAdvancedPricesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductUpdateBatchPayloadInnerAdvancedPricesInnerWithDefaults() *ProductUpdateBatchPayloadInnerAdvancedPricesInner {
	this := ProductUpdateBatchPayloadInnerAdvancedPricesInner{}
	return &this
}

// GetValue returns the Value field value
func (o *ProductUpdateBatchPayloadInnerAdvancedPricesInner) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProductUpdateBatchPayloadInnerAdvancedPricesInner) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ProductUpdateBatchPayloadInnerAdvancedPricesInner) SetValue(v float32) {
	o.Value = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ProductUpdateBatchPayloadInnerAdvancedPricesInner) GetGroupId() int32 {
	if o == nil || IsNil(o.GroupId) {
		var ret int32
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductUpdateBatchPayloadInnerAdvancedPricesInner) GetGroupIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ProductUpdateBatchPayloadInnerAdvancedPricesInner) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.
func (o *ProductUpdateBatchPayloadInnerAdvancedPricesInner) SetGroupId(v int32) {
	o.GroupId = &v
}

// GetQuantity returns the Quantity field value
func (o *ProductUpdateBatchPayloadInnerAdvancedPricesInner) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ProductUpdateBatchPayloadInnerAdvancedPricesInner) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ProductUpdateBatchPayloadInnerAdvancedPricesInner) SetQuantity(v float32) {
	o.Quantity = v
}

func (o ProductUpdateBatchPayloadInnerAdvancedPricesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductUpdateBatchPayloadInnerAdvancedPricesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

func (o *ProductUpdateBatchPayloadInnerAdvancedPricesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varProductUpdateBatchPayloadInnerAdvancedPricesInner := _ProductUpdateBatchPayloadInnerAdvancedPricesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductUpdateBatchPayloadInnerAdvancedPricesInner)

	if err != nil {
		return err
	}

	*o = ProductUpdateBatchPayloadInnerAdvancedPricesInner(varProductUpdateBatchPayloadInnerAdvancedPricesInner)

	return err
}

type NullableProductUpdateBatchPayloadInnerAdvancedPricesInner struct {
	value *ProductUpdateBatchPayloadInnerAdvancedPricesInner
	isSet bool
}

func (v NullableProductUpdateBatchPayloadInnerAdvancedPricesInner) Get() *ProductUpdateBatchPayloadInnerAdvancedPricesInner {
	return v.value
}

func (v *NullableProductUpdateBatchPayloadInnerAdvancedPricesInner) Set(val *ProductUpdateBatchPayloadInnerAdvancedPricesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductUpdateBatchPayloadInnerAdvancedPricesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductUpdateBatchPayloadInnerAdvancedPricesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductUpdateBatchPayloadInnerAdvancedPricesInner(val *ProductUpdateBatchPayloadInnerAdvancedPricesInner) *NullableProductUpdateBatchPayloadInnerAdvancedPricesInner {
	return &NullableProductUpdateBatchPayloadInnerAdvancedPricesInner{value: val, isSet: true}
}

func (v NullableProductUpdateBatchPayloadInnerAdvancedPricesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductUpdateBatchPayloadInnerAdvancedPricesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


