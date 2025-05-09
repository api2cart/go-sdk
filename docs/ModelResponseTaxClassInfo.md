# ModelResponseTaxClassInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | Pointer to **int32** |  | [optional] 
**ReturnMessage** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**ResponseTaxClassInfoResult**](ResponseTaxClassInfoResult.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModelResponseTaxClassInfo

`func NewModelResponseTaxClassInfo() *ModelResponseTaxClassInfo`

NewModelResponseTaxClassInfo instantiates a new ModelResponseTaxClassInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelResponseTaxClassInfoWithDefaults

`func NewModelResponseTaxClassInfoWithDefaults() *ModelResponseTaxClassInfo`

NewModelResponseTaxClassInfoWithDefaults instantiates a new ModelResponseTaxClassInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *ModelResponseTaxClassInfo) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ModelResponseTaxClassInfo) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ModelResponseTaxClassInfo) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ModelResponseTaxClassInfo) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnMessage

`func (o *ModelResponseTaxClassInfo) GetReturnMessage() string`

GetReturnMessage returns the ReturnMessage field if non-nil, zero value otherwise.

### GetReturnMessageOk

`func (o *ModelResponseTaxClassInfo) GetReturnMessageOk() (*string, bool)`

GetReturnMessageOk returns a tuple with the ReturnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnMessage

`func (o *ModelResponseTaxClassInfo) SetReturnMessage(v string)`

SetReturnMessage sets ReturnMessage field to given value.

### HasReturnMessage

`func (o *ModelResponseTaxClassInfo) HasReturnMessage() bool`

HasReturnMessage returns a boolean if a field has been set.

### GetResult

`func (o *ModelResponseTaxClassInfo) GetResult() ResponseTaxClassInfoResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModelResponseTaxClassInfo) GetResultOk() (*ResponseTaxClassInfoResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModelResponseTaxClassInfo) SetResult(v ResponseTaxClassInfoResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ModelResponseTaxClassInfo) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ModelResponseTaxClassInfo) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ModelResponseTaxClassInfo) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ModelResponseTaxClassInfo) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ModelResponseTaxClassInfo) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ModelResponseTaxClassInfo) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModelResponseTaxClassInfo) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModelResponseTaxClassInfo) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModelResponseTaxClassInfo) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


