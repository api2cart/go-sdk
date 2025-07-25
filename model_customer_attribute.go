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

// checks if the CustomerAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerAttribute{}

// CustomerAttribute struct for CustomerAttribute
type CustomerAttribute struct {
	AttributeId NullableString `json:"attribute_id,omitempty"`
	Code NullableString `json:"code,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Values []CustomerAttributeValue `json:"values,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewCustomerAttribute instantiates a new CustomerAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAttribute() *CustomerAttribute {
	this := CustomerAttribute{}
	return &this
}

// NewCustomerAttributeWithDefaults instantiates a new CustomerAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAttributeWithDefaults() *CustomerAttribute {
	this := CustomerAttribute{}
	return &this
}

// GetAttributeId returns the AttributeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerAttribute) GetAttributeId() string {
	if o == nil || IsNil(o.AttributeId.Get()) {
		var ret string
		return ret
	}
	return *o.AttributeId.Get()
}

// GetAttributeIdOk returns a tuple with the AttributeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerAttribute) GetAttributeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeId.Get(), o.AttributeId.IsSet()
}

// HasAttributeId returns a boolean if a field has been set.
func (o *CustomerAttribute) HasAttributeId() bool {
	if o != nil && o.AttributeId.IsSet() {
		return true
	}

	return false
}

// SetAttributeId gets a reference to the given NullableString and assigns it to the AttributeId field.
func (o *CustomerAttribute) SetAttributeId(v string) {
	o.AttributeId.Set(&v)
}
// SetAttributeIdNil sets the value for AttributeId to be an explicit nil
func (o *CustomerAttribute) SetAttributeIdNil() {
	o.AttributeId.Set(nil)
}

// UnsetAttributeId ensures that no value is present for AttributeId, not even an explicit nil
func (o *CustomerAttribute) UnsetAttributeId() {
	o.AttributeId.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerAttribute) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerAttribute) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *CustomerAttribute) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *CustomerAttribute) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *CustomerAttribute) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *CustomerAttribute) UnsetCode() {
	o.Code.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerAttribute) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CustomerAttribute) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CustomerAttribute) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CustomerAttribute) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CustomerAttribute) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerAttribute) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerAttribute) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *CustomerAttribute) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *CustomerAttribute) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *CustomerAttribute) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *CustomerAttribute) UnsetType() {
	o.Type.Unset()
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *CustomerAttribute) GetValues() []CustomerAttributeValue {
	if o == nil || IsNil(o.Values) {
		var ret []CustomerAttributeValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAttribute) GetValuesOk() ([]CustomerAttributeValue, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *CustomerAttribute) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []CustomerAttributeValue and assigns it to the Values field.
func (o *CustomerAttribute) SetValues(v []CustomerAttributeValue) {
	o.Values = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerAttribute) GetAdditionalFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerAttribute) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *CustomerAttribute) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *CustomerAttribute) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerAttribute) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerAttribute) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CustomerAttribute) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CustomerAttribute) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o CustomerAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AttributeId.IsSet() {
		toSerialize["attribute_id"] = o.AttributeId.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if o.AdditionalFields != nil {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableCustomerAttribute struct {
	value *CustomerAttribute
	isSet bool
}

func (v NullableCustomerAttribute) Get() *CustomerAttribute {
	return v.value
}

func (v *NullableCustomerAttribute) Set(val *CustomerAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAttribute(val *CustomerAttribute) *NullableCustomerAttribute {
	return &NullableCustomerAttribute{value: val, isSet: true}
}

func (v NullableCustomerAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


