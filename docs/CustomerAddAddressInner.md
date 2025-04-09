# CustomerAddAddressInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressBookType** | Pointer to **string** | Specifies customer&#39;s address type | [optional] 
**AddressBookFirstName** | Pointer to **string** | Specifies customer&#39;s first name in the address book | [optional] 
**AddressBookLastName** | Pointer to **string** | Specifies customer&#39;s last name in the address book | [optional] 
**AddressBookCompany** | Pointer to **string** | Specifies customer&#39;s company name in the address book | [optional] 
**AddressBookFax** | Pointer to **string** | Specifies customer&#39;s fax in the address book | [optional] 
**AddressBookPhone** | Pointer to **string** | Specifies customer&#39;s phone number in the address book | [optional] 
**AddressBookPhoneMobile** | Pointer to **string** | Specifies customer&#39;s mobile phone number in the address book | [optional] 
**AddressBookWebsite** | Pointer to **string** | Specifies customer&#39;s website in the address book | [optional] 
**AddressBookAddress1** | Pointer to **string** | Specifies customer&#39;s first address in the address book | [optional] 
**AddressBookAddress2** | Pointer to **string** | Specifies customer&#39;s second address in the address book | [optional] 
**AddressBookCity** | Pointer to **string** | Specifies customer&#39;s city in the address book | [optional] 
**AddressBookCountry** | Pointer to **string** | ISO code or name of country | [optional] 
**AddressBookState** | Pointer to **string** | ISO code or name of state. | [optional] 
**AddressBookPostcode** | Pointer to **string** | Specifies customer&#39;s postcode | [optional] 
**AddressBookGender** | Pointer to **string** | Specifies customer&#39;s gender | [optional] 
**AddressBookRegion** | Pointer to **string** | Specifies customer&#39;s region | [optional] 
**AddressBookDefault** | Pointer to **bool** | Defines whether the address is used by default | [optional] 
**AddressBookTaxId** | Pointer to **string** | Add Tax Id | [optional] 
**AddressBookIdentificationNumber** | Pointer to **string** | The national ID card number of this person, or a unique tax identification number. | [optional] 
**AddressBookAlias** | Pointer to **string** | Specifies customer&#39;s alias in the address book | [optional] 

## Methods

### NewCustomerAddAddressInner

`func NewCustomerAddAddressInner() *CustomerAddAddressInner`

NewCustomerAddAddressInner instantiates a new CustomerAddAddressInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddAddressInnerWithDefaults

`func NewCustomerAddAddressInnerWithDefaults() *CustomerAddAddressInner`

NewCustomerAddAddressInnerWithDefaults instantiates a new CustomerAddAddressInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressBookType

`func (o *CustomerAddAddressInner) GetAddressBookType() string`

GetAddressBookType returns the AddressBookType field if non-nil, zero value otherwise.

### GetAddressBookTypeOk

`func (o *CustomerAddAddressInner) GetAddressBookTypeOk() (*string, bool)`

GetAddressBookTypeOk returns a tuple with the AddressBookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookType

`func (o *CustomerAddAddressInner) SetAddressBookType(v string)`

SetAddressBookType sets AddressBookType field to given value.

### HasAddressBookType

`func (o *CustomerAddAddressInner) HasAddressBookType() bool`

HasAddressBookType returns a boolean if a field has been set.

### GetAddressBookFirstName

`func (o *CustomerAddAddressInner) GetAddressBookFirstName() string`

GetAddressBookFirstName returns the AddressBookFirstName field if non-nil, zero value otherwise.

### GetAddressBookFirstNameOk

`func (o *CustomerAddAddressInner) GetAddressBookFirstNameOk() (*string, bool)`

GetAddressBookFirstNameOk returns a tuple with the AddressBookFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookFirstName

`func (o *CustomerAddAddressInner) SetAddressBookFirstName(v string)`

SetAddressBookFirstName sets AddressBookFirstName field to given value.

### HasAddressBookFirstName

`func (o *CustomerAddAddressInner) HasAddressBookFirstName() bool`

HasAddressBookFirstName returns a boolean if a field has been set.

### GetAddressBookLastName

`func (o *CustomerAddAddressInner) GetAddressBookLastName() string`

GetAddressBookLastName returns the AddressBookLastName field if non-nil, zero value otherwise.

### GetAddressBookLastNameOk

`func (o *CustomerAddAddressInner) GetAddressBookLastNameOk() (*string, bool)`

GetAddressBookLastNameOk returns a tuple with the AddressBookLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookLastName

`func (o *CustomerAddAddressInner) SetAddressBookLastName(v string)`

SetAddressBookLastName sets AddressBookLastName field to given value.

### HasAddressBookLastName

`func (o *CustomerAddAddressInner) HasAddressBookLastName() bool`

HasAddressBookLastName returns a boolean if a field has been set.

### GetAddressBookCompany

`func (o *CustomerAddAddressInner) GetAddressBookCompany() string`

GetAddressBookCompany returns the AddressBookCompany field if non-nil, zero value otherwise.

### GetAddressBookCompanyOk

`func (o *CustomerAddAddressInner) GetAddressBookCompanyOk() (*string, bool)`

GetAddressBookCompanyOk returns a tuple with the AddressBookCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookCompany

`func (o *CustomerAddAddressInner) SetAddressBookCompany(v string)`

SetAddressBookCompany sets AddressBookCompany field to given value.

### HasAddressBookCompany

`func (o *CustomerAddAddressInner) HasAddressBookCompany() bool`

HasAddressBookCompany returns a boolean if a field has been set.

### GetAddressBookFax

`func (o *CustomerAddAddressInner) GetAddressBookFax() string`

GetAddressBookFax returns the AddressBookFax field if non-nil, zero value otherwise.

### GetAddressBookFaxOk

`func (o *CustomerAddAddressInner) GetAddressBookFaxOk() (*string, bool)`

GetAddressBookFaxOk returns a tuple with the AddressBookFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookFax

`func (o *CustomerAddAddressInner) SetAddressBookFax(v string)`

SetAddressBookFax sets AddressBookFax field to given value.

### HasAddressBookFax

`func (o *CustomerAddAddressInner) HasAddressBookFax() bool`

HasAddressBookFax returns a boolean if a field has been set.

### GetAddressBookPhone

`func (o *CustomerAddAddressInner) GetAddressBookPhone() string`

GetAddressBookPhone returns the AddressBookPhone field if non-nil, zero value otherwise.

### GetAddressBookPhoneOk

`func (o *CustomerAddAddressInner) GetAddressBookPhoneOk() (*string, bool)`

GetAddressBookPhoneOk returns a tuple with the AddressBookPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookPhone

`func (o *CustomerAddAddressInner) SetAddressBookPhone(v string)`

SetAddressBookPhone sets AddressBookPhone field to given value.

### HasAddressBookPhone

`func (o *CustomerAddAddressInner) HasAddressBookPhone() bool`

HasAddressBookPhone returns a boolean if a field has been set.

### GetAddressBookPhoneMobile

`func (o *CustomerAddAddressInner) GetAddressBookPhoneMobile() string`

GetAddressBookPhoneMobile returns the AddressBookPhoneMobile field if non-nil, zero value otherwise.

### GetAddressBookPhoneMobileOk

`func (o *CustomerAddAddressInner) GetAddressBookPhoneMobileOk() (*string, bool)`

GetAddressBookPhoneMobileOk returns a tuple with the AddressBookPhoneMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookPhoneMobile

`func (o *CustomerAddAddressInner) SetAddressBookPhoneMobile(v string)`

SetAddressBookPhoneMobile sets AddressBookPhoneMobile field to given value.

### HasAddressBookPhoneMobile

`func (o *CustomerAddAddressInner) HasAddressBookPhoneMobile() bool`

HasAddressBookPhoneMobile returns a boolean if a field has been set.

### GetAddressBookWebsite

`func (o *CustomerAddAddressInner) GetAddressBookWebsite() string`

GetAddressBookWebsite returns the AddressBookWebsite field if non-nil, zero value otherwise.

### GetAddressBookWebsiteOk

`func (o *CustomerAddAddressInner) GetAddressBookWebsiteOk() (*string, bool)`

GetAddressBookWebsiteOk returns a tuple with the AddressBookWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookWebsite

`func (o *CustomerAddAddressInner) SetAddressBookWebsite(v string)`

SetAddressBookWebsite sets AddressBookWebsite field to given value.

### HasAddressBookWebsite

`func (o *CustomerAddAddressInner) HasAddressBookWebsite() bool`

HasAddressBookWebsite returns a boolean if a field has been set.

### GetAddressBookAddress1

`func (o *CustomerAddAddressInner) GetAddressBookAddress1() string`

GetAddressBookAddress1 returns the AddressBookAddress1 field if non-nil, zero value otherwise.

### GetAddressBookAddress1Ok

`func (o *CustomerAddAddressInner) GetAddressBookAddress1Ok() (*string, bool)`

GetAddressBookAddress1Ok returns a tuple with the AddressBookAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookAddress1

`func (o *CustomerAddAddressInner) SetAddressBookAddress1(v string)`

SetAddressBookAddress1 sets AddressBookAddress1 field to given value.

### HasAddressBookAddress1

`func (o *CustomerAddAddressInner) HasAddressBookAddress1() bool`

HasAddressBookAddress1 returns a boolean if a field has been set.

### GetAddressBookAddress2

`func (o *CustomerAddAddressInner) GetAddressBookAddress2() string`

GetAddressBookAddress2 returns the AddressBookAddress2 field if non-nil, zero value otherwise.

### GetAddressBookAddress2Ok

`func (o *CustomerAddAddressInner) GetAddressBookAddress2Ok() (*string, bool)`

GetAddressBookAddress2Ok returns a tuple with the AddressBookAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookAddress2

`func (o *CustomerAddAddressInner) SetAddressBookAddress2(v string)`

SetAddressBookAddress2 sets AddressBookAddress2 field to given value.

### HasAddressBookAddress2

`func (o *CustomerAddAddressInner) HasAddressBookAddress2() bool`

HasAddressBookAddress2 returns a boolean if a field has been set.

### GetAddressBookCity

`func (o *CustomerAddAddressInner) GetAddressBookCity() string`

GetAddressBookCity returns the AddressBookCity field if non-nil, zero value otherwise.

### GetAddressBookCityOk

`func (o *CustomerAddAddressInner) GetAddressBookCityOk() (*string, bool)`

GetAddressBookCityOk returns a tuple with the AddressBookCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookCity

`func (o *CustomerAddAddressInner) SetAddressBookCity(v string)`

SetAddressBookCity sets AddressBookCity field to given value.

### HasAddressBookCity

`func (o *CustomerAddAddressInner) HasAddressBookCity() bool`

HasAddressBookCity returns a boolean if a field has been set.

### GetAddressBookCountry

`func (o *CustomerAddAddressInner) GetAddressBookCountry() string`

GetAddressBookCountry returns the AddressBookCountry field if non-nil, zero value otherwise.

### GetAddressBookCountryOk

`func (o *CustomerAddAddressInner) GetAddressBookCountryOk() (*string, bool)`

GetAddressBookCountryOk returns a tuple with the AddressBookCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookCountry

`func (o *CustomerAddAddressInner) SetAddressBookCountry(v string)`

SetAddressBookCountry sets AddressBookCountry field to given value.

### HasAddressBookCountry

`func (o *CustomerAddAddressInner) HasAddressBookCountry() bool`

HasAddressBookCountry returns a boolean if a field has been set.

### GetAddressBookState

`func (o *CustomerAddAddressInner) GetAddressBookState() string`

GetAddressBookState returns the AddressBookState field if non-nil, zero value otherwise.

### GetAddressBookStateOk

`func (o *CustomerAddAddressInner) GetAddressBookStateOk() (*string, bool)`

GetAddressBookStateOk returns a tuple with the AddressBookState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookState

`func (o *CustomerAddAddressInner) SetAddressBookState(v string)`

SetAddressBookState sets AddressBookState field to given value.

### HasAddressBookState

`func (o *CustomerAddAddressInner) HasAddressBookState() bool`

HasAddressBookState returns a boolean if a field has been set.

### GetAddressBookPostcode

`func (o *CustomerAddAddressInner) GetAddressBookPostcode() string`

GetAddressBookPostcode returns the AddressBookPostcode field if non-nil, zero value otherwise.

### GetAddressBookPostcodeOk

`func (o *CustomerAddAddressInner) GetAddressBookPostcodeOk() (*string, bool)`

GetAddressBookPostcodeOk returns a tuple with the AddressBookPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookPostcode

`func (o *CustomerAddAddressInner) SetAddressBookPostcode(v string)`

SetAddressBookPostcode sets AddressBookPostcode field to given value.

### HasAddressBookPostcode

`func (o *CustomerAddAddressInner) HasAddressBookPostcode() bool`

HasAddressBookPostcode returns a boolean if a field has been set.

### GetAddressBookGender

`func (o *CustomerAddAddressInner) GetAddressBookGender() string`

GetAddressBookGender returns the AddressBookGender field if non-nil, zero value otherwise.

### GetAddressBookGenderOk

`func (o *CustomerAddAddressInner) GetAddressBookGenderOk() (*string, bool)`

GetAddressBookGenderOk returns a tuple with the AddressBookGender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookGender

`func (o *CustomerAddAddressInner) SetAddressBookGender(v string)`

SetAddressBookGender sets AddressBookGender field to given value.

### HasAddressBookGender

`func (o *CustomerAddAddressInner) HasAddressBookGender() bool`

HasAddressBookGender returns a boolean if a field has been set.

### GetAddressBookRegion

`func (o *CustomerAddAddressInner) GetAddressBookRegion() string`

GetAddressBookRegion returns the AddressBookRegion field if non-nil, zero value otherwise.

### GetAddressBookRegionOk

`func (o *CustomerAddAddressInner) GetAddressBookRegionOk() (*string, bool)`

GetAddressBookRegionOk returns a tuple with the AddressBookRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookRegion

`func (o *CustomerAddAddressInner) SetAddressBookRegion(v string)`

SetAddressBookRegion sets AddressBookRegion field to given value.

### HasAddressBookRegion

`func (o *CustomerAddAddressInner) HasAddressBookRegion() bool`

HasAddressBookRegion returns a boolean if a field has been set.

### GetAddressBookDefault

`func (o *CustomerAddAddressInner) GetAddressBookDefault() bool`

GetAddressBookDefault returns the AddressBookDefault field if non-nil, zero value otherwise.

### GetAddressBookDefaultOk

`func (o *CustomerAddAddressInner) GetAddressBookDefaultOk() (*bool, bool)`

GetAddressBookDefaultOk returns a tuple with the AddressBookDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookDefault

`func (o *CustomerAddAddressInner) SetAddressBookDefault(v bool)`

SetAddressBookDefault sets AddressBookDefault field to given value.

### HasAddressBookDefault

`func (o *CustomerAddAddressInner) HasAddressBookDefault() bool`

HasAddressBookDefault returns a boolean if a field has been set.

### GetAddressBookTaxId

`func (o *CustomerAddAddressInner) GetAddressBookTaxId() string`

GetAddressBookTaxId returns the AddressBookTaxId field if non-nil, zero value otherwise.

### GetAddressBookTaxIdOk

`func (o *CustomerAddAddressInner) GetAddressBookTaxIdOk() (*string, bool)`

GetAddressBookTaxIdOk returns a tuple with the AddressBookTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookTaxId

`func (o *CustomerAddAddressInner) SetAddressBookTaxId(v string)`

SetAddressBookTaxId sets AddressBookTaxId field to given value.

### HasAddressBookTaxId

`func (o *CustomerAddAddressInner) HasAddressBookTaxId() bool`

HasAddressBookTaxId returns a boolean if a field has been set.

### GetAddressBookIdentificationNumber

`func (o *CustomerAddAddressInner) GetAddressBookIdentificationNumber() string`

GetAddressBookIdentificationNumber returns the AddressBookIdentificationNumber field if non-nil, zero value otherwise.

### GetAddressBookIdentificationNumberOk

`func (o *CustomerAddAddressInner) GetAddressBookIdentificationNumberOk() (*string, bool)`

GetAddressBookIdentificationNumberOk returns a tuple with the AddressBookIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookIdentificationNumber

`func (o *CustomerAddAddressInner) SetAddressBookIdentificationNumber(v string)`

SetAddressBookIdentificationNumber sets AddressBookIdentificationNumber field to given value.

### HasAddressBookIdentificationNumber

`func (o *CustomerAddAddressInner) HasAddressBookIdentificationNumber() bool`

HasAddressBookIdentificationNumber returns a boolean if a field has been set.

### GetAddressBookAlias

`func (o *CustomerAddAddressInner) GetAddressBookAlias() string`

GetAddressBookAlias returns the AddressBookAlias field if non-nil, zero value otherwise.

### GetAddressBookAliasOk

`func (o *CustomerAddAddressInner) GetAddressBookAliasOk() (*string, bool)`

GetAddressBookAliasOk returns a tuple with the AddressBookAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookAlias

`func (o *CustomerAddAddressInner) SetAddressBookAlias(v string)`

SetAddressBookAlias sets AddressBookAlias field to given value.

### HasAddressBookAlias

`func (o *CustomerAddAddressInner) HasAddressBookAlias() bool`

HasAddressBookAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


