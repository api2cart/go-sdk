# CategoryDeleteBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | [**[]CategoryDeleteBatchPayloadInner**](CategoryDeleteBatchPayloadInner.md) | Contains an array of category IDs to delete. | 

## Methods

### NewCategoryDeleteBatch

`func NewCategoryDeleteBatch(payload []CategoryDeleteBatchPayloadInner, ) *CategoryDeleteBatch`

NewCategoryDeleteBatch instantiates a new CategoryDeleteBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryDeleteBatchWithDefaults

`func NewCategoryDeleteBatchWithDefaults() *CategoryDeleteBatch`

NewCategoryDeleteBatchWithDefaults instantiates a new CategoryDeleteBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *CategoryDeleteBatch) GetPayload() []CategoryDeleteBatchPayloadInner`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CategoryDeleteBatch) GetPayloadOk() (*[]CategoryDeleteBatchPayloadInner, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CategoryDeleteBatch) SetPayload(v []CategoryDeleteBatchPayloadInner)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


