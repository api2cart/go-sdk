# CouponAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **string** |  | [optional] 
**ApplyTo** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**IncludeTax** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DiscountedQuantity** | Pointer to **float32** |  | [optional] 
**DiscountQuantityStep** | Pointer to **int32** |  | [optional] 
**LogicOperator** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]CouponCondition**](CouponCondition.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCouponAction

`func NewCouponAction() *CouponAction`

NewCouponAction instantiates a new CouponAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponActionWithDefaults

`func NewCouponActionWithDefaults() *CouponAction`

NewCouponActionWithDefaults instantiates a new CouponAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *CouponAction) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CouponAction) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CouponAction) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CouponAction) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetApplyTo

`func (o *CouponAction) GetApplyTo() string`

GetApplyTo returns the ApplyTo field if non-nil, zero value otherwise.

### GetApplyToOk

`func (o *CouponAction) GetApplyToOk() (*string, bool)`

GetApplyToOk returns a tuple with the ApplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTo

`func (o *CouponAction) SetApplyTo(v string)`

SetApplyTo sets ApplyTo field to given value.

### HasApplyTo

`func (o *CouponAction) HasApplyTo() bool`

HasApplyTo returns a boolean if a field has been set.

### GetAmount

`func (o *CouponAction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CouponAction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CouponAction) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CouponAction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *CouponAction) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CouponAction) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CouponAction) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CouponAction) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetIncludeTax

`func (o *CouponAction) GetIncludeTax() bool`

GetIncludeTax returns the IncludeTax field if non-nil, zero value otherwise.

### GetIncludeTaxOk

`func (o *CouponAction) GetIncludeTaxOk() (*bool, bool)`

GetIncludeTaxOk returns a tuple with the IncludeTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTax

`func (o *CouponAction) SetIncludeTax(v bool)`

SetIncludeTax sets IncludeTax field to given value.

### HasIncludeTax

`func (o *CouponAction) HasIncludeTax() bool`

HasIncludeTax returns a boolean if a field has been set.

### GetType

`func (o *CouponAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CouponAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDiscountedQuantity

`func (o *CouponAction) GetDiscountedQuantity() float32`

GetDiscountedQuantity returns the DiscountedQuantity field if non-nil, zero value otherwise.

### GetDiscountedQuantityOk

`func (o *CouponAction) GetDiscountedQuantityOk() (*float32, bool)`

GetDiscountedQuantityOk returns a tuple with the DiscountedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedQuantity

`func (o *CouponAction) SetDiscountedQuantity(v float32)`

SetDiscountedQuantity sets DiscountedQuantity field to given value.

### HasDiscountedQuantity

`func (o *CouponAction) HasDiscountedQuantity() bool`

HasDiscountedQuantity returns a boolean if a field has been set.

### GetDiscountQuantityStep

`func (o *CouponAction) GetDiscountQuantityStep() int32`

GetDiscountQuantityStep returns the DiscountQuantityStep field if non-nil, zero value otherwise.

### GetDiscountQuantityStepOk

`func (o *CouponAction) GetDiscountQuantityStepOk() (*int32, bool)`

GetDiscountQuantityStepOk returns a tuple with the DiscountQuantityStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountQuantityStep

`func (o *CouponAction) SetDiscountQuantityStep(v int32)`

SetDiscountQuantityStep sets DiscountQuantityStep field to given value.

### HasDiscountQuantityStep

`func (o *CouponAction) HasDiscountQuantityStep() bool`

HasDiscountQuantityStep returns a boolean if a field has been set.

### GetLogicOperator

`func (o *CouponAction) GetLogicOperator() string`

GetLogicOperator returns the LogicOperator field if non-nil, zero value otherwise.

### GetLogicOperatorOk

`func (o *CouponAction) GetLogicOperatorOk() (*string, bool)`

GetLogicOperatorOk returns a tuple with the LogicOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicOperator

`func (o *CouponAction) SetLogicOperator(v string)`

SetLogicOperator sets LogicOperator field to given value.

### HasLogicOperator

`func (o *CouponAction) HasLogicOperator() bool`

HasLogicOperator returns a boolean if a field has been set.

### GetConditions

`func (o *CouponAction) GetConditions() []CouponCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CouponAction) GetConditionsOk() (*[]CouponCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CouponAction) SetConditions(v []CouponCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CouponAction) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *CouponAction) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CouponAction) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CouponAction) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CouponAction) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *CouponAction) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CouponAction) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CouponAction) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CouponAction) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


