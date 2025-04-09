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

// checks if the CategoryAddBatchPayloadInnerImagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryAddBatchPayloadInnerImagesInner{}

// CategoryAddBatchPayloadInnerImagesInner struct for CategoryAddBatchPayloadInnerImagesInner
type CategoryAddBatchPayloadInnerImagesInner struct {
	Url string `json:"url"`
	ImageName *string `json:"image_name,omitempty"`
	Type *string `json:"type,omitempty"`
	Label *string `json:"label,omitempty"`
}

type _CategoryAddBatchPayloadInnerImagesInner CategoryAddBatchPayloadInnerImagesInner

// NewCategoryAddBatchPayloadInnerImagesInner instantiates a new CategoryAddBatchPayloadInnerImagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryAddBatchPayloadInnerImagesInner(url string) *CategoryAddBatchPayloadInnerImagesInner {
	this := CategoryAddBatchPayloadInnerImagesInner{}
	this.Url = url
	return &this
}

// NewCategoryAddBatchPayloadInnerImagesInnerWithDefaults instantiates a new CategoryAddBatchPayloadInnerImagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryAddBatchPayloadInnerImagesInnerWithDefaults() *CategoryAddBatchPayloadInnerImagesInner {
	this := CategoryAddBatchPayloadInnerImagesInner{}
	return &this
}

// GetUrl returns the Url field value
func (o *CategoryAddBatchPayloadInnerImagesInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CategoryAddBatchPayloadInnerImagesInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CategoryAddBatchPayloadInnerImagesInner) SetUrl(v string) {
	o.Url = v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *CategoryAddBatchPayloadInnerImagesInner) GetImageName() string {
	if o == nil || IsNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryAddBatchPayloadInnerImagesInner) GetImageNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *CategoryAddBatchPayloadInnerImagesInner) HasImageName() bool {
	if o != nil && !IsNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *CategoryAddBatchPayloadInnerImagesInner) SetImageName(v string) {
	o.ImageName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CategoryAddBatchPayloadInnerImagesInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryAddBatchPayloadInnerImagesInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CategoryAddBatchPayloadInnerImagesInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CategoryAddBatchPayloadInnerImagesInner) SetType(v string) {
	o.Type = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CategoryAddBatchPayloadInnerImagesInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryAddBatchPayloadInnerImagesInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CategoryAddBatchPayloadInnerImagesInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CategoryAddBatchPayloadInnerImagesInner) SetLabel(v string) {
	o.Label = &v
}

func (o CategoryAddBatchPayloadInnerImagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryAddBatchPayloadInnerImagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if !IsNil(o.ImageName) {
		toSerialize["image_name"] = o.ImageName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

func (o *CategoryAddBatchPayloadInnerImagesInner) UnmarshalJSON(data []byte) (err error) {
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

	varCategoryAddBatchPayloadInnerImagesInner := _CategoryAddBatchPayloadInnerImagesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategoryAddBatchPayloadInnerImagesInner)

	if err != nil {
		return err
	}

	*o = CategoryAddBatchPayloadInnerImagesInner(varCategoryAddBatchPayloadInnerImagesInner)

	return err
}

type NullableCategoryAddBatchPayloadInnerImagesInner struct {
	value *CategoryAddBatchPayloadInnerImagesInner
	isSet bool
}

func (v NullableCategoryAddBatchPayloadInnerImagesInner) Get() *CategoryAddBatchPayloadInnerImagesInner {
	return v.value
}

func (v *NullableCategoryAddBatchPayloadInnerImagesInner) Set(val *CategoryAddBatchPayloadInnerImagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryAddBatchPayloadInnerImagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryAddBatchPayloadInnerImagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryAddBatchPayloadInnerImagesInner(val *CategoryAddBatchPayloadInnerImagesInner) *NullableCategoryAddBatchPayloadInnerImagesInner {
	return &NullableCategoryAddBatchPayloadInnerImagesInner{value: val, isSet: true}
}

func (v NullableCategoryAddBatchPayloadInnerImagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryAddBatchPayloadInnerImagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


