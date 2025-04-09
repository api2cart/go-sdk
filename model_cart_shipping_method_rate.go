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

// checks if the CartShippingMethodRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartShippingMethodRate{}

// CartShippingMethodRate struct for CartShippingMethodRate
type CartShippingMethodRate struct {
	MinWeight *string `json:"min_weight,omitempty"`
	MaxWeight *string `json:"max_weight,omitempty"`
	MinOrderAmount *string `json:"min_order_amount,omitempty"`
	MaxOrderAmount *string `json:"max_order_amount,omitempty"`
	MinItemsCount *string `json:"min_items_count,omitempty"`
	MaxItemsCount *string `json:"max_items_count,omitempty"`
	Price *string `json:"price,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewCartShippingMethodRate instantiates a new CartShippingMethodRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartShippingMethodRate() *CartShippingMethodRate {
	this := CartShippingMethodRate{}
	return &this
}

// NewCartShippingMethodRateWithDefaults instantiates a new CartShippingMethodRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartShippingMethodRateWithDefaults() *CartShippingMethodRate {
	this := CartShippingMethodRate{}
	return &this
}

// GetMinWeight returns the MinWeight field value if set, zero value otherwise.
func (o *CartShippingMethodRate) GetMinWeight() string {
	if o == nil || IsNil(o.MinWeight) {
		var ret string
		return ret
	}
	return *o.MinWeight
}

// GetMinWeightOk returns a tuple with the MinWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartShippingMethodRate) GetMinWeightOk() (*string, bool) {
	if o == nil || IsNil(o.MinWeight) {
		return nil, false
	}
	return o.MinWeight, true
}

// HasMinWeight returns a boolean if a field has been set.
func (o *CartShippingMethodRate) HasMinWeight() bool {
	if o != nil && !IsNil(o.MinWeight) {
		return true
	}

	return false
}

// SetMinWeight gets a reference to the given string and assigns it to the MinWeight field.
func (o *CartShippingMethodRate) SetMinWeight(v string) {
	o.MinWeight = &v
}

// GetMaxWeight returns the MaxWeight field value if set, zero value otherwise.
func (o *CartShippingMethodRate) GetMaxWeight() string {
	if o == nil || IsNil(o.MaxWeight) {
		var ret string
		return ret
	}
	return *o.MaxWeight
}

// GetMaxWeightOk returns a tuple with the MaxWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartShippingMethodRate) GetMaxWeightOk() (*string, bool) {
	if o == nil || IsNil(o.MaxWeight) {
		return nil, false
	}
	return o.MaxWeight, true
}

// HasMaxWeight returns a boolean if a field has been set.
func (o *CartShippingMethodRate) HasMaxWeight() bool {
	if o != nil && !IsNil(o.MaxWeight) {
		return true
	}

	return false
}

// SetMaxWeight gets a reference to the given string and assigns it to the MaxWeight field.
func (o *CartShippingMethodRate) SetMaxWeight(v string) {
	o.MaxWeight = &v
}

// GetMinOrderAmount returns the MinOrderAmount field value if set, zero value otherwise.
func (o *CartShippingMethodRate) GetMinOrderAmount() string {
	if o == nil || IsNil(o.MinOrderAmount) {
		var ret string
		return ret
	}
	return *o.MinOrderAmount
}

// GetMinOrderAmountOk returns a tuple with the MinOrderAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartShippingMethodRate) GetMinOrderAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MinOrderAmount) {
		return nil, false
	}
	return o.MinOrderAmount, true
}

// HasMinOrderAmount returns a boolean if a field has been set.
func (o *CartShippingMethodRate) HasMinOrderAmount() bool {
	if o != nil && !IsNil(o.MinOrderAmount) {
		return true
	}

	return false
}

// SetMinOrderAmount gets a reference to the given string and assigns it to the MinOrderAmount field.
func (o *CartShippingMethodRate) SetMinOrderAmount(v string) {
	o.MinOrderAmount = &v
}

// GetMaxOrderAmount returns the MaxOrderAmount field value if set, zero value otherwise.
func (o *CartShippingMethodRate) GetMaxOrderAmount() string {
	if o == nil || IsNil(o.MaxOrderAmount) {
		var ret string
		return ret
	}
	return *o.MaxOrderAmount
}

// GetMaxOrderAmountOk returns a tuple with the MaxOrderAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartShippingMethodRate) GetMaxOrderAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MaxOrderAmount) {
		return nil, false
	}
	return o.MaxOrderAmount, true
}

// HasMaxOrderAmount returns a boolean if a field has been set.
func (o *CartShippingMethodRate) HasMaxOrderAmount() bool {
	if o != nil && !IsNil(o.MaxOrderAmount) {
		return true
	}

	return false
}

// SetMaxOrderAmount gets a reference to the given string and assigns it to the MaxOrderAmount field.
func (o *CartShippingMethodRate) SetMaxOrderAmount(v string) {
	o.MaxOrderAmount = &v
}

// GetMinItemsCount returns the MinItemsCount field value if set, zero value otherwise.
func (o *CartShippingMethodRate) GetMinItemsCount() string {
	if o == nil || IsNil(o.MinItemsCount) {
		var ret string
		return ret
	}
	return *o.MinItemsCount
}

// GetMinItemsCountOk returns a tuple with the MinItemsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartShippingMethodRate) GetMinItemsCountOk() (*string, bool) {
	if o == nil || IsNil(o.MinItemsCount) {
		return nil, false
	}
	return o.MinItemsCount, true
}

// HasMinItemsCount returns a boolean if a field has been set.
func (o *CartShippingMethodRate) HasMinItemsCount() bool {
	if o != nil && !IsNil(o.MinItemsCount) {
		return true
	}

	return false
}

// SetMinItemsCount gets a reference to the given string and assigns it to the MinItemsCount field.
func (o *CartShippingMethodRate) SetMinItemsCount(v string) {
	o.MinItemsCount = &v
}

// GetMaxItemsCount returns the MaxItemsCount field value if set, zero value otherwise.
func (o *CartShippingMethodRate) GetMaxItemsCount() string {
	if o == nil || IsNil(o.MaxItemsCount) {
		var ret string
		return ret
	}
	return *o.MaxItemsCount
}

// GetMaxItemsCountOk returns a tuple with the MaxItemsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartShippingMethodRate) GetMaxItemsCountOk() (*string, bool) {
	if o == nil || IsNil(o.MaxItemsCount) {
		return nil, false
	}
	return o.MaxItemsCount, true
}

// HasMaxItemsCount returns a boolean if a field has been set.
func (o *CartShippingMethodRate) HasMaxItemsCount() bool {
	if o != nil && !IsNil(o.MaxItemsCount) {
		return true
	}

	return false
}

// SetMaxItemsCount gets a reference to the given string and assigns it to the MaxItemsCount field.
func (o *CartShippingMethodRate) SetMaxItemsCount(v string) {
	o.MaxItemsCount = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CartShippingMethodRate) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartShippingMethodRate) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CartShippingMethodRate) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *CartShippingMethodRate) SetPrice(v string) {
	o.Price = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *CartShippingMethodRate) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartShippingMethodRate) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *CartShippingMethodRate) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *CartShippingMethodRate) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CartShippingMethodRate) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartShippingMethodRate) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CartShippingMethodRate) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CartShippingMethodRate) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o CartShippingMethodRate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartShippingMethodRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinWeight) {
		toSerialize["min_weight"] = o.MinWeight
	}
	if !IsNil(o.MaxWeight) {
		toSerialize["max_weight"] = o.MaxWeight
	}
	if !IsNil(o.MinOrderAmount) {
		toSerialize["min_order_amount"] = o.MinOrderAmount
	}
	if !IsNil(o.MaxOrderAmount) {
		toSerialize["max_order_amount"] = o.MaxOrderAmount
	}
	if !IsNil(o.MinItemsCount) {
		toSerialize["min_items_count"] = o.MinItemsCount
	}
	if !IsNil(o.MaxItemsCount) {
		toSerialize["max_items_count"] = o.MaxItemsCount
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableCartShippingMethodRate struct {
	value *CartShippingMethodRate
	isSet bool
}

func (v NullableCartShippingMethodRate) Get() *CartShippingMethodRate {
	return v.value
}

func (v *NullableCartShippingMethodRate) Set(val *CartShippingMethodRate) {
	v.value = val
	v.isSet = true
}

func (v NullableCartShippingMethodRate) IsSet() bool {
	return v.isSet
}

func (v *NullableCartShippingMethodRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartShippingMethodRate(val *CartShippingMethodRate) *NullableCartShippingMethodRate {
	return &NullableCartShippingMethodRate{value: val, isSet: true}
}

func (v NullableCartShippingMethodRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartShippingMethodRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


