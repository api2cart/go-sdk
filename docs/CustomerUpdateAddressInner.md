# CustomerUpdateAddressInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressBookId** | Pointer to **string** | The ID of the address. | [optional] 
**AddressBookFirstName** | Pointer to **string** | Specifies customer&#39;s first name in the address book | [optional] 
**AddressBookLastName** | Pointer to **string** | Specifies customer&#39;s last name in the address book | [optional] 
**AddressBookCompany** | Pointer to **string** | Specifies customer&#39;s company name in the address book | [optional] 
**AddressBookFax** | Pointer to **string** | Specifies customer&#39;s fax in the address book | [optional] 
**AddressBookPhone** | Pointer to **string** | Specifies customer&#39;s phone number in the address book | [optional] 
**AddressBookPhoneMobile** | Pointer to **string** | Specifies customer&#39;s mobile phone number in the address book | [optional] 
**AddressBookAddress1** | Pointer to **string** | Specifies customer&#39;s first address in the address book | [optional] 
**AddressBookAddress2** | Pointer to **string** | Specifies customer&#39;s second address in the address book | [optional] 
**AddressBookCity** | Pointer to **string** | Specifies customer&#39;s city in the address book | [optional] 
**AddressBookCountry** | Pointer to **string** | ISO code or name of country | [optional] 
**AddressBookState** | Pointer to **string** | ISO code or name of state. | [optional] 
**AddressBookPostcode** | Pointer to **string** | Specifies customer&#39;s postcode | [optional] 
**AddressBookTaxId** | Pointer to **string** | Add Tax Id | [optional] 
**AddressBookIdentificationNumber** | Pointer to **string** | The national ID card number of this person, or a unique tax identification number. | [optional] 
**AddressBookGender** | Pointer to **string** | Specifies customer&#39;s gender | [optional] 
**AddressBookAlias** | Pointer to **string** | Specifies customer&#39;s alias in the address book | [optional] 
**AddressBookType** | Pointer to **string** | Specifies customer&#39;s address type | [optional] 
**AddressBookDefault** | Pointer to **bool** | Defines whether the address is used by default | [optional] 

## Methods

### NewCustomerUpdateAddressInner

`func NewCustomerUpdateAddressInner() *CustomerUpdateAddressInner`

NewCustomerUpdateAddressInner instantiates a new CustomerUpdateAddressInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerUpdateAddressInnerWithDefaults

`func NewCustomerUpdateAddressInnerWithDefaults() *CustomerUpdateAddressInner`

NewCustomerUpdateAddressInnerWithDefaults instantiates a new CustomerUpdateAddressInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressBookId

`func (o *CustomerUpdateAddressInner) GetAddressBookId() string`

GetAddressBookId returns the AddressBookId field if non-nil, zero value otherwise.

### GetAddressBookIdOk

`func (o *CustomerUpdateAddressInner) GetAddressBookIdOk() (*string, bool)`

GetAddressBookIdOk returns a tuple with the AddressBookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookId

`func (o *CustomerUpdateAddressInner) SetAddressBookId(v string)`

SetAddressBookId sets AddressBookId field to given value.

### HasAddressBookId

`func (o *CustomerUpdateAddressInner) HasAddressBookId() bool`

HasAddressBookId returns a boolean if a field has been set.

### GetAddressBookFirstName

`func (o *CustomerUpdateAddressInner) GetAddressBookFirstName() string`

GetAddressBookFirstName returns the AddressBookFirstName field if non-nil, zero value otherwise.

### GetAddressBookFirstNameOk

`func (o *CustomerUpdateAddressInner) GetAddressBookFirstNameOk() (*string, bool)`

GetAddressBookFirstNameOk returns a tuple with the AddressBookFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookFirstName

`func (o *CustomerUpdateAddressInner) SetAddressBookFirstName(v string)`

SetAddressBookFirstName sets AddressBookFirstName field to given value.

### HasAddressBookFirstName

`func (o *CustomerUpdateAddressInner) HasAddressBookFirstName() bool`

HasAddressBookFirstName returns a boolean if a field has been set.

### GetAddressBookLastName

`func (o *CustomerUpdateAddressInner) GetAddressBookLastName() string`

GetAddressBookLastName returns the AddressBookLastName field if non-nil, zero value otherwise.

### GetAddressBookLastNameOk

`func (o *CustomerUpdateAddressInner) GetAddressBookLastNameOk() (*string, bool)`

GetAddressBookLastNameOk returns a tuple with the AddressBookLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookLastName

`func (o *CustomerUpdateAddressInner) SetAddressBookLastName(v string)`

SetAddressBookLastName sets AddressBookLastName field to given value.

### HasAddressBookLastName

`func (o *CustomerUpdateAddressInner) HasAddressBookLastName() bool`

HasAddressBookLastName returns a boolean if a field has been set.

### GetAddressBookCompany

`func (o *CustomerUpdateAddressInner) GetAddressBookCompany() string`

GetAddressBookCompany returns the AddressBookCompany field if non-nil, zero value otherwise.

### GetAddressBookCompanyOk

`func (o *CustomerUpdateAddressInner) GetAddressBookCompanyOk() (*string, bool)`

GetAddressBookCompanyOk returns a tuple with the AddressBookCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookCompany

`func (o *CustomerUpdateAddressInner) SetAddressBookCompany(v string)`

SetAddressBookCompany sets AddressBookCompany field to given value.

### HasAddressBookCompany

`func (o *CustomerUpdateAddressInner) HasAddressBookCompany() bool`

HasAddressBookCompany returns a boolean if a field has been set.

### GetAddressBookFax

`func (o *CustomerUpdateAddressInner) GetAddressBookFax() string`

GetAddressBookFax returns the AddressBookFax field if non-nil, zero value otherwise.

### GetAddressBookFaxOk

`func (o *CustomerUpdateAddressInner) GetAddressBookFaxOk() (*string, bool)`

GetAddressBookFaxOk returns a tuple with the AddressBookFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookFax

`func (o *CustomerUpdateAddressInner) SetAddressBookFax(v string)`

SetAddressBookFax sets AddressBookFax field to given value.

### HasAddressBookFax

`func (o *CustomerUpdateAddressInner) HasAddressBookFax() bool`

HasAddressBookFax returns a boolean if a field has been set.

### GetAddressBookPhone

`func (o *CustomerUpdateAddressInner) GetAddressBookPhone() string`

GetAddressBookPhone returns the AddressBookPhone field if non-nil, zero value otherwise.

### GetAddressBookPhoneOk

`func (o *CustomerUpdateAddressInner) GetAddressBookPhoneOk() (*string, bool)`

GetAddressBookPhoneOk returns a tuple with the AddressBookPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookPhone

`func (o *CustomerUpdateAddressInner) SetAddressBookPhone(v string)`

SetAddressBookPhone sets AddressBookPhone field to given value.

### HasAddressBookPhone

`func (o *CustomerUpdateAddressInner) HasAddressBookPhone() bool`

HasAddressBookPhone returns a boolean if a field has been set.

### GetAddressBookPhoneMobile

`func (o *CustomerUpdateAddressInner) GetAddressBookPhoneMobile() string`

GetAddressBookPhoneMobile returns the AddressBookPhoneMobile field if non-nil, zero value otherwise.

### GetAddressBookPhoneMobileOk

`func (o *CustomerUpdateAddressInner) GetAddressBookPhoneMobileOk() (*string, bool)`

GetAddressBookPhoneMobileOk returns a tuple with the AddressBookPhoneMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookPhoneMobile

`func (o *CustomerUpdateAddressInner) SetAddressBookPhoneMobile(v string)`

SetAddressBookPhoneMobile sets AddressBookPhoneMobile field to given value.

### HasAddressBookPhoneMobile

`func (o *CustomerUpdateAddressInner) HasAddressBookPhoneMobile() bool`

HasAddressBookPhoneMobile returns a boolean if a field has been set.

### GetAddressBookAddress1

`func (o *CustomerUpdateAddressInner) GetAddressBookAddress1() string`

GetAddressBookAddress1 returns the AddressBookAddress1 field if non-nil, zero value otherwise.

### GetAddressBookAddress1Ok

`func (o *CustomerUpdateAddressInner) GetAddressBookAddress1Ok() (*string, bool)`

GetAddressBookAddress1Ok returns a tuple with the AddressBookAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookAddress1

`func (o *CustomerUpdateAddressInner) SetAddressBookAddress1(v string)`

SetAddressBookAddress1 sets AddressBookAddress1 field to given value.

### HasAddressBookAddress1

`func (o *CustomerUpdateAddressInner) HasAddressBookAddress1() bool`

HasAddressBookAddress1 returns a boolean if a field has been set.

### GetAddressBookAddress2

`func (o *CustomerUpdateAddressInner) GetAddressBookAddress2() string`

GetAddressBookAddress2 returns the AddressBookAddress2 field if non-nil, zero value otherwise.

### GetAddressBookAddress2Ok

`func (o *CustomerUpdateAddressInner) GetAddressBookAddress2Ok() (*string, bool)`

GetAddressBookAddress2Ok returns a tuple with the AddressBookAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookAddress2

`func (o *CustomerUpdateAddressInner) SetAddressBookAddress2(v string)`

SetAddressBookAddress2 sets AddressBookAddress2 field to given value.

### HasAddressBookAddress2

`func (o *CustomerUpdateAddressInner) HasAddressBookAddress2() bool`

HasAddressBookAddress2 returns a boolean if a field has been set.

### GetAddressBookCity

`func (o *CustomerUpdateAddressInner) GetAddressBookCity() string`

GetAddressBookCity returns the AddressBookCity field if non-nil, zero value otherwise.

### GetAddressBookCityOk

`func (o *CustomerUpdateAddressInner) GetAddressBookCityOk() (*string, bool)`

GetAddressBookCityOk returns a tuple with the AddressBookCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookCity

`func (o *CustomerUpdateAddressInner) SetAddressBookCity(v string)`

SetAddressBookCity sets AddressBookCity field to given value.

### HasAddressBookCity

`func (o *CustomerUpdateAddressInner) HasAddressBookCity() bool`

HasAddressBookCity returns a boolean if a field has been set.

### GetAddressBookCountry

`func (o *CustomerUpdateAddressInner) GetAddressBookCountry() string`

GetAddressBookCountry returns the AddressBookCountry field if non-nil, zero value otherwise.

### GetAddressBookCountryOk

`func (o *CustomerUpdateAddressInner) GetAddressBookCountryOk() (*string, bool)`

GetAddressBookCountryOk returns a tuple with the AddressBookCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookCountry

`func (o *CustomerUpdateAddressInner) SetAddressBookCountry(v string)`

SetAddressBookCountry sets AddressBookCountry field to given value.

### HasAddressBookCountry

`func (o *CustomerUpdateAddressInner) HasAddressBookCountry() bool`

HasAddressBookCountry returns a boolean if a field has been set.

### GetAddressBookState

`func (o *CustomerUpdateAddressInner) GetAddressBookState() string`

GetAddressBookState returns the AddressBookState field if non-nil, zero value otherwise.

### GetAddressBookStateOk

`func (o *CustomerUpdateAddressInner) GetAddressBookStateOk() (*string, bool)`

GetAddressBookStateOk returns a tuple with the AddressBookState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookState

`func (o *CustomerUpdateAddressInner) SetAddressBookState(v string)`

SetAddressBookState sets AddressBookState field to given value.

### HasAddressBookState

`func (o *CustomerUpdateAddressInner) HasAddressBookState() bool`

HasAddressBookState returns a boolean if a field has been set.

### GetAddressBookPostcode

`func (o *CustomerUpdateAddressInner) GetAddressBookPostcode() string`

GetAddressBookPostcode returns the AddressBookPostcode field if non-nil, zero value otherwise.

### GetAddressBookPostcodeOk

`func (o *CustomerUpdateAddressInner) GetAddressBookPostcodeOk() (*string, bool)`

GetAddressBookPostcodeOk returns a tuple with the AddressBookPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookPostcode

`func (o *CustomerUpdateAddressInner) SetAddressBookPostcode(v string)`

SetAddressBookPostcode sets AddressBookPostcode field to given value.

### HasAddressBookPostcode

`func (o *CustomerUpdateAddressInner) HasAddressBookPostcode() bool`

HasAddressBookPostcode returns a boolean if a field has been set.

### GetAddressBookTaxId

`func (o *CustomerUpdateAddressInner) GetAddressBookTaxId() string`

GetAddressBookTaxId returns the AddressBookTaxId field if non-nil, zero value otherwise.

### GetAddressBookTaxIdOk

`func (o *CustomerUpdateAddressInner) GetAddressBookTaxIdOk() (*string, bool)`

GetAddressBookTaxIdOk returns a tuple with the AddressBookTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookTaxId

`func (o *CustomerUpdateAddressInner) SetAddressBookTaxId(v string)`

SetAddressBookTaxId sets AddressBookTaxId field to given value.

### HasAddressBookTaxId

`func (o *CustomerUpdateAddressInner) HasAddressBookTaxId() bool`

HasAddressBookTaxId returns a boolean if a field has been set.

### GetAddressBookIdentificationNumber

`func (o *CustomerUpdateAddressInner) GetAddressBookIdentificationNumber() string`

GetAddressBookIdentificationNumber returns the AddressBookIdentificationNumber field if non-nil, zero value otherwise.

### GetAddressBookIdentificationNumberOk

`func (o *CustomerUpdateAddressInner) GetAddressBookIdentificationNumberOk() (*string, bool)`

GetAddressBookIdentificationNumberOk returns a tuple with the AddressBookIdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookIdentificationNumber

`func (o *CustomerUpdateAddressInner) SetAddressBookIdentificationNumber(v string)`

SetAddressBookIdentificationNumber sets AddressBookIdentificationNumber field to given value.

### HasAddressBookIdentificationNumber

`func (o *CustomerUpdateAddressInner) HasAddressBookIdentificationNumber() bool`

HasAddressBookIdentificationNumber returns a boolean if a field has been set.

### GetAddressBookGender

`func (o *CustomerUpdateAddressInner) GetAddressBookGender() string`

GetAddressBookGender returns the AddressBookGender field if non-nil, zero value otherwise.

### GetAddressBookGenderOk

`func (o *CustomerUpdateAddressInner) GetAddressBookGenderOk() (*string, bool)`

GetAddressBookGenderOk returns a tuple with the AddressBookGender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookGender

`func (o *CustomerUpdateAddressInner) SetAddressBookGender(v string)`

SetAddressBookGender sets AddressBookGender field to given value.

### HasAddressBookGender

`func (o *CustomerUpdateAddressInner) HasAddressBookGender() bool`

HasAddressBookGender returns a boolean if a field has been set.

### GetAddressBookAlias

`func (o *CustomerUpdateAddressInner) GetAddressBookAlias() string`

GetAddressBookAlias returns the AddressBookAlias field if non-nil, zero value otherwise.

### GetAddressBookAliasOk

`func (o *CustomerUpdateAddressInner) GetAddressBookAliasOk() (*string, bool)`

GetAddressBookAliasOk returns a tuple with the AddressBookAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookAlias

`func (o *CustomerUpdateAddressInner) SetAddressBookAlias(v string)`

SetAddressBookAlias sets AddressBookAlias field to given value.

### HasAddressBookAlias

`func (o *CustomerUpdateAddressInner) HasAddressBookAlias() bool`

HasAddressBookAlias returns a boolean if a field has been set.

### GetAddressBookType

`func (o *CustomerUpdateAddressInner) GetAddressBookType() string`

GetAddressBookType returns the AddressBookType field if non-nil, zero value otherwise.

### GetAddressBookTypeOk

`func (o *CustomerUpdateAddressInner) GetAddressBookTypeOk() (*string, bool)`

GetAddressBookTypeOk returns a tuple with the AddressBookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookType

`func (o *CustomerUpdateAddressInner) SetAddressBookType(v string)`

SetAddressBookType sets AddressBookType field to given value.

### HasAddressBookType

`func (o *CustomerUpdateAddressInner) HasAddressBookType() bool`

HasAddressBookType returns a boolean if a field has been set.

### GetAddressBookDefault

`func (o *CustomerUpdateAddressInner) GetAddressBookDefault() bool`

GetAddressBookDefault returns the AddressBookDefault field if non-nil, zero value otherwise.

### GetAddressBookDefaultOk

`func (o *CustomerUpdateAddressInner) GetAddressBookDefaultOk() (*bool, bool)`

GetAddressBookDefaultOk returns a tuple with the AddressBookDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBookDefault

`func (o *CustomerUpdateAddressInner) SetAddressBookDefault(v bool)`

SetAddressBookDefault sets AddressBookDefault field to given value.

### HasAddressBookDefault

`func (o *CustomerUpdateAddressInner) HasAddressBookDefault() bool`

HasAddressBookDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


