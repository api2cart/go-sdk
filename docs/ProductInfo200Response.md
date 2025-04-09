# ProductInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **int32** |  | [optional] 
**ReturnMessage** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**Product**](Product.md) |  | [optional] 

## Methods

### NewProductInfo200Response

`func NewProductInfo200Response() *ProductInfo200Response`

NewProductInfo200Response instantiates a new ProductInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductInfo200ResponseWithDefaults

`func NewProductInfo200ResponseWithDefaults() *ProductInfo200Response`

NewProductInfo200ResponseWithDefaults instantiates a new ProductInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *ProductInfo200Response) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ProductInfo200Response) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ProductInfo200Response) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ProductInfo200Response) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnMessage

`func (o *ProductInfo200Response) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *ProductInfo200Response) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *ProductInfo200Response) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *ProductInfo200Response) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### GetResult

`func (o *ProductInfo200Response) GetResult() Product`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ProductInfo200Response) GetResultOk() (*Product, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ProductInfo200Response) SetResult(v Product)`

SetResult sets Result field to given value.

### HasResult

`func (o *ProductInfo200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


