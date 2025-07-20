# TaxClassCountries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Tax** | Pointer to **float32** |  | [optional] 
**TaxType** | Pointer to **int32** |  | [optional] 
**States** | Pointer to [**[]TaxClassStates**](TaxClassStates.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTaxClassCountries

`func NewTaxClassCountries() *TaxClassCountries`

NewTaxClassCountries instantiates a new TaxClassCountries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxClassCountriesWithDefaults

`func NewTaxClassCountriesWithDefaults() *TaxClassCountries`

NewTaxClassCountriesWithDefaults instantiates a new TaxClassCountries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxClassCountries) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxClassCountries) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxClassCountries) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxClassCountries) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaxClassCountries) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaxClassCountries) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TaxClassCountries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxClassCountries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxClassCountries) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxClassCountries) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *TaxClassCountries) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaxClassCountries) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaxClassCountries) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaxClassCountries) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTax

`func (o *TaxClassCountries) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *TaxClassCountries) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *TaxClassCountries) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *TaxClassCountries) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxType

`func (o *TaxClassCountries) GetTaxType() int32`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *TaxClassCountries) GetTaxTypeOk() (*int32, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *TaxClassCountries) SetTaxType(v int32)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *TaxClassCountries) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetStates

`func (o *TaxClassCountries) GetStates() []TaxClassStates`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *TaxClassCountries) GetStatesOk() (*[]TaxClassStates, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *TaxClassCountries) SetStates(v []TaxClassStates)`

SetStates sets States field to given value.

### HasStates

`func (o *TaxClassCountries) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *TaxClassCountries) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *TaxClassCountries) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *TaxClassCountries) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *TaxClassCountries) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *TaxClassCountries) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *TaxClassCountries) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *TaxClassCountries) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TaxClassCountries) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TaxClassCountries) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TaxClassCountries) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *TaxClassCountries) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *TaxClassCountries) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


