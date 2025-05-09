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

// checks if the ProductVariantPriceUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductVariantPriceUpdate{}

// ProductVariantPriceUpdate struct for ProductVariantPriceUpdate
type ProductVariantPriceUpdate struct {
	// Defines the variant where the price has to be updated
	Id *string `json:"id,omitempty"`
	// Product id
	ProductId *string `json:"product_id,omitempty"`
	// Defines variants's group prices
	GroupPrices []ProductPriceUpdateGroupPricesInner `json:"group_prices"`
}

type _ProductVariantPriceUpdate ProductVariantPriceUpdate

// NewProductVariantPriceUpdate instantiates a new ProductVariantPriceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductVariantPriceUpdate(groupPrices []ProductPriceUpdateGroupPricesInner) *ProductVariantPriceUpdate {
	this := ProductVariantPriceUpdate{}
	this.GroupPrices = groupPrices
	return &this
}

// NewProductVariantPriceUpdateWithDefaults instantiates a new ProductVariantPriceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductVariantPriceUpdateWithDefaults() *ProductVariantPriceUpdate {
	this := ProductVariantPriceUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductVariantPriceUpdate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVariantPriceUpdate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductVariantPriceUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductVariantPriceUpdate) SetId(v string) {
	o.Id = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ProductVariantPriceUpdate) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVariantPriceUpdate) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ProductVariantPriceUpdate) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ProductVariantPriceUpdate) SetProductId(v string) {
	o.ProductId = &v
}

// GetGroupPrices returns the GroupPrices field value
func (o *ProductVariantPriceUpdate) GetGroupPrices() []ProductPriceUpdateGroupPricesInner {
	if o == nil {
		var ret []ProductPriceUpdateGroupPricesInner
		return ret
	}

	return o.GroupPrices
}

// GetGroupPricesOk returns a tuple with the GroupPrices field value
// and a boolean to check if the value has been set.
func (o *ProductVariantPriceUpdate) GetGroupPricesOk() ([]ProductPriceUpdateGroupPricesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupPrices, true
}

// SetGroupPrices sets field value
func (o *ProductVariantPriceUpdate) SetGroupPrices(v []ProductPriceUpdateGroupPricesInner) {
	o.GroupPrices = v
}

func (o ProductVariantPriceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductVariantPriceUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	toSerialize["group_prices"] = o.GroupPrices
	return toSerialize, nil
}

func (o *ProductVariantPriceUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group_prices",
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

	varProductVariantPriceUpdate := _ProductVariantPriceUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductVariantPriceUpdate)

	if err != nil {
		return err
	}

	*o = ProductVariantPriceUpdate(varProductVariantPriceUpdate)

	return err
}

type NullableProductVariantPriceUpdate struct {
	value *ProductVariantPriceUpdate
	isSet bool
}

func (v NullableProductVariantPriceUpdate) Get() *ProductVariantPriceUpdate {
	return v.value
}

func (v *NullableProductVariantPriceUpdate) Set(val *ProductVariantPriceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableProductVariantPriceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableProductVariantPriceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductVariantPriceUpdate(val *ProductVariantPriceUpdate) *NullableProductVariantPriceUpdate {
	return &NullableProductVariantPriceUpdate{value: val, isSet: true}
}

func (v NullableProductVariantPriceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductVariantPriceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


