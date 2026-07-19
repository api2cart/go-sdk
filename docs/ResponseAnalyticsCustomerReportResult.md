# ResponseAnalyticsCustomerReportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to [**AnalyticsPeriod**](AnalyticsPeriod.md) |  | [optional] 
**Items** | Pointer to [**[]AnalyticsCustomer**](AnalyticsCustomer.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseAnalyticsCustomerReportResult

`func NewResponseAnalyticsCustomerReportResult() *ResponseAnalyticsCustomerReportResult`

NewResponseAnalyticsCustomerReportResult instantiates a new ResponseAnalyticsCustomerReportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAnalyticsCustomerReportResultWithDefaults

`func NewResponseAnalyticsCustomerReportResultWithDefaults() *ResponseAnalyticsCustomerReportResult`

NewResponseAnalyticsCustomerReportResultWithDefaults instantiates a new ResponseAnalyticsCustomerReportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *ResponseAnalyticsCustomerReportResult) GetPeriod() AnalyticsPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ResponseAnalyticsCustomerReportResult) GetPeriodOk() (*AnalyticsPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ResponseAnalyticsCustomerReportResult) SetPeriod(v AnalyticsPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ResponseAnalyticsCustomerReportResult) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetItems

`func (o *ResponseAnalyticsCustomerReportResult) GetItems() []AnalyticsCustomer`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResponseAnalyticsCustomerReportResult) GetItemsOk() (*[]AnalyticsCustomer, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResponseAnalyticsCustomerReportResult) SetItems(v []AnalyticsCustomer)`

SetItems sets Items field to given value.

### HasItems

`func (o *ResponseAnalyticsCustomerReportResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseAnalyticsCustomerReportResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseAnalyticsCustomerReportResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseAnalyticsCustomerReportResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseAnalyticsCustomerReportResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseAnalyticsCustomerReportResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseAnalyticsCustomerReportResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseAnalyticsCustomerReportResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseAnalyticsCustomerReportResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseAnalyticsCustomerReportResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseAnalyticsCustomerReportResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseAnalyticsCustomerReportResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseAnalyticsCustomerReportResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


