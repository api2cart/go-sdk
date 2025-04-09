# TaxClassRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tax** | Pointer to **float32** |  | [optional] 
**TaxType** | Pointer to **int32** |  | [optional] 
**TaxBasedOn** | Pointer to **string** |  | [optional] 
**Countries** | Pointer to [**[]TaxClassCountries**](TaxClassCountries.md) |  | [optional] 
**Cities** | Pointer to **[]string** |  | [optional] 
**Address** | Pointer to **[]string** |  | [optional] 
**ZipCodes** | Pointer to [**[]TaxClassZipCodes**](TaxClassZipCodes.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTaxClassRate

`func NewTaxClassRate() *TaxClassRate`

NewTaxClassRate instantiates a new TaxClassRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxClassRateWithDefaults

`func NewTaxClassRateWithDefaults() *TaxClassRate`

NewTaxClassRateWithDefaults instantiates a new TaxClassRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxClassRate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxClassRate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxClassRate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxClassRate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaxClassRate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxClassRate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxClassRate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxClassRate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTax

`func (o *TaxClassRate) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *TaxClassRate) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *TaxClassRate) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *TaxClassRate) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxType

`func (o *TaxClassRate) GetTaxType() int32`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *TaxClassRate) GetTaxTypeOk() (*int32, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *TaxClassRate) SetTaxType(v int32)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *TaxClassRate) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetTaxBasedOn

`func (o *TaxClassRate) GetTaxBasedOn() string`

GetTaxBasedOn returns the TaxBasedOn field if non-nil, zero value otherwise.

### GetTaxBasedOnOk

`func (o *TaxClassRate) GetTaxBasedOnOk() (*string, bool)`

GetTaxBasedOnOk returns a tuple with the TaxBasedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBasedOn

`func (o *TaxClassRate) SetTaxBasedOn(v string)`

SetTaxBasedOn sets TaxBasedOn field to given value.

### HasTaxBasedOn

`func (o *TaxClassRate) HasTaxBasedOn() bool`

HasTaxBasedOn returns a boolean if a field has been set.

### GetCountries

`func (o *TaxClassRate) GetCountries() []TaxClassCountries`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *TaxClassRate) GetCountriesOk() (*[]TaxClassCountries, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *TaxClassRate) SetCountries(v []TaxClassCountries)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *TaxClassRate) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetCities

`func (o *TaxClassRate) GetCities() []string`

GetCities returns the Cities field if non-nil, zero value otherwise.

### GetCitiesOk

`func (o *TaxClassRate) GetCitiesOk() (*[]string, bool)`

GetCitiesOk returns a tuple with the Cities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCities

`func (o *TaxClassRate) SetCities(v []string)`

SetCities sets Cities field to given value.

### HasCities

`func (o *TaxClassRate) HasCities() bool`

HasCities returns a boolean if a field has been set.

### GetAddress

`func (o *TaxClassRate) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TaxClassRate) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TaxClassRate) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TaxClassRate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetZipCodes

`func (o *TaxClassRate) GetZipCodes() []TaxClassZipCodes`

GetZipCodes returns the ZipCodes field if non-nil, zero value otherwise.

### GetZipCodesOk

`func (o *TaxClassRate) GetZipCodesOk() (*[]TaxClassZipCodes, bool)`

GetZipCodesOk returns a tuple with the ZipCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodes

`func (o *TaxClassRate) SetZipCodes(v []TaxClassZipCodes)`

SetZipCodes sets ZipCodes field to given value.

### HasZipCodes

`func (o *TaxClassRate) HasZipCodes() bool`

HasZipCodes returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *TaxClassRate) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *TaxClassRate) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *TaxClassRate) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *TaxClassRate) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *TaxClassRate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TaxClassRate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TaxClassRate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TaxClassRate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


