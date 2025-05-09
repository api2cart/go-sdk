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

// checks if the Info type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Info{}

// Info struct for Info
type Info struct {
	Owner *string `json:"owner,omitempty"`
	Country *string `json:"country,omitempty"`
	State *string `json:"state,omitempty"`
	StateCode *string `json:"state_code,omitempty"`
	City *string `json:"city,omitempty"`
	StreetAddress *string `json:"street_address,omitempty"`
	StreetAddressLine2 *string `json:"street_address_line_2,omitempty"`
	ZipCode *string `json:"zip_code,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewInfo instantiates a new Info object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfo() *Info {
	this := Info{}
	return &this
}

// NewInfoWithDefaults instantiates a new Info object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfoWithDefaults() *Info {
	this := Info{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Info) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Info) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *Info) SetOwner(v string) {
	o.Owner = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Info) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Info) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Info) SetCountry(v string) {
	o.Country = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Info) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Info) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Info) SetState(v string) {
	o.State = &v
}

// GetStateCode returns the StateCode field value if set, zero value otherwise.
func (o *Info) GetStateCode() string {
	if o == nil || IsNil(o.StateCode) {
		var ret string
		return ret
	}
	return *o.StateCode
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetStateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StateCode) {
		return nil, false
	}
	return o.StateCode, true
}

// HasStateCode returns a boolean if a field has been set.
func (o *Info) HasStateCode() bool {
	if o != nil && !IsNil(o.StateCode) {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given string and assigns it to the StateCode field.
func (o *Info) SetStateCode(v string) {
	o.StateCode = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Info) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Info) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Info) SetCity(v string) {
	o.City = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *Info) GetStreetAddress() string {
	if o == nil || IsNil(o.StreetAddress) {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetStreetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *Info) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *Info) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetStreetAddressLine2 returns the StreetAddressLine2 field value if set, zero value otherwise.
func (o *Info) GetStreetAddressLine2() string {
	if o == nil || IsNil(o.StreetAddressLine2) {
		var ret string
		return ret
	}
	return *o.StreetAddressLine2
}

// GetStreetAddressLine2Ok returns a tuple with the StreetAddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetStreetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.StreetAddressLine2) {
		return nil, false
	}
	return o.StreetAddressLine2, true
}

// HasStreetAddressLine2 returns a boolean if a field has been set.
func (o *Info) HasStreetAddressLine2() bool {
	if o != nil && !IsNil(o.StreetAddressLine2) {
		return true
	}

	return false
}

// SetStreetAddressLine2 gets a reference to the given string and assigns it to the StreetAddressLine2 field.
func (o *Info) SetStreetAddressLine2(v string) {
	o.StreetAddressLine2 = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *Info) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode) {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetZipCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *Info) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *Info) SetZipCode(v string) {
	o.ZipCode = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Info) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Info) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Info) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Info) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Info) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Info) SetPhone(v string) {
	o.Phone = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *Info) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *Info) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *Info) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Info) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Info) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Info) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Info) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Info) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StateCode) {
		toSerialize["state_code"] = o.StateCode
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.StreetAddress) {
		toSerialize["street_address"] = o.StreetAddress
	}
	if !IsNil(o.StreetAddressLine2) {
		toSerialize["street_address_line_2"] = o.StreetAddressLine2
	}
	if !IsNil(o.ZipCode) {
		toSerialize["zip_code"] = o.ZipCode
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableInfo struct {
	value *Info
	isSet bool
}

func (v NullableInfo) Get() *Info {
	return v.value
}

func (v *NullableInfo) Set(val *Info) {
	v.value = val
	v.isSet = true
}

func (v NullableInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfo(val *Info) *NullableInfo {
	return &NullableInfo{value: val, isSet: true}
}

func (v NullableInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


