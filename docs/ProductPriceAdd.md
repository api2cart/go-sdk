# ProductPriceAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** | Defines the product to which the price has to be added | [optional] 
**GroupPrices** | Pointer to [**[]ProductAddGroupPricesInner**](ProductAddGroupPricesInner.md) | Defines product&#39;s group prices | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 

## Methods

### NewProductPriceAdd

`func NewProductPriceAdd() *ProductPriceAdd`

NewProductPriceAdd instantiates a new ProductPriceAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductPriceAddWithDefaults

`func NewProductPriceAddWithDefaults() *ProductPriceAdd`

NewProductPriceAddWithDefaults instantiates a new ProductPriceAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProductPriceAdd) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductPriceAdd) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductPriceAdd) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductPriceAdd) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetGroupPrices

`func (o *ProductPriceAdd) GetGroupPrices() []ProductAddGroupPricesInner`

GetGroupPrices returns the GroupPrices field if non-nil, zero value otherwise.

### GetGroupPricesOk

`func (o *ProductPriceAdd) GetGroupPricesOk() (*[]ProductAddGroupPricesInner, bool)`

GetGroupPricesOk returns a tuple with the GroupPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPrices

`func (o *ProductPriceAdd) SetGroupPrices(v []ProductAddGroupPricesInner)`

SetGroupPrices sets GroupPrices field to given value.

### HasGroupPrices

`func (o *ProductPriceAdd) HasGroupPrices() bool`

HasGroupPrices returns a boolean if a field has been set.

### GetStoreId

`func (o *ProductPriceAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductPriceAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductPriceAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductPriceAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


