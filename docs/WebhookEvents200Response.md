# WebhookEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **int32** |  | [optional] 
**ReturnMessage** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**WebhookEvents200ResponseResult**](WebhookEvents200ResponseResult.md) |  | [optional] 

## Methods

### NewWebhookEvents200Response

`func NewWebhookEvents200Response() *WebhookEvents200Response`

NewWebhookEvents200Response instantiates a new WebhookEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEvents200ResponseWithDefaults

`func NewWebhookEvents200ResponseWithDefaults() *WebhookEvents200Response`

NewWebhookEvents200ResponseWithDefaults instantiates a new WebhookEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *WebhookEvents200Response) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *WebhookEvents200Response) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *WebhookEvents200Response) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *WebhookEvents200Response) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnMessage

`func (o *WebhookEvents200Response) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *WebhookEvents200Response) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *WebhookEvents200Response) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *WebhookEvents200Response) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### GetResult

`func (o *WebhookEvents200Response) GetResult() WebhookEvents200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WebhookEvents200Response) GetResultOk() (*WebhookEvents200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WebhookEvents200Response) SetResult(v WebhookEvents200ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *WebhookEvents200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


