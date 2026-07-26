# OrderShipmentAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Defines the order for which the shipment will be created | [optional] 
**WarehouseId** | Pointer to **string** | This parameter is used for selecting a warehouse where you need to set/modify a product quantity. | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**ShipmentProvider** | Pointer to **string** | Defines company name that provide tracking of shipment | [optional] 
**ShippingMethod** | Pointer to **string** | Define shipping method | [optional] 
**Items** | Pointer to [**[]OrderShipmentAddItemsInner**](OrderShipmentAddItemsInner.md) | Defines items in the order that will be shipped | [optional] 
**TrackingNumbers** | Pointer to [**[]OrderShipmentAddTrackingNumbersInner**](OrderShipmentAddTrackingNumbersInner.md) | Defines shipment&#39;s tracking numbers that have to be added&lt;/br&gt; How set tracking numbers to appropriate carrier:&lt;ul&gt;&lt;li&gt;tracking_numbers[]&#x3D;a2c.demo1,a2c.demo2 - set default carrier&lt;/li&gt;&lt;li&gt;tracking_numbers[&lt;b&gt;carrier_id&lt;/b&gt;]&#x3D;a2c.demo - set appropriate carrier&lt;/li&gt;&lt;/ul&gt;To get the list of carriers IDs that are available in your store, use the &lt;a href &#x3D; \&quot;https://api2cart.com/docs/#/cart/CartInfo\&quot;&gt;cart.info&lt;/a &gt; method | [optional] 
**TrackingLink** | Pointer to **string** | Defines custom tracking link | [optional] 
**IsShipped** | Pointer to **bool** | Defines shipment&#39;s status | [optional] [default to true]
**SendNotifications** | Pointer to **bool** | Send notifications to customer after shipment was created | [optional] [default to false]
**AdjustStock** | Pointer to **bool** | This parameter is used for adjust stock. | [optional] [default to false]
**CheckProcessStatus** | Pointer to **bool** | Disable or enable check process status. Please note that the response will be slower due to additional requests to the store. | [optional] [default to false]
**TrackingProvider** | Pointer to **string** | Defines name of the company which provides shipment tracking | [optional] 
**AdminComment** | Pointer to **string** | Specifies admin&#39;s order comment | [optional] 
**MailClass** | Pointer to **string** | Mail class for the shipment (e.g., priority, express). | [optional] 
**ShipDate** | Pointer to **string** | Ship date. | [optional] 
**Weight** | Pointer to **float32** | Weight | [optional] 
**WeightUnit** | Pointer to **string** | Weight Unit | [optional] 
**Length** | Pointer to **float32** | Defines product&#39;s length | [optional] 
**Width** | Pointer to **float32** | Defines product&#39;s width | [optional] 
**Height** | Pointer to **float32** | Defines product&#39;s height | [optional] 
**DimensionsUnit** | Pointer to **string** | Weight Unit | [optional] 
**ShippingLabelCost** | Pointer to **float32** | Cost of the shipping label. | [optional] 
**ShippingLabelCurrency** | Pointer to **string** | Currency code for the shipping label cost (3-letter ISO code). | [optional] 
**RevenueEligibility** | Pointer to **bool** | Revenue eligibility flag. | [optional] 
**ShipFromCountry** | Pointer to **string** | Country code the shipment is sent from (2-letter ISO code). | [optional] 
**ShipToCountry** | Pointer to **string** | Country code the shipment is sent to (2-letter ISO code). | [optional] 
**Incoterm** | Pointer to **string** | International commercial term for the shipment (e.g., DAP, DDP). | [optional] 
**DutyAmount** | Pointer to **float32** | Duty amount for international shipment. | [optional] 
**DutyCurrency** | Pointer to **string** | Currency code for the duty amount (3-letter ISO code). | [optional] 
**EnableCache** | Pointer to **bool** | If the value is &#39;true&#39; and order exist in our cache, we will use order.info from cache to prepare shipment items. | [optional] [default to false]
**UseLatestApiVersion** | Pointer to **bool** | Use the latest platform API version | [optional] [default to false]
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 

## Methods

### NewOrderShipmentAdd

`func NewOrderShipmentAdd() *OrderShipmentAdd`

NewOrderShipmentAdd instantiates a new OrderShipmentAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderShipmentAddWithDefaults

`func NewOrderShipmentAddWithDefaults() *OrderShipmentAdd`

NewOrderShipmentAddWithDefaults instantiates a new OrderShipmentAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *OrderShipmentAdd) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderShipmentAdd) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderShipmentAdd) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderShipmentAdd) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *OrderShipmentAdd) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *OrderShipmentAdd) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *OrderShipmentAdd) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *OrderShipmentAdd) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetStoreId

`func (o *OrderShipmentAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderShipmentAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderShipmentAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderShipmentAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetShipmentProvider

`func (o *OrderShipmentAdd) GetShipmentProvider() string`

GetShipmentProvider returns the ShipmentProvider field if non-nil, zero value otherwise.

### GetShipmentProviderOk

`func (o *OrderShipmentAdd) GetShipmentProviderOk() (*string, bool)`

GetShipmentProviderOk returns a tuple with the ShipmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentProvider

`func (o *OrderShipmentAdd) SetShipmentProvider(v string)`

SetShipmentProvider sets ShipmentProvider field to given value.

### HasShipmentProvider

`func (o *OrderShipmentAdd) HasShipmentProvider() bool`

HasShipmentProvider returns a boolean if a field has been set.

### GetShippingMethod

`func (o *OrderShipmentAdd) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *OrderShipmentAdd) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *OrderShipmentAdd) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *OrderShipmentAdd) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetItems

`func (o *OrderShipmentAdd) GetItems() []OrderShipmentAddItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrderShipmentAdd) GetItemsOk() (*[]OrderShipmentAddItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrderShipmentAdd) SetItems(v []OrderShipmentAddItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *OrderShipmentAdd) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTrackingNumbers

`func (o *OrderShipmentAdd) GetTrackingNumbers() []OrderShipmentAddTrackingNumbersInner`

GetTrackingNumbers returns the TrackingNumbers field if non-nil, zero value otherwise.

### GetTrackingNumbersOk

`func (o *OrderShipmentAdd) GetTrackingNumbersOk() (*[]OrderShipmentAddTrackingNumbersInner, bool)`

GetTrackingNumbersOk returns a tuple with the TrackingNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumbers

`func (o *OrderShipmentAdd) SetTrackingNumbers(v []OrderShipmentAddTrackingNumbersInner)`

SetTrackingNumbers sets TrackingNumbers field to given value.

### HasTrackingNumbers

`func (o *OrderShipmentAdd) HasTrackingNumbers() bool`

HasTrackingNumbers returns a boolean if a field has been set.

### GetTrackingLink

`func (o *OrderShipmentAdd) GetTrackingLink() string`

GetTrackingLink returns the TrackingLink field if non-nil, zero value otherwise.

### GetTrackingLinkOk

`func (o *OrderShipmentAdd) GetTrackingLinkOk() (*string, bool)`

GetTrackingLinkOk returns a tuple with the TrackingLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingLink

`func (o *OrderShipmentAdd) SetTrackingLink(v string)`

SetTrackingLink sets TrackingLink field to given value.

### HasTrackingLink

`func (o *OrderShipmentAdd) HasTrackingLink() bool`

HasTrackingLink returns a boolean if a field has been set.

### GetIsShipped

`func (o *OrderShipmentAdd) GetIsShipped() bool`

GetIsShipped returns the IsShipped field if non-nil, zero value otherwise.

### GetIsShippedOk

`func (o *OrderShipmentAdd) GetIsShippedOk() (*bool, bool)`

GetIsShippedOk returns a tuple with the IsShipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShipped

`func (o *OrderShipmentAdd) SetIsShipped(v bool)`

SetIsShipped sets IsShipped field to given value.

### HasIsShipped

`func (o *OrderShipmentAdd) HasIsShipped() bool`

HasIsShipped returns a boolean if a field has been set.

### GetSendNotifications

`func (o *OrderShipmentAdd) GetSendNotifications() bool`

GetSendNotifications returns the SendNotifications field if non-nil, zero value otherwise.

### GetSendNotificationsOk

`func (o *OrderShipmentAdd) GetSendNotificationsOk() (*bool, bool)`

GetSendNotificationsOk returns a tuple with the SendNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotifications

`func (o *OrderShipmentAdd) SetSendNotifications(v bool)`

SetSendNotifications sets SendNotifications field to given value.

### HasSendNotifications

`func (o *OrderShipmentAdd) HasSendNotifications() bool`

HasSendNotifications returns a boolean if a field has been set.

### GetAdjustStock

`func (o *OrderShipmentAdd) GetAdjustStock() bool`

GetAdjustStock returns the AdjustStock field if non-nil, zero value otherwise.

### GetAdjustStockOk

`func (o *OrderShipmentAdd) GetAdjustStockOk() (*bool, bool)`

GetAdjustStockOk returns a tuple with the AdjustStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustStock

`func (o *OrderShipmentAdd) SetAdjustStock(v bool)`

SetAdjustStock sets AdjustStock field to given value.

### HasAdjustStock

`func (o *OrderShipmentAdd) HasAdjustStock() bool`

HasAdjustStock returns a boolean if a field has been set.

### GetCheckProcessStatus

`func (o *OrderShipmentAdd) GetCheckProcessStatus() bool`

GetCheckProcessStatus returns the CheckProcessStatus field if non-nil, zero value otherwise.

### GetCheckProcessStatusOk

`func (o *OrderShipmentAdd) GetCheckProcessStatusOk() (*bool, bool)`

GetCheckProcessStatusOk returns a tuple with the CheckProcessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckProcessStatus

`func (o *OrderShipmentAdd) SetCheckProcessStatus(v bool)`

SetCheckProcessStatus sets CheckProcessStatus field to given value.

### HasCheckProcessStatus

`func (o *OrderShipmentAdd) HasCheckProcessStatus() bool`

HasCheckProcessStatus returns a boolean if a field has been set.

### GetTrackingProvider

`func (o *OrderShipmentAdd) GetTrackingProvider() string`

GetTrackingProvider returns the TrackingProvider field if non-nil, zero value otherwise.

### GetTrackingProviderOk

`func (o *OrderShipmentAdd) GetTrackingProviderOk() (*string, bool)`

GetTrackingProviderOk returns a tuple with the TrackingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingProvider

`func (o *OrderShipmentAdd) SetTrackingProvider(v string)`

SetTrackingProvider sets TrackingProvider field to given value.

### HasTrackingProvider

`func (o *OrderShipmentAdd) HasTrackingProvider() bool`

HasTrackingProvider returns a boolean if a field has been set.

### GetAdminComment

`func (o *OrderShipmentAdd) GetAdminComment() string`

GetAdminComment returns the AdminComment field if non-nil, zero value otherwise.

### GetAdminCommentOk

`func (o *OrderShipmentAdd) GetAdminCommentOk() (*string, bool)`

GetAdminCommentOk returns a tuple with the AdminComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminComment

`func (o *OrderShipmentAdd) SetAdminComment(v string)`

SetAdminComment sets AdminComment field to given value.

### HasAdminComment

`func (o *OrderShipmentAdd) HasAdminComment() bool`

HasAdminComment returns a boolean if a field has been set.

### GetMailClass

`func (o *OrderShipmentAdd) GetMailClass() string`

GetMailClass returns the MailClass field if non-nil, zero value otherwise.

### GetMailClassOk

`func (o *OrderShipmentAdd) GetMailClassOk() (*string, bool)`

GetMailClassOk returns a tuple with the MailClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailClass

`func (o *OrderShipmentAdd) SetMailClass(v string)`

SetMailClass sets MailClass field to given value.

### HasMailClass

`func (o *OrderShipmentAdd) HasMailClass() bool`

HasMailClass returns a boolean if a field has been set.

### GetShipDate

`func (o *OrderShipmentAdd) GetShipDate() string`

GetShipDate returns the ShipDate field if non-nil, zero value otherwise.

### GetShipDateOk

`func (o *OrderShipmentAdd) GetShipDateOk() (*string, bool)`

GetShipDateOk returns a tuple with the ShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipDate

`func (o *OrderShipmentAdd) SetShipDate(v string)`

SetShipDate sets ShipDate field to given value.

### HasShipDate

`func (o *OrderShipmentAdd) HasShipDate() bool`

HasShipDate returns a boolean if a field has been set.

### GetWeight

`func (o *OrderShipmentAdd) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OrderShipmentAdd) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OrderShipmentAdd) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *OrderShipmentAdd) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *OrderShipmentAdd) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *OrderShipmentAdd) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *OrderShipmentAdd) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *OrderShipmentAdd) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetLength

`func (o *OrderShipmentAdd) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *OrderShipmentAdd) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *OrderShipmentAdd) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *OrderShipmentAdd) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWidth

`func (o *OrderShipmentAdd) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *OrderShipmentAdd) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *OrderShipmentAdd) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *OrderShipmentAdd) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *OrderShipmentAdd) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *OrderShipmentAdd) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *OrderShipmentAdd) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *OrderShipmentAdd) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetDimensionsUnit

`func (o *OrderShipmentAdd) GetDimensionsUnit() string`

GetDimensionsUnit returns the DimensionsUnit field if non-nil, zero value otherwise.

### GetDimensionsUnitOk

`func (o *OrderShipmentAdd) GetDimensionsUnitOk() (*string, bool)`

GetDimensionsUnitOk returns a tuple with the DimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsUnit

`func (o *OrderShipmentAdd) SetDimensionsUnit(v string)`

SetDimensionsUnit sets DimensionsUnit field to given value.

### HasDimensionsUnit

`func (o *OrderShipmentAdd) HasDimensionsUnit() bool`

HasDimensionsUnit returns a boolean if a field has been set.

### GetShippingLabelCost

`func (o *OrderShipmentAdd) GetShippingLabelCost() float32`

GetShippingLabelCost returns the ShippingLabelCost field if non-nil, zero value otherwise.

### GetShippingLabelCostOk

`func (o *OrderShipmentAdd) GetShippingLabelCostOk() (*float32, bool)`

GetShippingLabelCostOk returns a tuple with the ShippingLabelCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelCost

`func (o *OrderShipmentAdd) SetShippingLabelCost(v float32)`

SetShippingLabelCost sets ShippingLabelCost field to given value.

### HasShippingLabelCost

`func (o *OrderShipmentAdd) HasShippingLabelCost() bool`

HasShippingLabelCost returns a boolean if a field has been set.

### GetShippingLabelCurrency

`func (o *OrderShipmentAdd) GetShippingLabelCurrency() string`

GetShippingLabelCurrency returns the ShippingLabelCurrency field if non-nil, zero value otherwise.

### GetShippingLabelCurrencyOk

`func (o *OrderShipmentAdd) GetShippingLabelCurrencyOk() (*string, bool)`

GetShippingLabelCurrencyOk returns a tuple with the ShippingLabelCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelCurrency

`func (o *OrderShipmentAdd) SetShippingLabelCurrency(v string)`

SetShippingLabelCurrency sets ShippingLabelCurrency field to given value.

### HasShippingLabelCurrency

`func (o *OrderShipmentAdd) HasShippingLabelCurrency() bool`

HasShippingLabelCurrency returns a boolean if a field has been set.

### GetRevenueEligibility

`func (o *OrderShipmentAdd) GetRevenueEligibility() bool`

GetRevenueEligibility returns the RevenueEligibility field if non-nil, zero value otherwise.

### GetRevenueEligibilityOk

`func (o *OrderShipmentAdd) GetRevenueEligibilityOk() (*bool, bool)`

GetRevenueEligibilityOk returns a tuple with the RevenueEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueEligibility

`func (o *OrderShipmentAdd) SetRevenueEligibility(v bool)`

SetRevenueEligibility sets RevenueEligibility field to given value.

### HasRevenueEligibility

`func (o *OrderShipmentAdd) HasRevenueEligibility() bool`

HasRevenueEligibility returns a boolean if a field has been set.

### GetShipFromCountry

`func (o *OrderShipmentAdd) GetShipFromCountry() string`

GetShipFromCountry returns the ShipFromCountry field if non-nil, zero value otherwise.

### GetShipFromCountryOk

`func (o *OrderShipmentAdd) GetShipFromCountryOk() (*string, bool)`

GetShipFromCountryOk returns a tuple with the ShipFromCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromCountry

`func (o *OrderShipmentAdd) SetShipFromCountry(v string)`

SetShipFromCountry sets ShipFromCountry field to given value.

### HasShipFromCountry

`func (o *OrderShipmentAdd) HasShipFromCountry() bool`

HasShipFromCountry returns a boolean if a field has been set.

### GetShipToCountry

`func (o *OrderShipmentAdd) GetShipToCountry() string`

GetShipToCountry returns the ShipToCountry field if non-nil, zero value otherwise.

### GetShipToCountryOk

`func (o *OrderShipmentAdd) GetShipToCountryOk() (*string, bool)`

GetShipToCountryOk returns a tuple with the ShipToCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCountry

`func (o *OrderShipmentAdd) SetShipToCountry(v string)`

SetShipToCountry sets ShipToCountry field to given value.

### HasShipToCountry

`func (o *OrderShipmentAdd) HasShipToCountry() bool`

HasShipToCountry returns a boolean if a field has been set.

### GetIncoterm

`func (o *OrderShipmentAdd) GetIncoterm() string`

GetIncoterm returns the Incoterm field if non-nil, zero value otherwise.

### GetIncotermOk

`func (o *OrderShipmentAdd) GetIncotermOk() (*string, bool)`

GetIncotermOk returns a tuple with the Incoterm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncoterm

`func (o *OrderShipmentAdd) SetIncoterm(v string)`

SetIncoterm sets Incoterm field to given value.

### HasIncoterm

`func (o *OrderShipmentAdd) HasIncoterm() bool`

HasIncoterm returns a boolean if a field has been set.

### GetDutyAmount

`func (o *OrderShipmentAdd) GetDutyAmount() float32`

GetDutyAmount returns the DutyAmount field if non-nil, zero value otherwise.

### GetDutyAmountOk

`func (o *OrderShipmentAdd) GetDutyAmountOk() (*float32, bool)`

GetDutyAmountOk returns a tuple with the DutyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyAmount

`func (o *OrderShipmentAdd) SetDutyAmount(v float32)`

SetDutyAmount sets DutyAmount field to given value.

### HasDutyAmount

`func (o *OrderShipmentAdd) HasDutyAmount() bool`

HasDutyAmount returns a boolean if a field has been set.

### GetDutyCurrency

`func (o *OrderShipmentAdd) GetDutyCurrency() string`

GetDutyCurrency returns the DutyCurrency field if non-nil, zero value otherwise.

### GetDutyCurrencyOk

`func (o *OrderShipmentAdd) GetDutyCurrencyOk() (*string, bool)`

GetDutyCurrencyOk returns a tuple with the DutyCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyCurrency

`func (o *OrderShipmentAdd) SetDutyCurrency(v string)`

SetDutyCurrency sets DutyCurrency field to given value.

### HasDutyCurrency

`func (o *OrderShipmentAdd) HasDutyCurrency() bool`

HasDutyCurrency returns a boolean if a field has been set.

### GetEnableCache

`func (o *OrderShipmentAdd) GetEnableCache() bool`

GetEnableCache returns the EnableCache field if non-nil, zero value otherwise.

### GetEnableCacheOk

`func (o *OrderShipmentAdd) GetEnableCacheOk() (*bool, bool)`

GetEnableCacheOk returns a tuple with the EnableCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCache

`func (o *OrderShipmentAdd) SetEnableCache(v bool)`

SetEnableCache sets EnableCache field to given value.

### HasEnableCache

`func (o *OrderShipmentAdd) HasEnableCache() bool`

HasEnableCache returns a boolean if a field has been set.

### GetUseLatestApiVersion

`func (o *OrderShipmentAdd) GetUseLatestApiVersion() bool`

GetUseLatestApiVersion returns the UseLatestApiVersion field if non-nil, zero value otherwise.

### GetUseLatestApiVersionOk

`func (o *OrderShipmentAdd) GetUseLatestApiVersionOk() (*bool, bool)`

GetUseLatestApiVersionOk returns a tuple with the UseLatestApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLatestApiVersion

`func (o *OrderShipmentAdd) SetUseLatestApiVersion(v bool)`

SetUseLatestApiVersion sets UseLatestApiVersion field to given value.

### HasUseLatestApiVersion

`func (o *OrderShipmentAdd) HasUseLatestApiVersion() bool`

HasUseLatestApiVersion returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *OrderShipmentAdd) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *OrderShipmentAdd) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *OrderShipmentAdd) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *OrderShipmentAdd) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


