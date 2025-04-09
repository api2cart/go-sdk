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

// checks if the Customer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Customer{}

// Customer struct for Customer
type Customer struct {
	Id *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Phone *string `json:"phone,omitempty"`
	CreatedTime *A2CDateTime `json:"created_time,omitempty"`
	ModifiedTime *A2CDateTime `json:"modified_time,omitempty"`
	Group []CustomerGroup `json:"group,omitempty"`
	Login *string `json:"login,omitempty"`
	LastLogin *A2CDateTime `json:"last_login,omitempty"`
	BirthDay *A2CDateTime `json:"birth_day,omitempty"`
	Status *string `json:"status,omitempty"`
	NewsLetterSubscription *bool `json:"news_letter_subscription,omitempty"`
	Consents []CustomerConsent `json:"consents,omitempty"`
	Gender *string `json:"gender,omitempty"`
	StoresIds []string `json:"stores_ids,omitempty"`
	Website *string `json:"website,omitempty"`
	Fax *string `json:"fax,omitempty"`
	Company *string `json:"company,omitempty"`
	IpAddress *string `json:"ip_address,omitempty"`
	AddressBook []CustomerAddress `json:"address_book,omitempty"`
	LangId *string `json:"lang_id,omitempty"`
	OrdersCount *int32 `json:"orders_count,omitempty"`
	LastOrderId *string `json:"last_order_id,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewCustomer instantiates a new Customer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomer() *Customer {
	this := Customer{}
	return &this
}

// NewCustomerWithDefaults instantiates a new Customer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerWithDefaults() *Customer {
	this := Customer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Customer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Customer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Customer) SetId(v string) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Customer) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Customer) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Customer) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Customer) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Customer) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Customer) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Customer) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Customer) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Customer) SetLastName(v string) {
	o.LastName = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Customer) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Customer) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Customer) SetPhone(v string) {
	o.Phone = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Customer) GetCreatedTime() A2CDateTime {
	if o == nil || IsNil(o.CreatedTime) {
		var ret A2CDateTime
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetCreatedTimeOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Customer) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given A2CDateTime and assigns it to the CreatedTime field.
func (o *Customer) SetCreatedTime(v A2CDateTime) {
	o.CreatedTime = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Customer) GetModifiedTime() A2CDateTime {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret A2CDateTime
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetModifiedTimeOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Customer) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given A2CDateTime and assigns it to the ModifiedTime field.
func (o *Customer) SetModifiedTime(v A2CDateTime) {
	o.ModifiedTime = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *Customer) GetGroup() []CustomerGroup {
	if o == nil || IsNil(o.Group) {
		var ret []CustomerGroup
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetGroupOk() ([]CustomerGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *Customer) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given []CustomerGroup and assigns it to the Group field.
func (o *Customer) SetGroup(v []CustomerGroup) {
	o.Group = v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *Customer) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *Customer) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *Customer) SetLogin(v string) {
	o.Login = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *Customer) GetLastLogin() A2CDateTime {
	if o == nil || IsNil(o.LastLogin) {
		var ret A2CDateTime
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetLastLoginOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *Customer) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given A2CDateTime and assigns it to the LastLogin field.
func (o *Customer) SetLastLogin(v A2CDateTime) {
	o.LastLogin = &v
}

// GetBirthDay returns the BirthDay field value if set, zero value otherwise.
func (o *Customer) GetBirthDay() A2CDateTime {
	if o == nil || IsNil(o.BirthDay) {
		var ret A2CDateTime
		return ret
	}
	return *o.BirthDay
}

// GetBirthDayOk returns a tuple with the BirthDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetBirthDayOk() (*A2CDateTime, bool) {
	if o == nil || IsNil(o.BirthDay) {
		return nil, false
	}
	return o.BirthDay, true
}

// HasBirthDay returns a boolean if a field has been set.
func (o *Customer) HasBirthDay() bool {
	if o != nil && !IsNil(o.BirthDay) {
		return true
	}

	return false
}

// SetBirthDay gets a reference to the given A2CDateTime and assigns it to the BirthDay field.
func (o *Customer) SetBirthDay(v A2CDateTime) {
	o.BirthDay = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Customer) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Customer) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Customer) SetStatus(v string) {
	o.Status = &v
}

// GetNewsLetterSubscription returns the NewsLetterSubscription field value if set, zero value otherwise.
func (o *Customer) GetNewsLetterSubscription() bool {
	if o == nil || IsNil(o.NewsLetterSubscription) {
		var ret bool
		return ret
	}
	return *o.NewsLetterSubscription
}

// GetNewsLetterSubscriptionOk returns a tuple with the NewsLetterSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetNewsLetterSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.NewsLetterSubscription) {
		return nil, false
	}
	return o.NewsLetterSubscription, true
}

// HasNewsLetterSubscription returns a boolean if a field has been set.
func (o *Customer) HasNewsLetterSubscription() bool {
	if o != nil && !IsNil(o.NewsLetterSubscription) {
		return true
	}

	return false
}

// SetNewsLetterSubscription gets a reference to the given bool and assigns it to the NewsLetterSubscription field.
func (o *Customer) SetNewsLetterSubscription(v bool) {
	o.NewsLetterSubscription = &v
}

// GetConsents returns the Consents field value if set, zero value otherwise.
func (o *Customer) GetConsents() []CustomerConsent {
	if o == nil || IsNil(o.Consents) {
		var ret []CustomerConsent
		return ret
	}
	return o.Consents
}

// GetConsentsOk returns a tuple with the Consents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetConsentsOk() ([]CustomerConsent, bool) {
	if o == nil || IsNil(o.Consents) {
		return nil, false
	}
	return o.Consents, true
}

// HasConsents returns a boolean if a field has been set.
func (o *Customer) HasConsents() bool {
	if o != nil && !IsNil(o.Consents) {
		return true
	}

	return false
}

// SetConsents gets a reference to the given []CustomerConsent and assigns it to the Consents field.
func (o *Customer) SetConsents(v []CustomerConsent) {
	o.Consents = v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *Customer) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *Customer) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *Customer) SetGender(v string) {
	o.Gender = &v
}

// GetStoresIds returns the StoresIds field value if set, zero value otherwise.
func (o *Customer) GetStoresIds() []string {
	if o == nil || IsNil(o.StoresIds) {
		var ret []string
		return ret
	}
	return o.StoresIds
}

// GetStoresIdsOk returns a tuple with the StoresIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetStoresIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StoresIds) {
		return nil, false
	}
	return o.StoresIds, true
}

// HasStoresIds returns a boolean if a field has been set.
func (o *Customer) HasStoresIds() bool {
	if o != nil && !IsNil(o.StoresIds) {
		return true
	}

	return false
}

// SetStoresIds gets a reference to the given []string and assigns it to the StoresIds field.
func (o *Customer) SetStoresIds(v []string) {
	o.StoresIds = v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *Customer) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *Customer) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *Customer) SetWebsite(v string) {
	o.Website = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *Customer) GetFax() string {
	if o == nil || IsNil(o.Fax) {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetFaxOk() (*string, bool) {
	if o == nil || IsNil(o.Fax) {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *Customer) HasFax() bool {
	if o != nil && !IsNil(o.Fax) {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *Customer) SetFax(v string) {
	o.Fax = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *Customer) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *Customer) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *Customer) SetCompany(v string) {
	o.Company = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *Customer) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *Customer) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *Customer) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetAddressBook returns the AddressBook field value if set, zero value otherwise.
func (o *Customer) GetAddressBook() []CustomerAddress {
	if o == nil || IsNil(o.AddressBook) {
		var ret []CustomerAddress
		return ret
	}
	return o.AddressBook
}

// GetAddressBookOk returns a tuple with the AddressBook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetAddressBookOk() ([]CustomerAddress, bool) {
	if o == nil || IsNil(o.AddressBook) {
		return nil, false
	}
	return o.AddressBook, true
}

// HasAddressBook returns a boolean if a field has been set.
func (o *Customer) HasAddressBook() bool {
	if o != nil && !IsNil(o.AddressBook) {
		return true
	}

	return false
}

// SetAddressBook gets a reference to the given []CustomerAddress and assigns it to the AddressBook field.
func (o *Customer) SetAddressBook(v []CustomerAddress) {
	o.AddressBook = v
}

// GetLangId returns the LangId field value if set, zero value otherwise.
func (o *Customer) GetLangId() string {
	if o == nil || IsNil(o.LangId) {
		var ret string
		return ret
	}
	return *o.LangId
}

// GetLangIdOk returns a tuple with the LangId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetLangIdOk() (*string, bool) {
	if o == nil || IsNil(o.LangId) {
		return nil, false
	}
	return o.LangId, true
}

// HasLangId returns a boolean if a field has been set.
func (o *Customer) HasLangId() bool {
	if o != nil && !IsNil(o.LangId) {
		return true
	}

	return false
}

// SetLangId gets a reference to the given string and assigns it to the LangId field.
func (o *Customer) SetLangId(v string) {
	o.LangId = &v
}

// GetOrdersCount returns the OrdersCount field value if set, zero value otherwise.
func (o *Customer) GetOrdersCount() int32 {
	if o == nil || IsNil(o.OrdersCount) {
		var ret int32
		return ret
	}
	return *o.OrdersCount
}

// GetOrdersCountOk returns a tuple with the OrdersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetOrdersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OrdersCount) {
		return nil, false
	}
	return o.OrdersCount, true
}

// HasOrdersCount returns a boolean if a field has been set.
func (o *Customer) HasOrdersCount() bool {
	if o != nil && !IsNil(o.OrdersCount) {
		return true
	}

	return false
}

// SetOrdersCount gets a reference to the given int32 and assigns it to the OrdersCount field.
func (o *Customer) SetOrdersCount(v int32) {
	o.OrdersCount = &v
}

// GetLastOrderId returns the LastOrderId field value if set, zero value otherwise.
func (o *Customer) GetLastOrderId() string {
	if o == nil || IsNil(o.LastOrderId) {
		var ret string
		return ret
	}
	return *o.LastOrderId
}

// GetLastOrderIdOk returns a tuple with the LastOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetLastOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastOrderId) {
		return nil, false
	}
	return o.LastOrderId, true
}

// HasLastOrderId returns a boolean if a field has been set.
func (o *Customer) HasLastOrderId() bool {
	if o != nil && !IsNil(o.LastOrderId) {
		return true
	}

	return false
}

// SetLastOrderId gets a reference to the given string and assigns it to the LastOrderId field.
func (o *Customer) SetLastOrderId(v string) {
	o.LastOrderId = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *Customer) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *Customer) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *Customer) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Customer) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Customer) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Customer) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Customer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Customer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !IsNil(o.LastLogin) {
		toSerialize["last_login"] = o.LastLogin
	}
	if !IsNil(o.BirthDay) {
		toSerialize["birth_day"] = o.BirthDay
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.NewsLetterSubscription) {
		toSerialize["news_letter_subscription"] = o.NewsLetterSubscription
	}
	if !IsNil(o.Consents) {
		toSerialize["consents"] = o.Consents
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.StoresIds) {
		toSerialize["stores_ids"] = o.StoresIds
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.Fax) {
		toSerialize["fax"] = o.Fax
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.AddressBook) {
		toSerialize["address_book"] = o.AddressBook
	}
	if !IsNil(o.LangId) {
		toSerialize["lang_id"] = o.LangId
	}
	if !IsNil(o.OrdersCount) {
		toSerialize["orders_count"] = o.OrdersCount
	}
	if !IsNil(o.LastOrderId) {
		toSerialize["last_order_id"] = o.LastOrderId
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableCustomer struct {
	value *Customer
	isSet bool
}

func (v NullableCustomer) Get() *Customer {
	return v.value
}

func (v *NullableCustomer) Set(val *Customer) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomer(val *Customer) *NullableCustomer {
	return &NullableCustomer{value: val, isSet: true}
}

func (v NullableCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


