# CustomerInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **int32** |  | [optional] 
**ReturnMessage** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**Customer**](Customer.md) |  | [optional] 

## Methods

### NewCustomerInfo200Response

`func NewCustomerInfo200Response() *CustomerInfo200Response`

NewCustomerInfo200Response instantiates a new CustomerInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerInfo200ResponseWithDefaults

`func NewCustomerInfo200ResponseWithDefaults() *CustomerInfo200Response`

NewCustomerInfo200ResponseWithDefaults instantiates a new CustomerInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *CustomerInfo200Response) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *CustomerInfo200Response) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *CustomerInfo200Response) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *CustomerInfo200Response) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnMessage

`func (o *CustomerInfo200Response) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *CustomerInfo200Response) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *CustomerInfo200Response) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *CustomerInfo200Response) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### GetResult

`func (o *CustomerInfo200Response) GetResult() Customer`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CustomerInfo200Response) GetResultOk() (*Customer, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CustomerInfo200Response) SetResult(v Customer)`

SetResult sets Result field to given value.

### HasResult

`func (o *CustomerInfo200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


