# AnalyticsProductMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemsSold** | Pointer to **NullableInt32** |  | [optional] 
**OrdersCount** | Pointer to **NullableInt32** |  | [optional] 
**GrossSales** | Pointer to **NullableFloat32** |  | [optional] 
**NetSales** | Pointer to **NullableFloat32** |  | [optional] 
**Discounts** | Pointer to **NullableFloat32** |  | [optional] 
**Refunds** | Pointer to **NullableFloat32** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAnalyticsProductMetric

`func NewAnalyticsProductMetric() *AnalyticsProductMetric`

NewAnalyticsProductMetric instantiates a new AnalyticsProductMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsProductMetricWithDefaults

`func NewAnalyticsProductMetricWithDefaults() *AnalyticsProductMetric`

NewAnalyticsProductMetricWithDefaults instantiates a new AnalyticsProductMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemsSold

`func (o *AnalyticsProductMetric) GetItemsSold() int32`

GetItemsSold returns the ItemsSold field if non-nil, zero value otherwise.

### GetItemsSoldOk

`func (o *AnalyticsProductMetric) GetItemsSoldOk() (*int32, bool)`

GetItemsSoldOk returns a tuple with the ItemsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsSold

`func (o *AnalyticsProductMetric) SetItemsSold(v int32)`

SetItemsSold sets ItemsSold field to given value.

### HasItemsSold

`func (o *AnalyticsProductMetric) HasItemsSold() bool`

HasItemsSold returns a boolean if a field has been set.

### SetItemsSoldNil

`func (o *AnalyticsProductMetric) SetItemsSoldNil(b bool)`

 SetItemsSoldNil sets the value for ItemsSold to be an explicit nil

### UnsetItemsSold
`func (o *AnalyticsProductMetric) UnsetItemsSold()`

UnsetItemsSold ensures that no value is present for ItemsSold, not even an explicit nil
### GetOrdersCount

`func (o *AnalyticsProductMetric) GetOrdersCount() int32`

GetOrdersCount returns the OrdersCount field if non-nil, zero value otherwise.

### GetOrdersCountOk

`func (o *AnalyticsProductMetric) GetOrdersCountOk() (*int32, bool)`

GetOrdersCountOk returns a tuple with the OrdersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersCount

`func (o *AnalyticsProductMetric) SetOrdersCount(v int32)`

SetOrdersCount sets OrdersCount field to given value.

### HasOrdersCount

`func (o *AnalyticsProductMetric) HasOrdersCount() bool`

HasOrdersCount returns a boolean if a field has been set.

### SetOrdersCountNil

`func (o *AnalyticsProductMetric) SetOrdersCountNil(b bool)`

 SetOrdersCountNil sets the value for OrdersCount to be an explicit nil

### UnsetOrdersCount
`func (o *AnalyticsProductMetric) UnsetOrdersCount()`

UnsetOrdersCount ensures that no value is present for OrdersCount, not even an explicit nil
### GetGrossSales

`func (o *AnalyticsProductMetric) GetGrossSales() float32`

GetGrossSales returns the GrossSales field if non-nil, zero value otherwise.

### GetGrossSalesOk

`func (o *AnalyticsProductMetric) GetGrossSalesOk() (*float32, bool)`

GetGrossSalesOk returns a tuple with the GrossSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossSales

`func (o *AnalyticsProductMetric) SetGrossSales(v float32)`

SetGrossSales sets GrossSales field to given value.

### HasGrossSales

`func (o *AnalyticsProductMetric) HasGrossSales() bool`

HasGrossSales returns a boolean if a field has been set.

### SetGrossSalesNil

`func (o *AnalyticsProductMetric) SetGrossSalesNil(b bool)`

 SetGrossSalesNil sets the value for GrossSales to be an explicit nil

### UnsetGrossSales
`func (o *AnalyticsProductMetric) UnsetGrossSales()`

UnsetGrossSales ensures that no value is present for GrossSales, not even an explicit nil
### GetNetSales

`func (o *AnalyticsProductMetric) GetNetSales() float32`

GetNetSales returns the NetSales field if non-nil, zero value otherwise.

### GetNetSalesOk

`func (o *AnalyticsProductMetric) GetNetSalesOk() (*float32, bool)`

GetNetSalesOk returns a tuple with the NetSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetSales

`func (o *AnalyticsProductMetric) SetNetSales(v float32)`

SetNetSales sets NetSales field to given value.

### HasNetSales

`func (o *AnalyticsProductMetric) HasNetSales() bool`

HasNetSales returns a boolean if a field has been set.

### SetNetSalesNil

`func (o *AnalyticsProductMetric) SetNetSalesNil(b bool)`

 SetNetSalesNil sets the value for NetSales to be an explicit nil

### UnsetNetSales
`func (o *AnalyticsProductMetric) UnsetNetSales()`

UnsetNetSales ensures that no value is present for NetSales, not even an explicit nil
### GetDiscounts

`func (o *AnalyticsProductMetric) GetDiscounts() float32`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *AnalyticsProductMetric) GetDiscountsOk() (*float32, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *AnalyticsProductMetric) SetDiscounts(v float32)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *AnalyticsProductMetric) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### SetDiscountsNil

`func (o *AnalyticsProductMetric) SetDiscountsNil(b bool)`

 SetDiscountsNil sets the value for Discounts to be an explicit nil

### UnsetDiscounts
`func (o *AnalyticsProductMetric) UnsetDiscounts()`

UnsetDiscounts ensures that no value is present for Discounts, not even an explicit nil
### GetRefunds

`func (o *AnalyticsProductMetric) GetRefunds() float32`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *AnalyticsProductMetric) GetRefundsOk() (*float32, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *AnalyticsProductMetric) SetRefunds(v float32)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *AnalyticsProductMetric) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### SetRefundsNil

`func (o *AnalyticsProductMetric) SetRefundsNil(b bool)`

 SetRefundsNil sets the value for Refunds to be an explicit nil

### UnsetRefunds
`func (o *AnalyticsProductMetric) UnsetRefunds()`

UnsetRefunds ensures that no value is present for Refunds, not even an explicit nil
### GetCurrencyId

`func (o *AnalyticsProductMetric) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AnalyticsProductMetric) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AnalyticsProductMetric) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AnalyticsProductMetric) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *AnalyticsProductMetric) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *AnalyticsProductMetric) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetAdditionalFields

`func (o *AnalyticsProductMetric) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *AnalyticsProductMetric) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *AnalyticsProductMetric) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *AnalyticsProductMetric) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *AnalyticsProductMetric) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *AnalyticsProductMetric) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *AnalyticsProductMetric) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AnalyticsProductMetric) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AnalyticsProductMetric) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AnalyticsProductMetric) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *AnalyticsProductMetric) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *AnalyticsProductMetric) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


