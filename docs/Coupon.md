# Coupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Codes** | Pointer to [**[]CouponCode**](CouponCode.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Actions** | Pointer to [**[]CouponAction**](CouponAction.md) |  | [optional] 
**DateStart** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**DateEnd** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**Avail** | Pointer to **bool** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**UsedTimes** | Pointer to **int32** |  | [optional] 
**UsageLimit** | Pointer to **int32** |  | [optional] 
**UsageLimitPerCustomer** | Pointer to **int32** |  | [optional] 
**LogicOperator** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]CouponCondition**](CouponCondition.md) |  | [optional] 
**UsageHistory** | Pointer to [**[]CouponHistory**](CouponHistory.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCoupon

`func NewCoupon() *Coupon`

NewCoupon instantiates a new Coupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponWithDefaults

`func NewCouponWithDefaults() *Coupon`

NewCouponWithDefaults instantiates a new Coupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Coupon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Coupon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Coupon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Coupon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *Coupon) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Coupon) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Coupon) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Coupon) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCodes

`func (o *Coupon) GetCodes() []CouponCode`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *Coupon) GetCodesOk() (*[]CouponCode, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *Coupon) SetCodes(v []CouponCode)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *Coupon) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetName

`func (o *Coupon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Coupon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Coupon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Coupon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Coupon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Coupon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Coupon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Coupon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActions

`func (o *Coupon) GetActions() []CouponAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Coupon) GetActionsOk() (*[]CouponAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Coupon) SetActions(v []CouponAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Coupon) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDateStart

`func (o *Coupon) GetDateStart() A2CDateTime`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *Coupon) GetDateStartOk() (*A2CDateTime, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *Coupon) SetDateStart(v A2CDateTime)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *Coupon) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateEnd

`func (o *Coupon) GetDateEnd() A2CDateTime`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *Coupon) GetDateEndOk() (*A2CDateTime, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *Coupon) SetDateEnd(v A2CDateTime)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *Coupon) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### GetAvail

`func (o *Coupon) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *Coupon) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *Coupon) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *Coupon) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetPriority

`func (o *Coupon) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Coupon) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Coupon) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Coupon) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUsedTimes

`func (o *Coupon) GetUsedTimes() int32`

GetUsedTimes returns the UsedTimes field if non-nil, zero value otherwise.

### GetUsedTimesOk

`func (o *Coupon) GetUsedTimesOk() (*int32, bool)`

GetUsedTimesOk returns a tuple with the UsedTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTimes

`func (o *Coupon) SetUsedTimes(v int32)`

SetUsedTimes sets UsedTimes field to given value.

### HasUsedTimes

`func (o *Coupon) HasUsedTimes() bool`

HasUsedTimes returns a boolean if a field has been set.

### GetUsageLimit

`func (o *Coupon) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *Coupon) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *Coupon) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *Coupon) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetUsageLimitPerCustomer

`func (o *Coupon) GetUsageLimitPerCustomer() int32`

GetUsageLimitPerCustomer returns the UsageLimitPerCustomer field if non-nil, zero value otherwise.

### GetUsageLimitPerCustomerOk

`func (o *Coupon) GetUsageLimitPerCustomerOk() (*int32, bool)`

GetUsageLimitPerCustomerOk returns a tuple with the UsageLimitPerCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimitPerCustomer

`func (o *Coupon) SetUsageLimitPerCustomer(v int32)`

SetUsageLimitPerCustomer sets UsageLimitPerCustomer field to given value.

### HasUsageLimitPerCustomer

`func (o *Coupon) HasUsageLimitPerCustomer() bool`

HasUsageLimitPerCustomer returns a boolean if a field has been set.

### GetLogicOperator

`func (o *Coupon) GetLogicOperator() string`

GetLogicOperator returns the LogicOperator field if non-nil, zero value otherwise.

### GetLogicOperatorOk

`func (o *Coupon) GetLogicOperatorOk() (*string, bool)`

GetLogicOperatorOk returns a tuple with the LogicOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicOperator

`func (o *Coupon) SetLogicOperator(v string)`

SetLogicOperator sets LogicOperator field to given value.

### HasLogicOperator

`func (o *Coupon) HasLogicOperator() bool`

HasLogicOperator returns a boolean if a field has been set.

### GetConditions

`func (o *Coupon) GetConditions() []CouponCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Coupon) GetConditionsOk() (*[]CouponCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Coupon) SetConditions(v []CouponCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Coupon) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetUsageHistory

`func (o *Coupon) GetUsageHistory() []CouponHistory`

GetUsageHistory returns the UsageHistory field if non-nil, zero value otherwise.

### GetUsageHistoryOk

`func (o *Coupon) GetUsageHistoryOk() (*[]CouponHistory, bool)`

GetUsageHistoryOk returns a tuple with the UsageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageHistory

`func (o *Coupon) SetUsageHistory(v []CouponHistory)`

SetUsageHistory sets UsageHistory field to given value.

### HasUsageHistory

`func (o *Coupon) HasUsageHistory() bool`

HasUsageHistory returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Coupon) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Coupon) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Coupon) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Coupon) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *Coupon) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Coupon) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Coupon) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Coupon) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


