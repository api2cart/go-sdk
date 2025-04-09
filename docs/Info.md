# Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StateCode** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**StreetAddress** | Pointer to **string** |  | [optional] 
**StreetAddressLine2** | Pointer to **string** |  | [optional] 
**ZipCode** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInfo

`func NewInfo() *Info`

NewInfo instantiates a new Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoWithDefaults

`func NewInfoWithDefaults() *Info`

NewInfoWithDefaults instantiates a new Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *Info) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Info) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Info) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Info) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCountry

`func (o *Info) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Info) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Info) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Info) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *Info) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Info) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Info) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Info) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateCode

`func (o *Info) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *Info) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *Info) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *Info) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### GetCity

`func (o *Info) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Info) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Info) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Info) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStreetAddress

`func (o *Info) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *Info) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *Info) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *Info) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetStreetAddressLine2

`func (o *Info) GetStreetAddressLine2() string`

GetStreetAddressLine2 returns the StreetAddressLine2 field if non-nil, zero value otherwise.

### GetStreetAddressLine2Ok

`func (o *Info) GetStreetAddressLine2Ok() (*string, bool)`

GetStreetAddressLine2Ok returns a tuple with the StreetAddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddressLine2

`func (o *Info) SetStreetAddressLine2(v string)`

SetStreetAddressLine2 sets StreetAddressLine2 field to given value.

### HasStreetAddressLine2

`func (o *Info) HasStreetAddressLine2() bool`

HasStreetAddressLine2 returns a boolean if a field has been set.

### GetZipCode

`func (o *Info) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *Info) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *Info) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *Info) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetEmail

`func (o *Info) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Info) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Info) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Info) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *Info) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Info) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Info) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Info) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Info) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Info) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Info) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Info) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *Info) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Info) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Info) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Info) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


