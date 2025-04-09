# ModelResponseCartScriptList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **int32** |  | [optional] 
**ReturnMessage** | Pointer to **string** |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Result** | Pointer to [**ResponseCartScriptListResult**](ResponseCartScriptListResult.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModelResponseCartScriptList

`func NewModelResponseCartScriptList() *ModelResponseCartScriptList`

NewModelResponseCartScriptList instantiates a new ModelResponseCartScriptList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelResponseCartScriptListWithDefaults

`func NewModelResponseCartScriptListWithDefaults() *ModelResponseCartScriptList`

NewModelResponseCartScriptListWithDefaults instantiates a new ModelResponseCartScriptList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *ModelResponseCartScriptList) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ModelResponseCartScriptList) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ModelResponseCartScriptList) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ModelResponseCartScriptList) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnMessage

`func (o *ModelResponseCartScriptList) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *ModelResponseCartScriptList) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *ModelResponseCartScriptList) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *ModelResponseCartScriptList) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### GetPagination

`func (o *ModelResponseCartScriptList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ModelResponseCartScriptList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ModelResponseCartScriptList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ModelResponseCartScriptList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResult

`func (o *ModelResponseCartScriptList) GetResult() ResponseCartScriptListResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModelResponseCartScriptList) GetResultOk() (*ResponseCartScriptListResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModelResponseCartScriptList) SetResult(v ResponseCartScriptListResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ModelResponseCartScriptList) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ModelResponseCartScriptList) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ModelResponseCartScriptList) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ModelResponseCartScriptList) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ModelResponseCartScriptList) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ModelResponseCartScriptList) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModelResponseCartScriptList) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModelResponseCartScriptList) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModelResponseCartScriptList) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


