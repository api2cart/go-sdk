# ProductDeleteBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | [**[]ProductDeleteBatchPayloadInner**](ProductDeleteBatchPayloadInner.md) | Contains an array of product deletion requests, each including the product ID. | 

## Methods

### NewProductDeleteBatch

`func NewProductDeleteBatch(payload []ProductDeleteBatchPayloadInner, ) *ProductDeleteBatch`

NewProductDeleteBatch instantiates a new ProductDeleteBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDeleteBatchWithDefaults

`func NewProductDeleteBatchWithDefaults() *ProductDeleteBatch`

NewProductDeleteBatchWithDefaults instantiates a new ProductDeleteBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *ProductDeleteBatch) GetPayload() []ProductDeleteBatchPayloadInner`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ProductDeleteBatch) GetPayloadOk() (*[]ProductDeleteBatchPayloadInner, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ProductDeleteBatch) SetPayload(v []ProductDeleteBatchPayloadInner)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


