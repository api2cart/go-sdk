# OrderRefund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Shipping** | Pointer to **NullableFloat32** |  | [optional] 
**Fee** | Pointer to **NullableFloat32** |  | [optional] 
**Tax** | Pointer to **NullableFloat32** |  | [optional] 
**Total** | Pointer to **NullableFloat32** |  | [optional] 
**ModifiedTime** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]OrderStatusRefundItem**](OrderStatusRefundItem.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderRefund

`func NewOrderRefund() *OrderRefund`

NewOrderRefund instantiates a new OrderRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRefundWithDefaults

`func NewOrderRefundWithDefaults() *OrderRefund`

NewOrderRefundWithDefaults instantiates a new OrderRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderRefund) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderRefund) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderRefund) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderRefund) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OrderRefund) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OrderRefund) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetShipping

`func (o *OrderRefund) GetShipping() float32`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *OrderRefund) GetShippingOk() (*float32, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *OrderRefund) SetShipping(v float32)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *OrderRefund) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### SetShippingNil

`func (o *OrderRefund) SetShippingNil(b bool)`

 SetShippingNil sets the value for Shipping to be an explicit nil

### UnsetShipping
`func (o *OrderRefund) UnsetShipping()`

UnsetShipping ensures that no value is present for Shipping, not even an explicit nil
### GetFee

`func (o *OrderRefund) GetFee() float32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *OrderRefund) GetFeeOk() (*float32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *OrderRefund) SetFee(v float32)`

SetFee sets Fee field to given value.

### HasFee

`func (o *OrderRefund) HasFee() bool`

HasFee returns a boolean if a field has been set.

### SetFeeNil

`func (o *OrderRefund) SetFeeNil(b bool)`

 SetFeeNil sets the value for Fee to be an explicit nil

### UnsetFee
`func (o *OrderRefund) UnsetFee()`

UnsetFee ensures that no value is present for Fee, not even an explicit nil
### GetTax

`func (o *OrderRefund) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *OrderRefund) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *OrderRefund) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *OrderRefund) HasTax() bool`

HasTax returns a boolean if a field has been set.

### SetTaxNil

`func (o *OrderRefund) SetTaxNil(b bool)`

 SetTaxNil sets the value for Tax to be an explicit nil

### UnsetTax
`func (o *OrderRefund) UnsetTax()`

UnsetTax ensures that no value is present for Tax, not even an explicit nil
### GetTotal

`func (o *OrderRefund) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrderRefund) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrderRefund) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OrderRefund) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *OrderRefund) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *OrderRefund) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetModifiedTime

`func (o *OrderRefund) GetModifiedTime() A2CDateTime`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *OrderRefund) GetModifiedTimeOk() (*A2CDateTime, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *OrderRefund) SetModifiedTime(v A2CDateTime)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *OrderRefund) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### SetModifiedTimeNil

`func (o *OrderRefund) SetModifiedTimeNil(b bool)`

 SetModifiedTimeNil sets the value for ModifiedTime to be an explicit nil

### UnsetModifiedTime
`func (o *OrderRefund) UnsetModifiedTime()`

UnsetModifiedTime ensures that no value is present for ModifiedTime, not even an explicit nil
### GetComment

`func (o *OrderRefund) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *OrderRefund) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *OrderRefund) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *OrderRefund) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *OrderRefund) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *OrderRefund) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetItems

`func (o *OrderRefund) GetItems() []OrderStatusRefundItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrderRefund) GetItemsOk() (*[]OrderStatusRefundItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrderRefund) SetItems(v []OrderStatusRefundItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *OrderRefund) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *OrderRefund) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderRefund) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderRefund) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderRefund) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *OrderRefund) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *OrderRefund) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *OrderRefund) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderRefund) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderRefund) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderRefund) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *OrderRefund) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *OrderRefund) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


