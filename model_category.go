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

// checks if the Category type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Category{}

// Category struct for Category
type Category struct {
	Id *string `json:"id,omitempty"`
	ParentId *string `json:"parent_id,omitempty"`
	CreatedAt *A2CDateTime `json:"created_at,omitempty"`
	ModifiedAt *A2CDateTime `json:"modified_at,omitempty"`
	Name *string `json:"name,omitempty"`
	ShortDescription *string `json:"short_description,omitempty"`
	Description *string `json:"description,omitempty"`
	StoresIds []string `json:"stores_ids,omitempty"`
	Keywords *string `json:"keywords,omitempty"`
	MetaDescription *string `json:"meta_description,omitempty"`
	MetaTitle *string `json:"meta_title,omitempty"`
	Avail *bool `json:"avail,omitempty"`
	Path *string `json:"path,omitempty"`
	SeoUrl *string `json:"seo_url,omitempty"`
	SortOrder *int32 `json:"sort_order,omitempty"`
	Images []Image `json:"images,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewCategory instantiates a new Category object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategory() *Category {
	this := Category{}
	return &this
}

// NewCategoryWithDefaults instantiates a new Category object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryWithDefaults() *Category {
	this := Category{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Category) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Category) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Category) SetId(v string) {
	o.Id = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *Category) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *Category) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *Category) SetParentId(v string) {
	o.ParentId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Category) GetCreatedAt() A2CDateTime {
	if o == nil || IsNil(o.CreatedAt) {
		var ret A2CDateTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetCreatedAtOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Category) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given A2CDateTime and assigns it to the CreatedAt field.
func (o *Category) SetCreatedAt(v A2CDateTime) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Category) GetModifiedAt() A2CDateTime {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret A2CDateTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetModifiedAtOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Category) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given A2CDateTime and assigns it to the ModifiedAt field.
func (o *Category) SetModifiedAt(v A2CDateTime) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Category) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Category) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Category) SetName(v string) {
	o.Name = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *Category) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *Category) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *Category) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Category) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Category) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Category) SetDescription(v string) {
	o.Description = &v
}

// GetStoresIds returns the StoresIds field value if set, zero value otherwise.
func (o *Category) GetStoresIds() []string {
	if o == nil || IsNil(o.StoresIds) {
		var ret []string
		return ret
	}
	return o.StoresIds
}

// GetStoresIdsOk returns a tuple with the StoresIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetStoresIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StoresIds) {
		return nil, false
	}
	return o.StoresIds, true
}

// HasStoresIds returns a boolean if a field has been set.
func (o *Category) HasStoresIds() bool {
	if o != nil && !IsNil(o.StoresIds) {
		return true
	}

	return false
}

// SetStoresIds gets a reference to the given []string and assigns it to the StoresIds field.
func (o *Category) SetStoresIds(v []string) {
	o.StoresIds = v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *Category) GetKeywords() string {
	if o == nil || IsNil(o.Keywords) {
		var ret string
		return ret
	}
	return *o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetKeywordsOk() (*string, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *Category) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given string and assigns it to the Keywords field.
func (o *Category) SetKeywords(v string) {
	o.Keywords = &v
}

// GetMetaDescription returns the MetaDescription field value if set, zero value otherwise.
func (o *Category) GetMetaDescription() string {
	if o == nil || IsNil(o.MetaDescription) {
		var ret string
		return ret
	}
	return *o.MetaDescription
}

// GetMetaDescriptionOk returns a tuple with the MetaDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetMetaDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.MetaDescription) {
		return nil, false
	}
	return o.MetaDescription, true
}

// HasMetaDescription returns a boolean if a field has been set.
func (o *Category) HasMetaDescription() bool {
	if o != nil && !IsNil(o.MetaDescription) {
		return true
	}

	return false
}

// SetMetaDescription gets a reference to the given string and assigns it to the MetaDescription field.
func (o *Category) SetMetaDescription(v string) {
	o.MetaDescription = &v
}

// GetMetaTitle returns the MetaTitle field value if set, zero value otherwise.
func (o *Category) GetMetaTitle() string {
	if o == nil || IsNil(o.MetaTitle) {
		var ret string
		return ret
	}
	return *o.MetaTitle
}

// GetMetaTitleOk returns a tuple with the MetaTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetMetaTitleOk() (*string, bool) {
	if o == nil || IsNil(o.MetaTitle) {
		return nil, false
	}
	return o.MetaTitle, true
}

// HasMetaTitle returns a boolean if a field has been set.
func (o *Category) HasMetaTitle() bool {
	if o != nil && !IsNil(o.MetaTitle) {
		return true
	}

	return false
}

// SetMetaTitle gets a reference to the given string and assigns it to the MetaTitle field.
func (o *Category) SetMetaTitle(v string) {
	o.MetaTitle = &v
}

// GetAvail returns the Avail field value if set, zero value otherwise.
func (o *Category) GetAvail() bool {
	if o == nil || IsNil(o.Avail) {
		var ret bool
		return ret
	}
	return *o.Avail
}

// GetAvailOk returns a tuple with the Avail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetAvailOk() (*bool, bool) {
	if o == nil || IsNil(o.Avail) {
		return nil, false
	}
	return o.Avail, true
}

// HasAvail returns a boolean if a field has been set.
func (o *Category) HasAvail() bool {
	if o != nil && !IsNil(o.Avail) {
		return true
	}

	return false
}

// SetAvail gets a reference to the given bool and assigns it to the Avail field.
func (o *Category) SetAvail(v bool) {
	o.Avail = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Category) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Category) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Category) SetPath(v string) {
	o.Path = &v
}

// GetSeoUrl returns the SeoUrl field value if set, zero value otherwise.
func (o *Category) GetSeoUrl() string {
	if o == nil || IsNil(o.SeoUrl) {
		var ret string
		return ret
	}
	return *o.SeoUrl
}

// GetSeoUrlOk returns a tuple with the SeoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetSeoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SeoUrl) {
		return nil, false
	}
	return o.SeoUrl, true
}

// HasSeoUrl returns a boolean if a field has been set.
func (o *Category) HasSeoUrl() bool {
	if o != nil && !IsNil(o.SeoUrl) {
		return true
	}

	return false
}

// SetSeoUrl gets a reference to the given string and assigns it to the SeoUrl field.
func (o *Category) SetSeoUrl(v string) {
	o.SeoUrl = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *Category) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *Category) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *Category) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *Category) GetImages() []Image {
	if o == nil || IsNil(o.Images) {
		var ret []Image
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetImagesOk() ([]Image, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Category) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []Image and assigns it to the Images field.
func (o *Category) SetImages(v []Image) {
	o.Images = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *Category) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *Category) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *Category) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Category) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Category) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Category) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Category) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Category) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["short_description"] = o.ShortDescription
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.StoresIds) {
		toSerialize["stores_ids"] = o.StoresIds
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.MetaDescription) {
		toSerialize["meta_description"] = o.MetaDescription
	}
	if !IsNil(o.MetaTitle) {
		toSerialize["meta_title"] = o.MetaTitle
	}
	if !IsNil(o.Avail) {
		toSerialize["avail"] = o.Avail
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.SeoUrl) {
		toSerialize["seo_url"] = o.SeoUrl
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sort_order"] = o.SortOrder
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

type NullableCategory struct {
	value *Category
	isSet bool
}

func (v NullableCategory) Get() *Category {
	return v.value
}

func (v *NullableCategory) Set(val *Category) {
	v.value = val
	v.isSet = true
}

func (v NullableCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategory(val *Category) *NullableCategory {
	return &NullableCategory{value: val, isSet: true}
}

func (v NullableCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


