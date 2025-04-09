# ProductInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarehouseId** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**InStock** | Pointer to **bool** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProductInventory

`func NewProductInventory() *ProductInventory`

NewProductInventory instantiates a new ProductInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductInventoryWithDefaults

`func NewProductInventoryWithDefaults() *ProductInventory`

NewProductInventoryWithDefaults instantiates a new ProductInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouseId

`func (o *ProductInventory) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ProductInventory) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ProductInventory) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ProductInventory) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductInventory) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductInventory) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductInventory) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductInventory) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetInStock

`func (o *ProductInventory) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *ProductInventory) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *ProductInventory) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *ProductInventory) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### GetPriority

`func (o *ProductInventory) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ProductInventory) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ProductInventory) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ProductInventory) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ProductInventory) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ProductInventory) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ProductInventory) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ProductInventory) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ProductInventory) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProductInventory) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProductInventory) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProductInventory) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


