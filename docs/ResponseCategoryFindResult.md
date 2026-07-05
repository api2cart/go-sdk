# ResponseCategoryFindResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoriesCount** | Pointer to **NullableInt32** |  | [optional] 
**Category** | Pointer to [**[]Category**](Category.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseCategoryFindResult

`func NewResponseCategoryFindResult() *ResponseCategoryFindResult`

NewResponseCategoryFindResult instantiates a new ResponseCategoryFindResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCategoryFindResultWithDefaults

`func NewResponseCategoryFindResultWithDefaults() *ResponseCategoryFindResult`

NewResponseCategoryFindResultWithDefaults instantiates a new ResponseCategoryFindResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoriesCount

`func (o *ResponseCategoryFindResult) GetCategoriesCount() int32`

GetCategoriesCount returns the CategoriesCount field if non-nil, zero value otherwise.

### GetCategoriesCountOk

`func (o *ResponseCategoryFindResult) GetCategoriesCountOk() (*int32, bool)`

GetCategoriesCountOk returns a tuple with the CategoriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesCount

`func (o *ResponseCategoryFindResult) SetCategoriesCount(v int32)`

SetCategoriesCount sets CategoriesCount field to given value.

### HasCategoriesCount

`func (o *ResponseCategoryFindResult) HasCategoriesCount() bool`

HasCategoriesCount returns a boolean if a field has been set.

### SetCategoriesCountNil

`func (o *ResponseCategoryFindResult) SetCategoriesCountNil(b bool)`

 SetCategoriesCountNil sets the value for CategoriesCount to be an explicit nil

### UnsetCategoriesCount
`func (o *ResponseCategoryFindResult) UnsetCategoriesCount()`

UnsetCategoriesCount ensures that no value is present for CategoriesCount, not even an explicit nil
### GetCategory

`func (o *ResponseCategoryFindResult) GetCategory() []Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ResponseCategoryFindResult) GetCategoryOk() (*[]Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ResponseCategoryFindResult) SetCategory(v []Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ResponseCategoryFindResult) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseCategoryFindResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseCategoryFindResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseCategoryFindResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseCategoryFindResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseCategoryFindResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseCategoryFindResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseCategoryFindResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseCategoryFindResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseCategoryFindResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseCategoryFindResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseCategoryFindResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseCategoryFindResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


