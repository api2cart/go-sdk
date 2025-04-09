# OrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** |  | [optional] 
**OrderProductId** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**PriceIncTax** | Pointer to **float32** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**DiscountAmount** | Pointer to **float32** |  | [optional] 
**TotalPrice** | Pointer to **float32** |  | [optional] 
**TaxPercent** | Pointer to **float32** |  | [optional] 
**TaxValue** | Pointer to **float32** |  | [optional] 
**TaxValueAfterDiscount** | Pointer to **float32** |  | [optional] 
**Options** | Pointer to [**[]OrderItemOption**](OrderItemOption.md) |  | [optional] 
**VariantId** | Pointer to **string** |  | [optional] 
**WeightUnit** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] 
**Barcode** | Pointer to **string** |  | [optional] 
**ParentOrderProductId** | Pointer to **string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderItem

`func NewOrderItem() *OrderItem`

NewOrderItem instantiates a new OrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemWithDefaults

`func NewOrderItemWithDefaults() *OrderItem`

NewOrderItemWithDefaults instantiates a new OrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *OrderItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *OrderItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *OrderItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *OrderItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetOrderProductId

`func (o *OrderItem) GetOrderProductId() string`

GetOrderProductId returns the OrderProductId field if non-nil, zero value otherwise.

### GetOrderProductIdOk

`func (o *OrderItem) GetOrderProductIdOk() (*string, bool)`

GetOrderProductIdOk returns a tuple with the OrderProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductId

`func (o *OrderItem) SetOrderProductId(v string)`

SetOrderProductId sets OrderProductId field to given value.

### HasOrderProductId

`func (o *OrderItem) HasOrderProductId() bool`

HasOrderProductId returns a boolean if a field has been set.

### GetModel

`func (o *OrderItem) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OrderItem) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OrderItem) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OrderItem) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *OrderItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *OrderItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderItem) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceIncTax

`func (o *OrderItem) GetPriceIncTax() float32`

GetPriceIncTax returns the PriceIncTax field if non-nil, zero value otherwise.

### GetPriceIncTaxOk

`func (o *OrderItem) GetPriceIncTaxOk() (*float32, bool)`

GetPriceIncTaxOk returns a tuple with the PriceIncTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceIncTax

`func (o *OrderItem) SetPriceIncTax(v float32)`

SetPriceIncTax sets PriceIncTax field to given value.

### HasPriceIncTax

`func (o *OrderItem) HasPriceIncTax() bool`

HasPriceIncTax returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *OrderItem) GetDiscountAmount() float32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *OrderItem) GetDiscountAmountOk() (*float32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *OrderItem) SetDiscountAmount(v float32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *OrderItem) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetTotalPrice

`func (o *OrderItem) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *OrderItem) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *OrderItem) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *OrderItem) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetTaxPercent

`func (o *OrderItem) GetTaxPercent() float32`

GetTaxPercent returns the TaxPercent field if non-nil, zero value otherwise.

### GetTaxPercentOk

`func (o *OrderItem) GetTaxPercentOk() (*float32, bool)`

GetTaxPercentOk returns a tuple with the TaxPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercent

`func (o *OrderItem) SetTaxPercent(v float32)`

SetTaxPercent sets TaxPercent field to given value.

### HasTaxPercent

`func (o *OrderItem) HasTaxPercent() bool`

HasTaxPercent returns a boolean if a field has been set.

### GetTaxValue

`func (o *OrderItem) GetTaxValue() float32`

GetTaxValue returns the TaxValue field if non-nil, zero value otherwise.

### GetTaxValueOk

`func (o *OrderItem) GetTaxValueOk() (*float32, bool)`

GetTaxValueOk returns a tuple with the TaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxValue

`func (o *OrderItem) SetTaxValue(v float32)`

SetTaxValue sets TaxValue field to given value.

### HasTaxValue

`func (o *OrderItem) HasTaxValue() bool`

HasTaxValue returns a boolean if a field has been set.

### GetTaxValueAfterDiscount

`func (o *OrderItem) GetTaxValueAfterDiscount() float32`

GetTaxValueAfterDiscount returns the TaxValueAfterDiscount field if non-nil, zero value otherwise.

### GetTaxValueAfterDiscountOk

`func (o *OrderItem) GetTaxValueAfterDiscountOk() (*float32, bool)`

GetTaxValueAfterDiscountOk returns a tuple with the TaxValueAfterDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxValueAfterDiscount

`func (o *OrderItem) SetTaxValueAfterDiscount(v float32)`

SetTaxValueAfterDiscount sets TaxValueAfterDiscount field to given value.

### HasTaxValueAfterDiscount

`func (o *OrderItem) HasTaxValueAfterDiscount() bool`

HasTaxValueAfterDiscount returns a boolean if a field has been set.

### GetOptions

`func (o *OrderItem) GetOptions() []OrderItemOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OrderItem) GetOptionsOk() (*[]OrderItemOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OrderItem) SetOptions(v []OrderItemOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OrderItem) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetVariantId

`func (o *OrderItem) GetVariantId() string`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *OrderItem) GetVariantIdOk() (*string, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *OrderItem) SetVariantId(v string)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *OrderItem) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.

### GetWeightUnit

`func (o *OrderItem) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *OrderItem) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *OrderItem) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *OrderItem) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetWeight

`func (o *OrderItem) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OrderItem) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OrderItem) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *OrderItem) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetBarcode

`func (o *OrderItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *OrderItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *OrderItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *OrderItem) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetParentOrderProductId

`func (o *OrderItem) GetParentOrderProductId() string`

GetParentOrderProductId returns the ParentOrderProductId field if non-nil, zero value otherwise.

### GetParentOrderProductIdOk

`func (o *OrderItem) GetParentOrderProductIdOk() (*string, bool)`

GetParentOrderProductIdOk returns a tuple with the ParentOrderProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrderProductId

`func (o *OrderItem) SetParentOrderProductId(v string)`

SetParentOrderProductId sets ParentOrderProductId field to given value.

### HasParentOrderProductId

`func (o *OrderItem) HasParentOrderProductId() bool`

HasParentOrderProductId returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *OrderItem) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderItem) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderItem) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderItem) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *OrderItem) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderItem) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderItem) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


