# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Iso3** | Pointer to **string** |  | [optional] 
**SymbolLeft** | Pointer to **string** |  | [optional] 
**SymbolRight** | Pointer to **NullableString** |  | [optional] 
**Rate** | Pointer to **NullableFloat32** |  | [optional] 
**Avail** | Pointer to **NullableBool** |  | [optional] 
**Default** | Pointer to **NullableBool** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCurrency

`func NewCurrency() *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Currency) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Currency) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Currency) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Currency) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Currency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Currency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Currency) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Currency) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Currency) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Currency) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIso3

`func (o *Currency) GetIso3() string`

GetIso3 returns the Iso3 field if non-nil, zero value otherwise.

### GetIso3Ok

`func (o *Currency) GetIso3Ok() (*string, bool)`

GetIso3Ok returns a tuple with the Iso3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso3

`func (o *Currency) SetIso3(v string)`

SetIso3 sets Iso3 field to given value.

### HasIso3

`func (o *Currency) HasIso3() bool`

HasIso3 returns a boolean if a field has been set.

### GetSymbolLeft

`func (o *Currency) GetSymbolLeft() string`

GetSymbolLeft returns the SymbolLeft field if non-nil, zero value otherwise.

### GetSymbolLeftOk

`func (o *Currency) GetSymbolLeftOk() (*string, bool)`

GetSymbolLeftOk returns a tuple with the SymbolLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolLeft

`func (o *Currency) SetSymbolLeft(v string)`

SetSymbolLeft sets SymbolLeft field to given value.

### HasSymbolLeft

`func (o *Currency) HasSymbolLeft() bool`

HasSymbolLeft returns a boolean if a field has been set.

### GetSymbolRight

`func (o *Currency) GetSymbolRight() string`

GetSymbolRight returns the SymbolRight field if non-nil, zero value otherwise.

### GetSymbolRightOk

`func (o *Currency) GetSymbolRightOk() (*string, bool)`

GetSymbolRightOk returns a tuple with the SymbolRight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolRight

`func (o *Currency) SetSymbolRight(v string)`

SetSymbolRight sets SymbolRight field to given value.

### HasSymbolRight

`func (o *Currency) HasSymbolRight() bool`

HasSymbolRight returns a boolean if a field has been set.

### SetSymbolRightNil

`func (o *Currency) SetSymbolRightNil(b bool)`

 SetSymbolRightNil sets the value for SymbolRight to be an explicit nil

### UnsetSymbolRight
`func (o *Currency) UnsetSymbolRight()`

UnsetSymbolRight ensures that no value is present for SymbolRight, not even an explicit nil
### GetRate

`func (o *Currency) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *Currency) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *Currency) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *Currency) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRateNil

`func (o *Currency) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *Currency) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetAvail

`func (o *Currency) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *Currency) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *Currency) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *Currency) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### SetAvailNil

`func (o *Currency) SetAvailNil(b bool)`

 SetAvailNil sets the value for Avail to be an explicit nil

### UnsetAvail
`func (o *Currency) UnsetAvail()`

UnsetAvail ensures that no value is present for Avail, not even an explicit nil
### GetDefault

`func (o *Currency) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Currency) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Currency) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Currency) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *Currency) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *Currency) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetAdditionalFields

`func (o *Currency) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Currency) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Currency) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Currency) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *Currency) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *Currency) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *Currency) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Currency) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Currency) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Currency) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *Currency) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *Currency) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


