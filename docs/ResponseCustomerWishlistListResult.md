# ResponseCustomerWishlistListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**WishLists** | Pointer to [**[]CustomerWishList**](CustomerWishList.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseCustomerWishlistListResult

`func NewResponseCustomerWishlistListResult() *ResponseCustomerWishlistListResult`

NewResponseCustomerWishlistListResult instantiates a new ResponseCustomerWishlistListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCustomerWishlistListResultWithDefaults

`func NewResponseCustomerWishlistListResultWithDefaults() *ResponseCustomerWishlistListResult`

NewResponseCustomerWishlistListResultWithDefaults instantiates a new ResponseCustomerWishlistListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ResponseCustomerWishlistListResult) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResponseCustomerWishlistListResult) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResponseCustomerWishlistListResult) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ResponseCustomerWishlistListResult) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetWishLists

`func (o *ResponseCustomerWishlistListResult) GetWishLists() []CustomerWishList`

GetWishLists returns the WishLists field if non-nil, zero value otherwise.

### GetWishListsOk

`func (o *ResponseCustomerWishlistListResult) GetWishListsOk() (*[]CustomerWishList, bool)`

GetWishListsOk returns a tuple with the WishLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWishLists

`func (o *ResponseCustomerWishlistListResult) SetWishLists(v []CustomerWishList)`

SetWishLists sets WishLists field to given value.

### HasWishLists

`func (o *ResponseCustomerWishlistListResult) HasWishLists() bool`

HasWishLists returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseCustomerWishlistListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseCustomerWishlistListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseCustomerWishlistListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseCustomerWishlistListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseCustomerWishlistListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseCustomerWishlistListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseCustomerWishlistListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseCustomerWishlistListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


