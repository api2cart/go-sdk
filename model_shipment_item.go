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

// checks if the ShipmentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentItem{}

// ShipmentItem struct for ShipmentItem
type ShipmentItem struct {
	OrderProductId *string `json:"order_product_id,omitempty"`
	ProductId *string `json:"product_id,omitempty"`
	VariantId *string `json:"variant_id,omitempty"`
	Model *string `json:"model,omitempty"`
	Name *string `json:"name,omitempty"`
	Price *float32 `json:"price,omitempty"`
	Quantity *float32 `json:"quantity,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewShipmentItem instantiates a new ShipmentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentItem() *ShipmentItem {
	this := ShipmentItem{}
	return &this
}

// NewShipmentItemWithDefaults instantiates a new ShipmentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentItemWithDefaults() *ShipmentItem {
	this := ShipmentItem{}
	return &this
}

// GetOrderProductId returns the OrderProductId field value if set, zero value otherwise.
func (o *ShipmentItem) GetOrderProductId() string {
	if o == nil || IsNil(o.OrderProductId) {
		var ret string
		return ret
	}
	return *o.OrderProductId
}

// GetOrderProductIdOk returns a tuple with the OrderProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentItem) GetOrderProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderProductId) {
		return nil, false
	}
	return o.OrderProductId, true
}

// HasOrderProductId returns a boolean if a field has been set.
func (o *ShipmentItem) HasOrderProductId() bool {
	if o != nil && !IsNil(o.OrderProductId) {
		return true
	}

	return false
}

// SetOrderProductId gets a reference to the given string and assigns it to the OrderProductId field.
func (o *ShipmentItem) SetOrderProductId(v string) {
	o.OrderProductId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ShipmentItem) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentItem) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ShipmentItem) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ShipmentItem) SetProductId(v string) {
	o.ProductId = &v
}

// GetVariantId returns the VariantId field value if set, zero value otherwise.
func (o *ShipmentItem) GetVariantId() string {
	if o == nil || IsNil(o.VariantId) {
		var ret string
		return ret
	}
	return *o.VariantId
}

// GetVariantIdOk returns a tuple with the VariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentItem) GetVariantIdOk() (*string, bool) {
	if o == nil || IsNil(o.VariantId) {
		return nil, false
	}
	return o.VariantId, true
}

// HasVariantId returns a boolean if a field has been set.
func (o *ShipmentItem) HasVariantId() bool {
	if o != nil && !IsNil(o.VariantId) {
		return true
	}

	return false
}

// SetVariantId gets a reference to the given string and assigns it to the VariantId field.
func (o *ShipmentItem) SetVariantId(v string) {
	o.VariantId = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ShipmentItem) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentItem) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ShipmentItem) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ShipmentItem) SetModel(v string) {
	o.Model = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShipmentItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShipmentItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShipmentItem) SetName(v string) {
	o.Name = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ShipmentItem) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentItem) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ShipmentItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ShipmentItem) SetPrice(v float32) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ShipmentItem) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentItem) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ShipmentItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *ShipmentItem) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *ShipmentItem) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentItem) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ShipmentItem) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ShipmentItem) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ShipmentItem) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentItem) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ShipmentItem) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ShipmentItem) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ShipmentItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderProductId) {
		toSerialize["order_product_id"] = o.OrderProductId
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.VariantId) {
		toSerialize["variant_id"] = o.VariantId
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableShipmentItem struct {
	value *ShipmentItem
	isSet bool
}

func (v NullableShipmentItem) Get() *ShipmentItem {
	return v.value
}

func (v *NullableShipmentItem) Set(val *ShipmentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentItem(val *ShipmentItem) *NullableShipmentItem {
	return &NullableShipmentItem{value: val, isSet: true}
}

func (v NullableShipmentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


