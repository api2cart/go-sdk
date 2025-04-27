# OrderAddOrderItemInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderItemId** | **string** | Defines orders specified by order item id | 
**OrderItemName** | **string** | Defines orders specified by order item name | 
**OrderItemModel** | Pointer to **string** | Defines orders specified by order item model | [optional] 
**OrderItemPrice** | **float32** | Defines orders specified by order item price | 
**OrderItemQuantity** | **int32** | Defines orders specified by order item quantity | 
**OrderItemWeight** | Pointer to **float32** | Defines orders specified by order item weight | [optional] 
**OrderItemVariantId** | Pointer to **string** | Ordered product variant. Where x is order item ID | [optional] 
**OrderItemTax** | Pointer to **float32** | Percentage of tax for product order | [optional] [default to 0]
**OrderItemPriceIncludesTax** | Pointer to **bool** | Defines if item price includes tax | [optional] [default to false]
**OrderItemParent** | Pointer to **int32** | Index of the parent grouped/bundle product | [optional] 
**OrderItemParentOptionName** | Pointer to **string** | Option name of the parent grouped/bundle product | [optional] 
**OrderItemAllowRefundItemsSeparately** | Pointer to **bool** | Indicates whether subitems of the grouped/bundle product can be refunded separately | [optional] 
**OrderItemAllowShipItemsSeparately** | Pointer to **bool** | Indicates whether subitems of the grouped/bundle product can be shipped separately | [optional] 
**OrderItemOption** | Pointer to [**[]OrderAddOrderItemInnerOrderItemOptionInner**](OrderAddOrderItemInnerOrderItemOptionInner.md) |  | [optional] 
**OrderItemProperty** | Pointer to [**[]OrderAddOrderItemInnerOrderItemPropertyInner**](OrderAddOrderItemInnerOrderItemPropertyInner.md) |  | [optional] 

## Methods

### NewOrderAddOrderItemInner

`func NewOrderAddOrderItemInner(orderItemId string, orderItemName string, orderItemPrice float32, orderItemQuantity int32, ) *OrderAddOrderItemInner`

NewOrderAddOrderItemInner instantiates a new OrderAddOrderItemInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAddOrderItemInnerWithDefaults

`func NewOrderAddOrderItemInnerWithDefaults() *OrderAddOrderItemInner`

NewOrderAddOrderItemInnerWithDefaults instantiates a new OrderAddOrderItemInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderItemId

`func (o *OrderAddOrderItemInner) GetOrderItemId() string`

GetOrderItemId returns the OrderItemId field if non-nil, zero value otherwise.

### GetOrderItemIdOk

`func (o *OrderAddOrderItemInner) GetOrderItemIdOk() (*string, bool)`

GetOrderItemIdOk returns a tuple with the OrderItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemId

`func (o *OrderAddOrderItemInner) SetOrderItemId(v string)`

SetOrderItemId sets OrderItemId field to given value.


### GetOrderItemName

`func (o *OrderAddOrderItemInner) GetOrderItemName() string`

GetOrderItemName returns the OrderItemName field if non-nil, zero value otherwise.

### GetOrderItemNameOk

`func (o *OrderAddOrderItemInner) GetOrderItemNameOk() (*string, bool)`

GetOrderItemNameOk returns a tuple with the OrderItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemName

`func (o *OrderAddOrderItemInner) SetOrderItemName(v string)`

SetOrderItemName sets OrderItemName field to given value.


### GetOrderItemModel

`func (o *OrderAddOrderItemInner) GetOrderItemModel() string`

GetOrderItemModel returns the OrderItemModel field if non-nil, zero value otherwise.

### GetOrderItemModelOk

`func (o *OrderAddOrderItemInner) GetOrderItemModelOk() (*string, bool)`

GetOrderItemModelOk returns a tuple with the OrderItemModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemModel

`func (o *OrderAddOrderItemInner) SetOrderItemModel(v string)`

SetOrderItemModel sets OrderItemModel field to given value.

### HasOrderItemModel

`func (o *OrderAddOrderItemInner) HasOrderItemModel() bool`

HasOrderItemModel returns a boolean if a field has been set.

### GetOrderItemPrice

`func (o *OrderAddOrderItemInner) GetOrderItemPrice() float32`

GetOrderItemPrice returns the OrderItemPrice field if non-nil, zero value otherwise.

### GetOrderItemPriceOk

`func (o *OrderAddOrderItemInner) GetOrderItemPriceOk() (*float32, bool)`

GetOrderItemPriceOk returns a tuple with the OrderItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemPrice

`func (o *OrderAddOrderItemInner) SetOrderItemPrice(v float32)`

SetOrderItemPrice sets OrderItemPrice field to given value.


### GetOrderItemQuantity

`func (o *OrderAddOrderItemInner) GetOrderItemQuantity() int32`

GetOrderItemQuantity returns the OrderItemQuantity field if non-nil, zero value otherwise.

### GetOrderItemQuantityOk

`func (o *OrderAddOrderItemInner) GetOrderItemQuantityOk() (*int32, bool)`

GetOrderItemQuantityOk returns a tuple with the OrderItemQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemQuantity

`func (o *OrderAddOrderItemInner) SetOrderItemQuantity(v int32)`

SetOrderItemQuantity sets OrderItemQuantity field to given value.


### GetOrderItemWeight

`func (o *OrderAddOrderItemInner) GetOrderItemWeight() float32`

GetOrderItemWeight returns the OrderItemWeight field if non-nil, zero value otherwise.

### GetOrderItemWeightOk

`func (o *OrderAddOrderItemInner) GetOrderItemWeightOk() (*float32, bool)`

GetOrderItemWeightOk returns a tuple with the OrderItemWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemWeight

`func (o *OrderAddOrderItemInner) SetOrderItemWeight(v float32)`

SetOrderItemWeight sets OrderItemWeight field to given value.

### HasOrderItemWeight

`func (o *OrderAddOrderItemInner) HasOrderItemWeight() bool`

HasOrderItemWeight returns a boolean if a field has been set.

### GetOrderItemVariantId

`func (o *OrderAddOrderItemInner) GetOrderItemVariantId() string`

GetOrderItemVariantId returns the OrderItemVariantId field if non-nil, zero value otherwise.

### GetOrderItemVariantIdOk

`func (o *OrderAddOrderItemInner) GetOrderItemVariantIdOk() (*string, bool)`

GetOrderItemVariantIdOk returns a tuple with the OrderItemVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemVariantId

`func (o *OrderAddOrderItemInner) SetOrderItemVariantId(v string)`

SetOrderItemVariantId sets OrderItemVariantId field to given value.

### HasOrderItemVariantId

`func (o *OrderAddOrderItemInner) HasOrderItemVariantId() bool`

HasOrderItemVariantId returns a boolean if a field has been set.

### GetOrderItemTax

`func (o *OrderAddOrderItemInner) GetOrderItemTax() float32`

GetOrderItemTax returns the OrderItemTax field if non-nil, zero value otherwise.

### GetOrderItemTaxOk

`func (o *OrderAddOrderItemInner) GetOrderItemTaxOk() (*float32, bool)`

GetOrderItemTaxOk returns a tuple with the OrderItemTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemTax

`func (o *OrderAddOrderItemInner) SetOrderItemTax(v float32)`

SetOrderItemTax sets OrderItemTax field to given value.

### HasOrderItemTax

`func (o *OrderAddOrderItemInner) HasOrderItemTax() bool`

HasOrderItemTax returns a boolean if a field has been set.

### GetOrderItemPriceIncludesTax

`func (o *OrderAddOrderItemInner) GetOrderItemPriceIncludesTax() bool`

GetOrderItemPriceIncludesTax returns the OrderItemPriceIncludesTax field if non-nil, zero value otherwise.

### GetOrderItemPriceIncludesTaxOk

`func (o *OrderAddOrderItemInner) GetOrderItemPriceIncludesTaxOk() (*bool, bool)`

GetOrderItemPriceIncludesTaxOk returns a tuple with the OrderItemPriceIncludesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemPriceIncludesTax

`func (o *OrderAddOrderItemInner) SetOrderItemPriceIncludesTax(v bool)`

SetOrderItemPriceIncludesTax sets OrderItemPriceIncludesTax field to given value.

### HasOrderItemPriceIncludesTax

`func (o *OrderAddOrderItemInner) HasOrderItemPriceIncludesTax() bool`

HasOrderItemPriceIncludesTax returns a boolean if a field has been set.

### GetOrderItemParent

`func (o *OrderAddOrderItemInner) GetOrderItemParent() int32`

GetOrderItemParent returns the OrderItemParent field if non-nil, zero value otherwise.

### GetOrderItemParentOk

`func (o *OrderAddOrderItemInner) GetOrderItemParentOk() (*int32, bool)`

GetOrderItemParentOk returns a tuple with the OrderItemParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemParent

`func (o *OrderAddOrderItemInner) SetOrderItemParent(v int32)`

SetOrderItemParent sets OrderItemParent field to given value.

### HasOrderItemParent

`func (o *OrderAddOrderItemInner) HasOrderItemParent() bool`

HasOrderItemParent returns a boolean if a field has been set.

### GetOrderItemParentOptionName

`func (o *OrderAddOrderItemInner) GetOrderItemParentOptionName() string`

GetOrderItemParentOptionName returns the OrderItemParentOptionName field if non-nil, zero value otherwise.

### GetOrderItemParentOptionNameOk

`func (o *OrderAddOrderItemInner) GetOrderItemParentOptionNameOk() (*string, bool)`

GetOrderItemParentOptionNameOk returns a tuple with the OrderItemParentOptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemParentOptionName

`func (o *OrderAddOrderItemInner) SetOrderItemParentOptionName(v string)`

SetOrderItemParentOptionName sets OrderItemParentOptionName field to given value.

### HasOrderItemParentOptionName

`func (o *OrderAddOrderItemInner) HasOrderItemParentOptionName() bool`

HasOrderItemParentOptionName returns a boolean if a field has been set.

### GetOrderItemAllowRefundItemsSeparately

`func (o *OrderAddOrderItemInner) GetOrderItemAllowRefundItemsSeparately() bool`

GetOrderItemAllowRefundItemsSeparately returns the OrderItemAllowRefundItemsSeparately field if non-nil, zero value otherwise.

### GetOrderItemAllowRefundItemsSeparatelyOk

`func (o *OrderAddOrderItemInner) GetOrderItemAllowRefundItemsSeparatelyOk() (*bool, bool)`

GetOrderItemAllowRefundItemsSeparatelyOk returns a tuple with the OrderItemAllowRefundItemsSeparately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemAllowRefundItemsSeparately

`func (o *OrderAddOrderItemInner) SetOrderItemAllowRefundItemsSeparately(v bool)`

SetOrderItemAllowRefundItemsSeparately sets OrderItemAllowRefundItemsSeparately field to given value.

### HasOrderItemAllowRefundItemsSeparately

`func (o *OrderAddOrderItemInner) HasOrderItemAllowRefundItemsSeparately() bool`

HasOrderItemAllowRefundItemsSeparately returns a boolean if a field has been set.

### GetOrderItemAllowShipItemsSeparately

`func (o *OrderAddOrderItemInner) GetOrderItemAllowShipItemsSeparately() bool`

GetOrderItemAllowShipItemsSeparately returns the OrderItemAllowShipItemsSeparately field if non-nil, zero value otherwise.

### GetOrderItemAllowShipItemsSeparatelyOk

`func (o *OrderAddOrderItemInner) GetOrderItemAllowShipItemsSeparatelyOk() (*bool, bool)`

GetOrderItemAllowShipItemsSeparatelyOk returns a tuple with the OrderItemAllowShipItemsSeparately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemAllowShipItemsSeparately

`func (o *OrderAddOrderItemInner) SetOrderItemAllowShipItemsSeparately(v bool)`

SetOrderItemAllowShipItemsSeparately sets OrderItemAllowShipItemsSeparately field to given value.

### HasOrderItemAllowShipItemsSeparately

`func (o *OrderAddOrderItemInner) HasOrderItemAllowShipItemsSeparately() bool`

HasOrderItemAllowShipItemsSeparately returns a boolean if a field has been set.

### GetOrderItemOption

`func (o *OrderAddOrderItemInner) GetOrderItemOption() []OrderAddOrderItemInnerOrderItemOptionInner`

GetOrderItemOption returns the OrderItemOption field if non-nil, zero value otherwise.

### GetOrderItemOptionOk

`func (o *OrderAddOrderItemInner) GetOrderItemOptionOk() (*[]OrderAddOrderItemInnerOrderItemOptionInner, bool)`

GetOrderItemOptionOk returns a tuple with the OrderItemOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemOption

`func (o *OrderAddOrderItemInner) SetOrderItemOption(v []OrderAddOrderItemInnerOrderItemOptionInner)`

SetOrderItemOption sets OrderItemOption field to given value.

### HasOrderItemOption

`func (o *OrderAddOrderItemInner) HasOrderItemOption() bool`

HasOrderItemOption returns a boolean if a field has been set.

### GetOrderItemProperty

`func (o *OrderAddOrderItemInner) GetOrderItemProperty() []OrderAddOrderItemInnerOrderItemPropertyInner`

GetOrderItemProperty returns the OrderItemProperty field if non-nil, zero value otherwise.

### GetOrderItemPropertyOk

`func (o *OrderAddOrderItemInner) GetOrderItemPropertyOk() (*[]OrderAddOrderItemInnerOrderItemPropertyInner, bool)`

GetOrderItemPropertyOk returns a tuple with the OrderItemProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemProperty

`func (o *OrderAddOrderItemInner) SetOrderItemProperty(v []OrderAddOrderItemInnerOrderItemPropertyInner)`

SetOrderItemProperty sets OrderItemProperty field to given value.

### HasOrderItemProperty

`func (o *OrderAddOrderItemInner) HasOrderItemProperty() bool`

HasOrderItemProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


