# Basket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**BaseCustomer**](BaseCustomer.md) |  | [optional] 
**BasketUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**BasketProducts** | Pointer to [**[]BasketItem**](BasketItem.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBasket

`func NewBasket() *Basket`

NewBasket instantiates a new Basket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasketWithDefaults

`func NewBasketWithDefaults() *Basket`

NewBasketWithDefaults instantiates a new Basket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Basket) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Basket) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Basket) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Basket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomer

`func (o *Basket) GetCustomer() BaseCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Basket) GetCustomerOk() (*BaseCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Basket) SetCustomer(v BaseCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Basket) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetBasketUrl

`func (o *Basket) GetBasketUrl() string`

GetBasketUrl returns the BasketUrl field if non-nil, zero value otherwise.

### GetBasketUrlOk

`func (o *Basket) GetBasketUrlOk() (*string, bool)`

GetBasketUrlOk returns a tuple with the BasketUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasketUrl

`func (o *Basket) SetBasketUrl(v string)`

SetBasketUrl sets BasketUrl field to given value.

### HasBasketUrl

`func (o *Basket) HasBasketUrl() bool`

HasBasketUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Basket) GetCreatedAt() A2CDateTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Basket) GetCreatedAtOk() (*A2CDateTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Basket) SetCreatedAt(v A2CDateTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Basket) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Basket) GetModifiedAt() A2CDateTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Basket) GetModifiedAtOk() (*A2CDateTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Basket) SetModifiedAt(v A2CDateTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Basket) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *Basket) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Basket) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Basket) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Basket) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBasketProducts

`func (o *Basket) GetBasketProducts() []BasketItem`

GetBasketProducts returns the BasketProducts field if non-nil, zero value otherwise.

### GetBasketProductsOk

`func (o *Basket) GetBasketProductsOk() (*[]BasketItem, bool)`

GetBasketProductsOk returns a tuple with the BasketProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasketProducts

`func (o *Basket) SetBasketProducts(v []BasketItem)`

SetBasketProducts sets BasketProducts field to given value.

### HasBasketProducts

`func (o *Basket) HasBasketProducts() bool`

HasBasketProducts returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Basket) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Basket) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Basket) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Basket) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *Basket) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Basket) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Basket) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Basket) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


