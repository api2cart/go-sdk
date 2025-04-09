# Child

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Upc** | Pointer to **string** |  | [optional] 
**Ean** | Pointer to **string** |  | [optional] 
**Mpn** | Pointer to **string** |  | [optional] 
**Gtin** | Pointer to **string** |  | [optional] 
**Isbn** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**SeoUrl** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**CreatedTime** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedTime** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**FullDescription** | Pointer to **string** |  | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) |  | [optional] 
**Combination** | Pointer to [**[]ProductChildItemCombination**](ProductChildItemCombination.md) |  | [optional] 
**DefaultPrice** | Pointer to **float32** |  | [optional] 
**CostPrice** | Pointer to **float32** |  | [optional] 
**ListPrice** | Pointer to **float32** |  | [optional] 
**WholesalePrice** | Pointer to **float32** |  | [optional] 
**AdvancedPrice** | Pointer to [**[]ProductAdvancedPrice**](ProductAdvancedPrice.md) |  | [optional] 
**TaxClassId** | Pointer to **string** |  | [optional] 
**AvailForSale** | Pointer to **bool** |  | [optional] 
**AllowBackorders** | Pointer to **bool** |  | [optional] 
**InStock** | Pointer to **bool** |  | [optional] 
**ManageStock** | Pointer to **bool** |  | [optional] 
**InventoryLevel** | Pointer to **float32** |  | [optional] 
**Inventory** | Pointer to [**[]ProductInventory**](ProductInventory.md) |  | [optional] 
**MinQuantity** | Pointer to **float32** |  | [optional] 
**DefaultQtyInPack** | Pointer to **float32** |  | [optional] 
**IsQtyInPackFixed** | Pointer to **bool** |  | [optional] 
**WeightUnit** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] 
**DimensionsUnit** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **float32** |  | [optional] 
**Height** | Pointer to **float32** |  | [optional] 
**Length** | Pointer to **float32** |  | [optional] 
**MetaTitle** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**MetaKeywords** | Pointer to **string** |  | [optional] 
**Discounts** | Pointer to [**[]Discount**](Discount.md) |  | [optional] 
**IsVirtual** | Pointer to **bool** |  | [optional] 
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


