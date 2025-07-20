# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**StoreId** | Pointer to **NullableString** |  | [optional] 
**LangId** | Pointer to **NullableString** |  | [optional] 
**Active** | Pointer to **NullableBool** |  | [optional] 
**Callback** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to **NullableString** |  | [optional] 
**ResponseFields** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**Entity** | Pointer to **NullableString** |  | [optional] 
**Action** | Pointer to **NullableString** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWebhook

`func NewWebhook() *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Webhook) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Webhook) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Webhook) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLabel

`func (o *Webhook) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Webhook) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Webhook) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Webhook) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *Webhook) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *Webhook) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetStoreId

`func (o *Webhook) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *Webhook) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *Webhook) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *Webhook) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### SetStoreIdNil

`func (o *Webhook) SetStoreIdNil(b bool)`

 SetStoreIdNil sets the value for StoreId to be an explicit nil

### UnsetStoreId
`func (o *Webhook) UnsetStoreId()`

UnsetStoreId ensures that no value is present for StoreId, not even an explicit nil
### GetLangId

`func (o *Webhook) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *Webhook) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *Webhook) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *Webhook) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### SetLangIdNil

`func (o *Webhook) SetLangIdNil(b bool)`

 SetLangIdNil sets the value for LangId to be an explicit nil

### UnsetLangId
`func (o *Webhook) UnsetLangId()`

UnsetLangId ensures that no value is present for LangId, not even an explicit nil
### GetActive

`func (o *Webhook) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Webhook) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Webhook) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Webhook) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *Webhook) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *Webhook) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetCallback

`func (o *Webhook) GetCallback() string`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *Webhook) GetCallbackOk() (*string, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *Webhook) SetCallback(v string)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *Webhook) HasCallback() bool`

HasCallback returns a boolean if a field has been set.

### SetCallbackNil

`func (o *Webhook) SetCallbackNil(b bool)`

 SetCallbackNil sets the value for Callback to be an explicit nil

### UnsetCallback
`func (o *Webhook) UnsetCallback()`

UnsetCallback ensures that no value is present for Callback, not even an explicit nil
### GetFields

`func (o *Webhook) GetFields() string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Webhook) GetFieldsOk() (*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Webhook) SetFields(v string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Webhook) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *Webhook) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *Webhook) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetResponseFields

`func (o *Webhook) GetResponseFields() string`

GetResponseFields returns the ResponseFields field if non-nil, zero value otherwise.

### GetResponseFieldsOk

`func (o *Webhook) GetResponseFieldsOk() (*string, bool)`

GetResponseFieldsOk returns a tuple with the ResponseFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFields

`func (o *Webhook) SetResponseFields(v string)`

SetResponseFields sets ResponseFields field to given value.

### HasResponseFields

`func (o *Webhook) HasResponseFields() bool`

HasResponseFields returns a boolean if a field has been set.

### SetResponseFieldsNil

`func (o *Webhook) SetResponseFieldsNil(b bool)`

 SetResponseFieldsNil sets the value for ResponseFields to be an explicit nil

### UnsetResponseFields
`func (o *Webhook) UnsetResponseFields()`

UnsetResponseFields ensures that no value is present for ResponseFields, not even an explicit nil
### GetCreatedAt

`func (o *Webhook) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Webhook) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Webhook) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Webhook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Webhook) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Webhook) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Webhook) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Webhook) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Webhook) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Webhook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Webhook) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Webhook) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetEntity

`func (o *Webhook) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Webhook) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Webhook) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *Webhook) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntityNil

`func (o *Webhook) SetEntityNil(b bool)`

 SetEntityNil sets the value for Entity to be an explicit nil

### UnsetEntity
`func (o *Webhook) UnsetEntity()`

UnsetEntity ensures that no value is present for Entity, not even an explicit nil
### GetAction

`func (o *Webhook) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Webhook) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Webhook) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Webhook) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *Webhook) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *Webhook) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetAdditionalFields

`func (o *Webhook) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Webhook) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Webhook) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Webhook) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *Webhook) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *Webhook) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *Webhook) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Webhook) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Webhook) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Webhook) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *Webhook) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *Webhook) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


