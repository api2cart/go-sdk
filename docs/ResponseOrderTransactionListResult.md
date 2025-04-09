# ResponseOrderTransactionListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionsCount** | Pointer to **int32** |  | [optional] 
**Transactions** | Pointer to [**[]OrderTransaction**](OrderTransaction.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseOrderTransactionListResult

`func NewResponseOrderTransactionListResult() *ResponseOrderTransactionListResult`

NewResponseOrderTransactionListResult instantiates a new ResponseOrderTransactionListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseOrderTransactionListResultWithDefaults

`func NewResponseOrderTransactionListResultWithDefaults() *ResponseOrderTransactionListResult`

NewResponseOrderTransactionListResultWithDefaults instantiates a new ResponseOrderTransactionListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionsCount

`func (o *ResponseOrderTransactionListResult) GetTransactionsCount() int32`

GetTransactionsCount returns the TransactionsCount field if non-nil, zero value otherwise.

### GetTransactionsCountOk

`func (o *ResponseOrderTransactionListResult) GetTransactionsCountOk() (*int32, bool)`

GetTransactionsCountOk returns a tuple with the TransactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsCount

`func (o *ResponseOrderTransactionListResult) SetTransactionsCount(v int32)`

SetTransactionsCount sets TransactionsCount field to given value.

### HasTransactionsCount

`func (o *ResponseOrderTransactionListResult) HasTransactionsCount() bool`

HasTransactionsCount returns a boolean if a field has been set.

### GetTransactions

`func (o *ResponseOrderTransactionListResult) GetTransactions() []OrderTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ResponseOrderTransactionListResult) GetTransactionsOk() (*[]OrderTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ResponseOrderTransactionListResult) SetTransactions(v []OrderTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *ResponseOrderTransactionListResult) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseOrderTransactionListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseOrderTransactionListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseOrderTransactionListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseOrderTransactionListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseOrderTransactionListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseOrderTransactionListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseOrderTransactionListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseOrderTransactionListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


