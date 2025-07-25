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
	FirstName NullableString `json:"first_name,omitempty"`
	LastName NullableString `json:"last_name,omitempty"`
	Phone NullableString `json:"phone,omitempty"`
	CreatedTime NullableA2CDateTime `json:"created_time,omitempty"`
	ModifiedTime NullableA2CDateTime `json:"modified_time,omitempty"`
	Group []CustomerGroup `json:"group,omitempty"`
	Login NullableString `json:"login,omitempty"`
	LastLogin NullableA2CDateTime `json:"last_login,omitempty"`
	BirthDay NullableA2CDateTime `json:"birth_day,omitempty"`
	Status NullableString `json:"status,omitempty"`
	IsGuest NullableBool `json:"is_guest,omitempty"`
	NewsLetterSubscription NullableBool `json:"news_letter_subscription,omitempty"`
	Consents []CustomerConsent `json:"consents,omitempty"`
	Gender NullableString `json:"gender,omitempty"`
	StoresIds []string `json:"stores_ids,omitempty"`
	Website NullableString `json:"website,omitempty"`
	Fax NullableString `json:"fax,omitempty"`
	Company NullableString `json:"company,omitempty"`
	IpAddress NullableString `json:"ip_address,omitempty"`
	AddressBook []CustomerAddress `json:"address_book,omitempty"`
	LangId NullableString `json:"lang_id,omitempty"`
	OrdersCount NullableInt32 `json:"orders_count,omitempty"`
	LastOrderId NullableString `json:"last_order_id,omitempty"`
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

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *Customer) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *Customer) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *Customer) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *Customer) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *Customer) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *Customer) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *Customer) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *Customer) UnsetLastName() {
	o.LastName.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *Customer) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *Customer) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *Customer) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *Customer) UnsetPhone() {
	o.Phone.Unset()
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetCreatedTime() A2CDateTime {
	if o == nil || IsNil(o.CreatedTime.Get()) {
		var ret A2CDateTime
		return ret
	}
	return *o.CreatedTime.Get()
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetCreatedTimeOk() (*A2CDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedTime.Get(), o.CreatedTime.IsSet()
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Customer) HasCreatedTime() bool {
	if o != nil && o.CreatedTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given NullableA2CDateTime and assigns it to the CreatedTime field.
func (o *Customer) SetCreatedTime(v A2CDateTime) {
	o.CreatedTime.Set(&v)
}
// SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil
func (o *Customer) SetCreatedTimeNil() {
	o.CreatedTime.Set(nil)
}

// UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
func (o *Customer) UnsetCreatedTime() {
	o.CreatedTime.Unset()
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetModifiedTime() A2CDateTime {
	if o == nil || IsNil(o.ModifiedTime.Get()) {
		var ret A2CDateTime
		return ret
	}
	return *o.ModifiedTime.Get()
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetModifiedTimeOk() (*A2CDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModifiedTime.Get(), o.ModifiedTime.IsSet()
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Customer) HasModifiedTime() bool {
	if o != nil && o.ModifiedTime.IsSet() {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given NullableA2CDateTime and assigns it to the ModifiedTime field.
func (o *Customer) SetModifiedTime(v A2CDateTime) {
	o.ModifiedTime.Set(&v)
}
// SetModifiedTimeNil sets the value for ModifiedTime to be an explicit nil
func (o *Customer) SetModifiedTimeNil() {
	o.ModifiedTime.Set(nil)
}

// UnsetModifiedTime ensures that no value is present for ModifiedTime, not even an explicit nil
func (o *Customer) UnsetModifiedTime() {
	o.ModifiedTime.Unset()
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

// GetLogin returns the Login field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetLogin() string {
	if o == nil || IsNil(o.Login.Get()) {
		var ret string
		return ret
	}
	return *o.Login.Get()
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Login.Get(), o.Login.IsSet()
}

// HasLogin returns a boolean if a field has been set.
func (o *Customer) HasLogin() bool {
	if o != nil && o.Login.IsSet() {
		return true
	}

	return false
}

// SetLogin gets a reference to the given NullableString and assigns it to the Login field.
func (o *Customer) SetLogin(v string) {
	o.Login.Set(&v)
}
// SetLoginNil sets the value for Login to be an explicit nil
func (o *Customer) SetLoginNil() {
	o.Login.Set(nil)
}

// UnsetLogin ensures that no value is present for Login, not even an explicit nil
func (o *Customer) UnsetLogin() {
	o.Login.Unset()
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetLastLogin() A2CDateTime {
	if o == nil || IsNil(o.LastLogin.Get()) {
		var ret A2CDateTime
		return ret
	}
	return *o.LastLogin.Get()
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetLastLoginOk() (*A2CDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLogin.Get(), o.LastLogin.IsSet()
}

// HasLastLogin returns a boolean if a field has been set.
func (o *Customer) HasLastLogin() bool {
	if o != nil && o.LastLogin.IsSet() {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given NullableA2CDateTime and assigns it to the LastLogin field.
func (o *Customer) SetLastLogin(v A2CDateTime) {
	o.LastLogin.Set(&v)
}
// SetLastLoginNil sets the value for LastLogin to be an explicit nil
func (o *Customer) SetLastLoginNil() {
	o.LastLogin.Set(nil)
}

// UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
func (o *Customer) UnsetLastLogin() {
	o.LastLogin.Unset()
}

// GetBirthDay returns the BirthDay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetBirthDay() A2CDateTime {
	if o == nil || IsNil(o.BirthDay.Get()) {
		var ret A2CDateTime
		return ret
	}
	return *o.BirthDay.Get()
}

// GetBirthDayOk returns a tuple with the BirthDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetBirthDayOk() (*A2CDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthDay.Get(), o.BirthDay.IsSet()
}

// HasBirthDay returns a boolean if a field has been set.
func (o *Customer) HasBirthDay() bool {
	if o != nil && o.BirthDay.IsSet() {
		return true
	}

	return false
}

// SetBirthDay gets a reference to the given NullableA2CDateTime and assigns it to the BirthDay field.
func (o *Customer) SetBirthDay(v A2CDateTime) {
	o.BirthDay.Set(&v)
}
// SetBirthDayNil sets the value for BirthDay to be an explicit nil
func (o *Customer) SetBirthDayNil() {
	o.BirthDay.Set(nil)
}

// UnsetBirthDay ensures that no value is present for BirthDay, not even an explicit nil
func (o *Customer) UnsetBirthDay() {
	o.BirthDay.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Customer) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *Customer) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Customer) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Customer) UnsetStatus() {
	o.Status.Unset()
}

// GetIsGuest returns the IsGuest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetIsGuest() bool {
	if o == nil || IsNil(o.IsGuest.Get()) {
		var ret bool
		return ret
	}
	return *o.IsGuest.Get()
}

// GetIsGuestOk returns a tuple with the IsGuest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetIsGuestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsGuest.Get(), o.IsGuest.IsSet()
}

// HasIsGuest returns a boolean if a field has been set.
func (o *Customer) HasIsGuest() bool {
	if o != nil && o.IsGuest.IsSet() {
		return true
	}

	return false
}

// SetIsGuest gets a reference to the given NullableBool and assigns it to the IsGuest field.
func (o *Customer) SetIsGuest(v bool) {
	o.IsGuest.Set(&v)
}
// SetIsGuestNil sets the value for IsGuest to be an explicit nil
func (o *Customer) SetIsGuestNil() {
	o.IsGuest.Set(nil)
}

// UnsetIsGuest ensures that no value is present for IsGuest, not even an explicit nil
func (o *Customer) UnsetIsGuest() {
	o.IsGuest.Unset()
}

// GetNewsLetterSubscription returns the NewsLetterSubscription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetNewsLetterSubscription() bool {
	if o == nil || IsNil(o.NewsLetterSubscription.Get()) {
		var ret bool
		return ret
	}
	return *o.NewsLetterSubscription.Get()
}

// GetNewsLetterSubscriptionOk returns a tuple with the NewsLetterSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetNewsLetterSubscriptionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewsLetterSubscription.Get(), o.NewsLetterSubscription.IsSet()
}

// HasNewsLetterSubscription returns a boolean if a field has been set.
func (o *Customer) HasNewsLetterSubscription() bool {
	if o != nil && o.NewsLetterSubscription.IsSet() {
		return true
	}

	return false
}

// SetNewsLetterSubscription gets a reference to the given NullableBool and assigns it to the NewsLetterSubscription field.
func (o *Customer) SetNewsLetterSubscription(v bool) {
	o.NewsLetterSubscription.Set(&v)
}
// SetNewsLetterSubscriptionNil sets the value for NewsLetterSubscription to be an explicit nil
func (o *Customer) SetNewsLetterSubscriptionNil() {
	o.NewsLetterSubscription.Set(nil)
}

// UnsetNewsLetterSubscription ensures that no value is present for NewsLetterSubscription, not even an explicit nil
func (o *Customer) UnsetNewsLetterSubscription() {
	o.NewsLetterSubscription.Unset()
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

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetGender() string {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret string
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetGenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *Customer) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableString and assigns it to the Gender field.
func (o *Customer) SetGender(v string) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *Customer) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *Customer) UnsetGender() {
	o.Gender.Unset()
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

// GetWebsite returns the Website field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetWebsite() string {
	if o == nil || IsNil(o.Website.Get()) {
		var ret string
		return ret
	}
	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetWebsiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// HasWebsite returns a boolean if a field has been set.
func (o *Customer) HasWebsite() bool {
	if o != nil && o.Website.IsSet() {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given NullableString and assigns it to the Website field.
func (o *Customer) SetWebsite(v string) {
	o.Website.Set(&v)
}
// SetWebsiteNil sets the value for Website to be an explicit nil
func (o *Customer) SetWebsiteNil() {
	o.Website.Set(nil)
}

// UnsetWebsite ensures that no value is present for Website, not even an explicit nil
func (o *Customer) UnsetWebsite() {
	o.Website.Unset()
}

// GetFax returns the Fax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetFax() string {
	if o == nil || IsNil(o.Fax.Get()) {
		var ret string
		return ret
	}
	return *o.Fax.Get()
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetFaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fax.Get(), o.Fax.IsSet()
}

// HasFax returns a boolean if a field has been set.
func (o *Customer) HasFax() bool {
	if o != nil && o.Fax.IsSet() {
		return true
	}

	return false
}

// SetFax gets a reference to the given NullableString and assigns it to the Fax field.
func (o *Customer) SetFax(v string) {
	o.Fax.Set(&v)
}
// SetFaxNil sets the value for Fax to be an explicit nil
func (o *Customer) SetFaxNil() {
	o.Fax.Set(nil)
}

// UnsetFax ensures that no value is present for Fax, not even an explicit nil
func (o *Customer) UnsetFax() {
	o.Fax.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetCompany() string {
	if o == nil || IsNil(o.Company.Get()) {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *Customer) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *Customer) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *Customer) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *Customer) UnsetCompany() {
	o.Company.Unset()
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *Customer) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *Customer) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}
// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *Customer) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *Customer) UnsetIpAddress() {
	o.IpAddress.Unset()
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

// GetLangId returns the LangId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetLangId() string {
	if o == nil || IsNil(o.LangId.Get()) {
		var ret string
		return ret
	}
	return *o.LangId.Get()
}

// GetLangIdOk returns a tuple with the LangId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetLangIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LangId.Get(), o.LangId.IsSet()
}

// HasLangId returns a boolean if a field has been set.
func (o *Customer) HasLangId() bool {
	if o != nil && o.LangId.IsSet() {
		return true
	}

	return false
}

// SetLangId gets a reference to the given NullableString and assigns it to the LangId field.
func (o *Customer) SetLangId(v string) {
	o.LangId.Set(&v)
}
// SetLangIdNil sets the value for LangId to be an explicit nil
func (o *Customer) SetLangIdNil() {
	o.LangId.Set(nil)
}

// UnsetLangId ensures that no value is present for LangId, not even an explicit nil
func (o *Customer) UnsetLangId() {
	o.LangId.Unset()
}

// GetOrdersCount returns the OrdersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetOrdersCount() int32 {
	if o == nil || IsNil(o.OrdersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.OrdersCount.Get()
}

// GetOrdersCountOk returns a tuple with the OrdersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetOrdersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrdersCount.Get(), o.OrdersCount.IsSet()
}

// HasOrdersCount returns a boolean if a field has been set.
func (o *Customer) HasOrdersCount() bool {
	if o != nil && o.OrdersCount.IsSet() {
		return true
	}

	return false
}

// SetOrdersCount gets a reference to the given NullableInt32 and assigns it to the OrdersCount field.
func (o *Customer) SetOrdersCount(v int32) {
	o.OrdersCount.Set(&v)
}
// SetOrdersCountNil sets the value for OrdersCount to be an explicit nil
func (o *Customer) SetOrdersCountNil() {
	o.OrdersCount.Set(nil)
}

// UnsetOrdersCount ensures that no value is present for OrdersCount, not even an explicit nil
func (o *Customer) UnsetOrdersCount() {
	o.OrdersCount.Unset()
}

// GetLastOrderId returns the LastOrderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetLastOrderId() string {
	if o == nil || IsNil(o.LastOrderId.Get()) {
		var ret string
		return ret
	}
	return *o.LastOrderId.Get()
}

// GetLastOrderIdOk returns a tuple with the LastOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetLastOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastOrderId.Get(), o.LastOrderId.IsSet()
}

// HasLastOrderId returns a boolean if a field has been set.
func (o *Customer) HasLastOrderId() bool {
	if o != nil && o.LastOrderId.IsSet() {
		return true
	}

	return false
}

// SetLastOrderId gets a reference to the given NullableString and assigns it to the LastOrderId field.
func (o *Customer) SetLastOrderId(v string) {
	o.LastOrderId.Set(&v)
}
// SetLastOrderIdNil sets the value for LastOrderId to be an explicit nil
func (o *Customer) SetLastOrderIdNil() {
	o.LastOrderId.Set(nil)
}

// UnsetLastOrderId ensures that no value is present for LastOrderId, not even an explicit nil
func (o *Customer) UnsetLastOrderId() {
	o.LastOrderId.Unset()
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetAdditionalFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
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

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
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
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.CreatedTime.IsSet() {
		toSerialize["created_time"] = o.CreatedTime.Get()
	}
	if o.ModifiedTime.IsSet() {
		toSerialize["modified_time"] = o.ModifiedTime.Get()
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if o.Login.IsSet() {
		toSerialize["login"] = o.Login.Get()
	}
	if o.LastLogin.IsSet() {
		toSerialize["last_login"] = o.LastLogin.Get()
	}
	if o.BirthDay.IsSet() {
		toSerialize["birth_day"] = o.BirthDay.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.IsGuest.IsSet() {
		toSerialize["is_guest"] = o.IsGuest.Get()
	}
	if o.NewsLetterSubscription.IsSet() {
		toSerialize["news_letter_subscription"] = o.NewsLetterSubscription.Get()
	}
	if !IsNil(o.Consents) {
		toSerialize["consents"] = o.Consents
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	if !IsNil(o.StoresIds) {
		toSerialize["stores_ids"] = o.StoresIds
	}
	if o.Website.IsSet() {
		toSerialize["website"] = o.Website.Get()
	}
	if o.Fax.IsSet() {
		toSerialize["fax"] = o.Fax.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.IpAddress.IsSet() {
		toSerialize["ip_address"] = o.IpAddress.Get()
	}
	if !IsNil(o.AddressBook) {
		toSerialize["address_book"] = o.AddressBook
	}
	if o.LangId.IsSet() {
		toSerialize["lang_id"] = o.LangId.Get()
	}
	if o.OrdersCount.IsSet() {
		toSerialize["orders_count"] = o.OrdersCount.Get()
	}
	if o.LastOrderId.IsSet() {
		toSerialize["last_order_id"] = o.LastOrderId.Get()
	}
	if o.AdditionalFields != nil {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if o.CustomFields != nil {
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


