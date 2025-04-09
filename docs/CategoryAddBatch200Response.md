# CategoryAddBatch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **int32** |  | [optional] 
**ReturnMessage** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**CategoryAddBatch200ResponseResult**](CategoryAddBatch200ResponseResult.md) |  | [optional] 

## Methods

### NewCategoryAddBatch200Response

`func NewCategoryAddBatch200Response() *CategoryAddBatch200Response`

NewCategoryAddBatch200Response instantiates a new CategoryAddBatch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryAddBatch200ResponseWithDefaults

`func NewCategoryAddBatch200ResponseWithDefaults() *CategoryAddBatch200Response`

NewCategoryAddBatch200ResponseWithDefaults instantiates a new CategoryAddBatch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *CategoryAddBatch200Response) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *CategoryAddBatch200Response) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *CategoryAddBatch200Response) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *CategoryAddBatch200Response) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnMessage

`func (o *CategoryAddBatch200Response) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *CategoryAddBatch200Response) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *CategoryAddBatch200Response) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *CategoryAddBatch200Response) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### GetResult

`func (o *CategoryAddBatch200Response) GetResult() CategoryAddBatch200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CategoryAddBatch200Response) GetResultOk() (*CategoryAddBatch200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CategoryAddBatch200Response) SetResult(v CategoryAddBatch200ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CategoryAddBatch200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


