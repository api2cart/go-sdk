# ProductGroupItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildItemId** | Pointer to **NullableString** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**DefaultQtyInPack** | Pointer to **string** |  | [optional] 
**IsQtyInPackFixed** | Pointer to **NullableBool** |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProductGroupItem

`func NewProductGroupItem() *ProductGroupItem`

NewProductGroupItem instantiates a new ProductGroupItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductGroupItemWithDefaults

`func NewProductGroupItemWithDefaults() *ProductGroupItem`

NewProductGroupItemWithDefaults instantiates a new ProductGroupItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildItemId

`func (o *ProductGroupItem) GetChildItemId() string`

GetChildItemId returns the ChildItemId field if non-nil, zero value otherwise.

### GetChildItemIdOk

`func (o *ProductGroupItem) GetChildItemIdOk() (*string, bool)`

GetChildItemIdOk returns a tuple with the ChildItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildItemId

`func (o *ProductGroupItem) SetChildItemId(v string)`

SetChildItemId sets ChildItemId field to given value.

### HasChildItemId

`func (o *ProductGroupItem) HasChildItemId() bool`

HasChildItemId returns a boolean if a field has been set.

### SetChildItemIdNil

`func (o *ProductGroupItem) SetChildItemIdNil(b bool)`

 SetChildItemIdNil sets the value for ChildItemId to be an explicit nil

### UnsetChildItemId
`func (o *ProductGroupItem) UnsetChildItemId()`

UnsetChildItemId ensures that no value is present for ChildItemId, not even an explicit nil
### GetProductId

`func (o *ProductGroupItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductGroupItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductGroupItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductGroupItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetDefaultQtyInPack

`func (o *ProductGroupItem) GetDefaultQtyInPack() string`

GetDefaultQtyInPack returns the DefaultQtyInPack field if non-nil, zero value otherwise.

### GetDefaultQtyInPackOk

`func (o *ProductGroupItem) GetDefaultQtyInPackOk() (*string, bool)`

GetDefaultQtyInPackOk returns a tuple with the DefaultQtyInPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQtyInPack

`func (o *ProductGroupItem) SetDefaultQtyInPack(v string)`

SetDefaultQtyInPack sets DefaultQtyInPack field to given value.

### HasDefaultQtyInPack

`func (o *ProductGroupItem) HasDefaultQtyInPack() bool`

HasDefaultQtyInPack returns a boolean if a field has been set.

### GetIsQtyInPackFixed

`func (o *ProductGroupItem) GetIsQtyInPackFixed() bool`

GetIsQtyInPackFixed returns the IsQtyInPackFixed field if non-nil, zero value otherwise.

### GetIsQtyInPackFixedOk

`func (o *ProductGroupItem) GetIsQtyInPackFixedOk() (*bool, bool)`

GetIsQtyInPackFixedOk returns a tuple with the IsQtyInPackFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQtyInPackFixed

`func (o *ProductGroupItem) SetIsQtyInPackFixed(v bool)`

SetIsQtyInPackFixed sets IsQtyInPackFixed field to given value.

### HasIsQtyInPackFixed

`func (o *ProductGroupItem) HasIsQtyInPackFixed() bool`

HasIsQtyInPackFixed returns a boolean if a field has been set.

### SetIsQtyInPackFixedNil

`func (o *ProductGroupItem) SetIsQtyInPackFixedNil(b bool)`

 SetIsQtyInPackFixedNil sets the value for IsQtyInPackFixed to be an explicit nil

### UnsetIsQtyInPackFixed
`func (o *ProductGroupItem) UnsetIsQtyInPackFixed()`

UnsetIsQtyInPackFixed ensures that no value is present for IsQtyInPackFixed, not even an explicit nil
### GetPrice

`func (o *ProductGroupItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductGroupItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductGroupItem) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductGroupItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ProductGroupItem) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ProductGroupItem) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetAdditionalFields

`func (o *ProductGroupItem) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ProductGroupItem) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ProductGroupItem) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ProductGroupItem) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ProductGroupItem) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ProductGroupItem) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ProductGroupItem) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProductGroupItem) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProductGroupItem) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProductGroupItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ProductGroupItem) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ProductGroupItem) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


