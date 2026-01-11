# ProductVariantPriceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Defines the variant where the price has to be updated | [optional] 
**ProductId** | Pointer to **string** | Product id | [optional] 
**GroupPrices** | [**[]ProductPriceUpdateGroupPricesInner**](ProductPriceUpdateGroupPricesInner.md) | Defines variants&#39;s group prices | 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 

## Methods

### NewProductVariantPriceUpdate

`func NewProductVariantPriceUpdate(groupPrices []ProductPriceUpdateGroupPricesInner, ) *ProductVariantPriceUpdate`

NewProductVariantPriceUpdate instantiates a new ProductVariantPriceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariantPriceUpdateWithDefaults

`func NewProductVariantPriceUpdateWithDefaults() *ProductVariantPriceUpdate`

NewProductVariantPriceUpdateWithDefaults instantiates a new ProductVariantPriceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductVariantPriceUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductVariantPriceUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductVariantPriceUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductVariantPriceUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *ProductVariantPriceUpdate) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductVariantPriceUpdate) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductVariantPriceUpdate) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductVariantPriceUpdate) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetGroupPrices

`func (o *ProductVariantPriceUpdate) GetGroupPrices() []ProductPriceUpdateGroupPricesInner`

GetGroupPrices returns the GroupPrices field if non-nil, zero value otherwise.

### GetGroupPricesOk

`func (o *ProductVariantPriceUpdate) GetGroupPricesOk() (*[]ProductPriceUpdateGroupPricesInner, bool)`

GetGroupPricesOk returns a tuple with the GroupPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPrices

`func (o *ProductVariantPriceUpdate) SetGroupPrices(v []ProductPriceUpdateGroupPricesInner)`

SetGroupPrices sets GroupPrices field to given value.


### GetIdempotencyKey

`func (o *ProductVariantPriceUpdate) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ProductVariantPriceUpdate) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ProductVariantPriceUpdate) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *ProductVariantPriceUpdate) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


