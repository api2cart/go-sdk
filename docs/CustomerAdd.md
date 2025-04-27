# CustomerAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Defines customer&#39;s email | 
**FirstName** | Pointer to **string** | Defines customer&#39;s first name | [optional] 
**LastName** | Pointer to **string** | Defines customer&#39;s last name | [optional] 
**Password** | Pointer to **string** | Defines customer&#39;s unique password | [optional] 
**Group** | Pointer to **string** | Defines the group where the customer | [optional] 
**GroupIds** | Pointer to **string** | Groups that will be assigned to a customer | [optional] 
**Status** | Pointer to **string** | Defines customer&#39;s status | [optional] [default to "enabled"]
**CreatedTime** | Pointer to **string** | Entity&#39;s date creation | [optional] 
**ModifiedTime** | Pointer to **string** | Entity&#39;s date modification | [optional] 
**Login** | Pointer to **string** | Specifies customer&#39;s login name | [optional] 
**LastLogin** | Pointer to **string** | Defines customer&#39;s last login time | [optional] 
**BirthDay** | Pointer to **string** | Defines customer&#39;s birthday | [optional] 
**NewsLetterSubscription** | Pointer to **bool** | Defines whether the newsletter subscription is available for the user | [optional] 
**Consents** | Pointer to [**[]CustomerAddConsentsInner**](CustomerAddConsentsInner.md) | Defines consents to notifications | [optional] 
**Gender** | Pointer to **string** | Defines customer&#39;s gender | [optional] 
**Website** | Pointer to **string** | Link to customer website | [optional] 
**Fax** | Pointer to **string** | Defines customer&#39;s fax | [optional] 
**Company** | Pointer to **string** | Defines customer&#39;s company | [optional] 
**Phone** | Pointer to **string** | Defines customer&#39;s phone number | [optional] 
**Note** | Pointer to **string** | The customer note. | [optional] 
**Country** | Pointer to **string** | Specifies ISO code or name of country | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**Address** | Pointer to [**[]CustomerAddAddressInner**](CustomerAddAddressInner.md) |  | [optional] 

## Methods

### NewCustomerAdd

`func NewCustomerAdd(email string, ) *CustomerAdd`

NewCustomerAdd instantiates a new CustomerAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddWithDefaults

`func NewCustomerAddWithDefaults() *CustomerAdd`

NewCustomerAddWithDefaults instantiates a new CustomerAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CustomerAdd) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerAdd) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerAdd) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *CustomerAdd) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerAdd) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerAdd) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CustomerAdd) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CustomerAdd) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerAdd) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerAdd) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CustomerAdd) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPassword

`func (o *CustomerAdd) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CustomerAdd) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CustomerAdd) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CustomerAdd) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetGroup

`func (o *CustomerAdd) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CustomerAdd) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CustomerAdd) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CustomerAdd) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupIds

`func (o *CustomerAdd) GetGroupIds() string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *CustomerAdd) GetGroupIdsOk() (*string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *CustomerAdd) SetGroupIds(v string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *CustomerAdd) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerAdd) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerAdd) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerAdd) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerAdd) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CustomerAdd) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CustomerAdd) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CustomerAdd) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CustomerAdd) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetModifiedTime

`func (o *CustomerAdd) GetModifiedTime() string`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *CustomerAdd) GetModifiedTimeOk() (*string, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *CustomerAdd) SetModifiedTime(v string)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *CustomerAdd) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetLogin

`func (o *CustomerAdd) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CustomerAdd) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *CustomerAdd) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *CustomerAdd) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetLastLogin

`func (o *CustomerAdd) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *CustomerAdd) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *CustomerAdd) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *CustomerAdd) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetBirthDay

`func (o *CustomerAdd) GetBirthDay() string`

GetBirthDay returns the BirthDay field if non-nil, zero value otherwise.

### GetBirthDayOk

`func (o *CustomerAdd) GetBirthDayOk() (*string, bool)`

GetBirthDayOk returns a tuple with the BirthDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDay

`func (o *CustomerAdd) SetBirthDay(v string)`

SetBirthDay sets BirthDay field to given value.

### HasBirthDay

`func (o *CustomerAdd) HasBirthDay() bool`

HasBirthDay returns a boolean if a field has been set.

### GetNewsLetterSubscription

`func (o *CustomerAdd) GetNewsLetterSubscription() bool`

GetNewsLetterSubscription returns the NewsLetterSubscription field if non-nil, zero value otherwise.

### GetNewsLetterSubscriptionOk

`func (o *CustomerAdd) GetNewsLetterSubscriptionOk() (*bool, bool)`

GetNewsLetterSubscriptionOk returns a tuple with the NewsLetterSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsLetterSubscription

`func (o *CustomerAdd) SetNewsLetterSubscription(v bool)`

SetNewsLetterSubscription sets NewsLetterSubscription field to given value.

### HasNewsLetterSubscription

`func (o *CustomerAdd) HasNewsLetterSubscription() bool`

HasNewsLetterSubscription returns a boolean if a field has been set.

### GetConsents

`func (o *CustomerAdd) GetConsents() []CustomerAddConsentsInner`

GetConsents returns the Consents field if non-nil, zero value otherwise.

### GetConsentsOk

`func (o *CustomerAdd) GetConsentsOk() (*[]CustomerAddConsentsInner, bool)`

GetConsentsOk returns a tuple with the Consents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsents

`func (o *CustomerAdd) SetConsents(v []CustomerAddConsentsInner)`

SetConsents sets Consents field to given value.

### HasConsents

`func (o *CustomerAdd) HasConsents() bool`

HasConsents returns a boolean if a field has been set.

### GetGender

`func (o *CustomerAdd) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CustomerAdd) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CustomerAdd) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CustomerAdd) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetWebsite

`func (o *CustomerAdd) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CustomerAdd) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CustomerAdd) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CustomerAdd) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetFax

`func (o *CustomerAdd) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *CustomerAdd) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *CustomerAdd) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *CustomerAdd) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetCompany

`func (o *CustomerAdd) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CustomerAdd) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CustomerAdd) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CustomerAdd) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerAdd) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerAdd) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerAdd) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerAdd) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetNote

`func (o *CustomerAdd) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CustomerAdd) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CustomerAdd) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CustomerAdd) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetCountry

`func (o *CustomerAdd) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerAdd) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerAdd) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CustomerAdd) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetStoreId

`func (o *CustomerAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CustomerAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CustomerAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *CustomerAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetAddress

`func (o *CustomerAdd) GetAddress() []CustomerAddAddressInner`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerAdd) GetAddressOk() (*[]CustomerAddAddressInner, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerAdd) SetAddress(v []CustomerAddAddressInner)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustomerAdd) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


