# OrderInvoiceAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Defines the order the document is uploaded for | [optional] 
**ShipmentId** | Pointer to **string** | Defines the shipment of the order the document is for. Required when the order was shipped in several shipments | [optional] 
**StoreId** | Pointer to **string** | Defines the store of the order. The store of the order is taken when not specified | [optional] 
**Number** | **string** | Defines the number of the invoice or credit note as printed on the document | 
**DocumentType** | Pointer to **string** | Defines whether the document is an invoice or a credit note for a refund or return | [optional] [default to "invoice"]
**TransactionId** | Pointer to **string** | Defines the refund or return transaction the credit note is issued for. Required when document_type is credit_note | [optional] 
**TotalPrice** | Pointer to **float32** | Defines the total amount of the document, VAT included. When specified, it must match the total the store calculated for the order or shipment. For sellers with VAT Calculation Service Option 2 Amazon requires both total_price and tax_price, otherwise it rejects the feed after processing | [optional] 
**TaxPrice** | Pointer to **float32** | Defines the total VAT amount of the document. When specified, it must match the VAT the store calculated for the order or shipment. For sellers with VAT Calculation Service Option 2 Amazon requires both total_price and tax_price, otherwise it rejects the feed after processing | [optional] 
**Url** | Pointer to **string** | Defines the public URL of the PDF document. Required if content is not specified | [optional] 
**Content** | Pointer to **string** | Defines the PDF document encoded in base64. Required if url is not specified | [optional] 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 

## Methods

### NewOrderInvoiceAdd

`func NewOrderInvoiceAdd(number string, ) *OrderInvoiceAdd`

NewOrderInvoiceAdd instantiates a new OrderInvoiceAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderInvoiceAddWithDefaults

`func NewOrderInvoiceAddWithDefaults() *OrderInvoiceAdd`

NewOrderInvoiceAddWithDefaults instantiates a new OrderInvoiceAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *OrderInvoiceAdd) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderInvoiceAdd) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderInvoiceAdd) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderInvoiceAdd) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetShipmentId

`func (o *OrderInvoiceAdd) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *OrderInvoiceAdd) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *OrderInvoiceAdd) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *OrderInvoiceAdd) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetStoreId

`func (o *OrderInvoiceAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderInvoiceAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderInvoiceAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderInvoiceAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetNumber

`func (o *OrderInvoiceAdd) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OrderInvoiceAdd) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OrderInvoiceAdd) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetDocumentType

`func (o *OrderInvoiceAdd) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *OrderInvoiceAdd) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *OrderInvoiceAdd) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *OrderInvoiceAdd) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetTransactionId

`func (o *OrderInvoiceAdd) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *OrderInvoiceAdd) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *OrderInvoiceAdd) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *OrderInvoiceAdd) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTotalPrice

`func (o *OrderInvoiceAdd) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *OrderInvoiceAdd) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *OrderInvoiceAdd) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *OrderInvoiceAdd) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetTaxPrice

`func (o *OrderInvoiceAdd) GetTaxPrice() float32`

GetTaxPrice returns the TaxPrice field if non-nil, zero value otherwise.

### GetTaxPriceOk

`func (o *OrderInvoiceAdd) GetTaxPriceOk() (*float32, bool)`

GetTaxPriceOk returns a tuple with the TaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPrice

`func (o *OrderInvoiceAdd) SetTaxPrice(v float32)`

SetTaxPrice sets TaxPrice field to given value.

### HasTaxPrice

`func (o *OrderInvoiceAdd) HasTaxPrice() bool`

HasTaxPrice returns a boolean if a field has been set.

### GetUrl

`func (o *OrderInvoiceAdd) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrderInvoiceAdd) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrderInvoiceAdd) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OrderInvoiceAdd) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetContent

`func (o *OrderInvoiceAdd) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OrderInvoiceAdd) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OrderInvoiceAdd) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OrderInvoiceAdd) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *OrderInvoiceAdd) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *OrderInvoiceAdd) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *OrderInvoiceAdd) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *OrderInvoiceAdd) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


