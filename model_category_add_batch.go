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

// checks if the CategoryAddBatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryAddBatch{}

// CategoryAddBatch struct for CategoryAddBatch
type CategoryAddBatch struct {
	// Contains an array of categories objects. The list of properties may vary depending on the specific platform.
	Payload []CategoryAddBatchPayloadInner `json:"payload"`
}

type _CategoryAddBatch CategoryAddBatch

// NewCategoryAddBatch instantiates a new CategoryAddBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryAddBatch(payload []CategoryAddBatchPayloadInner) *CategoryAddBatch {
	this := CategoryAddBatch{}
	this.Payload = payload
	return &this
}

// NewCategoryAddBatchWithDefaults instantiates a new CategoryAddBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryAddBatchWithDefaults() *CategoryAddBatch {
	this := CategoryAddBatch{}
	return &this
}

// GetPayload returns the Payload field value
func (o *CategoryAddBatch) GetPayload() []CategoryAddBatchPayloadInner {
	if o == nil {
		var ret []CategoryAddBatchPayloadInner
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *CategoryAddBatch) GetPayloadOk() ([]CategoryAddBatchPayloadInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *CategoryAddBatch) SetPayload(v []CategoryAddBatchPayloadInner) {
	o.Payload = v
}

func (o CategoryAddBatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryAddBatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

func (o *CategoryAddBatch) UnmarshalJSON(data []byte) (err error) {
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

	varCategoryAddBatch := _CategoryAddBatch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategoryAddBatch)

	if err != nil {
		return err
	}

	*o = CategoryAddBatch(varCategoryAddBatch)

	return err
}

type NullableCategoryAddBatch struct {
	value *CategoryAddBatch
	isSet bool
}

func (v NullableCategoryAddBatch) Get() *CategoryAddBatch {
	return v.value
}

func (v *NullableCategoryAddBatch) Set(val *CategoryAddBatch) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryAddBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryAddBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryAddBatch(val *CategoryAddBatch) *NullableCategoryAddBatch {
	return &NullableCategoryAddBatch{value: val, isSet: true}
}

func (v NullableCategoryAddBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryAddBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


