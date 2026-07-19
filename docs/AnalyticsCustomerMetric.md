# AnalyticsCustomerMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrdersCount** | Pointer to **NullableInt32** |  | [optional] 
**TotalSpend** | Pointer to **NullableFloat32** |  | [optional] 
**AvgOrderValue** | Pointer to **NullableFloat32** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAnalyticsCustomerMetric

`func NewAnalyticsCustomerMetric() *AnalyticsCustomerMetric`

NewAnalyticsCustomerMetric instantiates a new AnalyticsCustomerMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsCustomerMetricWithDefaults

`func NewAnalyticsCustomerMetricWithDefaults() *AnalyticsCustomerMetric`

NewAnalyticsCustomerMetricWithDefaults instantiates a new AnalyticsCustomerMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrdersCount

`func (o *AnalyticsCustomerMetric) GetOrdersCount() int32`

GetOrdersCount returns the OrdersCount field if non-nil, zero value otherwise.

### GetOrdersCountOk

`func (o *AnalyticsCustomerMetric) GetOrdersCountOk() (*int32, bool)`

GetOrdersCountOk returns a tuple with the OrdersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersCount

`func (o *AnalyticsCustomerMetric) SetOrdersCount(v int32)`

SetOrdersCount sets OrdersCount field to given value.

### HasOrdersCount

`func (o *AnalyticsCustomerMetric) HasOrdersCount() bool`

HasOrdersCount returns a boolean if a field has been set.

### SetOrdersCountNil

`func (o *AnalyticsCustomerMetric) SetOrdersCountNil(b bool)`

 SetOrdersCountNil sets the value for OrdersCount to be an explicit nil

### UnsetOrdersCount
`func (o *AnalyticsCustomerMetric) UnsetOrdersCount()`

UnsetOrdersCount ensures that no value is present for OrdersCount, not even an explicit nil
### GetTotalSpend

`func (o *AnalyticsCustomerMetric) GetTotalSpend() float32`

GetTotalSpend returns the TotalSpend field if non-nil, zero value otherwise.

### GetTotalSpendOk

`func (o *AnalyticsCustomerMetric) GetTotalSpendOk() (*float32, bool)`

GetTotalSpendOk returns a tuple with the TotalSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpend

`func (o *AnalyticsCustomerMetric) SetTotalSpend(v float32)`

SetTotalSpend sets TotalSpend field to given value.

### HasTotalSpend

`func (o *AnalyticsCustomerMetric) HasTotalSpend() bool`

HasTotalSpend returns a boolean if a field has been set.

### SetTotalSpendNil

`func (o *AnalyticsCustomerMetric) SetTotalSpendNil(b bool)`

 SetTotalSpendNil sets the value for TotalSpend to be an explicit nil

### UnsetTotalSpend
`func (o *AnalyticsCustomerMetric) UnsetTotalSpend()`

UnsetTotalSpend ensures that no value is present for TotalSpend, not even an explicit nil
### GetAvgOrderValue

`func (o *AnalyticsCustomerMetric) GetAvgOrderValue() float32`

GetAvgOrderValue returns the AvgOrderValue field if non-nil, zero value otherwise.

### GetAvgOrderValueOk

`func (o *AnalyticsCustomerMetric) GetAvgOrderValueOk() (*float32, bool)`

GetAvgOrderValueOk returns a tuple with the AvgOrderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgOrderValue

`func (o *AnalyticsCustomerMetric) SetAvgOrderValue(v float32)`

SetAvgOrderValue sets AvgOrderValue field to given value.

### HasAvgOrderValue

`func (o *AnalyticsCustomerMetric) HasAvgOrderValue() bool`

HasAvgOrderValue returns a boolean if a field has been set.

### SetAvgOrderValueNil

`func (o *AnalyticsCustomerMetric) SetAvgOrderValueNil(b bool)`

 SetAvgOrderValueNil sets the value for AvgOrderValue to be an explicit nil

### UnsetAvgOrderValue
`func (o *AnalyticsCustomerMetric) UnsetAvgOrderValue()`

UnsetAvgOrderValue ensures that no value is present for AvgOrderValue, not even an explicit nil
### GetCurrencyId

`func (o *AnalyticsCustomerMetric) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AnalyticsCustomerMetric) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AnalyticsCustomerMetric) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AnalyticsCustomerMetric) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *AnalyticsCustomerMetric) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *AnalyticsCustomerMetric) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetAdditionalFields

`func (o *AnalyticsCustomerMetric) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *AnalyticsCustomerMetric) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *AnalyticsCustomerMetric) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *AnalyticsCustomerMetric) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *AnalyticsCustomerMetric) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *AnalyticsCustomerMetric) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *AnalyticsCustomerMetric) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AnalyticsCustomerMetric) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AnalyticsCustomerMetric) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AnalyticsCustomerMetric) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *AnalyticsCustomerMetric) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *AnalyticsCustomerMetric) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


