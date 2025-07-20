# BasketLiveShippingService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Callback** | Pointer to **string** |  | [optional] 
**CallbackErrCnt** | Pointer to **NullableInt32** |  | [optional] 
**EnabledOnStore** | Pointer to **NullableBool** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBasketLiveShippingService

`func NewBasketLiveShippingService() *BasketLiveShippingService`

NewBasketLiveShippingService instantiates a new BasketLiveShippingService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasketLiveShippingServiceWithDefaults

`func NewBasketLiveShippingServiceWithDefaults() *BasketLiveShippingService`

NewBasketLiveShippingServiceWithDefaults instantiates a new BasketLiveShippingService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasketLiveShippingService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasketLiveShippingService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasketLiveShippingService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BasketLiveShippingService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BasketLiveShippingService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasketLiveShippingService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasketLiveShippingService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BasketLiveShippingService) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BasketLiveShippingService) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BasketLiveShippingService) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCallback

`func (o *BasketLiveShippingService) GetCallback() string`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *BasketLiveShippingService) GetCallbackOk() (*string, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *BasketLiveShippingService) SetCallback(v string)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *BasketLiveShippingService) HasCallback() bool`

HasCallback returns a boolean if a field has been set.

### GetCallbackErrCnt

`func (o *BasketLiveShippingService) GetCallbackErrCnt() int32`

GetCallbackErrCnt returns the CallbackErrCnt field if non-nil, zero value otherwise.

### GetCallbackErrCntOk

`func (o *BasketLiveShippingService) GetCallbackErrCntOk() (*int32, bool)`

GetCallbackErrCntOk returns a tuple with the CallbackErrCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackErrCnt

`func (o *BasketLiveShippingService) SetCallbackErrCnt(v int32)`

SetCallbackErrCnt sets CallbackErrCnt field to given value.

### HasCallbackErrCnt

`func (o *BasketLiveShippingService) HasCallbackErrCnt() bool`

HasCallbackErrCnt returns a boolean if a field has been set.

### SetCallbackErrCntNil

`func (o *BasketLiveShippingService) SetCallbackErrCntNil(b bool)`

 SetCallbackErrCntNil sets the value for CallbackErrCnt to be an explicit nil

### UnsetCallbackErrCnt
`func (o *BasketLiveShippingService) UnsetCallbackErrCnt()`

UnsetCallbackErrCnt ensures that no value is present for CallbackErrCnt, not even an explicit nil
### GetEnabledOnStore

`func (o *BasketLiveShippingService) GetEnabledOnStore() bool`

GetEnabledOnStore returns the EnabledOnStore field if non-nil, zero value otherwise.

### GetEnabledOnStoreOk

`func (o *BasketLiveShippingService) GetEnabledOnStoreOk() (*bool, bool)`

GetEnabledOnStoreOk returns a tuple with the EnabledOnStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnStore

`func (o *BasketLiveShippingService) SetEnabledOnStore(v bool)`

SetEnabledOnStore sets EnabledOnStore field to given value.

### HasEnabledOnStore

`func (o *BasketLiveShippingService) HasEnabledOnStore() bool`

HasEnabledOnStore returns a boolean if a field has been set.

### SetEnabledOnStoreNil

`func (o *BasketLiveShippingService) SetEnabledOnStoreNil(b bool)`

 SetEnabledOnStoreNil sets the value for EnabledOnStore to be an explicit nil

### UnsetEnabledOnStore
`func (o *BasketLiveShippingService) UnsetEnabledOnStore()`

UnsetEnabledOnStore ensures that no value is present for EnabledOnStore, not even an explicit nil
### GetAdditionalFields

`func (o *BasketLiveShippingService) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *BasketLiveShippingService) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *BasketLiveShippingService) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *BasketLiveShippingService) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *BasketLiveShippingService) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *BasketLiveShippingService) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *BasketLiveShippingService) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BasketLiveShippingService) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BasketLiveShippingService) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BasketLiveShippingService) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *BasketLiveShippingService) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *BasketLiveShippingService) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


