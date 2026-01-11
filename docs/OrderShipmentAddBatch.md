# OrderShipmentAddBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | [**[]OrderShipmentAddBatchPayloadInner**](OrderShipmentAddBatchPayloadInner.md) | Contains an array of order shipment objects. The list of properties may vary depending on the specific platform. | 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 

## Methods

### NewOrderShipmentAddBatch

`func NewOrderShipmentAddBatch(payload []OrderShipmentAddBatchPayloadInner, ) *OrderShipmentAddBatch`

NewOrderShipmentAddBatch instantiates a new OrderShipmentAddBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderShipmentAddBatchWithDefaults

`func NewOrderShipmentAddBatchWithDefaults() *OrderShipmentAddBatch`

NewOrderShipmentAddBatchWithDefaults instantiates a new OrderShipmentAddBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *OrderShipmentAddBatch) GetPayload() []OrderShipmentAddBatchPayloadInner`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *OrderShipmentAddBatch) GetPayloadOk() (*[]OrderShipmentAddBatchPayloadInner, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *OrderShipmentAddBatch) SetPayload(v []OrderShipmentAddBatchPayloadInner)`

SetPayload sets Payload field to given value.


### GetIdempotencyKey

`func (o *OrderShipmentAddBatch) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *OrderShipmentAddBatch) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *OrderShipmentAddBatch) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *OrderShipmentAddBatch) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


