# ProductAddBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NestedItemsUpdateBehaviour** | Pointer to **string** |  Determines how updates to nested items should be handled.&lt;hr&gt;&lt;div style&#x3D;\&quot;font-style:normal\&quot;&gt;  Values description:  &lt;div style&#x3D;\&quot;margin-left: 2%; padding-top: 2%\&quot;&gt;    &lt;div style&#x3D;\&quot;font-size:85%\&quot;&gt;      &lt;b&gt;  replace&lt;/b&gt;: This option indicates that the nested items should be completely replaced with the new data provided. &lt;/br&gt;      &lt;b&gt;  merge&lt;/b&gt;: With this option, updates to nested items are merged with the existing data. &lt;/br&gt;    &lt;/div&gt;  &lt;/div&gt;&lt;/div&gt; | [optional] [default to "replace"]
**ClearCache** | Pointer to **bool** |  | [optional] [default to false]
**Reindex** | Pointer to **bool** |  | [optional] [default to false]
**Payload** | [**[]ProductAddBatchPayloadInner**](ProductAddBatchPayloadInner.md) | Contains an array of product objects. The list of properties may vary depending on the specific platform. | 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 

## Methods

### NewProductAddBatch

`func NewProductAddBatch(payload []ProductAddBatchPayloadInner, ) *ProductAddBatch`

NewProductAddBatch instantiates a new ProductAddBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAddBatchWithDefaults

`func NewProductAddBatchWithDefaults() *ProductAddBatch`

NewProductAddBatchWithDefaults instantiates a new ProductAddBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNestedItemsUpdateBehaviour

`func (o *ProductAddBatch) GetNestedItemsUpdateBehaviour() string`

GetNestedItemsUpdateBehaviour returns the NestedItemsUpdateBehaviour field if non-nil, zero value otherwise.

### GetNestedItemsUpdateBehaviourOk

`func (o *ProductAddBatch) GetNestedItemsUpdateBehaviourOk() (*string, bool)`

GetNestedItemsUpdateBehaviourOk returns a tuple with the NestedItemsUpdateBehaviour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedItemsUpdateBehaviour

`func (o *ProductAddBatch) SetNestedItemsUpdateBehaviour(v string)`

SetNestedItemsUpdateBehaviour sets NestedItemsUpdateBehaviour field to given value.

### HasNestedItemsUpdateBehaviour

`func (o *ProductAddBatch) HasNestedItemsUpdateBehaviour() bool`

HasNestedItemsUpdateBehaviour returns a boolean if a field has been set.

### GetClearCache

`func (o *ProductAddBatch) GetClearCache() bool`

GetClearCache returns the ClearCache field if non-nil, zero value otherwise.

### GetClearCacheOk

`func (o *ProductAddBatch) GetClearCacheOk() (*bool, bool)`

GetClearCacheOk returns a tuple with the ClearCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCache

`func (o *ProductAddBatch) SetClearCache(v bool)`

SetClearCache sets ClearCache field to given value.

### HasClearCache

`func (o *ProductAddBatch) HasClearCache() bool`

HasClearCache returns a boolean if a field has been set.

### GetReindex

`func (o *ProductAddBatch) GetReindex() bool`

GetReindex returns the Reindex field if non-nil, zero value otherwise.

### GetReindexOk

`func (o *ProductAddBatch) GetReindexOk() (*bool, bool)`

GetReindexOk returns a tuple with the Reindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReindex

`func (o *ProductAddBatch) SetReindex(v bool)`

SetReindex sets Reindex field to given value.

### HasReindex

`func (o *ProductAddBatch) HasReindex() bool`

HasReindex returns a boolean if a field has been set.

### GetPayload

`func (o *ProductAddBatch) GetPayload() []ProductAddBatchPayloadInner`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ProductAddBatch) GetPayloadOk() (*[]ProductAddBatchPayloadInner, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ProductAddBatch) SetPayload(v []ProductAddBatchPayloadInner)`

SetPayload sets Payload field to given value.


### GetIdempotencyKey

`func (o *ProductAddBatch) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ProductAddBatch) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ProductAddBatch) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *ProductAddBatch) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


