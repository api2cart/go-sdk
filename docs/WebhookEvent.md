# WebhookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookName** | Pointer to **NullableString** |  | [optional] 
**Entity** | Pointer to **NullableString** |  | [optional] 
**Action** | Pointer to **NullableString** |  | [optional] 
**FilterableFields** | Pointer to **map[string]interface{}** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWebhookEvent

`func NewWebhookEvent() *WebhookEvent`

NewWebhookEvent instantiates a new WebhookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventWithDefaults

`func NewWebhookEventWithDefaults() *WebhookEvent`

NewWebhookEventWithDefaults instantiates a new WebhookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookName

`func (o *WebhookEvent) GetWebhookName() string`

GetWebhookName returns the WebhookName field if non-nil, zero value otherwise.

### GetWebhookNameOk

`func (o *WebhookEvent) GetWebhookNameOk() (*string, bool)`

GetWebhookNameOk returns a tuple with the WebhookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookName

`func (o *WebhookEvent) SetWebhookName(v string)`

SetWebhookName sets WebhookName field to given value.

### HasWebhookName

`func (o *WebhookEvent) HasWebhookName() bool`

HasWebhookName returns a boolean if a field has been set.

### SetWebhookNameNil

`func (o *WebhookEvent) SetWebhookNameNil(b bool)`

 SetWebhookNameNil sets the value for WebhookName to be an explicit nil

### UnsetWebhookName
`func (o *WebhookEvent) UnsetWebhookName()`

UnsetWebhookName ensures that no value is present for WebhookName, not even an explicit nil
### GetEntity

`func (o *WebhookEvent) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *WebhookEvent) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *WebhookEvent) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *WebhookEvent) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntityNil

`func (o *WebhookEvent) SetEntityNil(b bool)`

 SetEntityNil sets the value for Entity to be an explicit nil

### UnsetEntity
`func (o *WebhookEvent) UnsetEntity()`

UnsetEntity ensures that no value is present for Entity, not even an explicit nil
### GetAction

`func (o *WebhookEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WebhookEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WebhookEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WebhookEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *WebhookEvent) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *WebhookEvent) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetFilterableFields

`func (o *WebhookEvent) GetFilterableFields() map[string]interface{}`

GetFilterableFields returns the FilterableFields field if non-nil, zero value otherwise.

### GetFilterableFieldsOk

`func (o *WebhookEvent) GetFilterableFieldsOk() (*map[string]interface{}, bool)`

GetFilterableFieldsOk returns a tuple with the FilterableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterableFields

`func (o *WebhookEvent) SetFilterableFields(v map[string]interface{})`

SetFilterableFields sets FilterableFields field to given value.

### HasFilterableFields

`func (o *WebhookEvent) HasFilterableFields() bool`

HasFilterableFields returns a boolean if a field has been set.

### SetFilterableFieldsNil

`func (o *WebhookEvent) SetFilterableFieldsNil(b bool)`

 SetFilterableFieldsNil sets the value for FilterableFields to be an explicit nil

### UnsetFilterableFields
`func (o *WebhookEvent) UnsetFilterableFields()`

UnsetFilterableFields ensures that no value is present for FilterableFields, not even an explicit nil
### GetAdditionalFields

`func (o *WebhookEvent) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *WebhookEvent) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *WebhookEvent) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *WebhookEvent) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *WebhookEvent) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *WebhookEvent) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *WebhookEvent) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WebhookEvent) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WebhookEvent) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WebhookEvent) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *WebhookEvent) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *WebhookEvent) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


