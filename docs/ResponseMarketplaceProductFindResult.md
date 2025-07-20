# ResponseMarketplaceProductFindResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceProductsCount** | Pointer to **NullableInt32** |  | [optional] 
**MarketplaceProduct** | Pointer to [**[]MarketplaceProduct**](MarketplaceProduct.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseMarketplaceProductFindResult

`func NewResponseMarketplaceProductFindResult() *ResponseMarketplaceProductFindResult`

NewResponseMarketplaceProductFindResult instantiates a new ResponseMarketplaceProductFindResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMarketplaceProductFindResultWithDefaults

`func NewResponseMarketplaceProductFindResultWithDefaults() *ResponseMarketplaceProductFindResult`

NewResponseMarketplaceProductFindResultWithDefaults instantiates a new ResponseMarketplaceProductFindResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceProductsCount

`func (o *ResponseMarketplaceProductFindResult) GetMarketplaceProductsCount() int32`

GetMarketplaceProductsCount returns the MarketplaceProductsCount field if non-nil, zero value otherwise.

### GetMarketplaceProductsCountOk

`func (o *ResponseMarketplaceProductFindResult) GetMarketplaceProductsCountOk() (*int32, bool)`

GetMarketplaceProductsCountOk returns a tuple with the MarketplaceProductsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceProductsCount

`func (o *ResponseMarketplaceProductFindResult) SetMarketplaceProductsCount(v int32)`

SetMarketplaceProductsCount sets MarketplaceProductsCount field to given value.

### HasMarketplaceProductsCount

`func (o *ResponseMarketplaceProductFindResult) HasMarketplaceProductsCount() bool`

HasMarketplaceProductsCount returns a boolean if a field has been set.

### SetMarketplaceProductsCountNil

`func (o *ResponseMarketplaceProductFindResult) SetMarketplaceProductsCountNil(b bool)`

 SetMarketplaceProductsCountNil sets the value for MarketplaceProductsCount to be an explicit nil

### UnsetMarketplaceProductsCount
`func (o *ResponseMarketplaceProductFindResult) UnsetMarketplaceProductsCount()`

UnsetMarketplaceProductsCount ensures that no value is present for MarketplaceProductsCount, not even an explicit nil
### GetMarketplaceProduct

`func (o *ResponseMarketplaceProductFindResult) GetMarketplaceProduct() []MarketplaceProduct`

GetMarketplaceProduct returns the MarketplaceProduct field if non-nil, zero value otherwise.

### GetMarketplaceProductOk

`func (o *ResponseMarketplaceProductFindResult) GetMarketplaceProductOk() (*[]MarketplaceProduct, bool)`

GetMarketplaceProductOk returns a tuple with the MarketplaceProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceProduct

`func (o *ResponseMarketplaceProductFindResult) SetMarketplaceProduct(v []MarketplaceProduct)`

SetMarketplaceProduct sets MarketplaceProduct field to given value.

### HasMarketplaceProduct

`func (o *ResponseMarketplaceProductFindResult) HasMarketplaceProduct() bool`

HasMarketplaceProduct returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseMarketplaceProductFindResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseMarketplaceProductFindResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseMarketplaceProductFindResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseMarketplaceProductFindResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseMarketplaceProductFindResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseMarketplaceProductFindResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseMarketplaceProductFindResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseMarketplaceProductFindResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseMarketplaceProductFindResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseMarketplaceProductFindResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseMarketplaceProductFindResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseMarketplaceProductFindResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


