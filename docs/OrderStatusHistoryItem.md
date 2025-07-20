# OrderStatusHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ModifiedTime** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**Notify** | Pointer to **NullableBool** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderStatusHistoryItem

`func NewOrderStatusHistoryItem() *OrderStatusHistoryItem`

NewOrderStatusHistoryItem instantiates a new OrderStatusHistoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusHistoryItemWithDefaults

`func NewOrderStatusHistoryItemWithDefaults() *OrderStatusHistoryItem`

NewOrderStatusHistoryItemWithDefaults instantiates a new OrderStatusHistoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderStatusHistoryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderStatusHistoryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderStatusHistoryItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderStatusHistoryItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrderStatusHistoryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderStatusHistoryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderStatusHistoryItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderStatusHistoryItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModifiedTime

`func (o *OrderStatusHistoryItem) GetModifiedTime() A2CDateTime`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *OrderStatusHistoryItem) GetModifiedTimeOk() (*A2CDateTime, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *OrderStatusHistoryItem) SetModifiedTime(v A2CDateTime)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *OrderStatusHistoryItem) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetNotify

`func (o *OrderStatusHistoryItem) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *OrderStatusHistoryItem) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *OrderStatusHistoryItem) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *OrderStatusHistoryItem) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### SetNotifyNil

`func (o *OrderStatusHistoryItem) SetNotifyNil(b bool)`

 SetNotifyNil sets the value for Notify to be an explicit nil

### UnsetNotify
`func (o *OrderStatusHistoryItem) UnsetNotify()`

UnsetNotify ensures that no value is present for Notify, not even an explicit nil
### GetComment

`func (o *OrderStatusHistoryItem) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *OrderStatusHistoryItem) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *OrderStatusHistoryItem) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *OrderStatusHistoryItem) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *OrderStatusHistoryItem) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *OrderStatusHistoryItem) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetAdditionalFields

`func (o *OrderStatusHistoryItem) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderStatusHistoryItem) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderStatusHistoryItem) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderStatusHistoryItem) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *OrderStatusHistoryItem) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *OrderStatusHistoryItem) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *OrderStatusHistoryItem) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderStatusHistoryItem) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderStatusHistoryItem) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderStatusHistoryItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *OrderStatusHistoryItem) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *OrderStatusHistoryItem) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


