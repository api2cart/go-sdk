# ProductPriceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** | Defines the product where the price has to be updated | [optional] 
**GroupPrices** | Pointer to [**[]ProductPriceUpdateGroupPricesInner**](ProductPriceUpdateGroupPricesInner.md) | Defines product&#39;s group prices | [optional] 

## Methods

### NewProductPriceUpdate

`func NewProductPriceUpdate() *ProductPriceUpdate`

NewProductPriceUpdate instantiates a new ProductPriceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductPriceUpdateWithDefaults

`func NewProductPriceUpdateWithDefaults() *ProductPriceUpdate`

NewProductPriceUpdateWithDefaults instantiates a new ProductPriceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProductPriceUpdate) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductPriceUpdate) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductPriceUpdate) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductPriceUpdate) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetGroupPrices

`func (o *ProductPriceUpdate) GetGroupPrices() []ProductPriceUpdateGroupPricesInner`

GetGroupPrices returns the GroupPrices field if non-nil, zero value otherwise.

### GetGroupPricesOk

`func (o *ProductPriceUpdate) GetGroupPricesOk() (*[]ProductPriceUpdateGroupPricesInner, bool)`

GetGroupPricesOk returns a tuple with the GroupPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPrices

`func (o *ProductPriceUpdate) SetGroupPrices(v []ProductPriceUpdateGroupPricesInner)`

SetGroupPrices sets GroupPrices field to given value.

### HasGroupPrices

`func (o *ProductPriceUpdate) HasGroupPrices() bool`

HasGroupPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


