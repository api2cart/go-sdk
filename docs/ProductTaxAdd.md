# ProductTaxAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** | Defines products specified by product id | [optional] 
**Name** | **string** | Defines tax class name where tax has to be added | 
**TaxRates** | [**[]ProductTaxAddTaxRatesInner**](ProductTaxAddTaxRatesInner.md) | Defines tax rates of specified tax classes | 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 

## Methods

### NewProductTaxAdd

`func NewProductTaxAdd(name string, taxRates []ProductTaxAddTaxRatesInner, ) *ProductTaxAdd`

NewProductTaxAdd instantiates a new ProductTaxAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTaxAddWithDefaults

`func NewProductTaxAddWithDefaults() *ProductTaxAdd`

NewProductTaxAddWithDefaults instantiates a new ProductTaxAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProductTaxAdd) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductTaxAdd) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductTaxAdd) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductTaxAdd) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetName

`func (o *ProductTaxAdd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductTaxAdd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductTaxAdd) SetName(v string)`

SetName sets Name field to given value.


### GetTaxRates

`func (o *ProductTaxAdd) GetTaxRates() []ProductTaxAddTaxRatesInner`

GetTaxRates returns the TaxRates field if non-nil, zero value otherwise.

### GetTaxRatesOk

`func (o *ProductTaxAdd) GetTaxRatesOk() (*[]ProductTaxAddTaxRatesInner, bool)`

GetTaxRatesOk returns a tuple with the TaxRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRates

`func (o *ProductTaxAdd) SetTaxRates(v []ProductTaxAddTaxRatesInner)`

SetTaxRates sets TaxRates field to given value.


### GetIdempotencyKey

`func (o *ProductTaxAdd) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ProductTaxAdd) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ProductTaxAdd) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *ProductTaxAdd) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


