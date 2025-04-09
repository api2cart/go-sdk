# BatchJobResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **int32** |  | [optional] 
**JobName** | Pointer to **string** |  | [optional] 
**ItemsProcessed** | Pointer to **int32** |  | [optional] 
**ItemsSucceed** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]BatchJobResultItem**](BatchJobResultItem.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBatchJobResult

`func NewBatchJobResult() *BatchJobResult`

NewBatchJobResult instantiates a new BatchJobResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchJobResultWithDefaults

`func NewBatchJobResultWithDefaults() *BatchJobResult`

NewBatchJobResultWithDefaults instantiates a new BatchJobResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *BatchJobResult) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BatchJobResult) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BatchJobResult) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BatchJobResult) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobName

`func (o *BatchJobResult) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *BatchJobResult) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *BatchJobResult) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *BatchJobResult) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetItemsProcessed

`func (o *BatchJobResult) GetItemsProcessed() int32`

GetItemsProcessed returns the ItemsProcessed field if non-nil, zero value otherwise.

### GetItemsProcessedOk

`func (o *BatchJobResult) GetItemsProcessedOk() (*int32, bool)`

GetItemsProcessedOk returns a tuple with the ItemsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsProcessed

`func (o *BatchJobResult) SetItemsProcessed(v int32)`

SetItemsProcessed sets ItemsProcessed field to given value.

### HasItemsProcessed

`func (o *BatchJobResult) HasItemsProcessed() bool`

HasItemsProcessed returns a boolean if a field has been set.

### GetItemsSucceed

`func (o *BatchJobResult) GetItemsSucceed() int32`

GetItemsSucceed returns the ItemsSucceed field if non-nil, zero value otherwise.

### GetItemsSucceedOk

`func (o *BatchJobResult) GetItemsSucceedOk() (*int32, bool)`

GetItemsSucceedOk returns a tuple with the ItemsSucceed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsSucceed

`func (o *BatchJobResult) SetItemsSucceed(v int32)`

SetItemsSucceed sets ItemsSucceed field to given value.

### HasItemsSucceed

`func (o *BatchJobResult) HasItemsSucceed() bool`

HasItemsSucceed returns a boolean if a field has been set.

### GetItems

`func (o *BatchJobResult) GetItems() []BatchJobResultItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BatchJobResult) GetItemsOk() (*[]BatchJobResultItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BatchJobResult) SetItems(v []BatchJobResultItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *BatchJobResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *BatchJobResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *BatchJobResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *BatchJobResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *BatchJobResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *BatchJobResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BatchJobResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BatchJobResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BatchJobResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


