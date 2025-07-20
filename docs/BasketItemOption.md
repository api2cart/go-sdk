# BasketItemOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**ValueId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**UsedInCombination** | Pointer to **NullableBool** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBasketItemOption

`func NewBasketItemOption() *BasketItemOption`

NewBasketItemOption instantiates a new BasketItemOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasketItemOptionWithDefaults

`func NewBasketItemOptionWithDefaults() *BasketItemOption`

NewBasketItemOptionWithDefaults instantiates a new BasketItemOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasketItemOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasketItemOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasketItemOption) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BasketItemOption) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BasketItemOption) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BasketItemOption) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetValueId

`func (o *BasketItemOption) GetValueId() string`

GetValueId returns the ValueId field if non-nil, zero value otherwise.

### GetValueIdOk

`func (o *BasketItemOption) GetValueIdOk() (*string, bool)`

GetValueIdOk returns a tuple with the ValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueId

`func (o *BasketItemOption) SetValueId(v string)`

SetValueId sets ValueId field to given value.

### HasValueId

`func (o *BasketItemOption) HasValueId() bool`

HasValueId returns a boolean if a field has been set.

### SetValueIdNil

`func (o *BasketItemOption) SetValueIdNil(b bool)`

 SetValueIdNil sets the value for ValueId to be an explicit nil

### UnsetValueId
`func (o *BasketItemOption) UnsetValueId()`

UnsetValueId ensures that no value is present for ValueId, not even an explicit nil
### GetName

`func (o *BasketItemOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasketItemOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasketItemOption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasketItemOption) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BasketItemOption) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BasketItemOption) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *BasketItemOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BasketItemOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BasketItemOption) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BasketItemOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *BasketItemOption) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *BasketItemOption) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetUsedInCombination

`func (o *BasketItemOption) GetUsedInCombination() bool`

GetUsedInCombination returns the UsedInCombination field if non-nil, zero value otherwise.

### GetUsedInCombinationOk

`func (o *BasketItemOption) GetUsedInCombinationOk() (*bool, bool)`

GetUsedInCombinationOk returns a tuple with the UsedInCombination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedInCombination

`func (o *BasketItemOption) SetUsedInCombination(v bool)`

SetUsedInCombination sets UsedInCombination field to given value.

### HasUsedInCombination

`func (o *BasketItemOption) HasUsedInCombination() bool`

HasUsedInCombination returns a boolean if a field has been set.

### SetUsedInCombinationNil

`func (o *BasketItemOption) SetUsedInCombinationNil(b bool)`

 SetUsedInCombinationNil sets the value for UsedInCombination to be an explicit nil

### UnsetUsedInCombination
`func (o *BasketItemOption) UnsetUsedInCombination()`

UnsetUsedInCombination ensures that no value is present for UsedInCombination, not even an explicit nil
### GetAdditionalFields

`func (o *BasketItemOption) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *BasketItemOption) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *BasketItemOption) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *BasketItemOption) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *BasketItemOption) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *BasketItemOption) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *BasketItemOption) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BasketItemOption) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BasketItemOption) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BasketItemOption) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *BasketItemOption) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *BasketItemOption) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


