# TaxClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Avail** | Pointer to **bool** |  | [optional] 
**Tax** | Pointer to **float32** |  | [optional] 
**TaxType** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTaxClass

`func NewTaxClass() *TaxClass`

NewTaxClass instantiates a new TaxClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxClassWithDefaults

`func NewTaxClassWithDefaults() *TaxClass`

NewTaxClassWithDefaults instantiates a new TaxClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxClass) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxClass) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaxClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxClass) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxClass) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvail

`func (o *TaxClass) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *TaxClass) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *TaxClass) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *TaxClass) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetTax

`func (o *TaxClass) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *TaxClass) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *TaxClass) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *TaxClass) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxType

`func (o *TaxClass) GetTaxType() int32`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *TaxClass) GetTaxTypeOk() (*int32, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *TaxClass) SetTaxType(v int32)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *TaxClass) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaxClass) GetCreatedAt() A2CDateTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaxClass) GetCreatedAtOk() (*A2CDateTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaxClass) SetCreatedAt(v A2CDateTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaxClass) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *TaxClass) GetModifiedAt() A2CDateTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *TaxClass) GetModifiedAtOk() (*A2CDateTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *TaxClass) SetModifiedAt(v A2CDateTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *TaxClass) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *TaxClass) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *TaxClass) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *TaxClass) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *TaxClass) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *TaxClass) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TaxClass) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TaxClass) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TaxClass) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


