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

// checks if the CouponAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponAction{}

// CouponAction struct for CouponAction
type CouponAction struct {
	Scope *string `json:"scope,omitempty"`
	ApplyTo *string `json:"apply_to,omitempty"`
	Amount *float32 `json:"amount,omitempty"`
	CurrencyCode *string `json:"currency_code,omitempty"`
	IncludeTax *bool `json:"include_tax,omitempty"`
	Type *string `json:"type,omitempty"`
	DiscountedQuantity *float32 `json:"discounted_quantity,omitempty"`
	DiscountQuantityStep *int32 `json:"discount_quantity_step,omitempty"`
	LogicOperator *string `json:"logic_operator,omitempty"`
	Conditions []CouponCondition `json:"conditions,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewCouponAction instantiates a new CouponAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponAction() *CouponAction {
	this := CouponAction{}
	return &this
}

// NewCouponActionWithDefaults instantiates a new CouponAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponActionWithDefaults() *CouponAction {
	this := CouponAction{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *CouponAction) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *CouponAction) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *CouponAction) SetScope(v string) {
	o.Scope = &v
}

// GetApplyTo returns the ApplyTo field value if set, zero value otherwise.
func (o *CouponAction) GetApplyTo() string {
	if o == nil || IsNil(o.ApplyTo) {
		var ret string
		return ret
	}
	return *o.ApplyTo
}

// GetApplyToOk returns a tuple with the ApplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetApplyToOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyTo) {
		return nil, false
	}
	return o.ApplyTo, true
}

// HasApplyTo returns a boolean if a field has been set.
func (o *CouponAction) HasApplyTo() bool {
	if o != nil && !IsNil(o.ApplyTo) {
		return true
	}

	return false
}

// SetApplyTo gets a reference to the given string and assigns it to the ApplyTo field.
func (o *CouponAction) SetApplyTo(v string) {
	o.ApplyTo = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CouponAction) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CouponAction) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *CouponAction) SetAmount(v float32) {
	o.Amount = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *CouponAction) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *CouponAction) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *CouponAction) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetIncludeTax returns the IncludeTax field value if set, zero value otherwise.
func (o *CouponAction) GetIncludeTax() bool {
	if o == nil || IsNil(o.IncludeTax) {
		var ret bool
		return ret
	}
	return *o.IncludeTax
}

// GetIncludeTaxOk returns a tuple with the IncludeTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetIncludeTaxOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeTax) {
		return nil, false
	}
	return o.IncludeTax, true
}

// HasIncludeTax returns a boolean if a field has been set.
func (o *CouponAction) HasIncludeTax() bool {
	if o != nil && !IsNil(o.IncludeTax) {
		return true
	}

	return false
}

// SetIncludeTax gets a reference to the given bool and assigns it to the IncludeTax field.
func (o *CouponAction) SetIncludeTax(v bool) {
	o.IncludeTax = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CouponAction) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CouponAction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CouponAction) SetType(v string) {
	o.Type = &v
}

// GetDiscountedQuantity returns the DiscountedQuantity field value if set, zero value otherwise.
func (o *CouponAction) GetDiscountedQuantity() float32 {
	if o == nil || IsNil(o.DiscountedQuantity) {
		var ret float32
		return ret
	}
	return *o.DiscountedQuantity
}

// GetDiscountedQuantityOk returns a tuple with the DiscountedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetDiscountedQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountedQuantity) {
		return nil, false
	}
	return o.DiscountedQuantity, true
}

// HasDiscountedQuantity returns a boolean if a field has been set.
func (o *CouponAction) HasDiscountedQuantity() bool {
	if o != nil && !IsNil(o.DiscountedQuantity) {
		return true
	}

	return false
}

// SetDiscountedQuantity gets a reference to the given float32 and assigns it to the DiscountedQuantity field.
func (o *CouponAction) SetDiscountedQuantity(v float32) {
	o.DiscountedQuantity = &v
}

// GetDiscountQuantityStep returns the DiscountQuantityStep field value if set, zero value otherwise.
func (o *CouponAction) GetDiscountQuantityStep() int32 {
	if o == nil || IsNil(o.DiscountQuantityStep) {
		var ret int32
		return ret
	}
	return *o.DiscountQuantityStep
}

// GetDiscountQuantityStepOk returns a tuple with the DiscountQuantityStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetDiscountQuantityStepOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountQuantityStep) {
		return nil, false
	}
	return o.DiscountQuantityStep, true
}

// HasDiscountQuantityStep returns a boolean if a field has been set.
func (o *CouponAction) HasDiscountQuantityStep() bool {
	if o != nil && !IsNil(o.DiscountQuantityStep) {
		return true
	}

	return false
}

// SetDiscountQuantityStep gets a reference to the given int32 and assigns it to the DiscountQuantityStep field.
func (o *CouponAction) SetDiscountQuantityStep(v int32) {
	o.DiscountQuantityStep = &v
}

// GetLogicOperator returns the LogicOperator field value if set, zero value otherwise.
func (o *CouponAction) GetLogicOperator() string {
	if o == nil || IsNil(o.LogicOperator) {
		var ret string
		return ret
	}
	return *o.LogicOperator
}

// GetLogicOperatorOk returns a tuple with the LogicOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetLogicOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.LogicOperator) {
		return nil, false
	}
	return o.LogicOperator, true
}

// HasLogicOperator returns a boolean if a field has been set.
func (o *CouponAction) HasLogicOperator() bool {
	if o != nil && !IsNil(o.LogicOperator) {
		return true
	}

	return false
}

// SetLogicOperator gets a reference to the given string and assigns it to the LogicOperator field.
func (o *CouponAction) SetLogicOperator(v string) {
	o.LogicOperator = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *CouponAction) GetConditions() []CouponCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []CouponCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetConditionsOk() ([]CouponCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *CouponAction) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []CouponCondition and assigns it to the Conditions field.
func (o *CouponAction) SetConditions(v []CouponCondition) {
	o.Conditions = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *CouponAction) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *CouponAction) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *CouponAction) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CouponAction) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponAction) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CouponAction) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CouponAction) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o CouponAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.ApplyTo) {
		toSerialize["apply_to"] = o.ApplyTo
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if !IsNil(o.IncludeTax) {
		toSerialize["include_tax"] = o.IncludeTax
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.DiscountedQuantity) {
		toSerialize["discounted_quantity"] = o.DiscountedQuantity
	}
	if !IsNil(o.DiscountQuantityStep) {
		toSerialize["discount_quantity_step"] = o.DiscountQuantityStep
	}
	if !IsNil(o.LogicOperator) {
		toSerialize["logic_operator"] = o.LogicOperator
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableCouponAction struct {
	value *CouponAction
	isSet bool
}

func (v NullableCouponAction) Get() *CouponAction {
	return v.value
}

func (v *NullableCouponAction) Set(val *CouponAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponAction(val *CouponAction) *NullableCouponAction {
	return &NullableCouponAction{value: val, isSet: true}
}

func (v NullableCouponAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


