# ResponseCustomerAttributeListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]CustomerAttribute**](CustomerAttribute.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseCustomerAttributeListResult

`func NewResponseCustomerAttributeListResult() *ResponseCustomerAttributeListResult`

NewResponseCustomerAttributeListResult instantiates a new ResponseCustomerAttributeListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCustomerAttributeListResultWithDefaults

`func NewResponseCustomerAttributeListResultWithDefaults() *ResponseCustomerAttributeListResult`

NewResponseCustomerAttributeListResultWithDefaults instantiates a new ResponseCustomerAttributeListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ResponseCustomerAttributeListResult) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResponseCustomerAttributeListResult) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResponseCustomerAttributeListResult) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ResponseCustomerAttributeListResult) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *ResponseCustomerAttributeListResult) GetItems() []CustomerAttribute`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResponseCustomerAttributeListResult) GetItemsOk() (*[]CustomerAttribute, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResponseCustomerAttributeListResult) SetItems(v []CustomerAttribute)`

SetItems sets Items field to given value.

### HasItems

`func (o *ResponseCustomerAttributeListResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseCustomerAttributeListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseCustomerAttributeListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseCustomerAttributeListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseCustomerAttributeListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseCustomerAttributeListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseCustomerAttributeListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseCustomerAttributeListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseCustomerAttributeListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


