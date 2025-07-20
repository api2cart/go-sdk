# ResponseOrderPreestimateShippingListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreestimateShippingsCount** | Pointer to **NullableInt32** |  | [optional] 
**PreestimateShippings** | Pointer to [**[]OrderPreestimateShipping**](OrderPreestimateShipping.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseOrderPreestimateShippingListResult

`func NewResponseOrderPreestimateShippingListResult() *ResponseOrderPreestimateShippingListResult`

NewResponseOrderPreestimateShippingListResult instantiates a new ResponseOrderPreestimateShippingListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrderPreestimateShippingListResultWithDefaults

`func NewResponseOrderPreestimateShippingListResultWithDefaults() *ResponseOrderPreestimateShippingListResult`

NewResponseOrderPreestimateShippingListResultWithDefaults instantiates a new ResponseOrderPreestimateShippingListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreestimateShippingsCount

`func (o *ResponseOrderPreestimateShippingListResult) GetPreestimateShippingsCount() int32`

GetPreestimateShippingsCount returns the PreestimateShippingsCount field if non-nil, zero value otherwise.

### GetPreestimateShippingsCountOk

`func (o *ResponseOrderPreestimateShippingListResult) GetPreestimateShippingsCountOk() (*int32, bool)`

GetPreestimateShippingsCountOk returns a tuple with the PreestimateShippingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreestimateShippingsCount

`func (o *ResponseOrderPreestimateShippingListResult) SetPreestimateShippingsCount(v int32)`

SetPreestimateShippingsCount sets PreestimateShippingsCount field to given value.

### HasPreestimateShippingsCount

`func (o *ResponseOrderPreestimateShippingListResult) HasPreestimateShippingsCount() bool`

HasPreestimateShippingsCount returns a boolean if a field has been set.

### SetPreestimateShippingsCountNil

`func (o *ResponseOrderPreestimateShippingListResult) SetPreestimateShippingsCountNil(b bool)`

 SetPreestimateShippingsCountNil sets the value for PreestimateShippingsCount to be an explicit nil

### UnsetPreestimateShippingsCount
`func (o *ResponseOrderPreestimateShippingListResult) UnsetPreestimateShippingsCount()`

UnsetPreestimateShippingsCount ensures that no value is present for PreestimateShippingsCount, not even an explicit nil
### GetPreestimateShippings

`func (o *ResponseOrderPreestimateShippingListResult) GetPreestimateShippings() []OrderPreestimateShipping`

GetPreestimateShippings returns the PreestimateShippings field if non-nil, zero value otherwise.

### GetPreestimateShippingsOk

`func (o *ResponseOrderPreestimateShippingListResult) GetPreestimateShippingsOk() (*[]OrderPreestimateShipping, bool)`

GetPreestimateShippingsOk returns a tuple with the PreestimateShippings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreestimateShippings

`func (o *ResponseOrderPreestimateShippingListResult) SetPreestimateShippings(v []OrderPreestimateShipping)`

SetPreestimateShippings sets PreestimateShippings field to given value.

### HasPreestimateShippings

`func (o *ResponseOrderPreestimateShippingListResult) HasPreestimateShippings() bool`

HasPreestimateShippings returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseOrderPreestimateShippingListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseOrderPreestimateShippingListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseOrderPreestimateShippingListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseOrderPreestimateShippingListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseOrderPreestimateShippingListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseOrderPreestimateShippingListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseOrderPreestimateShippingListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseOrderPreestimateShippingListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseOrderPreestimateShippingListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseOrderPreestimateShippingListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseOrderPreestimateShippingListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseOrderPreestimateShippingListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


