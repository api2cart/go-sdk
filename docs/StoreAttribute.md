# StoreAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DefaultValues** | Pointer to **[]string** |  | [optional] 
**Position** | Pointer to **NullableInt32** |  | [optional] 
**Visible** | Pointer to **NullableBool** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**System** | Pointer to **NullableBool** |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**StoreId** | Pointer to **NullableString** |  | [optional] 
**LangId** | Pointer to **NullableString** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewStoreAttribute

`func NewStoreAttribute() *StoreAttribute`

NewStoreAttribute instantiates a new StoreAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreAttributeWithDefaults

`func NewStoreAttributeWithDefaults() *StoreAttribute`

NewStoreAttributeWithDefaults instantiates a new StoreAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StoreAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoreAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoreAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StoreAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *StoreAttribute) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StoreAttribute) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StoreAttribute) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *StoreAttribute) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *StoreAttribute) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *StoreAttribute) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetType

`func (o *StoreAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoreAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoreAttribute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StoreAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *StoreAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoreAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultValues

`func (o *StoreAttribute) GetDefaultValues() []string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *StoreAttribute) GetDefaultValuesOk() (*[]string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *StoreAttribute) SetDefaultValues(v []string)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *StoreAttribute) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetPosition

`func (o *StoreAttribute) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *StoreAttribute) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *StoreAttribute) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *StoreAttribute) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *StoreAttribute) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *StoreAttribute) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetVisible

`func (o *StoreAttribute) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *StoreAttribute) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *StoreAttribute) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *StoreAttribute) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### SetVisibleNil

`func (o *StoreAttribute) SetVisibleNil(b bool)`

 SetVisibleNil sets the value for Visible to be an explicit nil

### UnsetVisible
`func (o *StoreAttribute) UnsetVisible()`

UnsetVisible ensures that no value is present for Visible, not even an explicit nil
### GetRequired

`func (o *StoreAttribute) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *StoreAttribute) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *StoreAttribute) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *StoreAttribute) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *StoreAttribute) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *StoreAttribute) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetSystem

`func (o *StoreAttribute) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *StoreAttribute) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *StoreAttribute) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *StoreAttribute) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *StoreAttribute) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *StoreAttribute) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetValues

`func (o *StoreAttribute) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StoreAttribute) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StoreAttribute) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *StoreAttribute) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStoreId

`func (o *StoreAttribute) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *StoreAttribute) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *StoreAttribute) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *StoreAttribute) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### SetStoreIdNil

`func (o *StoreAttribute) SetStoreIdNil(b bool)`

 SetStoreIdNil sets the value for StoreId to be an explicit nil

### UnsetStoreId
`func (o *StoreAttribute) UnsetStoreId()`

UnsetStoreId ensures that no value is present for StoreId, not even an explicit nil
### GetLangId

`func (o *StoreAttribute) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *StoreAttribute) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *StoreAttribute) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *StoreAttribute) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### SetLangIdNil

`func (o *StoreAttribute) SetLangIdNil(b bool)`

 SetLangIdNil sets the value for LangId to be an explicit nil

### UnsetLangId
`func (o *StoreAttribute) UnsetLangId()`

UnsetLangId ensures that no value is present for LangId, not even an explicit nil
### GetAdditionalFields

`func (o *StoreAttribute) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *StoreAttribute) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *StoreAttribute) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *StoreAttribute) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *StoreAttribute) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *StoreAttribute) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *StoreAttribute) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *StoreAttribute) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *StoreAttribute) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *StoreAttribute) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *StoreAttribute) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *StoreAttribute) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


