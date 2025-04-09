# OrderShipmentDelete200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **int32** |  | [optional] 
**ReturnMessage** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**OrderShipmentDelete200ResponseResult**](OrderShipmentDelete200ResponseResult.md) |  | [optional] 

## Methods

### NewOrderShipmentDelete200Response

`func NewOrderShipmentDelete200Response() *OrderShipmentDelete200Response`

NewOrderShipmentDelete200Response instantiates a new OrderShipmentDelete200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderShipmentDelete200ResponseWithDefaults

`func NewOrderShipmentDelete200ResponseWithDefaults() *OrderShipmentDelete200Response`

NewOrderShipmentDelete200ResponseWithDefaults instantiates a new OrderShipmentDelete200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *OrderShipmentDelete200Response) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *OrderShipmentDelete200Response) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *OrderShipmentDelete200Response) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *OrderShipmentDelete200Response) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnMessage

`func (o *OrderShipmentDelete200Response) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *OrderShipmentDelete200Response) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *OrderShipmentDelete200Response) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *OrderShipmentDelete200Response) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### GetResult

`func (o *OrderShipmentDelete200Response) GetResult() OrderShipmentDelete200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OrderShipmentDelete200Response) GetResultOk() (*OrderShipmentDelete200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OrderShipmentDelete200Response) SetResult(v OrderShipmentDelete200ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *OrderShipmentDelete200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


