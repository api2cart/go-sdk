# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UModel** | Pointer to **string** |  | [optional] 
**USku** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**AdvancedPrice** | Pointer to [**[]ProductAdvancedPrice**](ProductAdvancedPrice.md) |  | [optional] 
**CostPrice** | Pointer to **float32** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**Inventory** | Pointer to [**[]ProductInventory**](ProductInventory.md) |  | [optional] 
**GroupItems** | Pointer to [**[]ProductGroupItem**](ProductGroupItem.md) |  | [optional] 
**UBrandId** | Pointer to **string** |  | [optional] 
**UBrand** | Pointer to **string** |  | [optional] 
**CategoriesIds** | Pointer to **[]string** |  | [optional] 
**StoresIds** | Pointer to **[]string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**SeoUrl** | Pointer to **string** |  | [optional] 
**MetaTitle** | Pointer to **string** |  | [optional] 
**MetaKeywords** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**AvailSale** | Pointer to **bool** |  | [optional] 
**AvailView** | Pointer to **bool** |  | [optional] 
**IsVirtual** | Pointer to **bool** |  | [optional] 
**IsDownloadable** | Pointer to **bool** |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] 
**WeightUnit** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**InStock** | Pointer to **bool** |  | [optional] 
**Backorders** | Pointer to **string** |  | [optional] 
**ManageStock** | Pointer to **string** |  | [optional] 
**IsStockManaged** | Pointer to **bool** |  | [optional] 
**CreateAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**TaxClassId** | Pointer to **string** |  | [optional] 
**SpecialPrice** | Pointer to [**SpecialPrice**](SpecialPrice.md) |  | [optional] 
**TierPrice** | Pointer to [**[]ProductTierPrice**](ProductTierPrice.md) |  | [optional] 
**GroupPrice** | Pointer to [**[]ProductGroupPrice**](ProductGroupPrice.md) |  | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) |  | [optional] 
**ProductOptions** | Pointer to [**[]ProductOption**](ProductOption.md) |  | [optional] 
**UUpc** | Pointer to **string** |  | [optional] 
**UMpn** | Pointer to **string** |  | [optional] 
**UGtin** | Pointer to **string** |  | [optional] 
**UIsbn** | Pointer to **string** |  | [optional] 
**UEan** | Pointer to **string** |  | [optional] 
**RelatedProductsIds** | Pointer to **[]string** |  | [optional] 
**UpSellProductsIds** | Pointer to **[]string** |  | [optional] 
**CrossSellProductsIds** | Pointer to **[]string** |  | [optional] 
**DimensionsUnit** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **float32** |  | [optional] 
**Height** | Pointer to **float32** |  | [optional] 
**Length** | Pointer to **float32** |  | [optional] 
**Discounts** | Pointer to [**[]Discount**](Discount.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProduct

`func NewProduct() *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Product) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Product) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Product) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Product) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Product) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUModel

`func (o *Product) GetUModel() string`

GetUModel returns the UModel field if non-nil, zero value otherwise.

### GetUModelOk

`func (o *Product) GetUModelOk() (*string, bool)`

GetUModelOk returns a tuple with the UModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUModel

`func (o *Product) SetUModel(v string)`

SetUModel sets UModel field to given value.

### HasUModel

`func (o *Product) HasUModel() bool`

HasUModel returns a boolean if a field has been set.

### GetUSku

`func (o *Product) GetUSku() string`

GetUSku returns the USku field if non-nil, zero value otherwise.

### GetUSkuOk

`func (o *Product) GetUSkuOk() (*string, bool)`

GetUSkuOk returns a tuple with the USku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSku

`func (o *Product) SetUSku(v string)`

SetUSku sets USku field to given value.

### HasUSku

`func (o *Product) HasUSku() bool`

HasUSku returns a boolean if a field has been set.

### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Product) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Product) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortDescription

`func (o *Product) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *Product) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *Product) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *Product) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetPrice

`func (o *Product) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Product) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Product) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Product) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAdvancedPrice

`func (o *Product) GetAdvancedPrice() []ProductAdvancedPrice`

GetAdvancedPrice returns the AdvancedPrice field if non-nil, zero value otherwise.

### GetAdvancedPriceOk

`func (o *Product) GetAdvancedPriceOk() (*[]ProductAdvancedPrice, bool)`

GetAdvancedPriceOk returns a tuple with the AdvancedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPrice

`func (o *Product) SetAdvancedPrice(v []ProductAdvancedPrice)`

SetAdvancedPrice sets AdvancedPrice field to given value.

### HasAdvancedPrice

`func (o *Product) HasAdvancedPrice() bool`

HasAdvancedPrice returns a boolean if a field has been set.

### GetCostPrice

`func (o *Product) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *Product) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *Product) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *Product) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *Product) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Product) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Product) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Product) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetInventory

`func (o *Product) GetInventory() []ProductInventory`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *Product) GetInventoryOk() (*[]ProductInventory, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *Product) SetInventory(v []ProductInventory)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *Product) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetGroupItems

`func (o *Product) GetGroupItems() []ProductGroupItem`

GetGroupItems returns the GroupItems field if non-nil, zero value otherwise.

### GetGroupItemsOk

`func (o *Product) GetGroupItemsOk() (*[]ProductGroupItem, bool)`

GetGroupItemsOk returns a tuple with the GroupItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupItems

`func (o *Product) SetGroupItems(v []ProductGroupItem)`

SetGroupItems sets GroupItems field to given value.

### HasGroupItems

`func (o *Product) HasGroupItems() bool`

HasGroupItems returns a boolean if a field has been set.

### GetUBrandId

`func (o *Product) GetUBrandId() string`

GetUBrandId returns the UBrandId field if non-nil, zero value otherwise.

### GetUBrandIdOk

`func (o *Product) GetUBrandIdOk() (*string, bool)`

GetUBrandIdOk returns a tuple with the UBrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUBrandId

`func (o *Product) SetUBrandId(v string)`

SetUBrandId sets UBrandId field to given value.

### HasUBrandId

`func (o *Product) HasUBrandId() bool`

HasUBrandId returns a boolean if a field has been set.

### GetUBrand

`func (o *Product) GetUBrand() string`

GetUBrand returns the UBrand field if non-nil, zero value otherwise.

### GetUBrandOk

`func (o *Product) GetUBrandOk() (*string, bool)`

GetUBrandOk returns a tuple with the UBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUBrand

`func (o *Product) SetUBrand(v string)`

SetUBrand sets UBrand field to given value.

### HasUBrand

`func (o *Product) HasUBrand() bool`

HasUBrand returns a boolean if a field has been set.

### GetCategoriesIds

`func (o *Product) GetCategoriesIds() []string`

GetCategoriesIds returns the CategoriesIds field if non-nil, zero value otherwise.

### GetCategoriesIdsOk

`func (o *Product) GetCategoriesIdsOk() (*[]string, bool)`

GetCategoriesIdsOk returns a tuple with the CategoriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesIds

`func (o *Product) SetCategoriesIds(v []string)`

SetCategoriesIds sets CategoriesIds field to given value.

### HasCategoriesIds

`func (o *Product) HasCategoriesIds() bool`

HasCategoriesIds returns a boolean if a field has been set.

### GetStoresIds

`func (o *Product) GetStoresIds() []string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *Product) GetStoresIdsOk() (*[]string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *Product) SetStoresIds(v []string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *Product) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetUrl

`func (o *Product) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Product) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Product) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Product) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSeoUrl

`func (o *Product) GetSeoUrl() string`

GetSeoUrl returns the SeoUrl field if non-nil, zero value otherwise.

### GetSeoUrlOk

`func (o *Product) GetSeoUrlOk() (*string, bool)`

GetSeoUrlOk returns a tuple with the SeoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoUrl

`func (o *Product) SetSeoUrl(v string)`

SetSeoUrl sets SeoUrl field to given value.

### HasSeoUrl

`func (o *Product) HasSeoUrl() bool`

HasSeoUrl returns a boolean if a field has been set.

### GetMetaTitle

`func (o *Product) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *Product) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *Product) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *Product) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *Product) GetMetaKeywords() string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *Product) GetMetaKeywordsOk() (*string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *Product) SetMetaKeywords(v string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *Product) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetMetaDescription

`func (o *Product) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *Product) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *Product) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *Product) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetAvailSale

`func (o *Product) GetAvailSale() bool`

GetAvailSale returns the AvailSale field if non-nil, zero value otherwise.

### GetAvailSaleOk

`func (o *Product) GetAvailSaleOk() (*bool, bool)`

GetAvailSaleOk returns a tuple with the AvailSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailSale

`func (o *Product) SetAvailSale(v bool)`

SetAvailSale sets AvailSale field to given value.

### HasAvailSale

`func (o *Product) HasAvailSale() bool`

HasAvailSale returns a boolean if a field has been set.

### GetAvailView

`func (o *Product) GetAvailView() bool`

GetAvailView returns the AvailView field if non-nil, zero value otherwise.

### GetAvailViewOk

`func (o *Product) GetAvailViewOk() (*bool, bool)`

GetAvailViewOk returns a tuple with the AvailView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailView

`func (o *Product) SetAvailView(v bool)`

SetAvailView sets AvailView field to given value.

### HasAvailView

`func (o *Product) HasAvailView() bool`

HasAvailView returns a boolean if a field has been set.

### GetIsVirtual

`func (o *Product) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *Product) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *Product) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *Product) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetIsDownloadable

`func (o *Product) GetIsDownloadable() bool`

GetIsDownloadable returns the IsDownloadable field if non-nil, zero value otherwise.

### GetIsDownloadableOk

`func (o *Product) GetIsDownloadableOk() (*bool, bool)`

GetIsDownloadableOk returns a tuple with the IsDownloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDownloadable

`func (o *Product) SetIsDownloadable(v bool)`

SetIsDownloadable sets IsDownloadable field to given value.

### HasIsDownloadable

`func (o *Product) HasIsDownloadable() bool`

HasIsDownloadable returns a boolean if a field has been set.

### GetWeight

`func (o *Product) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Product) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Product) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Product) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *Product) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *Product) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *Product) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *Product) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetSortOrder

`func (o *Product) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Product) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Product) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Product) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetInStock

`func (o *Product) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *Product) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *Product) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *Product) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### GetBackorders

`func (o *Product) GetBackorders() string`

GetBackorders returns the Backorders field if non-nil, zero value otherwise.

### GetBackordersOk

`func (o *Product) GetBackordersOk() (*string, bool)`

GetBackordersOk returns a tuple with the Backorders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorders

`func (o *Product) SetBackorders(v string)`

SetBackorders sets Backorders field to given value.

### HasBackorders

`func (o *Product) HasBackorders() bool`

HasBackorders returns a boolean if a field has been set.

### GetManageStock

`func (o *Product) GetManageStock() string`

GetManageStock returns the ManageStock field if non-nil, zero value otherwise.

### GetManageStockOk

`func (o *Product) GetManageStockOk() (*string, bool)`

GetManageStockOk returns a tuple with the ManageStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStock

`func (o *Product) SetManageStock(v string)`

SetManageStock sets ManageStock field to given value.

### HasManageStock

`func (o *Product) HasManageStock() bool`

HasManageStock returns a boolean if a field has been set.

### GetIsStockManaged

`func (o *Product) GetIsStockManaged() bool`

GetIsStockManaged returns the IsStockManaged field if non-nil, zero value otherwise.

### GetIsStockManagedOk

`func (o *Product) GetIsStockManagedOk() (*bool, bool)`

GetIsStockManagedOk returns a tuple with the IsStockManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStockManaged

`func (o *Product) SetIsStockManaged(v bool)`

SetIsStockManaged sets IsStockManaged field to given value.

### HasIsStockManaged

`func (o *Product) HasIsStockManaged() bool`

HasIsStockManaged returns a boolean if a field has been set.

### GetCreateAt

`func (o *Product) GetCreateAt() A2CDateTime`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *Product) GetCreateAtOk() (*A2CDateTime, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *Product) SetCreateAt(v A2CDateTime)`

SetCreateAt sets CreateAt field to given value.

### HasCreateAt

`func (o *Product) HasCreateAt() bool`

HasCreateAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Product) GetModifiedAt() A2CDateTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Product) GetModifiedAtOk() (*A2CDateTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Product) SetModifiedAt(v A2CDateTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Product) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetTaxClassId

`func (o *Product) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *Product) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *Product) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *Product) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### GetSpecialPrice

`func (o *Product) GetSpecialPrice() SpecialPrice`

GetSpecialPrice returns the SpecialPrice field if non-nil, zero value otherwise.

### GetSpecialPriceOk

`func (o *Product) GetSpecialPriceOk() (*SpecialPrice, bool)`

GetSpecialPriceOk returns a tuple with the SpecialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPrice

`func (o *Product) SetSpecialPrice(v SpecialPrice)`

SetSpecialPrice sets SpecialPrice field to given value.

### HasSpecialPrice

`func (o *Product) HasSpecialPrice() bool`

HasSpecialPrice returns a boolean if a field has been set.

### GetTierPrice

`func (o *Product) GetTierPrice() []ProductTierPrice`

GetTierPrice returns the TierPrice field if non-nil, zero value otherwise.

### GetTierPriceOk

`func (o *Product) GetTierPriceOk() (*[]ProductTierPrice, bool)`

GetTierPriceOk returns a tuple with the TierPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierPrice

`func (o *Product) SetTierPrice(v []ProductTierPrice)`

SetTierPrice sets TierPrice field to given value.

### HasTierPrice

`func (o *Product) HasTierPrice() bool`

HasTierPrice returns a boolean if a field has been set.

### GetGroupPrice

`func (o *Product) GetGroupPrice() []ProductGroupPrice`

GetGroupPrice returns the GroupPrice field if non-nil, zero value otherwise.

### GetGroupPriceOk

`func (o *Product) GetGroupPriceOk() (*[]ProductGroupPrice, bool)`

GetGroupPriceOk returns a tuple with the GroupPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPrice

`func (o *Product) SetGroupPrice(v []ProductGroupPrice)`

SetGroupPrice sets GroupPrice field to given value.

### HasGroupPrice

`func (o *Product) HasGroupPrice() bool`

HasGroupPrice returns a boolean if a field has been set.

### GetImages

`func (o *Product) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Product) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Product) SetImages(v []Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *Product) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetProductOptions

`func (o *Product) GetProductOptions() []ProductOption`

GetProductOptions returns the ProductOptions field if non-nil, zero value otherwise.

### GetProductOptionsOk

`func (o *Product) GetProductOptionsOk() (*[]ProductOption, bool)`

GetProductOptionsOk returns a tuple with the ProductOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOptions

`func (o *Product) SetProductOptions(v []ProductOption)`

SetProductOptions sets ProductOptions field to given value.

### HasProductOptions

`func (o *Product) HasProductOptions() bool`

HasProductOptions returns a boolean if a field has been set.

### GetUUpc

`func (o *Product) GetUUpc() string`

GetUUpc returns the UUpc field if non-nil, zero value otherwise.

### GetUUpcOk

`func (o *Product) GetUUpcOk() (*string, bool)`

GetUUpcOk returns a tuple with the UUpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUpc

`func (o *Product) SetUUpc(v string)`

SetUUpc sets UUpc field to given value.

### HasUUpc

`func (o *Product) HasUUpc() bool`

HasUUpc returns a boolean if a field has been set.

### GetUMpn

`func (o *Product) GetUMpn() string`

GetUMpn returns the UMpn field if non-nil, zero value otherwise.

### GetUMpnOk

`func (o *Product) GetUMpnOk() (*string, bool)`

GetUMpnOk returns a tuple with the UMpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUMpn

`func (o *Product) SetUMpn(v string)`

SetUMpn sets UMpn field to given value.

### HasUMpn

`func (o *Product) HasUMpn() bool`

HasUMpn returns a boolean if a field has been set.

### GetUGtin

`func (o *Product) GetUGtin() string`

GetUGtin returns the UGtin field if non-nil, zero value otherwise.

### GetUGtinOk

`func (o *Product) GetUGtinOk() (*string, bool)`

GetUGtinOk returns a tuple with the UGtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUGtin

`func (o *Product) SetUGtin(v string)`

SetUGtin sets UGtin field to given value.

### HasUGtin

`func (o *Product) HasUGtin() bool`

HasUGtin returns a boolean if a field has been set.

### GetUIsbn

`func (o *Product) GetUIsbn() string`

GetUIsbn returns the UIsbn field if non-nil, zero value otherwise.

### GetUIsbnOk

`func (o *Product) GetUIsbnOk() (*string, bool)`

GetUIsbnOk returns a tuple with the UIsbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUIsbn

`func (o *Product) SetUIsbn(v string)`

SetUIsbn sets UIsbn field to given value.

### HasUIsbn

`func (o *Product) HasUIsbn() bool`

HasUIsbn returns a boolean if a field has been set.

### GetUEan

`func (o *Product) GetUEan() string`

GetUEan returns the UEan field if non-nil, zero value otherwise.

### GetUEanOk

`func (o *Product) GetUEanOk() (*string, bool)`

GetUEanOk returns a tuple with the UEan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUEan

`func (o *Product) SetUEan(v string)`

SetUEan sets UEan field to given value.

### HasUEan

`func (o *Product) HasUEan() bool`

HasUEan returns a boolean if a field has been set.

### GetRelatedProductsIds

`func (o *Product) GetRelatedProductsIds() []string`

GetRelatedProductsIds returns the RelatedProductsIds field if non-nil, zero value otherwise.

### GetRelatedProductsIdsOk

`func (o *Product) GetRelatedProductsIdsOk() (*[]string, bool)`

GetRelatedProductsIdsOk returns a tuple with the RelatedProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedProductsIds

`func (o *Product) SetRelatedProductsIds(v []string)`

SetRelatedProductsIds sets RelatedProductsIds field to given value.

### HasRelatedProductsIds

`func (o *Product) HasRelatedProductsIds() bool`

HasRelatedProductsIds returns a boolean if a field has been set.

### GetUpSellProductsIds

`func (o *Product) GetUpSellProductsIds() []string`

GetUpSellProductsIds returns the UpSellProductsIds field if non-nil, zero value otherwise.

### GetUpSellProductsIdsOk

`func (o *Product) GetUpSellProductsIdsOk() (*[]string, bool)`

GetUpSellProductsIdsOk returns a tuple with the UpSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSellProductsIds

`func (o *Product) SetUpSellProductsIds(v []string)`

SetUpSellProductsIds sets UpSellProductsIds field to given value.

### HasUpSellProductsIds

`func (o *Product) HasUpSellProductsIds() bool`

HasUpSellProductsIds returns a boolean if a field has been set.

### GetCrossSellProductsIds

`func (o *Product) GetCrossSellProductsIds() []string`

GetCrossSellProductsIds returns the CrossSellProductsIds field if non-nil, zero value otherwise.

### GetCrossSellProductsIdsOk

`func (o *Product) GetCrossSellProductsIdsOk() (*[]string, bool)`

GetCrossSellProductsIdsOk returns a tuple with the CrossSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSellProductsIds

`func (o *Product) SetCrossSellProductsIds(v []string)`

SetCrossSellProductsIds sets CrossSellProductsIds field to given value.

### HasCrossSellProductsIds

`func (o *Product) HasCrossSellProductsIds() bool`

HasCrossSellProductsIds returns a boolean if a field has been set.

### GetDimensionsUnit

`func (o *Product) GetDimensionsUnit() string`

GetDimensionsUnit returns the DimensionsUnit field if non-nil, zero value otherwise.

### GetDimensionsUnitOk

`func (o *Product) GetDimensionsUnitOk() (*string, bool)`

GetDimensionsUnitOk returns a tuple with the DimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsUnit

`func (o *Product) SetDimensionsUnit(v string)`

SetDimensionsUnit sets DimensionsUnit field to given value.

### HasDimensionsUnit

`func (o *Product) HasDimensionsUnit() bool`

HasDimensionsUnit returns a boolean if a field has been set.

### GetWidth

`func (o *Product) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Product) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Product) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Product) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *Product) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Product) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Product) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Product) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetLength

`func (o *Product) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Product) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Product) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *Product) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetDiscounts

`func (o *Product) GetDiscounts() []Discount`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *Product) GetDiscountsOk() (*[]Discount, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *Product) SetDiscounts(v []Discount)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *Product) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Product) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Product) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Product) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Product) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *Product) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Product) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Product) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Product) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


