# CartWarehouse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Avail** | Pointer to **NullableBool** |  | [optional] 
**Address** | Pointer to [**NullableCustomerAddress**](CustomerAddress.md) |  | [optional] 
**CarriersIds** | Pointer to **[]string** |  | [optional] 
**StoresIds** | Pointer to **[]string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCartWarehouse

`func NewCartWarehouse() *CartWarehouse`

NewCartWarehouse instantiates a new CartWarehouse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartWarehouseWithDefaults

`func NewCartWarehouseWithDefaults() *CartWarehouse`

NewCartWarehouseWithDefaults instantiates a new CartWarehouse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CartWarehouse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartWarehouse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartWarehouse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CartWarehouse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CartWarehouse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CartWarehouse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CartWarehouse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartWarehouse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartWarehouse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CartWarehouse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CartWarehouse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CartWarehouse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CartWarehouse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CartWarehouse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CartWarehouse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CartWarehouse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CartWarehouse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CartWarehouse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAvail

`func (o *CartWarehouse) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *CartWarehouse) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *CartWarehouse) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *CartWarehouse) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### SetAvailNil

`func (o *CartWarehouse) SetAvailNil(b bool)`

 SetAvailNil sets the value for Avail to be an explicit nil

### UnsetAvail
`func (o *CartWarehouse) UnsetAvail()`

UnsetAvail ensures that no value is present for Avail, not even an explicit nil
### GetAddress

`func (o *CartWarehouse) GetAddress() CustomerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CartWarehouse) GetAddressOk() (*CustomerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CartWarehouse) SetAddress(v CustomerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CartWarehouse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *CartWarehouse) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *CartWarehouse) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCarriersIds

`func (o *CartWarehouse) GetCarriersIds() []string`

GetCarriersIds returns the CarriersIds field if non-nil, zero value otherwise.

### GetCarriersIdsOk

`func (o *CartWarehouse) GetCarriersIdsOk() (*[]string, bool)`

GetCarriersIdsOk returns a tuple with the CarriersIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarriersIds

`func (o *CartWarehouse) SetCarriersIds(v []string)`

SetCarriersIds sets CarriersIds field to given value.

### HasCarriersIds

`func (o *CartWarehouse) HasCarriersIds() bool`

HasCarriersIds returns a boolean if a field has been set.

### GetStoresIds

`func (o *CartWarehouse) GetStoresIds() []string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *CartWarehouse) GetStoresIdsOk() (*[]string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *CartWarehouse) SetStoresIds(v []string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *CartWarehouse) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *CartWarehouse) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CartWarehouse) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CartWarehouse) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CartWarehouse) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *CartWarehouse) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *CartWarehouse) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *CartWarehouse) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CartWarehouse) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CartWarehouse) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CartWarehouse) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CartWarehouse) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CartWarehouse) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


