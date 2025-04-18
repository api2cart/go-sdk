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

// checks if the ProductAddBatchPayloadInnerImagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAddBatchPayloadInnerImagesInner{}

// ProductAddBatchPayloadInnerImagesInner struct for ProductAddBatchPayloadInnerImagesInner
type ProductAddBatchPayloadInnerImagesInner struct {
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Label *string `json:"label,omitempty"`
	Name *string `json:"name,omitempty"`
	Position *int32 `json:"position,omitempty"`
}

type _ProductAddBatchPayloadInnerImagesInner ProductAddBatchPayloadInnerImagesInner

// NewProductAddBatchPayloadInnerImagesInner instantiates a new ProductAddBatchPayloadInnerImagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAddBatchPayloadInnerImagesInner(url string) *ProductAddBatchPayloadInnerImagesInner {
	this := ProductAddBatchPayloadInnerImagesInner{}
	this.Url = url
	return &this
}

// NewProductAddBatchPayloadInnerImagesInnerWithDefaults instantiates a new ProductAddBatchPayloadInnerImagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAddBatchPayloadInnerImagesInnerWithDefaults() *ProductAddBatchPayloadInnerImagesInner {
	this := ProductAddBatchPayloadInnerImagesInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProductAddBatchPayloadInnerImagesInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddBatchPayloadInnerImagesInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProductAddBatchPayloadInnerImagesInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProductAddBatchPayloadInnerImagesInner) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value
func (o *ProductAddBatchPayloadInnerImagesInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ProductAddBatchPayloadInnerImagesInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ProductAddBatchPayloadInnerImagesInner) SetUrl(v string) {
	o.Url = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductAddBatchPayloadInnerImagesInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddBatchPayloadInnerImagesInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductAddBatchPayloadInnerImagesInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ProductAddBatchPayloadInnerImagesInner) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAddBatchPayloadInnerImagesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddBatchPayloadInnerImagesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAddBatchPayloadInnerImagesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAddBatchPayloadInnerImagesInner) SetName(v string) {
	o.Name = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductAddBatchPayloadInnerImagesInner) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAddBatchPayloadInnerImagesInner) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductAddBatchPayloadInnerImagesInner) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *ProductAddBatchPayloadInnerImagesInner) SetPosition(v int32) {
	o.Position = &v
}

func (o ProductAddBatchPayloadInnerImagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAddBatchPayloadInnerImagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["url"] = o.Url
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	return toSerialize, nil
}

func (o *ProductAddBatchPayloadInnerImagesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varProductAddBatchPayloadInnerImagesInner := _ProductAddBatchPayloadInnerImagesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductAddBatchPayloadInnerImagesInner)

	if err != nil {
		return err
	}

	*o = ProductAddBatchPayloadInnerImagesInner(varProductAddBatchPayloadInnerImagesInner)

	return err
}

type NullableProductAddBatchPayloadInnerImagesInner struct {
	value *ProductAddBatchPayloadInnerImagesInner
	isSet bool
}

func (v NullableProductAddBatchPayloadInnerImagesInner) Get() *ProductAddBatchPayloadInnerImagesInner {
	return v.value
}

func (v *NullableProductAddBatchPayloadInnerImagesInner) Set(val *ProductAddBatchPayloadInnerImagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAddBatchPayloadInnerImagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAddBatchPayloadInnerImagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAddBatchPayloadInnerImagesInner(val *ProductAddBatchPayloadInnerImagesInner) *NullableProductAddBatchPayloadInnerImagesInner {
	return &NullableProductAddBatchPayloadInnerImagesInner{value: val, isSet: true}
}

func (v NullableProductAddBatchPayloadInnerImagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAddBatchPayloadInnerImagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


