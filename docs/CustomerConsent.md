# CustomerConsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**OptInLevel** | Pointer to **NullableString** |  | [optional] 
**ModifiedTime** | Pointer to **NullableString** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCustomerConsent

`func NewCustomerConsent() *CustomerConsent`

NewCustomerConsent instantiates a new CustomerConsent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerConsentWithDefaults

`func NewCustomerConsentWithDefaults() *CustomerConsent`

NewCustomerConsentWithDefaults instantiates a new CustomerConsent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerConsent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerConsent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerConsent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerConsent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomerConsent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerConsent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerConsent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomerConsent) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CustomerConsent) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomerConsent) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStatus

`func (o *CustomerConsent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerConsent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerConsent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerConsent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CustomerConsent) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CustomerConsent) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSource

`func (o *CustomerConsent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CustomerConsent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CustomerConsent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CustomerConsent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *CustomerConsent) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *CustomerConsent) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetOptInLevel

`func (o *CustomerConsent) GetOptInLevel() string`

GetOptInLevel returns the OptInLevel field if non-nil, zero value otherwise.

### GetOptInLevelOk

`func (o *CustomerConsent) GetOptInLevelOk() (*string, bool)`

GetOptInLevelOk returns a tuple with the OptInLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInLevel

`func (o *CustomerConsent) SetOptInLevel(v string)`

SetOptInLevel sets OptInLevel field to given value.

### HasOptInLevel

`func (o *CustomerConsent) HasOptInLevel() bool`

HasOptInLevel returns a boolean if a field has been set.

### SetOptInLevelNil

`func (o *CustomerConsent) SetOptInLevelNil(b bool)`

 SetOptInLevelNil sets the value for OptInLevel to be an explicit nil

### UnsetOptInLevel
`func (o *CustomerConsent) UnsetOptInLevel()`

UnsetOptInLevel ensures that no value is present for OptInLevel, not even an explicit nil
### GetModifiedTime

`func (o *CustomerConsent) GetModifiedTime() string`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *CustomerConsent) GetModifiedTimeOk() (*string, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *CustomerConsent) SetModifiedTime(v string)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *CustomerConsent) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### SetModifiedTimeNil

`func (o *CustomerConsent) SetModifiedTimeNil(b bool)`

 SetModifiedTimeNil sets the value for ModifiedTime to be an explicit nil

### UnsetModifiedTime
`func (o *CustomerConsent) UnsetModifiedTime()`

UnsetModifiedTime ensures that no value is present for ModifiedTime, not even an explicit nil
### GetAdditionalFields

`func (o *CustomerConsent) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CustomerConsent) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CustomerConsent) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CustomerConsent) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *CustomerConsent) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *CustomerConsent) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *CustomerConsent) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CustomerConsent) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CustomerConsent) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CustomerConsent) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CustomerConsent) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CustomerConsent) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


