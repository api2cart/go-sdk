/*
API2Cart OpenAPI

API2Cart

API version: 1.1
Contact: contact@api2cart.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OrderTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderTransaction{}

// OrderTransaction struct for OrderTransaction
type OrderTransaction struct {
	Id *string `json:"id,omitempty"`
	TransactionId NullableString `json:"transaction_id,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
	ParentId NullableString `json:"parent_id,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Status NullableString `json:"status,omitempty"`
	Gateway NullableString `json:"gateway,omitempty"`
	ReferenceNumber NullableString `json:"reference_number,omitempty"`
	Currency NullableString `json:"currency,omitempty"`
	Amount NullableFloat32 `json:"amount,omitempty"`
	CreatedTime NullableA2CDateTime `json:"created_time,omitempty"`
	SettlementCurrency NullableString `json:"settlement_currency,omitempty"`
	SettlementAmount NullableFloat32 `json:"settlement_amount,omitempty"`
	SettlementCreatedTime NullableA2CDateTime `json:"settlement_created_time,omitempty"`
	CardBrand NullableString `json:"card_brand,omitempty"`
	CardBin NullableString `json:"card_bin,omitempty"`
	CardLastFour NullableString `json:"card_last_four,omitempty"`
	AvsStreetRespCode NullableString `json:"avs_street_resp_code,omitempty"`
	AvsPostalRespCode NullableString `json:"avs_postal_resp_code,omitempty"`
	AvsMessage NullableString `json:"avs_message,omitempty"`
	CvvCode NullableString `json:"cvv_code,omitempty"`
	CvvMessage NullableString `json:"cvv_message,omitempty"`
	IsTestMode NullableBool `json:"is_test_mode,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewOrderTransaction instantiates a new OrderTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTransaction() *OrderTransaction {
	this := OrderTransaction{}
	return &this
}

// NewOrderTransactionWithDefaults instantiates a new OrderTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTransactionWithDefaults() *OrderTransaction {
	this := OrderTransaction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderTransaction) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTransaction) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderTransaction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderTransaction) SetId(v string) {
	o.Id = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionId.Get()
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionId.Get(), o.TransactionId.IsSet()
}

// HasTransactionId returns a boolean if a field has been set.
func (o *OrderTransaction) HasTransactionId() bool {
	if o != nil && o.TransactionId.IsSet() {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given NullableString and assigns it to the TransactionId field.
func (o *OrderTransaction) SetTransactionId(v string) {
	o.TransactionId.Set(&v)
}
// SetTransactionIdNil sets the value for TransactionId to be an explicit nil
func (o *OrderTransaction) SetTransactionIdNil() {
	o.TransactionId.Set(nil)
}

// UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
func (o *OrderTransaction) UnsetTransactionId() {
	o.TransactionId.Unset()
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderTransaction) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTransaction) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderTransaction) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderTransaction) SetOrderId(v string) {
	o.OrderId = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *OrderTransaction) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *OrderTransaction) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *OrderTransaction) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *OrderTransaction) UnsetParentId() {
	o.ParentId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *OrderTransaction) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *OrderTransaction) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *OrderTransaction) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *OrderTransaction) UnsetDescription() {
	o.Description.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderTransaction) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *OrderTransaction) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *OrderTransaction) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *OrderTransaction) UnsetStatus() {
	o.Status.Unset()
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetGateway() string {
	if o == nil || IsNil(o.Gateway.Get()) {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *OrderTransaction) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *OrderTransaction) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *OrderTransaction) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *OrderTransaction) UnsetGateway() {
	o.Gateway.Unset()
}

// GetReferenceNumber returns the ReferenceNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetReferenceNumber() string {
	if o == nil || IsNil(o.ReferenceNumber.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceNumber.Get()
}

// GetReferenceNumberOk returns a tuple with the ReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetReferenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceNumber.Get(), o.ReferenceNumber.IsSet()
}

// HasReferenceNumber returns a boolean if a field has been set.
func (o *OrderTransaction) HasReferenceNumber() bool {
	if o != nil && o.ReferenceNumber.IsSet() {
		return true
	}

	return false
}

// SetReferenceNumber gets a reference to the given NullableString and assigns it to the ReferenceNumber field.
func (o *OrderTransaction) SetReferenceNumber(v string) {
	o.ReferenceNumber.Set(&v)
}
// SetReferenceNumberNil sets the value for ReferenceNumber to be an explicit nil
func (o *OrderTransaction) SetReferenceNumberNil() {
	o.ReferenceNumber.Set(nil)
}

// UnsetReferenceNumber ensures that no value is present for ReferenceNumber, not even an explicit nil
func (o *OrderTransaction) UnsetReferenceNumber() {
	o.ReferenceNumber.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *OrderTransaction) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *OrderTransaction) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *OrderTransaction) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *OrderTransaction) UnsetCurrency() {
	o.Currency.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetAmount() float32 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *OrderTransaction) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat32 and assigns it to the Amount field.
func (o *OrderTransaction) SetAmount(v float32) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *OrderTransaction) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *OrderTransaction) UnsetAmount() {
	o.Amount.Unset()
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetCreatedTime() A2CDateTime {
	if o == nil || IsNil(o.CreatedTime.Get()) {
		var ret A2CDateTime
		return ret
	}
	return *o.CreatedTime.Get()
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetCreatedTimeOk() (*A2CDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedTime.Get(), o.CreatedTime.IsSet()
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *OrderTransaction) HasCreatedTime() bool {
	if o != nil && o.CreatedTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given NullableA2CDateTime and assigns it to the CreatedTime field.
func (o *OrderTransaction) SetCreatedTime(v A2CDateTime) {
	o.CreatedTime.Set(&v)
}
// SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil
func (o *OrderTransaction) SetCreatedTimeNil() {
	o.CreatedTime.Set(nil)
}

// UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
func (o *OrderTransaction) UnsetCreatedTime() {
	o.CreatedTime.Unset()
}

// GetSettlementCurrency returns the SettlementCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetSettlementCurrency() string {
	if o == nil || IsNil(o.SettlementCurrency.Get()) {
		var ret string
		return ret
	}
	return *o.SettlementCurrency.Get()
}

// GetSettlementCurrencyOk returns a tuple with the SettlementCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetSettlementCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettlementCurrency.Get(), o.SettlementCurrency.IsSet()
}

// HasSettlementCurrency returns a boolean if a field has been set.
func (o *OrderTransaction) HasSettlementCurrency() bool {
	if o != nil && o.SettlementCurrency.IsSet() {
		return true
	}

	return false
}

// SetSettlementCurrency gets a reference to the given NullableString and assigns it to the SettlementCurrency field.
func (o *OrderTransaction) SetSettlementCurrency(v string) {
	o.SettlementCurrency.Set(&v)
}
// SetSettlementCurrencyNil sets the value for SettlementCurrency to be an explicit nil
func (o *OrderTransaction) SetSettlementCurrencyNil() {
	o.SettlementCurrency.Set(nil)
}

// UnsetSettlementCurrency ensures that no value is present for SettlementCurrency, not even an explicit nil
func (o *OrderTransaction) UnsetSettlementCurrency() {
	o.SettlementCurrency.Unset()
}

// GetSettlementAmount returns the SettlementAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetSettlementAmount() float32 {
	if o == nil || IsNil(o.SettlementAmount.Get()) {
		var ret float32
		return ret
	}
	return *o.SettlementAmount.Get()
}

// GetSettlementAmountOk returns a tuple with the SettlementAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetSettlementAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettlementAmount.Get(), o.SettlementAmount.IsSet()
}

// HasSettlementAmount returns a boolean if a field has been set.
func (o *OrderTransaction) HasSettlementAmount() bool {
	if o != nil && o.SettlementAmount.IsSet() {
		return true
	}

	return false
}

// SetSettlementAmount gets a reference to the given NullableFloat32 and assigns it to the SettlementAmount field.
func (o *OrderTransaction) SetSettlementAmount(v float32) {
	o.SettlementAmount.Set(&v)
}
// SetSettlementAmountNil sets the value for SettlementAmount to be an explicit nil
func (o *OrderTransaction) SetSettlementAmountNil() {
	o.SettlementAmount.Set(nil)
}

// UnsetSettlementAmount ensures that no value is present for SettlementAmount, not even an explicit nil
func (o *OrderTransaction) UnsetSettlementAmount() {
	o.SettlementAmount.Unset()
}

// GetSettlementCreatedTime returns the SettlementCreatedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetSettlementCreatedTime() A2CDateTime {
	if o == nil || IsNil(o.SettlementCreatedTime.Get()) {
		var ret A2CDateTime
		return ret
	}
	return *o.SettlementCreatedTime.Get()
}

// GetSettlementCreatedTimeOk returns a tuple with the SettlementCreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetSettlementCreatedTimeOk() (*A2CDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettlementCreatedTime.Get(), o.SettlementCreatedTime.IsSet()
}

// HasSettlementCreatedTime returns a boolean if a field has been set.
func (o *OrderTransaction) HasSettlementCreatedTime() bool {
	if o != nil && o.SettlementCreatedTime.IsSet() {
		return true
	}

	return false
}

// SetSettlementCreatedTime gets a reference to the given NullableA2CDateTime and assigns it to the SettlementCreatedTime field.
func (o *OrderTransaction) SetSettlementCreatedTime(v A2CDateTime) {
	o.SettlementCreatedTime.Set(&v)
}
// SetSettlementCreatedTimeNil sets the value for SettlementCreatedTime to be an explicit nil
func (o *OrderTransaction) SetSettlementCreatedTimeNil() {
	o.SettlementCreatedTime.Set(nil)
}

// UnsetSettlementCreatedTime ensures that no value is present for SettlementCreatedTime, not even an explicit nil
func (o *OrderTransaction) UnsetSettlementCreatedTime() {
	o.SettlementCreatedTime.Unset()
}

// GetCardBrand returns the CardBrand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetCardBrand() string {
	if o == nil || IsNil(o.CardBrand.Get()) {
		var ret string
		return ret
	}
	return *o.CardBrand.Get()
}

// GetCardBrandOk returns a tuple with the CardBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetCardBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardBrand.Get(), o.CardBrand.IsSet()
}

// HasCardBrand returns a boolean if a field has been set.
func (o *OrderTransaction) HasCardBrand() bool {
	if o != nil && o.CardBrand.IsSet() {
		return true
	}

	return false
}

// SetCardBrand gets a reference to the given NullableString and assigns it to the CardBrand field.
func (o *OrderTransaction) SetCardBrand(v string) {
	o.CardBrand.Set(&v)
}
// SetCardBrandNil sets the value for CardBrand to be an explicit nil
func (o *OrderTransaction) SetCardBrandNil() {
	o.CardBrand.Set(nil)
}

// UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
func (o *OrderTransaction) UnsetCardBrand() {
	o.CardBrand.Unset()
}

// GetCardBin returns the CardBin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetCardBin() string {
	if o == nil || IsNil(o.CardBin.Get()) {
		var ret string
		return ret
	}
	return *o.CardBin.Get()
}

// GetCardBinOk returns a tuple with the CardBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetCardBinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardBin.Get(), o.CardBin.IsSet()
}

// HasCardBin returns a boolean if a field has been set.
func (o *OrderTransaction) HasCardBin() bool {
	if o != nil && o.CardBin.IsSet() {
		return true
	}

	return false
}

// SetCardBin gets a reference to the given NullableString and assigns it to the CardBin field.
func (o *OrderTransaction) SetCardBin(v string) {
	o.CardBin.Set(&v)
}
// SetCardBinNil sets the value for CardBin to be an explicit nil
func (o *OrderTransaction) SetCardBinNil() {
	o.CardBin.Set(nil)
}

// UnsetCardBin ensures that no value is present for CardBin, not even an explicit nil
func (o *OrderTransaction) UnsetCardBin() {
	o.CardBin.Unset()
}

// GetCardLastFour returns the CardLastFour field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetCardLastFour() string {
	if o == nil || IsNil(o.CardLastFour.Get()) {
		var ret string
		return ret
	}
	return *o.CardLastFour.Get()
}

// GetCardLastFourOk returns a tuple with the CardLastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetCardLastFourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardLastFour.Get(), o.CardLastFour.IsSet()
}

// HasCardLastFour returns a boolean if a field has been set.
func (o *OrderTransaction) HasCardLastFour() bool {
	if o != nil && o.CardLastFour.IsSet() {
		return true
	}

	return false
}

// SetCardLastFour gets a reference to the given NullableString and assigns it to the CardLastFour field.
func (o *OrderTransaction) SetCardLastFour(v string) {
	o.CardLastFour.Set(&v)
}
// SetCardLastFourNil sets the value for CardLastFour to be an explicit nil
func (o *OrderTransaction) SetCardLastFourNil() {
	o.CardLastFour.Set(nil)
}

// UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
func (o *OrderTransaction) UnsetCardLastFour() {
	o.CardLastFour.Unset()
}

// GetAvsStreetRespCode returns the AvsStreetRespCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetAvsStreetRespCode() string {
	if o == nil || IsNil(o.AvsStreetRespCode.Get()) {
		var ret string
		return ret
	}
	return *o.AvsStreetRespCode.Get()
}

// GetAvsStreetRespCodeOk returns a tuple with the AvsStreetRespCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetAvsStreetRespCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvsStreetRespCode.Get(), o.AvsStreetRespCode.IsSet()
}

// HasAvsStreetRespCode returns a boolean if a field has been set.
func (o *OrderTransaction) HasAvsStreetRespCode() bool {
	if o != nil && o.AvsStreetRespCode.IsSet() {
		return true
	}

	return false
}

// SetAvsStreetRespCode gets a reference to the given NullableString and assigns it to the AvsStreetRespCode field.
func (o *OrderTransaction) SetAvsStreetRespCode(v string) {
	o.AvsStreetRespCode.Set(&v)
}
// SetAvsStreetRespCodeNil sets the value for AvsStreetRespCode to be an explicit nil
func (o *OrderTransaction) SetAvsStreetRespCodeNil() {
	o.AvsStreetRespCode.Set(nil)
}

// UnsetAvsStreetRespCode ensures that no value is present for AvsStreetRespCode, not even an explicit nil
func (o *OrderTransaction) UnsetAvsStreetRespCode() {
	o.AvsStreetRespCode.Unset()
}

// GetAvsPostalRespCode returns the AvsPostalRespCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetAvsPostalRespCode() string {
	if o == nil || IsNil(o.AvsPostalRespCode.Get()) {
		var ret string
		return ret
	}
	return *o.AvsPostalRespCode.Get()
}

// GetAvsPostalRespCodeOk returns a tuple with the AvsPostalRespCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetAvsPostalRespCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvsPostalRespCode.Get(), o.AvsPostalRespCode.IsSet()
}

// HasAvsPostalRespCode returns a boolean if a field has been set.
func (o *OrderTransaction) HasAvsPostalRespCode() bool {
	if o != nil && o.AvsPostalRespCode.IsSet() {
		return true
	}

	return false
}

// SetAvsPostalRespCode gets a reference to the given NullableString and assigns it to the AvsPostalRespCode field.
func (o *OrderTransaction) SetAvsPostalRespCode(v string) {
	o.AvsPostalRespCode.Set(&v)
}
// SetAvsPostalRespCodeNil sets the value for AvsPostalRespCode to be an explicit nil
func (o *OrderTransaction) SetAvsPostalRespCodeNil() {
	o.AvsPostalRespCode.Set(nil)
}

// UnsetAvsPostalRespCode ensures that no value is present for AvsPostalRespCode, not even an explicit nil
func (o *OrderTransaction) UnsetAvsPostalRespCode() {
	o.AvsPostalRespCode.Unset()
}

// GetAvsMessage returns the AvsMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetAvsMessage() string {
	if o == nil || IsNil(o.AvsMessage.Get()) {
		var ret string
		return ret
	}
	return *o.AvsMessage.Get()
}

// GetAvsMessageOk returns a tuple with the AvsMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetAvsMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvsMessage.Get(), o.AvsMessage.IsSet()
}

// HasAvsMessage returns a boolean if a field has been set.
func (o *OrderTransaction) HasAvsMessage() bool {
	if o != nil && o.AvsMessage.IsSet() {
		return true
	}

	return false
}

// SetAvsMessage gets a reference to the given NullableString and assigns it to the AvsMessage field.
func (o *OrderTransaction) SetAvsMessage(v string) {
	o.AvsMessage.Set(&v)
}
// SetAvsMessageNil sets the value for AvsMessage to be an explicit nil
func (o *OrderTransaction) SetAvsMessageNil() {
	o.AvsMessage.Set(nil)
}

// UnsetAvsMessage ensures that no value is present for AvsMessage, not even an explicit nil
func (o *OrderTransaction) UnsetAvsMessage() {
	o.AvsMessage.Unset()
}

// GetCvvCode returns the CvvCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetCvvCode() string {
	if o == nil || IsNil(o.CvvCode.Get()) {
		var ret string
		return ret
	}
	return *o.CvvCode.Get()
}

// GetCvvCodeOk returns a tuple with the CvvCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetCvvCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CvvCode.Get(), o.CvvCode.IsSet()
}

// HasCvvCode returns a boolean if a field has been set.
func (o *OrderTransaction) HasCvvCode() bool {
	if o != nil && o.CvvCode.IsSet() {
		return true
	}

	return false
}

// SetCvvCode gets a reference to the given NullableString and assigns it to the CvvCode field.
func (o *OrderTransaction) SetCvvCode(v string) {
	o.CvvCode.Set(&v)
}
// SetCvvCodeNil sets the value for CvvCode to be an explicit nil
func (o *OrderTransaction) SetCvvCodeNil() {
	o.CvvCode.Set(nil)
}

// UnsetCvvCode ensures that no value is present for CvvCode, not even an explicit nil
func (o *OrderTransaction) UnsetCvvCode() {
	o.CvvCode.Unset()
}

// GetCvvMessage returns the CvvMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetCvvMessage() string {
	if o == nil || IsNil(o.CvvMessage.Get()) {
		var ret string
		return ret
	}
	return *o.CvvMessage.Get()
}

// GetCvvMessageOk returns a tuple with the CvvMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetCvvMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CvvMessage.Get(), o.CvvMessage.IsSet()
}

// HasCvvMessage returns a boolean if a field has been set.
func (o *OrderTransaction) HasCvvMessage() bool {
	if o != nil && o.CvvMessage.IsSet() {
		return true
	}

	return false
}

// SetCvvMessage gets a reference to the given NullableString and assigns it to the CvvMessage field.
func (o *OrderTransaction) SetCvvMessage(v string) {
	o.CvvMessage.Set(&v)
}
// SetCvvMessageNil sets the value for CvvMessage to be an explicit nil
func (o *OrderTransaction) SetCvvMessageNil() {
	o.CvvMessage.Set(nil)
}

// UnsetCvvMessage ensures that no value is present for CvvMessage, not even an explicit nil
func (o *OrderTransaction) UnsetCvvMessage() {
	o.CvvMessage.Unset()
}

// GetIsTestMode returns the IsTestMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetIsTestMode() bool {
	if o == nil || IsNil(o.IsTestMode.Get()) {
		var ret bool
		return ret
	}
	return *o.IsTestMode.Get()
}

// GetIsTestModeOk returns a tuple with the IsTestMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetIsTestModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsTestMode.Get(), o.IsTestMode.IsSet()
}

// HasIsTestMode returns a boolean if a field has been set.
func (o *OrderTransaction) HasIsTestMode() bool {
	if o != nil && o.IsTestMode.IsSet() {
		return true
	}

	return false
}

// SetIsTestMode gets a reference to the given NullableBool and assigns it to the IsTestMode field.
func (o *OrderTransaction) SetIsTestMode(v bool) {
	o.IsTestMode.Set(&v)
}
// SetIsTestModeNil sets the value for IsTestMode to be an explicit nil
func (o *OrderTransaction) SetIsTestModeNil() {
	o.IsTestMode.Set(nil)
}

// UnsetIsTestMode ensures that no value is present for IsTestMode, not even an explicit nil
func (o *OrderTransaction) UnsetIsTestMode() {
	o.IsTestMode.Unset()
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetAdditionalFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *OrderTransaction) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *OrderTransaction) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderTransaction) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderTransaction) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *OrderTransaction) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *OrderTransaction) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o OrderTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.TransactionId.IsSet() {
		toSerialize["transaction_id"] = o.TransactionId.Get()
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if o.ParentId.IsSet() {
		toSerialize["parent_id"] = o.ParentId.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Gateway.IsSet() {
		toSerialize["gateway"] = o.Gateway.Get()
	}
	if o.ReferenceNumber.IsSet() {
		toSerialize["reference_number"] = o.ReferenceNumber.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.CreatedTime.IsSet() {
		toSerialize["created_time"] = o.CreatedTime.Get()
	}
	if o.SettlementCurrency.IsSet() {
		toSerialize["settlement_currency"] = o.SettlementCurrency.Get()
	}
	if o.SettlementAmount.IsSet() {
		toSerialize["settlement_amount"] = o.SettlementAmount.Get()
	}
	if o.SettlementCreatedTime.IsSet() {
		toSerialize["settlement_created_time"] = o.SettlementCreatedTime.Get()
	}
	if o.CardBrand.IsSet() {
		toSerialize["card_brand"] = o.CardBrand.Get()
	}
	if o.CardBin.IsSet() {
		toSerialize["card_bin"] = o.CardBin.Get()
	}
	if o.CardLastFour.IsSet() {
		toSerialize["card_last_four"] = o.CardLastFour.Get()
	}
	if o.AvsStreetRespCode.IsSet() {
		toSerialize["avs_street_resp_code"] = o.AvsStreetRespCode.Get()
	}
	if o.AvsPostalRespCode.IsSet() {
		toSerialize["avs_postal_resp_code"] = o.AvsPostalRespCode.Get()
	}
	if o.AvsMessage.IsSet() {
		toSerialize["avs_message"] = o.AvsMessage.Get()
	}
	if o.CvvCode.IsSet() {
		toSerialize["cvv_code"] = o.CvvCode.Get()
	}
	if o.CvvMessage.IsSet() {
		toSerialize["cvv_message"] = o.CvvMessage.Get()
	}
	if o.IsTestMode.IsSet() {
		toSerialize["is_test_mode"] = o.IsTestMode.Get()
	}
	if o.AdditionalFields != nil {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableOrderTransaction struct {
	value *OrderTransaction
	isSet bool
}

func (v NullableOrderTransaction) Get() *OrderTransaction {
	return v.value
}

func (v *NullableOrderTransaction) Set(val *OrderTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTransaction(val *OrderTransaction) *NullableOrderTransaction {
	return &NullableOrderTransaction{value: val, isSet: true}
}

func (v NullableOrderTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


