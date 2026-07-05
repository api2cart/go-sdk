# ResponseCustomerCountResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomersCount** | Pointer to **NullableInt32** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseCustomerCountResult

`func NewResponseCustomerCountResult() *ResponseCustomerCountResult`

NewResponseCustomerCountResult instantiates a new ResponseCustomerCountResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCustomerCountResultWithDefaults

`func NewResponseCustomerCountResultWithDefaults() *ResponseCustomerCountResult`

NewResponseCustomerCountResultWithDefaults instantiates a new ResponseCustomerCountResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomersCount

`func (o *ResponseCustomerCountResult) GetCustomersCount() int32`

GetCustomersCount returns the CustomersCount field if non-nil, zero value otherwise.

### GetCustomersCountOk

`func (o *ResponseCustomerCountResult) GetCustomersCountOk() (*int32, bool)`

GetCustomersCountOk returns a tuple with the CustomersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomersCount

`func (o *ResponseCustomerCountResult) SetCustomersCount(v int32)`

SetCustomersCount sets CustomersCount field to given value.

### HasCustomersCount

`func (o *ResponseCustomerCountResult) HasCustomersCount() bool`

HasCustomersCount returns a boolean if a field has been set.

### SetCustomersCountNil

`func (o *ResponseCustomerCountResult) SetCustomersCountNil(b bool)`

 SetCustomersCountNil sets the value for CustomersCount to be an explicit nil

### UnsetCustomersCount
`func (o *ResponseCustomerCountResult) UnsetCustomersCount()`

UnsetCustomersCount ensures that no value is present for CustomersCount, not even an explicit nil
### GetAdditionalFields

`func (o *ResponseCustomerCountResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseCustomerCountResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseCustomerCountResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseCustomerCountResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseCustomerCountResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseCustomerCountResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseCustomerCountResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseCustomerCountResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseCustomerCountResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseCustomerCountResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseCustomerCountResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseCustomerCountResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


