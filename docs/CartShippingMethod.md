# CartShippingMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**HandlingFee** | Pointer to **NullableString** |  | [optional] 
**HandlingEnabled** | Pointer to **NullableString** |  | [optional] 
**HandlingType** | Pointer to **NullableString** |  | [optional] 
**DefaultPrice** | Pointer to **NullableString** |  | [optional] 
**DefaultPriceType** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **NullableString** |  | [optional] 
**MinOrderAmount** | Pointer to **NullableString** |  | [optional] 
**Rates** | Pointer to [**[]CartShippingMethodRate**](CartShippingMethodRate.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCartShippingMethod

`func NewCartShippingMethod() *CartShippingMethod`

NewCartShippingMethod instantiates a new CartShippingMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartShippingMethodWithDefaults

`func NewCartShippingMethodWithDefaults() *CartShippingMethod`

NewCartShippingMethodWithDefaults instantiates a new CartShippingMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CartShippingMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartShippingMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartShippingMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CartShippingMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CartShippingMethod) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CartShippingMethod) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHandlingFee

`func (o *CartShippingMethod) GetHandlingFee() string`

GetHandlingFee returns the HandlingFee field if non-nil, zero value otherwise.

### GetHandlingFeeOk

`func (o *CartShippingMethod) GetHandlingFeeOk() (*string, bool)`

GetHandlingFeeOk returns a tuple with the HandlingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingFee

`func (o *CartShippingMethod) SetHandlingFee(v string)`

SetHandlingFee sets HandlingFee field to given value.

### HasHandlingFee

`func (o *CartShippingMethod) HasHandlingFee() bool`

HasHandlingFee returns a boolean if a field has been set.

### SetHandlingFeeNil

`func (o *CartShippingMethod) SetHandlingFeeNil(b bool)`

 SetHandlingFeeNil sets the value for HandlingFee to be an explicit nil

### UnsetHandlingFee
`func (o *CartShippingMethod) UnsetHandlingFee()`

UnsetHandlingFee ensures that no value is present for HandlingFee, not even an explicit nil
### GetHandlingEnabled

`func (o *CartShippingMethod) GetHandlingEnabled() string`

GetHandlingEnabled returns the HandlingEnabled field if non-nil, zero value otherwise.

### GetHandlingEnabledOk

`func (o *CartShippingMethod) GetHandlingEnabledOk() (*string, bool)`

GetHandlingEnabledOk returns a tuple with the HandlingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingEnabled

`func (o *CartShippingMethod) SetHandlingEnabled(v string)`

SetHandlingEnabled sets HandlingEnabled field to given value.

### HasHandlingEnabled

`func (o *CartShippingMethod) HasHandlingEnabled() bool`

HasHandlingEnabled returns a boolean if a field has been set.

### SetHandlingEnabledNil

`func (o *CartShippingMethod) SetHandlingEnabledNil(b bool)`

 SetHandlingEnabledNil sets the value for HandlingEnabled to be an explicit nil

### UnsetHandlingEnabled
`func (o *CartShippingMethod) UnsetHandlingEnabled()`

UnsetHandlingEnabled ensures that no value is present for HandlingEnabled, not even an explicit nil
### GetHandlingType

`func (o *CartShippingMethod) GetHandlingType() string`

GetHandlingType returns the HandlingType field if non-nil, zero value otherwise.

### GetHandlingTypeOk

`func (o *CartShippingMethod) GetHandlingTypeOk() (*string, bool)`

GetHandlingTypeOk returns a tuple with the HandlingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingType

`func (o *CartShippingMethod) SetHandlingType(v string)`

SetHandlingType sets HandlingType field to given value.

### HasHandlingType

`func (o *CartShippingMethod) HasHandlingType() bool`

HasHandlingType returns a boolean if a field has been set.

### SetHandlingTypeNil

`func (o *CartShippingMethod) SetHandlingTypeNil(b bool)`

 SetHandlingTypeNil sets the value for HandlingType to be an explicit nil

### UnsetHandlingType
`func (o *CartShippingMethod) UnsetHandlingType()`

UnsetHandlingType ensures that no value is present for HandlingType, not even an explicit nil
### GetDefaultPrice

`func (o *CartShippingMethod) GetDefaultPrice() string`

GetDefaultPrice returns the DefaultPrice field if non-nil, zero value otherwise.

### GetDefaultPriceOk

`func (o *CartShippingMethod) GetDefaultPriceOk() (*string, bool)`

GetDefaultPriceOk returns a tuple with the DefaultPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrice

`func (o *CartShippingMethod) SetDefaultPrice(v string)`

SetDefaultPrice sets DefaultPrice field to given value.

### HasDefaultPrice

`func (o *CartShippingMethod) HasDefaultPrice() bool`

HasDefaultPrice returns a boolean if a field has been set.

### SetDefaultPriceNil

`func (o *CartShippingMethod) SetDefaultPriceNil(b bool)`

 SetDefaultPriceNil sets the value for DefaultPrice to be an explicit nil

### UnsetDefaultPrice
`func (o *CartShippingMethod) UnsetDefaultPrice()`

UnsetDefaultPrice ensures that no value is present for DefaultPrice, not even an explicit nil
### GetDefaultPriceType

`func (o *CartShippingMethod) GetDefaultPriceType() string`

GetDefaultPriceType returns the DefaultPriceType field if non-nil, zero value otherwise.

### GetDefaultPriceTypeOk

`func (o *CartShippingMethod) GetDefaultPriceTypeOk() (*string, bool)`

GetDefaultPriceTypeOk returns a tuple with the DefaultPriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPriceType

`func (o *CartShippingMethod) SetDefaultPriceType(v string)`

SetDefaultPriceType sets DefaultPriceType field to given value.

### HasDefaultPriceType

`func (o *CartShippingMethod) HasDefaultPriceType() bool`

HasDefaultPriceType returns a boolean if a field has been set.

### SetDefaultPriceTypeNil

`func (o *CartShippingMethod) SetDefaultPriceTypeNil(b bool)`

 SetDefaultPriceTypeNil sets the value for DefaultPriceType to be an explicit nil

### UnsetDefaultPriceType
`func (o *CartShippingMethod) UnsetDefaultPriceType()`

UnsetDefaultPriceType ensures that no value is present for DefaultPriceType, not even an explicit nil
### GetType

`func (o *CartShippingMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CartShippingMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CartShippingMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CartShippingMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CartShippingMethod) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CartShippingMethod) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetEnabled

`func (o *CartShippingMethod) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CartShippingMethod) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CartShippingMethod) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CartShippingMethod) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *CartShippingMethod) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *CartShippingMethod) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetMinOrderAmount

`func (o *CartShippingMethod) GetMinOrderAmount() string`

GetMinOrderAmount returns the MinOrderAmount field if non-nil, zero value otherwise.

### GetMinOrderAmountOk

`func (o *CartShippingMethod) GetMinOrderAmountOk() (*string, bool)`

GetMinOrderAmountOk returns a tuple with the MinOrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrderAmount

`func (o *CartShippingMethod) SetMinOrderAmount(v string)`

SetMinOrderAmount sets MinOrderAmount field to given value.

### HasMinOrderAmount

`func (o *CartShippingMethod) HasMinOrderAmount() bool`

HasMinOrderAmount returns a boolean if a field has been set.

### SetMinOrderAmountNil

`func (o *CartShippingMethod) SetMinOrderAmountNil(b bool)`

 SetMinOrderAmountNil sets the value for MinOrderAmount to be an explicit nil

### UnsetMinOrderAmount
`func (o *CartShippingMethod) UnsetMinOrderAmount()`

UnsetMinOrderAmount ensures that no value is present for MinOrderAmount, not even an explicit nil
### GetRates

`func (o *CartShippingMethod) GetRates() []CartShippingMethodRate`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *CartShippingMethod) GetRatesOk() (*[]CartShippingMethodRate, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *CartShippingMethod) SetRates(v []CartShippingMethodRate)`

SetRates sets Rates field to given value.

### HasRates

`func (o *CartShippingMethod) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *CartShippingMethod) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CartShippingMethod) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CartShippingMethod) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CartShippingMethod) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *CartShippingMethod) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *CartShippingMethod) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *CartShippingMethod) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CartShippingMethod) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CartShippingMethod) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CartShippingMethod) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CartShippingMethod) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CartShippingMethod) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


