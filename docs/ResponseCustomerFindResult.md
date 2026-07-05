# ResponseCustomerFindResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomersCount** | Pointer to **NullableInt32** |  | [optional] 
**Customer** | Pointer to [**[]Customer**](Customer.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseCustomerFindResult

`func NewResponseCustomerFindResult() *ResponseCustomerFindResult`

NewResponseCustomerFindResult instantiates a new ResponseCustomerFindResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCustomerFindResultWithDefaults

`func NewResponseCustomerFindResultWithDefaults() *ResponseCustomerFindResult`

NewResponseCustomerFindResultWithDefaults instantiates a new ResponseCustomerFindResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomersCount

`func (o *ResponseCustomerFindResult) GetCustomersCount() int32`

GetCustomersCount returns the CustomersCount field if non-nil, zero value otherwise.

### GetCustomersCountOk

`func (o *ResponseCustomerFindResult) GetCustomersCountOk() (*int32, bool)`

GetCustomersCountOk returns a tuple with the CustomersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomersCount

`func (o *ResponseCustomerFindResult) SetCustomersCount(v int32)`

SetCustomersCount sets CustomersCount field to given value.

### HasCustomersCount

`func (o *ResponseCustomerFindResult) HasCustomersCount() bool`

HasCustomersCount returns a boolean if a field has been set.

### SetCustomersCountNil

`func (o *ResponseCustomerFindResult) SetCustomersCountNil(b bool)`

 SetCustomersCountNil sets the value for CustomersCount to be an explicit nil

### UnsetCustomersCount
`func (o *ResponseCustomerFindResult) UnsetCustomersCount()`

UnsetCustomersCount ensures that no value is present for CustomersCount, not even an explicit nil
### GetCustomer

`func (o *ResponseCustomerFindResult) GetCustomer() []Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *ResponseCustomerFindResult) GetCustomerOk() (*[]Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *ResponseCustomerFindResult) SetCustomer(v []Customer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *ResponseCustomerFindResult) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseCustomerFindResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseCustomerFindResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseCustomerFindResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseCustomerFindResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseCustomerFindResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseCustomerFindResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseCustomerFindResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseCustomerFindResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseCustomerFindResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseCustomerFindResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseCustomerFindResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseCustomerFindResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


