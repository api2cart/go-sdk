# ResponseSubscriberListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Subscribers** | Pointer to [**[]Subscriber**](Subscriber.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseSubscriberListResult

`func NewResponseSubscriberListResult() *ResponseSubscriberListResult`

NewResponseSubscriberListResult instantiates a new ResponseSubscriberListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSubscriberListResultWithDefaults

`func NewResponseSubscriberListResultWithDefaults() *ResponseSubscriberListResult`

NewResponseSubscriberListResultWithDefaults instantiates a new ResponseSubscriberListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ResponseSubscriberListResult) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResponseSubscriberListResult) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResponseSubscriberListResult) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ResponseSubscriberListResult) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetSubscribers

`func (o *ResponseSubscriberListResult) GetSubscribers() []Subscriber`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *ResponseSubscriberListResult) GetSubscribersOk() (*[]Subscriber, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *ResponseSubscriberListResult) SetSubscribers(v []Subscriber)`

SetSubscribers sets Subscribers field to given value.

### HasSubscribers

`func (o *ResponseSubscriberListResult) HasSubscribers() bool`

HasSubscribers returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseSubscriberListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseSubscriberListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseSubscriberListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseSubscriberListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseSubscriberListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseSubscriberListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseSubscriberListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseSubscriberListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


