# OrderPreestimateShippingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreId** | Pointer to **string** | Store Id | [optional] 
**WarehouseId** | Pointer to **string** | This parameter is used for selecting a warehouse where you need to set/modify a product quantity. | [optional] 
**CustomerEmail** | Pointer to **string** | Retrieves orders specified by customer email | [optional] 
**CustomerId** | Pointer to **string** | Retrieves orders specified by customer id | [optional] 
**ShippAddress1** | Pointer to **string** | Specifies first shipping address | [optional] 
**ShippCity** | Pointer to **string** | Specifies shipping city | [optional] 
**ShippPostcode** | Pointer to **string** | Specifies shipping postcode | [optional] 
**ShippState** | Pointer to **string** | Specifies shipping state code | [optional] 
**ShippCountry** | **string** | Specifies shipping country code | 
**Params** | Pointer to **string** | Set this parameter in order to choose which entity fields you want to retrieve | [optional] [default to "force_all"]
**Exclude** | Pointer to **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | [optional] 
**OrderItem** | [**[]OrderPreestimateShippingListOrderItemInner**](OrderPreestimateShippingListOrderItemInner.md) |  | 

## Methods

### NewOrderPreestimateShippingList

`func NewOrderPreestimateShippingList(shippCountry string, orderItem []OrderPreestimateShippingListOrderItemInner, ) *OrderPreestimateShippingList`

NewOrderPreestimateShippingList instantiates a new OrderPreestimateShippingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPreestimateShippingListWithDefaults

`func NewOrderPreestimateShippingListWithDefaults() *OrderPreestimateShippingList`

NewOrderPreestimateShippingListWithDefaults instantiates a new OrderPreestimateShippingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreId

`func (o *OrderPreestimateShippingList) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderPreestimateShippingList) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderPreestimateShippingList) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderPreestimateShippingList) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *OrderPreestimateShippingList) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *OrderPreestimateShippingList) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *OrderPreestimateShippingList) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *OrderPreestimateShippingList) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *OrderPreestimateShippingList) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *OrderPreestimateShippingList) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *OrderPreestimateShippingList) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *OrderPreestimateShippingList) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerId

`func (o *OrderPreestimateShippingList) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *OrderPreestimateShippingList) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *OrderPreestimateShippingList) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *OrderPreestimateShippingList) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetShippAddress1

`func (o *OrderPreestimateShippingList) GetShippAddress1() string`

GetShippAddress1 returns the ShippAddress1 field if non-nil, zero value otherwise.

### GetShippAddress1Ok

`func (o *OrderPreestimateShippingList) GetShippAddress1Ok() (*string, bool)`

GetShippAddress1Ok returns a tuple with the ShippAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippAddress1

`func (o *OrderPreestimateShippingList) SetShippAddress1(v string)`

SetShippAddress1 sets ShippAddress1 field to given value.

### HasShippAddress1

`func (o *OrderPreestimateShippingList) HasShippAddress1() bool`

HasShippAddress1 returns a boolean if a field has been set.

### GetShippCity

`func (o *OrderPreestimateShippingList) GetShippCity() string`

GetShippCity returns the ShippCity field if non-nil, zero value otherwise.

### GetShippCityOk

`func (o *OrderPreestimateShippingList) GetShippCityOk() (*string, bool)`

GetShippCityOk returns a tuple with the ShippCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippCity

`func (o *OrderPreestimateShippingList) SetShippCity(v string)`

SetShippCity sets ShippCity field to given value.

### HasShippCity

`func (o *OrderPreestimateShippingList) HasShippCity() bool`

HasShippCity returns a boolean if a field has been set.

### GetShippPostcode

`func (o *OrderPreestimateShippingList) GetShippPostcode() string`

GetShippPostcode returns the ShippPostcode field if non-nil, zero value otherwise.

### GetShippPostcodeOk

`func (o *OrderPreestimateShippingList) GetShippPostcodeOk() (*string, bool)`

GetShippPostcodeOk returns a tuple with the ShippPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippPostcode

`func (o *OrderPreestimateShippingList) SetShippPostcode(v string)`

SetShippPostcode sets ShippPostcode field to given value.

### HasShippPostcode

`func (o *OrderPreestimateShippingList) HasShippPostcode() bool`

HasShippPostcode returns a boolean if a field has been set.

### GetShippState

`func (o *OrderPreestimateShippingList) GetShippState() string`

GetShippState returns the ShippState field if non-nil, zero value otherwise.

### GetShippStateOk

`func (o *OrderPreestimateShippingList) GetShippStateOk() (*string, bool)`

GetShippStateOk returns a tuple with the ShippState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippState

`func (o *OrderPreestimateShippingList) SetShippState(v string)`

SetShippState sets ShippState field to given value.

### HasShippState

`func (o *OrderPreestimateShippingList) HasShippState() bool`

HasShippState returns a boolean if a field has been set.

### GetShippCountry

`func (o *OrderPreestimateShippingList) GetShippCountry() string`

GetShippCountry returns the ShippCountry field if non-nil, zero value otherwise.

### GetShippCountryOk

`func (o *OrderPreestimateShippingList) GetShippCountryOk() (*string, bool)`

GetShippCountryOk returns a tuple with the ShippCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippCountry

`func (o *OrderPreestimateShippingList) SetShippCountry(v string)`

SetShippCountry sets ShippCountry field to given value.


### GetParams

`func (o *OrderPreestimateShippingList) GetParams() string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *OrderPreestimateShippingList) GetParamsOk() (*string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *OrderPreestimateShippingList) SetParams(v string)`

SetParams sets Params field to given value.

### HasParams

`func (o *OrderPreestimateShippingList) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetExclude

`func (o *OrderPreestimateShippingList) GetExclude() string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *OrderPreestimateShippingList) GetExcludeOk() (*string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *OrderPreestimateShippingList) SetExclude(v string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *OrderPreestimateShippingList) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetOrderItem

`func (o *OrderPreestimateShippingList) GetOrderItem() []OrderPreestimateShippingListOrderItemInner`

GetOrderItem returns the OrderItem field if non-nil, zero value otherwise.

### GetOrderItemOk

`func (o *OrderPreestimateShippingList) GetOrderItemOk() (*[]OrderPreestimateShippingListOrderItemInner, bool)`

GetOrderItemOk returns a tuple with the OrderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItem

`func (o *OrderPreestimateShippingList) SetOrderItem(v []OrderPreestimateShippingListOrderItemInner)`

SetOrderItem sets OrderItem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


