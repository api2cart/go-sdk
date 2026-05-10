# OrderShipmentEventAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentId** | **string** | Defines the shipment to which the tracking event will be added | 
**OrderId** | Pointer to **string** | Defines the order to which the shipment belongs | [optional] 
**Status** | **string** | Defines the tracking event status (e.g. in_transit, delivered, out_for_delivery) | 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**Address1** | Pointer to **string** | Specifies the street address of the event location | [optional] 
**City** | Pointer to **string** | Specifies city | [optional] 
**Country** | Pointer to **string** | Specifies ISO code or name of country | [optional] 
**State** | Pointer to **string** | Specifies ISO code or name of state | [optional] 
**Postcode** | Pointer to **string** | Specifies postcode | [optional] 
**Message** | Pointer to **string** | Defines a message associated with the tracking event. | [optional] 
**Latitude** | Pointer to **float32** | Latitude coordinate of the event location. | [optional] 
**Longitude** | Pointer to **float32** | Longitude coordinate of the event location. | [optional] 
**CreatedAt** | Pointer to **string** | Defines the date of entity creation | [optional] 
**EstimatedDeliveryAt** | Pointer to **string** | Estimated delivery date and time in ISO 8601 format. | [optional] 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 

## Methods

### NewOrderShipmentEventAdd

`func NewOrderShipmentEventAdd(shipmentId string, status string, ) *OrderShipmentEventAdd`

NewOrderShipmentEventAdd instantiates a new OrderShipmentEventAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderShipmentEventAddWithDefaults

`func NewOrderShipmentEventAddWithDefaults() *OrderShipmentEventAdd`

NewOrderShipmentEventAddWithDefaults instantiates a new OrderShipmentEventAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipmentId

`func (o *OrderShipmentEventAdd) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *OrderShipmentEventAdd) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *OrderShipmentEventAdd) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.


### GetOrderId

`func (o *OrderShipmentEventAdd) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderShipmentEventAdd) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderShipmentEventAdd) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderShipmentEventAdd) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *OrderShipmentEventAdd) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderShipmentEventAdd) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderShipmentEventAdd) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStoreId

`func (o *OrderShipmentEventAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderShipmentEventAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderShipmentEventAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderShipmentEventAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetAddress1

`func (o *OrderShipmentEventAdd) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *OrderShipmentEventAdd) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *OrderShipmentEventAdd) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *OrderShipmentEventAdd) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetCity

`func (o *OrderShipmentEventAdd) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderShipmentEventAdd) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderShipmentEventAdd) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrderShipmentEventAdd) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *OrderShipmentEventAdd) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrderShipmentEventAdd) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrderShipmentEventAdd) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OrderShipmentEventAdd) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *OrderShipmentEventAdd) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderShipmentEventAdd) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderShipmentEventAdd) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderShipmentEventAdd) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostcode

`func (o *OrderShipmentEventAdd) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *OrderShipmentEventAdd) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *OrderShipmentEventAdd) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *OrderShipmentEventAdd) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetMessage

`func (o *OrderShipmentEventAdd) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OrderShipmentEventAdd) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OrderShipmentEventAdd) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OrderShipmentEventAdd) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLatitude

`func (o *OrderShipmentEventAdd) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *OrderShipmentEventAdd) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *OrderShipmentEventAdd) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *OrderShipmentEventAdd) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *OrderShipmentEventAdd) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *OrderShipmentEventAdd) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *OrderShipmentEventAdd) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *OrderShipmentEventAdd) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderShipmentEventAdd) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderShipmentEventAdd) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderShipmentEventAdd) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderShipmentEventAdd) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEstimatedDeliveryAt

`func (o *OrderShipmentEventAdd) GetEstimatedDeliveryAt() string`

GetEstimatedDeliveryAt returns the EstimatedDeliveryAt field if non-nil, zero value otherwise.

### GetEstimatedDeliveryAtOk

`func (o *OrderShipmentEventAdd) GetEstimatedDeliveryAtOk() (*string, bool)`

GetEstimatedDeliveryAtOk returns a tuple with the EstimatedDeliveryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDeliveryAt

`func (o *OrderShipmentEventAdd) SetEstimatedDeliveryAt(v string)`

SetEstimatedDeliveryAt sets EstimatedDeliveryAt field to given value.

### HasEstimatedDeliveryAt

`func (o *OrderShipmentEventAdd) HasEstimatedDeliveryAt() bool`

HasEstimatedDeliveryAt returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *OrderShipmentEventAdd) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *OrderShipmentEventAdd) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *OrderShipmentEventAdd) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *OrderShipmentEventAdd) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


