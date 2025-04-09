# CartShippingZone2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Countries** | Pointer to [**[]Country**](Country.md) |  | [optional] 
**ShippingMethods** | Pointer to [**[]CartShippingMethod**](CartShippingMethod.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCartShippingZone2

`func NewCartShippingZone2() *CartShippingZone2`

NewCartShippingZone2 instantiates a new CartShippingZone2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartShippingZone2WithDefaults

`func NewCartShippingZone2WithDefaults() *CartShippingZone2`

NewCartShippingZone2WithDefaults instantiates a new CartShippingZone2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CartShippingZone2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartShippingZone2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartShippingZone2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CartShippingZone2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CartShippingZone2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartShippingZone2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartShippingZone2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CartShippingZone2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *CartShippingZone2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CartShippingZone2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CartShippingZone2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CartShippingZone2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCountries

`func (o *CartShippingZone2) GetCountries() []Country`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *CartShippingZone2) GetCountriesOk() (*[]Country, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *CartShippingZone2) SetCountries(v []Country)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *CartShippingZone2) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetShippingMethods

`func (o *CartShippingZone2) GetShippingMethods() []CartShippingMethod`

GetShippingMethods returns the ShippingMethods field if non-nil, zero value otherwise.

### GetShippingMethodsOk

`func (o *CartShippingZone2) GetShippingMethodsOk() (*[]CartShippingMethod, bool)`

GetShippingMethodsOk returns a tuple with the ShippingMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethods

`func (o *CartShippingZone2) SetShippingMethods(v []CartShippingMethod)`

SetShippingMethods sets ShippingMethods field to given value.

### HasShippingMethods

`func (o *CartShippingZone2) HasShippingMethods() bool`

HasShippingMethods returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *CartShippingZone2) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CartShippingZone2) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CartShippingZone2) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CartShippingZone2) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *CartShippingZone2) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CartShippingZone2) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CartShippingZone2) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CartShippingZone2) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


