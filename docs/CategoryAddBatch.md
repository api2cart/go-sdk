# CategoryAddBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | [**[]CategoryAddBatchPayloadInner**](CategoryAddBatchPayloadInner.md) | Contains an array of categories objects. The list of properties may vary depending on the specific platform. | 

## Methods

### NewCategoryAddBatch

`func NewCategoryAddBatch(payload []CategoryAddBatchPayloadInner, ) *CategoryAddBatch`

NewCategoryAddBatch instantiates a new CategoryAddBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryAddBatchWithDefaults

`func NewCategoryAddBatchWithDefaults() *CategoryAddBatch`

NewCategoryAddBatchWithDefaults instantiates a new CategoryAddBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *CategoryAddBatch) GetPayload() []CategoryAddBatchPayloadInner`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CategoryAddBatch) GetPayloadOk() (*[]CategoryAddBatchPayloadInner, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CategoryAddBatch) SetPayload(v []CategoryAddBatchPayloadInner)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


