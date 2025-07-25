# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**BasketId** | Pointer to **NullableString** |  | [optional] 
**ChannelId** | Pointer to **NullableString** |  | [optional] 
**Customer** | Pointer to [**BaseCustomer**](BaseCustomer.md) |  | [optional] 
**CreateAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**Currency** | Pointer to [**NullableCurrency**](Currency.md) |  | [optional] 
**ShippingAddress** | Pointer to [**NullableCustomerAddress**](CustomerAddress.md) |  | [optional] 
**BillingAddress** | Pointer to [**NullableCustomerAddress**](CustomerAddress.md) |  | [optional] 
**PaymentMethod** | Pointer to [**NullableOrderPaymentMethod**](OrderPaymentMethod.md) |  | [optional] 
**ShippingMethod** | Pointer to [**NullableOrderShippingMethod**](OrderShippingMethod.md) |  | [optional] 
**ShippingMethods** | Pointer to [**[]OrderShippingMethod**](OrderShippingMethod.md) |  | [optional] 
**Status** | Pointer to [**OrderStatus**](OrderStatus.md) |  | [optional] 
**Totals** | Pointer to [**NullableOrderTotals**](OrderTotals.md) |  | [optional] 
**Total** | Pointer to [**NullableOrderTotal**](OrderTotal.md) |  | [optional] 
**Discounts** | Pointer to [**[]OrderTotalsNewDiscount**](OrderTotalsNewDiscount.md) |  | [optional] 
**OrderProducts** | Pointer to [**[]OrderItem**](OrderItem.md) |  | [optional] 
**Bundles** | Pointer to [**[]OrderItem**](OrderItem.md) |  | [optional] 
**ModifiedAt** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**FinishedTime** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**StoreId** | Pointer to **NullableString** |  | [optional] 
**WarehousesIds** | Pointer to **[]string** |  | [optional] 
**Refunds** | Pointer to [**[]OrderRefund**](OrderRefund.md) |  | [optional] 
**GiftMessage** | Pointer to **NullableString** |  | [optional] 
**OrderDetailsUrl** | Pointer to **NullableString** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Order) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *Order) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Order) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Order) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Order) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetBasketId

`func (o *Order) GetBasketId() string`

GetBasketId returns the BasketId field if non-nil, zero value otherwise.

### GetBasketIdOk

`func (o *Order) GetBasketIdOk() (*string, bool)`

GetBasketIdOk returns a tuple with the BasketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasketId

`func (o *Order) SetBasketId(v string)`

SetBasketId sets BasketId field to given value.

### HasBasketId

`func (o *Order) HasBasketId() bool`

HasBasketId returns a boolean if a field has been set.

### SetBasketIdNil

`func (o *Order) SetBasketIdNil(b bool)`

 SetBasketIdNil sets the value for BasketId to be an explicit nil

### UnsetBasketId
`func (o *Order) UnsetBasketId()`

UnsetBasketId ensures that no value is present for BasketId, not even an explicit nil
### GetChannelId

`func (o *Order) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *Order) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *Order) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *Order) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *Order) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *Order) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetCustomer

`func (o *Order) GetCustomer() BaseCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Order) GetCustomerOk() (*BaseCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Order) SetCustomer(v BaseCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Order) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCreateAt

`func (o *Order) GetCreateAt() A2CDateTime`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *Order) GetCreateAtOk() (*A2CDateTime, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *Order) SetCreateAt(v A2CDateTime)`

SetCreateAt sets CreateAt field to given value.

### HasCreateAt

`func (o *Order) HasCreateAt() bool`

HasCreateAt returns a boolean if a field has been set.

### GetCurrency

`func (o *Order) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Order) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Order) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Order) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *Order) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *Order) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetShippingAddress

`func (o *Order) GetShippingAddress() CustomerAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *Order) GetShippingAddressOk() (*CustomerAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *Order) SetShippingAddress(v CustomerAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *Order) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### SetShippingAddressNil

`func (o *Order) SetShippingAddressNil(b bool)`

 SetShippingAddressNil sets the value for ShippingAddress to be an explicit nil

### UnsetShippingAddress
`func (o *Order) UnsetShippingAddress()`

UnsetShippingAddress ensures that no value is present for ShippingAddress, not even an explicit nil
### GetBillingAddress

`func (o *Order) GetBillingAddress() CustomerAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *Order) GetBillingAddressOk() (*CustomerAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *Order) SetBillingAddress(v CustomerAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *Order) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### SetBillingAddressNil

`func (o *Order) SetBillingAddressNil(b bool)`

 SetBillingAddressNil sets the value for BillingAddress to be an explicit nil

### UnsetBillingAddress
`func (o *Order) UnsetBillingAddress()`

UnsetBillingAddress ensures that no value is present for BillingAddress, not even an explicit nil
### GetPaymentMethod

`func (o *Order) GetPaymentMethod() OrderPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *Order) GetPaymentMethodOk() (*OrderPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *Order) SetPaymentMethod(v OrderPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *Order) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *Order) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *Order) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetShippingMethod

`func (o *Order) GetShippingMethod() OrderShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *Order) GetShippingMethodOk() (*OrderShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *Order) SetShippingMethod(v OrderShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *Order) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### SetShippingMethodNil

`func (o *Order) SetShippingMethodNil(b bool)`

 SetShippingMethodNil sets the value for ShippingMethod to be an explicit nil

### UnsetShippingMethod
`func (o *Order) UnsetShippingMethod()`

UnsetShippingMethod ensures that no value is present for ShippingMethod, not even an explicit nil
### GetShippingMethods

`func (o *Order) GetShippingMethods() []OrderShippingMethod`

GetShippingMethods returns the ShippingMethods field if non-nil, zero value otherwise.

### GetShippingMethodsOk

`func (o *Order) GetShippingMethodsOk() (*[]OrderShippingMethod, bool)`

GetShippingMethodsOk returns a tuple with the ShippingMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethods

`func (o *Order) SetShippingMethods(v []OrderShippingMethod)`

SetShippingMethods sets ShippingMethods field to given value.

### HasShippingMethods

`func (o *Order) HasShippingMethods() bool`

HasShippingMethods returns a boolean if a field has been set.

### GetStatus

`func (o *Order) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Order) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotals

`func (o *Order) GetTotals() OrderTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *Order) GetTotalsOk() (*OrderTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *Order) SetTotals(v OrderTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *Order) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### SetTotalsNil

`func (o *Order) SetTotalsNil(b bool)`

 SetTotalsNil sets the value for Totals to be an explicit nil

### UnsetTotals
`func (o *Order) UnsetTotals()`

UnsetTotals ensures that no value is present for Totals, not even an explicit nil
### GetTotal

`func (o *Order) GetTotal() OrderTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Order) GetTotalOk() (*OrderTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Order) SetTotal(v OrderTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Order) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *Order) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *Order) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetDiscounts

`func (o *Order) GetDiscounts() []OrderTotalsNewDiscount`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *Order) GetDiscountsOk() (*[]OrderTotalsNewDiscount, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *Order) SetDiscounts(v []OrderTotalsNewDiscount)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *Order) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetOrderProducts

`func (o *Order) GetOrderProducts() []OrderItem`

GetOrderProducts returns the OrderProducts field if non-nil, zero value otherwise.

### GetOrderProductsOk

`func (o *Order) GetOrderProductsOk() (*[]OrderItem, bool)`

GetOrderProductsOk returns a tuple with the OrderProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProducts

`func (o *Order) SetOrderProducts(v []OrderItem)`

SetOrderProducts sets OrderProducts field to given value.

### HasOrderProducts

`func (o *Order) HasOrderProducts() bool`

HasOrderProducts returns a boolean if a field has been set.

### GetBundles

`func (o *Order) GetBundles() []OrderItem`

GetBundles returns the Bundles field if non-nil, zero value otherwise.

### GetBundlesOk

`func (o *Order) GetBundlesOk() (*[]OrderItem, bool)`

GetBundlesOk returns a tuple with the Bundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundles

`func (o *Order) SetBundles(v []OrderItem)`

SetBundles sets Bundles field to given value.

### HasBundles

`func (o *Order) HasBundles() bool`

HasBundles returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Order) GetModifiedAt() A2CDateTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Order) GetModifiedAtOk() (*A2CDateTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Order) SetModifiedAt(v A2CDateTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Order) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *Order) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *Order) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetFinishedTime

`func (o *Order) GetFinishedTime() A2CDateTime`

GetFinishedTime returns the FinishedTime field if non-nil, zero value otherwise.

### GetFinishedTimeOk

`func (o *Order) GetFinishedTimeOk() (*A2CDateTime, bool)`

GetFinishedTimeOk returns a tuple with the FinishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedTime

`func (o *Order) SetFinishedTime(v A2CDateTime)`

SetFinishedTime sets FinishedTime field to given value.

### HasFinishedTime

`func (o *Order) HasFinishedTime() bool`

HasFinishedTime returns a boolean if a field has been set.

### SetFinishedTimeNil

`func (o *Order) SetFinishedTimeNil(b bool)`

 SetFinishedTimeNil sets the value for FinishedTime to be an explicit nil

### UnsetFinishedTime
`func (o *Order) UnsetFinishedTime()`

UnsetFinishedTime ensures that no value is present for FinishedTime, not even an explicit nil
### GetComment

`func (o *Order) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Order) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Order) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Order) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Order) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Order) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetStoreId

`func (o *Order) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *Order) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *Order) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *Order) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### SetStoreIdNil

`func (o *Order) SetStoreIdNil(b bool)`

 SetStoreIdNil sets the value for StoreId to be an explicit nil

### UnsetStoreId
`func (o *Order) UnsetStoreId()`

UnsetStoreId ensures that no value is present for StoreId, not even an explicit nil
### GetWarehousesIds

`func (o *Order) GetWarehousesIds() []string`

GetWarehousesIds returns the WarehousesIds field if non-nil, zero value otherwise.

### GetWarehousesIdsOk

`func (o *Order) GetWarehousesIdsOk() (*[]string, bool)`

GetWarehousesIdsOk returns a tuple with the WarehousesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehousesIds

`func (o *Order) SetWarehousesIds(v []string)`

SetWarehousesIds sets WarehousesIds field to given value.

### HasWarehousesIds

`func (o *Order) HasWarehousesIds() bool`

HasWarehousesIds returns a boolean if a field has been set.

### GetRefunds

`func (o *Order) GetRefunds() []OrderRefund`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *Order) GetRefundsOk() (*[]OrderRefund, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *Order) SetRefunds(v []OrderRefund)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *Order) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetGiftMessage

`func (o *Order) GetGiftMessage() string`

GetGiftMessage returns the GiftMessage field if non-nil, zero value otherwise.

### GetGiftMessageOk

`func (o *Order) GetGiftMessageOk() (*string, bool)`

GetGiftMessageOk returns a tuple with the GiftMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftMessage

`func (o *Order) SetGiftMessage(v string)`

SetGiftMessage sets GiftMessage field to given value.

### HasGiftMessage

`func (o *Order) HasGiftMessage() bool`

HasGiftMessage returns a boolean if a field has been set.

### SetGiftMessageNil

`func (o *Order) SetGiftMessageNil(b bool)`

 SetGiftMessageNil sets the value for GiftMessage to be an explicit nil

### UnsetGiftMessage
`func (o *Order) UnsetGiftMessage()`

UnsetGiftMessage ensures that no value is present for GiftMessage, not even an explicit nil
### GetOrderDetailsUrl

`func (o *Order) GetOrderDetailsUrl() string`

GetOrderDetailsUrl returns the OrderDetailsUrl field if non-nil, zero value otherwise.

### GetOrderDetailsUrlOk

`func (o *Order) GetOrderDetailsUrlOk() (*string, bool)`

GetOrderDetailsUrlOk returns a tuple with the OrderDetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDetailsUrl

`func (o *Order) SetOrderDetailsUrl(v string)`

SetOrderDetailsUrl sets OrderDetailsUrl field to given value.

### HasOrderDetailsUrl

`func (o *Order) HasOrderDetailsUrl() bool`

HasOrderDetailsUrl returns a boolean if a field has been set.

### SetOrderDetailsUrlNil

`func (o *Order) SetOrderDetailsUrlNil(b bool)`

 SetOrderDetailsUrlNil sets the value for OrderDetailsUrl to be an explicit nil

### UnsetOrderDetailsUrl
`func (o *Order) UnsetOrderDetailsUrl()`

UnsetOrderDetailsUrl ensures that no value is present for OrderDetailsUrl, not even an explicit nil
### GetAdditionalFields

`func (o *Order) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Order) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Order) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Order) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *Order) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *Order) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *Order) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Order) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Order) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Order) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *Order) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *Order) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


