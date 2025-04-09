# ResponseCustomerGroupListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupCount** | Pointer to **int32** |  | [optional] 
**Group** | Pointer to [**[]CustomerGroup**](CustomerGroup.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseCustomerGroupListResult

`func NewResponseCustomerGroupListResult() *ResponseCustomerGroupListResult`

NewResponseCustomerGroupListResult instantiates a new ResponseCustomerGroupListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCustomerGroupListResultWithDefaults

`func NewResponseCustomerGroupListResultWithDefaults() *ResponseCustomerGroupListResult`

NewResponseCustomerGroupListResultWithDefaults instantiates a new ResponseCustomerGroupListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupCount

`func (o *ResponseCustomerGroupListResult) GetGroupCount() int32`

GetGroupCount returns the GroupCount field if non-nil, zero value otherwise.

### GetGroupCountOk

`func (o *ResponseCustomerGroupListResult) GetGroupCountOk() (*int32, bool)`

GetGroupCountOk returns a tuple with the GroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCount

`func (o *ResponseCustomerGroupListResult) SetGroupCount(v int32)`

SetGroupCount sets GroupCount field to given value.

### HasGroupCount

`func (o *ResponseCustomerGroupListResult) HasGroupCount() bool`

HasGroupCount returns a boolean if a field has been set.

### GetGroup

`func (o *ResponseCustomerGroupListResult) GetGroup() []CustomerGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ResponseCustomerGroupListResult) GetGroupOk() (*[]CustomerGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ResponseCustomerGroupListResult) SetGroup(v []CustomerGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ResponseCustomerGroupListResult) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseCustomerGroupListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseCustomerGroupListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseCustomerGroupListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseCustomerGroupListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseCustomerGroupListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseCustomerGroupListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseCustomerGroupListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseCustomerGroupListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


