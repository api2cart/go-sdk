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

// checks if the OrderPreestimateShippingListOrderItemInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderPreestimateShippingListOrderItemInner{}

// OrderPreestimateShippingListOrderItemInner struct for OrderPreestimateShippingListOrderItemInner
type OrderPreestimateShippingListOrderItemInner struct {
	// Defines orders specified by order item id
	OrderItemId string `json:"order_item_id"`
	// Defines orders specified by order item model
	OrderItemModel *string `json:"order_item_model,omitempty"`
	// Defines orders specified by order item quantity
	OrderItemQuantity int32 `json:"order_item_quantity"`
	// Defines orders specified by order item weight
	OrderItemWeight *float32 `json:"order_item_weight,omitempty"`
	// Ordered product variant. Where x is order item ID
	OrderItemVariantId *string `json:"order_item_variant_id,omitempty"`
	OrderItemOption []OrderPreestimateShippingListOrderItemInnerOrderItemOptionInner `json:"order_item_option,omitempty"`
}

type _OrderPreestimateShippingListOrderItemInner OrderPreestimateShippingListOrderItemInner

// NewOrderPreestimateShippingListOrderItemInner instantiates a new OrderPreestimateShippingListOrderItemInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderPreestimateShippingListOrderItemInner(orderItemId string, orderItemQuantity int32) *OrderPreestimateShippingListOrderItemInner {
	this := OrderPreestimateShippingListOrderItemInner{}
	this.OrderItemId = orderItemId
	this.OrderItemQuantity = orderItemQuantity
	return &this
}

// NewOrderPreestimateShippingListOrderItemInnerWithDefaults instantiates a new OrderPreestimateShippingListOrderItemInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderPreestimateShippingListOrderItemInnerWithDefaults() *OrderPreestimateShippingListOrderItemInner {
	this := OrderPreestimateShippingListOrderItemInner{}
	return &this
}

// GetOrderItemId returns the OrderItemId field value
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value
// and a boolean to check if the value has been set.
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemId, true
}

// SetOrderItemId sets field value
func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemId(v string) {
	o.OrderItemId = v
}

// GetOrderItemModel returns the OrderItemModel field value if set, zero value otherwise.
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemModel() string {
	if o == nil || IsNil(o.OrderItemModel) {
		var ret string
		return ret
	}
	return *o.OrderItemModel
}

// GetOrderItemModelOk returns a tuple with the OrderItemModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemModelOk() (*string, bool) {
	if o == nil || IsNil(o.OrderItemModel) {
		return nil, false
	}
	return o.OrderItemModel, true
}

// HasOrderItemModel returns a boolean if a field has been set.
func (o *OrderPreestimateShippingListOrderItemInner) HasOrderItemModel() bool {
	if o != nil && !IsNil(o.OrderItemModel) {
		return true
	}

	return false
}

// SetOrderItemModel gets a reference to the given string and assigns it to the OrderItemModel field.
func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemModel(v string) {
	o.OrderItemModel = &v
}

// GetOrderItemQuantity returns the OrderItemQuantity field value
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderItemQuantity
}

// GetOrderItemQuantityOk returns a tuple with the OrderItemQuantity field value
// and a boolean to check if the value has been set.
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemQuantity, true
}

// SetOrderItemQuantity sets field value
func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemQuantity(v int32) {
	o.OrderItemQuantity = v
}

// GetOrderItemWeight returns the OrderItemWeight field value if set, zero value otherwise.
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemWeight() float32 {
	if o == nil || IsNil(o.OrderItemWeight) {
		var ret float32
		return ret
	}
	return *o.OrderItemWeight
}

// GetOrderItemWeightOk returns a tuple with the OrderItemWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderItemWeight) {
		return nil, false
	}
	return o.OrderItemWeight, true
}

// HasOrderItemWeight returns a boolean if a field has been set.
func (o *OrderPreestimateShippingListOrderItemInner) HasOrderItemWeight() bool {
	if o != nil && !IsNil(o.OrderItemWeight) {
		return true
	}

	return false
}

// SetOrderItemWeight gets a reference to the given float32 and assigns it to the OrderItemWeight field.
func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemWeight(v float32) {
	o.OrderItemWeight = &v
}

// GetOrderItemVariantId returns the OrderItemVariantId field value if set, zero value otherwise.
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemVariantId() string {
	if o == nil || IsNil(o.OrderItemVariantId) {
		var ret string
		return ret
	}
	return *o.OrderItemVariantId
}

// GetOrderItemVariantIdOk returns a tuple with the OrderItemVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemVariantIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderItemVariantId) {
		return nil, false
	}
	return o.OrderItemVariantId, true
}

// HasOrderItemVariantId returns a boolean if a field has been set.
func (o *OrderPreestimateShippingListOrderItemInner) HasOrderItemVariantId() bool {
	if o != nil && !IsNil(o.OrderItemVariantId) {
		return true
	}

	return false
}

// SetOrderItemVariantId gets a reference to the given string and assigns it to the OrderItemVariantId field.
func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemVariantId(v string) {
	o.OrderItemVariantId = &v
}

// GetOrderItemOption returns the OrderItemOption field value if set, zero value otherwise.
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemOption() []OrderPreestimateShippingListOrderItemInnerOrderItemOptionInner {
	if o == nil || IsNil(o.OrderItemOption) {
		var ret []OrderPreestimateShippingListOrderItemInnerOrderItemOptionInner
		return ret
	}
	return o.OrderItemOption
}

// GetOrderItemOptionOk returns a tuple with the OrderItemOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPreestimateShippingListOrderItemInner) GetOrderItemOptionOk() ([]OrderPreestimateShippingListOrderItemInnerOrderItemOptionInner, bool) {
	if o == nil || IsNil(o.OrderItemOption) {
		return nil, false
	}
	return o.OrderItemOption, true
}

// HasOrderItemOption returns a boolean if a field has been set.
func (o *OrderPreestimateShippingListOrderItemInner) HasOrderItemOption() bool {
	if o != nil && !IsNil(o.OrderItemOption) {
		return true
	}

	return false
}

// SetOrderItemOption gets a reference to the given []OrderPreestimateShippingListOrderItemInnerOrderItemOptionInner and assigns it to the OrderItemOption field.
func (o *OrderPreestimateShippingListOrderItemInner) SetOrderItemOption(v []OrderPreestimateShippingListOrderItemInnerOrderItemOptionInner) {
	o.OrderItemOption = v
}

func (o OrderPreestimateShippingListOrderItemInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderPreestimateShippingListOrderItemInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order_item_id"] = o.OrderItemId
	if !IsNil(o.OrderItemModel) {
		toSerialize["order_item_model"] = o.OrderItemModel
	}
	toSerialize["order_item_quantity"] = o.OrderItemQuantity
	if !IsNil(o.OrderItemWeight) {
		toSerialize["order_item_weight"] = o.OrderItemWeight
	}
	if !IsNil(o.OrderItemVariantId) {
		toSerialize["order_item_variant_id"] = o.OrderItemVariantId
	}
	if !IsNil(o.OrderItemOption) {
		toSerialize["order_item_option"] = o.OrderItemOption
	}
	return toSerialize, nil
}

func (o *OrderPreestimateShippingListOrderItemInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"order_item_id",
		"order_item_quantity",
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

	varOrderPreestimateShippingListOrderItemInner := _OrderPreestimateShippingListOrderItemInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderPreestimateShippingListOrderItemInner)

	if err != nil {
		return err
	}

	*o = OrderPreestimateShippingListOrderItemInner(varOrderPreestimateShippingListOrderItemInner)

	return err
}

type NullableOrderPreestimateShippingListOrderItemInner struct {
	value *OrderPreestimateShippingListOrderItemInner
	isSet bool
}

func (v NullableOrderPreestimateShippingListOrderItemInner) Get() *OrderPreestimateShippingListOrderItemInner {
	return v.value
}

func (v *NullableOrderPreestimateShippingListOrderItemInner) Set(val *OrderPreestimateShippingListOrderItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPreestimateShippingListOrderItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPreestimateShippingListOrderItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPreestimateShippingListOrderItemInner(val *OrderPreestimateShippingListOrderItemInner) *NullableOrderPreestimateShippingListOrderItemInner {
	return &NullableOrderPreestimateShippingListOrderItemInner{value: val, isSet: true}
}

func (v NullableOrderPreestimateShippingListOrderItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderPreestimateShippingListOrderItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


