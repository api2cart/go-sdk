# ResponseProductOptionListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Option** | Pointer to [**[]ProductOption**](ProductOption.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseProductOptionListResult

`func NewResponseProductOptionListResult() *ResponseProductOptionListResult`

NewResponseProductOptionListResult instantiates a new ResponseProductOptionListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseProductOptionListResultWithDefaults

`func NewResponseProductOptionListResultWithDefaults() *ResponseProductOptionListResult`

NewResponseProductOptionListResultWithDefaults instantiates a new ResponseProductOptionListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOption

`func (o *ResponseProductOptionListResult) GetOption() []ProductOption`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *ResponseProductOptionListResult) GetOptionOk() (*[]ProductOption, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *ResponseProductOptionListResult) SetOption(v []ProductOption)`

SetOption sets Option field to given value.

### HasOption

`func (o *ResponseProductOptionListResult) HasOption() bool`

HasOption returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseProductOptionListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseProductOptionListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseProductOptionListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseProductOptionListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseProductOptionListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseProductOptionListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseProductOptionListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseProductOptionListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


