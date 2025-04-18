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

// checks if the ProductAddSpecificsInnerBookingDetailsOverridesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddSpecificsInnerBookingDetailsOverridesInner{}

// ProductAddSpecificsInnerBookingDetailsOverridesInner struct for ProductAddSpecificsInnerBookingDetailsOverridesInner
type ProductAddSpecificsInnerBookingDetailsOverridesInner struct {
	Day string `json:"day"`
	Date string `json:"date" validate:"regexp=^(19|20)\\\\d\\\\d-(0[1-9]|1[0-2])-(0[1-9]|[12]\\\\d|3[01])$"`
}

type _ProductAddSpecificsInnerBookingDetailsOverridesInner ProductAddSpecificsInnerBookingDetailsOverridesInner

// NewProductAddSpecificsInnerBookingDetailsOverridesInner instantiates a new ProductAddSpecificsInnerBookingDetailsOverridesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddSpecificsInnerBookingDetailsOverridesInner(day string, date string) *ProductAddSpecificsInnerBookingDetailsOverridesInner {
	this := ProductAddSpecificsInnerBookingDetailsOverridesInner{}
	this.Day = day
	this.Date = date
	return &this
}

// NewProductAddSpecificsInnerBookingDetailsOverridesInnerWithDefaults instantiates a new ProductAddSpecificsInnerBookingDetailsOverridesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddSpecificsInnerBookingDetailsOverridesInnerWithDefaults() *ProductAddSpecificsInnerBookingDetailsOverridesInner {
	this := ProductAddSpecificsInnerBookingDetailsOverridesInner{}
	return &this
}

// GetDay returns the Day field value
func (o *ProductAddSpecificsInnerBookingDetailsOverridesInner) GetDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInnerBookingDetailsOverridesInner) GetDayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *ProductAddSpecificsInnerBookingDetailsOverridesInner) SetDay(v string) {
	o.Day = v
}

// GetDate returns the Date field value
func (o *ProductAddSpecificsInnerBookingDetailsOverridesInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInnerBookingDetailsOverridesInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *ProductAddSpecificsInnerBookingDetailsOverridesInner) SetDate(v string) {
	o.Date = v
}

func (o ProductAddSpecificsInnerBookingDetailsOverridesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddSpecificsInnerBookingDetailsOverridesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["day"] = o.Day
	toSerialize["date"] = o.Date
	return toSerialize, nil
}

func (o *ProductAddSpecificsInnerBookingDetailsOverridesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"day",
		"date",
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

	varProductAddSpecificsInnerBookingDetailsOverridesInner := _ProductAddSpecificsInnerBookingDetailsOverridesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductAddSpecificsInnerBookingDetailsOverridesInner)

	if err != nil {
		return err
	}

	*o = ProductAddSpecificsInnerBookingDetailsOverridesInner(varProductAddSpecificsInnerBookingDetailsOverridesInner)

	return err
}

type NullableProductAddSpecificsInnerBookingDetailsOverridesInner struct {
	value *ProductAddSpecificsInnerBookingDetailsOverridesInner
	isSet bool
}

func (v NullableProductAddSpecificsInnerBookingDetailsOverridesInner) Get() *ProductAddSpecificsInnerBookingDetailsOverridesInner {
	return v.value
}

func (v *NullableProductAddSpecificsInnerBookingDetailsOverridesInner) Set(val *ProductAddSpecificsInnerBookingDetailsOverridesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddSpecificsInnerBookingDetailsOverridesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddSpecificsInnerBookingDetailsOverridesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddSpecificsInnerBookingDetailsOverridesInner(val *ProductAddSpecificsInnerBookingDetailsOverridesInner) *NullableProductAddSpecificsInnerBookingDetailsOverridesInner {
	return &NullableProductAddSpecificsInnerBookingDetailsOverridesInner{value: val, isSet: true}
}

func (v NullableProductAddSpecificsInnerBookingDetailsOverridesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddSpecificsInnerBookingDetailsOverridesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


