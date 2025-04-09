# BatchJobResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**Warnings** | Pointer to **[]string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBatchJobResultItem

`func NewBatchJobResultItem() *BatchJobResultItem`

NewBatchJobResultItem instantiates a new BatchJobResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchJobResultItemWithDefaults

`func NewBatchJobResultItemWithDefaults() *BatchJobResultItem`

NewBatchJobResultItemWithDefaults instantiates a new BatchJobResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchJobResultItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchJobResultItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchJobResultItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BatchJobResultItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *BatchJobResultItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchJobResultItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchJobResultItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchJobResultItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntityId

`func (o *BatchJobResultItem) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *BatchJobResultItem) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *BatchJobResultItem) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *BatchJobResultItem) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetErrors

`func (o *BatchJobResultItem) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchJobResultItem) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchJobResultItem) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchJobResultItem) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *BatchJobResultItem) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *BatchJobResultItem) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *BatchJobResultItem) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *BatchJobResultItem) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *BatchJobResultItem) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *BatchJobResultItem) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *BatchJobResultItem) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *BatchJobResultItem) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *BatchJobResultItem) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BatchJobResultItem) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BatchJobResultItem) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BatchJobResultItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


