# BasketItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**VariantId** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Tax** | Pointer to **float32** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**WeightUnit** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] 
**Options** | Pointer to [**[]BasketItemOption**](BasketItemOption.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBasketItem

`func NewBasketItem() *BasketItem`

NewBasketItem instantiates a new BasketItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasketItemWithDefaults

`func NewBasketItemWithDefaults() *BasketItem`

NewBasketItemWithDefaults instantiates a new BasketItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasketItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasketItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasketItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BasketItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *BasketItem) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BasketItem) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BasketItem) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BasketItem) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProductId

`func (o *BasketItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *BasketItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *BasketItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *BasketItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetVariantId

`func (o *BasketItem) GetVariantId() string`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *BasketItem) GetVariantIdOk() (*string, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *BasketItem) SetVariantId(v string)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *BasketItem) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.

### GetSku

`func (o *BasketItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *BasketItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *BasketItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *BasketItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetName

`func (o *BasketItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasketItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasketItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasketItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *BasketItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BasketItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BasketItem) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BasketItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTax

`func (o *BasketItem) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *BasketItem) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *BasketItem) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *BasketItem) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetQuantity

`func (o *BasketItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BasketItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BasketItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BasketItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetWeightUnit

`func (o *BasketItem) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *BasketItem) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *BasketItem) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *BasketItem) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetWeight

`func (o *BasketItem) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BasketItem) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BasketItem) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BasketItem) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetOptions

`func (o *BasketItem) GetOptions() []BasketItemOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BasketItem) GetOptionsOk() (*[]BasketItemOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BasketItem) SetOptions(v []BasketItemOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BasketItem) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *BasketItem) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *BasketItem) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *BasketItem) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *BasketItem) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *BasketItem) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BasketItem) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BasketItem) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BasketItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


