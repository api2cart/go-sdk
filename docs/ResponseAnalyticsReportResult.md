# ResponseAnalyticsReportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to [**AnalyticsPeriod**](AnalyticsPeriod.md) |  | [optional] 
**TotalMetrics** | Pointer to [**AnalyticsMetric**](AnalyticsMetric.md) |  | [optional] 
**Intervals** | Pointer to [**[]AnalyticsInterval**](AnalyticsInterval.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseAnalyticsReportResult

`func NewResponseAnalyticsReportResult() *ResponseAnalyticsReportResult`

NewResponseAnalyticsReportResult instantiates a new ResponseAnalyticsReportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAnalyticsReportResultWithDefaults

`func NewResponseAnalyticsReportResultWithDefaults() *ResponseAnalyticsReportResult`

NewResponseAnalyticsReportResultWithDefaults instantiates a new ResponseAnalyticsReportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *ResponseAnalyticsReportResult) GetPeriod() AnalyticsPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ResponseAnalyticsReportResult) GetPeriodOk() (*AnalyticsPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ResponseAnalyticsReportResult) SetPeriod(v AnalyticsPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ResponseAnalyticsReportResult) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetTotalMetrics

`func (o *ResponseAnalyticsReportResult) GetTotalMetrics() AnalyticsMetric`

GetTotalMetrics returns the TotalMetrics field if non-nil, zero value otherwise.

### GetTotalMetricsOk

`func (o *ResponseAnalyticsReportResult) GetTotalMetricsOk() (*AnalyticsMetric, bool)`

GetTotalMetricsOk returns a tuple with the TotalMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMetrics

`func (o *ResponseAnalyticsReportResult) SetTotalMetrics(v AnalyticsMetric)`

SetTotalMetrics sets TotalMetrics field to given value.

### HasTotalMetrics

`func (o *ResponseAnalyticsReportResult) HasTotalMetrics() bool`

HasTotalMetrics returns a boolean if a field has been set.

### GetIntervals

`func (o *ResponseAnalyticsReportResult) GetIntervals() []AnalyticsInterval`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *ResponseAnalyticsReportResult) GetIntervalsOk() (*[]AnalyticsInterval, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *ResponseAnalyticsReportResult) SetIntervals(v []AnalyticsInterval)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *ResponseAnalyticsReportResult) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseAnalyticsReportResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseAnalyticsReportResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseAnalyticsReportResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseAnalyticsReportResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseAnalyticsReportResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseAnalyticsReportResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseAnalyticsReportResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseAnalyticsReportResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseAnalyticsReportResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseAnalyticsReportResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseAnalyticsReportResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseAnalyticsReportResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


