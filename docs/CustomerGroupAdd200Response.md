# CustomerGroupAdd200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **int32** |  | [optional] 
**ReturnMessage** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**CustomerGroupAdd200ResponseResult**](CustomerGroupAdd200ResponseResult.md) |  | [optional] 

## Methods

### NewCustomerGroupAdd200Response

`func NewCustomerGroupAdd200Response() *CustomerGroupAdd200Response`

NewCustomerGroupAdd200Response instantiates a new CustomerGroupAdd200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerGroupAdd200ResponseWithDefaults

`func NewCustomerGroupAdd200ResponseWithDefaults() *CustomerGroupAdd200Response`

NewCustomerGroupAdd200ResponseWithDefaults instantiates a new CustomerGroupAdd200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *CustomerGroupAdd200Response) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *CustomerGroupAdd200Response) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *CustomerGroupAdd200Response) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *CustomerGroupAdd200Response) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnMessage

`func (o *CustomerGroupAdd200Response) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *CustomerGroupAdd200Response) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *CustomerGroupAdd200Response) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *CustomerGroupAdd200Response) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### GetResult

`func (o *CustomerGroupAdd200Response) GetResult() CustomerGroupAdd200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CustomerGroupAdd200Response) GetResultOk() (*CustomerGroupAdd200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CustomerGroupAdd200Response) SetResult(v CustomerGroupAdd200ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CustomerGroupAdd200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


