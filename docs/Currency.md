# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Iso3** | Pointer to **string** |  | [optional] 
**SymbolLeft** | Pointer to **string** |  | [optional] 
**SymbolRight** | Pointer to **string** |  | [optional] 
**Rate** | Pointer to **float32** |  | [optional] 
**Avail** | Pointer to **bool** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


