# AccountFailedWebhooks200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFailedWebhook** | Pointer to **string** |  | [optional] 
**Webhook** | Pointer to [**[]AccountFailedWebhooks200ResponseResultWebhookInner**](AccountFailedWebhooks200ResponseResultWebhookInner.md) |  | [optional] 

## Methods

### NewAccountFailedWebhooks200ResponseResult

`func NewAccountFailedWebhooks200ResponseResult() *AccountFailedWebhooks200ResponseResult`

NewAccountFailedWebhooks200ResponseResult instantiates a new AccountFailedWebhooks200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFailedWebhooks200ResponseResultWithDefaults

`func NewAccountFailedWebhooks200ResponseResultWithDefaults() *AccountFailedWebhooks200ResponseResult`

NewAccountFailedWebhooks200ResponseResultWithDefaults instantiates a new AccountFailedWebhooks200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFailedWebhook

`func (o *AccountFailedWebhooks200ResponseResult) GetAllFailedWebhook() string`

GetAllFailedWebhook returns the AllFailedWebhook field if non-nil, zero value otherwise.

### GetAllFailedWebhookOk

`func (o *AccountFailedWebhooks200ResponseResult) GetAllFailedWebhookOk() (*string, bool)`

GetAllFailedWebhookOk returns a tuple with the AllFailedWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFailedWebhook

`func (o *AccountFailedWebhooks200ResponseResult) SetAllFailedWebhook(v string)`

SetAllFailedWebhook sets AllFailedWebhook field to given value.

### HasAllFailedWebhook

`func (o *AccountFailedWebhooks200ResponseResult) HasAllFailedWebhook() bool`

HasAllFailedWebhook returns a boolean if a field has been set.

### GetWebhook

`func (o *AccountFailedWebhooks200ResponseResult) GetWebhook() []AccountFailedWebhooks200ResponseResultWebhookInner`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *AccountFailedWebhooks200ResponseResult) GetWebhookOk() (*[]AccountFailedWebhooks200ResponseResultWebhookInner, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *AccountFailedWebhooks200ResponseResult) SetWebhook(v []AccountFailedWebhooks200ResponseResultWebhookInner)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *AccountFailedWebhooks200ResponseResult) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


