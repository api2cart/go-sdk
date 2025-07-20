# Script

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Src** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Event** | Pointer to **NullableString** |  | [optional] 
**LoadMethod** | Pointer to **NullableString** |  | [optional] 
**Html** | Pointer to **NullableString** |  | [optional] 
**CreatedTime** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedTime** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewScript

`func NewScript() *Script`

NewScript instantiates a new Script object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptWithDefaults

`func NewScriptWithDefaults() *Script`

NewScriptWithDefaults instantiates a new Script object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Script) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Script) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Script) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Script) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Script) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Script) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Script) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Script) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Script) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Script) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Script) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Script) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Script) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Script) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Script) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Script) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSrc

`func (o *Script) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *Script) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *Script) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *Script) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### SetSrcNil

`func (o *Script) SetSrcNil(b bool)`

 SetSrcNil sets the value for Src to be an explicit nil

### UnsetSrc
`func (o *Script) UnsetSrc()`

UnsetSrc ensures that no value is present for Src, not even an explicit nil
### GetScope

`func (o *Script) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Script) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Script) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Script) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *Script) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *Script) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetEvent

`func (o *Script) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Script) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Script) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Script) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *Script) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *Script) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetLoadMethod

`func (o *Script) GetLoadMethod() string`

GetLoadMethod returns the LoadMethod field if non-nil, zero value otherwise.

### GetLoadMethodOk

`func (o *Script) GetLoadMethodOk() (*string, bool)`

GetLoadMethodOk returns a tuple with the LoadMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadMethod

`func (o *Script) SetLoadMethod(v string)`

SetLoadMethod sets LoadMethod field to given value.

### HasLoadMethod

`func (o *Script) HasLoadMethod() bool`

HasLoadMethod returns a boolean if a field has been set.

### SetLoadMethodNil

`func (o *Script) SetLoadMethodNil(b bool)`

 SetLoadMethodNil sets the value for LoadMethod to be an explicit nil

### UnsetLoadMethod
`func (o *Script) UnsetLoadMethod()`

UnsetLoadMethod ensures that no value is present for LoadMethod, not even an explicit nil
### GetHtml

`func (o *Script) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *Script) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *Script) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *Script) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### SetHtmlNil

`func (o *Script) SetHtmlNil(b bool)`

 SetHtmlNil sets the value for Html to be an explicit nil

### UnsetHtml
`func (o *Script) UnsetHtml()`

UnsetHtml ensures that no value is present for Html, not even an explicit nil
### GetCreatedTime

`func (o *Script) GetCreatedTime() A2CDateTime`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Script) GetCreatedTimeOk() (*A2CDateTime, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Script) SetCreatedTime(v A2CDateTime)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Script) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *Script) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *Script) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetModifiedTime

`func (o *Script) GetModifiedTime() A2CDateTime`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Script) GetModifiedTimeOk() (*A2CDateTime, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Script) SetModifiedTime(v A2CDateTime)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Script) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### SetModifiedTimeNil

`func (o *Script) SetModifiedTimeNil(b bool)`

 SetModifiedTimeNil sets the value for ModifiedTime to be an explicit nil

### UnsetModifiedTime
`func (o *Script) UnsetModifiedTime()`

UnsetModifiedTime ensures that no value is present for ModifiedTime, not even an explicit nil
### GetAdditionalFields

`func (o *Script) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Script) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Script) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Script) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *Script) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *Script) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *Script) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Script) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Script) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Script) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *Script) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *Script) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


