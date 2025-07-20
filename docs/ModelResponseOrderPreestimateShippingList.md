# ModelResponseOrderPreestimateShippingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **NullableInt32** |  | [optional] 
**ReturnMessage** | Pointer to **NullableString** |  | [optional] 
**Pagination** | Pointer to [**NullablePagination**](Pagination.md) |  | [optional] 
**Result** | Pointer to [**NullableResponseOrderPreestimateShippingListResult**](ResponseOrderPreestimateShippingListResult.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModelResponseOrderPreestimateShippingList

`func NewModelResponseOrderPreestimateShippingList() *ModelResponseOrderPreestimateShippingList`

NewModelResponseOrderPreestimateShippingList instantiates a new ModelResponseOrderPreestimateShippingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelResponseOrderPreestimateShippingListWithDefaults

`func NewModelResponseOrderPreestimateShippingListWithDefaults() *ModelResponseOrderPreestimateShippingList`

NewModelResponseOrderPreestimateShippingListWithDefaults instantiates a new ModelResponseOrderPreestimateShippingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *ModelResponseOrderPreestimateShippingList) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ModelResponseOrderPreestimateShippingList) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ModelResponseOrderPreestimateShippingList) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ModelResponseOrderPreestimateShippingList) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### SetReturnCodeNil

`func (o *ModelResponseOrderPreestimateShippingList) SetReturnCodeNil(b bool)`

 SetReturnCodeNil sets the value for ReturnCode to be an explicit nil

### UnsetReturnCode
`func (o *ModelResponseOrderPreestimateShippingList) UnsetReturnCode()`

UnsetReturnCode ensures that no value is present for ReturnCode, not even an explicit nil
### GetReturnMessage

`func (o *ModelResponseOrderPreestimateShippingList) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *ModelResponseOrderPreestimateShippingList) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *ModelResponseOrderPreestimateShippingList) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *ModelResponseOrderPreestimateShippingList) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### SetReturnMessageNil

`func (o *ModelResponseOrderPreestimateShippingList) SetReturnMessageNil(b bool)`

 SetReturnMessageNil sets the value for ReturnMessage to be an explicit nil

### UnsetReturnMessage
`func (o *ModelResponseOrderPreestimateShippingList) UnsetReturnMessage()`

UnsetReturnMessage ensures that no value is present for ReturnMessage, not even an explicit nil
### GetPagination

`func (o *ModelResponseOrderPreestimateShippingList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ModelResponseOrderPreestimateShippingList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ModelResponseOrderPreestimateShippingList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ModelResponseOrderPreestimateShippingList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### SetPaginationNil

`func (o *ModelResponseOrderPreestimateShippingList) SetPaginationNil(b bool)`

 SetPaginationNil sets the value for Pagination to be an explicit nil

### UnsetPagination
`func (o *ModelResponseOrderPreestimateShippingList) UnsetPagination()`

UnsetPagination ensures that no value is present for Pagination, not even an explicit nil
### GetResult

`func (o *ModelResponseOrderPreestimateShippingList) GetResult() ResponseOrderPreestimateShippingListResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModelResponseOrderPreestimateShippingList) GetResultOk() (*ResponseOrderPreestimateShippingListResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModelResponseOrderPreestimateShippingList) SetResult(v ResponseOrderPreestimateShippingListResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ModelResponseOrderPreestimateShippingList) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ModelResponseOrderPreestimateShippingList) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ModelResponseOrderPreestimateShippingList) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetAdditionalFields

`func (o *ModelResponseOrderPreestimateShippingList) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ModelResponseOrderPreestimateShippingList) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ModelResponseOrderPreestimateShippingList) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ModelResponseOrderPreestimateShippingList) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ModelResponseOrderPreestimateShippingList) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ModelResponseOrderPreestimateShippingList) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ModelResponseOrderPreestimateShippingList) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModelResponseOrderPreestimateShippingList) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModelResponseOrderPreestimateShippingList) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModelResponseOrderPreestimateShippingList) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ModelResponseOrderPreestimateShippingList) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ModelResponseOrderPreestimateShippingList) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


