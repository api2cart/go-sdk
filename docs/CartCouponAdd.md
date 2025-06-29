# CartCouponAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Coupon code | 
**ActionType** | **string** | Coupon discount type | 
**ActionApplyTo** | **string** | Defines where discount should be applied | 
**ActionScope** | **string** | Specify how discount should be applied. If scope&#x3D;matching_items, then discount will be applied to each of the items that match action conditions. Scope order means that discount will be applied once. | 
**ActionAmount** | **float32** | Defines the discount amount value. | 
**Codes** | Pointer to **[]string** | Entity codes | [optional] 
**Name** | Pointer to **string** | Coupon name | [optional] 
**DateStart** | Pointer to **string** | Date start | [optional] [default to "now"]
**DateEnd** | Pointer to **string** | Defines when discount code will be expired. | [optional] 
**UsageLimit** | Pointer to **int32** | Usage limit for coupon. | [optional] 
**UsageLimitPerCustomer** | Pointer to **int32** | Usage limit per customer. | [optional] 
**ActionConditionEntity** | Pointer to **string** | Defines entity for action condition. | [optional] 
**ActionConditionKey** | Pointer to **string** | Defines entity attribute code for action condition. | [optional] 
**ActionConditionOperator** | Pointer to **string** | Defines condition operator. | [optional] 
**ActionConditionValue** | Pointer to **string** | Defines condition attribute value/s. Can be comma separated string. | [optional] 
**IncludeTax** | Pointer to **bool** | Indicates whether to apply a discount for taxes. | [optional] [default to false]
**StoreId** | Pointer to **string** | Store Id | [optional] 
**FreeCashOnDelivery** | Pointer to **bool** | Defines whether the coupon provides free cash on delivery | [optional] 

## Methods

### NewCartCouponAdd

`func NewCartCouponAdd(code string, actionType string, actionApplyTo string, actionScope string, actionAmount float32, ) *CartCouponAdd`

NewCartCouponAdd instantiates a new CartCouponAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartCouponAddWithDefaults

`func NewCartCouponAddWithDefaults() *CartCouponAdd`

NewCartCouponAddWithDefaults instantiates a new CartCouponAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CartCouponAdd) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CartCouponAdd) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CartCouponAdd) SetCode(v string)`

SetCode sets Code field to given value.


### GetActionType

`func (o *CartCouponAdd) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *CartCouponAdd) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *CartCouponAdd) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetActionApplyTo

`func (o *CartCouponAdd) GetActionApplyTo() string`

GetActionApplyTo returns the ActionApplyTo field if non-nil, zero value otherwise.

### GetActionApplyToOk

`func (o *CartCouponAdd) GetActionApplyToOk() (*string, bool)`

GetActionApplyToOk returns a tuple with the ActionApplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionApplyTo

`func (o *CartCouponAdd) SetActionApplyTo(v string)`

SetActionApplyTo sets ActionApplyTo field to given value.


### GetActionScope

`func (o *CartCouponAdd) GetActionScope() string`

GetActionScope returns the ActionScope field if non-nil, zero value otherwise.

### GetActionScopeOk

`func (o *CartCouponAdd) GetActionScopeOk() (*string, bool)`

GetActionScopeOk returns a tuple with the ActionScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionScope

`func (o *CartCouponAdd) SetActionScope(v string)`

SetActionScope sets ActionScope field to given value.


### GetActionAmount

`func (o *CartCouponAdd) GetActionAmount() float32`

GetActionAmount returns the ActionAmount field if non-nil, zero value otherwise.

### GetActionAmountOk

`func (o *CartCouponAdd) GetActionAmountOk() (*float32, bool)`

GetActionAmountOk returns a tuple with the ActionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionAmount

`func (o *CartCouponAdd) SetActionAmount(v float32)`

SetActionAmount sets ActionAmount field to given value.


### GetCodes

`func (o *CartCouponAdd) GetCodes() []string`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *CartCouponAdd) GetCodesOk() (*[]string, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *CartCouponAdd) SetCodes(v []string)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *CartCouponAdd) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetName

`func (o *CartCouponAdd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartCouponAdd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartCouponAdd) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CartCouponAdd) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDateStart

`func (o *CartCouponAdd) GetDateStart() string`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *CartCouponAdd) GetDateStartOk() (*string, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *CartCouponAdd) SetDateStart(v string)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *CartCouponAdd) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateEnd

`func (o *CartCouponAdd) GetDateEnd() string`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *CartCouponAdd) GetDateEndOk() (*string, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *CartCouponAdd) SetDateEnd(v string)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *CartCouponAdd) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### GetUsageLimit

`func (o *CartCouponAdd) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *CartCouponAdd) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *CartCouponAdd) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *CartCouponAdd) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetUsageLimitPerCustomer

`func (o *CartCouponAdd) GetUsageLimitPerCustomer() int32`

GetUsageLimitPerCustomer returns the UsageLimitPerCustomer field if non-nil, zero value otherwise.

### GetUsageLimitPerCustomerOk

`func (o *CartCouponAdd) GetUsageLimitPerCustomerOk() (*int32, bool)`

GetUsageLimitPerCustomerOk returns a tuple with the UsageLimitPerCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimitPerCustomer

`func (o *CartCouponAdd) SetUsageLimitPerCustomer(v int32)`

SetUsageLimitPerCustomer sets UsageLimitPerCustomer field to given value.

### HasUsageLimitPerCustomer

`func (o *CartCouponAdd) HasUsageLimitPerCustomer() bool`

HasUsageLimitPerCustomer returns a boolean if a field has been set.

### GetActionConditionEntity

`func (o *CartCouponAdd) GetActionConditionEntity() string`

GetActionConditionEntity returns the ActionConditionEntity field if non-nil, zero value otherwise.

### GetActionConditionEntityOk

`func (o *CartCouponAdd) GetActionConditionEntityOk() (*string, bool)`

GetActionConditionEntityOk returns a tuple with the ActionConditionEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConditionEntity

`func (o *CartCouponAdd) SetActionConditionEntity(v string)`

SetActionConditionEntity sets ActionConditionEntity field to given value.

### HasActionConditionEntity

`func (o *CartCouponAdd) HasActionConditionEntity() bool`

HasActionConditionEntity returns a boolean if a field has been set.

### GetActionConditionKey

`func (o *CartCouponAdd) GetActionConditionKey() string`

GetActionConditionKey returns the ActionConditionKey field if non-nil, zero value otherwise.

### GetActionConditionKeyOk

`func (o *CartCouponAdd) GetActionConditionKeyOk() (*string, bool)`

GetActionConditionKeyOk returns a tuple with the ActionConditionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConditionKey

`func (o *CartCouponAdd) SetActionConditionKey(v string)`

SetActionConditionKey sets ActionConditionKey field to given value.

### HasActionConditionKey

`func (o *CartCouponAdd) HasActionConditionKey() bool`

HasActionConditionKey returns a boolean if a field has been set.

### GetActionConditionOperator

`func (o *CartCouponAdd) GetActionConditionOperator() string`

GetActionConditionOperator returns the ActionConditionOperator field if non-nil, zero value otherwise.

### GetActionConditionOperatorOk

`func (o *CartCouponAdd) GetActionConditionOperatorOk() (*string, bool)`

GetActionConditionOperatorOk returns a tuple with the ActionConditionOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConditionOperator

`func (o *CartCouponAdd) SetActionConditionOperator(v string)`

SetActionConditionOperator sets ActionConditionOperator field to given value.

### HasActionConditionOperator

`func (o *CartCouponAdd) HasActionConditionOperator() bool`

HasActionConditionOperator returns a boolean if a field has been set.

### GetActionConditionValue

`func (o *CartCouponAdd) GetActionConditionValue() string`

GetActionConditionValue returns the ActionConditionValue field if non-nil, zero value otherwise.

### GetActionConditionValueOk

`func (o *CartCouponAdd) GetActionConditionValueOk() (*string, bool)`

GetActionConditionValueOk returns a tuple with the ActionConditionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionConditionValue

`func (o *CartCouponAdd) SetActionConditionValue(v string)`

SetActionConditionValue sets ActionConditionValue field to given value.

### HasActionConditionValue

`func (o *CartCouponAdd) HasActionConditionValue() bool`

HasActionConditionValue returns a boolean if a field has been set.

### GetIncludeTax

`func (o *CartCouponAdd) GetIncludeTax() bool`

GetIncludeTax returns the IncludeTax field if non-nil, zero value otherwise.

### GetIncludeTaxOk

`func (o *CartCouponAdd) GetIncludeTaxOk() (*bool, bool)`

GetIncludeTaxOk returns a tuple with the IncludeTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTax

`func (o *CartCouponAdd) SetIncludeTax(v bool)`

SetIncludeTax sets IncludeTax field to given value.

### HasIncludeTax

`func (o *CartCouponAdd) HasIncludeTax() bool`

HasIncludeTax returns a boolean if a field has been set.

### GetStoreId

`func (o *CartCouponAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CartCouponAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CartCouponAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *CartCouponAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetFreeCashOnDelivery

`func (o *CartCouponAdd) GetFreeCashOnDelivery() bool`

GetFreeCashOnDelivery returns the FreeCashOnDelivery field if non-nil, zero value otherwise.

### GetFreeCashOnDeliveryOk

`func (o *CartCouponAdd) GetFreeCashOnDeliveryOk() (*bool, bool)`

GetFreeCashOnDeliveryOk returns a tuple with the FreeCashOnDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCashOnDelivery

`func (o *CartCouponAdd) SetFreeCashOnDelivery(v bool)`

SetFreeCashOnDelivery sets FreeCashOnDelivery field to given value.

### HasFreeCashOnDelivery

`func (o *CartCouponAdd) HasFreeCashOnDelivery() bool`

HasFreeCashOnDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


