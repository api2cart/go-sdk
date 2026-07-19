# AnalyticsProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** |  | [optional] 
**VariantId** | Pointer to **NullableString** |  | [optional] 
**Sku** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CategoryIds** | Pointer to **[]string** |  | [optional] 
**Metrics** | Pointer to [**AnalyticsProductMetric**](AnalyticsProductMetric.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAnalyticsProduct

`func NewAnalyticsProduct() *AnalyticsProduct`

NewAnalyticsProduct instantiates a new AnalyticsProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsProductWithDefaults

`func NewAnalyticsProductWithDefaults() *AnalyticsProduct`

NewAnalyticsProductWithDefaults instantiates a new AnalyticsProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *AnalyticsProduct) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *AnalyticsProduct) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *AnalyticsProduct) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *AnalyticsProduct) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetVariantId

`func (o *AnalyticsProduct) GetVariantId() string`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *AnalyticsProduct) GetVariantIdOk() (*string, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *AnalyticsProduct) SetVariantId(v string)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *AnalyticsProduct) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.

### SetVariantIdNil

`func (o *AnalyticsProduct) SetVariantIdNil(b bool)`

 SetVariantIdNil sets the value for VariantId to be an explicit nil

### UnsetVariantId
`func (o *AnalyticsProduct) UnsetVariantId()`

UnsetVariantId ensures that no value is present for VariantId, not even an explicit nil
### GetSku

`func (o *AnalyticsProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *AnalyticsProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *AnalyticsProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *AnalyticsProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *AnalyticsProduct) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *AnalyticsProduct) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetName

`func (o *AnalyticsProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnalyticsProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnalyticsProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnalyticsProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AnalyticsProduct) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AnalyticsProduct) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCategoryIds

`func (o *AnalyticsProduct) GetCategoryIds() []string`

GetCategoryIds returns the CategoryIds field if non-nil, zero value otherwise.

### GetCategoryIdsOk

`func (o *AnalyticsProduct) GetCategoryIdsOk() (*[]string, bool)`

GetCategoryIdsOk returns a tuple with the CategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryIds

`func (o *AnalyticsProduct) SetCategoryIds(v []string)`

SetCategoryIds sets CategoryIds field to given value.

### HasCategoryIds

`func (o *AnalyticsProduct) HasCategoryIds() bool`

HasCategoryIds returns a boolean if a field has been set.

### GetMetrics

`func (o *AnalyticsProduct) GetMetrics() AnalyticsProductMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AnalyticsProduct) GetMetricsOk() (*AnalyticsProductMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AnalyticsProduct) SetMetrics(v AnalyticsProductMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AnalyticsProduct) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *AnalyticsProduct) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *AnalyticsProduct) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *AnalyticsProduct) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *AnalyticsProduct) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *AnalyticsProduct) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *AnalyticsProduct) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *AnalyticsProduct) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AnalyticsProduct) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AnalyticsProduct) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AnalyticsProduct) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *AnalyticsProduct) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *AnalyticsProduct) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


