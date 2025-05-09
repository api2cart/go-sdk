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

// checks if the CartList200ResponseResultSupportedCartsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartList200ResponseResultSupportedCartsInner{}

// CartList200ResponseResultSupportedCartsInner struct for CartList200ResponseResultSupportedCartsInner
type CartList200ResponseResultSupportedCartsInner struct {
	CartId *string `json:"cart_id,omitempty"`
	CartName *string `json:"cart_name,omitempty"`
	CartVersions *string `json:"cart_versions,omitempty"`
	Params []string `json:"params,omitempty"`
}

// NewCartList200ResponseResultSupportedCartsInner instantiates a new CartList200ResponseResultSupportedCartsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartList200ResponseResultSupportedCartsInner() *CartList200ResponseResultSupportedCartsInner {
	this := CartList200ResponseResultSupportedCartsInner{}
	return &this
}

// NewCartList200ResponseResultSupportedCartsInnerWithDefaults instantiates a new CartList200ResponseResultSupportedCartsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartList200ResponseResultSupportedCartsInnerWithDefaults() *CartList200ResponseResultSupportedCartsInner {
	this := CartList200ResponseResultSupportedCartsInner{}
	return &this
}

// GetCartId returns the CartId field value if set, zero value otherwise.
func (o *CartList200ResponseResultSupportedCartsInner) GetCartId() string {
	if o == nil || IsNil(o.CartId) {
		var ret string
		return ret
	}
	return *o.CartId
}

// GetCartIdOk returns a tuple with the CartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartList200ResponseResultSupportedCartsInner) GetCartIdOk() (*string, bool) {
	if o == nil || IsNil(o.CartId) {
		return nil, false
	}
	return o.CartId, true
}

// HasCartId returns a boolean if a field has been set.
func (o *CartList200ResponseResultSupportedCartsInner) HasCartId() bool {
	if o != nil && !IsNil(o.CartId) {
		return true
	}

	return false
}

// SetCartId gets a reference to the given string and assigns it to the CartId field.
func (o *CartList200ResponseResultSupportedCartsInner) SetCartId(v string) {
	o.CartId = &v
}

// GetCartName returns the CartName field value if set, zero value otherwise.
func (o *CartList200ResponseResultSupportedCartsInner) GetCartName() string {
	if o == nil || IsNil(o.CartName) {
		var ret string
		return ret
	}
	return *o.CartName
}

// GetCartNameOk returns a tuple with the CartName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartList200ResponseResultSupportedCartsInner) GetCartNameOk() (*string, bool) {
	if o == nil || IsNil(o.CartName) {
		return nil, false
	}
	return o.CartName, true
}

// HasCartName returns a boolean if a field has been set.
func (o *CartList200ResponseResultSupportedCartsInner) HasCartName() bool {
	if o != nil && !IsNil(o.CartName) {
		return true
	}

	return false
}

// SetCartName gets a reference to the given string and assigns it to the CartName field.
func (o *CartList200ResponseResultSupportedCartsInner) SetCartName(v string) {
	o.CartName = &v
}

// GetCartVersions returns the CartVersions field value if set, zero value otherwise.
func (o *CartList200ResponseResultSupportedCartsInner) GetCartVersions() string {
	if o == nil || IsNil(o.CartVersions) {
		var ret string
		return ret
	}
	return *o.CartVersions
}

// GetCartVersionsOk returns a tuple with the CartVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartList200ResponseResultSupportedCartsInner) GetCartVersionsOk() (*string, bool) {
	if o == nil || IsNil(o.CartVersions) {
		return nil, false
	}
	return o.CartVersions, true
}

// HasCartVersions returns a boolean if a field has been set.
func (o *CartList200ResponseResultSupportedCartsInner) HasCartVersions() bool {
	if o != nil && !IsNil(o.CartVersions) {
		return true
	}

	return false
}

// SetCartVersions gets a reference to the given string and assigns it to the CartVersions field.
func (o *CartList200ResponseResultSupportedCartsInner) SetCartVersions(v string) {
	o.CartVersions = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *CartList200ResponseResultSupportedCartsInner) GetParams() []string {
	if o == nil || IsNil(o.Params) {
		var ret []string
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartList200ResponseResultSupportedCartsInner) GetParamsOk() ([]string, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *CartList200ResponseResultSupportedCartsInner) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given []string and assigns it to the Params field.
func (o *CartList200ResponseResultSupportedCartsInner) SetParams(v []string) {
	o.Params = v
}

func (o CartList200ResponseResultSupportedCartsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartList200ResponseResultSupportedCartsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CartId) {
		toSerialize["cart_id"] = o.CartId
	}
	if !IsNil(o.CartName) {
		toSerialize["cart_name"] = o.CartName
	}
	if !IsNil(o.CartVersions) {
		toSerialize["cart_versions"] = o.CartVersions
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

type NullableCartList200ResponseResultSupportedCartsInner struct {
	value *CartList200ResponseResultSupportedCartsInner
	isSet bool
}

func (v NullableCartList200ResponseResultSupportedCartsInner) Get() *CartList200ResponseResultSupportedCartsInner {
	return v.value
}

func (v *NullableCartList200ResponseResultSupportedCartsInner) Set(val *CartList200ResponseResultSupportedCartsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCartList200ResponseResultSupportedCartsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCartList200ResponseResultSupportedCartsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartList200ResponseResultSupportedCartsInner(val *CartList200ResponseResultSupportedCartsInner) *NullableCartList200ResponseResultSupportedCartsInner {
	return &NullableCartList200ResponseResultSupportedCartsInner{value: val, isSet: true}
}

func (v NullableCartList200ResponseResultSupportedCartsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartList200ResponseResultSupportedCartsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


