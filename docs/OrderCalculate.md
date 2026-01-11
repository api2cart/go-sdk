# OrderCalculate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerEmail** | **string** | Defines the customer specified by email for whom the order needs to be calculated | 
**CurrencyId** | Pointer to **string** | Currency Id | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**Coupons** | Pointer to **[]string** | Coupons that will be applied to order. If the order isn&#39;t eligible for any given discount code or there is no discount with such a code it will be skipped during calculation | [optional] 
**RoundingPrecision** | Pointer to **int32** | &lt;p&gt;Specifies the rounding precision for fractional numeric values (such as prices, taxes, and weights).&lt;/p&gt; &lt;p&gt;Supported values range from &lt;b&gt;1&lt;/b&gt; to &lt;b&gt;6&lt;/b&gt;.&lt;/p&gt; &lt;p&gt;The default rounding precision may vary depending on the platform. You can retrieve the default value using the &lt;strong&gt;cart.info&lt;/strong&gt; method in the &lt;code&gt;default_rounding_precision&lt;/code&gt; field. &lt;/p&gt;&lt;p&gt;Values are rounded to the nearest number at the specified precision. Fractions of .5 or higher are rounded up, while fractions lower than .5 are rounded down.&lt;/p&gt; | [optional] 
**ShippFirstName** | **string** | Specifies shipping first name | 
**ShippLastName** | **string** | Specifies shipping last name | 
**ShippAddress1** | **string** | Specifies first shipping address | 
**ShippAddress2** | Pointer to **string** | Specifies second address line of a shipping street address | [optional] 
**ShippCity** | **string** | Specifies shipping city | 
**ShippPostcode** | **string** | Specifies shipping postcode | 
**ShippState** | Pointer to **string** | Specifies shipping state code | [optional] 
**ShippCountry** | **string** | Specifies shipping country code | 
**ShippCompany** | Pointer to **string** | Specifies shipping company | [optional] 
**ShippPhone** | Pointer to **string** | Specifies shipping phone | [optional] 
**BillFirstName** | Pointer to **string** | Specifies billing first name | [optional] 
**BillLastName** | Pointer to **string** | Specifies billing last name | [optional] 
**BillAddress1** | Pointer to **string** | Specifies first billing address | [optional] 
**BillAddress2** | Pointer to **string** | Specifies second billing address | [optional] 
**BillCity** | Pointer to **string** | Specifies billing city | [optional] 
**BillPostcode** | Pointer to **string** | Specifies billing postcode | [optional] 
**BillState** | Pointer to **string** | Specifies billing state code | [optional] 
**BillCountry** | Pointer to **string** | Specifies billing country code | [optional] 
**BillCompany** | Pointer to **string** | Specifies billing company | [optional] 
**BillPhone** | Pointer to **string** | Specifies billing phone | [optional] 
**ResponseFields** | Pointer to **string** | Set this parameter in order to choose which entity fields you want to retrieve | [optional] 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 
**OrderItem** | [**[]OrderCalculateOrderItemInner**](OrderCalculateOrderItemInner.md) |  | 

## Methods

### NewOrderCalculate

`func NewOrderCalculate(customerEmail string, shippFirstName string, shippLastName string, shippAddress1 string, shippCity string, shippPostcode string, shippCountry string, orderItem []OrderCalculateOrderItemInner, ) *OrderCalculate`

NewOrderCalculate instantiates a new OrderCalculate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCalculateWithDefaults

`func NewOrderCalculateWithDefaults() *OrderCalculate`

NewOrderCalculateWithDefaults instantiates a new OrderCalculate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerEmail

`func (o *OrderCalculate) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *OrderCalculate) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *OrderCalculate) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.


### GetCurrencyId

`func (o *OrderCalculate) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *OrderCalculate) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *OrderCalculate) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *OrderCalculate) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetStoreId

`func (o *OrderCalculate) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderCalculate) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderCalculate) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderCalculate) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetCoupons

`func (o *OrderCalculate) GetCoupons() []string`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *OrderCalculate) GetCouponsOk() (*[]string, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *OrderCalculate) SetCoupons(v []string)`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *OrderCalculate) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### GetRoundingPrecision

`func (o *OrderCalculate) GetRoundingPrecision() int32`

GetRoundingPrecision returns the RoundingPrecision field if non-nil, zero value otherwise.

### GetRoundingPrecisionOk

`func (o *OrderCalculate) GetRoundingPrecisionOk() (*int32, bool)`

GetRoundingPrecisionOk returns a tuple with the RoundingPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingPrecision

`func (o *OrderCalculate) SetRoundingPrecision(v int32)`

SetRoundingPrecision sets RoundingPrecision field to given value.

### HasRoundingPrecision

`func (o *OrderCalculate) HasRoundingPrecision() bool`

HasRoundingPrecision returns a boolean if a field has been set.

### GetShippFirstName

`func (o *OrderCalculate) GetShippFirstName() string`

GetShippFirstName returns the ShippFirstName field if non-nil, zero value otherwise.

### GetShippFirstNameOk

`func (o *OrderCalculate) GetShippFirstNameOk() (*string, bool)`

GetShippFirstNameOk returns a tuple with the ShippFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippFirstName

`func (o *OrderCalculate) SetShippFirstName(v string)`

SetShippFirstName sets ShippFirstName field to given value.


### GetShippLastName

`func (o *OrderCalculate) GetShippLastName() string`

GetShippLastName returns the ShippLastName field if non-nil, zero value otherwise.

### GetShippLastNameOk

`func (o *OrderCalculate) GetShippLastNameOk() (*string, bool)`

GetShippLastNameOk returns a tuple with the ShippLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippLastName

`func (o *OrderCalculate) SetShippLastName(v string)`

SetShippLastName sets ShippLastName field to given value.


### GetShippAddress1

`func (o *OrderCalculate) GetShippAddress1() string`

GetShippAddress1 returns the ShippAddress1 field if non-nil, zero value otherwise.

### GetShippAddress1Ok

`func (o *OrderCalculate) GetShippAddress1Ok() (*string, bool)`

GetShippAddress1Ok returns a tuple with the ShippAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippAddress1

`func (o *OrderCalculate) SetShippAddress1(v string)`

SetShippAddress1 sets ShippAddress1 field to given value.


### GetShippAddress2

`func (o *OrderCalculate) GetShippAddress2() string`

GetShippAddress2 returns the ShippAddress2 field if non-nil, zero value otherwise.

### GetShippAddress2Ok

`func (o *OrderCalculate) GetShippAddress2Ok() (*string, bool)`

GetShippAddress2Ok returns a tuple with the ShippAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippAddress2

`func (o *OrderCalculate) SetShippAddress2(v string)`

SetShippAddress2 sets ShippAddress2 field to given value.

### HasShippAddress2

`func (o *OrderCalculate) HasShippAddress2() bool`

HasShippAddress2 returns a boolean if a field has been set.

### GetShippCity

`func (o *OrderCalculate) GetShippCity() string`

GetShippCity returns the ShippCity field if non-nil, zero value otherwise.

### GetShippCityOk

`func (o *OrderCalculate) GetShippCityOk() (*string, bool)`

GetShippCityOk returns a tuple with the ShippCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippCity

`func (o *OrderCalculate) SetShippCity(v string)`

SetShippCity sets ShippCity field to given value.


### GetShippPostcode

`func (o *OrderCalculate) GetShippPostcode() string`

GetShippPostcode returns the ShippPostcode field if non-nil, zero value otherwise.

### GetShippPostcodeOk

`func (o *OrderCalculate) GetShippPostcodeOk() (*string, bool)`

GetShippPostcodeOk returns a tuple with the ShippPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippPostcode

`func (o *OrderCalculate) SetShippPostcode(v string)`

SetShippPostcode sets ShippPostcode field to given value.


### GetShippState

`func (o *OrderCalculate) GetShippState() string`

GetShippState returns the ShippState field if non-nil, zero value otherwise.

### GetShippStateOk

`func (o *OrderCalculate) GetShippStateOk() (*string, bool)`

GetShippStateOk returns a tuple with the ShippState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippState

`func (o *OrderCalculate) SetShippState(v string)`

SetShippState sets ShippState field to given value.

### HasShippState

`func (o *OrderCalculate) HasShippState() bool`

HasShippState returns a boolean if a field has been set.

### GetShippCountry

`func (o *OrderCalculate) GetShippCountry() string`

GetShippCountry returns the ShippCountry field if non-nil, zero value otherwise.

### GetShippCountryOk

`func (o *OrderCalculate) GetShippCountryOk() (*string, bool)`

GetShippCountryOk returns a tuple with the ShippCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippCountry

`func (o *OrderCalculate) SetShippCountry(v string)`

SetShippCountry sets ShippCountry field to given value.


### GetShippCompany

`func (o *OrderCalculate) GetShippCompany() string`

GetShippCompany returns the ShippCompany field if non-nil, zero value otherwise.

### GetShippCompanyOk

`func (o *OrderCalculate) GetShippCompanyOk() (*string, bool)`

GetShippCompanyOk returns a tuple with the ShippCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippCompany

`func (o *OrderCalculate) SetShippCompany(v string)`

SetShippCompany sets ShippCompany field to given value.

### HasShippCompany

`func (o *OrderCalculate) HasShippCompany() bool`

HasShippCompany returns a boolean if a field has been set.

### GetShippPhone

`func (o *OrderCalculate) GetShippPhone() string`

GetShippPhone returns the ShippPhone field if non-nil, zero value otherwise.

### GetShippPhoneOk

`func (o *OrderCalculate) GetShippPhoneOk() (*string, bool)`

GetShippPhoneOk returns a tuple with the ShippPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippPhone

`func (o *OrderCalculate) SetShippPhone(v string)`

SetShippPhone sets ShippPhone field to given value.

### HasShippPhone

`func (o *OrderCalculate) HasShippPhone() bool`

HasShippPhone returns a boolean if a field has been set.

### GetBillFirstName

`func (o *OrderCalculate) GetBillFirstName() string`

GetBillFirstName returns the BillFirstName field if non-nil, zero value otherwise.

### GetBillFirstNameOk

`func (o *OrderCalculate) GetBillFirstNameOk() (*string, bool)`

GetBillFirstNameOk returns a tuple with the BillFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillFirstName

`func (o *OrderCalculate) SetBillFirstName(v string)`

SetBillFirstName sets BillFirstName field to given value.

### HasBillFirstName

`func (o *OrderCalculate) HasBillFirstName() bool`

HasBillFirstName returns a boolean if a field has been set.

### GetBillLastName

`func (o *OrderCalculate) GetBillLastName() string`

GetBillLastName returns the BillLastName field if non-nil, zero value otherwise.

### GetBillLastNameOk

`func (o *OrderCalculate) GetBillLastNameOk() (*string, bool)`

GetBillLastNameOk returns a tuple with the BillLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillLastName

`func (o *OrderCalculate) SetBillLastName(v string)`

SetBillLastName sets BillLastName field to given value.

### HasBillLastName

`func (o *OrderCalculate) HasBillLastName() bool`

HasBillLastName returns a boolean if a field has been set.

### GetBillAddress1

`func (o *OrderCalculate) GetBillAddress1() string`

GetBillAddress1 returns the BillAddress1 field if non-nil, zero value otherwise.

### GetBillAddress1Ok

`func (o *OrderCalculate) GetBillAddress1Ok() (*string, bool)`

GetBillAddress1Ok returns a tuple with the BillAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillAddress1

`func (o *OrderCalculate) SetBillAddress1(v string)`

SetBillAddress1 sets BillAddress1 field to given value.

### HasBillAddress1

`func (o *OrderCalculate) HasBillAddress1() bool`

HasBillAddress1 returns a boolean if a field has been set.

### GetBillAddress2

`func (o *OrderCalculate) GetBillAddress2() string`

GetBillAddress2 returns the BillAddress2 field if non-nil, zero value otherwise.

### GetBillAddress2Ok

`func (o *OrderCalculate) GetBillAddress2Ok() (*string, bool)`

GetBillAddress2Ok returns a tuple with the BillAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillAddress2

`func (o *OrderCalculate) SetBillAddress2(v string)`

SetBillAddress2 sets BillAddress2 field to given value.

### HasBillAddress2

`func (o *OrderCalculate) HasBillAddress2() bool`

HasBillAddress2 returns a boolean if a field has been set.

### GetBillCity

`func (o *OrderCalculate) GetBillCity() string`

GetBillCity returns the BillCity field if non-nil, zero value otherwise.

### GetBillCityOk

`func (o *OrderCalculate) GetBillCityOk() (*string, bool)`

GetBillCityOk returns a tuple with the BillCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCity

`func (o *OrderCalculate) SetBillCity(v string)`

SetBillCity sets BillCity field to given value.

### HasBillCity

`func (o *OrderCalculate) HasBillCity() bool`

HasBillCity returns a boolean if a field has been set.

### GetBillPostcode

`func (o *OrderCalculate) GetBillPostcode() string`

GetBillPostcode returns the BillPostcode field if non-nil, zero value otherwise.

### GetBillPostcodeOk

`func (o *OrderCalculate) GetBillPostcodeOk() (*string, bool)`

GetBillPostcodeOk returns a tuple with the BillPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillPostcode

`func (o *OrderCalculate) SetBillPostcode(v string)`

SetBillPostcode sets BillPostcode field to given value.

### HasBillPostcode

`func (o *OrderCalculate) HasBillPostcode() bool`

HasBillPostcode returns a boolean if a field has been set.

### GetBillState

`func (o *OrderCalculate) GetBillState() string`

GetBillState returns the BillState field if non-nil, zero value otherwise.

### GetBillStateOk

`func (o *OrderCalculate) GetBillStateOk() (*string, bool)`

GetBillStateOk returns a tuple with the BillState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillState

`func (o *OrderCalculate) SetBillState(v string)`

SetBillState sets BillState field to given value.

### HasBillState

`func (o *OrderCalculate) HasBillState() bool`

HasBillState returns a boolean if a field has been set.

### GetBillCountry

`func (o *OrderCalculate) GetBillCountry() string`

GetBillCountry returns the BillCountry field if non-nil, zero value otherwise.

### GetBillCountryOk

`func (o *OrderCalculate) GetBillCountryOk() (*string, bool)`

GetBillCountryOk returns a tuple with the BillCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCountry

`func (o *OrderCalculate) SetBillCountry(v string)`

SetBillCountry sets BillCountry field to given value.

### HasBillCountry

`func (o *OrderCalculate) HasBillCountry() bool`

HasBillCountry returns a boolean if a field has been set.

### GetBillCompany

`func (o *OrderCalculate) GetBillCompany() string`

GetBillCompany returns the BillCompany field if non-nil, zero value otherwise.

### GetBillCompanyOk

`func (o *OrderCalculate) GetBillCompanyOk() (*string, bool)`

GetBillCompanyOk returns a tuple with the BillCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCompany

`func (o *OrderCalculate) SetBillCompany(v string)`

SetBillCompany sets BillCompany field to given value.

### HasBillCompany

`func (o *OrderCalculate) HasBillCompany() bool`

HasBillCompany returns a boolean if a field has been set.

### GetBillPhone

`func (o *OrderCalculate) GetBillPhone() string`

GetBillPhone returns the BillPhone field if non-nil, zero value otherwise.

### GetBillPhoneOk

`func (o *OrderCalculate) GetBillPhoneOk() (*string, bool)`

GetBillPhoneOk returns a tuple with the BillPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillPhone

`func (o *OrderCalculate) SetBillPhone(v string)`

SetBillPhone sets BillPhone field to given value.

### HasBillPhone

`func (o *OrderCalculate) HasBillPhone() bool`

HasBillPhone returns a boolean if a field has been set.

### GetResponseFields

`func (o *OrderCalculate) GetResponseFields() string`

GetResponseFields returns the ResponseFields field if non-nil, zero value otherwise.

### GetResponseFieldsOk

`func (o *OrderCalculate) GetResponseFieldsOk() (*string, bool)`

GetResponseFieldsOk returns a tuple with the ResponseFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFields

`func (o *OrderCalculate) SetResponseFields(v string)`

SetResponseFields sets ResponseFields field to given value.

### HasResponseFields

`func (o *OrderCalculate) HasResponseFields() bool`

HasResponseFields returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *OrderCalculate) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *OrderCalculate) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *OrderCalculate) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *OrderCalculate) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetOrderItem

`func (o *OrderCalculate) GetOrderItem() []OrderCalculateOrderItemInner`

GetOrderItem returns the OrderItem field if non-nil, zero value otherwise.

### GetOrderItemOk

`func (o *OrderCalculate) GetOrderItemOk() (*[]OrderCalculateOrderItemInner, bool)`

GetOrderItemOk returns a tuple with the OrderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItem

`func (o *OrderCalculate) SetOrderItem(v []OrderCalculateOrderItemInner)`

SetOrderItem sets OrderItem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


