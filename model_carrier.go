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

// checks if the Carrier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Carrier{}

// Carrier struct for Carrier
type Carrier struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Active *bool `json:"active,omitempty"`
	ShippingMethods []OrderShippingMethod `json:"shipping_methods,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewCarrier instantiates a new Carrier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrier() *Carrier {
	this := Carrier{}
	return &this
}

// NewCarrierWithDefaults instantiates a new Carrier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierWithDefaults() *Carrier {
	this := Carrier{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Carrier) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Carrier) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Carrier) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Carrier) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Carrier) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Carrier) SetName(v string) {
	o.Name = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Carrier) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Carrier) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Carrier) SetActive(v bool) {
	o.Active = &v
}

// GetShippingMethods returns the ShippingMethods field value if set, zero value otherwise.
func (o *Carrier) GetShippingMethods() []OrderShippingMethod {
	if o == nil || IsNil(o.ShippingMethods) {
		var ret []OrderShippingMethod
		return ret
	}
	return o.ShippingMethods
}

// GetShippingMethodsOk returns a tuple with the ShippingMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetShippingMethodsOk() ([]OrderShippingMethod, bool) {
	if o == nil || IsNil(o.ShippingMethods) {
		return nil, false
	}
	return o.ShippingMethods, true
}

// HasShippingMethods returns a boolean if a field has been set.
func (o *Carrier) HasShippingMethods() bool {
	if o != nil && !IsNil(o.ShippingMethods) {
		return true
	}

	return false
}

// SetShippingMethods gets a reference to the given []OrderShippingMethod and assigns it to the ShippingMethods field.
func (o *Carrier) SetShippingMethods(v []OrderShippingMethod) {
	o.ShippingMethods = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *Carrier) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *Carrier) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *Carrier) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Carrier) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Carrier) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Carrier) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Carrier) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Carrier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Carrier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ShippingMethods) {
		toSerialize["shipping_methods"] = o.ShippingMethods
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableCarrier struct {
	value *Carrier
	isSet bool
}

func (v NullableCarrier) Get() *Carrier {
	return v.value
}

func (v *NullableCarrier) Set(val *Carrier) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrier) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrier(val *Carrier) *NullableCarrier {
	return &NullableCarrier{value: val, isSet: true}
}

func (v NullableCarrier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarrier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


