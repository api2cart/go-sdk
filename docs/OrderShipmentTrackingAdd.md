# OrderShipmentTrackingAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreId** | Pointer to **string** | Store Id | [optional] 
**OrderId** | Pointer to **string** | Defines the order id | [optional] 
**ShipmentId** | **string** | Shipment id indicates the number of delivery | 
**CarrierId** | Pointer to **string** | Defines tracking carrier id | [optional] 
**TrackingProvider** | Pointer to **string** | Defines name of the company which provides shipment tracking | [optional] 
**TrackingNumber** | **string** | Defines tracking number | 
**TrackingLink** | Pointer to **string** | Defines custom tracking link | [optional] 
**SendNotifications** | Pointer to **bool** | Send notifications to customer after tracking was created | [optional] [default to false]

## Methods

### NewOrderShipmentTrackingAdd

`func NewOrderShipmentTrackingAdd(shipmentId string, trackingNumber string, ) *OrderShipmentTrackingAdd`

NewOrderShipmentTrackingAdd instantiates a new OrderShipmentTrackingAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderShipmentTrackingAddWithDefaults

`func NewOrderShipmentTrackingAddWithDefaults() *OrderShipmentTrackingAdd`

NewOrderShipmentTrackingAddWithDefaults instantiates a new OrderShipmentTrackingAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreId

`func (o *OrderShipmentTrackingAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderShipmentTrackingAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderShipmentTrackingAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderShipmentTrackingAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderShipmentTrackingAdd) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderShipmentTrackingAdd) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderShipmentTrackingAdd) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderShipmentTrackingAdd) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetShipmentId

`func (o *OrderShipmentTrackingAdd) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *OrderShipmentTrackingAdd) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *OrderShipmentTrackingAdd) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.


### GetCarrierId

`func (o *OrderShipmentTrackingAdd) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *OrderShipmentTrackingAdd) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *OrderShipmentTrackingAdd) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *OrderShipmentTrackingAdd) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetTrackingProvider

`func (o *OrderShipmentTrackingAdd) GetTrackingProvider() string`

GetTrackingProvider returns the TrackingProvider field if non-nil, zero value otherwise.

### GetTrackingProviderOk

`func (o *OrderShipmentTrackingAdd) GetTrackingProviderOk() (*string, bool)`

GetTrackingProviderOk returns a tuple with the TrackingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingProvider

`func (o *OrderShipmentTrackingAdd) SetTrackingProvider(v string)`

SetTrackingProvider sets TrackingProvider field to given value.

### HasTrackingProvider

`func (o *OrderShipmentTrackingAdd) HasTrackingProvider() bool`

HasTrackingProvider returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *OrderShipmentTrackingAdd) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *OrderShipmentTrackingAdd) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *OrderShipmentTrackingAdd) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.


### GetTrackingLink

`func (o *OrderShipmentTrackingAdd) GetTrackingLink() string`

GetTrackingLink returns the TrackingLink field if non-nil, zero value otherwise.

### GetTrackingLinkOk

`func (o *OrderShipmentTrackingAdd) GetTrackingLinkOk() (*string, bool)`

GetTrackingLinkOk returns a tuple with the TrackingLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingLink

`func (o *OrderShipmentTrackingAdd) SetTrackingLink(v string)`

SetTrackingLink sets TrackingLink field to given value.

### HasTrackingLink

`func (o *OrderShipmentTrackingAdd) HasTrackingLink() bool`

HasTrackingLink returns a boolean if a field has been set.

### GetSendNotifications

`func (o *OrderShipmentTrackingAdd) GetSendNotifications() bool`

GetSendNotifications returns the SendNotifications field if non-nil, zero value otherwise.

### GetSendNotificationsOk

`func (o *OrderShipmentTrackingAdd) GetSendNotificationsOk() (*bool, bool)`

GetSendNotificationsOk returns a tuple with the SendNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotifications

`func (o *OrderShipmentTrackingAdd) SetSendNotifications(v bool)`

SetSendNotifications sets SendNotifications field to given value.

### HasSendNotifications

`func (o *OrderShipmentTrackingAdd) HasSendNotifications() bool`

HasSendNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


