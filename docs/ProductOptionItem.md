# ProductOptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProductOptionItemId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**TypePrice** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProductOptionItem

`func NewProductOptionItem() *ProductOptionItem`

NewProductOptionItem instantiates a new ProductOptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductOptionItemWithDefaults

`func NewProductOptionItemWithDefaults() *ProductOptionItem`

NewProductOptionItemWithDefaults instantiates a new ProductOptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductOptionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductOptionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductOptionItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductOptionItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductOptionItemId

`func (o *ProductOptionItem) GetProductOptionItemId() string`

GetProductOptionItemId returns the ProductOptionItemId field if non-nil, zero value otherwise.

### GetProductOptionItemIdOk

`func (o *ProductOptionItem) GetProductOptionItemIdOk() (*string, bool)`

GetProductOptionItemIdOk returns a tuple with the ProductOptionItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOptionItemId

`func (o *ProductOptionItem) SetProductOptionItemId(v string)`

SetProductOptionItemId sets ProductOptionItemId field to given value.

### HasProductOptionItemId

`func (o *ProductOptionItem) HasProductOptionItemId() bool`

HasProductOptionItemId returns a boolean if a field has been set.

### GetName

`func (o *ProductOptionItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductOptionItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductOptionItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductOptionItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSortOrder

`func (o *ProductOptionItem) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ProductOptionItem) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ProductOptionItem) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ProductOptionItem) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetPrice

`func (o *ProductOptionItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductOptionItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductOptionItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductOptionItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetWeight

`func (o *ProductOptionItem) GetWeight() string`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductOptionItem) GetWeightOk() (*string, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductOptionItem) SetWeight(v string)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductOptionItem) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductOptionItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductOptionItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductOptionItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductOptionItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTypePrice

`func (o *ProductOptionItem) GetTypePrice() string`

GetTypePrice returns the TypePrice field if non-nil, zero value otherwise.

### GetTypePriceOk

`func (o *ProductOptionItem) GetTypePriceOk() (*string, bool)`

GetTypePriceOk returns a tuple with the TypePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypePrice

`func (o *ProductOptionItem) SetTypePrice(v string)`

SetTypePrice sets TypePrice field to given value.

### HasTypePrice

`func (o *ProductOptionItem) HasTypePrice() bool`

HasTypePrice returns a boolean if a field has been set.

### GetSku

`func (o *ProductOptionItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductOptionItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductOptionItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductOptionItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetIsDefault

`func (o *ProductOptionItem) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ProductOptionItem) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ProductOptionItem) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ProductOptionItem) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ProductOptionItem) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ProductOptionItem) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ProductOptionItem) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ProductOptionItem) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ProductOptionItem) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProductOptionItem) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProductOptionItem) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProductOptionItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


