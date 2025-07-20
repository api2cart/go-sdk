# TaxClassZipCodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRange** | Pointer to **bool** |  | [optional] 
**Range** | Pointer to **[]string** |  | [optional] 
**Fields** | Pointer to [**[]TaxClassZipCodesRange**](TaxClassZipCodesRange.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTaxClassZipCodes

`func NewTaxClassZipCodes() *TaxClassZipCodes`

NewTaxClassZipCodes instantiates a new TaxClassZipCodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxClassZipCodesWithDefaults

`func NewTaxClassZipCodesWithDefaults() *TaxClassZipCodes`

NewTaxClassZipCodesWithDefaults instantiates a new TaxClassZipCodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRange

`func (o *TaxClassZipCodes) GetIsRange() bool`

GetIsRange returns the IsRange field if non-nil, zero value otherwise.

### GetIsRangeOk

`func (o *TaxClassZipCodes) GetIsRangeOk() (*bool, bool)`

GetIsRangeOk returns a tuple with the IsRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRange

`func (o *TaxClassZipCodes) SetIsRange(v bool)`

SetIsRange sets IsRange field to given value.

### HasIsRange

`func (o *TaxClassZipCodes) HasIsRange() bool`

HasIsRange returns a boolean if a field has been set.

### GetRange

`func (o *TaxClassZipCodes) GetRange() []string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *TaxClassZipCodes) GetRangeOk() (*[]string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *TaxClassZipCodes) SetRange(v []string)`

SetRange sets Range field to given value.

### HasRange

`func (o *TaxClassZipCodes) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetFields

`func (o *TaxClassZipCodes) GetFields() []TaxClassZipCodesRange`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TaxClassZipCodes) GetFieldsOk() (*[]TaxClassZipCodesRange, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TaxClassZipCodes) SetFields(v []TaxClassZipCodesRange)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TaxClassZipCodes) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *TaxClassZipCodes) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *TaxClassZipCodes) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *TaxClassZipCodes) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *TaxClassZipCodes) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *TaxClassZipCodes) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *TaxClassZipCodes) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *TaxClassZipCodes) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TaxClassZipCodes) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TaxClassZipCodes) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TaxClassZipCodes) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *TaxClassZipCodes) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *TaxClassZipCodes) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


