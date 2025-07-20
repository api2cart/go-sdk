# CatalogPriceRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Gid** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ShortDescription** | Pointer to **NullableString** |  | [optional] 
**Avail** | Pointer to **NullableBool** |  | [optional] 
**Actions** | Pointer to [**[]CatalogPriceRuleAction**](CatalogPriceRuleAction.md) |  | [optional] 
**CreatedTime** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**DateStart** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**DateEnd** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**UsageCount** | Pointer to **NullableFloat32** |  | [optional] 
**Conditions** | Pointer to [**[]CouponCondition**](CouponCondition.md) |  | [optional] 
**UsesPerOrderLimit** | Pointer to **NullableInt32** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCatalogPriceRule

`func NewCatalogPriceRule() *CatalogPriceRule`

NewCatalogPriceRule instantiates a new CatalogPriceRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogPriceRuleWithDefaults

`func NewCatalogPriceRuleWithDefaults() *CatalogPriceRule`

NewCatalogPriceRuleWithDefaults instantiates a new CatalogPriceRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogPriceRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogPriceRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogPriceRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogPriceRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGid

`func (o *CatalogPriceRule) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *CatalogPriceRule) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *CatalogPriceRule) SetGid(v string)`

SetGid sets Gid field to given value.

### HasGid

`func (o *CatalogPriceRule) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *CatalogPriceRule) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *CatalogPriceRule) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetType

`func (o *CatalogPriceRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogPriceRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogPriceRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogPriceRule) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CatalogPriceRule) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CatalogPriceRule) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *CatalogPriceRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogPriceRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogPriceRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogPriceRule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CatalogPriceRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CatalogPriceRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CatalogPriceRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogPriceRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogPriceRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogPriceRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogPriceRule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogPriceRule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetShortDescription

`func (o *CatalogPriceRule) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *CatalogPriceRule) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *CatalogPriceRule) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *CatalogPriceRule) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### SetShortDescriptionNil

`func (o *CatalogPriceRule) SetShortDescriptionNil(b bool)`

 SetShortDescriptionNil sets the value for ShortDescription to be an explicit nil

### UnsetShortDescription
`func (o *CatalogPriceRule) UnsetShortDescription()`

UnsetShortDescription ensures that no value is present for ShortDescription, not even an explicit nil
### GetAvail

`func (o *CatalogPriceRule) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *CatalogPriceRule) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *CatalogPriceRule) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *CatalogPriceRule) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### SetAvailNil

`func (o *CatalogPriceRule) SetAvailNil(b bool)`

 SetAvailNil sets the value for Avail to be an explicit nil

### UnsetAvail
`func (o *CatalogPriceRule) UnsetAvail()`

UnsetAvail ensures that no value is present for Avail, not even an explicit nil
### GetActions

`func (o *CatalogPriceRule) GetActions() []CatalogPriceRuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CatalogPriceRule) GetActionsOk() (*[]CatalogPriceRuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CatalogPriceRule) SetActions(v []CatalogPriceRuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CatalogPriceRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CatalogPriceRule) GetCreatedTime() A2CDateTime`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CatalogPriceRule) GetCreatedTimeOk() (*A2CDateTime, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CatalogPriceRule) SetCreatedTime(v A2CDateTime)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CatalogPriceRule) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *CatalogPriceRule) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *CatalogPriceRule) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetDateStart

`func (o *CatalogPriceRule) GetDateStart() A2CDateTime`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *CatalogPriceRule) GetDateStartOk() (*A2CDateTime, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *CatalogPriceRule) SetDateStart(v A2CDateTime)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *CatalogPriceRule) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### SetDateStartNil

`func (o *CatalogPriceRule) SetDateStartNil(b bool)`

 SetDateStartNil sets the value for DateStart to be an explicit nil

### UnsetDateStart
`func (o *CatalogPriceRule) UnsetDateStart()`

UnsetDateStart ensures that no value is present for DateStart, not even an explicit nil
### GetDateEnd

`func (o *CatalogPriceRule) GetDateEnd() A2CDateTime`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *CatalogPriceRule) GetDateEndOk() (*A2CDateTime, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *CatalogPriceRule) SetDateEnd(v A2CDateTime)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *CatalogPriceRule) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### SetDateEndNil

`func (o *CatalogPriceRule) SetDateEndNil(b bool)`

 SetDateEndNil sets the value for DateEnd to be an explicit nil

### UnsetDateEnd
`func (o *CatalogPriceRule) UnsetDateEnd()`

UnsetDateEnd ensures that no value is present for DateEnd, not even an explicit nil
### GetUsageCount

`func (o *CatalogPriceRule) GetUsageCount() float32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *CatalogPriceRule) GetUsageCountOk() (*float32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *CatalogPriceRule) SetUsageCount(v float32)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *CatalogPriceRule) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### SetUsageCountNil

`func (o *CatalogPriceRule) SetUsageCountNil(b bool)`

 SetUsageCountNil sets the value for UsageCount to be an explicit nil

### UnsetUsageCount
`func (o *CatalogPriceRule) UnsetUsageCount()`

UnsetUsageCount ensures that no value is present for UsageCount, not even an explicit nil
### GetConditions

`func (o *CatalogPriceRule) GetConditions() []CouponCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CatalogPriceRule) GetConditionsOk() (*[]CouponCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CatalogPriceRule) SetConditions(v []CouponCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CatalogPriceRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetUsesPerOrderLimit

`func (o *CatalogPriceRule) GetUsesPerOrderLimit() int32`

GetUsesPerOrderLimit returns the UsesPerOrderLimit field if non-nil, zero value otherwise.

### GetUsesPerOrderLimitOk

`func (o *CatalogPriceRule) GetUsesPerOrderLimitOk() (*int32, bool)`

GetUsesPerOrderLimitOk returns a tuple with the UsesPerOrderLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesPerOrderLimit

`func (o *CatalogPriceRule) SetUsesPerOrderLimit(v int32)`

SetUsesPerOrderLimit sets UsesPerOrderLimit field to given value.

### HasUsesPerOrderLimit

`func (o *CatalogPriceRule) HasUsesPerOrderLimit() bool`

HasUsesPerOrderLimit returns a boolean if a field has been set.

### SetUsesPerOrderLimitNil

`func (o *CatalogPriceRule) SetUsesPerOrderLimitNil(b bool)`

 SetUsesPerOrderLimitNil sets the value for UsesPerOrderLimit to be an explicit nil

### UnsetUsesPerOrderLimit
`func (o *CatalogPriceRule) UnsetUsesPerOrderLimit()`

UnsetUsesPerOrderLimit ensures that no value is present for UsesPerOrderLimit, not even an explicit nil
### GetAdditionalFields

`func (o *CatalogPriceRule) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CatalogPriceRule) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CatalogPriceRule) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CatalogPriceRule) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *CatalogPriceRule) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *CatalogPriceRule) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *CatalogPriceRule) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CatalogPriceRule) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CatalogPriceRule) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CatalogPriceRule) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CatalogPriceRule) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CatalogPriceRule) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


