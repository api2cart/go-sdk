# GiftCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**InitialAmount** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**AvailTo** | Pointer to **NullableString** |  | [optional] 
**FreeProductIds** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**IssuerEmail** | Pointer to **NullableString** |  | [optional] 
**RecipientEmail** | Pointer to **NullableString** |  | [optional] 
**IssuerName** | Pointer to **NullableString** |  | [optional] 
**RecipientName** | Pointer to **NullableString** |  | [optional] 
**UsageHistory** | Pointer to [**[]CouponHistory**](CouponHistory.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGiftCard

`func NewGiftCard() *GiftCard`

NewGiftCard instantiates a new GiftCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardWithDefaults

`func NewGiftCardWithDefaults() *GiftCard`

NewGiftCardWithDefaults instantiates a new GiftCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GiftCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GiftCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GiftCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GiftCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GiftCard) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GiftCard) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GiftCard) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GiftCard) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GiftCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GiftCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GiftCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GiftCard) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GiftCard) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GiftCard) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *GiftCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GiftCard) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *GiftCard) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GiftCard) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GiftCard) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GiftCard) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GiftCard) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GiftCard) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetAmount

`func (o *GiftCard) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GiftCard) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GiftCard) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GiftCard) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetInitialAmount

`func (o *GiftCard) GetInitialAmount() float32`

GetInitialAmount returns the InitialAmount field if non-nil, zero value otherwise.

### GetInitialAmountOk

`func (o *GiftCard) GetInitialAmountOk() (*float32, bool)`

GetInitialAmountOk returns a tuple with the InitialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAmount

`func (o *GiftCard) SetInitialAmount(v float32)`

SetInitialAmount sets InitialAmount field to given value.

### HasInitialAmount

`func (o *GiftCard) HasInitialAmount() bool`

HasInitialAmount returns a boolean if a field has been set.

### GetStatus

`func (o *GiftCard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GiftCard) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GiftCard) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GiftCard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GiftCard) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GiftCard) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GiftCard) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GiftCard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GiftCard) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GiftCard) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetAvailTo

`func (o *GiftCard) GetAvailTo() string`

GetAvailTo returns the AvailTo field if non-nil, zero value otherwise.

### GetAvailToOk

`func (o *GiftCard) GetAvailToOk() (*string, bool)`

GetAvailToOk returns a tuple with the AvailTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailTo

`func (o *GiftCard) SetAvailTo(v string)`

SetAvailTo sets AvailTo field to given value.

### HasAvailTo

`func (o *GiftCard) HasAvailTo() bool`

HasAvailTo returns a boolean if a field has been set.

### SetAvailToNil

`func (o *GiftCard) SetAvailToNil(b bool)`

 SetAvailToNil sets the value for AvailTo to be an explicit nil

### UnsetAvailTo
`func (o *GiftCard) UnsetAvailTo()`

UnsetAvailTo ensures that no value is present for AvailTo, not even an explicit nil
### GetFreeProductIds

`func (o *GiftCard) GetFreeProductIds() string`

GetFreeProductIds returns the FreeProductIds field if non-nil, zero value otherwise.

### GetFreeProductIdsOk

`func (o *GiftCard) GetFreeProductIdsOk() (*string, bool)`

GetFreeProductIdsOk returns a tuple with the FreeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeProductIds

`func (o *GiftCard) SetFreeProductIds(v string)`

SetFreeProductIds sets FreeProductIds field to given value.

### HasFreeProductIds

`func (o *GiftCard) HasFreeProductIds() bool`

HasFreeProductIds returns a boolean if a field has been set.

### SetFreeProductIdsNil

`func (o *GiftCard) SetFreeProductIdsNil(b bool)`

 SetFreeProductIdsNil sets the value for FreeProductIds to be an explicit nil

### UnsetFreeProductIds
`func (o *GiftCard) UnsetFreeProductIds()`

UnsetFreeProductIds ensures that no value is present for FreeProductIds, not even an explicit nil
### GetMessage

`func (o *GiftCard) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GiftCard) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GiftCard) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GiftCard) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GiftCard) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GiftCard) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetIssuerEmail

`func (o *GiftCard) GetIssuerEmail() string`

GetIssuerEmail returns the IssuerEmail field if non-nil, zero value otherwise.

### GetIssuerEmailOk

`func (o *GiftCard) GetIssuerEmailOk() (*string, bool)`

GetIssuerEmailOk returns a tuple with the IssuerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerEmail

`func (o *GiftCard) SetIssuerEmail(v string)`

SetIssuerEmail sets IssuerEmail field to given value.

### HasIssuerEmail

`func (o *GiftCard) HasIssuerEmail() bool`

HasIssuerEmail returns a boolean if a field has been set.

### SetIssuerEmailNil

`func (o *GiftCard) SetIssuerEmailNil(b bool)`

 SetIssuerEmailNil sets the value for IssuerEmail to be an explicit nil

### UnsetIssuerEmail
`func (o *GiftCard) UnsetIssuerEmail()`

UnsetIssuerEmail ensures that no value is present for IssuerEmail, not even an explicit nil
### GetRecipientEmail

`func (o *GiftCard) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *GiftCard) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *GiftCard) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *GiftCard) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *GiftCard) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *GiftCard) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetIssuerName

`func (o *GiftCard) GetIssuerName() string`

GetIssuerName returns the IssuerName field if non-nil, zero value otherwise.

### GetIssuerNameOk

`func (o *GiftCard) GetIssuerNameOk() (*string, bool)`

GetIssuerNameOk returns a tuple with the IssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerName

`func (o *GiftCard) SetIssuerName(v string)`

SetIssuerName sets IssuerName field to given value.

### HasIssuerName

`func (o *GiftCard) HasIssuerName() bool`

HasIssuerName returns a boolean if a field has been set.

### SetIssuerNameNil

`func (o *GiftCard) SetIssuerNameNil(b bool)`

 SetIssuerNameNil sets the value for IssuerName to be an explicit nil

### UnsetIssuerName
`func (o *GiftCard) UnsetIssuerName()`

UnsetIssuerName ensures that no value is present for IssuerName, not even an explicit nil
### GetRecipientName

`func (o *GiftCard) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *GiftCard) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *GiftCard) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *GiftCard) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### SetRecipientNameNil

`func (o *GiftCard) SetRecipientNameNil(b bool)`

 SetRecipientNameNil sets the value for RecipientName to be an explicit nil

### UnsetRecipientName
`func (o *GiftCard) UnsetRecipientName()`

UnsetRecipientName ensures that no value is present for RecipientName, not even an explicit nil
### GetUsageHistory

`func (o *GiftCard) GetUsageHistory() []CouponHistory`

GetUsageHistory returns the UsageHistory field if non-nil, zero value otherwise.

### GetUsageHistoryOk

`func (o *GiftCard) GetUsageHistoryOk() (*[]CouponHistory, bool)`

GetUsageHistoryOk returns a tuple with the UsageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageHistory

`func (o *GiftCard) SetUsageHistory(v []CouponHistory)`

SetUsageHistory sets UsageHistory field to given value.

### HasUsageHistory

`func (o *GiftCard) HasUsageHistory() bool`

HasUsageHistory returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *GiftCard) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *GiftCard) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *GiftCard) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *GiftCard) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *GiftCard) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *GiftCard) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *GiftCard) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GiftCard) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GiftCard) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GiftCard) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *GiftCard) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *GiftCard) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


