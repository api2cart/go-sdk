# ResponseCartCouponListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponCount** | Pointer to **NullableInt32** |  | [optional] 
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

### SetCouponCountNil

`func (o *ResponseCartCouponListResult) SetCouponCountNil(b bool)`

 SetCouponCountNil sets the value for CouponCount to be an explicit nil

### UnsetCouponCount
`func (o *ResponseCartCouponListResult) UnsetCouponCount()`

UnsetCouponCount ensures that no value is present for CouponCount, not even an explicit nil
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

### SetAdditionalFieldsNil

`func (o *ResponseCartCouponListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseCartCouponListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
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

### SetCustomFieldsNil

`func (o *ResponseCartCouponListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseCartCouponListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


