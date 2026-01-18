# CartPickupLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Avail** | Pointer to **NullableBool** |  | [optional] 
**Address** | Pointer to [**NullableCustomerAddress**](CustomerAddress.md) |  | [optional] 
**PickupInstructions** | Pointer to **NullableString** |  | [optional] 
**ExpectedReadyTime** | Pointer to **NullableString** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCartPickupLocation

`func NewCartPickupLocation() *CartPickupLocation`

NewCartPickupLocation instantiates a new CartPickupLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartPickupLocationWithDefaults

`func NewCartPickupLocationWithDefaults() *CartPickupLocation`

NewCartPickupLocationWithDefaults instantiates a new CartPickupLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CartPickupLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartPickupLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartPickupLocation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CartPickupLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CartPickupLocation) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CartPickupLocation) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CartPickupLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartPickupLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartPickupLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CartPickupLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CartPickupLocation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CartPickupLocation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAvail

`func (o *CartPickupLocation) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *CartPickupLocation) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *CartPickupLocation) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *CartPickupLocation) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### SetAvailNil

`func (o *CartPickupLocation) SetAvailNil(b bool)`

 SetAvailNil sets the value for Avail to be an explicit nil

### UnsetAvail
`func (o *CartPickupLocation) UnsetAvail()`

UnsetAvail ensures that no value is present for Avail, not even an explicit nil
### GetAddress

`func (o *CartPickupLocation) GetAddress() CustomerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CartPickupLocation) GetAddressOk() (*CustomerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CartPickupLocation) SetAddress(v CustomerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CartPickupLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *CartPickupLocation) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *CartPickupLocation) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPickupInstructions

`func (o *CartPickupLocation) GetPickupInstructions() string`

GetPickupInstructions returns the PickupInstructions field if non-nil, zero value otherwise.

### GetPickupInstructionsOk

`func (o *CartPickupLocation) GetPickupInstructionsOk() (*string, bool)`

GetPickupInstructionsOk returns a tuple with the PickupInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupInstructions

`func (o *CartPickupLocation) SetPickupInstructions(v string)`

SetPickupInstructions sets PickupInstructions field to given value.

### HasPickupInstructions

`func (o *CartPickupLocation) HasPickupInstructions() bool`

HasPickupInstructions returns a boolean if a field has been set.

### SetPickupInstructionsNil

`func (o *CartPickupLocation) SetPickupInstructionsNil(b bool)`

 SetPickupInstructionsNil sets the value for PickupInstructions to be an explicit nil

### UnsetPickupInstructions
`func (o *CartPickupLocation) UnsetPickupInstructions()`

UnsetPickupInstructions ensures that no value is present for PickupInstructions, not even an explicit nil
### GetExpectedReadyTime

`func (o *CartPickupLocation) GetExpectedReadyTime() string`

GetExpectedReadyTime returns the ExpectedReadyTime field if non-nil, zero value otherwise.

### GetExpectedReadyTimeOk

`func (o *CartPickupLocation) GetExpectedReadyTimeOk() (*string, bool)`

GetExpectedReadyTimeOk returns a tuple with the ExpectedReadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedReadyTime

`func (o *CartPickupLocation) SetExpectedReadyTime(v string)`

SetExpectedReadyTime sets ExpectedReadyTime field to given value.

### HasExpectedReadyTime

`func (o *CartPickupLocation) HasExpectedReadyTime() bool`

HasExpectedReadyTime returns a boolean if a field has been set.

### SetExpectedReadyTimeNil

`func (o *CartPickupLocation) SetExpectedReadyTimeNil(b bool)`

 SetExpectedReadyTimeNil sets the value for ExpectedReadyTime to be an explicit nil

### UnsetExpectedReadyTime
`func (o *CartPickupLocation) UnsetExpectedReadyTime()`

UnsetExpectedReadyTime ensures that no value is present for ExpectedReadyTime, not even an explicit nil
### GetAdditionalFields

`func (o *CartPickupLocation) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CartPickupLocation) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CartPickupLocation) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CartPickupLocation) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *CartPickupLocation) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *CartPickupLocation) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *CartPickupLocation) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CartPickupLocation) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CartPickupLocation) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CartPickupLocation) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *CartPickupLocation) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *CartPickupLocation) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


