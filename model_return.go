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

// checks if the Return type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Return{}

// Return struct for Return
type Return struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	StoreId *string `json:"store_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	ModifiedAt *string `json:"modified_at,omitempty"`
	Status *ReturnStatus `json:"status,omitempty"`
	OrderProducts []ReturnOrderProduct `json:"order_products,omitempty"`
	Comment *string `json:"comment,omitempty"`
	StaffNote *string `json:"staff_note,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewReturn instantiates a new Return object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturn() *Return {
	this := Return{}
	return &this
}

// NewReturnWithDefaults instantiates a new Return object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnWithDefaults() *Return {
	this := Return{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Return) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Return) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Return) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Return) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Return) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Return) SetName(v string) {
	o.Name = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *Return) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *Return) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *Return) SetOrderId(v string) {
	o.OrderId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *Return) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *Return) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *Return) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *Return) GetStoreId() string {
	if o == nil || IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *Return) HasStoreId() bool {
	if o != nil && !IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *Return) SetStoreId(v string) {
	o.StoreId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Return) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Return) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Return) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Return) GetModifiedAt() string {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetModifiedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Return) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *Return) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Return) GetStatus() ReturnStatus {
	if o == nil || IsNil(o.Status) {
		var ret ReturnStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetStatusOk() (*ReturnStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Return) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ReturnStatus and assigns it to the Status field.
func (o *Return) SetStatus(v ReturnStatus) {
	o.Status = &v
}

// GetOrderProducts returns the OrderProducts field value if set, zero value otherwise.
func (o *Return) GetOrderProducts() []ReturnOrderProduct {
	if o == nil || IsNil(o.OrderProducts) {
		var ret []ReturnOrderProduct
		return ret
	}
	return o.OrderProducts
}

// GetOrderProductsOk returns a tuple with the OrderProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetOrderProductsOk() ([]ReturnOrderProduct, bool) {
	if o == nil || IsNil(o.OrderProducts) {
		return nil, false
	}
	return o.OrderProducts, true
}

// HasOrderProducts returns a boolean if a field has been set.
func (o *Return) HasOrderProducts() bool {
	if o != nil && !IsNil(o.OrderProducts) {
		return true
	}

	return false
}

// SetOrderProducts gets a reference to the given []ReturnOrderProduct and assigns it to the OrderProducts field.
func (o *Return) SetOrderProducts(v []ReturnOrderProduct) {
	o.OrderProducts = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Return) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Return) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Return) SetComment(v string) {
	o.Comment = &v
}

// GetStaffNote returns the StaffNote field value if set, zero value otherwise.
func (o *Return) GetStaffNote() string {
	if o == nil || IsNil(o.StaffNote) {
		var ret string
		return ret
	}
	return *o.StaffNote
}

// GetStaffNoteOk returns a tuple with the StaffNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetStaffNoteOk() (*string, bool) {
	if o == nil || IsNil(o.StaffNote) {
		return nil, false
	}
	return o.StaffNote, true
}

// HasStaffNote returns a boolean if a field has been set.
func (o *Return) HasStaffNote() bool {
	if o != nil && !IsNil(o.StaffNote) {
		return true
	}

	return false
}

// SetStaffNote gets a reference to the given string and assigns it to the StaffNote field.
func (o *Return) SetStaffNote(v string) {
	o.StaffNote = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *Return) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *Return) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *Return) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Return) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Return) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Return) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Return) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Return) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Return) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.StoreId) {
		toSerialize["store_id"] = o.StoreId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.OrderProducts) {
		toSerialize["order_products"] = o.OrderProducts
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.StaffNote) {
		toSerialize["staff_note"] = o.StaffNote
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableReturn struct {
	value *Return
	isSet bool
}

func (v NullableReturn) Get() *Return {
	return v.value
}

func (v *NullableReturn) Set(val *Return) {
	v.value = val
	v.isSet = true
}

func (v NullableReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturn(val *Return) *NullableReturn {
	return &NullableReturn{value: val, isSet: true}
}

func (v NullableReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


