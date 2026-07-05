# ResponseAttributeTypeListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemsCount** | Pointer to **NullableInt32** |  | [optional] 
**AttributeType** | Pointer to **[]string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseAttributeTypeListResult

`func NewResponseAttributeTypeListResult() *ResponseAttributeTypeListResult`

NewResponseAttributeTypeListResult instantiates a new ResponseAttributeTypeListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAttributeTypeListResultWithDefaults

`func NewResponseAttributeTypeListResultWithDefaults() *ResponseAttributeTypeListResult`

NewResponseAttributeTypeListResultWithDefaults instantiates a new ResponseAttributeTypeListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemsCount

`func (o *ResponseAttributeTypeListResult) GetItemsCount() int32`

GetItemsCount returns the ItemsCount field if non-nil, zero value otherwise.

### GetItemsCountOk

`func (o *ResponseAttributeTypeListResult) GetItemsCountOk() (*int32, bool)`

GetItemsCountOk returns a tuple with the ItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsCount

`func (o *ResponseAttributeTypeListResult) SetItemsCount(v int32)`

SetItemsCount sets ItemsCount field to given value.

### HasItemsCount

`func (o *ResponseAttributeTypeListResult) HasItemsCount() bool`

HasItemsCount returns a boolean if a field has been set.

### SetItemsCountNil

`func (o *ResponseAttributeTypeListResult) SetItemsCountNil(b bool)`

 SetItemsCountNil sets the value for ItemsCount to be an explicit nil

### UnsetItemsCount
`func (o *ResponseAttributeTypeListResult) UnsetItemsCount()`

UnsetItemsCount ensures that no value is present for ItemsCount, not even an explicit nil
### GetAttributeType

`func (o *ResponseAttributeTypeListResult) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *ResponseAttributeTypeListResult) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *ResponseAttributeTypeListResult) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.

### HasAttributeType

`func (o *ResponseAttributeTypeListResult) HasAttributeType() bool`

HasAttributeType returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseAttributeTypeListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseAttributeTypeListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseAttributeTypeListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseAttributeTypeListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseAttributeTypeListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseAttributeTypeListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseAttributeTypeListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseAttributeTypeListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseAttributeTypeListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseAttributeTypeListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseAttributeTypeListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseAttributeTypeListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


