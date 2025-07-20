# OrderPreestimateShipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodCode** | Pointer to **string** |  | [optional] 
**MethodName** | Pointer to **NullableString** |  | [optional] 
**CarrierCode** | Pointer to **NullableString** |  | [optional] 
**CarrierName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**PriceIncTax** | Pointer to **NullableFloat32** |  | [optional] 
**DeliveryTime** | Pointer to **NullableString** |  | [optional] 
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

### SetMethodNameNil

`func (o *OrderPreestimateShipping) SetMethodNameNil(b bool)`

 SetMethodNameNil sets the value for MethodName to be an explicit nil

### UnsetMethodName
`func (o *OrderPreestimateShipping) UnsetMethodName()`

UnsetMethodName ensures that no value is present for MethodName, not even an explicit nil
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

### SetCarrierCodeNil

`func (o *OrderPreestimateShipping) SetCarrierCodeNil(b bool)`

 SetCarrierCodeNil sets the value for CarrierCode to be an explicit nil

### UnsetCarrierCode
`func (o *OrderPreestimateShipping) UnsetCarrierCode()`

UnsetCarrierCode ensures that no value is present for CarrierCode, not even an explicit nil
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

### SetCarrierNameNil

`func (o *OrderPreestimateShipping) SetCarrierNameNil(b bool)`

 SetCarrierNameNil sets the value for CarrierName to be an explicit nil

### UnsetCarrierName
`func (o *OrderPreestimateShipping) UnsetCarrierName()`

UnsetCarrierName ensures that no value is present for CarrierName, not even an explicit nil
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

### SetDescriptionNil

`func (o *OrderPreestimateShipping) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OrderPreestimateShipping) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
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

### SetPriceNil

`func (o *OrderPreestimateShipping) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *OrderPreestimateShipping) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
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

### SetPriceIncTaxNil

`func (o *OrderPreestimateShipping) SetPriceIncTaxNil(b bool)`

 SetPriceIncTaxNil sets the value for PriceIncTax to be an explicit nil

### UnsetPriceIncTax
`func (o *OrderPreestimateShipping) UnsetPriceIncTax()`

UnsetPriceIncTax ensures that no value is present for PriceIncTax, not even an explicit nil
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

### SetDeliveryTimeNil

`func (o *OrderPreestimateShipping) SetDeliveryTimeNil(b bool)`

 SetDeliveryTimeNil sets the value for DeliveryTime to be an explicit nil

### UnsetDeliveryTime
`func (o *OrderPreestimateShipping) UnsetDeliveryTime()`

UnsetDeliveryTime ensures that no value is present for DeliveryTime, not even an explicit nil
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

### SetAdditionalFieldsNil

`func (o *OrderPreestimateShipping) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *OrderPreestimateShipping) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
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

### SetCustomFieldsNil

`func (o *OrderPreestimateShipping) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *OrderPreestimateShipping) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


