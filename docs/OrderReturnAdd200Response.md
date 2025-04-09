# OrderReturnAdd200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **int32** |  | [optional] 
**ReturnMessage** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**OrderReturnAdd200ResponseResult**](OrderReturnAdd200ResponseResult.md) |  | [optional] 

## Methods

### NewOrderReturnAdd200Response

`func NewOrderReturnAdd200Response() *OrderReturnAdd200Response`

NewOrderReturnAdd200Response instantiates a new OrderReturnAdd200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderReturnAdd200ResponseWithDefaults

`func NewOrderReturnAdd200ResponseWithDefaults() *OrderReturnAdd200Response`

NewOrderReturnAdd200ResponseWithDefaults instantiates a new OrderReturnAdd200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *OrderReturnAdd200Response) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *OrderReturnAdd200Response) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *OrderReturnAdd200Response) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *OrderReturnAdd200Response) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnMessage

`func (o *OrderReturnAdd200Response) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *OrderReturnAdd200Response) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *OrderReturnAdd200Response) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *OrderReturnAdd200Response) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### GetResult

`func (o *OrderReturnAdd200Response) GetResult() OrderReturnAdd200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OrderReturnAdd200Response) GetResultOk() (*OrderReturnAdd200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OrderReturnAdd200Response) SetResult(v OrderReturnAdd200ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *OrderReturnAdd200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


