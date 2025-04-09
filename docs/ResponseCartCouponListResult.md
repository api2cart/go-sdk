# ResponseCartCouponListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponCount** | Pointer to **int32** |  | [optional] 
**Coupon** | Pointer to [**[]Coupon**](Coupon.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseCartCouponListResult

`func NewResponseCartCouponListResult() *ResponseCartCouponListResult`

NewResponseCartCouponListResult instantiates a new ResponseCartCouponListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCartCouponListResultWithDefaults

`func NewResponseCartCouponListResultWithDefaults() *ResponseCartCouponListResult`

NewResponseCartCouponListResultWithDefaults instantiates a new ResponseCartCouponListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouponCount

`func (o *ResponseCartCouponListResult) GetCouponCount() int32`

GetCouponCount returns the CouponCount field if non-nil, zero value otherwise.

### GetCouponCountOk

`func (o *ResponseCartCouponListResult) GetCouponCountOk() (*int32, bool)`

GetCouponCountOk returns a tuple with the CouponCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCount

`func (o *ResponseCartCouponListResult) SetCouponCount(v int32)`

SetCouponCount sets CouponCount field to given value.

### HasCouponCount

`func (o *ResponseCartCouponListResult) HasCouponCount() bool`

HasCouponCount returns a boolean if a field has been set.

### GetCoupon

`func (o *ResponseCartCouponListResult) GetCoupon() []Coupon`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *ResponseCartCouponListResult) GetCouponOk() (*[]Coupon, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *ResponseCartCouponListResult) SetCoupon(v []Coupon)`

SetCoupon sets Coupon field to given value.

### HasCoupon

`func (o *ResponseCartCouponListResult) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseCartCouponListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseCartCouponListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseCartCouponListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseCartCouponListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseCartCouponListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseCartCouponListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseCartCouponListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseCartCouponListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


