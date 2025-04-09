# CartStoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**StoreLanguages** | Pointer to [**[]Language**](Language.md) |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**StoreCurrencies** | Pointer to [**[]Currency**](Currency.md) |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**RootCategoryId** | Pointer to **string** |  | [optional] 
**MultiStoreUrl** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**WeightUnit** | Pointer to **string** |  | [optional] 
**DimensionUnit** | Pointer to **string** |  | [optional] 
**PricesIncludeTax** | Pointer to **bool** |  | [optional] 
**CarrierInfo** | Pointer to [**[]Carrier**](Carrier.md) |  | [optional] 
**StoreOwnerInfo** | Pointer to [**Info**](Info.md) |  | [optional] 
**DefaultWarehouseId** | Pointer to **string** |  | [optional] 
**Channels** | Pointer to [**[]CartChannel**](CartChannel.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCartStoreInfo

`func NewCartStoreInfo() *CartStoreInfo`

NewCartStoreInfo instantiates a new CartStoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartStoreInfoWithDefaults

`func NewCartStoreInfoWithDefaults() *CartStoreInfo`

NewCartStoreInfoWithDefaults instantiates a new CartStoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreId

`func (o *CartStoreInfo) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CartStoreInfo) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CartStoreInfo) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *CartStoreInfo) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetName

`func (o *CartStoreInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartStoreInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartStoreInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CartStoreInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanguage

`func (o *CartStoreInfo) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CartStoreInfo) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CartStoreInfo) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CartStoreInfo) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetStoreLanguages

`func (o *CartStoreInfo) GetStoreLanguages() []Language`

GetStoreLanguages returns the StoreLanguages field if non-nil, zero value otherwise.

### GetStoreLanguagesOk

`func (o *CartStoreInfo) GetStoreLanguagesOk() (*[]Language, bool)`

GetStoreLanguagesOk returns a tuple with the StoreLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreLanguages

`func (o *CartStoreInfo) SetStoreLanguages(v []Language)`

SetStoreLanguages sets StoreLanguages field to given value.

### HasStoreLanguages

`func (o *CartStoreInfo) HasStoreLanguages() bool`

HasStoreLanguages returns a boolean if a field has been set.

### GetCurrency

`func (o *CartStoreInfo) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CartStoreInfo) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CartStoreInfo) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CartStoreInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStoreCurrencies

`func (o *CartStoreInfo) GetStoreCurrencies() []Currency`

GetStoreCurrencies returns the StoreCurrencies field if non-nil, zero value otherwise.

### GetStoreCurrenciesOk

`func (o *CartStoreInfo) GetStoreCurrenciesOk() (*[]Currency, bool)`

GetStoreCurrenciesOk returns a tuple with the StoreCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreCurrencies

`func (o *CartStoreInfo) SetStoreCurrencies(v []Currency)`

SetStoreCurrencies sets StoreCurrencies field to given value.

### HasStoreCurrencies

`func (o *CartStoreInfo) HasStoreCurrencies() bool`

HasStoreCurrencies returns a boolean if a field has been set.

### GetTimezone

`func (o *CartStoreInfo) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CartStoreInfo) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CartStoreInfo) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CartStoreInfo) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetCountry

`func (o *CartStoreInfo) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CartStoreInfo) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CartStoreInfo) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CartStoreInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRootCategoryId

`func (o *CartStoreInfo) GetRootCategoryId() string`

GetRootCategoryId returns the RootCategoryId field if non-nil, zero value otherwise.

### GetRootCategoryIdOk

`func (o *CartStoreInfo) GetRootCategoryIdOk() (*string, bool)`

GetRootCategoryIdOk returns a tuple with the RootCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCategoryId

`func (o *CartStoreInfo) SetRootCategoryId(v string)`

SetRootCategoryId sets RootCategoryId field to given value.

### HasRootCategoryId

`func (o *CartStoreInfo) HasRootCategoryId() bool`

HasRootCategoryId returns a boolean if a field has been set.

### GetMultiStoreUrl

`func (o *CartStoreInfo) GetMultiStoreUrl() string`

GetMultiStoreUrl returns the MultiStoreUrl field if non-nil, zero value otherwise.

### GetMultiStoreUrlOk

`func (o *CartStoreInfo) GetMultiStoreUrlOk() (*string, bool)`

GetMultiStoreUrlOk returns a tuple with the MultiStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStoreUrl

`func (o *CartStoreInfo) SetMultiStoreUrl(v string)`

SetMultiStoreUrl sets MultiStoreUrl field to given value.

### HasMultiStoreUrl

`func (o *CartStoreInfo) HasMultiStoreUrl() bool`

HasMultiStoreUrl returns a boolean if a field has been set.

### GetActive

`func (o *CartStoreInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CartStoreInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CartStoreInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CartStoreInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetWeightUnit

`func (o *CartStoreInfo) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *CartStoreInfo) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *CartStoreInfo) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *CartStoreInfo) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetDimensionUnit

`func (o *CartStoreInfo) GetDimensionUnit() string`

GetDimensionUnit returns the DimensionUnit field if non-nil, zero value otherwise.

### GetDimensionUnitOk

`func (o *CartStoreInfo) GetDimensionUnitOk() (*string, bool)`

GetDimensionUnitOk returns a tuple with the DimensionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionUnit

`func (o *CartStoreInfo) SetDimensionUnit(v string)`

SetDimensionUnit sets DimensionUnit field to given value.

### HasDimensionUnit

`func (o *CartStoreInfo) HasDimensionUnit() bool`

HasDimensionUnit returns a boolean if a field has been set.

### GetPricesIncludeTax

`func (o *CartStoreInfo) GetPricesIncludeTax() bool`

GetPricesIncludeTax returns the PricesIncludeTax field if non-nil, zero value otherwise.

### GetPricesIncludeTaxOk

`func (o *CartStoreInfo) GetPricesIncludeTaxOk() (*bool, bool)`

GetPricesIncludeTaxOk returns a tuple with the PricesIncludeTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesIncludeTax

`func (o *CartStoreInfo) SetPricesIncludeTax(v bool)`

SetPricesIncludeTax sets PricesIncludeTax field to given value.

### HasPricesIncludeTax

`func (o *CartStoreInfo) HasPricesIncludeTax() bool`

HasPricesIncludeTax returns a boolean if a field has been set.

### GetCarrierInfo

`func (o *CartStoreInfo) GetCarrierInfo() []Carrier`

GetCarrierInfo returns the CarrierInfo field if non-nil, zero value otherwise.

### GetCarrierInfoOk

`func (o *CartStoreInfo) GetCarrierInfoOk() (*[]Carrier, bool)`

GetCarrierInfoOk returns a tuple with the CarrierInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierInfo

`func (o *CartStoreInfo) SetCarrierInfo(v []Carrier)`

SetCarrierInfo sets CarrierInfo field to given value.

### HasCarrierInfo

`func (o *CartStoreInfo) HasCarrierInfo() bool`

HasCarrierInfo returns a boolean if a field has been set.

### GetStoreOwnerInfo

`func (o *CartStoreInfo) GetStoreOwnerInfo() Info`

GetStoreOwnerInfo returns the StoreOwnerInfo field if non-nil, zero value otherwise.

### GetStoreOwnerInfoOk

`func (o *CartStoreInfo) GetStoreOwnerInfoOk() (*Info, bool)`

GetStoreOwnerInfoOk returns a tuple with the StoreOwnerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreOwnerInfo

`func (o *CartStoreInfo) SetStoreOwnerInfo(v Info)`

SetStoreOwnerInfo sets StoreOwnerInfo field to given value.

### HasStoreOwnerInfo

`func (o *CartStoreInfo) HasStoreOwnerInfo() bool`

HasStoreOwnerInfo returns a boolean if a field has been set.

### GetDefaultWarehouseId

`func (o *CartStoreInfo) GetDefaultWarehouseId() string`

GetDefaultWarehouseId returns the DefaultWarehouseId field if non-nil, zero value otherwise.

### GetDefaultWarehouseIdOk

`func (o *CartStoreInfo) GetDefaultWarehouseIdOk() (*string, bool)`

GetDefaultWarehouseIdOk returns a tuple with the DefaultWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWarehouseId

`func (o *CartStoreInfo) SetDefaultWarehouseId(v string)`

SetDefaultWarehouseId sets DefaultWarehouseId field to given value.

### HasDefaultWarehouseId

`func (o *CartStoreInfo) HasDefaultWarehouseId() bool`

HasDefaultWarehouseId returns a boolean if a field has been set.

### GetChannels

`func (o *CartStoreInfo) GetChannels() []CartChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CartStoreInfo) GetChannelsOk() (*[]CartChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CartStoreInfo) SetChannels(v []CartChannel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CartStoreInfo) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *CartStoreInfo) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CartStoreInfo) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CartStoreInfo) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CartStoreInfo) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *CartStoreInfo) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CartStoreInfo) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CartStoreInfo) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CartStoreInfo) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


