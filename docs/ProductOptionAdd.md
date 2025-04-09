# ProductOptionAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Defines option&#39;s name | 
**Type** | **string** | Defines option&#39;s type that has to be added | 
**ProductId** | Pointer to **string** | Defines product id where the option should be added | [optional] 
**DefaultOptionValue** | Pointer to **string** | Defines default option value that has to be added | [optional] 
**OptionValues** | Pointer to **string** | Defines option values that has to be added | [optional] 
**Description** | Pointer to **string** | Defines option&#39;s description | [optional] 
**Avail** | Pointer to **bool** | Defines whether the option is available | [optional] [default to true]
**SortOrder** | Pointer to **int32** | Sort number in the list | [optional] [default to 0]
**Required** | Pointer to **bool** | Defines if the option is required | [optional] [default to false]
**Values** | Pointer to [**[]ProductOptionAddValuesInner**](ProductOptionAddValuesInner.md) | An array of option values.&lt;/b&gt; | [optional] 
**ClearCache** | Pointer to **bool** | Is cache clear required | [optional] [default to true]

## Methods

### NewProductOptionAdd

`func NewProductOptionAdd(name string, type_ string, ) *ProductOptionAdd`

NewProductOptionAdd instantiates a new ProductOptionAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductOptionAddWithDefaults

`func NewProductOptionAddWithDefaults() *ProductOptionAdd`

NewProductOptionAddWithDefaults instantiates a new ProductOptionAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductOptionAdd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductOptionAdd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductOptionAdd) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ProductOptionAdd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductOptionAdd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductOptionAdd) SetType(v string)`

SetType sets Type field to given value.


### GetProductId

`func (o *ProductOptionAdd) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductOptionAdd) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductOptionAdd) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductOptionAdd) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetDefaultOptionValue

`func (o *ProductOptionAdd) GetDefaultOptionValue() string`

GetDefaultOptionValue returns the DefaultOptionValue field if non-nil, zero value otherwise.

### GetDefaultOptionValueOk

`func (o *ProductOptionAdd) GetDefaultOptionValueOk() (*string, bool)`

GetDefaultOptionValueOk returns a tuple with the DefaultOptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOptionValue

`func (o *ProductOptionAdd) SetDefaultOptionValue(v string)`

SetDefaultOptionValue sets DefaultOptionValue field to given value.

### HasDefaultOptionValue

`func (o *ProductOptionAdd) HasDefaultOptionValue() bool`

HasDefaultOptionValue returns a boolean if a field has been set.

### GetOptionValues

`func (o *ProductOptionAdd) GetOptionValues() string`

GetOptionValues returns the OptionValues field if non-nil, zero value otherwise.

### GetOptionValuesOk

`func (o *ProductOptionAdd) GetOptionValuesOk() (*string, bool)`

GetOptionValuesOk returns a tuple with the OptionValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionValues

`func (o *ProductOptionAdd) SetOptionValues(v string)`

SetOptionValues sets OptionValues field to given value.

### HasOptionValues

`func (o *ProductOptionAdd) HasOptionValues() bool`

HasOptionValues returns a boolean if a field has been set.

### GetDescription

`func (o *ProductOptionAdd) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductOptionAdd) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductOptionAdd) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductOptionAdd) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAvail

`func (o *ProductOptionAdd) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *ProductOptionAdd) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *ProductOptionAdd) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *ProductOptionAdd) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetSortOrder

`func (o *ProductOptionAdd) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ProductOptionAdd) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ProductOptionAdd) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ProductOptionAdd) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetRequired

`func (o *ProductOptionAdd) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ProductOptionAdd) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ProductOptionAdd) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ProductOptionAdd) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetValues

`func (o *ProductOptionAdd) GetValues() []ProductOptionAddValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ProductOptionAdd) GetValuesOk() (*[]ProductOptionAddValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ProductOptionAdd) SetValues(v []ProductOptionAddValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *ProductOptionAdd) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetClearCache

`func (o *ProductOptionAdd) GetClearCache() bool`

GetClearCache returns the ClearCache field if non-nil, zero value otherwise.

### GetClearCacheOk

`func (o *ProductOptionAdd) GetClearCacheOk() (*bool, bool)`

GetClearCacheOk returns a tuple with the ClearCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCache

`func (o *ProductOptionAdd) SetClearCache(v bool)`

SetClearCache sets ClearCache field to given value.

### HasClearCache

`func (o *ProductOptionAdd) HasClearCache() bool`

HasClearCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


