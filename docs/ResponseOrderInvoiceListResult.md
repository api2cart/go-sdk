# ResponseOrderInvoiceListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceCount** | Pointer to **NullableInt32** |  | [optional] 
**Invoice** | Pointer to [**[]OrderInvoice**](OrderInvoice.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseOrderInvoiceListResult

`func NewResponseOrderInvoiceListResult() *ResponseOrderInvoiceListResult`

NewResponseOrderInvoiceListResult instantiates a new ResponseOrderInvoiceListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrderInvoiceListResultWithDefaults

`func NewResponseOrderInvoiceListResultWithDefaults() *ResponseOrderInvoiceListResult`

NewResponseOrderInvoiceListResultWithDefaults instantiates a new ResponseOrderInvoiceListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceCount

`func (o *ResponseOrderInvoiceListResult) GetInvoiceCount() int32`

GetInvoiceCount returns the InvoiceCount field if non-nil, zero value otherwise.

### GetInvoiceCountOk

`func (o *ResponseOrderInvoiceListResult) GetInvoiceCountOk() (*int32, bool)`

GetInvoiceCountOk returns a tuple with the InvoiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCount

`func (o *ResponseOrderInvoiceListResult) SetInvoiceCount(v int32)`

SetInvoiceCount sets InvoiceCount field to given value.

### HasInvoiceCount

`func (o *ResponseOrderInvoiceListResult) HasInvoiceCount() bool`

HasInvoiceCount returns a boolean if a field has been set.

### SetInvoiceCountNil

`func (o *ResponseOrderInvoiceListResult) SetInvoiceCountNil(b bool)`

 SetInvoiceCountNil sets the value for InvoiceCount to be an explicit nil

### UnsetInvoiceCount
`func (o *ResponseOrderInvoiceListResult) UnsetInvoiceCount()`

UnsetInvoiceCount ensures that no value is present for InvoiceCount, not even an explicit nil
### GetInvoice

`func (o *ResponseOrderInvoiceListResult) GetInvoice() []OrderInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *ResponseOrderInvoiceListResult) GetInvoiceOk() (*[]OrderInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *ResponseOrderInvoiceListResult) SetInvoice(v []OrderInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *ResponseOrderInvoiceListResult) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseOrderInvoiceListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseOrderInvoiceListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseOrderInvoiceListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseOrderInvoiceListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ResponseOrderInvoiceListResult) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ResponseOrderInvoiceListResult) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ResponseOrderInvoiceListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseOrderInvoiceListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseOrderInvoiceListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseOrderInvoiceListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ResponseOrderInvoiceListResult) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ResponseOrderInvoiceListResult) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


