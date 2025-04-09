# AccountCartList200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartsCount** | Pointer to **int32** |  | [optional] 
**Carts** | Pointer to [**[]AccountCartList200ResponseResultCartsInner**](AccountCartList200ResponseResultCartsInner.md) |  | [optional] 

## Methods

### NewAccountCartList200ResponseResult

`func NewAccountCartList200ResponseResult() *AccountCartList200ResponseResult`

NewAccountCartList200ResponseResult instantiates a new AccountCartList200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCartList200ResponseResultWithDefaults

`func NewAccountCartList200ResponseResultWithDefaults() *AccountCartList200ResponseResult`

NewAccountCartList200ResponseResultWithDefaults instantiates a new AccountCartList200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartsCount

`func (o *AccountCartList200ResponseResult) GetCartsCount() int32`

GetCartsCount returns the CartsCount field if non-nil, zero value otherwise.

### GetCartsCountOk

`func (o *AccountCartList200ResponseResult) GetCartsCountOk() (*int32, bool)`

GetCartsCountOk returns a tuple with the CartsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartsCount

`func (o *AccountCartList200ResponseResult) SetCartsCount(v int32)`

SetCartsCount sets CartsCount field to given value.

### HasCartsCount

`func (o *AccountCartList200ResponseResult) HasCartsCount() bool`

HasCartsCount returns a boolean if a field has been set.

### GetCarts

`func (o *AccountCartList200ResponseResult) GetCarts() []AccountCartList200ResponseResultCartsInner`

GetCarts returns the Carts field if non-nil, zero value otherwise.

### GetCartsOk

`func (o *AccountCartList200ResponseResult) GetCartsOk() (*[]AccountCartList200ResponseResultCartsInner, bool)`

GetCartsOk returns a tuple with the Carts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarts

`func (o *AccountCartList200ResponseResult) SetCarts(v []AccountCartList200ResponseResultCartsInner)`

SetCarts sets Carts field to given value.

### HasCarts

`func (o *AccountCartList200ResponseResult) HasCarts() bool`

HasCarts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


