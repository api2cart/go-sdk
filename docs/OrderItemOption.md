# OrderItemOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Weight** | Pointer to **NullableFloat32** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ProductOptionValueId** | Pointer to **NullableString** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderItemOption

`func NewOrderItemOption() *OrderItemOption`

NewOrderItemOption instantiates a new OrderItemOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemOptionWithDefaults

`func NewOrderItemOptionWithDefaults() *OrderItemOption`

NewOrderItemOptionWithDefaults instantiates a new OrderItemOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionId

`func (o *OrderItemOption) GetOptionId() string`

GetOptionId returns the OptionId field if non-nil, zero value otherwise.

### GetOptionIdOk

`func (o *OrderItemOption) GetOptionIdOk() (*string, bool)`

GetOptionIdOk returns a tuple with the OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionId

`func (o *OrderItemOption) SetOptionId(v string)`

SetOptionId sets OptionId field to given value.

### HasOptionId

`func (o *OrderItemOption) HasOptionId() bool`

HasOptionId returns a boolean if a field has been set.

### SetOptionIdNil

`func (o *OrderItemOption) SetOptionIdNil(b bool)`

 SetOptionIdNil sets the value for OptionId to be an explicit nil

### UnsetOptionId
`func (o *OrderItemOption) UnsetOptionId()`

UnsetOptionId ensures that no value is present for OptionId, not even an explicit nil
### GetName

`func (o *OrderItemOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderItemOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderItemOption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderItemOption) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *OrderItemOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OrderItemOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OrderItemOption) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OrderItemOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPrice

`func (o *OrderItemOption) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderItemOption) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderItemOption) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderItemOption) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetWeight

`func (o *OrderItemOption) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OrderItemOption) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OrderItemOption) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *OrderItemOption) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *OrderItemOption) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *OrderItemOption) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetType

`func (o *OrderItemOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderItemOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderItemOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderItemOption) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *OrderItemOption) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderItemOption) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetProductOptionValueId

`func (o *OrderItemOption) GetProductOptionValueId() string`

GetProductOptionValueId returns the ProductOptionValueId field if non-nil, zero value otherwise.

### GetProductOptionValueIdOk

`func (o *OrderItemOption) GetProductOptionValueIdOk() (*string, bool)`

GetProductOptionValueIdOk returns a tuple with the ProductOptionValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOptionValueId

`func (o *OrderItemOption) SetProductOptionValueId(v string)`

SetProductOptionValueId sets ProductOptionValueId field to given value.

### HasProductOptionValueId

`func (o *OrderItemOption) HasProductOptionValueId() bool`

HasProductOptionValueId returns a boolean if a field has been set.

### SetProductOptionValueIdNil

`func (o *OrderItemOption) SetProductOptionValueIdNil(b bool)`

 SetProductOptionValueIdNil sets the value for ProductOptionValueId to be an explicit nil

### UnsetProductOptionValueId
`func (o *OrderItemOption) UnsetProductOptionValueId()`

UnsetProductOptionValueId ensures that no value is present for ProductOptionValueId, not even an explicit nil
### GetAdditionalFields

`func (o *OrderItemOption) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderItemOption) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderItemOption) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderItemOption) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *OrderItemOption) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *OrderItemOption) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *OrderItemOption) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderItemOption) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderItemOption) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderItemOption) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *OrderItemOption) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *OrderItemOption) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


