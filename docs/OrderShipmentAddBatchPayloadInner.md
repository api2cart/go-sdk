# OrderShipmentAddBatchPayloadInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** |  | 
**StoreId** | Pointer to **string** |  | [optional] 
**WarehouseId** | Pointer to **string** |  | [optional] 
**CarrierId** | Pointer to **string** |  | [optional] 
**CarrierName** | Pointer to **string** |  | [optional] 
**TrackingNumber** | **string** |  | 
**TrackingLink** | Pointer to **string** |  | [optional] 
**ShipmentProvider** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]OrderShipmentAddBatchPayloadInnerItemsInner**](OrderShipmentAddBatchPayloadInnerItemsInner.md) |  | [optional] 
**SendNotifications** | Pointer to **bool** |  | [optional] 

## Methods

### NewOrderShipmentAddBatchPayloadInner

`func NewOrderShipmentAddBatchPayloadInner(orderId string, trackingNumber string, ) *OrderShipmentAddBatchPayloadInner`

NewOrderShipmentAddBatchPayloadInner instantiates a new OrderShipmentAddBatchPayloadInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderShipmentAddBatchPayloadInnerWithDefaults

`func NewOrderShipmentAddBatchPayloadInnerWithDefaults() *OrderShipmentAddBatchPayloadInner`

NewOrderShipmentAddBatchPayloadInnerWithDefaults instantiates a new OrderShipmentAddBatchPayloadInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *OrderShipmentAddBatchPayloadInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderShipmentAddBatchPayloadInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderShipmentAddBatchPayloadInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetStoreId

`func (o *OrderShipmentAddBatchPayloadInner) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderShipmentAddBatchPayloadInner) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderShipmentAddBatchPayloadInner) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderShipmentAddBatchPayloadInner) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *OrderShipmentAddBatchPayloadInner) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *OrderShipmentAddBatchPayloadInner) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *OrderShipmentAddBatchPayloadInner) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *OrderShipmentAddBatchPayloadInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetCarrierId

`func (o *OrderShipmentAddBatchPayloadInner) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *OrderShipmentAddBatchPayloadInner) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *OrderShipmentAddBatchPayloadInner) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *OrderShipmentAddBatchPayloadInner) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetCarrierName

`func (o *OrderShipmentAddBatchPayloadInner) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *OrderShipmentAddBatchPayloadInner) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *OrderShipmentAddBatchPayloadInner) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *OrderShipmentAddBatchPayloadInner) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *OrderShipmentAddBatchPayloadInner) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *OrderShipmentAddBatchPayloadInner) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *OrderShipmentAddBatchPayloadInner) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.


### GetTrackingLink

`func (o *OrderShipmentAddBatchPayloadInner) GetTrackingLink() string`

GetTrackingLink returns the TrackingLink field if non-nil, zero value otherwise.

### GetTrackingLinkOk

`func (o *OrderShipmentAddBatchPayloadInner) GetTrackingLinkOk() (*string, bool)`

GetTrackingLinkOk returns a tuple with the TrackingLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingLink

`func (o *OrderShipmentAddBatchPayloadInner) SetTrackingLink(v string)`

SetTrackingLink sets TrackingLink field to given value.

### HasTrackingLink

`func (o *OrderShipmentAddBatchPayloadInner) HasTrackingLink() bool`

HasTrackingLink returns a boolean if a field has been set.

### GetShipmentProvider

`func (o *OrderShipmentAddBatchPayloadInner) GetShipmentProvider() string`

GetShipmentProvider returns the ShipmentProvider field if non-nil, zero value otherwise.

### GetShipmentProviderOk

`func (o *OrderShipmentAddBatchPayloadInner) GetShipmentProviderOk() (*string, bool)`

GetShipmentProviderOk returns a tuple with the ShipmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentProvider

`func (o *OrderShipmentAddBatchPayloadInner) SetShipmentProvider(v string)`

SetShipmentProvider sets ShipmentProvider field to given value.

### HasShipmentProvider

`func (o *OrderShipmentAddBatchPayloadInner) HasShipmentProvider() bool`

HasShipmentProvider returns a boolean if a field has been set.

### GetItems

`func (o *OrderShipmentAddBatchPayloadInner) GetItems() []OrderShipmentAddBatchPayloadInnerItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrderShipmentAddBatchPayloadInner) GetItemsOk() (*[]OrderShipmentAddBatchPayloadInnerItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrderShipmentAddBatchPayloadInner) SetItems(v []OrderShipmentAddBatchPayloadInnerItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *OrderShipmentAddBatchPayloadInner) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSendNotifications

`func (o *OrderShipmentAddBatchPayloadInner) GetSendNotifications() bool`

GetSendNotifications returns the SendNotifications field if non-nil, zero value otherwise.

### GetSendNotificationsOk

`func (o *OrderShipmentAddBatchPayloadInner) GetSendNotificationsOk() (*bool, bool)`

GetSendNotificationsOk returns a tuple with the SendNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotifications

`func (o *OrderShipmentAddBatchPayloadInner) SetSendNotifications(v bool)`

SetSendNotifications sets SendNotifications field to given value.

### HasSendNotifications

`func (o *OrderShipmentAddBatchPayloadInner) HasSendNotifications() bool`

HasSendNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


