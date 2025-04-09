# ResponseOrderShipmentListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentCount** | Pointer to **int32** |  | [optional] 
**Shipment** | Pointer to [**[]Shipment**](Shipment.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseOrderShipmentListResult

`func NewResponseOrderShipmentListResult() *ResponseOrderShipmentListResult`

NewResponseOrderShipmentListResult instantiates a new ResponseOrderShipmentListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrderShipmentListResultWithDefaults

`func NewResponseOrderShipmentListResultWithDefaults() *ResponseOrderShipmentListResult`

NewResponseOrderShipmentListResultWithDefaults instantiates a new ResponseOrderShipmentListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipmentCount

`func (o *ResponseOrderShipmentListResult) GetShipmentCount() int32`

GetShipmentCount returns the ShipmentCount field if non-nil, zero value otherwise.

### GetShipmentCountOk

`func (o *ResponseOrderShipmentListResult) GetShipmentCountOk() (*int32, bool)`

GetShipmentCountOk returns a tuple with the ShipmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentCount

`func (o *ResponseOrderShipmentListResult) SetShipmentCount(v int32)`

SetShipmentCount sets ShipmentCount field to given value.

### HasShipmentCount

`func (o *ResponseOrderShipmentListResult) HasShipmentCount() bool`

HasShipmentCount returns a boolean if a field has been set.

### GetShipment

`func (o *ResponseOrderShipmentListResult) GetShipment() []Shipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *ResponseOrderShipmentListResult) GetShipmentOk() (*[]Shipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *ResponseOrderShipmentListResult) SetShipment(v []Shipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *ResponseOrderShipmentListResult) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseOrderShipmentListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseOrderShipmentListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseOrderShipmentListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseOrderShipmentListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseOrderShipmentListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseOrderShipmentListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseOrderShipmentListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseOrderShipmentListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


