# CustomerAddressAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | Defines customer id | 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**FirstName** | Pointer to **string** | Defines customer&#39;s address first name | [optional] 
**LastName** | Pointer to **string** | Defines customer&#39;s address last name | [optional] 
**Company** | Pointer to **string** | Defines customer&#39;s address company | [optional] 
**Address1** | **string** | Specifies customer&#39;s address address1 | 
**Address2** | Pointer to **string** | Specifies customer&#39;s address address2 | [optional] 
**City** | **string** | Specifies customer&#39;s address city | 
**Country** | **string** | Specifies customer&#39;s address ISO code or name of country | 
**State** | Pointer to **string** | Specifies customer&#39;s address ISO code or name of state | [optional] 
**Postcode** | **string** | Specifies customer&#39;s address postcode | 
**IdentificationNumber** | Pointer to **string** | Specifies the national ID card number of this person, or a unique tax identification number for customer&#39;s address | [optional] 
**Types** | Pointer to **[]string** | Specifies customer&#39;s address types | [optional] 
**Default** | Pointer to **bool** | Specifies whether the customer&#39;s address is used by default | [optional] 
**Phone** | Pointer to **string** | Defines customer&#39;s address phone number | [optional] 
**PhoneMobile** | Pointer to **string** | Defines customer&#39;s address mobile phone number | [optional] 
**Fax** | Pointer to **string** | Defines customer&#39;s address fax | [optional] 
**Website** | Pointer to **string** | Defines Link to customer&#39;s address website | [optional] 
**Gender** | Pointer to **string** | Defines customer&#39;s address gender | [optional] 
**TaxId** | Pointer to **string** | Add Tax Id | [optional] 
**Alias** | Pointer to **string** | Specifies customer&#39;s alias in the address book | [optional] 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 

## Methods

### NewCustomerAddressAdd

`func NewCustomerAddressAdd(customerId string, address1 string, city string, country string, postcode string, ) *CustomerAddressAdd`

NewCustomerAddressAdd instantiates a new CustomerAddressAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddressAddWithDefaults

`func NewCustomerAddressAddWithDefaults() *CustomerAddressAdd`

NewCustomerAddressAddWithDefaults instantiates a new CustomerAddressAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CustomerAddressAdd) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerAddressAdd) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerAddressAdd) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetStoreId

`func (o *CustomerAddressAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CustomerAddressAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CustomerAddressAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *CustomerAddressAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetFirstName

`func (o *CustomerAddressAdd) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerAddressAdd) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerAddressAdd) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CustomerAddressAdd) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CustomerAddressAdd) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerAddressAdd) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerAddressAdd) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CustomerAddressAdd) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompany

`func (o *CustomerAddressAdd) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CustomerAddressAdd) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CustomerAddressAdd) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CustomerAddressAdd) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAddress1

`func (o *CustomerAddressAdd) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *CustomerAddressAdd) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *CustomerAddressAdd) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *CustomerAddressAdd) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *CustomerAddressAdd) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *CustomerAddressAdd) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *CustomerAddressAdd) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *CustomerAddressAdd) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CustomerAddressAdd) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CustomerAddressAdd) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *CustomerAddressAdd) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerAddressAdd) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerAddressAdd) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetState

`func (o *CustomerAddressAdd) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerAddressAdd) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerAddressAdd) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerAddressAdd) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostcode

`func (o *CustomerAddressAdd) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *CustomerAddressAdd) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *CustomerAddressAdd) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### GetIdentificationNumber

`func (o *CustomerAddressAdd) GetIdentificationNumber() string`

GetIdentificationNumber returns the IdentificationNumber field if non-nil, zero value otherwise.

### GetIdentificationNumberOk

`func (o *CustomerAddressAdd) GetIdentificationNumberOk() (*string, bool)`

GetIdentificationNumberOk returns a tuple with the IdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationNumber

`func (o *CustomerAddressAdd) SetIdentificationNumber(v string)`

SetIdentificationNumber sets IdentificationNumber field to given value.

### HasIdentificationNumber

`func (o *CustomerAddressAdd) HasIdentificationNumber() bool`

HasIdentificationNumber returns a boolean if a field has been set.

### GetTypes

`func (o *CustomerAddressAdd) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *CustomerAddressAdd) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *CustomerAddressAdd) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *CustomerAddressAdd) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetDefault

`func (o *CustomerAddressAdd) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomerAddressAdd) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomerAddressAdd) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomerAddressAdd) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerAddressAdd) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerAddressAdd) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerAddressAdd) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerAddressAdd) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPhoneMobile

`func (o *CustomerAddressAdd) GetPhoneMobile() string`

GetPhoneMobile returns the PhoneMobile field if non-nil, zero value otherwise.

### GetPhoneMobileOk

`func (o *CustomerAddressAdd) GetPhoneMobileOk() (*string, bool)`

GetPhoneMobileOk returns a tuple with the PhoneMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneMobile

`func (o *CustomerAddressAdd) SetPhoneMobile(v string)`

SetPhoneMobile sets PhoneMobile field to given value.

### HasPhoneMobile

`func (o *CustomerAddressAdd) HasPhoneMobile() bool`

HasPhoneMobile returns a boolean if a field has been set.

### GetFax

`func (o *CustomerAddressAdd) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *CustomerAddressAdd) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *CustomerAddressAdd) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *CustomerAddressAdd) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetWebsite

`func (o *CustomerAddressAdd) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CustomerAddressAdd) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CustomerAddressAdd) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CustomerAddressAdd) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetGender

`func (o *CustomerAddressAdd) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CustomerAddressAdd) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CustomerAddressAdd) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CustomerAddressAdd) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetTaxId

`func (o *CustomerAddressAdd) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CustomerAddressAdd) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CustomerAddressAdd) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CustomerAddressAdd) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetAlias

`func (o *CustomerAddressAdd) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CustomerAddressAdd) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CustomerAddressAdd) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *CustomerAddressAdd) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *CustomerAddressAdd) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *CustomerAddressAdd) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *CustomerAddressAdd) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *CustomerAddressAdd) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


