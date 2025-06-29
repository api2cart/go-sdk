# ProductAddSpecificsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 
**UsedForVariations** | Pointer to **bool** |  | [optional] [default to false]
**ScaleId** | Pointer to **NullableInt32** |  | [optional] 
**InputValue** | Pointer to **NullableString** |  | [optional] 
**FoodDetails** | Pointer to [**ProductAddSpecificsInnerFoodDetails**](ProductAddSpecificsInnerFoodDetails.md) |  | [optional] 
**GroupProductsDetails** | Pointer to [**[]ProductAddSpecificsInnerGroupProductsDetailsInner**](ProductAddSpecificsInnerGroupProductsDetailsInner.md) |  | [optional] 
**BookingDetails** | Pointer to [**ProductAddSpecificsInnerBookingDetails**](ProductAddSpecificsInnerBookingDetails.md) |  | [optional] 

## Methods

### NewProductAddSpecificsInner

`func NewProductAddSpecificsInner() *ProductAddSpecificsInner`

NewProductAddSpecificsInner instantiates a new ProductAddSpecificsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAddSpecificsInnerWithDefaults

`func NewProductAddSpecificsInnerWithDefaults() *ProductAddSpecificsInner`

NewProductAddSpecificsInnerWithDefaults instantiates a new ProductAddSpecificsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductAddSpecificsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductAddSpecificsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductAddSpecificsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductAddSpecificsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ProductAddSpecificsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProductAddSpecificsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProductAddSpecificsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProductAddSpecificsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *ProductAddSpecificsInner) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ProductAddSpecificsInner) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ProductAddSpecificsInner) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ProductAddSpecificsInner) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetUsedForVariations

`func (o *ProductAddSpecificsInner) GetUsedForVariations() bool`

GetUsedForVariations returns the UsedForVariations field if non-nil, zero value otherwise.

### GetUsedForVariationsOk

`func (o *ProductAddSpecificsInner) GetUsedForVariationsOk() (*bool, bool)`

GetUsedForVariationsOk returns a tuple with the UsedForVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedForVariations

`func (o *ProductAddSpecificsInner) SetUsedForVariations(v bool)`

SetUsedForVariations sets UsedForVariations field to given value.

### HasUsedForVariations

`func (o *ProductAddSpecificsInner) HasUsedForVariations() bool`

HasUsedForVariations returns a boolean if a field has been set.

### GetScaleId

`func (o *ProductAddSpecificsInner) GetScaleId() int32`

GetScaleId returns the ScaleId field if non-nil, zero value otherwise.

### GetScaleIdOk

`func (o *ProductAddSpecificsInner) GetScaleIdOk() (*int32, bool)`

GetScaleIdOk returns a tuple with the ScaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleId

`func (o *ProductAddSpecificsInner) SetScaleId(v int32)`

SetScaleId sets ScaleId field to given value.

### HasScaleId

`func (o *ProductAddSpecificsInner) HasScaleId() bool`

HasScaleId returns a boolean if a field has been set.

### SetScaleIdNil

`func (o *ProductAddSpecificsInner) SetScaleIdNil(b bool)`

 SetScaleIdNil sets the value for ScaleId to be an explicit nil

### UnsetScaleId
`func (o *ProductAddSpecificsInner) UnsetScaleId()`

UnsetScaleId ensures that no value is present for ScaleId, not even an explicit nil
### GetInputValue

`func (o *ProductAddSpecificsInner) GetInputValue() string`

GetInputValue returns the InputValue field if non-nil, zero value otherwise.

### GetInputValueOk

`func (o *ProductAddSpecificsInner) GetInputValueOk() (*string, bool)`

GetInputValueOk returns a tuple with the InputValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValue

`func (o *ProductAddSpecificsInner) SetInputValue(v string)`

SetInputValue sets InputValue field to given value.

### HasInputValue

`func (o *ProductAddSpecificsInner) HasInputValue() bool`

HasInputValue returns a boolean if a field has been set.

### SetInputValueNil

`func (o *ProductAddSpecificsInner) SetInputValueNil(b bool)`

 SetInputValueNil sets the value for InputValue to be an explicit nil

### UnsetInputValue
`func (o *ProductAddSpecificsInner) UnsetInputValue()`

UnsetInputValue ensures that no value is present for InputValue, not even an explicit nil
### GetFoodDetails

`func (o *ProductAddSpecificsInner) GetFoodDetails() ProductAddSpecificsInnerFoodDetails`

GetFoodDetails returns the FoodDetails field if non-nil, zero value otherwise.

### GetFoodDetailsOk

`func (o *ProductAddSpecificsInner) GetFoodDetailsOk() (*ProductAddSpecificsInnerFoodDetails, bool)`

GetFoodDetailsOk returns a tuple with the FoodDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoodDetails

`func (o *ProductAddSpecificsInner) SetFoodDetails(v ProductAddSpecificsInnerFoodDetails)`

SetFoodDetails sets FoodDetails field to given value.

### HasFoodDetails

`func (o *ProductAddSpecificsInner) HasFoodDetails() bool`

HasFoodDetails returns a boolean if a field has been set.

### GetGroupProductsDetails

`func (o *ProductAddSpecificsInner) GetGroupProductsDetails() []ProductAddSpecificsInnerGroupProductsDetailsInner`

GetGroupProductsDetails returns the GroupProductsDetails field if non-nil, zero value otherwise.

### GetGroupProductsDetailsOk

`func (o *ProductAddSpecificsInner) GetGroupProductsDetailsOk() (*[]ProductAddSpecificsInnerGroupProductsDetailsInner, bool)`

GetGroupProductsDetailsOk returns a tuple with the GroupProductsDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupProductsDetails

`func (o *ProductAddSpecificsInner) SetGroupProductsDetails(v []ProductAddSpecificsInnerGroupProductsDetailsInner)`

SetGroupProductsDetails sets GroupProductsDetails field to given value.

### HasGroupProductsDetails

`func (o *ProductAddSpecificsInner) HasGroupProductsDetails() bool`

HasGroupProductsDetails returns a boolean if a field has been set.

### GetBookingDetails

`func (o *ProductAddSpecificsInner) GetBookingDetails() ProductAddSpecificsInnerBookingDetails`

GetBookingDetails returns the BookingDetails field if non-nil, zero value otherwise.

### GetBookingDetailsOk

`func (o *ProductAddSpecificsInner) GetBookingDetailsOk() (*ProductAddSpecificsInnerBookingDetails, bool)`

GetBookingDetailsOk returns a tuple with the BookingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingDetails

`func (o *ProductAddSpecificsInner) SetBookingDetails(v ProductAddSpecificsInnerBookingDetails)`

SetBookingDetails sets BookingDetails field to given value.

### HasBookingDetails

`func (o *ProductAddSpecificsInner) HasBookingDetails() bool`

HasBookingDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


