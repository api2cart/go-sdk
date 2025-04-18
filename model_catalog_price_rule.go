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

// checks if the CatalogPriceRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogPriceRule{}

// CatalogPriceRule struct for CatalogPriceRule
type CatalogPriceRule struct {
	Id *string `json:"id,omitempty"`
	Gid *string `json:"gid,omitempty"`
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ShortDescription *string `json:"short_description,omitempty"`
	Avail *bool `json:"avail,omitempty"`
	Actions []CatalogPriceRuleAction `json:"actions,omitempty"`
	CreatedTime *A2CDateTime `json:"created_time,omitempty"`
	DateStart *A2CDateTime `json:"date_start,omitempty"`
	DateEnd *A2CDateTime `json:"date_end,omitempty"`
	UsageCount *float32 `json:"usage_count,omitempty"`
	Conditions []CouponCondition `json:"conditions,omitempty"`
	UsesPerOrderLimit *int32 `json:"uses_per_order_limit,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewCatalogPriceRule instantiates a new CatalogPriceRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogPriceRule() *CatalogPriceRule {
	this := CatalogPriceRule{}
	return &this
}

// NewCatalogPriceRuleWithDefaults instantiates a new CatalogPriceRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogPriceRuleWithDefaults() *CatalogPriceRule {
	this := CatalogPriceRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CatalogPriceRule) SetId(v string) {
	o.Id = &v
}

// GetGid returns the Gid field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetGid() string {
	if o == nil || IsNil(o.Gid) {
		var ret string
		return ret
	}
	return *o.Gid
}

// GetGidOk returns a tuple with the Gid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetGidOk() (*string, bool) {
	if o == nil || IsNil(o.Gid) {
		return nil, false
	}
	return o.Gid, true
}

// HasGid returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasGid() bool {
	if o != nil && !IsNil(o.Gid) {
		return true
	}

	return false
}

// SetGid gets a reference to the given string and assigns it to the Gid field.
func (o *CatalogPriceRule) SetGid(v string) {
	o.Gid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CatalogPriceRule) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogPriceRule) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CatalogPriceRule) SetDescription(v string) {
	o.Description = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *CatalogPriceRule) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetAvail returns the Avail field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetAvail() bool {
	if o == nil || IsNil(o.Avail) {
		var ret bool
		return ret
	}
	return *o.Avail
}

// GetAvailOk returns a tuple with the Avail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetAvailOk() (*bool, bool) {
	if o == nil || IsNil(o.Avail) {
		return nil, false
	}
	return o.Avail, true
}

// HasAvail returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasAvail() bool {
	if o != nil && !IsNil(o.Avail) {
		return true
	}

	return false
}

// SetAvail gets a reference to the given bool and assigns it to the Avail field.
func (o *CatalogPriceRule) SetAvail(v bool) {
	o.Avail = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetActions() []CatalogPriceRuleAction {
	if o == nil || IsNil(o.Actions) {
		var ret []CatalogPriceRuleAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetActionsOk() ([]CatalogPriceRuleAction, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []CatalogPriceRuleAction and assigns it to the Actions field.
func (o *CatalogPriceRule) SetActions(v []CatalogPriceRuleAction) {
	o.Actions = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetCreatedTime() A2CDateTime {
	if o == nil || IsNil(o.CreatedTime) {
		var ret A2CDateTime
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetCreatedTimeOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given A2CDateTime and assigns it to the CreatedTime field.
func (o *CatalogPriceRule) SetCreatedTime(v A2CDateTime) {
	o.CreatedTime = &v
}

// GetDateStart returns the DateStart field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetDateStart() A2CDateTime {
	if o == nil || IsNil(o.DateStart) {
		var ret A2CDateTime
		return ret
	}
	return *o.DateStart
}

// GetDateStartOk returns a tuple with the DateStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetDateStartOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.DateStart) {
		return nil, false
	}
	return o.DateStart, true
}

// HasDateStart returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasDateStart() bool {
	if o != nil && !IsNil(o.DateStart) {
		return true
	}

	return false
}

// SetDateStart gets a reference to the given A2CDateTime and assigns it to the DateStart field.
func (o *CatalogPriceRule) SetDateStart(v A2CDateTime) {
	o.DateStart = &v
}

// GetDateEnd returns the DateEnd field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetDateEnd() A2CDateTime {
	if o == nil || IsNil(o.DateEnd) {
		var ret A2CDateTime
		return ret
	}
	return *o.DateEnd
}

// GetDateEndOk returns a tuple with the DateEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetDateEndOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.DateEnd) {
		return nil, false
	}
	return o.DateEnd, true
}

// HasDateEnd returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasDateEnd() bool {
	if o != nil && !IsNil(o.DateEnd) {
		return true
	}

	return false
}

// SetDateEnd gets a reference to the given A2CDateTime and assigns it to the DateEnd field.
func (o *CatalogPriceRule) SetDateEnd(v A2CDateTime) {
	o.DateEnd = &v
}

// GetUsageCount returns the UsageCount field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetUsageCount() float32 {
	if o == nil || IsNil(o.UsageCount) {
		var ret float32
		return ret
	}
	return *o.UsageCount
}

// GetUsageCountOk returns a tuple with the UsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetUsageCountOk() (*float32, bool) {
	if o == nil || IsNil(o.UsageCount) {
		return nil, false
	}
	return o.UsageCount, true
}

// HasUsageCount returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasUsageCount() bool {
	if o != nil && !IsNil(o.UsageCount) {
		return true
	}

	return false
}

// SetUsageCount gets a reference to the given float32 and assigns it to the UsageCount field.
func (o *CatalogPriceRule) SetUsageCount(v float32) {
	o.UsageCount = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetConditions() []CouponCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []CouponCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetConditionsOk() ([]CouponCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []CouponCondition and assigns it to the Conditions field.
func (o *CatalogPriceRule) SetConditions(v []CouponCondition) {
	o.Conditions = v
}

// GetUsesPerOrderLimit returns the UsesPerOrderLimit field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetUsesPerOrderLimit() int32 {
	if o == nil || IsNil(o.UsesPerOrderLimit) {
		var ret int32
		return ret
	}
	return *o.UsesPerOrderLimit
}

// GetUsesPerOrderLimitOk returns a tuple with the UsesPerOrderLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetUsesPerOrderLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UsesPerOrderLimit) {
		return nil, false
	}
	return o.UsesPerOrderLimit, true
}

// HasUsesPerOrderLimit returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasUsesPerOrderLimit() bool {
	if o != nil && !IsNil(o.UsesPerOrderLimit) {
		return true
	}

	return false
}

// SetUsesPerOrderLimit gets a reference to the given int32 and assigns it to the UsesPerOrderLimit field.
func (o *CatalogPriceRule) SetUsesPerOrderLimit(v int32) {
	o.UsesPerOrderLimit = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *CatalogPriceRule) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CatalogPriceRule) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogPriceRule) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CatalogPriceRule) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CatalogPriceRule) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o CatalogPriceRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogPriceRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Gid) {
		toSerialize["gid"] = o.Gid
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["short_description"] = o.ShortDescription
	}
	if !IsNil(o.Avail) {
		toSerialize["avail"] = o.Avail
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.DateStart) {
		toSerialize["date_start"] = o.DateStart
	}
	if !IsNil(o.DateEnd) {
		toSerialize["date_end"] = o.DateEnd
	}
	if !IsNil(o.UsageCount) {
		toSerialize["usage_count"] = o.UsageCount
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.UsesPerOrderLimit) {
		toSerialize["uses_per_order_limit"] = o.UsesPerOrderLimit
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableCatalogPriceRule struct {
	value *CatalogPriceRule
	isSet bool
}

func (v NullableCatalogPriceRule) Get() *CatalogPriceRule {
	return v.value
}

func (v *NullableCatalogPriceRule) Set(val *CatalogPriceRule) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogPriceRule) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogPriceRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogPriceRule(val *CatalogPriceRule) *NullableCatalogPriceRule {
	return &NullableCatalogPriceRule{value: val, isSet: true}
}

func (v NullableCatalogPriceRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogPriceRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


