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

// checks if the ProductAddManufacturerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddManufacturerInfo{}

// ProductAddManufacturerInfo Manufacturer information.
type ProductAddManufacturerInfo struct {
	// Defines manufacturer`s name
	Name *string `json:"name,omitempty"`
	// Defines manufacturer`s address
	Address *string `json:"address,omitempty"`
	// Defines manufacturer`s phone
	Phone *string `json:"phone,omitempty"`
	// Defines manufacturer`s email
	Email *string `json:"email,omitempty"`
}

// NewProductAddManufacturerInfo instantiates a new ProductAddManufacturerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddManufacturerInfo() *ProductAddManufacturerInfo {
	this := ProductAddManufacturerInfo{}
	return &this
}

// NewProductAddManufacturerInfoWithDefaults instantiates a new ProductAddManufacturerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddManufacturerInfoWithDefaults() *ProductAddManufacturerInfo {
	this := ProductAddManufacturerInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAddManufacturerInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddManufacturerInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAddManufacturerInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAddManufacturerInfo) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ProductAddManufacturerInfo) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddManufacturerInfo) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ProductAddManufacturerInfo) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ProductAddManufacturerInfo) SetAddress(v string) {
	o.Address = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ProductAddManufacturerInfo) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddManufacturerInfo) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ProductAddManufacturerInfo) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ProductAddManufacturerInfo) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ProductAddManufacturerInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddManufacturerInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ProductAddManufacturerInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ProductAddManufacturerInfo) SetEmail(v string) {
	o.Email = &v
}

func (o ProductAddManufacturerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddManufacturerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

type NullableProductAddManufacturerInfo struct {
	value *ProductAddManufacturerInfo
	isSet bool
}

func (v NullableProductAddManufacturerInfo) Get() *ProductAddManufacturerInfo {
	return v.value
}

func (v *NullableProductAddManufacturerInfo) Set(val *ProductAddManufacturerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddManufacturerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddManufacturerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddManufacturerInfo(val *ProductAddManufacturerInfo) *NullableProductAddManufacturerInfo {
	return &NullableProductAddManufacturerInfo{value: val, isSet: true}
}

func (v NullableProductAddManufacturerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddManufacturerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


