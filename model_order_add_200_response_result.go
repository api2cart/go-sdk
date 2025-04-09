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

// checks if the OrderAdd200ResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderAdd200ResponseResult{}

// OrderAdd200ResponseResult struct for OrderAdd200ResponseResult
type OrderAdd200ResponseResult struct {
	OrderId *string `json:"order_id,omitempty"`
	Id NullableString `json:"id,omitempty"`
	CustomerId NullableString `json:"customer_id,omitempty"`
}

// NewOrderAdd200ResponseResult instantiates a new OrderAdd200ResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAdd200ResponseResult() *OrderAdd200ResponseResult {
	this := OrderAdd200ResponseResult{}
	return &this
}

// NewOrderAdd200ResponseResultWithDefaults instantiates a new OrderAdd200ResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAdd200ResponseResultWithDefaults() *OrderAdd200ResponseResult {
	this := OrderAdd200ResponseResult{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderAdd200ResponseResult) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAdd200ResponseResult) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderAdd200ResponseResult) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderAdd200ResponseResult) SetOrderId(v string) {
	o.OrderId = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderAdd200ResponseResult) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderAdd200ResponseResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *OrderAdd200ResponseResult) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *OrderAdd200ResponseResult) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *OrderAdd200ResponseResult) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *OrderAdd200ResponseResult) UnsetId() {
	o.Id.Unset()
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderAdd200ResponseResult) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderAdd200ResponseResult) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *OrderAdd200ResponseResult) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *OrderAdd200ResponseResult) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *OrderAdd200ResponseResult) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *OrderAdd200ResponseResult) UnsetCustomerId() {
	o.CustomerId.Unset()
}

func (o OrderAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAdd200ResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
	}
	return toSerialize, nil
}

type NullableOrderAdd200ResponseResult struct {
	value *OrderAdd200ResponseResult
	isSet bool
}

func (v NullableOrderAdd200ResponseResult) Get() *OrderAdd200ResponseResult {
	return v.value
}

func (v *NullableOrderAdd200ResponseResult) Set(val *OrderAdd200ResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAdd200ResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAdd200ResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAdd200ResponseResult(val *OrderAdd200ResponseResult) *NullableOrderAdd200ResponseResult {
	return &NullableOrderAdd200ResponseResult{value: val, isSet: true}
}

func (v NullableOrderAdd200ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAdd200ResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


