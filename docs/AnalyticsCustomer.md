# AnalyticsCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**CustomerType** | Pointer to **NullableString** |  | [optional] 
**Metrics** | Pointer to [**AnalyticsCustomerMetric**](AnalyticsCustomerMetric.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAnalyticsCustomer

`func NewAnalyticsCustomer() *AnalyticsCustomer`

NewAnalyticsCustomer instantiates a new AnalyticsCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsCustomerWithDefaults

`func NewAnalyticsCustomerWithDefaults() *AnalyticsCustomer`

NewAnalyticsCustomerWithDefaults instantiates a new AnalyticsCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *AnalyticsCustomer) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AnalyticsCustomer) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AnalyticsCustomer) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AnalyticsCustomer) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetName

`func (o *AnalyticsCustomer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnalyticsCustomer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnalyticsCustomer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnalyticsCustomer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AnalyticsCustomer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AnalyticsCustomer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *AnalyticsCustomer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AnalyticsCustomer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AnalyticsCustomer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AnalyticsCustomer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AnalyticsCustomer) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AnalyticsCustomer) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCustomerType

`func (o *AnalyticsCustomer) GetCustomerType() string`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AnalyticsCustomer) GetCustomerTypeOk() (*string, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AnalyticsCustomer) SetCustomerType(v string)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *AnalyticsCustomer) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### SetCustomerTypeNil

`func (o *AnalyticsCustomer) SetCustomerTypeNil(b bool)`

 SetCustomerTypeNil sets the value for CustomerType to be an explicit nil

### UnsetCustomerType
`func (o *AnalyticsCustomer) UnsetCustomerType()`

UnsetCustomerType ensures that no value is present for CustomerType, not even an explicit nil
### GetMetrics

`func (o *AnalyticsCustomer) GetMetrics() AnalyticsCustomerMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AnalyticsCustomer) GetMetricsOk() (*AnalyticsCustomerMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AnalyticsCustomer) SetMetrics(v AnalyticsCustomerMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AnalyticsCustomer) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *AnalyticsCustomer) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *AnalyticsCustomer) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *AnalyticsCustomer) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *AnalyticsCustomer) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *AnalyticsCustomer) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *AnalyticsCustomer) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *AnalyticsCustomer) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AnalyticsCustomer) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AnalyticsCustomer) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AnalyticsCustomer) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *AnalyticsCustomer) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *AnalyticsCustomer) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


