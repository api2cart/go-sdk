# ResponseCategoryListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoriesCount** | Pointer to **int32** |  | [optional] 
**Category** | Pointer to [**[]Category**](Category.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseCategoryListResult

`func NewResponseCategoryListResult() *ResponseCategoryListResult`

NewResponseCategoryListResult instantiates a new ResponseCategoryListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCategoryListResultWithDefaults

`func NewResponseCategoryListResultWithDefaults() *ResponseCategoryListResult`

NewResponseCategoryListResultWithDefaults instantiates a new ResponseCategoryListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoriesCount

`func (o *ResponseCategoryListResult) GetCategoriesCount() int32`

GetCategoriesCount returns the CategoriesCount field if non-nil, zero value otherwise.

### GetCategoriesCountOk

`func (o *ResponseCategoryListResult) GetCategoriesCountOk() (*int32, bool)`

GetCategoriesCountOk returns a tuple with the CategoriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesCount

`func (o *ResponseCategoryListResult) SetCategoriesCount(v int32)`

SetCategoriesCount sets CategoriesCount field to given value.

### HasCategoriesCount

`func (o *ResponseCategoryListResult) HasCategoriesCount() bool`

HasCategoriesCount returns a boolean if a field has been set.

### GetCategory

`func (o *ResponseCategoryListResult) GetCategory() []Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ResponseCategoryListResult) GetCategoryOk() (*[]Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ResponseCategoryListResult) SetCategory(v []Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ResponseCategoryListResult) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseCategoryListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseCategoryListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseCategoryListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseCategoryListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseCategoryListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseCategoryListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseCategoryListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseCategoryListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


