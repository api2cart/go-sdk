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

// checks if the OrderRefundAddItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderRefundAddItemsInner{}

// OrderRefundAddItemsInner struct for OrderRefundAddItemsInner
type OrderRefundAddItemsInner struct {
	OrderProductId *string `json:"order_product_id,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	Price *float32 `json:"price,omitempty"`
}

// NewOrderRefundAddItemsInner instantiates a new OrderRefundAddItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderRefundAddItemsInner() *OrderRefundAddItemsInner {
	this := OrderRefundAddItemsInner{}
	return &this
}

// NewOrderRefundAddItemsInnerWithDefaults instantiates a new OrderRefundAddItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderRefundAddItemsInnerWithDefaults() *OrderRefundAddItemsInner {
	this := OrderRefundAddItemsInner{}
	return &this
}

// GetOrderProductId returns the OrderProductId field value if set, zero value otherwise.
func (o *OrderRefundAddItemsInner) GetOrderProductId() string {
	if o == nil || IsNil(o.OrderProductId) {
		var ret string
		return ret
	}
	return *o.OrderProductId
}

// GetOrderProductIdOk returns a tuple with the OrderProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefundAddItemsInner) GetOrderProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderProductId) {
		return nil, false
	}
	return o.OrderProductId, true
}

// HasOrderProductId returns a boolean if a field has been set.
func (o *OrderRefundAddItemsInner) HasOrderProductId() bool {
	if o != nil && !IsNil(o.OrderProductId) {
		return true
	}

	return false
}

// SetOrderProductId gets a reference to the given string and assigns it to the OrderProductId field.
func (o *OrderRefundAddItemsInner) SetOrderProductId(v string) {
	o.OrderProductId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderRefundAddItemsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefundAddItemsInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderRefundAddItemsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *OrderRefundAddItemsInner) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OrderRefundAddItemsInner) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderRefundAddItemsInner) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OrderRefundAddItemsInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *OrderRefundAddItemsInner) SetPrice(v float32) {
	o.Price = &v
}

func (o OrderRefundAddItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderRefundAddItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderProductId) {
		toSerialize["order_product_id"] = o.OrderProductId
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableOrderRefundAddItemsInner struct {
	value *OrderRefundAddItemsInner
	isSet bool
}

func (v NullableOrderRefundAddItemsInner) Get() *OrderRefundAddItemsInner {
	return v.value
}

func (v *NullableOrderRefundAddItemsInner) Set(val *OrderRefundAddItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRefundAddItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRefundAddItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRefundAddItemsInner(val *OrderRefundAddItemsInner) *NullableOrderRefundAddItemsInner {
	return &NullableOrderRefundAddItemsInner{value: val, isSet: true}
}

func (v NullableOrderRefundAddItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderRefundAddItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


