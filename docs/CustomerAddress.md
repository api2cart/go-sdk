# CustomerAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**Postcode** | Pointer to **string** |  | [optional] 
**Address1** | Pointer to **string** |  | [optional] 
**Address2** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**PhoneMobile** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 
**State** | Pointer to [**NullableState**](State.md) |  | [optional] 
**Company** | Pointer to **NullableString** |  | [optional] 
**Fax** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**TaxId** | Pointer to **NullableString** |  | [optional] 
**IdentificationNumber** | Pointer to **NullableString** |  | [optional] 
**Alias** | Pointer to **NullableString** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCustomerAddress

`func NewCustomerAddress() *CustomerAddress`

NewCustomerAddress instantiates a new CustomerAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddressWithDefaults

`func NewCustomerAddressWithDefaults() *CustomerAddress`

NewCustomerAddressWithDefaults instantiates a new CustomerAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomerAddress) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerAddress) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerAddress) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomerAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFirstName

`func (o *CustomerAddress) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerAddress) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerAddress) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CustomerAddress) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *CustomerAddress) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *CustomerAddress) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *CustomerAddress) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerAddress) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerAddress) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CustomerAddress) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *CustomerAddress) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *CustomerAddress) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPostcode

`func (o *CustomerAddress) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *CustomerAddress) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *CustomerAddress) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *CustomerAddress) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetAddress1

`func (o *CustomerAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *CustomerAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *CustomerAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *CustomerAddress) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *CustomerAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *CustomerAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *CustomerAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *CustomerAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *CustomerAddress) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *CustomerAddress) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetPhone

`func (o *CustomerAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerAddress) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *CustomerAddress) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *CustomerAddress) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPhoneMobile

`func (o *CustomerAddress) GetPhoneMobile() string`

GetPhoneMobile returns the PhoneMobile field if non-nil, zero value otherwise.

### GetPhoneMobileOk

`func (o *CustomerAddress) GetPhoneMobileOk() (*string, bool)`

GetPhoneMobileOk returns a tuple with the PhoneMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneMobile

`func (o *CustomerAddress) SetPhoneMobile(v string)`

SetPhoneMobile sets PhoneMobile field to given value.

### HasPhoneMobile

`func (o *CustomerAddress) HasPhoneMobile() bool`

HasPhoneMobile returns a boolean if a field has been set.

### SetPhoneMobileNil

`func (o *CustomerAddress) SetPhoneMobileNil(b bool)`

 SetPhoneMobileNil sets the value for PhoneMobile to be an explicit nil

### UnsetPhoneMobile
`func (o *CustomerAddress) UnsetPhoneMobile()`

UnsetPhoneMobile ensures that no value is present for PhoneMobile, not even an explicit nil
### GetCity

`func (o *CustomerAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CustomerAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CustomerAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CustomerAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *CustomerAddress) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerAddress) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerAddress) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CustomerAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *CustomerAddress) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerAddress) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerAddress) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *CustomerAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *CustomerAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CustomerAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCompany

`func (o *CustomerAddress) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CustomerAddress) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CustomerAddress) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CustomerAddress) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *CustomerAddress) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *CustomerAddress) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetFax

`func (o *CustomerAddress) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *CustomerAddress) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *CustomerAddress) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *CustomerAddress) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFaxNil

`func (o *CustomerAddress) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *CustomerAddress) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetWebsite

`func (o *CustomerAddress) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CustomerAddress) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CustomerAddress) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CustomerAddress) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *CustomerAddress) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *CustomerAddress) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetGender

`func (o *CustomerAddress) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CustomerAddress) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CustomerAddress) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CustomerAddress) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *CustomerAddress) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *CustomerAddress) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetRegion

`func (o *CustomerAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CustomerAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CustomerAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CustomerAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *CustomerAddress) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CustomerAddress) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetDefault

`func (o *CustomerAddress) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomerAddress) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomerAddress) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomerAddress) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetTaxId

`func (o *CustomerAddress) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CustomerAddress) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CustomerAddress) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CustomerAddress) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *CustomerAddress) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *CustomerAddress) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetIdentificationNumber

`func (o *CustomerAddress) GetIdentificationNumber() string`

GetIdentificationNumber returns the IdentificationNumber field if non-nil, zero value otherwise.

### GetIdentificationNumberOk

`func (o *CustomerAddress) GetIdentificationNumberOk() (*string, bool)`

GetIdentificationNumberOk returns a tuple with the IdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationNumber

`func (o *CustomerAddress) SetIdentificationNumber(v string)`

SetIdentificationNumber sets IdentificationNumber field to given value.

### HasIdentificationNumber

`func (o *CustomerAddress) HasIdentificationNumber() bool`

HasIdentificationNumber returns a boolean if a field has been set.

### SetIdentificationNumberNil

`func (o *CustomerAddress) SetIdentificationNumberNil(b bool)`

 SetIdentificationNumberNil sets the value for IdentificationNumber to be an explicit nil

### UnsetIdentificationNumber
`func (o *CustomerAddress) UnsetIdentificationNumber()`

UnsetIdentificationNumber ensures that no value is present for IdentificationNumber, not even an explicit nil
### GetAlias

`func (o *CustomerAddress) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CustomerAddress) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CustomerAddress) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *CustomerAddress) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *CustomerAddress) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *CustomerAddress) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetAdditionalFields

`func (o *CustomerAddress) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CustomerAddress) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CustomerAddress) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CustomerAddress) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *CustomerAddress) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *CustomerAddress) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *CustomerAddress) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CustomerAddress) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CustomerAddress) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CustomerAddress) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CustomerAddress) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CustomerAddress) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


