# OrderPreestimateShipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodCode** | Pointer to **string** |  | [optional] 
**MethodName** | Pointer to **string** |  | [optional] 
**CarrierCode** | Pointer to **string** |  | [optional] 
**CarrierName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**PriceIncTax** | Pointer to **float32** |  | [optional] 
**DeliveryTime** | Pointer to **string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOrderPreestimateShipping

`func NewOrderPreestimateShipping() *OrderPreestimateShipping`

NewOrderPreestimateShipping instantiates a new OrderPreestimateShipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPreestimateShippingWithDefaults

`func NewOrderPreestimateShippingWithDefaults() *OrderPreestimateShipping`

NewOrderPreestimateShippingWithDefaults instantiates a new OrderPreestimateShipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodCode

`func (o *OrderPreestimateShipping) GetMethodCode() string`

GetMethodCode returns the MethodCode field if non-nil, zero value otherwise.

### GetMethodCodeOk

`func (o *OrderPreestimateShipping) GetMethodCodeOk() (*string, bool)`

GetMethodCodeOk returns a tuple with the MethodCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodCode

`func (o *OrderPreestimateShipping) SetMethodCode(v string)`

SetMethodCode sets MethodCode field to given value.

### HasMethodCode

`func (o *OrderPreestimateShipping) HasMethodCode() bool`

HasMethodCode returns a boolean if a field has been set.

### GetMethodName

`func (o *OrderPreestimateShipping) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *OrderPreestimateShipping) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *OrderPreestimateShipping) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.

### HasMethodName

`func (o *OrderPreestimateShipping) HasMethodName() bool`

HasMethodName returns a boolean if a field has been set.

### GetCarrierCode

`func (o *OrderPreestimateShipping) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *OrderPreestimateShipping) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *OrderPreestimateShipping) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *OrderPreestimateShipping) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetCarrierName

`func (o *OrderPreestimateShipping) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *OrderPreestimateShipping) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *OrderPreestimateShipping) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *OrderPreestimateShipping) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetDescription

`func (o *OrderPreestimateShipping) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderPreestimateShipping) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderPreestimateShipping) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrderPreestimateShipping) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrice

`func (o *OrderPreestimateShipping) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderPreestimateShipping) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderPreestimateShipping) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderPreestimateShipping) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceIncTax

`func (o *OrderPreestimateShipping) GetPriceIncTax() float32`

GetPriceIncTax returns the PriceIncTax field if non-nil, zero value otherwise.

### GetPriceIncTaxOk

`func (o *OrderPreestimateShipping) GetPriceIncTaxOk() (*float32, bool)`

GetPriceIncTaxOk returns a tuple with the PriceIncTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceIncTax

`func (o *OrderPreestimateShipping) SetPriceIncTax(v float32)`

SetPriceIncTax sets PriceIncTax field to given value.

### HasPriceIncTax

`func (o *OrderPreestimateShipping) HasPriceIncTax() bool`

HasPriceIncTax returns a boolean if a field has been set.

### GetDeliveryTime

`func (o *OrderPreestimateShipping) GetDeliveryTime() string`

GetDeliveryTime returns the DeliveryTime field if non-nil, zero value otherwise.

### GetDeliveryTimeOk

`func (o *OrderPreestimateShipping) GetDeliveryTimeOk() (*string, bool)`

GetDeliveryTimeOk returns a tuple with the DeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTime

`func (o *OrderPreestimateShipping) SetDeliveryTime(v string)`

SetDeliveryTime sets DeliveryTime field to given value.

### HasDeliveryTime

`func (o *OrderPreestimateShipping) HasDeliveryTime() bool`

HasDeliveryTime returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *OrderPreestimateShipping) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *OrderPreestimateShipping) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *OrderPreestimateShipping) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *OrderPreestimateShipping) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *OrderPreestimateShipping) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *OrderPreestimateShipping) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *OrderPreestimateShipping) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *OrderPreestimateShipping) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


