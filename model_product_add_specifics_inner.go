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

// checks if the ProductAddSpecificsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddSpecificsInner{}

// ProductAddSpecificsInner struct for ProductAddSpecificsInner
type ProductAddSpecificsInner struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
	Values []string `json:"values,omitempty"`
	UsedForVariations *bool `json:"used_for_variations,omitempty"`
	ScaleId NullableInt32 `json:"scale_id,omitempty"`
	InputValue NullableString `json:"input_value,omitempty"`
	FoodDetails *ProductAddSpecificsInnerFoodDetails `json:"food_details,omitempty"`
	GroupProductsDetails []ProductAddSpecificsInnerGroupProductsDetailsInner `json:"group_products_details,omitempty"`
	BookingDetails *ProductAddSpecificsInnerBookingDetails `json:"booking_details,omitempty"`
}

// NewProductAddSpecificsInner instantiates a new ProductAddSpecificsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddSpecificsInner() *ProductAddSpecificsInner {
	this := ProductAddSpecificsInner{}
	var usedForVariations bool = false
	this.UsedForVariations = &usedForVariations
	return &this
}

// NewProductAddSpecificsInnerWithDefaults instantiates a new ProductAddSpecificsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddSpecificsInnerWithDefaults() *ProductAddSpecificsInner {
	this := ProductAddSpecificsInner{}
	var usedForVariations bool = false
	this.UsedForVariations = &usedForVariations
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAddSpecificsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAddSpecificsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAddSpecificsInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProductAddSpecificsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProductAddSpecificsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ProductAddSpecificsInner) SetValue(v string) {
	o.Value = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ProductAddSpecificsInner) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInner) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ProductAddSpecificsInner) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ProductAddSpecificsInner) SetValues(v []string) {
	o.Values = v
}

// GetUsedForVariations returns the UsedForVariations field value if set, zero value otherwise.
func (o *ProductAddSpecificsInner) GetUsedForVariations() bool {
	if o == nil || IsNil(o.UsedForVariations) {
		var ret bool
		return ret
	}
	return *o.UsedForVariations
}

// GetUsedForVariationsOk returns a tuple with the UsedForVariations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInner) GetUsedForVariationsOk() (*bool, bool) {
	if o == nil || IsNil(o.UsedForVariations) {
		return nil, false
	}
	return o.UsedForVariations, true
}

// HasUsedForVariations returns a boolean if a field has been set.
func (o *ProductAddSpecificsInner) HasUsedForVariations() bool {
	if o != nil && !IsNil(o.UsedForVariations) {
		return true
	}

	return false
}

// SetUsedForVariations gets a reference to the given bool and assigns it to the UsedForVariations field.
func (o *ProductAddSpecificsInner) SetUsedForVariations(v bool) {
	o.UsedForVariations = &v
}

// GetScaleId returns the ScaleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAddSpecificsInner) GetScaleId() int32 {
	if o == nil || IsNil(o.ScaleId.Get()) {
		var ret int32
		return ret
	}
	return *o.ScaleId.Get()
}

// GetScaleIdOk returns a tuple with the ScaleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAddSpecificsInner) GetScaleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScaleId.Get(), o.ScaleId.IsSet()
}

// HasScaleId returns a boolean if a field has been set.
func (o *ProductAddSpecificsInner) HasScaleId() bool {
	if o != nil && o.ScaleId.IsSet() {
		return true
	}

	return false
}

// SetScaleId gets a reference to the given NullableInt32 and assigns it to the ScaleId field.
func (o *ProductAddSpecificsInner) SetScaleId(v int32) {
	o.ScaleId.Set(&v)
}
// SetScaleIdNil sets the value for ScaleId to be an explicit nil
func (o *ProductAddSpecificsInner) SetScaleIdNil() {
	o.ScaleId.Set(nil)
}

// UnsetScaleId ensures that no value is present for ScaleId, not even an explicit nil
func (o *ProductAddSpecificsInner) UnsetScaleId() {
	o.ScaleId.Unset()
}

// GetInputValue returns the InputValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAddSpecificsInner) GetInputValue() string {
	if o == nil || IsNil(o.InputValue.Get()) {
		var ret string
		return ret
	}
	return *o.InputValue.Get()
}

// GetInputValueOk returns a tuple with the InputValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAddSpecificsInner) GetInputValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputValue.Get(), o.InputValue.IsSet()
}

// HasInputValue returns a boolean if a field has been set.
func (o *ProductAddSpecificsInner) HasInputValue() bool {
	if o != nil && o.InputValue.IsSet() {
		return true
	}

	return false
}

// SetInputValue gets a reference to the given NullableString and assigns it to the InputValue field.
func (o *ProductAddSpecificsInner) SetInputValue(v string) {
	o.InputValue.Set(&v)
}
// SetInputValueNil sets the value for InputValue to be an explicit nil
func (o *ProductAddSpecificsInner) SetInputValueNil() {
	o.InputValue.Set(nil)
}

// UnsetInputValue ensures that no value is present for InputValue, not even an explicit nil
func (o *ProductAddSpecificsInner) UnsetInputValue() {
	o.InputValue.Unset()
}

// GetFoodDetails returns the FoodDetails field value if set, zero value otherwise.
func (o *ProductAddSpecificsInner) GetFoodDetails() ProductAddSpecificsInnerFoodDetails {
	if o == nil || IsNil(o.FoodDetails) {
		var ret ProductAddSpecificsInnerFoodDetails
		return ret
	}
	return *o.FoodDetails
}

// GetFoodDetailsOk returns a tuple with the FoodDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInner) GetFoodDetailsOk() (*ProductAddSpecificsInnerFoodDetails, bool) {
	if o == nil || IsNil(o.FoodDetails) {
		return nil, false
	}
	return o.FoodDetails, true
}

// HasFoodDetails returns a boolean if a field has been set.
func (o *ProductAddSpecificsInner) HasFoodDetails() bool {
	if o != nil && !IsNil(o.FoodDetails) {
		return true
	}

	return false
}

// SetFoodDetails gets a reference to the given ProductAddSpecificsInnerFoodDetails and assigns it to the FoodDetails field.
func (o *ProductAddSpecificsInner) SetFoodDetails(v ProductAddSpecificsInnerFoodDetails) {
	o.FoodDetails = &v
}

// GetGroupProductsDetails returns the GroupProductsDetails field value if set, zero value otherwise.
func (o *ProductAddSpecificsInner) GetGroupProductsDetails() []ProductAddSpecificsInnerGroupProductsDetailsInner {
	if o == nil || IsNil(o.GroupProductsDetails) {
		var ret []ProductAddSpecificsInnerGroupProductsDetailsInner
		return ret
	}
	return o.GroupProductsDetails
}

// GetGroupProductsDetailsOk returns a tuple with the GroupProductsDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInner) GetGroupProductsDetailsOk() ([]ProductAddSpecificsInnerGroupProductsDetailsInner, bool) {
	if o == nil || IsNil(o.GroupProductsDetails) {
		return nil, false
	}
	return o.GroupProductsDetails, true
}

// HasGroupProductsDetails returns a boolean if a field has been set.
func (o *ProductAddSpecificsInner) HasGroupProductsDetails() bool {
	if o != nil && !IsNil(o.GroupProductsDetails) {
		return true
	}

	return false
}

// SetGroupProductsDetails gets a reference to the given []ProductAddSpecificsInnerGroupProductsDetailsInner and assigns it to the GroupProductsDetails field.
func (o *ProductAddSpecificsInner) SetGroupProductsDetails(v []ProductAddSpecificsInnerGroupProductsDetailsInner) {
	o.GroupProductsDetails = v
}

// GetBookingDetails returns the BookingDetails field value if set, zero value otherwise.
func (o *ProductAddSpecificsInner) GetBookingDetails() ProductAddSpecificsInnerBookingDetails {
	if o == nil || IsNil(o.BookingDetails) {
		var ret ProductAddSpecificsInnerBookingDetails
		return ret
	}
	return *o.BookingDetails
}

// GetBookingDetailsOk returns a tuple with the BookingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddSpecificsInner) GetBookingDetailsOk() (*ProductAddSpecificsInnerBookingDetails, bool) {
	if o == nil || IsNil(o.BookingDetails) {
		return nil, false
	}
	return o.BookingDetails, true
}

// HasBookingDetails returns a boolean if a field has been set.
func (o *ProductAddSpecificsInner) HasBookingDetails() bool {
	if o != nil && !IsNil(o.BookingDetails) {
		return true
	}

	return false
}

// SetBookingDetails gets a reference to the given ProductAddSpecificsInnerBookingDetails and assigns it to the BookingDetails field.
func (o *ProductAddSpecificsInner) SetBookingDetails(v ProductAddSpecificsInnerBookingDetails) {
	o.BookingDetails = &v
}

func (o ProductAddSpecificsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddSpecificsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.UsedForVariations) {
		toSerialize["used_for_variations"] = o.UsedForVariations
	}
	if o.ScaleId.IsSet() {
		toSerialize["scale_id"] = o.ScaleId.Get()
	}
	if o.InputValue.IsSet() {
		toSerialize["input_value"] = o.InputValue.Get()
	}
	if !IsNil(o.FoodDetails) {
		toSerialize["food_details"] = o.FoodDetails
	}
	if !IsNil(o.GroupProductsDetails) {
		toSerialize["group_products_details"] = o.GroupProductsDetails
	}
	if !IsNil(o.BookingDetails) {
		toSerialize["booking_details"] = o.BookingDetails
	}
	return toSerialize, nil
}

type NullableProductAddSpecificsInner struct {
	value *ProductAddSpecificsInner
	isSet bool
}

func (v NullableProductAddSpecificsInner) Get() *ProductAddSpecificsInner {
	return v.value
}

func (v *NullableProductAddSpecificsInner) Set(val *ProductAddSpecificsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddSpecificsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddSpecificsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddSpecificsInner(val *ProductAddSpecificsInner) *NullableProductAddSpecificsInner {
	return &NullableProductAddSpecificsInner{value: val, isSet: true}
}

func (v NullableProductAddSpecificsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddSpecificsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


