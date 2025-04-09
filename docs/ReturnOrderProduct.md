# ReturnOrderProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** |  | [optional] 
**OrderProductId** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Reason** | Pointer to [**ReturnReason**](ReturnReason.md) |  | [optional] 
**Action** | Pointer to [**ReturnAction**](ReturnAction.md) |  | [optional] 
**Condition** | Pointer to **string** |  | [optional] 
**CustomerComment** | Pointer to **string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewReturnOrderProduct

`func NewReturnOrderProduct() *ReturnOrderProduct`

NewReturnOrderProduct instantiates a new ReturnOrderProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnOrderProductWithDefaults

`func NewReturnOrderProductWithDefaults() *ReturnOrderProduct`

NewReturnOrderProductWithDefaults instantiates a new ReturnOrderProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ReturnOrderProduct) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ReturnOrderProduct) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ReturnOrderProduct) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ReturnOrderProduct) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetOrderProductId

`func (o *ReturnOrderProduct) GetOrderProductId() string`

GetOrderProductId returns the OrderProductId field if non-nil, zero value otherwise.

### GetOrderProductIdOk

`func (o *ReturnOrderProduct) GetOrderProductIdOk() (*string, bool)`

GetOrderProductIdOk returns a tuple with the OrderProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProductId

`func (o *ReturnOrderProduct) SetOrderProductId(v string)`

SetOrderProductId sets OrderProductId field to given value.

### HasOrderProductId

`func (o *ReturnOrderProduct) HasOrderProductId() bool`

HasOrderProductId returns a boolean if a field has been set.

### GetSku

`func (o *ReturnOrderProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ReturnOrderProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ReturnOrderProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ReturnOrderProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetName

`func (o *ReturnOrderProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReturnOrderProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReturnOrderProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReturnOrderProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *ReturnOrderProduct) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ReturnOrderProduct) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ReturnOrderProduct) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ReturnOrderProduct) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReason

`func (o *ReturnOrderProduct) GetReason() ReturnReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ReturnOrderProduct) GetReasonOk() (*ReturnReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ReturnOrderProduct) SetReason(v ReturnReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ReturnOrderProduct) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAction

`func (o *ReturnOrderProduct) GetAction() ReturnAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ReturnOrderProduct) GetActionOk() (*ReturnAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ReturnOrderProduct) SetAction(v ReturnAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ReturnOrderProduct) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCondition

`func (o *ReturnOrderProduct) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ReturnOrderProduct) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ReturnOrderProduct) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ReturnOrderProduct) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetCustomerComment

`func (o *ReturnOrderProduct) GetCustomerComment() string`

GetCustomerComment returns the CustomerComment field if non-nil, zero value otherwise.

### GetCustomerCommentOk

`func (o *ReturnOrderProduct) GetCustomerCommentOk() (*string, bool)`

GetCustomerCommentOk returns a tuple with the CustomerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerComment

`func (o *ReturnOrderProduct) SetCustomerComment(v string)`

SetCustomerComment sets CustomerComment field to given value.

### HasCustomerComment

`func (o *ReturnOrderProduct) HasCustomerComment() bool`

HasCustomerComment returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ReturnOrderProduct) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ReturnOrderProduct) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ReturnOrderProduct) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ReturnOrderProduct) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ReturnOrderProduct) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ReturnOrderProduct) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ReturnOrderProduct) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ReturnOrderProduct) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


