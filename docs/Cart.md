# Cart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**DbPrefix** | Pointer to **NullableString** |  | [optional] 
**StoresInfo** | Pointer to [**[]CartStoreInfo**](CartStoreInfo.md) |  | [optional] 
**Warehouses** | Pointer to [**[]CartWarehouse**](CartWarehouse.md) |  | [optional] 
**ShippingZones** | Pointer to [**[]CartShippingZone**](CartShippingZone.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCart

`func NewCart() *Cart`

NewCart instantiates a new Cart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartWithDefaults

`func NewCartWithDefaults() *Cart`

NewCartWithDefaults instantiates a new Cart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cart) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cart) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cart) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cart) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Cart) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Cart) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *Cart) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Cart) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Cart) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Cart) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *Cart) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Cart) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVersion

`func (o *Cart) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cart) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cart) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Cart) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Cart) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Cart) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetDbPrefix

`func (o *Cart) GetDbPrefix() string`

GetDbPrefix returns the DbPrefix field if non-nil, zero value otherwise.

### GetDbPrefixOk

`func (o *Cart) GetDbPrefixOk() (*string, bool)`

GetDbPrefixOk returns a tuple with the DbPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPrefix

`func (o *Cart) SetDbPrefix(v string)`

SetDbPrefix sets DbPrefix field to given value.

### HasDbPrefix

`func (o *Cart) HasDbPrefix() bool`

HasDbPrefix returns a boolean if a field has been set.

### SetDbPrefixNil

`func (o *Cart) SetDbPrefixNil(b bool)`

 SetDbPrefixNil sets the value for DbPrefix to be an explicit nil

### UnsetDbPrefix
`func (o *Cart) UnsetDbPrefix()`

UnsetDbPrefix ensures that no value is present for DbPrefix, not even an explicit nil
### GetStoresInfo

`func (o *Cart) GetStoresInfo() []CartStoreInfo`

GetStoresInfo returns the StoresInfo field if non-nil, zero value otherwise.

### GetStoresInfoOk

`func (o *Cart) GetStoresInfoOk() (*[]CartStoreInfo, bool)`

GetStoresInfoOk returns a tuple with the StoresInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresInfo

`func (o *Cart) SetStoresInfo(v []CartStoreInfo)`

SetStoresInfo sets StoresInfo field to given value.

### HasStoresInfo

`func (o *Cart) HasStoresInfo() bool`

HasStoresInfo returns a boolean if a field has been set.

### GetWarehouses

`func (o *Cart) GetWarehouses() []CartWarehouse`

GetWarehouses returns the Warehouses field if non-nil, zero value otherwise.

### GetWarehousesOk

`func (o *Cart) GetWarehousesOk() (*[]CartWarehouse, bool)`

GetWarehousesOk returns a tuple with the Warehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouses

`func (o *Cart) SetWarehouses(v []CartWarehouse)`

SetWarehouses sets Warehouses field to given value.

### HasWarehouses

`func (o *Cart) HasWarehouses() bool`

HasWarehouses returns a boolean if a field has been set.

### GetShippingZones

`func (o *Cart) GetShippingZones() []CartShippingZone`

GetShippingZones returns the ShippingZones field if non-nil, zero value otherwise.

### GetShippingZonesOk

`func (o *Cart) GetShippingZonesOk() (*[]CartShippingZone, bool)`

GetShippingZonesOk returns a tuple with the ShippingZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingZones

`func (o *Cart) SetShippingZones(v []CartShippingZone)`

SetShippingZones sets ShippingZones field to given value.

### HasShippingZones

`func (o *Cart) HasShippingZones() bool`

HasShippingZones returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Cart) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Cart) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Cart) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Cart) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *Cart) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *Cart) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *Cart) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Cart) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Cart) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Cart) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *Cart) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *Cart) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


