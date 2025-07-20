# TaxClassStates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Tax** | Pointer to **float32** |  | [optional] 
**TaxType** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ZipCodes** | Pointer to [**[]TaxClassZipCodes**](TaxClassZipCodes.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTaxClassStates

`func NewTaxClassStates() *TaxClassStates`

NewTaxClassStates instantiates a new TaxClassStates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxClassStatesWithDefaults

`func NewTaxClassStatesWithDefaults() *TaxClassStates`

NewTaxClassStatesWithDefaults instantiates a new TaxClassStates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxClassStates) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxClassStates) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxClassStates) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxClassStates) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaxClassStates) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaxClassStates) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTax

`func (o *TaxClassStates) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *TaxClassStates) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *TaxClassStates) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *TaxClassStates) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxType

`func (o *TaxClassStates) GetTaxType() int32`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *TaxClassStates) GetTaxTypeOk() (*int32, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *TaxClassStates) SetTaxType(v int32)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *TaxClassStates) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetName

`func (o *TaxClassStates) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxClassStates) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxClassStates) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxClassStates) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *TaxClassStates) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaxClassStates) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaxClassStates) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaxClassStates) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetZipCodes

`func (o *TaxClassStates) GetZipCodes() []TaxClassZipCodes`

GetZipCodes returns the ZipCodes field if non-nil, zero value otherwise.

### GetZipCodesOk

`func (o *TaxClassStates) GetZipCodesOk() (*[]TaxClassZipCodes, bool)`

GetZipCodesOk returns a tuple with the ZipCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodes

`func (o *TaxClassStates) SetZipCodes(v []TaxClassZipCodes)`

SetZipCodes sets ZipCodes field to given value.

### HasZipCodes

`func (o *TaxClassStates) HasZipCodes() bool`

HasZipCodes returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *TaxClassStates) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *TaxClassStates) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *TaxClassStates) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *TaxClassStates) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *TaxClassStates) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *TaxClassStates) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *TaxClassStates) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TaxClassStates) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TaxClassStates) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TaxClassStates) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *TaxClassStates) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *TaxClassStates) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


