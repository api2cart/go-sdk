# ModelResponseProductList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **NullableInt32** |  | [optional] 
**ReturnMessage** | Pointer to **NullableString** |  | [optional] 
**Pagination** | Pointer to [**NullablePagination**](Pagination.md) |  | [optional] 
**Result** | Pointer to [**NullableResponseProductListResult**](ResponseProductListResult.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModelResponseProductList

`func NewModelResponseProductList() *ModelResponseProductList`

NewModelResponseProductList instantiates a new ModelResponseProductList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelResponseProductListWithDefaults

`func NewModelResponseProductListWithDefaults() *ModelResponseProductList`

NewModelResponseProductListWithDefaults instantiates a new ModelResponseProductList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *ModelResponseProductList) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ModelResponseProductList) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ModelResponseProductList) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ModelResponseProductList) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### SetReturnCodeNil

`func (o *ModelResponseProductList) SetReturnCodeNil(b bool)`

 SetReturnCodeNil sets the value for ReturnCode to be an explicit nil

### UnsetReturnCode
`func (o *ModelResponseProductList) UnsetReturnCode()`

UnsetReturnCode ensures that no value is present for ReturnCode, not even an explicit nil
### GetReturnMessage

`func (o *ModelResponseProductList) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *ModelResponseProductList) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *ModelResponseProductList) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *ModelResponseProductList) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### SetReturnMessageNil

`func (o *ModelResponseProductList) SetReturnMessageNil(b bool)`

 SetReturnMessageNil sets the value for ReturnMessage to be an explicit nil

### UnsetReturnMessage
`func (o *ModelResponseProductList) UnsetReturnMessage()`

UnsetReturnMessage ensures that no value is present for ReturnMessage, not even an explicit nil
### GetPagination

`func (o *ModelResponseProductList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ModelResponseProductList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ModelResponseProductList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ModelResponseProductList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### SetPaginationNil

`func (o *ModelResponseProductList) SetPaginationNil(b bool)`

 SetPaginationNil sets the value for Pagination to be an explicit nil

### UnsetPagination
`func (o *ModelResponseProductList) UnsetPagination()`

UnsetPagination ensures that no value is present for Pagination, not even an explicit nil
### GetResult

`func (o *ModelResponseProductList) GetResult() ResponseProductListResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModelResponseProductList) GetResultOk() (*ResponseProductListResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModelResponseProductList) SetResult(v ResponseProductListResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ModelResponseProductList) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ModelResponseProductList) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ModelResponseProductList) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetAdditionalFields

`func (o *ModelResponseProductList) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ModelResponseProductList) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ModelResponseProductList) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ModelResponseProductList) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ModelResponseProductList) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ModelResponseProductList) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ModelResponseProductList) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModelResponseProductList) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModelResponseProductList) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModelResponseProductList) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ModelResponseProductList) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ModelResponseProductList) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


