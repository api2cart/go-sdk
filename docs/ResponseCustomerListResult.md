# ResponseCustomerListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomersCount** | Pointer to **NullableInt32** |  | [optional] 
**Customer** | Pointer to [**[]Customer**](Customer.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseCustomerListResult

`func NewResponseCustomerListResult() *ResponseCustomerListResult`

NewResponseCustomerListResult instantiates a new ResponseCustomerListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCustomerListResultWithDefaults

`func NewResponseCustomerListResultWithDefaults() *ResponseCustomerListResult`

NewResponseCustomerListResultWithDefaults instantiates a new ResponseCustomerListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomersCount

`func (o *ResponseCustomerListResult) GetCustomersCount() int32`

GetCustomersCount returns the CustomersCount field if non-nil, zero value otherwise.

### GetCustomersCountOk

`func (o *ResponseCustomerListResult) GetCustomersCountOk() (*int32, bool)`

GetCustomersCountOk returns a tuple with the CustomersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomersCount

`func (o *ResponseCustomerListResult) SetCustomersCount(v int32)`

SetCustomersCount sets CustomersCount field to given value.

### HasCustomersCount

`func (o *ResponseCustomerListResult) HasCustomersCount() bool`

HasCustomersCount returns a boolean if a field has been set.

### SetCustomersCountNil

`func (o *ResponseCustomerListResult) SetCustomersCountNil(b bool)`

 SetCustomersCountNil sets the value for CustomersCount to be an explicit nil

### UnsetCustomersCount
`func (o *ResponseCustomerListResult) UnsetCustomersCount()`

UnsetCustomersCount ensures that no value is present for CustomersCount, not even an explicit nil
### GetCustomer

`func (o *ResponseCustomerListResult) GetCustomer() []Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *ResponseCustomerListResult) GetCustomerOk() (*[]Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *ResponseCustomerListResult) SetCustomer(v []Customer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *ResponseCustomerListResult) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseCustomerListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseCustomerListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseCustomerListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseCustomerListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseCustomerListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseCustomerListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseCustomerListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseCustomerListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseCustomerListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseCustomerListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseCustomerListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseCustomerListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


