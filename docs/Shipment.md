# Shipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**WarehouseId** | Pointer to **string** |  | [optional] 
**ShipmentProvider** | Pointer to **string** |  | [optional] 
**TrackingNumbers** | Pointer to [**[]ShipmentTrackingNumber**](ShipmentTrackingNumber.md) |  | [optional] 
**CreatedAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedTime** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**Items** | Pointer to [**[]ShipmentItem**](ShipmentItem.md) |  | [optional] 
**IsShipped** | Pointer to **bool** |  | [optional] 
**DeliveredAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewShipment

`func NewShipment() *Shipment`

NewShipment instantiates a new Shipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentWithDefaults

`func NewShipmentWithDefaults() *Shipment`

NewShipmentWithDefaults instantiates a new Shipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Shipment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Shipment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Shipment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Shipment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *Shipment) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Shipment) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Shipment) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Shipment) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetName

`func (o *Shipment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Shipment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Shipment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Shipment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWarehouseId

`func (o *Shipment) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *Shipment) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *Shipment) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *Shipment) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetShipmentProvider

`func (o *Shipment) GetShipmentProvider() string`

GetShipmentProvider returns the ShipmentProvider field if non-nil, zero value otherwise.

### GetShipmentProviderOk

`func (o *Shipment) GetShipmentProviderOk() (*string, bool)`

GetShipmentProviderOk returns a tuple with the ShipmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentProvider

`func (o *Shipment) SetShipmentProvider(v string)`

SetShipmentProvider sets ShipmentProvider field to given value.

### HasShipmentProvider

`func (o *Shipment) HasShipmentProvider() bool`

HasShipmentProvider returns a boolean if a field has been set.

### GetTrackingNumbers

`func (o *Shipment) GetTrackingNumbers() []ShipmentTrackingNumber`

GetTrackingNumbers returns the TrackingNumbers field if non-nil, zero value otherwise.

### GetTrackingNumbersOk

`func (o *Shipment) GetTrackingNumbersOk() (*[]ShipmentTrackingNumber, bool)`

GetTrackingNumbersOk returns a tuple with the TrackingNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumbers

`func (o *Shipment) SetTrackingNumbers(v []ShipmentTrackingNumber)`

SetTrackingNumbers sets TrackingNumbers field to given value.

### HasTrackingNumbers

`func (o *Shipment) HasTrackingNumbers() bool`

HasTrackingNumbers returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Shipment) GetCreatedAt() A2CDateTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Shipment) GetCreatedAtOk() (*A2CDateTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Shipment) SetCreatedAt(v A2CDateTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Shipment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Shipment) GetModifiedTime() A2CDateTime`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Shipment) GetModifiedTimeOk() (*A2CDateTime, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Shipment) SetModifiedTime(v A2CDateTime)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Shipment) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetItems

`func (o *Shipment) GetItems() []ShipmentItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Shipment) GetItemsOk() (*[]ShipmentItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Shipment) SetItems(v []ShipmentItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Shipment) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetIsShipped

`func (o *Shipment) GetIsShipped() bool`

GetIsShipped returns the IsShipped field if non-nil, zero value otherwise.

### GetIsShippedOk

`func (o *Shipment) GetIsShippedOk() (*bool, bool)`

GetIsShippedOk returns a tuple with the IsShipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShipped

`func (o *Shipment) SetIsShipped(v bool)`

SetIsShipped sets IsShipped field to given value.

### HasIsShipped

`func (o *Shipment) HasIsShipped() bool`

HasIsShipped returns a boolean if a field has been set.

### GetDeliveredAt

`func (o *Shipment) GetDeliveredAt() A2CDateTime`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *Shipment) GetDeliveredAtOk() (*A2CDateTime, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *Shipment) SetDeliveredAt(v A2CDateTime)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *Shipment) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Shipment) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Shipment) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Shipment) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Shipment) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *Shipment) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Shipment) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Shipment) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Shipment) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


