# OrderAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Defines order&#39;s id | [optional] 
**OrderId** | Pointer to **string** | Defines the order id if it is supported by the cart | [optional] 
**StoreId** | Pointer to **string** | Defines store id where the order should be assigned | [optional] 
**ChannelId** | Pointer to **string** | Channel ID | [optional] 
**OrderStatus** | **string** | Defines order status. | 
**FulfillmentStatus** | Pointer to **string** | Create order with fulfillment status | [optional] 
**FinancialStatus** | Pointer to **string** | Create order with financial status | [optional] 
**CustomerEmail** | **string** | Defines the customer specified by email for whom order has to be created | 
**CustomerFirstName** | Pointer to **string** | Specifies customer&#39;s first name | [optional] 
**CustomerLastName** | Pointer to **string** | Specifies customer’s last name | [optional] 
**CustomerPhone** | Pointer to **string** | Specifies customer’s phone | [optional] 
**CustomerCountry** | Pointer to **string** | Specifies customer&#39;s address ISO code or name of country | [optional] 
**CustomerBirthday** | Pointer to **string** | Specifies customer’s birthday | [optional] 
**CustomerFax** | Pointer to **string** | Specifies customer’s fax | [optional] 
**IsGuest** | Pointer to **bool** | Indicates whether the customer is a guest customer | [optional] [default to false]
**OrderPaymentMethod** | Pointer to **string** | Defines order payment method.&lt;br/&gt;Setting order_payment_method on Shopify will also change financial_status field value to &#39;paid&#39; | [optional] 
**TransactionId** | Pointer to **string** | Payment transaction id | [optional] 
**Currency** | Pointer to **string** | Currency code of order | [optional] 
**Date** | Pointer to **string** | Specifies an order creation date in format Y-m-d H:i:s | [optional] 
**DateModified** | Pointer to **string** | Specifies order&#39;s  modification date | [optional] 
**DateFinished** | Pointer to **string** | Specifies order&#39;s  finished date | [optional] 
**BillFirstName** | **string** | Specifies billing first name | 
**BillLastName** | **string** | Specifies billing last name | 
**BillAddress1** | **string** | Specifies first billing address | 
**BillAddress2** | Pointer to **string** | Specifies second billing address | [optional] 
**BillCity** | **string** | Specifies billing city | 
**BillPostcode** | **string** | Specifies billing postcode | 
**BillState** | **string** | Specifies billing state code | 
**BillCountry** | **string** | Specifies billing country code | 
**BillCompany** | Pointer to **string** | Specifies billing company | [optional] 
**BillPhone** | Pointer to **string** | Specifies billing phone | [optional] 
**BillFax** | Pointer to **string** | Specifies billing fax | [optional] 
**ShippFirstName** | Pointer to **string** | Specifies shipping first name | [optional] 
**ShippLastName** | Pointer to **string** | Specifies shipping last name | [optional] 
**ShippAddress1** | Pointer to **string** | Specifies first shipping address | [optional] 
**ShippAddress2** | Pointer to **string** | Specifies second address line of a shipping street address | [optional] 
**ShippCity** | Pointer to **string** | Specifies shipping city | [optional] 
**ShippPostcode** | Pointer to **string** | Specifies shipping postcode | [optional] 
**ShippState** | Pointer to **string** | Specifies shipping state code | [optional] 
**ShippCountry** | Pointer to **string** | Specifies shipping country code | [optional] 
**ShippCompany** | Pointer to **string** | Specifies shipping company | [optional] 
**ShippPhone** | Pointer to **string** | Specifies shipping phone | [optional] 
**ShippFax** | Pointer to **string** | Specifies shipping fax | [optional] 
**SubtotalPrice** | Pointer to **float32** | Total price of all ordered products multiplied by their number, excluding tax, shipping price and discounts | [optional] 
**TaxPrice** | Pointer to **float32** | The value of tax cost for order | [optional] [default to 0]
**TotalPrice** | Pointer to **float32** | Defines order&#39;s total price | [optional] 
**TotalPaid** | Pointer to **float32** | Defines total paid amount for the order | [optional] 
**TotalWeight** | Pointer to **int32** | Defines the sum of all line item weights in grams for the order | [optional] 
**PricesIncTax** | Pointer to **bool** | Indicates whether prices and subtotal includes tax. | [optional] [default to false]
**ShippingPrice** | Pointer to **float32** | Specifies order&#39;s shipping price | [optional] [default to 0]
**ShippingTax** | Pointer to **float32** | Specifies order&#39;s shipping price tax | [optional] 
**Discount** | Pointer to **float32** | Specifies order&#39;s discount | [optional] 
**CouponDiscount** | Pointer to **float32** | Specifies order&#39;s coupon discount | [optional] 
**GiftCertificateDiscount** | Pointer to **float32** | Discounts for order with gift certificates | [optional] 
**OrderShippingMethod** | Pointer to **string** | Defines order shipping method | [optional] 
**CarrierId** | Pointer to **string** | Defines tracking carrier id | [optional] 
**WarehouseId** | Pointer to **string** | This parameter is used for selecting a warehouse where you need to set/modify a product quantity. | [optional] 
**Coupons** | Pointer to **[]string** | Coupons that will be applied to order | [optional] 
**Tags** | Pointer to **string** | Order tags | [optional] 
**Comment** | Pointer to **string** | Specifies order comment | [optional] 
**AdminComment** | Pointer to **string** | Specifies admin&#39;s order comment | [optional] 
**AdminPrivateComment** | Pointer to **string** | Specifies private admin&#39;s order comment | [optional] 
**SendNotifications** | Pointer to **bool** | Send notifications to customer after order was created | [optional] [default to false]
**SendAdminNotifications** | Pointer to **bool** | Notify admin when new order was created. | [optional] [default to false]
**ExternalSource** | Pointer to **string** | Identifying the system used to generate the order | [optional] 
**InventoryBehaviour** | Pointer to **string** | The behaviour to use when updating inventory.&lt;hr&gt;&lt;div style&#x3D;\&quot;font-style:normal\&quot;&gt;Values description:&lt;div style&#x3D;\&quot;margin-left: 2%; padding-top: 2%\&quot;&gt;&lt;div style&#x3D;\&quot;font-size:85%\&quot;&gt;&lt;b&gt;bypass&lt;/b&gt; &#x3D; Do not claim inventory &lt;/br&gt;&lt;/br&gt;&lt;b&gt;decrement_ignoring_policy&lt;/b&gt; &#x3D; Ignore the product&#39;s &lt;/br&gt; inventory policy and claim amounts&lt;/br&gt;&lt;/br&gt;&lt;b&gt;decrement_obeying_policy&lt;/b&gt; &#x3D;  Obey the product&#39;s &lt;/br&gt; inventory policy.&lt;/br&gt;&lt;/br&gt;&lt;/div&gt;&lt;/div&gt;&lt;/div&gt; | [optional] [default to "bypass"]
**CreateInvoice** | Pointer to **bool** | Defines whether the invoice is created automatically along with the order | [optional] [default to false]
**NoteAttributes** | Pointer to [**[]OrderAddNoteAttributesInner**](OrderAddNoteAttributesInner.md) | Defines note attributes | [optional] 
**ClearCache** | Pointer to **bool** | Is cache clear required | [optional] [default to true]
**Origin** | Pointer to **string** | The source of the order | [optional] 
**FeePrice** | Pointer to **float32** | Specifies refund&#39;s fee price | [optional] 
**OrderItem** | [**[]OrderAddOrderItemInner**](OrderAddOrderItemInner.md) |  | 

## Methods

### NewOrderAdd

`func NewOrderAdd(orderStatus string, customerEmail string, billFirstName string, billLastName string, billAddress1 string, billCity string, billPostcode string, billState string, billCountry string, orderItem []OrderAddOrderItemInner, ) *OrderAdd`

NewOrderAdd instantiates a new OrderAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAddWithDefaults

`func NewOrderAddWithDefaults() *OrderAdd`

NewOrderAddWithDefaults instantiates a new OrderAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderAdd) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderAdd) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderAdd) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderAdd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderAdd) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderAdd) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderAdd) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderAdd) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStoreId

`func (o *OrderAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetChannelId

`func (o *OrderAdd) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *OrderAdd) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *OrderAdd) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *OrderAdd) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetOrderStatus

`func (o *OrderAdd) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderAdd) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderAdd) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.


### GetFulfillmentStatus

`func (o *OrderAdd) GetFulfillmentStatus() string`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *OrderAdd) GetFulfillmentStatusOk() (*string, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *OrderAdd) SetFulfillmentStatus(v string)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *OrderAdd) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.

### GetFinancialStatus

`func (o *OrderAdd) GetFinancialStatus() string`

GetFinancialStatus returns the FinancialStatus field if non-nil, zero value otherwise.

### GetFinancialStatusOk

`func (o *OrderAdd) GetFinancialStatusOk() (*string, bool)`

GetFinancialStatusOk returns a tuple with the FinancialStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialStatus

`func (o *OrderAdd) SetFinancialStatus(v string)`

SetFinancialStatus sets FinancialStatus field to given value.

### HasFinancialStatus

`func (o *OrderAdd) HasFinancialStatus() bool`

HasFinancialStatus returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *OrderAdd) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *OrderAdd) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *OrderAdd) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.


### GetCustomerFirstName

`func (o *OrderAdd) GetCustomerFirstName() string`

GetCustomerFirstName returns the CustomerFirstName field if non-nil, zero value otherwise.

### GetCustomerFirstNameOk

`func (o *OrderAdd) GetCustomerFirstNameOk() (*string, bool)`

GetCustomerFirstNameOk returns a tuple with the CustomerFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFirstName

`func (o *OrderAdd) SetCustomerFirstName(v string)`

SetCustomerFirstName sets CustomerFirstName field to given value.

### HasCustomerFirstName

`func (o *OrderAdd) HasCustomerFirstName() bool`

HasCustomerFirstName returns a boolean if a field has been set.

### GetCustomerLastName

`func (o *OrderAdd) GetCustomerLastName() string`

GetCustomerLastName returns the CustomerLastName field if non-nil, zero value otherwise.

### GetCustomerLastNameOk

`func (o *OrderAdd) GetCustomerLastNameOk() (*string, bool)`

GetCustomerLastNameOk returns a tuple with the CustomerLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLastName

`func (o *OrderAdd) SetCustomerLastName(v string)`

SetCustomerLastName sets CustomerLastName field to given value.

### HasCustomerLastName

`func (o *OrderAdd) HasCustomerLastName() bool`

HasCustomerLastName returns a boolean if a field has been set.

### GetCustomerPhone

`func (o *OrderAdd) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *OrderAdd) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *OrderAdd) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.

### HasCustomerPhone

`func (o *OrderAdd) HasCustomerPhone() bool`

HasCustomerPhone returns a boolean if a field has been set.

### GetCustomerCountry

`func (o *OrderAdd) GetCustomerCountry() string`

GetCustomerCountry returns the CustomerCountry field if non-nil, zero value otherwise.

### GetCustomerCountryOk

`func (o *OrderAdd) GetCustomerCountryOk() (*string, bool)`

GetCustomerCountryOk returns a tuple with the CustomerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCountry

`func (o *OrderAdd) SetCustomerCountry(v string)`

SetCustomerCountry sets CustomerCountry field to given value.

### HasCustomerCountry

`func (o *OrderAdd) HasCustomerCountry() bool`

HasCustomerCountry returns a boolean if a field has been set.

### GetCustomerBirthday

`func (o *OrderAdd) GetCustomerBirthday() string`

GetCustomerBirthday returns the CustomerBirthday field if non-nil, zero value otherwise.

### GetCustomerBirthdayOk

`func (o *OrderAdd) GetCustomerBirthdayOk() (*string, bool)`

GetCustomerBirthdayOk returns a tuple with the CustomerBirthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerBirthday

`func (o *OrderAdd) SetCustomerBirthday(v string)`

SetCustomerBirthday sets CustomerBirthday field to given value.

### HasCustomerBirthday

`func (o *OrderAdd) HasCustomerBirthday() bool`

HasCustomerBirthday returns a boolean if a field has been set.

### GetCustomerFax

`func (o *OrderAdd) GetCustomerFax() string`

GetCustomerFax returns the CustomerFax field if non-nil, zero value otherwise.

### GetCustomerFaxOk

`func (o *OrderAdd) GetCustomerFaxOk() (*string, bool)`

GetCustomerFaxOk returns a tuple with the CustomerFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFax

`func (o *OrderAdd) SetCustomerFax(v string)`

SetCustomerFax sets CustomerFax field to given value.

### HasCustomerFax

`func (o *OrderAdd) HasCustomerFax() bool`

HasCustomerFax returns a boolean if a field has been set.

### GetIsGuest

`func (o *OrderAdd) GetIsGuest() bool`

GetIsGuest returns the IsGuest field if non-nil, zero value otherwise.

### GetIsGuestOk

`func (o *OrderAdd) GetIsGuestOk() (*bool, bool)`

GetIsGuestOk returns a tuple with the IsGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGuest

`func (o *OrderAdd) SetIsGuest(v bool)`

SetIsGuest sets IsGuest field to given value.

### HasIsGuest

`func (o *OrderAdd) HasIsGuest() bool`

HasIsGuest returns a boolean if a field has been set.

### GetOrderPaymentMethod

`func (o *OrderAdd) GetOrderPaymentMethod() string`

GetOrderPaymentMethod returns the OrderPaymentMethod field if non-nil, zero value otherwise.

### GetOrderPaymentMethodOk

`func (o *OrderAdd) GetOrderPaymentMethodOk() (*string, bool)`

GetOrderPaymentMethodOk returns a tuple with the OrderPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderPaymentMethod

`func (o *OrderAdd) SetOrderPaymentMethod(v string)`

SetOrderPaymentMethod sets OrderPaymentMethod field to given value.

### HasOrderPaymentMethod

`func (o *OrderAdd) HasOrderPaymentMethod() bool`

HasOrderPaymentMethod returns a boolean if a field has been set.

### GetTransactionId

`func (o *OrderAdd) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *OrderAdd) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *OrderAdd) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *OrderAdd) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetCurrency

`func (o *OrderAdd) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderAdd) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderAdd) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderAdd) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDate

`func (o *OrderAdd) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *OrderAdd) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *OrderAdd) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *OrderAdd) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDateModified

`func (o *OrderAdd) GetDateModified() string`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *OrderAdd) GetDateModifiedOk() (*string, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *OrderAdd) SetDateModified(v string)`

SetDateModified sets DateModified field to given value.

### HasDateModified

`func (o *OrderAdd) HasDateModified() bool`

HasDateModified returns a boolean if a field has been set.

### GetDateFinished

`func (o *OrderAdd) GetDateFinished() string`

GetDateFinished returns the DateFinished field if non-nil, zero value otherwise.

### GetDateFinishedOk

`func (o *OrderAdd) GetDateFinishedOk() (*string, bool)`

GetDateFinishedOk returns a tuple with the DateFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFinished

`func (o *OrderAdd) SetDateFinished(v string)`

SetDateFinished sets DateFinished field to given value.

### HasDateFinished

`func (o *OrderAdd) HasDateFinished() bool`

HasDateFinished returns a boolean if a field has been set.

### GetBillFirstName

`func (o *OrderAdd) GetBillFirstName() string`

GetBillFirstName returns the BillFirstName field if non-nil, zero value otherwise.

### GetBillFirstNameOk

`func (o *OrderAdd) GetBillFirstNameOk() (*string, bool)`

GetBillFirstNameOk returns a tuple with the BillFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillFirstName

`func (o *OrderAdd) SetBillFirstName(v string)`

SetBillFirstName sets BillFirstName field to given value.


### GetBillLastName

`func (o *OrderAdd) GetBillLastName() string`

GetBillLastName returns the BillLastName field if non-nil, zero value otherwise.

### GetBillLastNameOk

`func (o *OrderAdd) GetBillLastNameOk() (*string, bool)`

GetBillLastNameOk returns a tuple with the BillLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillLastName

`func (o *OrderAdd) SetBillLastName(v string)`

SetBillLastName sets BillLastName field to given value.


### GetBillAddress1

`func (o *OrderAdd) GetBillAddress1() string`

GetBillAddress1 returns the BillAddress1 field if non-nil, zero value otherwise.

### GetBillAddress1Ok

`func (o *OrderAdd) GetBillAddress1Ok() (*string, bool)`

GetBillAddress1Ok returns a tuple with the BillAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillAddress1

`func (o *OrderAdd) SetBillAddress1(v string)`

SetBillAddress1 sets BillAddress1 field to given value.


### GetBillAddress2

`func (o *OrderAdd) GetBillAddress2() string`

GetBillAddress2 returns the BillAddress2 field if non-nil, zero value otherwise.

### GetBillAddress2Ok

`func (o *OrderAdd) GetBillAddress2Ok() (*string, bool)`

GetBillAddress2Ok returns a tuple with the BillAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillAddress2

`func (o *OrderAdd) SetBillAddress2(v string)`

SetBillAddress2 sets BillAddress2 field to given value.

### HasBillAddress2

`func (o *OrderAdd) HasBillAddress2() bool`

HasBillAddress2 returns a boolean if a field has been set.

### GetBillCity

`func (o *OrderAdd) GetBillCity() string`

GetBillCity returns the BillCity field if non-nil, zero value otherwise.

### GetBillCityOk

`func (o *OrderAdd) GetBillCityOk() (*string, bool)`

GetBillCityOk returns a tuple with the BillCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCity

`func (o *OrderAdd) SetBillCity(v string)`

SetBillCity sets BillCity field to given value.


### GetBillPostcode

`func (o *OrderAdd) GetBillPostcode() string`

GetBillPostcode returns the BillPostcode field if non-nil, zero value otherwise.

### GetBillPostcodeOk

`func (o *OrderAdd) GetBillPostcodeOk() (*string, bool)`

GetBillPostcodeOk returns a tuple with the BillPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillPostcode

`func (o *OrderAdd) SetBillPostcode(v string)`

SetBillPostcode sets BillPostcode field to given value.


### GetBillState

`func (o *OrderAdd) GetBillState() string`

GetBillState returns the BillState field if non-nil, zero value otherwise.

### GetBillStateOk

`func (o *OrderAdd) GetBillStateOk() (*string, bool)`

GetBillStateOk returns a tuple with the BillState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillState

`func (o *OrderAdd) SetBillState(v string)`

SetBillState sets BillState field to given value.


### GetBillCountry

`func (o *OrderAdd) GetBillCountry() string`

GetBillCountry returns the BillCountry field if non-nil, zero value otherwise.

### GetBillCountryOk

`func (o *OrderAdd) GetBillCountryOk() (*string, bool)`

GetBillCountryOk returns a tuple with the BillCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCountry

`func (o *OrderAdd) SetBillCountry(v string)`

SetBillCountry sets BillCountry field to given value.


### GetBillCompany

`func (o *OrderAdd) GetBillCompany() string`

GetBillCompany returns the BillCompany field if non-nil, zero value otherwise.

### GetBillCompanyOk

`func (o *OrderAdd) GetBillCompanyOk() (*string, bool)`

GetBillCompanyOk returns a tuple with the BillCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCompany

`func (o *OrderAdd) SetBillCompany(v string)`

SetBillCompany sets BillCompany field to given value.

### HasBillCompany

`func (o *OrderAdd) HasBillCompany() bool`

HasBillCompany returns a boolean if a field has been set.

### GetBillPhone

`func (o *OrderAdd) GetBillPhone() string`

GetBillPhone returns the BillPhone field if non-nil, zero value otherwise.

### GetBillPhoneOk

`func (o *OrderAdd) GetBillPhoneOk() (*string, bool)`

GetBillPhoneOk returns a tuple with the BillPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillPhone

`func (o *OrderAdd) SetBillPhone(v string)`

SetBillPhone sets BillPhone field to given value.

### HasBillPhone

`func (o *OrderAdd) HasBillPhone() bool`

HasBillPhone returns a boolean if a field has been set.

### GetBillFax

`func (o *OrderAdd) GetBillFax() string`

GetBillFax returns the BillFax field if non-nil, zero value otherwise.

### GetBillFaxOk

`func (o *OrderAdd) GetBillFaxOk() (*string, bool)`

GetBillFaxOk returns a tuple with the BillFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillFax

`func (o *OrderAdd) SetBillFax(v string)`

SetBillFax sets BillFax field to given value.

### HasBillFax

`func (o *OrderAdd) HasBillFax() bool`

HasBillFax returns a boolean if a field has been set.

### GetShippFirstName

`func (o *OrderAdd) GetShippFirstName() string`

GetShippFirstName returns the ShippFirstName field if non-nil, zero value otherwise.

### GetShippFirstNameOk

`func (o *OrderAdd) GetShippFirstNameOk() (*string, bool)`

GetShippFirstNameOk returns a tuple with the ShippFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippFirstName

`func (o *OrderAdd) SetShippFirstName(v string)`

SetShippFirstName sets ShippFirstName field to given value.

### HasShippFirstName

`func (o *OrderAdd) HasShippFirstName() bool`

HasShippFirstName returns a boolean if a field has been set.

### GetShippLastName

`func (o *OrderAdd) GetShippLastName() string`

GetShippLastName returns the ShippLastName field if non-nil, zero value otherwise.

### GetShippLastNameOk

`func (o *OrderAdd) GetShippLastNameOk() (*string, bool)`

GetShippLastNameOk returns a tuple with the ShippLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippLastName

`func (o *OrderAdd) SetShippLastName(v string)`

SetShippLastName sets ShippLastName field to given value.

### HasShippLastName

`func (o *OrderAdd) HasShippLastName() bool`

HasShippLastName returns a boolean if a field has been set.

### GetShippAddress1

`func (o *OrderAdd) GetShippAddress1() string`

GetShippAddress1 returns the ShippAddress1 field if non-nil, zero value otherwise.

### GetShippAddress1Ok

`func (o *OrderAdd) GetShippAddress1Ok() (*string, bool)`

GetShippAddress1Ok returns a tuple with the ShippAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippAddress1

`func (o *OrderAdd) SetShippAddress1(v string)`

SetShippAddress1 sets ShippAddress1 field to given value.

### HasShippAddress1

`func (o *OrderAdd) HasShippAddress1() bool`

HasShippAddress1 returns a boolean if a field has been set.

### GetShippAddress2

`func (o *OrderAdd) GetShippAddress2() string`

GetShippAddress2 returns the ShippAddress2 field if non-nil, zero value otherwise.

### GetShippAddress2Ok

`func (o *OrderAdd) GetShippAddress2Ok() (*string, bool)`

GetShippAddress2Ok returns a tuple with the ShippAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippAddress2

`func (o *OrderAdd) SetShippAddress2(v string)`

SetShippAddress2 sets ShippAddress2 field to given value.

### HasShippAddress2

`func (o *OrderAdd) HasShippAddress2() bool`

HasShippAddress2 returns a boolean if a field has been set.

### GetShippCity

`func (o *OrderAdd) GetShippCity() string`

GetShippCity returns the ShippCity field if non-nil, zero value otherwise.

### GetShippCityOk

`func (o *OrderAdd) GetShippCityOk() (*string, bool)`

GetShippCityOk returns a tuple with the ShippCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippCity

`func (o *OrderAdd) SetShippCity(v string)`

SetShippCity sets ShippCity field to given value.

### HasShippCity

`func (o *OrderAdd) HasShippCity() bool`

HasShippCity returns a boolean if a field has been set.

### GetShippPostcode

`func (o *OrderAdd) GetShippPostcode() string`

GetShippPostcode returns the ShippPostcode field if non-nil, zero value otherwise.

### GetShippPostcodeOk

`func (o *OrderAdd) GetShippPostcodeOk() (*string, bool)`

GetShippPostcodeOk returns a tuple with the ShippPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippPostcode

`func (o *OrderAdd) SetShippPostcode(v string)`

SetShippPostcode sets ShippPostcode field to given value.

### HasShippPostcode

`func (o *OrderAdd) HasShippPostcode() bool`

HasShippPostcode returns a boolean if a field has been set.

### GetShippState

`func (o *OrderAdd) GetShippState() string`

GetShippState returns the ShippState field if non-nil, zero value otherwise.

### GetShippStateOk

`func (o *OrderAdd) GetShippStateOk() (*string, bool)`

GetShippStateOk returns a tuple with the ShippState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippState

`func (o *OrderAdd) SetShippState(v string)`

SetShippState sets ShippState field to given value.

### HasShippState

`func (o *OrderAdd) HasShippState() bool`

HasShippState returns a boolean if a field has been set.

### GetShippCountry

`func (o *OrderAdd) GetShippCountry() string`

GetShippCountry returns the ShippCountry field if non-nil, zero value otherwise.

### GetShippCountryOk

`func (o *OrderAdd) GetShippCountryOk() (*string, bool)`

GetShippCountryOk returns a tuple with the ShippCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippCountry

`func (o *OrderAdd) SetShippCountry(v string)`

SetShippCountry sets ShippCountry field to given value.

### HasShippCountry

`func (o *OrderAdd) HasShippCountry() bool`

HasShippCountry returns a boolean if a field has been set.

### GetShippCompany

`func (o *OrderAdd) GetShippCompany() string`

GetShippCompany returns the ShippCompany field if non-nil, zero value otherwise.

### GetShippCompanyOk

`func (o *OrderAdd) GetShippCompanyOk() (*string, bool)`

GetShippCompanyOk returns a tuple with the ShippCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippCompany

`func (o *OrderAdd) SetShippCompany(v string)`

SetShippCompany sets ShippCompany field to given value.

### HasShippCompany

`func (o *OrderAdd) HasShippCompany() bool`

HasShippCompany returns a boolean if a field has been set.

### GetShippPhone

`func (o *OrderAdd) GetShippPhone() string`

GetShippPhone returns the ShippPhone field if non-nil, zero value otherwise.

### GetShippPhoneOk

`func (o *OrderAdd) GetShippPhoneOk() (*string, bool)`

GetShippPhoneOk returns a tuple with the ShippPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippPhone

`func (o *OrderAdd) SetShippPhone(v string)`

SetShippPhone sets ShippPhone field to given value.

### HasShippPhone

`func (o *OrderAdd) HasShippPhone() bool`

HasShippPhone returns a boolean if a field has been set.

### GetShippFax

`func (o *OrderAdd) GetShippFax() string`

GetShippFax returns the ShippFax field if non-nil, zero value otherwise.

### GetShippFaxOk

`func (o *OrderAdd) GetShippFaxOk() (*string, bool)`

GetShippFaxOk returns a tuple with the ShippFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippFax

`func (o *OrderAdd) SetShippFax(v string)`

SetShippFax sets ShippFax field to given value.

### HasShippFax

`func (o *OrderAdd) HasShippFax() bool`

HasShippFax returns a boolean if a field has been set.

### GetSubtotalPrice

`func (o *OrderAdd) GetSubtotalPrice() float32`

GetSubtotalPrice returns the SubtotalPrice field if non-nil, zero value otherwise.

### GetSubtotalPriceOk

`func (o *OrderAdd) GetSubtotalPriceOk() (*float32, bool)`

GetSubtotalPriceOk returns a tuple with the SubtotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalPrice

`func (o *OrderAdd) SetSubtotalPrice(v float32)`

SetSubtotalPrice sets SubtotalPrice field to given value.

### HasSubtotalPrice

`func (o *OrderAdd) HasSubtotalPrice() bool`

HasSubtotalPrice returns a boolean if a field has been set.

### GetTaxPrice

`func (o *OrderAdd) GetTaxPrice() float32`

GetTaxPrice returns the TaxPrice field if non-nil, zero value otherwise.

### GetTaxPriceOk

`func (o *OrderAdd) GetTaxPriceOk() (*float32, bool)`

GetTaxPriceOk returns a tuple with the TaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPrice

`func (o *OrderAdd) SetTaxPrice(v float32)`

SetTaxPrice sets TaxPrice field to given value.

### HasTaxPrice

`func (o *OrderAdd) HasTaxPrice() bool`

HasTaxPrice returns a boolean if a field has been set.

### GetTotalPrice

`func (o *OrderAdd) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *OrderAdd) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *OrderAdd) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *OrderAdd) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetTotalPaid

`func (o *OrderAdd) GetTotalPaid() float32`

GetTotalPaid returns the TotalPaid field if non-nil, zero value otherwise.

### GetTotalPaidOk

`func (o *OrderAdd) GetTotalPaidOk() (*float32, bool)`

GetTotalPaidOk returns a tuple with the TotalPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaid

`func (o *OrderAdd) SetTotalPaid(v float32)`

SetTotalPaid sets TotalPaid field to given value.

### HasTotalPaid

`func (o *OrderAdd) HasTotalPaid() bool`

HasTotalPaid returns a boolean if a field has been set.

### GetTotalWeight

`func (o *OrderAdd) GetTotalWeight() int32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *OrderAdd) GetTotalWeightOk() (*int32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *OrderAdd) SetTotalWeight(v int32)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *OrderAdd) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetPricesIncTax

`func (o *OrderAdd) GetPricesIncTax() bool`

GetPricesIncTax returns the PricesIncTax field if non-nil, zero value otherwise.

### GetPricesIncTaxOk

`func (o *OrderAdd) GetPricesIncTaxOk() (*bool, bool)`

GetPricesIncTaxOk returns a tuple with the PricesIncTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesIncTax

`func (o *OrderAdd) SetPricesIncTax(v bool)`

SetPricesIncTax sets PricesIncTax field to given value.

### HasPricesIncTax

`func (o *OrderAdd) HasPricesIncTax() bool`

HasPricesIncTax returns a boolean if a field has been set.

### GetShippingPrice

`func (o *OrderAdd) GetShippingPrice() float32`

GetShippingPrice returns the ShippingPrice field if non-nil, zero value otherwise.

### GetShippingPriceOk

`func (o *OrderAdd) GetShippingPriceOk() (*float32, bool)`

GetShippingPriceOk returns a tuple with the ShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPrice

`func (o *OrderAdd) SetShippingPrice(v float32)`

SetShippingPrice sets ShippingPrice field to given value.

### HasShippingPrice

`func (o *OrderAdd) HasShippingPrice() bool`

HasShippingPrice returns a boolean if a field has been set.

### GetShippingTax

`func (o *OrderAdd) GetShippingTax() float32`

GetShippingTax returns the ShippingTax field if non-nil, zero value otherwise.

### GetShippingTaxOk

`func (o *OrderAdd) GetShippingTaxOk() (*float32, bool)`

GetShippingTaxOk returns a tuple with the ShippingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTax

`func (o *OrderAdd) SetShippingTax(v float32)`

SetShippingTax sets ShippingTax field to given value.

### HasShippingTax

`func (o *OrderAdd) HasShippingTax() bool`

HasShippingTax returns a boolean if a field has been set.

### GetDiscount

`func (o *OrderAdd) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *OrderAdd) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *OrderAdd) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *OrderAdd) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetCouponDiscount

`func (o *OrderAdd) GetCouponDiscount() float32`

GetCouponDiscount returns the CouponDiscount field if non-nil, zero value otherwise.

### GetCouponDiscountOk

`func (o *OrderAdd) GetCouponDiscountOk() (*float32, bool)`

GetCouponDiscountOk returns a tuple with the CouponDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponDiscount

`func (o *OrderAdd) SetCouponDiscount(v float32)`

SetCouponDiscount sets CouponDiscount field to given value.

### HasCouponDiscount

`func (o *OrderAdd) HasCouponDiscount() bool`

HasCouponDiscount returns a boolean if a field has been set.

### GetGiftCertificateDiscount

`func (o *OrderAdd) GetGiftCertificateDiscount() float32`

GetGiftCertificateDiscount returns the GiftCertificateDiscount field if non-nil, zero value otherwise.

### GetGiftCertificateDiscountOk

`func (o *OrderAdd) GetGiftCertificateDiscountOk() (*float32, bool)`

GetGiftCertificateDiscountOk returns a tuple with the GiftCertificateDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCertificateDiscount

`func (o *OrderAdd) SetGiftCertificateDiscount(v float32)`

SetGiftCertificateDiscount sets GiftCertificateDiscount field to given value.

### HasGiftCertificateDiscount

`func (o *OrderAdd) HasGiftCertificateDiscount() bool`

HasGiftCertificateDiscount returns a boolean if a field has been set.

### GetOrderShippingMethod

`func (o *OrderAdd) GetOrderShippingMethod() string`

GetOrderShippingMethod returns the OrderShippingMethod field if non-nil, zero value otherwise.

### GetOrderShippingMethodOk

`func (o *OrderAdd) GetOrderShippingMethodOk() (*string, bool)`

GetOrderShippingMethodOk returns a tuple with the OrderShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderShippingMethod

`func (o *OrderAdd) SetOrderShippingMethod(v string)`

SetOrderShippingMethod sets OrderShippingMethod field to given value.

### HasOrderShippingMethod

`func (o *OrderAdd) HasOrderShippingMethod() bool`

HasOrderShippingMethod returns a boolean if a field has been set.

### GetCarrierId

`func (o *OrderAdd) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *OrderAdd) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *OrderAdd) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *OrderAdd) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *OrderAdd) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *OrderAdd) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *OrderAdd) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *OrderAdd) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetCoupons

`func (o *OrderAdd) GetCoupons() []string`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *OrderAdd) GetCouponsOk() (*[]string, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *OrderAdd) SetCoupons(v []string)`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *OrderAdd) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### GetTags

`func (o *OrderAdd) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrderAdd) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrderAdd) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrderAdd) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetComment

`func (o *OrderAdd) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *OrderAdd) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *OrderAdd) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *OrderAdd) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAdminComment

`func (o *OrderAdd) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *OrderAdd) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *OrderAdd) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.

### HasAdminComment

`func (o *OrderAdd) HasAdminComment() bool`

HasAdminComment returns a boolean if a field has been set.

### GetAdminPrivateComment

`func (o *OrderAdd) GetAdminPrivateComment() string`

GetAdminPrivateComment returns the AdminPrivateComment field if non-nil, zero value otherwise.

### GetAdminPrivateCommentOk

`func (o *OrderAdd) GetAdminPrivateCommentOk() (*string, bool)`

GetAdminPrivateCommentOk returns a tuple with the AdminPrivateComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPrivateComment

`func (o *OrderAdd) SetAdminPrivateComment(v string)`

SetAdminPrivateComment sets AdminPrivateComment field to given value.

### HasAdminPrivateComment

`func (o *OrderAdd) HasAdminPrivateComment() bool`

HasAdminPrivateComment returns a boolean if a field has been set.

### GetSendNotifications

`func (o *OrderAdd) GetSendNotifications() bool`

GetSendNotifications returns the SendNotifications field if non-nil, zero value otherwise.

### GetSendNotificationsOk

`func (o *OrderAdd) GetSendNotificationsOk() (*bool, bool)`

GetSendNotificationsOk returns a tuple with the SendNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotifications

`func (o *OrderAdd) SetSendNotifications(v bool)`

SetSendNotifications sets SendNotifications field to given value.

### HasSendNotifications

`func (o *OrderAdd) HasSendNotifications() bool`

HasSendNotifications returns a boolean if a field has been set.

### GetSendAdminNotifications

`func (o *OrderAdd) GetSendAdminNotifications() bool`

GetSendAdminNotifications returns the SendAdminNotifications field if non-nil, zero value otherwise.

### GetSendAdminNotificationsOk

`func (o *OrderAdd) GetSendAdminNotificationsOk() (*bool, bool)`

GetSendAdminNotificationsOk returns a tuple with the SendAdminNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAdminNotifications

`func (o *OrderAdd) SetSendAdminNotifications(v bool)`

SetSendAdminNotifications sets SendAdminNotifications field to given value.

### HasSendAdminNotifications

`func (o *OrderAdd) HasSendAdminNotifications() bool`

HasSendAdminNotifications returns a boolean if a field has been set.

### GetExternalSource

`func (o *OrderAdd) GetExternalSource() string`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *OrderAdd) GetExternalSourceOk() (*string, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSource

`func (o *OrderAdd) SetExternalSource(v string)`

SetExternalSource sets ExternalSource field to given value.

### HasExternalSource

`func (o *OrderAdd) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### GetInventoryBehaviour

`func (o *OrderAdd) GetInventoryBehaviour() string`

GetInventoryBehaviour returns the InventoryBehaviour field if non-nil, zero value otherwise.

### GetInventoryBehaviourOk

`func (o *OrderAdd) GetInventoryBehaviourOk() (*string, bool)`

GetInventoryBehaviourOk returns a tuple with the InventoryBehaviour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryBehaviour

`func (o *OrderAdd) SetInventoryBehaviour(v string)`

SetInventoryBehaviour sets InventoryBehaviour field to given value.

### HasInventoryBehaviour

`func (o *OrderAdd) HasInventoryBehaviour() bool`

HasInventoryBehaviour returns a boolean if a field has been set.

### GetCreateInvoice

`func (o *OrderAdd) GetCreateInvoice() bool`

GetCreateInvoice returns the CreateInvoice field if non-nil, zero value otherwise.

### GetCreateInvoiceOk

`func (o *OrderAdd) GetCreateInvoiceOk() (*bool, bool)`

GetCreateInvoiceOk returns a tuple with the CreateInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInvoice

`func (o *OrderAdd) SetCreateInvoice(v bool)`

SetCreateInvoice sets CreateInvoice field to given value.

### HasCreateInvoice

`func (o *OrderAdd) HasCreateInvoice() bool`

HasCreateInvoice returns a boolean if a field has been set.

### GetNoteAttributes

`func (o *OrderAdd) GetNoteAttributes() []OrderAddNoteAttributesInner`

GetNoteAttributes returns the NoteAttributes field if non-nil, zero value otherwise.

### GetNoteAttributesOk

`func (o *OrderAdd) GetNoteAttributesOk() (*[]OrderAddNoteAttributesInner, bool)`

GetNoteAttributesOk returns a tuple with the NoteAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteAttributes

`func (o *OrderAdd) SetNoteAttributes(v []OrderAddNoteAttributesInner)`

SetNoteAttributes sets NoteAttributes field to given value.

### HasNoteAttributes

`func (o *OrderAdd) HasNoteAttributes() bool`

HasNoteAttributes returns a boolean if a field has been set.

### GetClearCache

`func (o *OrderAdd) GetClearCache() bool`

GetClearCache returns the ClearCache field if non-nil, zero value otherwise.

### GetClearCacheOk

`func (o *OrderAdd) GetClearCacheOk() (*bool, bool)`

GetClearCacheOk returns a tuple with the ClearCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCache

`func (o *OrderAdd) SetClearCache(v bool)`

SetClearCache sets ClearCache field to given value.

### HasClearCache

`func (o *OrderAdd) HasClearCache() bool`

HasClearCache returns a boolean if a field has been set.

### GetOrigin

`func (o *OrderAdd) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *OrderAdd) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *OrderAdd) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *OrderAdd) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetFeePrice

`func (o *OrderAdd) GetFeePrice() float32`

GetFeePrice returns the FeePrice field if non-nil, zero value otherwise.

### GetFeePriceOk

`func (o *OrderAdd) GetFeePriceOk() (*float32, bool)`

GetFeePriceOk returns a tuple with the FeePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePrice

`func (o *OrderAdd) SetFeePrice(v float32)`

SetFeePrice sets FeePrice field to given value.

### HasFeePrice

`func (o *OrderAdd) HasFeePrice() bool`

HasFeePrice returns a boolean if a field has been set.

### GetOrderItem

`func (o *OrderAdd) GetOrderItem() []OrderAddOrderItemInner`

GetOrderItem returns the OrderItem field if non-nil, zero value otherwise.

### GetOrderItemOk

`func (o *OrderAdd) GetOrderItemOk() (*[]OrderAddOrderItemInner, bool)`

GetOrderItemOk returns a tuple with the OrderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItem

`func (o *OrderAdd) SetOrderItem(v []OrderAddOrderItemInner)`

SetOrderItem sets OrderItem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


