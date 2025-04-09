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

// checks if the CustomerWishList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerWishList{}

// CustomerWishList struct for CustomerWishList
type CustomerWishList struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsPublic *string `json:"is_public,omitempty"`
	CreatedAt *A2CDateTime `json:"created_at,omitempty"`
	ModifiedAt *A2CDateTime `json:"modified_at,omitempty"`
	Products []CustomerWishListItem `json:"products,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewCustomerWishList instantiates a new CustomerWishList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerWishList() *CustomerWishList {
	this := CustomerWishList{}
	return &this
}

// NewCustomerWishListWithDefaults instantiates a new CustomerWishList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerWishListWithDefaults() *CustomerWishList {
	this := CustomerWishList{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerWishList) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerWishList) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerWishList) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerWishList) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerWishList) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerWishList) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerWishList) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerWishList) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomerWishList) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerWishList) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomerWishList) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomerWishList) SetDescription(v string) {
	o.Description = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *CustomerWishList) GetIsPublic() string {
	if o == nil || IsNil(o.IsPublic) {
		var ret string
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerWishList) GetIsPublicOk() (*string, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *CustomerWishList) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given string and assigns it to the IsPublic field.
func (o *CustomerWishList) SetIsPublic(v string) {
	o.IsPublic = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomerWishList) GetCreatedAt() A2CDateTime {
	if o == nil || IsNil(o.CreatedAt) {
		var ret A2CDateTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerWishList) GetCreatedAtOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerWishList) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given A2CDateTime and assigns it to the CreatedAt field.
func (o *CustomerWishList) SetCreatedAt(v A2CDateTime) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *CustomerWishList) GetModifiedAt() A2CDateTime {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret A2CDateTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerWishList) GetModifiedAtOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *CustomerWishList) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given A2CDateTime and assigns it to the ModifiedAt field.
func (o *CustomerWishList) SetModifiedAt(v A2CDateTime) {
	o.ModifiedAt = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *CustomerWishList) GetProducts() []CustomerWishListItem {
	if o == nil || IsNil(o.Products) {
		var ret []CustomerWishListItem
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerWishList) GetProductsOk() ([]CustomerWishListItem, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *CustomerWishList) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []CustomerWishListItem and assigns it to the Products field.
func (o *CustomerWishList) SetProducts(v []CustomerWishListItem) {
	o.Products = v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *CustomerWishList) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerWishList) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *CustomerWishList) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *CustomerWishList) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CustomerWishList) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerWishList) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CustomerWishList) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CustomerWishList) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o CustomerWishList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerWishList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableCustomerWishList struct {
	value *CustomerWishList
	isSet bool
}

func (v NullableCustomerWishList) Get() *CustomerWishList {
	return v.value
}

func (v *NullableCustomerWishList) Set(val *CustomerWishList) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerWishList) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerWishList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerWishList(val *CustomerWishList) *NullableCustomerWishList {
	return &NullableCustomerWishList{value: val, isSet: true}
}

func (v NullableCustomerWishList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerWishList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


