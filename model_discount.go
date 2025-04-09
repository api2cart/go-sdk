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

// checks if the Discount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Discount{}

// Discount struct for Discount
type Discount struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ModifierType *string `json:"modifier_type,omitempty"`
	Value *float32 `json:"value,omitempty"`
	FromTime *string `json:"from_time,omitempty"`
	ToTime *string `json:"to_time,omitempty"`
	CustomerGroupIds *string `json:"customer_group_ids,omitempty"`
	SortOrder *int32 `json:"sort_order,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewDiscount instantiates a new Discount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscount() *Discount {
	this := Discount{}
	return &this
}

// NewDiscountWithDefaults instantiates a new Discount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountWithDefaults() *Discount {
	this := Discount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Discount) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Discount) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Discount) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Discount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Discount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Discount) SetName(v string) {
	o.Name = &v
}

// GetModifierType returns the ModifierType field value if set, zero value otherwise.
func (o *Discount) GetModifierType() string {
	if o == nil || IsNil(o.ModifierType) {
		var ret string
		return ret
	}
	return *o.ModifierType
}

// GetModifierTypeOk returns a tuple with the ModifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetModifierTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ModifierType) {
		return nil, false
	}
	return o.ModifierType, true
}

// HasModifierType returns a boolean if a field has been set.
func (o *Discount) HasModifierType() bool {
	if o != nil && !IsNil(o.ModifierType) {
		return true
	}

	return false
}

// SetModifierType gets a reference to the given string and assigns it to the ModifierType field.
func (o *Discount) SetModifierType(v string) {
	o.ModifierType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Discount) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Discount) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *Discount) SetValue(v float32) {
	o.Value = &v
}

// GetFromTime returns the FromTime field value if set, zero value otherwise.
func (o *Discount) GetFromTime() string {
	if o == nil || IsNil(o.FromTime) {
		var ret string
		return ret
	}
	return *o.FromTime
}

// GetFromTimeOk returns a tuple with the FromTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetFromTimeOk() (*string, bool) {
	if o == nil || IsNil(o.FromTime) {
		return nil, false
	}
	return o.FromTime, true
}

// HasFromTime returns a boolean if a field has been set.
func (o *Discount) HasFromTime() bool {
	if o != nil && !IsNil(o.FromTime) {
		return true
	}

	return false
}

// SetFromTime gets a reference to the given string and assigns it to the FromTime field.
func (o *Discount) SetFromTime(v string) {
	o.FromTime = &v
}

// GetToTime returns the ToTime field value if set, zero value otherwise.
func (o *Discount) GetToTime() string {
	if o == nil || IsNil(o.ToTime) {
		var ret string
		return ret
	}
	return *o.ToTime
}

// GetToTimeOk returns a tuple with the ToTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetToTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ToTime) {
		return nil, false
	}
	return o.ToTime, true
}

// HasToTime returns a boolean if a field has been set.
func (o *Discount) HasToTime() bool {
	if o != nil && !IsNil(o.ToTime) {
		return true
	}

	return false
}

// SetToTime gets a reference to the given string and assigns it to the ToTime field.
func (o *Discount) SetToTime(v string) {
	o.ToTime = &v
}

// GetCustomerGroupIds returns the CustomerGroupIds field value if set, zero value otherwise.
func (o *Discount) GetCustomerGroupIds() string {
	if o == nil || IsNil(o.CustomerGroupIds) {
		var ret string
		return ret
	}
	return *o.CustomerGroupIds
}

// GetCustomerGroupIdsOk returns a tuple with the CustomerGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetCustomerGroupIdsOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerGroupIds) {
		return nil, false
	}
	return o.CustomerGroupIds, true
}

// HasCustomerGroupIds returns a boolean if a field has been set.
func (o *Discount) HasCustomerGroupIds() bool {
	if o != nil && !IsNil(o.CustomerGroupIds) {
		return true
	}

	return false
}

// SetCustomerGroupIds gets a reference to the given string and assigns it to the CustomerGroupIds field.
func (o *Discount) SetCustomerGroupIds(v string) {
	o.CustomerGroupIds = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *Discount) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *Discount) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *Discount) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *Discount) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *Discount) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *Discount) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Discount) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Discount) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Discount) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Discount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Discount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ModifierType) {
		toSerialize["modifier_type"] = o.ModifierType
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.FromTime) {
		toSerialize["from_time"] = o.FromTime
	}
	if !IsNil(o.ToTime) {
		toSerialize["to_time"] = o.ToTime
	}
	if !IsNil(o.CustomerGroupIds) {
		toSerialize["customer_group_ids"] = o.CustomerGroupIds
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sort_order"] = o.SortOrder
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableDiscount struct {
	value *Discount
	isSet bool
}

func (v NullableDiscount) Get() *Discount {
	return v.value
}

func (v *NullableDiscount) Set(val *Discount) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscount(val *Discount) *NullableDiscount {
	return &NullableDiscount{value: val, isSet: true}
}

func (v NullableDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


