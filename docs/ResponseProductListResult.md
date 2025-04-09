# ResponseProductListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductsCount** | Pointer to **int32** |  | [optional] 
**Product** | Pointer to [**[]Product**](Product.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseProductListResult

`func NewResponseProductListResult() *ResponseProductListResult`

NewResponseProductListResult instantiates a new ResponseProductListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseProductListResultWithDefaults

`func NewResponseProductListResultWithDefaults() *ResponseProductListResult`

NewResponseProductListResultWithDefaults instantiates a new ResponseProductListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductsCount

`func (o *ResponseProductListResult) GetProductsCount() int32`

GetProductsCount returns the ProductsCount field if non-nil, zero value otherwise.

### GetProductsCountOk

`func (o *ResponseProductListResult) GetProductsCountOk() (*int32, bool)`

GetProductsCountOk returns a tuple with the ProductsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsCount

`func (o *ResponseProductListResult) SetProductsCount(v int32)`

SetProductsCount sets ProductsCount field to given value.

### HasProductsCount

`func (o *ResponseProductListResult) HasProductsCount() bool`

HasProductsCount returns a boolean if a field has been set.

### GetProduct

`func (o *ResponseProductListResult) GetProduct() []Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ResponseProductListResult) GetProductOk() (*[]Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ResponseProductListResult) SetProduct(v []Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ResponseProductListResult) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseProductListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseProductListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseProductListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseProductListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseProductListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseProductListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseProductListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseProductListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


