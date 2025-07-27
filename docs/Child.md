# Child

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **NullableString** |  | [optional] 
**Upc** | Pointer to **NullableString** |  | [optional] 
**Ean** | Pointer to **NullableString** |  | [optional] 
**Mpn** | Pointer to **NullableString** |  | [optional] 
**Gtin** | Pointer to **NullableString** |  | [optional] 
**Isbn** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**SeoUrl** | Pointer to **NullableString** |  | [optional] 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**CreatedTime** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedTime** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ShortDescription** | Pointer to **NullableString** |  | [optional] 
**FullDescription** | Pointer to **NullableString** |  | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) |  | [optional] 
**Combination** | Pointer to [**[]ProductChildItemCombination**](ProductChildItemCombination.md) |  | [optional] 
**DefaultPrice** | Pointer to **NullableFloat32** |  | [optional] 
**CostPrice** | Pointer to **NullableFloat32** |  | [optional] 
**ListPrice** | Pointer to **NullableFloat32** |  | [optional] 
**WholesalePrice** | Pointer to **NullableFloat32** |  | [optional] 
**AdvancedPrice** | Pointer to [**[]ProductAdvancedPrice**](ProductAdvancedPrice.md) |  | [optional] 
**TaxClassId** | Pointer to **NullableString** |  | [optional] 
**AvailForSale** | Pointer to **NullableBool** |  | [optional] 
**AllowBackorders** | Pointer to **NullableBool** |  | [optional] 
**InStock** | Pointer to **NullableBool** |  | [optional] 
**OnSale** | Pointer to **NullableBool** |  | [optional] 
**ManageStock** | Pointer to **NullableBool** |  | [optional] 
**InventoryLevel** | Pointer to **NullableFloat32** |  | [optional] 
**Inventory** | Pointer to [**[]ProductInventory**](ProductInventory.md) |  | [optional] 
**MinQuantity** | Pointer to **NullableFloat32** |  | [optional] 
**LowStockThreshold** | Pointer to **NullableFloat32** |  | [optional] 
**DefaultQtyInPack** | Pointer to **NullableFloat32** |  | [optional] 
**IsQtyInPackFixed** | Pointer to **NullableBool** |  | [optional] 
**WeightUnit** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **NullableFloat32** |  | [optional] 
**DimensionsUnit** | Pointer to **NullableString** |  | [optional] 
**Width** | Pointer to **NullableFloat32** |  | [optional] 
**Height** | Pointer to **NullableFloat32** |  | [optional] 
**Length** | Pointer to **NullableFloat32** |  | [optional] 
**MetaTitle** | Pointer to **NullableString** |  | [optional] 
**MetaDescription** | Pointer to **NullableString** |  | [optional] 
**MetaKeywords** | Pointer to **NullableString** |  | [optional] 
**Discounts** | Pointer to [**[]Discount**](Discount.md) |  | [optional] 
**IsVirtual** | Pointer to **NullableBool** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewChild

`func NewChild() *Child`

NewChild instantiates a new Child object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildWithDefaults

`func NewChildWithDefaults() *Child`

NewChildWithDefaults instantiates a new Child object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Child) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Child) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Child) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Child) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *Child) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Child) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Child) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Child) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetSku

`func (o *Child) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Child) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Child) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Child) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *Child) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *Child) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetUpc

`func (o *Child) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *Child) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *Child) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *Child) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### SetUpcNil

`func (o *Child) SetUpcNil(b bool)`

 SetUpcNil sets the value for Upc to be an explicit nil

### UnsetUpc
`func (o *Child) UnsetUpc()`

UnsetUpc ensures that no value is present for Upc, not even an explicit nil
### GetEan

`func (o *Child) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *Child) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *Child) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *Child) HasEan() bool`

HasEan returns a boolean if a field has been set.

### SetEanNil

`func (o *Child) SetEanNil(b bool)`

 SetEanNil sets the value for Ean to be an explicit nil

### UnsetEan
`func (o *Child) UnsetEan()`

UnsetEan ensures that no value is present for Ean, not even an explicit nil
### GetMpn

`func (o *Child) GetMpn() string`

GetMpn returns the Mpn field if non-nil, zero value otherwise.

### GetMpnOk

`func (o *Child) GetMpnOk() (*string, bool)`

GetMpnOk returns a tuple with the Mpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpn

`func (o *Child) SetMpn(v string)`

SetMpn sets Mpn field to given value.

### HasMpn

`func (o *Child) HasMpn() bool`

HasMpn returns a boolean if a field has been set.

### SetMpnNil

`func (o *Child) SetMpnNil(b bool)`

 SetMpnNil sets the value for Mpn to be an explicit nil

### UnsetMpn
`func (o *Child) UnsetMpn()`

UnsetMpn ensures that no value is present for Mpn, not even an explicit nil
### GetGtin

`func (o *Child) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *Child) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *Child) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *Child) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### SetGtinNil

`func (o *Child) SetGtinNil(b bool)`

 SetGtinNil sets the value for Gtin to be an explicit nil

### UnsetGtin
`func (o *Child) UnsetGtin()`

UnsetGtin ensures that no value is present for Gtin, not even an explicit nil
### GetIsbn

`func (o *Child) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *Child) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *Child) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *Child) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### SetIsbnNil

`func (o *Child) SetIsbnNil(b bool)`

 SetIsbnNil sets the value for Isbn to be an explicit nil

### UnsetIsbn
`func (o *Child) UnsetIsbn()`

UnsetIsbn ensures that no value is present for Isbn, not even an explicit nil
### GetUrl

`func (o *Child) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Child) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Child) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Child) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *Child) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Child) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSeoUrl

`func (o *Child) GetSeoUrl() string`

GetSeoUrl returns the SeoUrl field if non-nil, zero value otherwise.

### GetSeoUrlOk

`func (o *Child) GetSeoUrlOk() (*string, bool)`

GetSeoUrlOk returns a tuple with the SeoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoUrl

`func (o *Child) SetSeoUrl(v string)`

SetSeoUrl sets SeoUrl field to given value.

### HasSeoUrl

`func (o *Child) HasSeoUrl() bool`

HasSeoUrl returns a boolean if a field has been set.

### SetSeoUrlNil

`func (o *Child) SetSeoUrlNil(b bool)`

 SetSeoUrlNil sets the value for SeoUrl to be an explicit nil

### UnsetSeoUrl
`func (o *Child) UnsetSeoUrl()`

UnsetSeoUrl ensures that no value is present for SeoUrl, not even an explicit nil
### GetSortOrder

`func (o *Child) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Child) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Child) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Child) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *Child) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *Child) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetCreatedTime

`func (o *Child) GetCreatedTime() A2CDateTime`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Child) GetCreatedTimeOk() (*A2CDateTime, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Child) SetCreatedTime(v A2CDateTime)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Child) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *Child) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *Child) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetModifiedTime

`func (o *Child) GetModifiedTime() A2CDateTime`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Child) GetModifiedTimeOk() (*A2CDateTime, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Child) SetModifiedTime(v A2CDateTime)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Child) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### SetModifiedTimeNil

`func (o *Child) SetModifiedTimeNil(b bool)`

 SetModifiedTimeNil sets the value for ModifiedTime to be an explicit nil

### UnsetModifiedTime
`func (o *Child) UnsetModifiedTime()`

UnsetModifiedTime ensures that no value is present for ModifiedTime, not even an explicit nil
### GetName

`func (o *Child) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Child) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Child) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Child) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Child) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Child) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShortDescription

`func (o *Child) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *Child) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *Child) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *Child) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### SetShortDescriptionNil

`func (o *Child) SetShortDescriptionNil(b bool)`

 SetShortDescriptionNil sets the value for ShortDescription to be an explicit nil

### UnsetShortDescription
`func (o *Child) UnsetShortDescription()`

UnsetShortDescription ensures that no value is present for ShortDescription, not even an explicit nil
### GetFullDescription

`func (o *Child) GetFullDescription() string`

GetFullDescription returns the FullDescription field if non-nil, zero value otherwise.

### GetFullDescriptionOk

`func (o *Child) GetFullDescriptionOk() (*string, bool)`

GetFullDescriptionOk returns a tuple with the FullDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDescription

`func (o *Child) SetFullDescription(v string)`

SetFullDescription sets FullDescription field to given value.

### HasFullDescription

`func (o *Child) HasFullDescription() bool`

HasFullDescription returns a boolean if a field has been set.

### SetFullDescriptionNil

`func (o *Child) SetFullDescriptionNil(b bool)`

 SetFullDescriptionNil sets the value for FullDescription to be an explicit nil

### UnsetFullDescription
`func (o *Child) UnsetFullDescription()`

UnsetFullDescription ensures that no value is present for FullDescription, not even an explicit nil
### GetImages

`func (o *Child) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Child) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Child) SetImages(v []Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *Child) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetCombination

`func (o *Child) GetCombination() []ProductChildItemCombination`

GetCombination returns the Combination field if non-nil, zero value otherwise.

### GetCombinationOk

`func (o *Child) GetCombinationOk() (*[]ProductChildItemCombination, bool)`

GetCombinationOk returns a tuple with the Combination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombination

`func (o *Child) SetCombination(v []ProductChildItemCombination)`

SetCombination sets Combination field to given value.

### HasCombination

`func (o *Child) HasCombination() bool`

HasCombination returns a boolean if a field has been set.

### GetDefaultPrice

`func (o *Child) GetDefaultPrice() float32`

GetDefaultPrice returns the DefaultPrice field if non-nil, zero value otherwise.

### GetDefaultPriceOk

`func (o *Child) GetDefaultPriceOk() (*float32, bool)`

GetDefaultPriceOk returns a tuple with the DefaultPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrice

`func (o *Child) SetDefaultPrice(v float32)`

SetDefaultPrice sets DefaultPrice field to given value.

### HasDefaultPrice

`func (o *Child) HasDefaultPrice() bool`

HasDefaultPrice returns a boolean if a field has been set.

### SetDefaultPriceNil

`func (o *Child) SetDefaultPriceNil(b bool)`

 SetDefaultPriceNil sets the value for DefaultPrice to be an explicit nil

### UnsetDefaultPrice
`func (o *Child) UnsetDefaultPrice()`

UnsetDefaultPrice ensures that no value is present for DefaultPrice, not even an explicit nil
### GetCostPrice

`func (o *Child) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *Child) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *Child) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *Child) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### SetCostPriceNil

`func (o *Child) SetCostPriceNil(b bool)`

 SetCostPriceNil sets the value for CostPrice to be an explicit nil

### UnsetCostPrice
`func (o *Child) UnsetCostPrice()`

UnsetCostPrice ensures that no value is present for CostPrice, not even an explicit nil
### GetListPrice

`func (o *Child) GetListPrice() float32`

GetListPrice returns the ListPrice field if non-nil, zero value otherwise.

### GetListPriceOk

`func (o *Child) GetListPriceOk() (*float32, bool)`

GetListPriceOk returns a tuple with the ListPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPrice

`func (o *Child) SetListPrice(v float32)`

SetListPrice sets ListPrice field to given value.

### HasListPrice

`func (o *Child) HasListPrice() bool`

HasListPrice returns a boolean if a field has been set.

### SetListPriceNil

`func (o *Child) SetListPriceNil(b bool)`

 SetListPriceNil sets the value for ListPrice to be an explicit nil

### UnsetListPrice
`func (o *Child) UnsetListPrice()`

UnsetListPrice ensures that no value is present for ListPrice, not even an explicit nil
### GetWholesalePrice

`func (o *Child) GetWholesalePrice() float32`

GetWholesalePrice returns the WholesalePrice field if non-nil, zero value otherwise.

### GetWholesalePriceOk

`func (o *Child) GetWholesalePriceOk() (*float32, bool)`

GetWholesalePriceOk returns a tuple with the WholesalePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholesalePrice

`func (o *Child) SetWholesalePrice(v float32)`

SetWholesalePrice sets WholesalePrice field to given value.

### HasWholesalePrice

`func (o *Child) HasWholesalePrice() bool`

HasWholesalePrice returns a boolean if a field has been set.

### SetWholesalePriceNil

`func (o *Child) SetWholesalePriceNil(b bool)`

 SetWholesalePriceNil sets the value for WholesalePrice to be an explicit nil

### UnsetWholesalePrice
`func (o *Child) UnsetWholesalePrice()`

UnsetWholesalePrice ensures that no value is present for WholesalePrice, not even an explicit nil
### GetAdvancedPrice

`func (o *Child) GetAdvancedPrice() []ProductAdvancedPrice`

GetAdvancedPrice returns the AdvancedPrice field if non-nil, zero value otherwise.

### GetAdvancedPriceOk

`func (o *Child) GetAdvancedPriceOk() (*[]ProductAdvancedPrice, bool)`

GetAdvancedPriceOk returns a tuple with the AdvancedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPrice

`func (o *Child) SetAdvancedPrice(v []ProductAdvancedPrice)`

SetAdvancedPrice sets AdvancedPrice field to given value.

### HasAdvancedPrice

`func (o *Child) HasAdvancedPrice() bool`

HasAdvancedPrice returns a boolean if a field has been set.

### GetTaxClassId

`func (o *Child) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *Child) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *Child) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *Child) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### SetTaxClassIdNil

`func (o *Child) SetTaxClassIdNil(b bool)`

 SetTaxClassIdNil sets the value for TaxClassId to be an explicit nil

### UnsetTaxClassId
`func (o *Child) UnsetTaxClassId()`

UnsetTaxClassId ensures that no value is present for TaxClassId, not even an explicit nil
### GetAvailForSale

`func (o *Child) GetAvailForSale() bool`

GetAvailForSale returns the AvailForSale field if non-nil, zero value otherwise.

### GetAvailForSaleOk

`func (o *Child) GetAvailForSaleOk() (*bool, bool)`

GetAvailForSaleOk returns a tuple with the AvailForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailForSale

`func (o *Child) SetAvailForSale(v bool)`

SetAvailForSale sets AvailForSale field to given value.

### HasAvailForSale

`func (o *Child) HasAvailForSale() bool`

HasAvailForSale returns a boolean if a field has been set.

### SetAvailForSaleNil

`func (o *Child) SetAvailForSaleNil(b bool)`

 SetAvailForSaleNil sets the value for AvailForSale to be an explicit nil

### UnsetAvailForSale
`func (o *Child) UnsetAvailForSale()`

UnsetAvailForSale ensures that no value is present for AvailForSale, not even an explicit nil
### GetAllowBackorders

`func (o *Child) GetAllowBackorders() bool`

GetAllowBackorders returns the AllowBackorders field if non-nil, zero value otherwise.

### GetAllowBackordersOk

`func (o *Child) GetAllowBackordersOk() (*bool, bool)`

GetAllowBackordersOk returns a tuple with the AllowBackorders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBackorders

`func (o *Child) SetAllowBackorders(v bool)`

SetAllowBackorders sets AllowBackorders field to given value.

### HasAllowBackorders

`func (o *Child) HasAllowBackorders() bool`

HasAllowBackorders returns a boolean if a field has been set.

### SetAllowBackordersNil

`func (o *Child) SetAllowBackordersNil(b bool)`

 SetAllowBackordersNil sets the value for AllowBackorders to be an explicit nil

### UnsetAllowBackorders
`func (o *Child) UnsetAllowBackorders()`

UnsetAllowBackorders ensures that no value is present for AllowBackorders, not even an explicit nil
### GetInStock

`func (o *Child) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *Child) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *Child) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *Child) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### SetInStockNil

`func (o *Child) SetInStockNil(b bool)`

 SetInStockNil sets the value for InStock to be an explicit nil

### UnsetInStock
`func (o *Child) UnsetInStock()`

UnsetInStock ensures that no value is present for InStock, not even an explicit nil
### GetOnSale

`func (o *Child) GetOnSale() bool`

GetOnSale returns the OnSale field if non-nil, zero value otherwise.

### GetOnSaleOk

`func (o *Child) GetOnSaleOk() (*bool, bool)`

GetOnSaleOk returns a tuple with the OnSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSale

`func (o *Child) SetOnSale(v bool)`

SetOnSale sets OnSale field to given value.

### HasOnSale

`func (o *Child) HasOnSale() bool`

HasOnSale returns a boolean if a field has been set.

### SetOnSaleNil

`func (o *Child) SetOnSaleNil(b bool)`

 SetOnSaleNil sets the value for OnSale to be an explicit nil

### UnsetOnSale
`func (o *Child) UnsetOnSale()`

UnsetOnSale ensures that no value is present for OnSale, not even an explicit nil
### GetManageStock

`func (o *Child) GetManageStock() bool`

GetManageStock returns the ManageStock field if non-nil, zero value otherwise.

### GetManageStockOk

`func (o *Child) GetManageStockOk() (*bool, bool)`

GetManageStockOk returns a tuple with the ManageStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStock

`func (o *Child) SetManageStock(v bool)`

SetManageStock sets ManageStock field to given value.

### HasManageStock

`func (o *Child) HasManageStock() bool`

HasManageStock returns a boolean if a field has been set.

### SetManageStockNil

`func (o *Child) SetManageStockNil(b bool)`

 SetManageStockNil sets the value for ManageStock to be an explicit nil

### UnsetManageStock
`func (o *Child) UnsetManageStock()`

UnsetManageStock ensures that no value is present for ManageStock, not even an explicit nil
### GetInventoryLevel

`func (o *Child) GetInventoryLevel() float32`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *Child) GetInventoryLevelOk() (*float32, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *Child) SetInventoryLevel(v float32)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *Child) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### SetInventoryLevelNil

`func (o *Child) SetInventoryLevelNil(b bool)`

 SetInventoryLevelNil sets the value for InventoryLevel to be an explicit nil

### UnsetInventoryLevel
`func (o *Child) UnsetInventoryLevel()`

UnsetInventoryLevel ensures that no value is present for InventoryLevel, not even an explicit nil
### GetInventory

`func (o *Child) GetInventory() []ProductInventory`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *Child) GetInventoryOk() (*[]ProductInventory, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *Child) SetInventory(v []ProductInventory)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *Child) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetMinQuantity

`func (o *Child) GetMinQuantity() float32`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *Child) GetMinQuantityOk() (*float32, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *Child) SetMinQuantity(v float32)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *Child) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### SetMinQuantityNil

`func (o *Child) SetMinQuantityNil(b bool)`

 SetMinQuantityNil sets the value for MinQuantity to be an explicit nil

### UnsetMinQuantity
`func (o *Child) UnsetMinQuantity()`

UnsetMinQuantity ensures that no value is present for MinQuantity, not even an explicit nil
### GetLowStockThreshold

`func (o *Child) GetLowStockThreshold() float32`

GetLowStockThreshold returns the LowStockThreshold field if non-nil, zero value otherwise.

### GetLowStockThresholdOk

`func (o *Child) GetLowStockThresholdOk() (*float32, bool)`

GetLowStockThresholdOk returns a tuple with the LowStockThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowStockThreshold

`func (o *Child) SetLowStockThreshold(v float32)`

SetLowStockThreshold sets LowStockThreshold field to given value.

### HasLowStockThreshold

`func (o *Child) HasLowStockThreshold() bool`

HasLowStockThreshold returns a boolean if a field has been set.

### SetLowStockThresholdNil

`func (o *Child) SetLowStockThresholdNil(b bool)`

 SetLowStockThresholdNil sets the value for LowStockThreshold to be an explicit nil

### UnsetLowStockThreshold
`func (o *Child) UnsetLowStockThreshold()`

UnsetLowStockThreshold ensures that no value is present for LowStockThreshold, not even an explicit nil
### GetDefaultQtyInPack

`func (o *Child) GetDefaultQtyInPack() float32`

GetDefaultQtyInPack returns the DefaultQtyInPack field if non-nil, zero value otherwise.

### GetDefaultQtyInPackOk

`func (o *Child) GetDefaultQtyInPackOk() (*float32, bool)`

GetDefaultQtyInPackOk returns a tuple with the DefaultQtyInPack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQtyInPack

`func (o *Child) SetDefaultQtyInPack(v float32)`

SetDefaultQtyInPack sets DefaultQtyInPack field to given value.

### HasDefaultQtyInPack

`func (o *Child) HasDefaultQtyInPack() bool`

HasDefaultQtyInPack returns a boolean if a field has been set.

### SetDefaultQtyInPackNil

`func (o *Child) SetDefaultQtyInPackNil(b bool)`

 SetDefaultQtyInPackNil sets the value for DefaultQtyInPack to be an explicit nil

### UnsetDefaultQtyInPack
`func (o *Child) UnsetDefaultQtyInPack()`

UnsetDefaultQtyInPack ensures that no value is present for DefaultQtyInPack, not even an explicit nil
### GetIsQtyInPackFixed

`func (o *Child) GetIsQtyInPackFixed() bool`

GetIsQtyInPackFixed returns the IsQtyInPackFixed field if non-nil, zero value otherwise.

### GetIsQtyInPackFixedOk

`func (o *Child) GetIsQtyInPackFixedOk() (*bool, bool)`

GetIsQtyInPackFixedOk returns a tuple with the IsQtyInPackFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQtyInPackFixed

`func (o *Child) SetIsQtyInPackFixed(v bool)`

SetIsQtyInPackFixed sets IsQtyInPackFixed field to given value.

### HasIsQtyInPackFixed

`func (o *Child) HasIsQtyInPackFixed() bool`

HasIsQtyInPackFixed returns a boolean if a field has been set.

### SetIsQtyInPackFixedNil

`func (o *Child) SetIsQtyInPackFixedNil(b bool)`

 SetIsQtyInPackFixedNil sets the value for IsQtyInPackFixed to be an explicit nil

### UnsetIsQtyInPackFixed
`func (o *Child) UnsetIsQtyInPackFixed()`

UnsetIsQtyInPackFixed ensures that no value is present for IsQtyInPackFixed, not even an explicit nil
### GetWeightUnit

`func (o *Child) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *Child) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *Child) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *Child) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *Child) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *Child) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetWeight

`func (o *Child) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Child) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Child) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Child) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *Child) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *Child) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetDimensionsUnit

`func (o *Child) GetDimensionsUnit() string`

GetDimensionsUnit returns the DimensionsUnit field if non-nil, zero value otherwise.

### GetDimensionsUnitOk

`func (o *Child) GetDimensionsUnitOk() (*string, bool)`

GetDimensionsUnitOk returns a tuple with the DimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsUnit

`func (o *Child) SetDimensionsUnit(v string)`

SetDimensionsUnit sets DimensionsUnit field to given value.

### HasDimensionsUnit

`func (o *Child) HasDimensionsUnit() bool`

HasDimensionsUnit returns a boolean if a field has been set.

### SetDimensionsUnitNil

`func (o *Child) SetDimensionsUnitNil(b bool)`

 SetDimensionsUnitNil sets the value for DimensionsUnit to be an explicit nil

### UnsetDimensionsUnit
`func (o *Child) UnsetDimensionsUnit()`

UnsetDimensionsUnit ensures that no value is present for DimensionsUnit, not even an explicit nil
### GetWidth

`func (o *Child) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Child) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Child) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Child) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *Child) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *Child) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *Child) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Child) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Child) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Child) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *Child) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *Child) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetLength

`func (o *Child) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Child) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Child) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *Child) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *Child) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *Child) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetMetaTitle

`func (o *Child) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *Child) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *Child) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *Child) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### SetMetaTitleNil

`func (o *Child) SetMetaTitleNil(b bool)`

 SetMetaTitleNil sets the value for MetaTitle to be an explicit nil

### UnsetMetaTitle
`func (o *Child) UnsetMetaTitle()`

UnsetMetaTitle ensures that no value is present for MetaTitle, not even an explicit nil
### GetMetaDescription

`func (o *Child) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *Child) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *Child) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *Child) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### SetMetaDescriptionNil

`func (o *Child) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *Child) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
### GetMetaKeywords

`func (o *Child) GetMetaKeywords() string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *Child) GetMetaKeywordsOk() (*string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *Child) SetMetaKeywords(v string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *Child) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### SetMetaKeywordsNil

`func (o *Child) SetMetaKeywordsNil(b bool)`

 SetMetaKeywordsNil sets the value for MetaKeywords to be an explicit nil

### UnsetMetaKeywords
`func (o *Child) UnsetMetaKeywords()`

UnsetMetaKeywords ensures that no value is present for MetaKeywords, not even an explicit nil
### GetDiscounts

`func (o *Child) GetDiscounts() []Discount`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *Child) GetDiscountsOk() (*[]Discount, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *Child) SetDiscounts(v []Discount)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *Child) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetIsVirtual

`func (o *Child) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *Child) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *Child) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *Child) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### SetIsVirtualNil

`func (o *Child) SetIsVirtualNil(b bool)`

 SetIsVirtualNil sets the value for IsVirtual to be an explicit nil

### UnsetIsVirtual
`func (o *Child) UnsetIsVirtual()`

UnsetIsVirtual ensures that no value is present for IsVirtual, not even an explicit nil
### GetAdditionalFields

`func (o *Child) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Child) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Child) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Child) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *Child) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *Child) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *Child) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Child) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Child) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Child) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *Child) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *Child) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


