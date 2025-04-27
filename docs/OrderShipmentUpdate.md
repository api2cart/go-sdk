# OrderShipmentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentId** | **string** | Shipment id indicates the number of delivery | 
**OrderId** | Pointer to **string** | Defines the order that will be updated | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**ShipmentProvider** | Pointer to **string** | Defines company name that provide tracking of shipment | [optional] 
**TrackingNumbers** | Pointer to [**[]OrderShipmentAddTrackingNumbersInner**](OrderShipmentAddTrackingNumbersInner.md) | Defines shipment&#39;s tracking numbers that have to be added&lt;/br&gt; How set tracking numbers to appropriate carrier:&lt;ul&gt;&lt;li&gt;tracking_numbers[]&#x3D;a2c.demo1,a2c.demo2 - set default carrier&lt;/li&gt;&lt;li&gt;tracking_numbers[&lt;b&gt;carrier_id&lt;/b&gt;]&#x3D;a2c.demo - set appropriate carrier&lt;/li&gt;&lt;/ul&gt;To get the list of carriers IDs that are available in your store, use the &lt;a href &#x3D; \&quot;https://api2cart.com/docs/#/cart/CartInfo\&quot;&gt;cart.info&lt;/a &gt; method | [optional] 
**TrackingLink** | Pointer to **string** | Defines custom tracking link | [optional] 
**IsShipped** | Pointer to **bool** | Defines shipment&#39;s status | [optional] [default to true]
**DeliveredAt** | Pointer to **string** | Defines the date of delivery | [optional] 
**Replace** | Pointer to **bool** | Allows rewrite tracking numbers | [optional] [default to true]

## Methods

### NewOrderShipmentUpdate

`func NewOrderShipmentUpdate(shipmentId string, ) *OrderShipmentUpdate`

NewOrderShipmentUpdate instantiates a new OrderShipmentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderShipmentUpdateWithDefaults

`func NewOrderShipmentUpdateWithDefaults() *OrderShipmentUpdate`

NewOrderShipmentUpdateWithDefaults instantiates a new OrderShipmentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipmentId

`func (o *OrderShipmentUpdate) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *OrderShipmentUpdate) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *OrderShipmentUpdate) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.


### GetOrderId

`func (o *OrderShipmentUpdate) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderShipmentUpdate) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderShipmentUpdate) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderShipmentUpdate) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStoreId

`func (o *OrderShipmentUpdate) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderShipmentUpdate) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderShipmentUpdate) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderShipmentUpdate) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetShipmentProvider

`func (o *OrderShipmentUpdate) GetShipmentProvider() string`

GetShipmentProvider returns the ShipmentProvider field if non-nil, zero value otherwise.

### GetShipmentProviderOk

`func (o *OrderShipmentUpdate) GetShipmentProviderOk() (*string, bool)`

GetShipmentProviderOk returns a tuple with the ShipmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentProvider

`func (o *OrderShipmentUpdate) SetShipmentProvider(v string)`

SetShipmentProvider sets ShipmentProvider field to given value.

### HasShipmentProvider

`func (o *OrderShipmentUpdate) HasShipmentProvider() bool`

HasShipmentProvider returns a boolean if a field has been set.

### GetTrackingNumbers

`func (o *OrderShipmentUpdate) GetTrackingNumbers() []OrderShipmentAddTrackingNumbersInner`

GetTrackingNumbers returns the TrackingNumbers field if non-nil, zero value otherwise.

### GetTrackingNumbersOk

`func (o *OrderShipmentUpdate) GetTrackingNumbersOk() (*[]OrderShipmentAddTrackingNumbersInner, bool)`

GetTrackingNumbersOk returns a tuple with the TrackingNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumbers

`func (o *OrderShipmentUpdate) SetTrackingNumbers(v []OrderShipmentAddTrackingNumbersInner)`

SetTrackingNumbers sets TrackingNumbers field to given value.

### HasTrackingNumbers

`func (o *OrderShipmentUpdate) HasTrackingNumbers() bool`

HasTrackingNumbers returns a boolean if a field has been set.

### GetTrackingLink

`func (o *OrderShipmentUpdate) GetTrackingLink() string`

GetTrackingLink returns the TrackingLink field if non-nil, zero value otherwise.

### GetTrackingLinkOk

`func (o *OrderShipmentUpdate) GetTrackingLinkOk() (*string, bool)`

GetTrackingLinkOk returns a tuple with the TrackingLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingLink

`func (o *OrderShipmentUpdate) SetTrackingLink(v string)`

SetTrackingLink sets TrackingLink field to given value.

### HasTrackingLink

`func (o *OrderShipmentUpdate) HasTrackingLink() bool`

HasTrackingLink returns a boolean if a field has been set.

### GetIsShipped

`func (o *OrderShipmentUpdate) GetIsShipped() bool`

GetIsShipped returns the IsShipped field if non-nil, zero value otherwise.

### GetIsShippedOk

`func (o *OrderShipmentUpdate) GetIsShippedOk() (*bool, bool)`

GetIsShippedOk returns a tuple with the IsShipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShipped

`func (o *OrderShipmentUpdate) SetIsShipped(v bool)`

SetIsShipped sets IsShipped field to given value.

### HasIsShipped

`func (o *OrderShipmentUpdate) HasIsShipped() bool`

HasIsShipped returns a boolean if a field has been set.

### GetDeliveredAt

`func (o *OrderShipmentUpdate) GetDeliveredAt() string`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *OrderShipmentUpdate) GetDeliveredAtOk() (*string, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *OrderShipmentUpdate) SetDeliveredAt(v string)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *OrderShipmentUpdate) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### GetReplace

`func (o *OrderShipmentUpdate) GetReplace() bool`

GetReplace returns the Replace field if non-nil, zero value otherwise.

### GetReplaceOk

`func (o *OrderShipmentUpdate) GetReplaceOk() (*bool, bool)`

GetReplaceOk returns a tuple with the Replace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplace

`func (o *OrderShipmentUpdate) SetReplace(v bool)`

SetReplace sets Replace field to given value.

### HasReplace

`func (o *OrderShipmentUpdate) HasReplace() bool`

HasReplace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


