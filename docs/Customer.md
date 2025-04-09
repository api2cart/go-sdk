# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedTime** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**Group** | Pointer to [**[]CustomerGroup**](CustomerGroup.md) |  | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**LastLogin** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**BirthDay** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**NewsLetterSubscription** | Pointer to **bool** |  | [optional] 
**Consents** | Pointer to [**[]CustomerConsent**](CustomerConsent.md) |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**StoresIds** | Pointer to **[]string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Fax** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**AddressBook** | Pointer to [**[]CustomerAddress**](CustomerAddress.md) |  | [optional] 
**LangId** | Pointer to **string** |  | [optional] 
**OrdersCount** | Pointer to **int32** |  | [optional] 
**LastOrderId** | Pointer to **string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCustomer

`func NewCustomer() *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Customer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Customer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Customer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Customer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *Customer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Customer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Customer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Customer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *Customer) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Customer) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Customer) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Customer) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Customer) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Customer) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Customer) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Customer) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhone

`func (o *Customer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Customer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Customer) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Customer) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Customer) GetCreatedTime() A2CDateTime`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Customer) GetCreatedTimeOk() (*A2CDateTime, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Customer) SetCreatedTime(v A2CDateTime)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Customer) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Customer) GetModifiedTime() A2CDateTime`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Customer) GetModifiedTimeOk() (*A2CDateTime, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Customer) SetModifiedTime(v A2CDateTime)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Customer) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetGroup

`func (o *Customer) GetGroup() []CustomerGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Customer) GetGroupOk() (*[]CustomerGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Customer) SetGroup(v []CustomerGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Customer) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetLogin

`func (o *Customer) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *Customer) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *Customer) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *Customer) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetLastLogin

`func (o *Customer) GetLastLogin() A2CDateTime`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *Customer) GetLastLoginOk() (*A2CDateTime, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *Customer) SetLastLogin(v A2CDateTime)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *Customer) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetBirthDay

`func (o *Customer) GetBirthDay() A2CDateTime`

GetBirthDay returns the BirthDay field if non-nil, zero value otherwise.

### GetBirthDayOk

`func (o *Customer) GetBirthDayOk() (*A2CDateTime, bool)`

GetBirthDayOk returns a tuple with the BirthDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDay

`func (o *Customer) SetBirthDay(v A2CDateTime)`

SetBirthDay sets BirthDay field to given value.

### HasBirthDay

`func (o *Customer) HasBirthDay() bool`

HasBirthDay returns a boolean if a field has been set.

### GetStatus

`func (o *Customer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Customer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Customer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Customer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNewsLetterSubscription

`func (o *Customer) GetNewsLetterSubscription() bool`

GetNewsLetterSubscription returns the NewsLetterSubscription field if non-nil, zero value otherwise.

### GetNewsLetterSubscriptionOk

`func (o *Customer) GetNewsLetterSubscriptionOk() (*bool, bool)`

GetNewsLetterSubscriptionOk returns a tuple with the NewsLetterSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsLetterSubscription

`func (o *Customer) SetNewsLetterSubscription(v bool)`

SetNewsLetterSubscription sets NewsLetterSubscription field to given value.

### HasNewsLetterSubscription

`func (o *Customer) HasNewsLetterSubscription() bool`

HasNewsLetterSubscription returns a boolean if a field has been set.

### GetConsents

`func (o *Customer) GetConsents() []CustomerConsent`

GetConsents returns the Consents field if non-nil, zero value otherwise.

### GetConsentsOk

`func (o *Customer) GetConsentsOk() (*[]CustomerConsent, bool)`

GetConsentsOk returns a tuple with the Consents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsents

`func (o *Customer) SetConsents(v []CustomerConsent)`

SetConsents sets Consents field to given value.

### HasConsents

`func (o *Customer) HasConsents() bool`

HasConsents returns a boolean if a field has been set.

### GetGender

`func (o *Customer) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Customer) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Customer) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Customer) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetStoresIds

`func (o *Customer) GetStoresIds() []string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *Customer) GetStoresIdsOk() (*[]string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *Customer) SetStoresIds(v []string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *Customer) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetWebsite

`func (o *Customer) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Customer) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Customer) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Customer) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetFax

`func (o *Customer) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *Customer) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *Customer) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *Customer) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetCompany

`func (o *Customer) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Customer) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Customer) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Customer) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetIpAddress

`func (o *Customer) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Customer) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Customer) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Customer) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetAddressBook

`func (o *Customer) GetAddressBook() []CustomerAddress`

GetAddressBook returns the AddressBook field if non-nil, zero value otherwise.

### GetAddressBookOk

`func (o *Customer) GetAddressBookOk() (*[]CustomerAddress, bool)`

GetAddressBookOk returns a tuple with the AddressBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBook

`func (o *Customer) SetAddressBook(v []CustomerAddress)`

SetAddressBook sets AddressBook field to given value.

### HasAddressBook

`func (o *Customer) HasAddressBook() bool`

HasAddressBook returns a boolean if a field has been set.

### GetLangId

`func (o *Customer) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *Customer) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *Customer) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *Customer) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetOrdersCount

`func (o *Customer) GetOrdersCount() int32`

GetOrdersCount returns the OrdersCount field if non-nil, zero value otherwise.

### GetOrdersCountOk

`func (o *Customer) GetOrdersCountOk() (*int32, bool)`

GetOrdersCountOk returns a tuple with the OrdersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersCount

`func (o *Customer) SetOrdersCount(v int32)`

SetOrdersCount sets OrdersCount field to given value.

### HasOrdersCount

`func (o *Customer) HasOrdersCount() bool`

HasOrdersCount returns a boolean if a field has been set.

### GetLastOrderId

`func (o *Customer) GetLastOrderId() string`

GetLastOrderId returns the LastOrderId field if non-nil, zero value otherwise.

### GetLastOrderIdOk

`func (o *Customer) GetLastOrderIdOk() (*string, bool)`

GetLastOrderIdOk returns a tuple with the LastOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOrderId

`func (o *Customer) SetLastOrderId(v string)`

SetLastOrderId sets LastOrderId field to given value.

### HasLastOrderId

`func (o *Customer) HasLastOrderId() bool`

HasLastOrderId returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Customer) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Customer) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Customer) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Customer) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *Customer) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Customer) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Customer) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Customer) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


