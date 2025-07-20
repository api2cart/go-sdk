# Discount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ModifierType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 
**FromTime** | Pointer to **NullableString** |  | [optional] 
**ToTime** | Pointer to **NullableString** |  | [optional] 
**CustomerGroupIds** | Pointer to **NullableString** |  | [optional] 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDiscount

`func NewDiscount() *Discount`

NewDiscount instantiates a new Discount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountWithDefaults

`func NewDiscountWithDefaults() *Discount`

NewDiscountWithDefaults instantiates a new Discount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Discount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Discount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Discount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Discount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Discount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Discount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Discount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Discount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModifierType

`func (o *Discount) GetModifierType() string`

GetModifierType returns the ModifierType field if non-nil, zero value otherwise.

### GetModifierTypeOk

`func (o *Discount) GetModifierTypeOk() (*string, bool)`

GetModifierTypeOk returns a tuple with the ModifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierType

`func (o *Discount) SetModifierType(v string)`

SetModifierType sets ModifierType field to given value.

### HasModifierType

`func (o *Discount) HasModifierType() bool`

HasModifierType returns a boolean if a field has been set.

### GetValue

`func (o *Discount) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Discount) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Discount) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Discount) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetFromTime

`func (o *Discount) GetFromTime() string`

GetFromTime returns the FromTime field if non-nil, zero value otherwise.

### GetFromTimeOk

`func (o *Discount) GetFromTimeOk() (*string, bool)`

GetFromTimeOk returns a tuple with the FromTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTime

`func (o *Discount) SetFromTime(v string)`

SetFromTime sets FromTime field to given value.

### HasFromTime

`func (o *Discount) HasFromTime() bool`

HasFromTime returns a boolean if a field has been set.

### SetFromTimeNil

`func (o *Discount) SetFromTimeNil(b bool)`

 SetFromTimeNil sets the value for FromTime to be an explicit nil

### UnsetFromTime
`func (o *Discount) UnsetFromTime()`

UnsetFromTime ensures that no value is present for FromTime, not even an explicit nil
### GetToTime

`func (o *Discount) GetToTime() string`

GetToTime returns the ToTime field if non-nil, zero value otherwise.

### GetToTimeOk

`func (o *Discount) GetToTimeOk() (*string, bool)`

GetToTimeOk returns a tuple with the ToTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTime

`func (o *Discount) SetToTime(v string)`

SetToTime sets ToTime field to given value.

### HasToTime

`func (o *Discount) HasToTime() bool`

HasToTime returns a boolean if a field has been set.

### SetToTimeNil

`func (o *Discount) SetToTimeNil(b bool)`

 SetToTimeNil sets the value for ToTime to be an explicit nil

### UnsetToTime
`func (o *Discount) UnsetToTime()`

UnsetToTime ensures that no value is present for ToTime, not even an explicit nil
### GetCustomerGroupIds

`func (o *Discount) GetCustomerGroupIds() string`

GetCustomerGroupIds returns the CustomerGroupIds field if non-nil, zero value otherwise.

### GetCustomerGroupIdsOk

`func (o *Discount) GetCustomerGroupIdsOk() (*string, bool)`

GetCustomerGroupIdsOk returns a tuple with the CustomerGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerGroupIds

`func (o *Discount) SetCustomerGroupIds(v string)`

SetCustomerGroupIds sets CustomerGroupIds field to given value.

### HasCustomerGroupIds

`func (o *Discount) HasCustomerGroupIds() bool`

HasCustomerGroupIds returns a boolean if a field has been set.

### SetCustomerGroupIdsNil

`func (o *Discount) SetCustomerGroupIdsNil(b bool)`

 SetCustomerGroupIdsNil sets the value for CustomerGroupIds to be an explicit nil

### UnsetCustomerGroupIds
`func (o *Discount) UnsetCustomerGroupIds()`

UnsetCustomerGroupIds ensures that no value is present for CustomerGroupIds, not even an explicit nil
### GetSortOrder

`func (o *Discount) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Discount) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Discount) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Discount) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *Discount) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *Discount) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetAdditionalFields

`func (o *Discount) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Discount) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Discount) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Discount) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *Discount) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *Discount) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *Discount) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Discount) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Discount) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Discount) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *Discount) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *Discount) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


