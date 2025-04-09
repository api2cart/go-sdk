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

// checks if the ProductVariantImageAdd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductVariantImageAdd{}

// ProductVariantImageAdd struct for ProductVariantImageAdd
type ProductVariantImageAdd struct {
	// Defines product id where the variant image has to be added
	ProductId *string `json:"product_id,omitempty"`
	// Defines product's variants specified by variant id
	ProductVariantId string `json:"product_variant_id"`
	// Defines image's name
	ImageName string `json:"image_name"`
	// Defines image's types that are specified by comma-separated list
	Type string `json:"type"`
	// Defines URL of the image that has to be added
	Url *string `json:"url,omitempty"`
	// Content(body) encoded in base64 of image file
	Content *string `json:"content,omitempty"`
	// Defines alternative text that has to be attached to the picture
	Label *string `json:"label,omitempty"`
	// Mime type of image http://en.wikipedia.org/wiki/Internet_media_type.
	Mime *string `json:"mime,omitempty"`
	// Defines image’s position in the list
	Position *int32 `json:"position,omitempty"`
	// Store Id
	StoreId *string `json:"store_id,omitempty"`
	// Defines option id of the product variant for which the image will be added
	OptionId *string `json:"option_id,omitempty"`
}

type _ProductVariantImageAdd ProductVariantImageAdd

// NewProductVariantImageAdd instantiates a new ProductVariantImageAdd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductVariantImageAdd(productVariantId string, imageName string, type_ string) *ProductVariantImageAdd {
	this := ProductVariantImageAdd{}
	this.ProductVariantId = productVariantId
	this.ImageName = imageName
	this.Type = type_
	var position int32 = 0
	this.Position = &position
	return &this
}

// NewProductVariantImageAddWithDefaults instantiates a new ProductVariantImageAdd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductVariantImageAddWithDefaults() *ProductVariantImageAdd {
	this := ProductVariantImageAdd{}
	var type_ string = "base"
	this.Type = type_
	var position int32 = 0
	this.Position = &position
	return &this
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ProductVariantImageAdd) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVariantImageAdd) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ProductVariantImageAdd) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ProductVariantImageAdd) SetProductId(v string) {
	o.ProductId = &v
}

// GetProductVariantId returns the ProductVariantId field value
func (o *ProductVariantImageAdd) GetProductVariantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductVariantId
}

// GetProductVariantIdOk returns a tuple with the ProductVariantId field value
// and a boolean to check if the value has been set.
func (o *ProductVariantImageAdd) GetProductVariantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductVariantId, true
}

// SetProductVariantId sets field value
func (o *ProductVariantImageAdd) SetProductVariantId(v string) {
	o.ProductVariantId = v
}

// GetImageName returns the ImageName field value
func (o *ProductVariantImageAdd) GetImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value
// and a boolean to check if the value has been set.
func (o *ProductVariantImageAdd) GetImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageName, true
}

// SetImageName sets field value
func (o *ProductVariantImageAdd) SetImageName(v string) {
	o.ImageName = v
}

// GetType returns the Type field value
func (o *ProductVariantImageAdd) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductVariantImageAdd) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProductVariantImageAdd) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ProductVariantImageAdd) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVariantImageAdd) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ProductVariantImageAdd) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ProductVariantImageAdd) SetUrl(v string) {
	o.Url = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ProductVariantImageAdd) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVariantImageAdd) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ProductVariantImageAdd) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ProductVariantImageAdd) SetContent(v string) {
	o.Content = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductVariantImageAdd) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVariantImageAdd) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductVariantImageAdd) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ProductVariantImageAdd) SetLabel(v string) {
	o.Label = &v
}

// GetMime returns the Mime field value if set, zero value otherwise.
func (o *ProductVariantImageAdd) GetMime() string {
	if o == nil || IsNil(o.Mime) {
		var ret string
		return ret
	}
	return *o.Mime
}

// GetMimeOk returns a tuple with the Mime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVariantImageAdd) GetMimeOk() (*string, bool) {
	if o == nil || IsNil(o.Mime) {
		return nil, false
	}
	return o.Mime, true
}

// HasMime returns a boolean if a field has been set.
func (o *ProductVariantImageAdd) HasMime() bool {
	if o != nil && !IsNil(o.Mime) {
		return true
	}

	return false
}

// SetMime gets a reference to the given string and assigns it to the Mime field.
func (o *ProductVariantImageAdd) SetMime(v string) {
	o.Mime = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductVariantImageAdd) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVariantImageAdd) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductVariantImageAdd) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *ProductVariantImageAdd) SetPosition(v int32) {
	o.Position = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *ProductVariantImageAdd) GetStoreId() string {
	if o == nil || IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVariantImageAdd) GetStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *ProductVariantImageAdd) HasStoreId() bool {
	if o != nil && !IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *ProductVariantImageAdd) SetStoreId(v string) {
	o.StoreId = &v
}

// GetOptionId returns the OptionId field value if set, zero value otherwise.
func (o *ProductVariantImageAdd) GetOptionId() string {
	if o == nil || IsNil(o.OptionId) {
		var ret string
		return ret
	}
	return *o.OptionId
}

// GetOptionIdOk returns a tuple with the OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductVariantImageAdd) GetOptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptionId) {
		return nil, false
	}
	return o.OptionId, true
}

// HasOptionId returns a boolean if a field has been set.
func (o *ProductVariantImageAdd) HasOptionId() bool {
	if o != nil && !IsNil(o.OptionId) {
		return true
	}

	return false
}

// SetOptionId gets a reference to the given string and assigns it to the OptionId field.
func (o *ProductVariantImageAdd) SetOptionId(v string) {
	o.OptionId = &v
}

func (o ProductVariantImageAdd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductVariantImageAdd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	toSerialize["product_variant_id"] = o.ProductVariantId
	toSerialize["image_name"] = o.ImageName
	toSerialize["type"] = o.Type
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Mime) {
		toSerialize["mime"] = o.Mime
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.StoreId) {
		toSerialize["store_id"] = o.StoreId
	}
	if !IsNil(o.OptionId) {
		toSerialize["option_id"] = o.OptionId
	}
	return toSerialize, nil
}

func (o *ProductVariantImageAdd) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product_variant_id",
		"image_name",
		"type",
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

	varProductVariantImageAdd := _ProductVariantImageAdd{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductVariantImageAdd)

	if err != nil {
		return err
	}

	*o = ProductVariantImageAdd(varProductVariantImageAdd)

	return err
}

type NullableProductVariantImageAdd struct {
	value *ProductVariantImageAdd
	isSet bool
}

func (v NullableProductVariantImageAdd) Get() *ProductVariantImageAdd {
	return v.value
}

func (v *NullableProductVariantImageAdd) Set(val *ProductVariantImageAdd) {
	v.value = val
	v.isSet = true
}

func (v NullableProductVariantImageAdd) IsSet() bool {
	return v.isSet
}

func (v *NullableProductVariantImageAdd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductVariantImageAdd(val *ProductVariantImageAdd) *NullableProductVariantImageAdd {
	return &NullableProductVariantImageAdd{value: val, isSet: true}
}

func (v NullableProductVariantImageAdd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductVariantImageAdd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


