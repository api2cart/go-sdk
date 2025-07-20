# CouponCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**UsedTimes** | Pointer to **NullableInt32** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCouponCode

`func NewCouponCode() *CouponCode`

NewCouponCode instantiates a new CouponCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCodeWithDefaults

`func NewCouponCodeWithDefaults() *CouponCode`

NewCouponCodeWithDefaults instantiates a new CouponCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CouponCode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponCode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponCode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CouponCode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *CouponCode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CouponCode) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CouponCode) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CouponCode) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetUsedTimes

`func (o *CouponCode) GetUsedTimes() int32`

GetUsedTimes returns the UsedTimes field if non-nil, zero value otherwise.

### GetUsedTimesOk

`func (o *CouponCode) GetUsedTimesOk() (*int32, bool)`

GetUsedTimesOk returns a tuple with the UsedTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTimes

`func (o *CouponCode) SetUsedTimes(v int32)`

SetUsedTimes sets UsedTimes field to given value.

### HasUsedTimes

`func (o *CouponCode) HasUsedTimes() bool`

HasUsedTimes returns a boolean if a field has been set.

### SetUsedTimesNil

`func (o *CouponCode) SetUsedTimesNil(b bool)`

 SetUsedTimesNil sets the value for UsedTimes to be an explicit nil

### UnsetUsedTimes
`func (o *CouponCode) UnsetUsedTimes()`

UnsetUsedTimes ensures that no value is present for UsedTimes, not even an explicit nil
### GetAdditionalFields

`func (o *CouponCode) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CouponCode) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CouponCode) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CouponCode) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *CouponCode) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *CouponCode) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *CouponCode) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CouponCode) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CouponCode) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CouponCode) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CouponCode) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CouponCode) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


