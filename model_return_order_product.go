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

// checks if the ReturnOrderProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnOrderProduct{}

// ReturnOrderProduct struct for ReturnOrderProduct
type ReturnOrderProduct struct {
	ProductId *string `json:"product_id,omitempty"`
	OrderProductId *string `json:"order_product_id,omitempty"`
	Sku *string `json:"sku,omitempty"`
	Name *string `json:"name,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	Reason *ReturnReason `json:"reason,omitempty"`
	Action *ReturnAction `json:"action,omitempty"`
	Condition *string `json:"condition,omitempty"`
	CustomerComment *string `json:"customer_comment,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewReturnOrderProduct instantiates a new ReturnOrderProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnOrderProduct() *ReturnOrderProduct {
	this := ReturnOrderProduct{}
	return &this
}

// NewReturnOrderProductWithDefaults instantiates a new ReturnOrderProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnOrderProductWithDefaults() *ReturnOrderProduct {
	this := ReturnOrderProduct{}
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ReturnOrderProduct) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrderProduct) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ReturnOrderProduct) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ReturnOrderProduct) SetProductId(v string) {
	o.ProductId = &v
}

// GetOrderProductId returns the OrderProductId field value if set, zero value otherwise.
func (o *ReturnOrderProduct) GetOrderProductId() string {
	if o == nil || IsNil(o.OrderProductId) {
		var ret string
		return ret
	}
	return *o.OrderProductId
}

// GetOrderProductIdOk returns a tuple with the OrderProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrderProduct) GetOrderProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderProductId) {
		return nil, false
	}
	return o.OrderProductId, true
}

// HasOrderProductId returns a boolean if a field has been set.
func (o *ReturnOrderProduct) HasOrderProductId() bool {
	if o != nil && !IsNil(o.OrderProductId) {
		return true
	}

	return false
}

// SetOrderProductId gets a reference to the given string and assigns it to the OrderProductId field.
func (o *ReturnOrderProduct) SetOrderProductId(v string) {
	o.OrderProductId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ReturnOrderProduct) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrderProduct) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ReturnOrderProduct) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ReturnOrderProduct) SetSku(v string) {
	o.Sku = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReturnOrderProduct) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrderProduct) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReturnOrderProduct) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReturnOrderProduct) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ReturnOrderProduct) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrderProduct) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ReturnOrderProduct) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ReturnOrderProduct) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ReturnOrderProduct) GetReason() ReturnReason {
	if o == nil || IsNil(o.Reason) {
		var ret ReturnReason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrderProduct) GetReasonOk() (*ReturnReason, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ReturnOrderProduct) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given ReturnReason and assigns it to the Reason field.
func (o *ReturnOrderProduct) SetReason(v ReturnReason) {
	o.Reason = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ReturnOrderProduct) GetAction() ReturnAction {
	if o == nil || IsNil(o.Action) {
		var ret ReturnAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrderProduct) GetActionOk() (*ReturnAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ReturnOrderProduct) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given ReturnAction and assigns it to the Action field.
func (o *ReturnOrderProduct) SetAction(v ReturnAction) {
	o.Action = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ReturnOrderProduct) GetCondition() string {
	if o == nil || IsNil(o.Condition) {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrderProduct) GetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ReturnOrderProduct) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *ReturnOrderProduct) SetCondition(v string) {
	o.Condition = &v
}

// GetCustomerComment returns the CustomerComment field value if set, zero value otherwise.
func (o *ReturnOrderProduct) GetCustomerComment() string {
	if o == nil || IsNil(o.CustomerComment) {
		var ret string
		return ret
	}
	return *o.CustomerComment
}

// GetCustomerCommentOk returns a tuple with the CustomerComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrderProduct) GetCustomerCommentOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerComment) {
		return nil, false
	}
	return o.CustomerComment, true
}

// HasCustomerComment returns a boolean if a field has been set.
func (o *ReturnOrderProduct) HasCustomerComment() bool {
	if o != nil && !IsNil(o.CustomerComment) {
		return true
	}

	return false
}

// SetCustomerComment gets a reference to the given string and assigns it to the CustomerComment field.
func (o *ReturnOrderProduct) SetCustomerComment(v string) {
	o.CustomerComment = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *ReturnOrderProduct) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrderProduct) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *ReturnOrderProduct) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *ReturnOrderProduct) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ReturnOrderProduct) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnOrderProduct) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ReturnOrderProduct) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ReturnOrderProduct) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ReturnOrderProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnOrderProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.OrderProductId) {
		toSerialize["order_product_id"] = o.OrderProductId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.CustomerComment) {
		toSerialize["customer_comment"] = o.CustomerComment
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableReturnOrderProduct struct {
	value *ReturnOrderProduct
	isSet bool
}

func (v NullableReturnOrderProduct) Get() *ReturnOrderProduct {
	return v.value
}

func (v *NullableReturnOrderProduct) Set(val *ReturnOrderProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnOrderProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnOrderProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnOrderProduct(val *ReturnOrderProduct) *NullableReturnOrderProduct {
	return &NullableReturnOrderProduct{value: val, isSet: true}
}

func (v NullableReturnOrderProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnOrderProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


