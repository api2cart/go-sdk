# ResponseAttributeListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributesCount** | Pointer to **NullableInt32** |  | [optional] 
**Attribute** | Pointer to [**[]StoreAttribute**](StoreAttribute.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseAttributeListResult

`func NewResponseAttributeListResult() *ResponseAttributeListResult`

NewResponseAttributeListResult instantiates a new ResponseAttributeListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAttributeListResultWithDefaults

`func NewResponseAttributeListResultWithDefaults() *ResponseAttributeListResult`

NewResponseAttributeListResultWithDefaults instantiates a new ResponseAttributeListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributesCount

`func (o *ResponseAttributeListResult) GetAttributesCount() int32`

GetAttributesCount returns the AttributesCount field if non-nil, zero value otherwise.

### GetAttributesCountOk

`func (o *ResponseAttributeListResult) GetAttributesCountOk() (*int32, bool)`

GetAttributesCountOk returns a tuple with the AttributesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesCount

`func (o *ResponseAttributeListResult) SetAttributesCount(v int32)`

SetAttributesCount sets AttributesCount field to given value.

### HasAttributesCount

`func (o *ResponseAttributeListResult) HasAttributesCount() bool`

HasAttributesCount returns a boolean if a field has been set.

### SetAttributesCountNil

`func (o *ResponseAttributeListResult) SetAttributesCountNil(b bool)`

 SetAttributesCountNil sets the value for AttributesCount to be an explicit nil

### UnsetAttributesCount
`func (o *ResponseAttributeListResult) UnsetAttributesCount()`

UnsetAttributesCount ensures that no value is present for AttributesCount, not even an explicit nil
### GetAttribute

`func (o *ResponseAttributeListResult) GetAttribute() []StoreAttribute`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *ResponseAttributeListResult) GetAttributeOk() (*[]StoreAttribute, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *ResponseAttributeListResult) SetAttribute(v []StoreAttribute)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *ResponseAttributeListResult) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseAttributeListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseAttributeListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseAttributeListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseAttributeListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseAttributeListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseAttributeListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseAttributeListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseAttributeListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseAttributeListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseAttributeListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseAttributeListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseAttributeListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


