# BatchJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**ProcessedTime** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBatchJob

`func NewBatchJob() *BatchJob`

NewBatchJob instantiates a new BatchJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchJobWithDefaults

`func NewBatchJobWithDefaults() *BatchJob`

NewBatchJobWithDefaults instantiates a new BatchJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BatchJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *BatchJob) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *BatchJob) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *BatchJob) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *BatchJob) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetStatus

`func (o *BatchJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedTime

`func (o *BatchJob) GetCreatedTime() A2CDateTime`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BatchJob) GetCreatedTimeOk() (*A2CDateTime, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BatchJob) SetCreatedTime(v A2CDateTime)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *BatchJob) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetProcessedTime

`func (o *BatchJob) GetProcessedTime() A2CDateTime`

GetProcessedTime returns the ProcessedTime field if non-nil, zero value otherwise.

### GetProcessedTimeOk

`func (o *BatchJob) GetProcessedTimeOk() (*A2CDateTime, bool)`

GetProcessedTimeOk returns a tuple with the ProcessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedTime

`func (o *BatchJob) SetProcessedTime(v A2CDateTime)`

SetProcessedTime sets ProcessedTime field to given value.

### HasProcessedTime

`func (o *BatchJob) HasProcessedTime() bool`

HasProcessedTime returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *BatchJob) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *BatchJob) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *BatchJob) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *BatchJob) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *BatchJob) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BatchJob) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BatchJob) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BatchJob) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


