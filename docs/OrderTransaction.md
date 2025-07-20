# OrderTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **NullableString** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 
**ReferenceNumber** | Pointer to **NullableString** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedTime** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**SettlementCurrency** | Pointer to **NullableString** |  | [optional] 
**SettlementAmount** | Pointer to **NullableFloat32** |  | [optional] 
**SettlementCreatedTime** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**CardBrand** | Pointer to **NullableString** |  | [optional] 
**CardBin** | Pointer to **NullableString** |  | [optional] 
**CardLastFour** | Pointer to **NullableString** |  | [optional] 
**AvsStreetRespCode** | Pointer to **NullableString** |  | [optional] 
**AvsPostalRespCode** | Pointer to **NullableString** |  | [optional] 
**AvsMessage** | Pointer to **NullableString** |  | [optional] 
**CvvCode** | Pointer to **NullableString** |  | [optional] 
**CvvMessage** | Pointer to **NullableString** |  | [optional] 
**IsTestMode** | Pointer to **NullableBool** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderTransaction

`func NewOrderTransaction() *OrderTransaction`

NewOrderTransaction instantiates a new OrderTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTransactionWithDefaults

`func NewOrderTransactionWithDefaults() *OrderTransaction`

NewOrderTransactionWithDefaults instantiates a new OrderTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTransactionId

`func (o *OrderTransaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *OrderTransaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *OrderTransaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *OrderTransaction) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### SetTransactionIdNil

`func (o *OrderTransaction) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *OrderTransaction) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetOrderId

`func (o *OrderTransaction) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderTransaction) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderTransaction) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderTransaction) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetParentId

`func (o *OrderTransaction) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *OrderTransaction) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *OrderTransaction) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *OrderTransaction) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *OrderTransaction) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *OrderTransaction) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetDescription

`func (o *OrderTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrderTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OrderTransaction) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OrderTransaction) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *OrderTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderTransaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *OrderTransaction) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *OrderTransaction) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetGateway

`func (o *OrderTransaction) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *OrderTransaction) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *OrderTransaction) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *OrderTransaction) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *OrderTransaction) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *OrderTransaction) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetReferenceNumber

`func (o *OrderTransaction) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *OrderTransaction) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *OrderTransaction) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *OrderTransaction) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### SetReferenceNumberNil

`func (o *OrderTransaction) SetReferenceNumberNil(b bool)`

 SetReferenceNumberNil sets the value for ReferenceNumber to be an explicit nil

### UnsetReferenceNumber
`func (o *OrderTransaction) UnsetReferenceNumber()`

UnsetReferenceNumber ensures that no value is present for ReferenceNumber, not even an explicit nil
### GetCurrency

`func (o *OrderTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderTransaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *OrderTransaction) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *OrderTransaction) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetAmount

`func (o *OrderTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OrderTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *OrderTransaction) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *OrderTransaction) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCreatedTime

`func (o *OrderTransaction) GetCreatedTime() A2CDateTime`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *OrderTransaction) GetCreatedTimeOk() (*A2CDateTime, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *OrderTransaction) SetCreatedTime(v A2CDateTime)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *OrderTransaction) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *OrderTransaction) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *OrderTransaction) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetSettlementCurrency

`func (o *OrderTransaction) GetSettlementCurrency() string`

GetSettlementCurrency returns the SettlementCurrency field if non-nil, zero value otherwise.

### GetSettlementCurrencyOk

`func (o *OrderTransaction) GetSettlementCurrencyOk() (*string, bool)`

GetSettlementCurrencyOk returns a tuple with the SettlementCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCurrency

`func (o *OrderTransaction) SetSettlementCurrency(v string)`

SetSettlementCurrency sets SettlementCurrency field to given value.

### HasSettlementCurrency

`func (o *OrderTransaction) HasSettlementCurrency() bool`

HasSettlementCurrency returns a boolean if a field has been set.

### SetSettlementCurrencyNil

`func (o *OrderTransaction) SetSettlementCurrencyNil(b bool)`

 SetSettlementCurrencyNil sets the value for SettlementCurrency to be an explicit nil

### UnsetSettlementCurrency
`func (o *OrderTransaction) UnsetSettlementCurrency()`

UnsetSettlementCurrency ensures that no value is present for SettlementCurrency, not even an explicit nil
### GetSettlementAmount

`func (o *OrderTransaction) GetSettlementAmount() float32`

GetSettlementAmount returns the SettlementAmount field if non-nil, zero value otherwise.

### GetSettlementAmountOk

`func (o *OrderTransaction) GetSettlementAmountOk() (*float32, bool)`

GetSettlementAmountOk returns a tuple with the SettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAmount

`func (o *OrderTransaction) SetSettlementAmount(v float32)`

SetSettlementAmount sets SettlementAmount field to given value.

### HasSettlementAmount

`func (o *OrderTransaction) HasSettlementAmount() bool`

HasSettlementAmount returns a boolean if a field has been set.

### SetSettlementAmountNil

`func (o *OrderTransaction) SetSettlementAmountNil(b bool)`

 SetSettlementAmountNil sets the value for SettlementAmount to be an explicit nil

### UnsetSettlementAmount
`func (o *OrderTransaction) UnsetSettlementAmount()`

UnsetSettlementAmount ensures that no value is present for SettlementAmount, not even an explicit nil
### GetSettlementCreatedTime

`func (o *OrderTransaction) GetSettlementCreatedTime() A2CDateTime`

GetSettlementCreatedTime returns the SettlementCreatedTime field if non-nil, zero value otherwise.

### GetSettlementCreatedTimeOk

`func (o *OrderTransaction) GetSettlementCreatedTimeOk() (*A2CDateTime, bool)`

GetSettlementCreatedTimeOk returns a tuple with the SettlementCreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCreatedTime

`func (o *OrderTransaction) SetSettlementCreatedTime(v A2CDateTime)`

SetSettlementCreatedTime sets SettlementCreatedTime field to given value.

### HasSettlementCreatedTime

`func (o *OrderTransaction) HasSettlementCreatedTime() bool`

HasSettlementCreatedTime returns a boolean if a field has been set.

### SetSettlementCreatedTimeNil

`func (o *OrderTransaction) SetSettlementCreatedTimeNil(b bool)`

 SetSettlementCreatedTimeNil sets the value for SettlementCreatedTime to be an explicit nil

### UnsetSettlementCreatedTime
`func (o *OrderTransaction) UnsetSettlementCreatedTime()`

UnsetSettlementCreatedTime ensures that no value is present for SettlementCreatedTime, not even an explicit nil
### GetCardBrand

`func (o *OrderTransaction) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *OrderTransaction) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *OrderTransaction) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *OrderTransaction) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### SetCardBrandNil

`func (o *OrderTransaction) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *OrderTransaction) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardBin

`func (o *OrderTransaction) GetCardBin() string`

GetCardBin returns the CardBin field if non-nil, zero value otherwise.

### GetCardBinOk

`func (o *OrderTransaction) GetCardBinOk() (*string, bool)`

GetCardBinOk returns a tuple with the CardBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBin

`func (o *OrderTransaction) SetCardBin(v string)`

SetCardBin sets CardBin field to given value.

### HasCardBin

`func (o *OrderTransaction) HasCardBin() bool`

HasCardBin returns a boolean if a field has been set.

### SetCardBinNil

`func (o *OrderTransaction) SetCardBinNil(b bool)`

 SetCardBinNil sets the value for CardBin to be an explicit nil

### UnsetCardBin
`func (o *OrderTransaction) UnsetCardBin()`

UnsetCardBin ensures that no value is present for CardBin, not even an explicit nil
### GetCardLastFour

`func (o *OrderTransaction) GetCardLastFour() string`

GetCardLastFour returns the CardLastFour field if non-nil, zero value otherwise.

### GetCardLastFourOk

`func (o *OrderTransaction) GetCardLastFourOk() (*string, bool)`

GetCardLastFourOk returns a tuple with the CardLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLastFour

`func (o *OrderTransaction) SetCardLastFour(v string)`

SetCardLastFour sets CardLastFour field to given value.

### HasCardLastFour

`func (o *OrderTransaction) HasCardLastFour() bool`

HasCardLastFour returns a boolean if a field has been set.

### SetCardLastFourNil

`func (o *OrderTransaction) SetCardLastFourNil(b bool)`

 SetCardLastFourNil sets the value for CardLastFour to be an explicit nil

### UnsetCardLastFour
`func (o *OrderTransaction) UnsetCardLastFour()`

UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
### GetAvsStreetRespCode

`func (o *OrderTransaction) GetAvsStreetRespCode() string`

GetAvsStreetRespCode returns the AvsStreetRespCode field if non-nil, zero value otherwise.

### GetAvsStreetRespCodeOk

`func (o *OrderTransaction) GetAvsStreetRespCodeOk() (*string, bool)`

GetAvsStreetRespCodeOk returns a tuple with the AvsStreetRespCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsStreetRespCode

`func (o *OrderTransaction) SetAvsStreetRespCode(v string)`

SetAvsStreetRespCode sets AvsStreetRespCode field to given value.

### HasAvsStreetRespCode

`func (o *OrderTransaction) HasAvsStreetRespCode() bool`

HasAvsStreetRespCode returns a boolean if a field has been set.

### SetAvsStreetRespCodeNil

`func (o *OrderTransaction) SetAvsStreetRespCodeNil(b bool)`

 SetAvsStreetRespCodeNil sets the value for AvsStreetRespCode to be an explicit nil

### UnsetAvsStreetRespCode
`func (o *OrderTransaction) UnsetAvsStreetRespCode()`

UnsetAvsStreetRespCode ensures that no value is present for AvsStreetRespCode, not even an explicit nil
### GetAvsPostalRespCode

`func (o *OrderTransaction) GetAvsPostalRespCode() string`

GetAvsPostalRespCode returns the AvsPostalRespCode field if non-nil, zero value otherwise.

### GetAvsPostalRespCodeOk

`func (o *OrderTransaction) GetAvsPostalRespCodeOk() (*string, bool)`

GetAvsPostalRespCodeOk returns a tuple with the AvsPostalRespCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsPostalRespCode

`func (o *OrderTransaction) SetAvsPostalRespCode(v string)`

SetAvsPostalRespCode sets AvsPostalRespCode field to given value.

### HasAvsPostalRespCode

`func (o *OrderTransaction) HasAvsPostalRespCode() bool`

HasAvsPostalRespCode returns a boolean if a field has been set.

### SetAvsPostalRespCodeNil

`func (o *OrderTransaction) SetAvsPostalRespCodeNil(b bool)`

 SetAvsPostalRespCodeNil sets the value for AvsPostalRespCode to be an explicit nil

### UnsetAvsPostalRespCode
`func (o *OrderTransaction) UnsetAvsPostalRespCode()`

UnsetAvsPostalRespCode ensures that no value is present for AvsPostalRespCode, not even an explicit nil
### GetAvsMessage

`func (o *OrderTransaction) GetAvsMessage() string`

GetAvsMessage returns the AvsMessage field if non-nil, zero value otherwise.

### GetAvsMessageOk

`func (o *OrderTransaction) GetAvsMessageOk() (*string, bool)`

GetAvsMessageOk returns a tuple with the AvsMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsMessage

`func (o *OrderTransaction) SetAvsMessage(v string)`

SetAvsMessage sets AvsMessage field to given value.

### HasAvsMessage

`func (o *OrderTransaction) HasAvsMessage() bool`

HasAvsMessage returns a boolean if a field has been set.

### SetAvsMessageNil

`func (o *OrderTransaction) SetAvsMessageNil(b bool)`

 SetAvsMessageNil sets the value for AvsMessage to be an explicit nil

### UnsetAvsMessage
`func (o *OrderTransaction) UnsetAvsMessage()`

UnsetAvsMessage ensures that no value is present for AvsMessage, not even an explicit nil
### GetCvvCode

`func (o *OrderTransaction) GetCvvCode() string`

GetCvvCode returns the CvvCode field if non-nil, zero value otherwise.

### GetCvvCodeOk

`func (o *OrderTransaction) GetCvvCodeOk() (*string, bool)`

GetCvvCodeOk returns a tuple with the CvvCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvCode

`func (o *OrderTransaction) SetCvvCode(v string)`

SetCvvCode sets CvvCode field to given value.

### HasCvvCode

`func (o *OrderTransaction) HasCvvCode() bool`

HasCvvCode returns a boolean if a field has been set.

### SetCvvCodeNil

`func (o *OrderTransaction) SetCvvCodeNil(b bool)`

 SetCvvCodeNil sets the value for CvvCode to be an explicit nil

### UnsetCvvCode
`func (o *OrderTransaction) UnsetCvvCode()`

UnsetCvvCode ensures that no value is present for CvvCode, not even an explicit nil
### GetCvvMessage

`func (o *OrderTransaction) GetCvvMessage() string`

GetCvvMessage returns the CvvMessage field if non-nil, zero value otherwise.

### GetCvvMessageOk

`func (o *OrderTransaction) GetCvvMessageOk() (*string, bool)`

GetCvvMessageOk returns a tuple with the CvvMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvMessage

`func (o *OrderTransaction) SetCvvMessage(v string)`

SetCvvMessage sets CvvMessage field to given value.

### HasCvvMessage

`func (o *OrderTransaction) HasCvvMessage() bool`

HasCvvMessage returns a boolean if a field has been set.

### SetCvvMessageNil

`func (o *OrderTransaction) SetCvvMessageNil(b bool)`

 SetCvvMessageNil sets the value for CvvMessage to be an explicit nil

### UnsetCvvMessage
`func (o *OrderTransaction) UnsetCvvMessage()`

UnsetCvvMessage ensures that no value is present for CvvMessage, not even an explicit nil
### GetIsTestMode

`func (o *OrderTransaction) GetIsTestMode() bool`

GetIsTestMode returns the IsTestMode field if non-nil, zero value otherwise.

### GetIsTestModeOk

`func (o *OrderTransaction) GetIsTestModeOk() (*bool, bool)`

GetIsTestModeOk returns a tuple with the IsTestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestMode

`func (o *OrderTransaction) SetIsTestMode(v bool)`

SetIsTestMode sets IsTestMode field to given value.

### HasIsTestMode

`func (o *OrderTransaction) HasIsTestMode() bool`

HasIsTestMode returns a boolean if a field has been set.

### SetIsTestModeNil

`func (o *OrderTransaction) SetIsTestModeNil(b bool)`

 SetIsTestModeNil sets the value for IsTestMode to be an explicit nil

### UnsetIsTestMode
`func (o *OrderTransaction) UnsetIsTestMode()`

UnsetIsTestMode ensures that no value is present for IsTestMode, not even an explicit nil
### GetAdditionalFields

`func (o *OrderTransaction) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderTransaction) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderTransaction) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderTransaction) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *OrderTransaction) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *OrderTransaction) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *OrderTransaction) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderTransaction) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderTransaction) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderTransaction) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *OrderTransaction) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *OrderTransaction) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


