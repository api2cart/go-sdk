# PluginList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllPlugins** | Pointer to **int32** |  | [optional] 
**Plugins** | Pointer to [**[]Plugin**](Plugin.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPluginList

`func NewPluginList() *PluginList`

NewPluginList instantiates a new PluginList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginListWithDefaults

`func NewPluginListWithDefaults() *PluginList`

NewPluginListWithDefaults instantiates a new PluginList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllPlugins

`func (o *PluginList) GetAllPlugins() int32`

GetAllPlugins returns the AllPlugins field if non-nil, zero value otherwise.

### GetAllPluginsOk

`func (o *PluginList) GetAllPluginsOk() (*int32, bool)`

GetAllPluginsOk returns a tuple with the AllPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlugins

`func (o *PluginList) SetAllPlugins(v int32)`

SetAllPlugins sets AllPlugins field to given value.

### HasAllPlugins

`func (o *PluginList) HasAllPlugins() bool`

HasAllPlugins returns a boolean if a field has been set.

### GetPlugins

`func (o *PluginList) GetPlugins() []Plugin`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *PluginList) GetPluginsOk() (*[]Plugin, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *PluginList) SetPlugins(v []Plugin)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *PluginList) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *PluginList) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *PluginList) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *PluginList) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *PluginList) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *PluginList) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PluginList) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PluginList) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PluginList) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


