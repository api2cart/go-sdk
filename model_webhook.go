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

// checks if the Webhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Webhook{}

// Webhook struct for Webhook
type Webhook struct {
	Id *int32 `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	StoreId *string `json:"store_id,omitempty"`
	LangId *string `json:"lang_id,omitempty"`
	Active *bool `json:"active,omitempty"`
	Callback *string `json:"callback,omitempty"`
	Fields *string `json:"fields,omitempty"`
	ResponseFields *string `json:"response_fields,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Entity *string `json:"entity,omitempty"`
	Action *string `json:"action,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewWebhook instantiates a new Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhook() *Webhook {
	this := Webhook{}
	return &this
}

// NewWebhookWithDefaults instantiates a new Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookWithDefaults() *Webhook {
	this := Webhook{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Webhook) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Webhook) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Webhook) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Webhook) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Webhook) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Webhook) SetLabel(v string) {
	o.Label = &v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *Webhook) GetStoreId() string {
	if o == nil || IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *Webhook) HasStoreId() bool {
	if o != nil && !IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *Webhook) SetStoreId(v string) {
	o.StoreId = &v
}

// GetLangId returns the LangId field value if set, zero value otherwise.
func (o *Webhook) GetLangId() string {
	if o == nil || IsNil(o.LangId) {
		var ret string
		return ret
	}
	return *o.LangId
}

// GetLangIdOk returns a tuple with the LangId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetLangIdOk() (*string, bool) {
	if o == nil || IsNil(o.LangId) {
		return nil, false
	}
	return o.LangId, true
}

// HasLangId returns a boolean if a field has been set.
func (o *Webhook) HasLangId() bool {
	if o != nil && !IsNil(o.LangId) {
		return true
	}

	return false
}

// SetLangId gets a reference to the given string and assigns it to the LangId field.
func (o *Webhook) SetLangId(v string) {
	o.LangId = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Webhook) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Webhook) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Webhook) SetActive(v bool) {
	o.Active = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *Webhook) GetCallback() string {
	if o == nil || IsNil(o.Callback) {
		var ret string
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetCallbackOk() (*string, bool) {
	if o == nil || IsNil(o.Callback) {
		return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *Webhook) HasCallback() bool {
	if o != nil && !IsNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given string and assigns it to the Callback field.
func (o *Webhook) SetCallback(v string) {
	o.Callback = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *Webhook) GetFields() string {
	if o == nil || IsNil(o.Fields) {
		var ret string
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetFieldsOk() (*string, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *Webhook) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given string and assigns it to the Fields field.
func (o *Webhook) SetFields(v string) {
	o.Fields = &v
}

// GetResponseFields returns the ResponseFields field value if set, zero value otherwise.
func (o *Webhook) GetResponseFields() string {
	if o == nil || IsNil(o.ResponseFields) {
		var ret string
		return ret
	}
	return *o.ResponseFields
}

// GetResponseFieldsOk returns a tuple with the ResponseFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetResponseFieldsOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseFields) {
		return nil, false
	}
	return o.ResponseFields, true
}

// HasResponseFields returns a boolean if a field has been set.
func (o *Webhook) HasResponseFields() bool {
	if o != nil && !IsNil(o.ResponseFields) {
		return true
	}

	return false
}

// SetResponseFields gets a reference to the given string and assigns it to the ResponseFields field.
func (o *Webhook) SetResponseFields(v string) {
	o.ResponseFields = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Webhook) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Webhook) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Webhook) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Webhook) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Webhook) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Webhook) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *Webhook) GetEntity() string {
	if o == nil || IsNil(o.Entity) {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEntityOk() (*string, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *Webhook) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *Webhook) SetEntity(v string) {
	o.Entity = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Webhook) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Webhook) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Webhook) SetAction(v string) {
	o.Action = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *Webhook) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *Webhook) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *Webhook) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Webhook) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Webhook) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Webhook) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Webhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Webhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.StoreId) {
		toSerialize["store_id"] = o.StoreId
	}
	if !IsNil(o.LangId) {
		toSerialize["lang_id"] = o.LangId
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.ResponseFields) {
		toSerialize["response_fields"] = o.ResponseFields
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableWebhook struct {
	value *Webhook
	isSet bool
}

func (v NullableWebhook) Get() *Webhook {
	return v.value
}

func (v *NullableWebhook) Set(val *Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhook(val *Webhook) *NullableWebhook {
	return &NullableWebhook{value: val, isSet: true}
}

func (v NullableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


