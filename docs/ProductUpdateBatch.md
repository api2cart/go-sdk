# ProductUpdateBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NestedItemsUpdateBehaviour** | Pointer to **string** |  Determines how updates to nested items should be handled.&lt;hr&gt;&lt;div style&#x3D;\&quot;font-style:normal\&quot;&gt;  Values description:  &lt;div style&#x3D;\&quot;margin-left: 2%; padding-top: 2%\&quot;&gt;    &lt;div style&#x3D;\&quot;font-size:85%\&quot;&gt;      &lt;b&gt;  replace&lt;/b&gt;: This option indicates that the nested items should be completely replaced with the new data provided. &lt;/br&gt;      &lt;b&gt;  merge&lt;/b&gt;: With this option, updates to nested items are merged with the existing data. &lt;/br&gt;    &lt;/div&gt;  &lt;/div&gt;&lt;/div&gt; | [optional] [default to "replace"]
**ClearCache** | Pointer to **bool** |  | [optional] [default to false]
**Reindex** | Pointer to **bool** |  | [optional] [default to false]
**Payload** | [**[]ProductUpdateBatchPayloadInner**](ProductUpdateBatchPayloadInner.md) | Contains an array of product objects. The list of properties may vary depending on the specific platform. | 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 

## Methods

### NewProductUpdateBatch

`func NewProductUpdateBatch(payload []ProductUpdateBatchPayloadInner, ) *ProductUpdateBatch`

NewProductUpdateBatch instantiates a new ProductUpdateBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductUpdateBatchWithDefaults

`func NewProductUpdateBatchWithDefaults() *ProductUpdateBatch`

NewProductUpdateBatchWithDefaults instantiates a new ProductUpdateBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNestedItemsUpdateBehaviour

`func (o *ProductUpdateBatch) GetNestedItemsUpdateBehaviour() string`

GetNestedItemsUpdateBehaviour returns the NestedItemsUpdateBehaviour field if non-nil, zero value otherwise.

### GetNestedItemsUpdateBehaviourOk

`func (o *ProductUpdateBatch) GetNestedItemsUpdateBehaviourOk() (*string, bool)`

GetNestedItemsUpdateBehaviourOk returns a tuple with the NestedItemsUpdateBehaviour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedItemsUpdateBehaviour

`func (o *ProductUpdateBatch) SetNestedItemsUpdateBehaviour(v string)`

SetNestedItemsUpdateBehaviour sets NestedItemsUpdateBehaviour field to given value.

### HasNestedItemsUpdateBehaviour

`func (o *ProductUpdateBatch) HasNestedItemsUpdateBehaviour() bool`

HasNestedItemsUpdateBehaviour returns a boolean if a field has been set.

### GetClearCache

`func (o *ProductUpdateBatch) GetClearCache() bool`

GetClearCache returns the ClearCache field if non-nil, zero value otherwise.

### GetClearCacheOk

`func (o *ProductUpdateBatch) GetClearCacheOk() (*bool, bool)`

GetClearCacheOk returns a tuple with the ClearCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCache

`func (o *ProductUpdateBatch) SetClearCache(v bool)`

SetClearCache sets ClearCache field to given value.

### HasClearCache

`func (o *ProductUpdateBatch) HasClearCache() bool`

HasClearCache returns a boolean if a field has been set.

### GetReindex

`func (o *ProductUpdateBatch) GetReindex() bool`

GetReindex returns the Reindex field if non-nil, zero value otherwise.

### GetReindexOk

`func (o *ProductUpdateBatch) GetReindexOk() (*bool, bool)`

GetReindexOk returns a tuple with the Reindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReindex

`func (o *ProductUpdateBatch) SetReindex(v bool)`

SetReindex sets Reindex field to given value.

### HasReindex

`func (o *ProductUpdateBatch) HasReindex() bool`

HasReindex returns a boolean if a field has been set.

### GetPayload

`func (o *ProductUpdateBatch) GetPayload() []ProductUpdateBatchPayloadInner`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ProductUpdateBatch) GetPayloadOk() (*[]ProductUpdateBatchPayloadInner, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ProductUpdateBatch) SetPayload(v []ProductUpdateBatchPayloadInner)`

SetPayload sets Payload field to given value.


### GetIdempotencyKey

`func (o *ProductUpdateBatch) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ProductUpdateBatch) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ProductUpdateBatch) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *ProductUpdateBatch) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


