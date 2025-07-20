# CouponCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 
**MatchItems** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Operator** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**LogicOperator** | Pointer to **NullableString** |  | [optional] 
**SubConditions** | Pointer to [**[]CouponCondition**](CouponCondition.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCouponCondition

`func NewCouponCondition() *CouponCondition`

NewCouponCondition instantiates a new CouponCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponConditionWithDefaults

`func NewCouponConditionWithDefaults() *CouponCondition`

NewCouponConditionWithDefaults instantiates a new CouponCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CouponCondition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponCondition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponCondition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CouponCondition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEntity

`func (o *CouponCondition) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CouponCondition) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CouponCondition) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *CouponCondition) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetMatchItems

`func (o *CouponCondition) GetMatchItems() string`

GetMatchItems returns the MatchItems field if non-nil, zero value otherwise.

### GetMatchItemsOk

`func (o *CouponCondition) GetMatchItemsOk() (*string, bool)`

GetMatchItemsOk returns a tuple with the MatchItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchItems

`func (o *CouponCondition) SetMatchItems(v string)`

SetMatchItems sets MatchItems field to given value.

### HasMatchItems

`func (o *CouponCondition) HasMatchItems() bool`

HasMatchItems returns a boolean if a field has been set.

### SetMatchItemsNil

`func (o *CouponCondition) SetMatchItemsNil(b bool)`

 SetMatchItemsNil sets the value for MatchItems to be an explicit nil

### UnsetMatchItems
`func (o *CouponCondition) UnsetMatchItems()`

UnsetMatchItems ensures that no value is present for MatchItems, not even an explicit nil
### GetKey

`func (o *CouponCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CouponCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CouponCondition) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CouponCondition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CouponCondition) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CouponCondition) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetOperator

`func (o *CouponCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CouponCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CouponCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CouponCondition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *CouponCondition) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *CouponCondition) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetValue

`func (o *CouponCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CouponCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CouponCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CouponCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CouponCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CouponCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetLogicOperator

`func (o *CouponCondition) GetLogicOperator() string`

GetLogicOperator returns the LogicOperator field if non-nil, zero value otherwise.

### GetLogicOperatorOk

`func (o *CouponCondition) GetLogicOperatorOk() (*string, bool)`

GetLogicOperatorOk returns a tuple with the LogicOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicOperator

`func (o *CouponCondition) SetLogicOperator(v string)`

SetLogicOperator sets LogicOperator field to given value.

### HasLogicOperator

`func (o *CouponCondition) HasLogicOperator() bool`

HasLogicOperator returns a boolean if a field has been set.

### SetLogicOperatorNil

`func (o *CouponCondition) SetLogicOperatorNil(b bool)`

 SetLogicOperatorNil sets the value for LogicOperator to be an explicit nil

### UnsetLogicOperator
`func (o *CouponCondition) UnsetLogicOperator()`

UnsetLogicOperator ensures that no value is present for LogicOperator, not even an explicit nil
### GetSubConditions

`func (o *CouponCondition) GetSubConditions() []CouponCondition`

GetSubConditions returns the SubConditions field if non-nil, zero value otherwise.

### GetSubConditionsOk

`func (o *CouponCondition) GetSubConditionsOk() (*[]CouponCondition, bool)`

GetSubConditionsOk returns a tuple with the SubConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubConditions

`func (o *CouponCondition) SetSubConditions(v []CouponCondition)`

SetSubConditions sets SubConditions field to given value.

### HasSubConditions

`func (o *CouponCondition) HasSubConditions() bool`

HasSubConditions returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *CouponCondition) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CouponCondition) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CouponCondition) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CouponCondition) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *CouponCondition) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *CouponCondition) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *CouponCondition) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CouponCondition) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CouponCondition) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CouponCondition) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CouponCondition) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CouponCondition) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


