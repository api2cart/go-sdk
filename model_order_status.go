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

// checks if the OrderStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderStatus{}

// OrderStatus struct for OrderStatus
type OrderStatus struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	History []OrderStatusHistoryItem `json:"history,omitempty"`
	RefundInfo NullableOrderStatusRefund `json:"refund_info,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewOrderStatus instantiates a new OrderStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStatus() *OrderStatus {
	this := OrderStatus{}
	return &this
}

// NewOrderStatusWithDefaults instantiates a new OrderStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStatusWithDefaults() *OrderStatus {
	this := OrderStatus{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderStatus) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatus) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderStatus) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderStatus) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrderStatus) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatus) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrderStatus) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrderStatus) SetName(v string) {
	o.Name = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *OrderStatus) GetHistory() []OrderStatusHistoryItem {
	if o == nil || IsNil(o.History) {
		var ret []OrderStatusHistoryItem
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatus) GetHistoryOk() ([]OrderStatusHistoryItem, bool) {
	if o == nil || IsNil(o.History) {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *OrderStatus) HasHistory() bool {
	if o != nil && !IsNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []OrderStatusHistoryItem and assigns it to the History field.
func (o *OrderStatus) SetHistory(v []OrderStatusHistoryItem) {
	o.History = v
}

// GetRefundInfo returns the RefundInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderStatus) GetRefundInfo() OrderStatusRefund {
	if o == nil || IsNil(o.RefundInfo.Get()) {
		var ret OrderStatusRefund
		return ret
	}
	return *o.RefundInfo.Get()
}

// GetRefundInfoOk returns a tuple with the RefundInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderStatus) GetRefundInfoOk() (*OrderStatusRefund, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefundInfo.Get(), o.RefundInfo.IsSet()
}

// HasRefundInfo returns a boolean if a field has been set.
func (o *OrderStatus) HasRefundInfo() bool {
	if o != nil && o.RefundInfo.IsSet() {
		return true
	}

	return false
}

// SetRefundInfo gets a reference to the given NullableOrderStatusRefund and assigns it to the RefundInfo field.
func (o *OrderStatus) SetRefundInfo(v OrderStatusRefund) {
	o.RefundInfo.Set(&v)
}
// SetRefundInfoNil sets the value for RefundInfo to be an explicit nil
func (o *OrderStatus) SetRefundInfoNil() {
	o.RefundInfo.Set(nil)
}

// UnsetRefundInfo ensures that no value is present for RefundInfo, not even an explicit nil
func (o *OrderStatus) UnsetRefundInfo() {
	o.RefundInfo.Unset()
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderStatus) GetAdditionalFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderStatus) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *OrderStatus) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *OrderStatus) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderStatus) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderStatus) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *OrderStatus) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *OrderStatus) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o OrderStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.History) {
		toSerialize["history"] = o.History
	}
	if o.RefundInfo.IsSet() {
		toSerialize["refund_info"] = o.RefundInfo.Get()
	}
	if o.AdditionalFields != nil {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableOrderStatus struct {
	value *OrderStatus
	isSet bool
}

func (v NullableOrderStatus) Get() *OrderStatus {
	return v.value
}

func (v *NullableOrderStatus) Set(val *OrderStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatus(val *OrderStatus) *NullableOrderStatus {
	return &NullableOrderStatus{value: val, isSet: true}
}

func (v NullableOrderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


