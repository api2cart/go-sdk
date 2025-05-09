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

// checks if the ProductGroupItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductGroupItem{}

// ProductGroupItem struct for ProductGroupItem
type ProductGroupItem struct {
	ChildItemId *string `json:"child_item_id,omitempty"`
	ProductId *string `json:"product_id,omitempty"`
	DefaultQtyInPack *string `json:"default_qty_in_pack,omitempty"`
	IsQtyInPackFixed *bool `json:"is_qty_in_pack_fixed,omitempty"`
	Price *float32 `json:"price,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewProductGroupItem instantiates a new ProductGroupItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductGroupItem() *ProductGroupItem {
	this := ProductGroupItem{}
	return &this
}

// NewProductGroupItemWithDefaults instantiates a new ProductGroupItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductGroupItemWithDefaults() *ProductGroupItem {
	this := ProductGroupItem{}
	return &this
}

// GetChildItemId returns the ChildItemId field value if set, zero value otherwise.
func (o *ProductGroupItem) GetChildItemId() string {
	if o == nil || IsNil(o.ChildItemId) {
		var ret string
		return ret
	}
	return *o.ChildItemId
}

// GetChildItemIdOk returns a tuple with the ChildItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGroupItem) GetChildItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChildItemId) {
		return nil, false
	}
	return o.ChildItemId, true
}

// HasChildItemId returns a boolean if a field has been set.
func (o *ProductGroupItem) HasChildItemId() bool {
	if o != nil && !IsNil(o.ChildItemId) {
		return true
	}

	return false
}

// SetChildItemId gets a reference to the given string and assigns it to the ChildItemId field.
func (o *ProductGroupItem) SetChildItemId(v string) {
	o.ChildItemId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ProductGroupItem) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGroupItem) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ProductGroupItem) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ProductGroupItem) SetProductId(v string) {
	o.ProductId = &v
}

// GetDefaultQtyInPack returns the DefaultQtyInPack field value if set, zero value otherwise.
func (o *ProductGroupItem) GetDefaultQtyInPack() string {
	if o == nil || IsNil(o.DefaultQtyInPack) {
		var ret string
		return ret
	}
	return *o.DefaultQtyInPack
}

// GetDefaultQtyInPackOk returns a tuple with the DefaultQtyInPack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGroupItem) GetDefaultQtyInPackOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultQtyInPack) {
		return nil, false
	}
	return o.DefaultQtyInPack, true
}

// HasDefaultQtyInPack returns a boolean if a field has been set.
func (o *ProductGroupItem) HasDefaultQtyInPack() bool {
	if o != nil && !IsNil(o.DefaultQtyInPack) {
		return true
	}

	return false
}

// SetDefaultQtyInPack gets a reference to the given string and assigns it to the DefaultQtyInPack field.
func (o *ProductGroupItem) SetDefaultQtyInPack(v string) {
	o.DefaultQtyInPack = &v
}

// GetIsQtyInPackFixed returns the IsQtyInPackFixed field value if set, zero value otherwise.
func (o *ProductGroupItem) GetIsQtyInPackFixed() bool {
	if o == nil || IsNil(o.IsQtyInPackFixed) {
		var ret bool
		return ret
	}
	return *o.IsQtyInPackFixed
}

// GetIsQtyInPackFixedOk returns a tuple with the IsQtyInPackFixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGroupItem) GetIsQtyInPackFixedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsQtyInPackFixed) {
		return nil, false
	}
	return o.IsQtyInPackFixed, true
}

// HasIsQtyInPackFixed returns a boolean if a field has been set.
func (o *ProductGroupItem) HasIsQtyInPackFixed() bool {
	if o != nil && !IsNil(o.IsQtyInPackFixed) {
		return true
	}

	return false
}

// SetIsQtyInPackFixed gets a reference to the given bool and assigns it to the IsQtyInPackFixed field.
func (o *ProductGroupItem) SetIsQtyInPackFixed(v bool) {
	o.IsQtyInPackFixed = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProductGroupItem) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGroupItem) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProductGroupItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ProductGroupItem) SetPrice(v float32) {
	o.Price = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *ProductGroupItem) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGroupItem) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ProductGroupItem) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ProductGroupItem) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ProductGroupItem) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGroupItem) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ProductGroupItem) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ProductGroupItem) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ProductGroupItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductGroupItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChildItemId) {
		toSerialize["child_item_id"] = o.ChildItemId
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.DefaultQtyInPack) {
		toSerialize["default_qty_in_pack"] = o.DefaultQtyInPack
	}
	if !IsNil(o.IsQtyInPackFixed) {
		toSerialize["is_qty_in_pack_fixed"] = o.IsQtyInPackFixed
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

type NullableProductGroupItem struct {
	value *ProductGroupItem
	isSet bool
}

func (v NullableProductGroupItem) Get() *ProductGroupItem {
	return v.value
}

func (v *NullableProductGroupItem) Set(val *ProductGroupItem) {
	v.value = val
	v.isSet = true
}

func (v NullableProductGroupItem) IsSet() bool {
	return v.isSet
}

func (v *NullableProductGroupItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductGroupItem(val *ProductGroupItem) *NullableProductGroupItem {
	return &NullableProductGroupItem{value: val, isSet: true}
}

func (v NullableProductGroupItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductGroupItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


