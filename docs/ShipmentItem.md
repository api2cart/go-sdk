# ShipmentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderProductId** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**VariantId** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewShipmentItem

`func NewShipmentItem() *ShipmentItem`

NewShipmentItem instantiates a new ShipmentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentItemWithDefaults

`func NewShipmentItemWithDefaults() *ShipmentItem`

NewShipmentItemWithDefaults instantiates a new ShipmentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderProductId

`func (o *ShipmentItem) GetOrderProductId() string`

GetOrderProductId returns the OrderProductId field if non-nil, zero value otherwise.

### GetOrderProductIdOk

`func (o *ShipmentItem) GetOrderProductIdOk() (*string, bool)`

GetOrderProductIdOk returns a tuple with the OrderProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductId

`func (o *ShipmentItem) SetOrderProductId(v string)`

SetOrderProductId sets OrderProductId field to given value.

### HasOrderProductId

`func (o *ShipmentItem) HasOrderProductId() bool`

HasOrderProductId returns a boolean if a field has been set.

### GetProductId

`func (o *ShipmentItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ShipmentItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ShipmentItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ShipmentItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetVariantId

`func (o *ShipmentItem) GetVariantId() string`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *ShipmentItem) GetVariantIdOk() (*string, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *ShipmentItem) SetVariantId(v string)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *ShipmentItem) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.

### GetModel

`func (o *ShipmentItem) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ShipmentItem) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ShipmentItem) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ShipmentItem) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *ShipmentItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipmentItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipmentItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipmentItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *ShipmentItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ShipmentItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ShipmentItem) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ShipmentItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *ShipmentItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ShipmentItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ShipmentItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ShipmentItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ShipmentItem) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ShipmentItem) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ShipmentItem) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ShipmentItem) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ShipmentItem) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ShipmentItem) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ShipmentItem) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ShipmentItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


