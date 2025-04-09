# OrderTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**ReferenceNumber** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **float32** |  | [optional] 
**CreatedTime** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**SettlementCurrency** | Pointer to **string** |  | [optional] 
**SettlementAmount** | Pointer to **float32** |  | [optional] 
**SettlementCreatedTime** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**CardBrand** | Pointer to **string** |  | [optional] 
**CardBin** | Pointer to **string** |  | [optional] 
**CardLastFour** | Pointer to **string** |  | [optional] 
**AvsStreetRespCode** | Pointer to **string** |  | [optional] 
**AvsPostalRespCode** | Pointer to **string** |  | [optional] 
**AvsMessage** | Pointer to **string** |  | [optional] 
**CvvCode** | Pointer to **string** |  | [optional] 
**CvvMessage** | Pointer to **string** |  | [optional] 
**IsTestMode** | Pointer to **bool** |  | [optional] 
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


