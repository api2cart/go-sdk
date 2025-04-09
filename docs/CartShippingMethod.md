# CartShippingMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**HandlingFee** | Pointer to **string** |  | [optional] 
**HandlingEnabled** | Pointer to **string** |  | [optional] 
**HandlingType** | Pointer to **string** |  | [optional] 
**DefaultPrice** | Pointer to **string** |  | [optional] 
**DefaultPriceType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **string** |  | [optional] 
**MinOrderAmount** | Pointer to **string** |  | [optional] 
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


