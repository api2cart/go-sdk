# ModelResponseOrderShipmentEventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **NullableInt32** |  | [optional] 
**ReturnMessage** | Pointer to **NullableString** |  | [optional] 
**Pagination** | Pointer to [**NullablePagination**](Pagination.md) |  | [optional] 
**Result** | Pointer to [**NullableResponseOrderShipmentEventListResult**](ResponseOrderShipmentEventListResult.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModelResponseOrderShipmentEventList

`func NewModelResponseOrderShipmentEventList() *ModelResponseOrderShipmentEventList`

NewModelResponseOrderShipmentEventList instantiates a new ModelResponseOrderShipmentEventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelResponseOrderShipmentEventListWithDefaults

`func NewModelResponseOrderShipmentEventListWithDefaults() *ModelResponseOrderShipmentEventList`

NewModelResponseOrderShipmentEventListWithDefaults instantiates a new ModelResponseOrderShipmentEventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *ModelResponseOrderShipmentEventList) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ModelResponseOrderShipmentEventList) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ModelResponseOrderShipmentEventList) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ModelResponseOrderShipmentEventList) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### SetReturnCodeNil

`func (o *ModelResponseOrderShipmentEventList) SetReturnCodeNil(b bool)`

 SetReturnCodeNil sets the value for ReturnCode to be an explicit nil

### UnsetReturnCode
`func (o *ModelResponseOrderShipmentEventList) UnsetReturnCode()`

UnsetReturnCode ensures that no value is present for ReturnCode, not even an explicit nil
### GetReturnMessage

`func (o *ModelResponseOrderShipmentEventList) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *ModelResponseOrderShipmentEventList) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *ModelResponseOrderShipmentEventList) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *ModelResponseOrderShipmentEventList) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### SetReturnMessageNil

`func (o *ModelResponseOrderShipmentEventList) SetReturnMessageNil(b bool)`

 SetReturnMessageNil sets the value for ReturnMessage to be an explicit nil

### UnsetReturnMessage
`func (o *ModelResponseOrderShipmentEventList) UnsetReturnMessage()`

UnsetReturnMessage ensures that no value is present for ReturnMessage, not even an explicit nil
### GetPagination

`func (o *ModelResponseOrderShipmentEventList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ModelResponseOrderShipmentEventList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ModelResponseOrderShipmentEventList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ModelResponseOrderShipmentEventList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### SetPaginationNil

`func (o *ModelResponseOrderShipmentEventList) SetPaginationNil(b bool)`

 SetPaginationNil sets the value for Pagination to be an explicit nil

### UnsetPagination
`func (o *ModelResponseOrderShipmentEventList) UnsetPagination()`

UnsetPagination ensures that no value is present for Pagination, not even an explicit nil
### GetResult

`func (o *ModelResponseOrderShipmentEventList) GetResult() ResponseOrderShipmentEventListResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModelResponseOrderShipmentEventList) GetResultOk() (*ResponseOrderShipmentEventListResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModelResponseOrderShipmentEventList) SetResult(v ResponseOrderShipmentEventListResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ModelResponseOrderShipmentEventList) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ModelResponseOrderShipmentEventList) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ModelResponseOrderShipmentEventList) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetAdditionalFields

`func (o *ModelResponseOrderShipmentEventList) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ModelResponseOrderShipmentEventList) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ModelResponseOrderShipmentEventList) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ModelResponseOrderShipmentEventList) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ModelResponseOrderShipmentEventList) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ModelResponseOrderShipmentEventList) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ModelResponseOrderShipmentEventList) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModelResponseOrderShipmentEventList) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModelResponseOrderShipmentEventList) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModelResponseOrderShipmentEventList) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ModelResponseOrderShipmentEventList) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ModelResponseOrderShipmentEventList) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


