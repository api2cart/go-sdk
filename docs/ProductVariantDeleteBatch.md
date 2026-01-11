# ProductVariantDeleteBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClearCache** | Pointer to **bool** |  | [optional] [default to false]
**Reindex** | Pointer to **bool** |  | [optional] [default to false]
**Payload** | [**[]ProductVariantDeleteBatchPayloadInner**](ProductVariantDeleteBatchPayloadInner.md) | Contains an array of product variant deletion requests, each including the product ID and variant ID. The list of properties may vary depending on the specific platform. | 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 

## Methods

### NewProductVariantDeleteBatch

`func NewProductVariantDeleteBatch(payload []ProductVariantDeleteBatchPayloadInner, ) *ProductVariantDeleteBatch`

NewProductVariantDeleteBatch instantiates a new ProductVariantDeleteBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariantDeleteBatchWithDefaults

`func NewProductVariantDeleteBatchWithDefaults() *ProductVariantDeleteBatch`

NewProductVariantDeleteBatchWithDefaults instantiates a new ProductVariantDeleteBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClearCache

`func (o *ProductVariantDeleteBatch) GetClearCache() bool`

GetClearCache returns the ClearCache field if non-nil, zero value otherwise.

### GetClearCacheOk

`func (o *ProductVariantDeleteBatch) GetClearCacheOk() (*bool, bool)`

GetClearCacheOk returns a tuple with the ClearCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCache

`func (o *ProductVariantDeleteBatch) SetClearCache(v bool)`

SetClearCache sets ClearCache field to given value.

### HasClearCache

`func (o *ProductVariantDeleteBatch) HasClearCache() bool`

HasClearCache returns a boolean if a field has been set.

### GetReindex

`func (o *ProductVariantDeleteBatch) GetReindex() bool`

GetReindex returns the Reindex field if non-nil, zero value otherwise.

### GetReindexOk

`func (o *ProductVariantDeleteBatch) GetReindexOk() (*bool, bool)`

GetReindexOk returns a tuple with the Reindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReindex

`func (o *ProductVariantDeleteBatch) SetReindex(v bool)`

SetReindex sets Reindex field to given value.

### HasReindex

`func (o *ProductVariantDeleteBatch) HasReindex() bool`

HasReindex returns a boolean if a field has been set.

### GetPayload

`func (o *ProductVariantDeleteBatch) GetPayload() []ProductVariantDeleteBatchPayloadInner`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ProductVariantDeleteBatch) GetPayloadOk() (*[]ProductVariantDeleteBatchPayloadInner, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ProductVariantDeleteBatch) SetPayload(v []ProductVariantDeleteBatchPayloadInner)`

SetPayload sets Payload field to given value.


### GetIdempotencyKey

`func (o *ProductVariantDeleteBatch) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ProductVariantDeleteBatch) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ProductVariantDeleteBatch) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *ProductVariantDeleteBatch) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


