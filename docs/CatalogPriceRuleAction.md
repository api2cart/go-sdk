# CatalogPriceRuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **string** |  | [optional] 
**ApplyTo** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **NullableFloat32** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**IncludeTax** | Pointer to **NullableBool** |  | [optional] 
**Conditions** | Pointer to [**[]CouponCondition**](CouponCondition.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCatalogPriceRuleAction

`func NewCatalogPriceRuleAction() *CatalogPriceRuleAction`

NewCatalogPriceRuleAction instantiates a new CatalogPriceRuleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogPriceRuleActionWithDefaults

`func NewCatalogPriceRuleActionWithDefaults() *CatalogPriceRuleAction`

NewCatalogPriceRuleActionWithDefaults instantiates a new CatalogPriceRuleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *CatalogPriceRuleAction) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CatalogPriceRuleAction) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CatalogPriceRuleAction) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CatalogPriceRuleAction) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetApplyTo

`func (o *CatalogPriceRuleAction) GetApplyTo() string`

GetApplyTo returns the ApplyTo field if non-nil, zero value otherwise.

### GetApplyToOk

`func (o *CatalogPriceRuleAction) GetApplyToOk() (*string, bool)`

GetApplyToOk returns a tuple with the ApplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTo

`func (o *CatalogPriceRuleAction) SetApplyTo(v string)`

SetApplyTo sets ApplyTo field to given value.

### HasApplyTo

`func (o *CatalogPriceRuleAction) HasApplyTo() bool`

HasApplyTo returns a boolean if a field has been set.

### GetType

`func (o *CatalogPriceRuleAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogPriceRuleAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogPriceRuleAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogPriceRuleAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuantity

`func (o *CatalogPriceRuleAction) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CatalogPriceRuleAction) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CatalogPriceRuleAction) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CatalogPriceRuleAction) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *CatalogPriceRuleAction) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *CatalogPriceRuleAction) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetValue

`func (o *CatalogPriceRuleAction) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CatalogPriceRuleAction) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CatalogPriceRuleAction) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *CatalogPriceRuleAction) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *CatalogPriceRuleAction) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CatalogPriceRuleAction) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CatalogPriceRuleAction) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CatalogPriceRuleAction) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *CatalogPriceRuleAction) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *CatalogPriceRuleAction) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetIncludeTax

`func (o *CatalogPriceRuleAction) GetIncludeTax() bool`

GetIncludeTax returns the IncludeTax field if non-nil, zero value otherwise.

### GetIncludeTaxOk

`func (o *CatalogPriceRuleAction) GetIncludeTaxOk() (*bool, bool)`

GetIncludeTaxOk returns a tuple with the IncludeTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTax

`func (o *CatalogPriceRuleAction) SetIncludeTax(v bool)`

SetIncludeTax sets IncludeTax field to given value.

### HasIncludeTax

`func (o *CatalogPriceRuleAction) HasIncludeTax() bool`

HasIncludeTax returns a boolean if a field has been set.

### SetIncludeTaxNil

`func (o *CatalogPriceRuleAction) SetIncludeTaxNil(b bool)`

 SetIncludeTaxNil sets the value for IncludeTax to be an explicit nil

### UnsetIncludeTax
`func (o *CatalogPriceRuleAction) UnsetIncludeTax()`

UnsetIncludeTax ensures that no value is present for IncludeTax, not even an explicit nil
### GetConditions

`func (o *CatalogPriceRuleAction) GetConditions() []CouponCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CatalogPriceRuleAction) GetConditionsOk() (*[]CouponCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CatalogPriceRuleAction) SetConditions(v []CouponCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CatalogPriceRuleAction) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *CatalogPriceRuleAction) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CatalogPriceRuleAction) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CatalogPriceRuleAction) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CatalogPriceRuleAction) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *CatalogPriceRuleAction) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *CatalogPriceRuleAction) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *CatalogPriceRuleAction) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CatalogPriceRuleAction) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CatalogPriceRuleAction) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CatalogPriceRuleAction) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CatalogPriceRuleAction) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CatalogPriceRuleAction) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


