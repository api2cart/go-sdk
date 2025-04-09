# OrderAbandoned

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**BaseCustomer**](BaseCustomer.md) |  | [optional] 
**BasketId** | Pointer to **string** |  | [optional] 
**BasketUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Totals** | Pointer to [**OrderTotals**](OrderTotals.md) |  | [optional] 
**OrderProducts** | Pointer to [**[]OrderItem**](OrderItem.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderAbandoned

`func NewOrderAbandoned() *OrderAbandoned`

NewOrderAbandoned instantiates a new OrderAbandoned object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAbandonedWithDefaults

`func NewOrderAbandonedWithDefaults() *OrderAbandoned`

NewOrderAbandonedWithDefaults instantiates a new OrderAbandoned object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderAbandoned) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderAbandoned) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderAbandoned) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderAbandoned) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderAbandoned) GetCustomer() BaseCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderAbandoned) GetCustomerOk() (*BaseCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderAbandoned) SetCustomer(v BaseCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderAbandoned) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetBasketId

`func (o *OrderAbandoned) GetBasketId() string`

GetBasketId returns the BasketId field if non-nil, zero value otherwise.

### GetBasketIdOk

`func (o *OrderAbandoned) GetBasketIdOk() (*string, bool)`

GetBasketIdOk returns a tuple with the BasketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasketId

`func (o *OrderAbandoned) SetBasketId(v string)`

SetBasketId sets BasketId field to given value.

### HasBasketId

`func (o *OrderAbandoned) HasBasketId() bool`

HasBasketId returns a boolean if a field has been set.

### GetBasketUrl

`func (o *OrderAbandoned) GetBasketUrl() string`

GetBasketUrl returns the BasketUrl field if non-nil, zero value otherwise.

### GetBasketUrlOk

`func (o *OrderAbandoned) GetBasketUrlOk() (*string, bool)`

GetBasketUrlOk returns a tuple with the BasketUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasketUrl

`func (o *OrderAbandoned) SetBasketUrl(v string)`

SetBasketUrl sets BasketUrl field to given value.

### HasBasketUrl

`func (o *OrderAbandoned) HasBasketUrl() bool`

HasBasketUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderAbandoned) GetCreatedAt() A2CDateTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderAbandoned) GetCreatedAtOk() (*A2CDateTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderAbandoned) SetCreatedAt(v A2CDateTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderAbandoned) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *OrderAbandoned) GetModifiedAt() A2CDateTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *OrderAbandoned) GetModifiedAtOk() (*A2CDateTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *OrderAbandoned) SetModifiedAt(v A2CDateTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *OrderAbandoned) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *OrderAbandoned) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderAbandoned) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderAbandoned) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderAbandoned) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotals

`func (o *OrderAbandoned) GetTotals() OrderTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *OrderAbandoned) GetTotalsOk() (*OrderTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *OrderAbandoned) SetTotals(v OrderTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *OrderAbandoned) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetOrderProducts

`func (o *OrderAbandoned) GetOrderProducts() []OrderItem`

GetOrderProducts returns the OrderProducts field if non-nil, zero value otherwise.

### GetOrderProductsOk

`func (o *OrderAbandoned) GetOrderProductsOk() (*[]OrderItem, bool)`

GetOrderProductsOk returns a tuple with the OrderProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProducts

`func (o *OrderAbandoned) SetOrderProducts(v []OrderItem)`

SetOrderProducts sets OrderProducts field to given value.

### HasOrderProducts

`func (o *OrderAbandoned) HasOrderProducts() bool`

HasOrderProducts returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *OrderAbandoned) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderAbandoned) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderAbandoned) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderAbandoned) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *OrderAbandoned) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderAbandoned) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderAbandoned) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderAbandoned) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


