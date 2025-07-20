# ProductOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductOptionId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**Available** | Pointer to **NullableBool** |  | [optional] 
**UsedInCombination** | Pointer to **NullableBool** |  | [optional] 
**OptionItems** | Pointer to [**[]ProductOptionItem**](ProductOptionItem.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProductOption

`func NewProductOption() *ProductOption`

NewProductOption instantiates a new ProductOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductOptionWithDefaults

`func NewProductOptionWithDefaults() *ProductOption`

NewProductOptionWithDefaults instantiates a new ProductOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductOption) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductOption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductOptionId

`func (o *ProductOption) GetProductOptionId() string`

GetProductOptionId returns the ProductOptionId field if non-nil, zero value otherwise.

### GetProductOptionIdOk

`func (o *ProductOption) GetProductOptionIdOk() (*string, bool)`

GetProductOptionIdOk returns a tuple with the ProductOptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOptionId

`func (o *ProductOption) SetProductOptionId(v string)`

SetProductOptionId sets ProductOptionId field to given value.

### HasProductOptionId

`func (o *ProductOption) HasProductOptionId() bool`

HasProductOptionId returns a boolean if a field has been set.

### SetProductOptionIdNil

`func (o *ProductOption) SetProductOptionIdNil(b bool)`

 SetProductOptionIdNil sets the value for ProductOptionId to be an explicit nil

### UnsetProductOptionId
`func (o *ProductOption) UnsetProductOptionId()`

UnsetProductOptionId ensures that no value is present for ProductOptionId, not even an explicit nil
### GetName

`func (o *ProductOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductOption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductOption) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProductOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductOption) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductOption) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProductOption) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductOption) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSortOrder

`func (o *ProductOption) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ProductOption) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ProductOption) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ProductOption) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *ProductOption) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *ProductOption) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetType

`func (o *ProductOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequired

`func (o *ProductOption) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ProductOption) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ProductOption) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ProductOption) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ProductOption) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ProductOption) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetAvailable

`func (o *ProductOption) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ProductOption) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ProductOption) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ProductOption) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *ProductOption) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *ProductOption) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetUsedInCombination

`func (o *ProductOption) GetUsedInCombination() bool`

GetUsedInCombination returns the UsedInCombination field if non-nil, zero value otherwise.

### GetUsedInCombinationOk

`func (o *ProductOption) GetUsedInCombinationOk() (*bool, bool)`

GetUsedInCombinationOk returns a tuple with the UsedInCombination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedInCombination

`func (o *ProductOption) SetUsedInCombination(v bool)`

SetUsedInCombination sets UsedInCombination field to given value.

### HasUsedInCombination

`func (o *ProductOption) HasUsedInCombination() bool`

HasUsedInCombination returns a boolean if a field has been set.

### SetUsedInCombinationNil

`func (o *ProductOption) SetUsedInCombinationNil(b bool)`

 SetUsedInCombinationNil sets the value for UsedInCombination to be an explicit nil

### UnsetUsedInCombination
`func (o *ProductOption) UnsetUsedInCombination()`

UnsetUsedInCombination ensures that no value is present for UsedInCombination, not even an explicit nil
### GetOptionItems

`func (o *ProductOption) GetOptionItems() []ProductOptionItem`

GetOptionItems returns the OptionItems field if non-nil, zero value otherwise.

### GetOptionItemsOk

`func (o *ProductOption) GetOptionItemsOk() (*[]ProductOptionItem, bool)`

GetOptionItemsOk returns a tuple with the OptionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionItems

`func (o *ProductOption) SetOptionItems(v []ProductOptionItem)`

SetOptionItems sets OptionItems field to given value.

### HasOptionItems

`func (o *ProductOption) HasOptionItems() bool`

HasOptionItems returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ProductOption) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ProductOption) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ProductOption) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ProductOption) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ProductOption) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ProductOption) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ProductOption) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProductOption) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProductOption) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProductOption) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ProductOption) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ProductOption) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


