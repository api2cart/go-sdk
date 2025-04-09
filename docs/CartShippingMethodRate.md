# CartShippingMethodRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinWeight** | Pointer to **string** |  | [optional] 
**MaxWeight** | Pointer to **string** |  | [optional] 
**MinOrderAmount** | Pointer to **string** |  | [optional] 
**MaxOrderAmount** | Pointer to **string** |  | [optional] 
**MinItemsCount** | Pointer to **string** |  | [optional] 
**MaxItemsCount** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCartShippingMethodRate

`func NewCartShippingMethodRate() *CartShippingMethodRate`

NewCartShippingMethodRate instantiates a new CartShippingMethodRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartShippingMethodRateWithDefaults

`func NewCartShippingMethodRateWithDefaults() *CartShippingMethodRate`

NewCartShippingMethodRateWithDefaults instantiates a new CartShippingMethodRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinWeight

`func (o *CartShippingMethodRate) GetMinWeight() string`

GetMinWeight returns the MinWeight field if non-nil, zero value otherwise.

### GetMinWeightOk

`func (o *CartShippingMethodRate) GetMinWeightOk() (*string, bool)`

GetMinWeightOk returns a tuple with the MinWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWeight

`func (o *CartShippingMethodRate) SetMinWeight(v string)`

SetMinWeight sets MinWeight field to given value.

### HasMinWeight

`func (o *CartShippingMethodRate) HasMinWeight() bool`

HasMinWeight returns a boolean if a field has been set.

### GetMaxWeight

`func (o *CartShippingMethodRate) GetMaxWeight() string`

GetMaxWeight returns the MaxWeight field if non-nil, zero value otherwise.

### GetMaxWeightOk

`func (o *CartShippingMethodRate) GetMaxWeightOk() (*string, bool)`

GetMaxWeightOk returns a tuple with the MaxWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWeight

`func (o *CartShippingMethodRate) SetMaxWeight(v string)`

SetMaxWeight sets MaxWeight field to given value.

### HasMaxWeight

`func (o *CartShippingMethodRate) HasMaxWeight() bool`

HasMaxWeight returns a boolean if a field has been set.

### GetMinOrderAmount

`func (o *CartShippingMethodRate) GetMinOrderAmount() string`

GetMinOrderAmount returns the MinOrderAmount field if non-nil, zero value otherwise.

### GetMinOrderAmountOk

`func (o *CartShippingMethodRate) GetMinOrderAmountOk() (*string, bool)`

GetMinOrderAmountOk returns a tuple with the MinOrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrderAmount

`func (o *CartShippingMethodRate) SetMinOrderAmount(v string)`

SetMinOrderAmount sets MinOrderAmount field to given value.

### HasMinOrderAmount

`func (o *CartShippingMethodRate) HasMinOrderAmount() bool`

HasMinOrderAmount returns a boolean if a field has been set.

### GetMaxOrderAmount

`func (o *CartShippingMethodRate) GetMaxOrderAmount() string`

GetMaxOrderAmount returns the MaxOrderAmount field if non-nil, zero value otherwise.

### GetMaxOrderAmountOk

`func (o *CartShippingMethodRate) GetMaxOrderAmountOk() (*string, bool)`

GetMaxOrderAmountOk returns a tuple with the MaxOrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrderAmount

`func (o *CartShippingMethodRate) SetMaxOrderAmount(v string)`

SetMaxOrderAmount sets MaxOrderAmount field to given value.

### HasMaxOrderAmount

`func (o *CartShippingMethodRate) HasMaxOrderAmount() bool`

HasMaxOrderAmount returns a boolean if a field has been set.

### GetMinItemsCount

`func (o *CartShippingMethodRate) GetMinItemsCount() string`

GetMinItemsCount returns the MinItemsCount field if non-nil, zero value otherwise.

### GetMinItemsCountOk

`func (o *CartShippingMethodRate) GetMinItemsCountOk() (*string, bool)`

GetMinItemsCountOk returns a tuple with the MinItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItemsCount

`func (o *CartShippingMethodRate) SetMinItemsCount(v string)`

SetMinItemsCount sets MinItemsCount field to given value.

### HasMinItemsCount

`func (o *CartShippingMethodRate) HasMinItemsCount() bool`

HasMinItemsCount returns a boolean if a field has been set.

### GetMaxItemsCount

`func (o *CartShippingMethodRate) GetMaxItemsCount() string`

GetMaxItemsCount returns the MaxItemsCount field if non-nil, zero value otherwise.

### GetMaxItemsCountOk

`func (o *CartShippingMethodRate) GetMaxItemsCountOk() (*string, bool)`

GetMaxItemsCountOk returns a tuple with the MaxItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItemsCount

`func (o *CartShippingMethodRate) SetMaxItemsCount(v string)`

SetMaxItemsCount sets MaxItemsCount field to given value.

### HasMaxItemsCount

`func (o *CartShippingMethodRate) HasMaxItemsCount() bool`

HasMaxItemsCount returns a boolean if a field has been set.

### GetPrice

`func (o *CartShippingMethodRate) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CartShippingMethodRate) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CartShippingMethodRate) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CartShippingMethodRate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *CartShippingMethodRate) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CartShippingMethodRate) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CartShippingMethodRate) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CartShippingMethodRate) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *CartShippingMethodRate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CartShippingMethodRate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CartShippingMethodRate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CartShippingMethodRate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


