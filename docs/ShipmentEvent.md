# ShipmentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**EstimatedDeliveryAt** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to [**NullableCustomerAddress**](CustomerAddress.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewShipmentEvent

`func NewShipmentEvent() *ShipmentEvent`

NewShipmentEvent instantiates a new ShipmentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentEventWithDefaults

`func NewShipmentEventWithDefaults() *ShipmentEvent`

NewShipmentEventWithDefaults instantiates a new ShipmentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShipmentEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShipmentEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ShipmentEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipmentEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipmentEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShipmentEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ShipmentEvent) GetCreatedAt() A2CDateTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShipmentEvent) GetCreatedAtOk() (*A2CDateTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShipmentEvent) SetCreatedAt(v A2CDateTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ShipmentEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ShipmentEvent) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ShipmentEvent) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetEstimatedDeliveryAt

`func (o *ShipmentEvent) GetEstimatedDeliveryAt() A2CDateTime`

GetEstimatedDeliveryAt returns the EstimatedDeliveryAt field if non-nil, zero value otherwise.

### GetEstimatedDeliveryAtOk

`func (o *ShipmentEvent) GetEstimatedDeliveryAtOk() (*A2CDateTime, bool)`

GetEstimatedDeliveryAtOk returns a tuple with the EstimatedDeliveryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDeliveryAt

`func (o *ShipmentEvent) SetEstimatedDeliveryAt(v A2CDateTime)`

SetEstimatedDeliveryAt sets EstimatedDeliveryAt field to given value.

### HasEstimatedDeliveryAt

`func (o *ShipmentEvent) HasEstimatedDeliveryAt() bool`

HasEstimatedDeliveryAt returns a boolean if a field has been set.

### SetEstimatedDeliveryAtNil

`func (o *ShipmentEvent) SetEstimatedDeliveryAtNil(b bool)`

 SetEstimatedDeliveryAtNil sets the value for EstimatedDeliveryAt to be an explicit nil

### UnsetEstimatedDeliveryAt
`func (o *ShipmentEvent) UnsetEstimatedDeliveryAt()`

UnsetEstimatedDeliveryAt ensures that no value is present for EstimatedDeliveryAt, not even an explicit nil
### GetMessage

`func (o *ShipmentEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ShipmentEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ShipmentEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ShipmentEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ShipmentEvent) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ShipmentEvent) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetAddress

`func (o *ShipmentEvent) GetAddress() CustomerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ShipmentEvent) GetAddressOk() (*CustomerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ShipmentEvent) SetAddress(v CustomerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ShipmentEvent) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ShipmentEvent) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ShipmentEvent) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAdditionalFields

`func (o *ShipmentEvent) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ShipmentEvent) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ShipmentEvent) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ShipmentEvent) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ShipmentEvent) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ShipmentEvent) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ShipmentEvent) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ShipmentEvent) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ShipmentEvent) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ShipmentEvent) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ShipmentEvent) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ShipmentEvent) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


