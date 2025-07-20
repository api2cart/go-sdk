# ResponseCartShippingZonesListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingZone** | Pointer to [**[]CartShippingZone2**](CartShippingZone2.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseCartShippingZonesListResult

`func NewResponseCartShippingZonesListResult() *ResponseCartShippingZonesListResult`

NewResponseCartShippingZonesListResult instantiates a new ResponseCartShippingZonesListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCartShippingZonesListResultWithDefaults

`func NewResponseCartShippingZonesListResultWithDefaults() *ResponseCartShippingZonesListResult`

NewResponseCartShippingZonesListResultWithDefaults instantiates a new ResponseCartShippingZonesListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingZone

`func (o *ResponseCartShippingZonesListResult) GetShippingZone() []CartShippingZone2`

GetShippingZone returns the ShippingZone field if non-nil, zero value otherwise.

### GetShippingZoneOk

`func (o *ResponseCartShippingZonesListResult) GetShippingZoneOk() (*[]CartShippingZone2, bool)`

GetShippingZoneOk returns a tuple with the ShippingZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingZone

`func (o *ResponseCartShippingZonesListResult) SetShippingZone(v []CartShippingZone2)`

SetShippingZone sets ShippingZone field to given value.

### HasShippingZone

`func (o *ResponseCartShippingZonesListResult) HasShippingZone() bool`

HasShippingZone returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseCartShippingZonesListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseCartShippingZonesListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseCartShippingZonesListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseCartShippingZonesListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseCartShippingZonesListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseCartShippingZonesListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseCartShippingZonesListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseCartShippingZonesListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseCartShippingZonesListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseCartShippingZonesListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseCartShippingZonesListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseCartShippingZonesListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


