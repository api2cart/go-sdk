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

// checks if the OrderAddOrderItemInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderAddOrderItemInner{}

// OrderAddOrderItemInner struct for OrderAddOrderItemInner
type OrderAddOrderItemInner struct {
	// Defines orders specified by order item id
	OrderItemId string `json:"order_item_id"`
	// Defines orders specified by order item name
	OrderItemName string `json:"order_item_name"`
	// Defines orders specified by order item model
	OrderItemModel *string `json:"order_item_model,omitempty"`
	// Defines orders specified by order item price
	OrderItemPrice float32 `json:"order_item_price"`
	// Defines orders specified by order item quantity
	OrderItemQuantity int32 `json:"order_item_quantity"`
	// Defines orders specified by order item weight
	OrderItemWeight *float32 `json:"order_item_weight,omitempty"`
	// Ordered product variant. Where x is order item ID
	OrderItemVariantId *string `json:"order_item_variant_id,omitempty"`
	// Percentage of tax for product order
	OrderItemTax *float32 `json:"order_item_tax,omitempty"`
	// Index of the parent grouped/bundle product
	OrderItemParent *int32 `json:"order_item_parent,omitempty"`
	// Option name of the parent grouped/bundle product
	OrderItemParentOptionName *string `json:"order_item_parent_option_name,omitempty"`
	// Indicates whether subitems of the grouped/bundle product can be refunded separately
	OrderItemAllowRefundItemsSeparately *bool `json:"order_item_allow_refund_items_separately,omitempty"`
	// Indicates whether subitems of the grouped/bundle product can be shipped separately
	OrderItemAllowShipItemsSeparately *bool `json:"order_item_allow_ship_items_separately,omitempty"`
	// Defines if item price includes tax
	OrderItemPriceIncludesTax *bool `json:"order_item_price_includes_tax,omitempty"`
	OrderItemOption []OrderAddOrderItemInnerOrderItemOptionInner `json:"order_item_option,omitempty"`
	OrderItemProperty []OrderAddOrderItemInnerOrderItemPropertyInner `json:"order_item_property,omitempty"`
}

type _OrderAddOrderItemInner OrderAddOrderItemInner

// NewOrderAddOrderItemInner instantiates a new OrderAddOrderItemInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAddOrderItemInner(orderItemId string, orderItemName string, orderItemPrice float32, orderItemQuantity int32) *OrderAddOrderItemInner {
	this := OrderAddOrderItemInner{}
	this.OrderItemId = orderItemId
	this.OrderItemName = orderItemName
	this.OrderItemPrice = orderItemPrice
	this.OrderItemQuantity = orderItemQuantity
	var orderItemTax float32 = 0
	this.OrderItemTax = &orderItemTax
	var orderItemPriceIncludesTax bool = false
	this.OrderItemPriceIncludesTax = &orderItemPriceIncludesTax
	return &this
}

// NewOrderAddOrderItemInnerWithDefaults instantiates a new OrderAddOrderItemInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAddOrderItemInnerWithDefaults() *OrderAddOrderItemInner {
	this := OrderAddOrderItemInner{}
	var orderItemTax float32 = 0
	this.OrderItemTax = &orderItemTax
	var orderItemPriceIncludesTax bool = false
	this.OrderItemPriceIncludesTax = &orderItemPriceIncludesTax
	return &this
}

// GetOrderItemId returns the OrderItemId field value
func (o *OrderAddOrderItemInner) GetOrderItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderItemId
}

// GetOrderItemIdOk returns a tuple with the OrderItemId field value
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemId, true
}

// SetOrderItemId sets field value
func (o *OrderAddOrderItemInner) SetOrderItemId(v string) {
	o.OrderItemId = v
}

// GetOrderItemName returns the OrderItemName field value
func (o *OrderAddOrderItemInner) GetOrderItemName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderItemName
}

// GetOrderItemNameOk returns a tuple with the OrderItemName field value
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemName, true
}

// SetOrderItemName sets field value
func (o *OrderAddOrderItemInner) SetOrderItemName(v string) {
	o.OrderItemName = v
}

// GetOrderItemModel returns the OrderItemModel field value if set, zero value otherwise.
func (o *OrderAddOrderItemInner) GetOrderItemModel() string {
	if o == nil || IsNil(o.OrderItemModel) {
		var ret string
		return ret
	}
	return *o.OrderItemModel
}

// GetOrderItemModelOk returns a tuple with the OrderItemModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemModelOk() (*string, bool) {
	if o == nil || IsNil(o.OrderItemModel) {
		return nil, false
	}
	return o.OrderItemModel, true
}

// HasOrderItemModel returns a boolean if a field has been set.
func (o *OrderAddOrderItemInner) HasOrderItemModel() bool {
	if o != nil && !IsNil(o.OrderItemModel) {
		return true
	}

	return false
}

// SetOrderItemModel gets a reference to the given string and assigns it to the OrderItemModel field.
func (o *OrderAddOrderItemInner) SetOrderItemModel(v string) {
	o.OrderItemModel = &v
}

// GetOrderItemPrice returns the OrderItemPrice field value
func (o *OrderAddOrderItemInner) GetOrderItemPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OrderItemPrice
}

// GetOrderItemPriceOk returns a tuple with the OrderItemPrice field value
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemPrice, true
}

// SetOrderItemPrice sets field value
func (o *OrderAddOrderItemInner) SetOrderItemPrice(v float32) {
	o.OrderItemPrice = v
}

// GetOrderItemQuantity returns the OrderItemQuantity field value
func (o *OrderAddOrderItemInner) GetOrderItemQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderItemQuantity
}

// GetOrderItemQuantityOk returns a tuple with the OrderItemQuantity field value
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderItemQuantity, true
}

// SetOrderItemQuantity sets field value
func (o *OrderAddOrderItemInner) SetOrderItemQuantity(v int32) {
	o.OrderItemQuantity = v
}

// GetOrderItemWeight returns the OrderItemWeight field value if set, zero value otherwise.
func (o *OrderAddOrderItemInner) GetOrderItemWeight() float32 {
	if o == nil || IsNil(o.OrderItemWeight) {
		var ret float32
		return ret
	}
	return *o.OrderItemWeight
}

// GetOrderItemWeightOk returns a tuple with the OrderItemWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderItemWeight) {
		return nil, false
	}
	return o.OrderItemWeight, true
}

// HasOrderItemWeight returns a boolean if a field has been set.
func (o *OrderAddOrderItemInner) HasOrderItemWeight() bool {
	if o != nil && !IsNil(o.OrderItemWeight) {
		return true
	}

	return false
}

// SetOrderItemWeight gets a reference to the given float32 and assigns it to the OrderItemWeight field.
func (o *OrderAddOrderItemInner) SetOrderItemWeight(v float32) {
	o.OrderItemWeight = &v
}

// GetOrderItemVariantId returns the OrderItemVariantId field value if set, zero value otherwise.
func (o *OrderAddOrderItemInner) GetOrderItemVariantId() string {
	if o == nil || IsNil(o.OrderItemVariantId) {
		var ret string
		return ret
	}
	return *o.OrderItemVariantId
}

// GetOrderItemVariantIdOk returns a tuple with the OrderItemVariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemVariantIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderItemVariantId) {
		return nil, false
	}
	return o.OrderItemVariantId, true
}

// HasOrderItemVariantId returns a boolean if a field has been set.
func (o *OrderAddOrderItemInner) HasOrderItemVariantId() bool {
	if o != nil && !IsNil(o.OrderItemVariantId) {
		return true
	}

	return false
}

// SetOrderItemVariantId gets a reference to the given string and assigns it to the OrderItemVariantId field.
func (o *OrderAddOrderItemInner) SetOrderItemVariantId(v string) {
	o.OrderItemVariantId = &v
}

// GetOrderItemTax returns the OrderItemTax field value if set, zero value otherwise.
func (o *OrderAddOrderItemInner) GetOrderItemTax() float32 {
	if o == nil || IsNil(o.OrderItemTax) {
		var ret float32
		return ret
	}
	return *o.OrderItemTax
}

// GetOrderItemTaxOk returns a tuple with the OrderItemTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderItemTax) {
		return nil, false
	}
	return o.OrderItemTax, true
}

// HasOrderItemTax returns a boolean if a field has been set.
func (o *OrderAddOrderItemInner) HasOrderItemTax() bool {
	if o != nil && !IsNil(o.OrderItemTax) {
		return true
	}

	return false
}

// SetOrderItemTax gets a reference to the given float32 and assigns it to the OrderItemTax field.
func (o *OrderAddOrderItemInner) SetOrderItemTax(v float32) {
	o.OrderItemTax = &v
}

// GetOrderItemParent returns the OrderItemParent field value if set, zero value otherwise.
func (o *OrderAddOrderItemInner) GetOrderItemParent() int32 {
	if o == nil || IsNil(o.OrderItemParent) {
		var ret int32
		return ret
	}
	return *o.OrderItemParent
}

// GetOrderItemParentOk returns a tuple with the OrderItemParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemParentOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderItemParent) {
		return nil, false
	}
	return o.OrderItemParent, true
}

// HasOrderItemParent returns a boolean if a field has been set.
func (o *OrderAddOrderItemInner) HasOrderItemParent() bool {
	if o != nil && !IsNil(o.OrderItemParent) {
		return true
	}

	return false
}

// SetOrderItemParent gets a reference to the given int32 and assigns it to the OrderItemParent field.
func (o *OrderAddOrderItemInner) SetOrderItemParent(v int32) {
	o.OrderItemParent = &v
}

// GetOrderItemParentOptionName returns the OrderItemParentOptionName field value if set, zero value otherwise.
func (o *OrderAddOrderItemInner) GetOrderItemParentOptionName() string {
	if o == nil || IsNil(o.OrderItemParentOptionName) {
		var ret string
		return ret
	}
	return *o.OrderItemParentOptionName
}

// GetOrderItemParentOptionNameOk returns a tuple with the OrderItemParentOptionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemParentOptionNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrderItemParentOptionName) {
		return nil, false
	}
	return o.OrderItemParentOptionName, true
}

// HasOrderItemParentOptionName returns a boolean if a field has been set.
func (o *OrderAddOrderItemInner) HasOrderItemParentOptionName() bool {
	if o != nil && !IsNil(o.OrderItemParentOptionName) {
		return true
	}

	return false
}

// SetOrderItemParentOptionName gets a reference to the given string and assigns it to the OrderItemParentOptionName field.
func (o *OrderAddOrderItemInner) SetOrderItemParentOptionName(v string) {
	o.OrderItemParentOptionName = &v
}

// GetOrderItemAllowRefundItemsSeparately returns the OrderItemAllowRefundItemsSeparately field value if set, zero value otherwise.
func (o *OrderAddOrderItemInner) GetOrderItemAllowRefundItemsSeparately() bool {
	if o == nil || IsNil(o.OrderItemAllowRefundItemsSeparately) {
		var ret bool
		return ret
	}
	return *o.OrderItemAllowRefundItemsSeparately
}

// GetOrderItemAllowRefundItemsSeparatelyOk returns a tuple with the OrderItemAllowRefundItemsSeparately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemAllowRefundItemsSeparatelyOk() (*bool, bool) {
	if o == nil || IsNil(o.OrderItemAllowRefundItemsSeparately) {
		return nil, false
	}
	return o.OrderItemAllowRefundItemsSeparately, true
}

// HasOrderItemAllowRefundItemsSeparately returns a boolean if a field has been set.
func (o *OrderAddOrderItemInner) HasOrderItemAllowRefundItemsSeparately() bool {
	if o != nil && !IsNil(o.OrderItemAllowRefundItemsSeparately) {
		return true
	}

	return false
}

// SetOrderItemAllowRefundItemsSeparately gets a reference to the given bool and assigns it to the OrderItemAllowRefundItemsSeparately field.
func (o *OrderAddOrderItemInner) SetOrderItemAllowRefundItemsSeparately(v bool) {
	o.OrderItemAllowRefundItemsSeparately = &v
}

// GetOrderItemAllowShipItemsSeparately returns the OrderItemAllowShipItemsSeparately field value if set, zero value otherwise.
func (o *OrderAddOrderItemInner) GetOrderItemAllowShipItemsSeparately() bool {
	if o == nil || IsNil(o.OrderItemAllowShipItemsSeparately) {
		var ret bool
		return ret
	}
	return *o.OrderItemAllowShipItemsSeparately
}

// GetOrderItemAllowShipItemsSeparatelyOk returns a tuple with the OrderItemAllowShipItemsSeparately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemAllowShipItemsSeparatelyOk() (*bool, bool) {
	if o == nil || IsNil(o.OrderItemAllowShipItemsSeparately) {
		return nil, false
	}
	return o.OrderItemAllowShipItemsSeparately, true
}

// HasOrderItemAllowShipItemsSeparately returns a boolean if a field has been set.
func (o *OrderAddOrderItemInner) HasOrderItemAllowShipItemsSeparately() bool {
	if o != nil && !IsNil(o.OrderItemAllowShipItemsSeparately) {
		return true
	}

	return false
}

// SetOrderItemAllowShipItemsSeparately gets a reference to the given bool and assigns it to the OrderItemAllowShipItemsSeparately field.
func (o *OrderAddOrderItemInner) SetOrderItemAllowShipItemsSeparately(v bool) {
	o.OrderItemAllowShipItemsSeparately = &v
}

// GetOrderItemPriceIncludesTax returns the OrderItemPriceIncludesTax field value if set, zero value otherwise.
func (o *OrderAddOrderItemInner) GetOrderItemPriceIncludesTax() bool {
	if o == nil || IsNil(o.OrderItemPriceIncludesTax) {
		var ret bool
		return ret
	}
	return *o.OrderItemPriceIncludesTax
}

// GetOrderItemPriceIncludesTaxOk returns a tuple with the OrderItemPriceIncludesTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemPriceIncludesTaxOk() (*bool, bool) {
	if o == nil || IsNil(o.OrderItemPriceIncludesTax) {
		return nil, false
	}
	return o.OrderItemPriceIncludesTax, true
}

// HasOrderItemPriceIncludesTax returns a boolean if a field has been set.
func (o *OrderAddOrderItemInner) HasOrderItemPriceIncludesTax() bool {
	if o != nil && !IsNil(o.OrderItemPriceIncludesTax) {
		return true
	}

	return false
}

// SetOrderItemPriceIncludesTax gets a reference to the given bool and assigns it to the OrderItemPriceIncludesTax field.
func (o *OrderAddOrderItemInner) SetOrderItemPriceIncludesTax(v bool) {
	o.OrderItemPriceIncludesTax = &v
}

// GetOrderItemOption returns the OrderItemOption field value if set, zero value otherwise.
func (o *OrderAddOrderItemInner) GetOrderItemOption() []OrderAddOrderItemInnerOrderItemOptionInner {
	if o == nil || IsNil(o.OrderItemOption) {
		var ret []OrderAddOrderItemInnerOrderItemOptionInner
		return ret
	}
	return o.OrderItemOption
}

// GetOrderItemOptionOk returns a tuple with the OrderItemOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemOptionOk() ([]OrderAddOrderItemInnerOrderItemOptionInner, bool) {
	if o == nil || IsNil(o.OrderItemOption) {
		return nil, false
	}
	return o.OrderItemOption, true
}

// HasOrderItemOption returns a boolean if a field has been set.
func (o *OrderAddOrderItemInner) HasOrderItemOption() bool {
	if o != nil && !IsNil(o.OrderItemOption) {
		return true
	}

	return false
}

// SetOrderItemOption gets a reference to the given []OrderAddOrderItemInnerOrderItemOptionInner and assigns it to the OrderItemOption field.
func (o *OrderAddOrderItemInner) SetOrderItemOption(v []OrderAddOrderItemInnerOrderItemOptionInner) {
	o.OrderItemOption = v
}

// GetOrderItemProperty returns the OrderItemProperty field value if set, zero value otherwise.
func (o *OrderAddOrderItemInner) GetOrderItemProperty() []OrderAddOrderItemInnerOrderItemPropertyInner {
	if o == nil || IsNil(o.OrderItemProperty) {
		var ret []OrderAddOrderItemInnerOrderItemPropertyInner
		return ret
	}
	return o.OrderItemProperty
}

// GetOrderItemPropertyOk returns a tuple with the OrderItemProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddOrderItemInner) GetOrderItemPropertyOk() ([]OrderAddOrderItemInnerOrderItemPropertyInner, bool) {
	if o == nil || IsNil(o.OrderItemProperty) {
		return nil, false
	}
	return o.OrderItemProperty, true
}

// HasOrderItemProperty returns a boolean if a field has been set.
func (o *OrderAddOrderItemInner) HasOrderItemProperty() bool {
	if o != nil && !IsNil(o.OrderItemProperty) {
		return true
	}

	return false
}

// SetOrderItemProperty gets a reference to the given []OrderAddOrderItemInnerOrderItemPropertyInner and assigns it to the OrderItemProperty field.
func (o *OrderAddOrderItemInner) SetOrderItemProperty(v []OrderAddOrderItemInnerOrderItemPropertyInner) {
	o.OrderItemProperty = v
}

func (o OrderAddOrderItemInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAddOrderItemInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order_item_id"] = o.OrderItemId
	toSerialize["order_item_name"] = o.OrderItemName
	if !IsNil(o.OrderItemModel) {
		toSerialize["order_item_model"] = o.OrderItemModel
	}
	toSerialize["order_item_price"] = o.OrderItemPrice
	toSerialize["order_item_quantity"] = o.OrderItemQuantity
	if !IsNil(o.OrderItemWeight) {
		toSerialize["order_item_weight"] = o.OrderItemWeight
	}
	if !IsNil(o.OrderItemVariantId) {
		toSerialize["order_item_variant_id"] = o.OrderItemVariantId
	}
	if !IsNil(o.OrderItemTax) {
		toSerialize["order_item_tax"] = o.OrderItemTax
	}
	if !IsNil(o.OrderItemParent) {
		toSerialize["order_item_parent"] = o.OrderItemParent
	}
	if !IsNil(o.OrderItemParentOptionName) {
		toSerialize["order_item_parent_option_name"] = o.OrderItemParentOptionName
	}
	if !IsNil(o.OrderItemAllowRefundItemsSeparately) {
		toSerialize["order_item_allow_refund_items_separately"] = o.OrderItemAllowRefundItemsSeparately
	}
	if !IsNil(o.OrderItemAllowShipItemsSeparately) {
		toSerialize["order_item_allow_ship_items_separately"] = o.OrderItemAllowShipItemsSeparately
	}
	if !IsNil(o.OrderItemPriceIncludesTax) {
		toSerialize["order_item_price_includes_tax"] = o.OrderItemPriceIncludesTax
	}
	if !IsNil(o.OrderItemOption) {
		toSerialize["order_item_option"] = o.OrderItemOption
	}
	if !IsNil(o.OrderItemProperty) {
		toSerialize["order_item_property"] = o.OrderItemProperty
	}
	return toSerialize, nil
}

func (o *OrderAddOrderItemInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"order_item_id",
		"order_item_name",
		"order_item_price",
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

	varOrderAddOrderItemInner := _OrderAddOrderItemInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderAddOrderItemInner)

	if err != nil {
		return err
	}

	*o = OrderAddOrderItemInner(varOrderAddOrderItemInner)

	return err
}

type NullableOrderAddOrderItemInner struct {
	value *OrderAddOrderItemInner
	isSet bool
}

func (v NullableOrderAddOrderItemInner) Get() *OrderAddOrderItemInner {
	return v.value
}

func (v *NullableOrderAddOrderItemInner) Set(val *OrderAddOrderItemInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAddOrderItemInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAddOrderItemInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAddOrderItemInner(val *OrderAddOrderItemInner) *NullableOrderAddOrderItemInner {
	return &NullableOrderAddOrderItemInner{value: val, isSet: true}
}

func (v NullableOrderAddOrderItemInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAddOrderItemInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


