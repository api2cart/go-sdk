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

// checks if the ProductDeleteBatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDeleteBatch{}

// ProductDeleteBatch struct for ProductDeleteBatch
type ProductDeleteBatch struct {
	// Contains an array of product deletion requests, each including the product ID.
	Payload []ProductDeleteBatchPayloadInner `json:"payload"`
}

type _ProductDeleteBatch ProductDeleteBatch

// NewProductDeleteBatch instantiates a new ProductDeleteBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDeleteBatch(payload []ProductDeleteBatchPayloadInner) *ProductDeleteBatch {
	this := ProductDeleteBatch{}
	this.Payload = payload
	return &this
}

// NewProductDeleteBatchWithDefaults instantiates a new ProductDeleteBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDeleteBatchWithDefaults() *ProductDeleteBatch {
	this := ProductDeleteBatch{}
	return &this
}

// GetPayload returns the Payload field value
func (o *ProductDeleteBatch) GetPayload() []ProductDeleteBatchPayloadInner {
	if o == nil {
		var ret []ProductDeleteBatchPayloadInner
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *ProductDeleteBatch) GetPayloadOk() ([]ProductDeleteBatchPayloadInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *ProductDeleteBatch) SetPayload(v []ProductDeleteBatchPayloadInner) {
	o.Payload = v
}

func (o ProductDeleteBatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDeleteBatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

func (o *ProductDeleteBatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"payload",
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

	varProductDeleteBatch := _ProductDeleteBatch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductDeleteBatch)

	if err != nil {
		return err
	}

	*o = ProductDeleteBatch(varProductDeleteBatch)

	return err
}

type NullableProductDeleteBatch struct {
	value *ProductDeleteBatch
	isSet bool
}

func (v NullableProductDeleteBatch) Get() *ProductDeleteBatch {
	return v.value
}

func (v *NullableProductDeleteBatch) Set(val *ProductDeleteBatch) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDeleteBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDeleteBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDeleteBatch(val *ProductDeleteBatch) *NullableProductDeleteBatch {
	return &NullableProductDeleteBatch{value: val, isSet: true}
}

func (v NullableProductDeleteBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDeleteBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


