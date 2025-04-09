# AttributeValueDelete200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **int32** |  | [optional] 
**ReturnMessage** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**AttributeValueDelete200ResponseResult**](AttributeValueDelete200ResponseResult.md) |  | [optional] 

## Methods

### NewAttributeValueDelete200Response

`func NewAttributeValueDelete200Response() *AttributeValueDelete200Response`

NewAttributeValueDelete200Response instantiates a new AttributeValueDelete200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeValueDelete200ResponseWithDefaults

`func NewAttributeValueDelete200ResponseWithDefaults() *AttributeValueDelete200Response`

NewAttributeValueDelete200ResponseWithDefaults instantiates a new AttributeValueDelete200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *AttributeValueDelete200Response) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *AttributeValueDelete200Response) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *AttributeValueDelete200Response) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *AttributeValueDelete200Response) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnMessage

`func (o *AttributeValueDelete200Response) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *AttributeValueDelete200Response) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *AttributeValueDelete200Response) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *AttributeValueDelete200Response) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### GetResult

`func (o *AttributeValueDelete200Response) GetResult() AttributeValueDelete200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AttributeValueDelete200Response) GetResultOk() (*AttributeValueDelete200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AttributeValueDelete200Response) SetResult(v AttributeValueDelete200ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *AttributeValueDelete200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


