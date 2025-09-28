# OrderCalculateBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**PriceIncTax** | Pointer to **float32** |  | [optional] 
**TaxRate** | Pointer to **float32** |  | [optional] 
**UnitDiscount** | Pointer to **float32** |  | [optional] 
**Weight** | Pointer to **NullableFloat32** |  | [optional] 
**WeightUnit** | Pointer to **NullableString** |  | [optional] 
**Barcode** | Pointer to **NullableString** |  | [optional] 
**VariantId** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to [**[]OrderItemOption**](OrderItemOption.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderCalculateBundle

`func NewOrderCalculateBundle() *OrderCalculateBundle`

NewOrderCalculateBundle instantiates a new OrderCalculateBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCalculateBundleWithDefaults

`func NewOrderCalculateBundleWithDefaults() *OrderCalculateBundle`

NewOrderCalculateBundleWithDefaults instantiates a new OrderCalculateBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *OrderCalculateBundle) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *OrderCalculateBundle) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *OrderCalculateBundle) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *OrderCalculateBundle) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetSku

`func (o *OrderCalculateBundle) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *OrderCalculateBundle) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *OrderCalculateBundle) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *OrderCalculateBundle) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetName

`func (o *OrderCalculateBundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderCalculateBundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderCalculateBundle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderCalculateBundle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderCalculateBundle) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderCalculateBundle) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderCalculateBundle) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderCalculateBundle) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *OrderCalculateBundle) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderCalculateBundle) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderCalculateBundle) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderCalculateBundle) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceIncTax

`func (o *OrderCalculateBundle) GetPriceIncTax() float32`

GetPriceIncTax returns the PriceIncTax field if non-nil, zero value otherwise.

### GetPriceIncTaxOk

`func (o *OrderCalculateBundle) GetPriceIncTaxOk() (*float32, bool)`

GetPriceIncTaxOk returns a tuple with the PriceIncTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceIncTax

`func (o *OrderCalculateBundle) SetPriceIncTax(v float32)`

SetPriceIncTax sets PriceIncTax field to given value.

### HasPriceIncTax

`func (o *OrderCalculateBundle) HasPriceIncTax() bool`

HasPriceIncTax returns a boolean if a field has been set.

### GetTaxRate

`func (o *OrderCalculateBundle) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *OrderCalculateBundle) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *OrderCalculateBundle) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *OrderCalculateBundle) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetUnitDiscount

`func (o *OrderCalculateBundle) GetUnitDiscount() float32`

GetUnitDiscount returns the UnitDiscount field if non-nil, zero value otherwise.

### GetUnitDiscountOk

`func (o *OrderCalculateBundle) GetUnitDiscountOk() (*float32, bool)`

GetUnitDiscountOk returns a tuple with the UnitDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitDiscount

`func (o *OrderCalculateBundle) SetUnitDiscount(v float32)`

SetUnitDiscount sets UnitDiscount field to given value.

### HasUnitDiscount

`func (o *OrderCalculateBundle) HasUnitDiscount() bool`

HasUnitDiscount returns a boolean if a field has been set.

### GetWeight

`func (o *OrderCalculateBundle) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OrderCalculateBundle) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OrderCalculateBundle) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *OrderCalculateBundle) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *OrderCalculateBundle) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *OrderCalculateBundle) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetWeightUnit

`func (o *OrderCalculateBundle) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *OrderCalculateBundle) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *OrderCalculateBundle) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *OrderCalculateBundle) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *OrderCalculateBundle) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *OrderCalculateBundle) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetBarcode

`func (o *OrderCalculateBundle) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *OrderCalculateBundle) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *OrderCalculateBundle) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *OrderCalculateBundle) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### SetBarcodeNil

`func (o *OrderCalculateBundle) SetBarcodeNil(b bool)`

 SetBarcodeNil sets the value for Barcode to be an explicit nil

### UnsetBarcode
`func (o *OrderCalculateBundle) UnsetBarcode()`

UnsetBarcode ensures that no value is present for Barcode, not even an explicit nil
### GetVariantId

`func (o *OrderCalculateBundle) GetVariantId() string`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *OrderCalculateBundle) GetVariantIdOk() (*string, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *OrderCalculateBundle) SetVariantId(v string)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *OrderCalculateBundle) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.

### SetVariantIdNil

`func (o *OrderCalculateBundle) SetVariantIdNil(b bool)`

 SetVariantIdNil sets the value for VariantId to be an explicit nil

### UnsetVariantId
`func (o *OrderCalculateBundle) UnsetVariantId()`

UnsetVariantId ensures that no value is present for VariantId, not even an explicit nil
### GetOptions

`func (o *OrderCalculateBundle) GetOptions() []OrderItemOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OrderCalculateBundle) GetOptionsOk() (*[]OrderItemOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OrderCalculateBundle) SetOptions(v []OrderItemOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OrderCalculateBundle) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *OrderCalculateBundle) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderCalculateBundle) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderCalculateBundle) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderCalculateBundle) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *OrderCalculateBundle) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *OrderCalculateBundle) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *OrderCalculateBundle) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderCalculateBundle) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderCalculateBundle) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderCalculateBundle) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *OrderCalculateBundle) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *OrderCalculateBundle) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


