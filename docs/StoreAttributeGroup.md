# StoreAttributeGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **NullableInt32** |  | [optional] 
**AttributeSetId** | Pointer to **NullableString** |  | [optional] 
**AssignedAttributeIds** | Pointer to **[]string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewStoreAttributeGroup

`func NewStoreAttributeGroup() *StoreAttributeGroup`

NewStoreAttributeGroup instantiates a new StoreAttributeGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreAttributeGroupWithDefaults

`func NewStoreAttributeGroupWithDefaults() *StoreAttributeGroup`

NewStoreAttributeGroupWithDefaults instantiates a new StoreAttributeGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StoreAttributeGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoreAttributeGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoreAttributeGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StoreAttributeGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StoreAttributeGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreAttributeGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreAttributeGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoreAttributeGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *StoreAttributeGroup) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *StoreAttributeGroup) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *StoreAttributeGroup) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *StoreAttributeGroup) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *StoreAttributeGroup) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *StoreAttributeGroup) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetAttributeSetId

`func (o *StoreAttributeGroup) GetAttributeSetId() string`

GetAttributeSetId returns the AttributeSetId field if non-nil, zero value otherwise.

### GetAttributeSetIdOk

`func (o *StoreAttributeGroup) GetAttributeSetIdOk() (*string, bool)`

GetAttributeSetIdOk returns a tuple with the AttributeSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSetId

`func (o *StoreAttributeGroup) SetAttributeSetId(v string)`

SetAttributeSetId sets AttributeSetId field to given value.

### HasAttributeSetId

`func (o *StoreAttributeGroup) HasAttributeSetId() bool`

HasAttributeSetId returns a boolean if a field has been set.

### SetAttributeSetIdNil

`func (o *StoreAttributeGroup) SetAttributeSetIdNil(b bool)`

 SetAttributeSetIdNil sets the value for AttributeSetId to be an explicit nil

### UnsetAttributeSetId
`func (o *StoreAttributeGroup) UnsetAttributeSetId()`

UnsetAttributeSetId ensures that no value is present for AttributeSetId, not even an explicit nil
### GetAssignedAttributeIds

`func (o *StoreAttributeGroup) GetAssignedAttributeIds() []string`

GetAssignedAttributeIds returns the AssignedAttributeIds field if non-nil, zero value otherwise.

### GetAssignedAttributeIdsOk

`func (o *StoreAttributeGroup) GetAssignedAttributeIdsOk() (*[]string, bool)`

GetAssignedAttributeIdsOk returns a tuple with the AssignedAttributeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAttributeIds

`func (o *StoreAttributeGroup) SetAssignedAttributeIds(v []string)`

SetAssignedAttributeIds sets AssignedAttributeIds field to given value.

### HasAssignedAttributeIds

`func (o *StoreAttributeGroup) HasAssignedAttributeIds() bool`

HasAssignedAttributeIds returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *StoreAttributeGroup) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *StoreAttributeGroup) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *StoreAttributeGroup) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *StoreAttributeGroup) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *StoreAttributeGroup) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *StoreAttributeGroup) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *StoreAttributeGroup) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *StoreAttributeGroup) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *StoreAttributeGroup) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *StoreAttributeGroup) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *StoreAttributeGroup) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *StoreAttributeGroup) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


