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

// checks if the ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner{}

// ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner struct for ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner
type ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner struct {
	Day string `json:"day"`
	IsAvailable *bool `json:"is_available,omitempty"`
	Times []ProductAddSpecificsInnerBookingDetailsAvailabilitiesInnerTimesInner `json:"times,omitempty"`
}

type _ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner

// NewProductAddSpecificsInnerBookingDetailsAvailabilitiesInner instantiates a new ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddSpecificsInnerBookingDetailsAvailabilitiesInner(day string) *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner {
	this := ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner{}
	this.Day = day
	var isAvailable bool = true
	this.IsAvailable = &isAvailable
	return &this
}

// NewProductAddSpecificsInnerBookingDetailsAvailabilitiesInnerWithDefaults instantiates a new ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddSpecificsInnerBookingDetailsAvailabilitiesInnerWithDefaults() *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner {
	this := ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner{}
	var isAvailable bool = true
	this.IsAvailable = &isAvailable
	return &this
}

// GetDay returns the Day field value
func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) GetDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) GetDayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) SetDay(v string) {
	o.Day = v
}

// GetIsAvailable returns the IsAvailable field value if set, zero value otherwise.
func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) GetIsAvailable() bool {
	if o == nil || IsNil(o.IsAvailable) {
		var ret bool
		return ret
	}
	return *o.IsAvailable
}

// GetIsAvailableOk returns a tuple with the IsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) GetIsAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAvailable) {
		return nil, false
	}
	return o.IsAvailable, true
}

// HasIsAvailable returns a boolean if a field has been set.
func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) HasIsAvailable() bool {
	if o != nil && !IsNil(o.IsAvailable) {
		return true
	}

	return false
}

// SetIsAvailable gets a reference to the given bool and assigns it to the IsAvailable field.
func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) SetIsAvailable(v bool) {
	o.IsAvailable = &v
}

// GetTimes returns the Times field value if set, zero value otherwise.
func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) GetTimes() []ProductAddSpecificsInnerBookingDetailsAvailabilitiesInnerTimesInner {
	if o == nil || IsNil(o.Times) {
		var ret []ProductAddSpecificsInnerBookingDetailsAvailabilitiesInnerTimesInner
		return ret
	}
	return o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) GetTimesOk() ([]ProductAddSpecificsInnerBookingDetailsAvailabilitiesInnerTimesInner, bool) {
	if o == nil || IsNil(o.Times) {
		return nil, false
	}
	return o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) HasTimes() bool {
	if o != nil && !IsNil(o.Times) {
		return true
	}

	return false
}

// SetTimes gets a reference to the given []ProductAddSpecificsInnerBookingDetailsAvailabilitiesInnerTimesInner and assigns it to the Times field.
func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) SetTimes(v []ProductAddSpecificsInnerBookingDetailsAvailabilitiesInnerTimesInner) {
	o.Times = v
}

func (o ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["day"] = o.Day
	if !IsNil(o.IsAvailable) {
		toSerialize["is_available"] = o.IsAvailable
	}
	if !IsNil(o.Times) {
		toSerialize["times"] = o.Times
	}
	return toSerialize, nil
}

func (o *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"day",
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

	varProductAddSpecificsInnerBookingDetailsAvailabilitiesInner := _ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductAddSpecificsInnerBookingDetailsAvailabilitiesInner)

	if err != nil {
		return err
	}

	*o = ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner(varProductAddSpecificsInnerBookingDetailsAvailabilitiesInner)

	return err
}

type NullableProductAddSpecificsInnerBookingDetailsAvailabilitiesInner struct {
	value *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner
	isSet bool
}

func (v NullableProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) Get() *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner {
	return v.value
}

func (v *NullableProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) Set(val *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddSpecificsInnerBookingDetailsAvailabilitiesInner(val *ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) *NullableProductAddSpecificsInnerBookingDetailsAvailabilitiesInner {
	return &NullableProductAddSpecificsInnerBookingDetailsAvailabilitiesInner{value: val, isSet: true}
}

func (v NullableProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddSpecificsInnerBookingDetailsAvailabilitiesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


