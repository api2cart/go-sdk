# ModelResponseSubscriberList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **int32** |  | [optional] 
**ReturnMessage** | Pointer to **string** |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Result** | Pointer to [**ResponseSubscriberListResult**](ResponseSubscriberListResult.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModelResponseSubscriberList

`func NewModelResponseSubscriberList() *ModelResponseSubscriberList`

NewModelResponseSubscriberList instantiates a new ModelResponseSubscriberList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelResponseSubscriberListWithDefaults

`func NewModelResponseSubscriberListWithDefaults() *ModelResponseSubscriberList`

NewModelResponseSubscriberListWithDefaults instantiates a new ModelResponseSubscriberList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *ModelResponseSubscriberList) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ModelResponseSubscriberList) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ModelResponseSubscriberList) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ModelResponseSubscriberList) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnMessage

`func (o *ModelResponseSubscriberList) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *ModelResponseSubscriberList) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *ModelResponseSubscriberList) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *ModelResponseSubscriberList) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### GetPagination

`func (o *ModelResponseSubscriberList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ModelResponseSubscriberList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ModelResponseSubscriberList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ModelResponseSubscriberList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetResult

`func (o *ModelResponseSubscriberList) GetResult() ResponseSubscriberListResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModelResponseSubscriberList) GetResultOk() (*ResponseSubscriberListResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModelResponseSubscriberList) SetResult(v ResponseSubscriberListResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ModelResponseSubscriberList) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ModelResponseSubscriberList) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ModelResponseSubscriberList) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ModelResponseSubscriberList) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ModelResponseSubscriberList) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ModelResponseSubscriberList) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModelResponseSubscriberList) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModelResponseSubscriberList) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModelResponseSubscriberList) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


