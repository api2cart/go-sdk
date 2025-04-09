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

// checks if the SpecialPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpecialPrice{}

// SpecialPrice struct for SpecialPrice
type SpecialPrice struct {
	Value *float32 `json:"value,omitempty"`
	Avail *bool `json:"avail,omitempty"`
	CreatedAt *A2CDateTime `json:"created_at,omitempty"`
	ModifiedAt *A2CDateTime `json:"modified_at,omitempty"`
	ExpiredAt *A2CDateTime `json:"expired_at,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewSpecialPrice instantiates a new SpecialPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecialPrice() *SpecialPrice {
	this := SpecialPrice{}
	return &this
}

// NewSpecialPriceWithDefaults instantiates a new SpecialPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialPriceWithDefaults() *SpecialPrice {
	this := SpecialPrice{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SpecialPrice) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialPrice) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SpecialPrice) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *SpecialPrice) SetValue(v float32) {
	o.Value = &v
}

// GetAvail returns the Avail field value if set, zero value otherwise.
func (o *SpecialPrice) GetAvail() bool {
	if o == nil || IsNil(o.Avail) {
		var ret bool
		return ret
	}
	return *o.Avail
}

// GetAvailOk returns a tuple with the Avail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialPrice) GetAvailOk() (*bool, bool) {
	if o == nil || IsNil(o.Avail) {
		return nil, false
	}
	return o.Avail, true
}

// HasAvail returns a boolean if a field has been set.
func (o *SpecialPrice) HasAvail() bool {
	if o != nil && !IsNil(o.Avail) {
		return true
	}

	return false
}

// SetAvail gets a reference to the given bool and assigns it to the Avail field.
func (o *SpecialPrice) SetAvail(v bool) {
	o.Avail = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SpecialPrice) GetCreatedAt() A2CDateTime {
	if o == nil || IsNil(o.CreatedAt) {
		var ret A2CDateTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialPrice) GetCreatedAtOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SpecialPrice) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given A2CDateTime and assigns it to the CreatedAt field.
func (o *SpecialPrice) SetCreatedAt(v A2CDateTime) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SpecialPrice) GetModifiedAt() A2CDateTime {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret A2CDateTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialPrice) GetModifiedAtOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SpecialPrice) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given A2CDateTime and assigns it to the ModifiedAt field.
func (o *SpecialPrice) SetModifiedAt(v A2CDateTime) {
	o.ModifiedAt = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *SpecialPrice) GetExpiredAt() A2CDateTime {
	if o == nil || IsNil(o.ExpiredAt) {
		var ret A2CDateTime
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialPrice) GetExpiredAtOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.ExpiredAt) {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *SpecialPrice) HasExpiredAt() bool {
	if o != nil && !IsNil(o.ExpiredAt) {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given A2CDateTime and assigns it to the ExpiredAt field.
func (o *SpecialPrice) SetExpiredAt(v A2CDateTime) {
	o.ExpiredAt = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *SpecialPrice) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialPrice) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *SpecialPrice) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *SpecialPrice) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *SpecialPrice) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialPrice) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *SpecialPrice) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *SpecialPrice) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o SpecialPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpecialPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Avail) {
		toSerialize["avail"] = o.Avail
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if !IsNil(o.ExpiredAt) {
		toSerialize["expired_at"] = o.ExpiredAt
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableSpecialPrice struct {
	value *SpecialPrice
	isSet bool
}

func (v NullableSpecialPrice) Get() *SpecialPrice {
	return v.value
}

func (v *NullableSpecialPrice) Set(val *SpecialPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecialPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecialPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecialPrice(val *SpecialPrice) *NullableSpecialPrice {
	return &NullableSpecialPrice{value: val, isSet: true}
}

func (v NullableSpecialPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecialPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


