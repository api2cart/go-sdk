# AccountSupportedPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartId** | Pointer to **NullableString** |  | [optional] 
**CartName** | Pointer to **NullableString** |  | [optional] 
**CartVersions** | Pointer to **NullableString** |  | [optional] 
**CartMethod** | Pointer to **NullableString** |  | [optional] 
**Params** | Pointer to [**AccountSupportedPlatformParams**](AccountSupportedPlatformParams.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAccountSupportedPlatform

`func NewAccountSupportedPlatform() *AccountSupportedPlatform`

NewAccountSupportedPlatform instantiates a new AccountSupportedPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSupportedPlatformWithDefaults

`func NewAccountSupportedPlatformWithDefaults() *AccountSupportedPlatform`

NewAccountSupportedPlatformWithDefaults instantiates a new AccountSupportedPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartId

`func (o *AccountSupportedPlatform) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *AccountSupportedPlatform) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *AccountSupportedPlatform) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *AccountSupportedPlatform) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *AccountSupportedPlatform) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *AccountSupportedPlatform) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetCartName

`func (o *AccountSupportedPlatform) GetCartName() string`

GetCartName returns the CartName field if non-nil, zero value otherwise.

### GetCartNameOk

`func (o *AccountSupportedPlatform) GetCartNameOk() (*string, bool)`

GetCartNameOk returns a tuple with the CartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartName

`func (o *AccountSupportedPlatform) SetCartName(v string)`

SetCartName sets CartName field to given value.

### HasCartName

`func (o *AccountSupportedPlatform) HasCartName() bool`

HasCartName returns a boolean if a field has been set.

### SetCartNameNil

`func (o *AccountSupportedPlatform) SetCartNameNil(b bool)`

 SetCartNameNil sets the value for CartName to be an explicit nil

### UnsetCartName
`func (o *AccountSupportedPlatform) UnsetCartName()`

UnsetCartName ensures that no value is present for CartName, not even an explicit nil
### GetCartVersions

`func (o *AccountSupportedPlatform) GetCartVersions() string`

GetCartVersions returns the CartVersions field if non-nil, zero value otherwise.

### GetCartVersionsOk

`func (o *AccountSupportedPlatform) GetCartVersionsOk() (*string, bool)`

GetCartVersionsOk returns a tuple with the CartVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartVersions

`func (o *AccountSupportedPlatform) SetCartVersions(v string)`

SetCartVersions sets CartVersions field to given value.

### HasCartVersions

`func (o *AccountSupportedPlatform) HasCartVersions() bool`

HasCartVersions returns a boolean if a field has been set.

### SetCartVersionsNil

`func (o *AccountSupportedPlatform) SetCartVersionsNil(b bool)`

 SetCartVersionsNil sets the value for CartVersions to be an explicit nil

### UnsetCartVersions
`func (o *AccountSupportedPlatform) UnsetCartVersions()`

UnsetCartVersions ensures that no value is present for CartVersions, not even an explicit nil
### GetCartMethod

`func (o *AccountSupportedPlatform) GetCartMethod() string`

GetCartMethod returns the CartMethod field if non-nil, zero value otherwise.

### GetCartMethodOk

`func (o *AccountSupportedPlatform) GetCartMethodOk() (*string, bool)`

GetCartMethodOk returns a tuple with the CartMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartMethod

`func (o *AccountSupportedPlatform) SetCartMethod(v string)`

SetCartMethod sets CartMethod field to given value.

### HasCartMethod

`func (o *AccountSupportedPlatform) HasCartMethod() bool`

HasCartMethod returns a boolean if a field has been set.

### SetCartMethodNil

`func (o *AccountSupportedPlatform) SetCartMethodNil(b bool)`

 SetCartMethodNil sets the value for CartMethod to be an explicit nil

### UnsetCartMethod
`func (o *AccountSupportedPlatform) UnsetCartMethod()`

UnsetCartMethod ensures that no value is present for CartMethod, not even an explicit nil
### GetParams

`func (o *AccountSupportedPlatform) GetParams() AccountSupportedPlatformParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *AccountSupportedPlatform) GetParamsOk() (*AccountSupportedPlatformParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *AccountSupportedPlatform) SetParams(v AccountSupportedPlatformParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *AccountSupportedPlatform) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *AccountSupportedPlatform) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *AccountSupportedPlatform) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *AccountSupportedPlatform) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *AccountSupportedPlatform) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *AccountSupportedPlatform) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *AccountSupportedPlatform) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *AccountSupportedPlatform) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AccountSupportedPlatform) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AccountSupportedPlatform) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AccountSupportedPlatform) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *AccountSupportedPlatform) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *AccountSupportedPlatform) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


