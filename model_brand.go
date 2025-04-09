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

// checks if the Brand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Brand{}

// Brand struct for Brand
type Brand struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	CreatedTime *string `json:"created_time,omitempty"`
	ModifiedTime *string `json:"modified_time,omitempty"`
	FullDescription *string `json:"full_description,omitempty"`
	ShortDescription *string `json:"short_description,omitempty"`
	StoresIds []string `json:"stores_ids,omitempty"`
	Active *bool `json:"active,omitempty"`
	Url *string `json:"url,omitempty"`
	MetaTitle *string `json:"meta_title,omitempty"`
	MetaKeywords *string `json:"meta_keywords,omitempty"`
	MetaDescription *string `json:"meta_description,omitempty"`
	Images []Image `json:"images,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewBrand instantiates a new Brand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrand() *Brand {
	this := Brand{}
	return &this
}

// NewBrandWithDefaults instantiates a new Brand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandWithDefaults() *Brand {
	this := Brand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Brand) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Brand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Brand) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Brand) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Brand) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Brand) SetName(v string) {
	o.Name = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Brand) GetCreatedTime() string {
	if o == nil || IsNil(o.CreatedTime) {
		var ret string
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetCreatedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Brand) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given string and assigns it to the CreatedTime field.
func (o *Brand) SetCreatedTime(v string) {
	o.CreatedTime = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Brand) GetModifiedTime() string {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret string
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetModifiedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Brand) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given string and assigns it to the ModifiedTime field.
func (o *Brand) SetModifiedTime(v string) {
	o.ModifiedTime = &v
}

// GetFullDescription returns the FullDescription field value if set, zero value otherwise.
func (o *Brand) GetFullDescription() string {
	if o == nil || IsNil(o.FullDescription) {
		var ret string
		return ret
	}
	return *o.FullDescription
}

// GetFullDescriptionOk returns a tuple with the FullDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetFullDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.FullDescription) {
		return nil, false
	}
	return o.FullDescription, true
}

// HasFullDescription returns a boolean if a field has been set.
func (o *Brand) HasFullDescription() bool {
	if o != nil && !IsNil(o.FullDescription) {
		return true
	}

	return false
}

// SetFullDescription gets a reference to the given string and assigns it to the FullDescription field.
func (o *Brand) SetFullDescription(v string) {
	o.FullDescription = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *Brand) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *Brand) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *Brand) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetStoresIds returns the StoresIds field value if set, zero value otherwise.
func (o *Brand) GetStoresIds() []string {
	if o == nil || IsNil(o.StoresIds) {
		var ret []string
		return ret
	}
	return o.StoresIds
}

// GetStoresIdsOk returns a tuple with the StoresIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetStoresIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StoresIds) {
		return nil, false
	}
	return o.StoresIds, true
}

// HasStoresIds returns a boolean if a field has been set.
func (o *Brand) HasStoresIds() bool {
	if o != nil && !IsNil(o.StoresIds) {
		return true
	}

	return false
}

// SetStoresIds gets a reference to the given []string and assigns it to the StoresIds field.
func (o *Brand) SetStoresIds(v []string) {
	o.StoresIds = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Brand) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Brand) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Brand) SetActive(v bool) {
	o.Active = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Brand) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Brand) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Brand) SetUrl(v string) {
	o.Url = &v
}

// GetMetaTitle returns the MetaTitle field value if set, zero value otherwise.
func (o *Brand) GetMetaTitle() string {
	if o == nil || IsNil(o.MetaTitle) {
		var ret string
		return ret
	}
	return *o.MetaTitle
}

// GetMetaTitleOk returns a tuple with the MetaTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetMetaTitleOk() (*string, bool) {
	if o == nil || IsNil(o.MetaTitle) {
		return nil, false
	}
	return o.MetaTitle, true
}

// HasMetaTitle returns a boolean if a field has been set.
func (o *Brand) HasMetaTitle() bool {
	if o != nil && !IsNil(o.MetaTitle) {
		return true
	}

	return false
}

// SetMetaTitle gets a reference to the given string and assigns it to the MetaTitle field.
func (o *Brand) SetMetaTitle(v string) {
	o.MetaTitle = &v
}

// GetMetaKeywords returns the MetaKeywords field value if set, zero value otherwise.
func (o *Brand) GetMetaKeywords() string {
	if o == nil || IsNil(o.MetaKeywords) {
		var ret string
		return ret
	}
	return *o.MetaKeywords
}

// GetMetaKeywordsOk returns a tuple with the MetaKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetMetaKeywordsOk() (*string, bool) {
	if o == nil || IsNil(o.MetaKeywords) {
		return nil, false
	}
	return o.MetaKeywords, true
}

// HasMetaKeywords returns a boolean if a field has been set.
func (o *Brand) HasMetaKeywords() bool {
	if o != nil && !IsNil(o.MetaKeywords) {
		return true
	}

	return false
}

// SetMetaKeywords gets a reference to the given string and assigns it to the MetaKeywords field.
func (o *Brand) SetMetaKeywords(v string) {
	o.MetaKeywords = &v
}

// GetMetaDescription returns the MetaDescription field value if set, zero value otherwise.
func (o *Brand) GetMetaDescription() string {
	if o == nil || IsNil(o.MetaDescription) {
		var ret string
		return ret
	}
	return *o.MetaDescription
}

// GetMetaDescriptionOk returns a tuple with the MetaDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetMetaDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.MetaDescription) {
		return nil, false
	}
	return o.MetaDescription, true
}

// HasMetaDescription returns a boolean if a field has been set.
func (o *Brand) HasMetaDescription() bool {
	if o != nil && !IsNil(o.MetaDescription) {
		return true
	}

	return false
}

// SetMetaDescription gets a reference to the given string and assigns it to the MetaDescription field.
func (o *Brand) SetMetaDescription(v string) {
	o.MetaDescription = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *Brand) GetImages() []Image {
	if o == nil || IsNil(o.Images) {
		var ret []Image
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetImagesOk() ([]Image, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Brand) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []Image and assigns it to the Images field.
func (o *Brand) SetImages(v []Image) {
	o.Images = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *Brand) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *Brand) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *Brand) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Brand) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Brand) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Brand) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Brand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Brand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.FullDescription) {
		toSerialize["full_description"] = o.FullDescription
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["short_description"] = o.ShortDescription
	}
	if !IsNil(o.StoresIds) {
		toSerialize["stores_ids"] = o.StoresIds
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.MetaTitle) {
		toSerialize["meta_title"] = o.MetaTitle
	}
	if !IsNil(o.MetaKeywords) {
		toSerialize["meta_keywords"] = o.MetaKeywords
	}
	if !IsNil(o.MetaDescription) {
		toSerialize["meta_description"] = o.MetaDescription
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableBrand struct {
	value *Brand
	isSet bool
}

func (v NullableBrand) Get() *Brand {
	return v.value
}

func (v *NullableBrand) Set(val *Brand) {
	v.value = val
	v.isSet = true
}

func (v NullableBrand) IsSet() bool {
	return v.isSet
}

func (v *NullableBrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrand(val *Brand) *NullableBrand {
	return &NullableBrand{value: val, isSet: true}
}

func (v NullableBrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


