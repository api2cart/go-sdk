# ProductVariantPriceAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Defines the variant to which the price has to be added | [optional] 
**ProductId** | Pointer to **string** | Product id | [optional] 
**GroupPrices** | [**[]ProductAddGroupPricesInner**](ProductAddGroupPricesInner.md) | Defines variants&#39;s group prices | 
**StoreId** | Pointer to **string** | Store Id | [optional] 

## Methods

### NewProductVariantPriceAdd

`func NewProductVariantPriceAdd(groupPrices []ProductAddGroupPricesInner, ) *ProductVariantPriceAdd`

NewProductVariantPriceAdd instantiates a new ProductVariantPriceAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariantPriceAddWithDefaults

`func NewProductVariantPriceAddWithDefaults() *ProductVariantPriceAdd`

NewProductVariantPriceAddWithDefaults instantiates a new ProductVariantPriceAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductVariantPriceAdd) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductVariantPriceAdd) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductVariantPriceAdd) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductVariantPriceAdd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *ProductVariantPriceAdd) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductVariantPriceAdd) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductVariantPriceAdd) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductVariantPriceAdd) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetGroupPrices

`func (o *ProductVariantPriceAdd) GetGroupPrices() []ProductAddGroupPricesInner`

GetGroupPrices returns the GroupPrices field if non-nil, zero value otherwise.

### GetGroupPricesOk

`func (o *ProductVariantPriceAdd) GetGroupPricesOk() (*[]ProductAddGroupPricesInner, bool)`

GetGroupPricesOk returns a tuple with the GroupPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPrices

`func (o *ProductVariantPriceAdd) SetGroupPrices(v []ProductAddGroupPricesInner)`

SetGroupPrices sets GroupPrices field to given value.


### GetStoreId

`func (o *ProductVariantPriceAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductVariantPriceAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductVariantPriceAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductVariantPriceAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


