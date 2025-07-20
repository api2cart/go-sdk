# ProductAdvancedPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 
**Avail** | Pointer to **NullableBool** |  | [optional] 
**GroupId** | Pointer to **NullableString** |  | [optional] 
**QuantityFrom** | Pointer to **NullableFloat32** |  | [optional] 
**StartTime** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**ExpireTime** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProductAdvancedPrice

`func NewProductAdvancedPrice() *ProductAdvancedPrice`

NewProductAdvancedPrice instantiates a new ProductAdvancedPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAdvancedPriceWithDefaults

`func NewProductAdvancedPriceWithDefaults() *ProductAdvancedPrice`

NewProductAdvancedPriceWithDefaults instantiates a new ProductAdvancedPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductAdvancedPrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductAdvancedPrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductAdvancedPrice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductAdvancedPrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *ProductAdvancedPrice) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProductAdvancedPrice) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProductAdvancedPrice) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProductAdvancedPrice) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAvail

`func (o *ProductAdvancedPrice) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *ProductAdvancedPrice) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *ProductAdvancedPrice) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *ProductAdvancedPrice) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### SetAvailNil

`func (o *ProductAdvancedPrice) SetAvailNil(b bool)`

 SetAvailNil sets the value for Avail to be an explicit nil

### UnsetAvail
`func (o *ProductAdvancedPrice) UnsetAvail()`

UnsetAvail ensures that no value is present for Avail, not even an explicit nil
### GetGroupId

`func (o *ProductAdvancedPrice) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ProductAdvancedPrice) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ProductAdvancedPrice) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ProductAdvancedPrice) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *ProductAdvancedPrice) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *ProductAdvancedPrice) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetQuantityFrom

`func (o *ProductAdvancedPrice) GetQuantityFrom() float32`

GetQuantityFrom returns the QuantityFrom field if non-nil, zero value otherwise.

### GetQuantityFromOk

`func (o *ProductAdvancedPrice) GetQuantityFromOk() (*float32, bool)`

GetQuantityFromOk returns a tuple with the QuantityFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityFrom

`func (o *ProductAdvancedPrice) SetQuantityFrom(v float32)`

SetQuantityFrom sets QuantityFrom field to given value.

### HasQuantityFrom

`func (o *ProductAdvancedPrice) HasQuantityFrom() bool`

HasQuantityFrom returns a boolean if a field has been set.

### SetQuantityFromNil

`func (o *ProductAdvancedPrice) SetQuantityFromNil(b bool)`

 SetQuantityFromNil sets the value for QuantityFrom to be an explicit nil

### UnsetQuantityFrom
`func (o *ProductAdvancedPrice) UnsetQuantityFrom()`

UnsetQuantityFrom ensures that no value is present for QuantityFrom, not even an explicit nil
### GetStartTime

`func (o *ProductAdvancedPrice) GetStartTime() A2CDateTime`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ProductAdvancedPrice) GetStartTimeOk() (*A2CDateTime, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ProductAdvancedPrice) SetStartTime(v A2CDateTime)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ProductAdvancedPrice) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ProductAdvancedPrice) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ProductAdvancedPrice) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetExpireTime

`func (o *ProductAdvancedPrice) GetExpireTime() A2CDateTime`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *ProductAdvancedPrice) GetExpireTimeOk() (*A2CDateTime, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *ProductAdvancedPrice) SetExpireTime(v A2CDateTime)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *ProductAdvancedPrice) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### SetExpireTimeNil

`func (o *ProductAdvancedPrice) SetExpireTimeNil(b bool)`

 SetExpireTimeNil sets the value for ExpireTime to be an explicit nil

### UnsetExpireTime
`func (o *ProductAdvancedPrice) UnsetExpireTime()`

UnsetExpireTime ensures that no value is present for ExpireTime, not even an explicit nil
### GetAdditionalFields

`func (o *ProductAdvancedPrice) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ProductAdvancedPrice) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ProductAdvancedPrice) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ProductAdvancedPrice) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *ProductAdvancedPrice) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ProductAdvancedPrice) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *ProductAdvancedPrice) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProductAdvancedPrice) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProductAdvancedPrice) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProductAdvancedPrice) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *ProductAdvancedPrice) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ProductAdvancedPrice) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


