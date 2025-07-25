# ModelResponseBatchJobList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **NullableInt32** |  | [optional] 
**ReturnMessage** | Pointer to **NullableString** |  | [optional] 
**Pagination** | Pointer to [**NullablePagination**](Pagination.md) |  | [optional] 
**Result** | Pointer to [**NullableResponseBatchJobListResult**](ResponseBatchJobListResult.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModelResponseBatchJobList

`func NewModelResponseBatchJobList() *ModelResponseBatchJobList`

NewModelResponseBatchJobList instantiates a new ModelResponseBatchJobList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelResponseBatchJobListWithDefaults

`func NewModelResponseBatchJobListWithDefaults() *ModelResponseBatchJobList`

NewModelResponseBatchJobListWithDefaults instantiates a new ModelResponseBatchJobList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *ModelResponseBatchJobList) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ModelResponseBatchJobList) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ModelResponseBatchJobList) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ModelResponseBatchJobList) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### SetReturnCodeNil

`func (o *ModelResponseBatchJobList) SetReturnCodeNil(b bool)`

 SetReturnCodeNil sets the value for ReturnCode to be an explicit nil

### UnsetReturnCode
`func (o *ModelResponseBatchJobList) UnsetReturnCode()`

UnsetReturnCode ensures that no value is present for ReturnCode, not even an explicit nil
### GetReturnMessage

`func (o *ModelResponseBatchJobList) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *ModelResponseBatchJobList) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *ModelResponseBatchJobList) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *ModelResponseBatchJobList) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### SetReturnMessageNil

`func (o *ModelResponseBatchJobList) SetReturnMessageNil(b bool)`

 SetReturnMessageNil sets the value for ReturnMessage to be an explicit nil

### UnsetReturnMessage
`func (o *ModelResponseBatchJobList) UnsetReturnMessage()`

UnsetReturnMessage ensures that no value is present for ReturnMessage, not even an explicit nil
### GetPagination

`func (o *ModelResponseBatchJobList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ModelResponseBatchJobList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ModelResponseBatchJobList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ModelResponseBatchJobList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### SetPaginationNil

`func (o *ModelResponseBatchJobList) SetPaginationNil(b bool)`

 SetPaginationNil sets the value for Pagination to be an explicit nil

### UnsetPagination
`func (o *ModelResponseBatchJobList) UnsetPagination()`

UnsetPagination ensures that no value is present for Pagination, not even an explicit nil
### GetResult

`func (o *ModelResponseBatchJobList) GetResult() ResponseBatchJobListResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModelResponseBatchJobList) GetResultOk() (*ResponseBatchJobListResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModelResponseBatchJobList) SetResult(v ResponseBatchJobListResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ModelResponseBatchJobList) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ModelResponseBatchJobList) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ModelResponseBatchJobList) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetAdditionalFields

`func (o *ModelResponseBatchJobList) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ModelResponseBatchJobList) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ModelResponseBatchJobList) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ModelResponseBatchJobList) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ModelResponseBatchJobList) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ModelResponseBatchJobList) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ModelResponseBatchJobList) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModelResponseBatchJobList) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModelResponseBatchJobList) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModelResponseBatchJobList) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ModelResponseBatchJobList) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ModelResponseBatchJobList) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


