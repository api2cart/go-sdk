# AnalyticsMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrdersCount** | Pointer to **NullableInt32** |  | [optional] 
**ItemsSold** | Pointer to **NullableInt32** |  | [optional] 
**GrossSales** | Pointer to **NullableFloat32** |  | [optional] 
**NetSales** | Pointer to **NullableFloat32** |  | [optional] 
**TotalSales** | Pointer to **NullableFloat32** |  | [optional] 
**AvgOrderValue** | Pointer to **NullableFloat32** |  | [optional] 
**Discounts** | Pointer to **NullableFloat32** |  | [optional] 
**Refunds** | Pointer to **NullableFloat32** |  | [optional] 
**Tax** | Pointer to **NullableFloat32** |  | [optional] 
**Shipping** | Pointer to **NullableFloat32** |  | [optional] 
**Fees** | Pointer to **NullableFloat32** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAnalyticsMetric

`func NewAnalyticsMetric() *AnalyticsMetric`

NewAnalyticsMetric instantiates a new AnalyticsMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsMetricWithDefaults

`func NewAnalyticsMetricWithDefaults() *AnalyticsMetric`

NewAnalyticsMetricWithDefaults instantiates a new AnalyticsMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrdersCount

`func (o *AnalyticsMetric) GetOrdersCount() int32`

GetOrdersCount returns the OrdersCount field if non-nil, zero value otherwise.

### GetOrdersCountOk

`func (o *AnalyticsMetric) GetOrdersCountOk() (*int32, bool)`

GetOrdersCountOk returns a tuple with the OrdersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersCount

`func (o *AnalyticsMetric) SetOrdersCount(v int32)`

SetOrdersCount sets OrdersCount field to given value.

### HasOrdersCount

`func (o *AnalyticsMetric) HasOrdersCount() bool`

HasOrdersCount returns a boolean if a field has been set.

### SetOrdersCountNil

`func (o *AnalyticsMetric) SetOrdersCountNil(b bool)`

 SetOrdersCountNil sets the value for OrdersCount to be an explicit nil

### UnsetOrdersCount
`func (o *AnalyticsMetric) UnsetOrdersCount()`

UnsetOrdersCount ensures that no value is present for OrdersCount, not even an explicit nil
### GetItemsSold

`func (o *AnalyticsMetric) GetItemsSold() int32`

GetItemsSold returns the ItemsSold field if non-nil, zero value otherwise.

### GetItemsSoldOk

`func (o *AnalyticsMetric) GetItemsSoldOk() (*int32, bool)`

GetItemsSoldOk returns a tuple with the ItemsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsSold

`func (o *AnalyticsMetric) SetItemsSold(v int32)`

SetItemsSold sets ItemsSold field to given value.

### HasItemsSold

`func (o *AnalyticsMetric) HasItemsSold() bool`

HasItemsSold returns a boolean if a field has been set.

### SetItemsSoldNil

`func (o *AnalyticsMetric) SetItemsSoldNil(b bool)`

 SetItemsSoldNil sets the value for ItemsSold to be an explicit nil

### UnsetItemsSold
`func (o *AnalyticsMetric) UnsetItemsSold()`

UnsetItemsSold ensures that no value is present for ItemsSold, not even an explicit nil
### GetGrossSales

`func (o *AnalyticsMetric) GetGrossSales() float32`

GetGrossSales returns the GrossSales field if non-nil, zero value otherwise.

### GetGrossSalesOk

`func (o *AnalyticsMetric) GetGrossSalesOk() (*float32, bool)`

GetGrossSalesOk returns a tuple with the GrossSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossSales

`func (o *AnalyticsMetric) SetGrossSales(v float32)`

SetGrossSales sets GrossSales field to given value.

### HasGrossSales

`func (o *AnalyticsMetric) HasGrossSales() bool`

HasGrossSales returns a boolean if a field has been set.

### SetGrossSalesNil

`func (o *AnalyticsMetric) SetGrossSalesNil(b bool)`

 SetGrossSalesNil sets the value for GrossSales to be an explicit nil

### UnsetGrossSales
`func (o *AnalyticsMetric) UnsetGrossSales()`

UnsetGrossSales ensures that no value is present for GrossSales, not even an explicit nil
### GetNetSales

`func (o *AnalyticsMetric) GetNetSales() float32`

GetNetSales returns the NetSales field if non-nil, zero value otherwise.

### GetNetSalesOk

`func (o *AnalyticsMetric) GetNetSalesOk() (*float32, bool)`

GetNetSalesOk returns a tuple with the NetSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetSales

`func (o *AnalyticsMetric) SetNetSales(v float32)`

SetNetSales sets NetSales field to given value.

### HasNetSales

`func (o *AnalyticsMetric) HasNetSales() bool`

HasNetSales returns a boolean if a field has been set.

### SetNetSalesNil

`func (o *AnalyticsMetric) SetNetSalesNil(b bool)`

 SetNetSalesNil sets the value for NetSales to be an explicit nil

### UnsetNetSales
`func (o *AnalyticsMetric) UnsetNetSales()`

UnsetNetSales ensures that no value is present for NetSales, not even an explicit nil
### GetTotalSales

`func (o *AnalyticsMetric) GetTotalSales() float32`

GetTotalSales returns the TotalSales field if non-nil, zero value otherwise.

### GetTotalSalesOk

`func (o *AnalyticsMetric) GetTotalSalesOk() (*float32, bool)`

GetTotalSalesOk returns a tuple with the TotalSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSales

`func (o *AnalyticsMetric) SetTotalSales(v float32)`

SetTotalSales sets TotalSales field to given value.

### HasTotalSales

`func (o *AnalyticsMetric) HasTotalSales() bool`

HasTotalSales returns a boolean if a field has been set.

### SetTotalSalesNil

`func (o *AnalyticsMetric) SetTotalSalesNil(b bool)`

 SetTotalSalesNil sets the value for TotalSales to be an explicit nil

### UnsetTotalSales
`func (o *AnalyticsMetric) UnsetTotalSales()`

UnsetTotalSales ensures that no value is present for TotalSales, not even an explicit nil
### GetAvgOrderValue

`func (o *AnalyticsMetric) GetAvgOrderValue() float32`

GetAvgOrderValue returns the AvgOrderValue field if non-nil, zero value otherwise.

### GetAvgOrderValueOk

`func (o *AnalyticsMetric) GetAvgOrderValueOk() (*float32, bool)`

GetAvgOrderValueOk returns a tuple with the AvgOrderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgOrderValue

`func (o *AnalyticsMetric) SetAvgOrderValue(v float32)`

SetAvgOrderValue sets AvgOrderValue field to given value.

### HasAvgOrderValue

`func (o *AnalyticsMetric) HasAvgOrderValue() bool`

HasAvgOrderValue returns a boolean if a field has been set.

### SetAvgOrderValueNil

`func (o *AnalyticsMetric) SetAvgOrderValueNil(b bool)`

 SetAvgOrderValueNil sets the value for AvgOrderValue to be an explicit nil

### UnsetAvgOrderValue
`func (o *AnalyticsMetric) UnsetAvgOrderValue()`

UnsetAvgOrderValue ensures that no value is present for AvgOrderValue, not even an explicit nil
### GetDiscounts

`func (o *AnalyticsMetric) GetDiscounts() float32`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *AnalyticsMetric) GetDiscountsOk() (*float32, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *AnalyticsMetric) SetDiscounts(v float32)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *AnalyticsMetric) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### SetDiscountsNil

`func (o *AnalyticsMetric) SetDiscountsNil(b bool)`

 SetDiscountsNil sets the value for Discounts to be an explicit nil

### UnsetDiscounts
`func (o *AnalyticsMetric) UnsetDiscounts()`

UnsetDiscounts ensures that no value is present for Discounts, not even an explicit nil
### GetRefunds

`func (o *AnalyticsMetric) GetRefunds() float32`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *AnalyticsMetric) GetRefundsOk() (*float32, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *AnalyticsMetric) SetRefunds(v float32)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *AnalyticsMetric) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### SetRefundsNil

`func (o *AnalyticsMetric) SetRefundsNil(b bool)`

 SetRefundsNil sets the value for Refunds to be an explicit nil

### UnsetRefunds
`func (o *AnalyticsMetric) UnsetRefunds()`

UnsetRefunds ensures that no value is present for Refunds, not even an explicit nil
### GetTax

`func (o *AnalyticsMetric) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *AnalyticsMetric) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *AnalyticsMetric) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *AnalyticsMetric) HasTax() bool`

HasTax returns a boolean if a field has been set.

### SetTaxNil

`func (o *AnalyticsMetric) SetTaxNil(b bool)`

 SetTaxNil sets the value for Tax to be an explicit nil

### UnsetTax
`func (o *AnalyticsMetric) UnsetTax()`

UnsetTax ensures that no value is present for Tax, not even an explicit nil
### GetShipping

`func (o *AnalyticsMetric) GetShipping() float32`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *AnalyticsMetric) GetShippingOk() (*float32, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *AnalyticsMetric) SetShipping(v float32)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *AnalyticsMetric) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### SetShippingNil

`func (o *AnalyticsMetric) SetShippingNil(b bool)`

 SetShippingNil sets the value for Shipping to be an explicit nil

### UnsetShipping
`func (o *AnalyticsMetric) UnsetShipping()`

UnsetShipping ensures that no value is present for Shipping, not even an explicit nil
### GetFees

`func (o *AnalyticsMetric) GetFees() float32`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *AnalyticsMetric) GetFeesOk() (*float32, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *AnalyticsMetric) SetFees(v float32)`

SetFees sets Fees field to given value.

### HasFees

`func (o *AnalyticsMetric) HasFees() bool`

HasFees returns a boolean if a field has been set.

### SetFeesNil

`func (o *AnalyticsMetric) SetFeesNil(b bool)`

 SetFeesNil sets the value for Fees to be an explicit nil

### UnsetFees
`func (o *AnalyticsMetric) UnsetFees()`

UnsetFees ensures that no value is present for Fees, not even an explicit nil
### GetCurrencyId

`func (o *AnalyticsMetric) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AnalyticsMetric) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AnalyticsMetric) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AnalyticsMetric) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *AnalyticsMetric) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *AnalyticsMetric) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetAdditionalFields

`func (o *AnalyticsMetric) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *AnalyticsMetric) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *AnalyticsMetric) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *AnalyticsMetric) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *AnalyticsMetric) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *AnalyticsMetric) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *AnalyticsMetric) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AnalyticsMetric) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AnalyticsMetric) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AnalyticsMetric) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *AnalyticsMetric) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *AnalyticsMetric) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


