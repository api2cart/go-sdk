# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**UModel** | Pointer to **NullableString** |  | [optional] 
**USku** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ShortDescription** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**AdvancedPrice** | Pointer to [**[]ProductAdvancedPrice**](ProductAdvancedPrice.md) |  | [optional] 
**CostPrice** | Pointer to **NullableFloat32** |  | [optional] 
**Quantity** | Pointer to **NullableFloat32** |  | [optional] 
**Inventory** | Pointer to [**[]ProductInventory**](ProductInventory.md) |  | [optional] 
**GroupItems** | Pointer to [**[]ProductGroupItem**](ProductGroupItem.md) |  | [optional] 
**UBrandId** | Pointer to **NullableString** |  | [optional] 
**UBrand** | Pointer to **NullableString** |  | [optional] 
**CategoriesIds** | Pointer to **[]string** |  | [optional] 
**StoresIds** | Pointer to **[]string** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**SeoUrl** | Pointer to **NullableString** |  | [optional] 
**MetaTitle** | Pointer to **NullableString** |  | [optional] 
**MetaKeywords** | Pointer to **NullableString** |  | [optional] 
**MetaDescription** | Pointer to **NullableString** |  | [optional] 
**AvailSale** | Pointer to **NullableBool** |  | [optional] 
**AvailView** | Pointer to **NullableBool** |  | [optional] 
**IsVirtual** | Pointer to **NullableBool** |  | [optional] 
**IsDownloadable** | Pointer to **NullableBool** |  | [optional] 
**Weight** | Pointer to **NullableFloat32** |  | [optional] 
**WeightUnit** | Pointer to **NullableString** |  | [optional] 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**InStock** | Pointer to **NullableBool** |  | [optional] 
**Backorders** | Pointer to **NullableString** |  | [optional] 
**ManageStock** | Pointer to **NullableString** |  | [optional] 
**IsStockManaged** | Pointer to **NullableBool** |  | [optional] 
**OnSale** | Pointer to **NullableBool** |  | [optional] 
**CreateAt** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedAt** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**TaxClassId** | Pointer to **NullableString** |  | [optional] 
**SpecialPrice** | Pointer to [**NullableSpecialPrice**](SpecialPrice.md) |  | [optional] 
**TierPrice** | Pointer to [**[]ProductTierPrice**](ProductTierPrice.md) |  | [optional] 
**GroupPrice** | Pointer to [**[]ProductGroupPrice**](ProductGroupPrice.md) |  | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) |  | [optional] 
**ProductOptions** | Pointer to [**[]ProductOption**](ProductOption.md) |  | [optional] 
**UUpc** | Pointer to **NullableString** |  | [optional] 
**UMpn** | Pointer to **NullableString** |  | [optional] 
**UGtin** | Pointer to **NullableString** |  | [optional] 
**UIsbn** | Pointer to **NullableString** |  | [optional] 
**UEan** | Pointer to **NullableString** |  | [optional] 
**RelatedProductsIds** | Pointer to **[]string** |  | [optional] 
**UpSellProductsIds** | Pointer to **[]string** |  | [optional] 
**CrossSellProductsIds** | Pointer to **[]string** |  | [optional] 
**DimensionsUnit** | Pointer to **NullableString** |  | [optional] 
**Width** | Pointer to **NullableFloat32** |  | [optional] 
**Height** | Pointer to **NullableFloat32** |  | [optional] 
**Length** | Pointer to **NullableFloat32** |  | [optional] 
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

### SetTypeNil

`func (o *Product) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Product) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
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

### SetUModelNil

`func (o *Product) SetUModelNil(b bool)`

 SetUModelNil sets the value for UModel to be an explicit nil

### UnsetUModel
`func (o *Product) UnsetUModel()`

UnsetUModel ensures that no value is present for UModel, not even an explicit nil
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

### SetUSkuNil

`func (o *Product) SetUSkuNil(b bool)`

 SetUSkuNil sets the value for USku to be an explicit nil

### UnsetUSku
`func (o *Product) UnsetUSku()`

UnsetUSku ensures that no value is present for USku, not even an explicit nil
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

### SetNameNil

`func (o *Product) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Product) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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

### SetDescriptionNil

`func (o *Product) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Product) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
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

### SetShortDescriptionNil

`func (o *Product) SetShortDescriptionNil(b bool)`

 SetShortDescriptionNil sets the value for ShortDescription to be an explicit nil

### UnsetShortDescription
`func (o *Product) UnsetShortDescription()`

UnsetShortDescription ensures that no value is present for ShortDescription, not even an explicit nil
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

### SetPriceNil

`func (o *Product) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *Product) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
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

### SetCostPriceNil

`func (o *Product) SetCostPriceNil(b bool)`

 SetCostPriceNil sets the value for CostPrice to be an explicit nil

### UnsetCostPrice
`func (o *Product) UnsetCostPrice()`

UnsetCostPrice ensures that no value is present for CostPrice, not even an explicit nil
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

### SetQuantityNil

`func (o *Product) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *Product) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
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

### SetUBrandIdNil

`func (o *Product) SetUBrandIdNil(b bool)`

 SetUBrandIdNil sets the value for UBrandId to be an explicit nil

### UnsetUBrandId
`func (o *Product) UnsetUBrandId()`

UnsetUBrandId ensures that no value is present for UBrandId, not even an explicit nil
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

### SetUBrandNil

`func (o *Product) SetUBrandNil(b bool)`

 SetUBrandNil sets the value for UBrand to be an explicit nil

### UnsetUBrand
`func (o *Product) UnsetUBrand()`

UnsetUBrand ensures that no value is present for UBrand, not even an explicit nil
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

### SetUrlNil

`func (o *Product) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Product) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
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

### SetSeoUrlNil

`func (o *Product) SetSeoUrlNil(b bool)`

 SetSeoUrlNil sets the value for SeoUrl to be an explicit nil

### UnsetSeoUrl
`func (o *Product) UnsetSeoUrl()`

UnsetSeoUrl ensures that no value is present for SeoUrl, not even an explicit nil
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

### SetMetaTitleNil

`func (o *Product) SetMetaTitleNil(b bool)`

 SetMetaTitleNil sets the value for MetaTitle to be an explicit nil

### UnsetMetaTitle
`func (o *Product) UnsetMetaTitle()`

UnsetMetaTitle ensures that no value is present for MetaTitle, not even an explicit nil
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

### SetMetaKeywordsNil

`func (o *Product) SetMetaKeywordsNil(b bool)`

 SetMetaKeywordsNil sets the value for MetaKeywords to be an explicit nil

### UnsetMetaKeywords
`func (o *Product) UnsetMetaKeywords()`

UnsetMetaKeywords ensures that no value is present for MetaKeywords, not even an explicit nil
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

### SetMetaDescriptionNil

`func (o *Product) SetMetaDescriptionNil(b bool)`

 SetMetaDescriptionNil sets the value for MetaDescription to be an explicit nil

### UnsetMetaDescription
`func (o *Product) UnsetMetaDescription()`

UnsetMetaDescription ensures that no value is present for MetaDescription, not even an explicit nil
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

### SetAvailSaleNil

`func (o *Product) SetAvailSaleNil(b bool)`

 SetAvailSaleNil sets the value for AvailSale to be an explicit nil

### UnsetAvailSale
`func (o *Product) UnsetAvailSale()`

UnsetAvailSale ensures that no value is present for AvailSale, not even an explicit nil
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

### SetAvailViewNil

`func (o *Product) SetAvailViewNil(b bool)`

 SetAvailViewNil sets the value for AvailView to be an explicit nil

### UnsetAvailView
`func (o *Product) UnsetAvailView()`

UnsetAvailView ensures that no value is present for AvailView, not even an explicit nil
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

### SetIsVirtualNil

`func (o *Product) SetIsVirtualNil(b bool)`

 SetIsVirtualNil sets the value for IsVirtual to be an explicit nil

### UnsetIsVirtual
`func (o *Product) UnsetIsVirtual()`

UnsetIsVirtual ensures that no value is present for IsVirtual, not even an explicit nil
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

### SetIsDownloadableNil

`func (o *Product) SetIsDownloadableNil(b bool)`

 SetIsDownloadableNil sets the value for IsDownloadable to be an explicit nil

### UnsetIsDownloadable
`func (o *Product) UnsetIsDownloadable()`

UnsetIsDownloadable ensures that no value is present for IsDownloadable, not even an explicit nil
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

### SetWeightNil

`func (o *Product) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *Product) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
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

### SetWeightUnitNil

`func (o *Product) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *Product) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
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

### SetSortOrderNil

`func (o *Product) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *Product) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
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

### SetInStockNil

`func (o *Product) SetInStockNil(b bool)`

 SetInStockNil sets the value for InStock to be an explicit nil

### UnsetInStock
`func (o *Product) UnsetInStock()`

UnsetInStock ensures that no value is present for InStock, not even an explicit nil
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

### SetBackordersNil

`func (o *Product) SetBackordersNil(b bool)`

 SetBackordersNil sets the value for Backorders to be an explicit nil

### UnsetBackorders
`func (o *Product) UnsetBackorders()`

UnsetBackorders ensures that no value is present for Backorders, not even an explicit nil
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

### SetManageStockNil

`func (o *Product) SetManageStockNil(b bool)`

 SetManageStockNil sets the value for ManageStock to be an explicit nil

### UnsetManageStock
`func (o *Product) UnsetManageStock()`

UnsetManageStock ensures that no value is present for ManageStock, not even an explicit nil
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

### SetIsStockManagedNil

`func (o *Product) SetIsStockManagedNil(b bool)`

 SetIsStockManagedNil sets the value for IsStockManaged to be an explicit nil

### UnsetIsStockManaged
`func (o *Product) UnsetIsStockManaged()`

UnsetIsStockManaged ensures that no value is present for IsStockManaged, not even an explicit nil
### GetOnSale

`func (o *Product) GetOnSale() bool`

GetOnSale returns the OnSale field if non-nil, zero value otherwise.

### GetOnSaleOk

`func (o *Product) GetOnSaleOk() (*bool, bool)`

GetOnSaleOk returns a tuple with the OnSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSale

`func (o *Product) SetOnSale(v bool)`

SetOnSale sets OnSale field to given value.

### HasOnSale

`func (o *Product) HasOnSale() bool`

HasOnSale returns a boolean if a field has been set.

### SetOnSaleNil

`func (o *Product) SetOnSaleNil(b bool)`

 SetOnSaleNil sets the value for OnSale to be an explicit nil

### UnsetOnSale
`func (o *Product) UnsetOnSale()`

UnsetOnSale ensures that no value is present for OnSale, not even an explicit nil
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

### SetCreateAtNil

`func (o *Product) SetCreateAtNil(b bool)`

 SetCreateAtNil sets the value for CreateAt to be an explicit nil

### UnsetCreateAt
`func (o *Product) UnsetCreateAt()`

UnsetCreateAt ensures that no value is present for CreateAt, not even an explicit nil
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

### SetModifiedAtNil

`func (o *Product) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *Product) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
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

### SetTaxClassIdNil

`func (o *Product) SetTaxClassIdNil(b bool)`

 SetTaxClassIdNil sets the value for TaxClassId to be an explicit nil

### UnsetTaxClassId
`func (o *Product) UnsetTaxClassId()`

UnsetTaxClassId ensures that no value is present for TaxClassId, not even an explicit nil
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

### SetSpecialPriceNil

`func (o *Product) SetSpecialPriceNil(b bool)`

 SetSpecialPriceNil sets the value for SpecialPrice to be an explicit nil

### UnsetSpecialPrice
`func (o *Product) UnsetSpecialPrice()`

UnsetSpecialPrice ensures that no value is present for SpecialPrice, not even an explicit nil
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

### SetUUpcNil

`func (o *Product) SetUUpcNil(b bool)`

 SetUUpcNil sets the value for UUpc to be an explicit nil

### UnsetUUpc
`func (o *Product) UnsetUUpc()`

UnsetUUpc ensures that no value is present for UUpc, not even an explicit nil
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

### SetUMpnNil

`func (o *Product) SetUMpnNil(b bool)`

 SetUMpnNil sets the value for UMpn to be an explicit nil

### UnsetUMpn
`func (o *Product) UnsetUMpn()`

UnsetUMpn ensures that no value is present for UMpn, not even an explicit nil
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

### SetUGtinNil

`func (o *Product) SetUGtinNil(b bool)`

 SetUGtinNil sets the value for UGtin to be an explicit nil

### UnsetUGtin
`func (o *Product) UnsetUGtin()`

UnsetUGtin ensures that no value is present for UGtin, not even an explicit nil
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

### SetUIsbnNil

`func (o *Product) SetUIsbnNil(b bool)`

 SetUIsbnNil sets the value for UIsbn to be an explicit nil

### UnsetUIsbn
`func (o *Product) UnsetUIsbn()`

UnsetUIsbn ensures that no value is present for UIsbn, not even an explicit nil
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

### SetUEanNil

`func (o *Product) SetUEanNil(b bool)`

 SetUEanNil sets the value for UEan to be an explicit nil

### UnsetUEan
`func (o *Product) UnsetUEan()`

UnsetUEan ensures that no value is present for UEan, not even an explicit nil
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

### SetDimensionsUnitNil

`func (o *Product) SetDimensionsUnitNil(b bool)`

 SetDimensionsUnitNil sets the value for DimensionsUnit to be an explicit nil

### UnsetDimensionsUnit
`func (o *Product) UnsetDimensionsUnit()`

UnsetDimensionsUnit ensures that no value is present for DimensionsUnit, not even an explicit nil
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

### SetWidthNil

`func (o *Product) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *Product) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
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

### SetHeightNil

`func (o *Product) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *Product) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
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

### SetLengthNil

`func (o *Product) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *Product) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
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

### SetAdditionalFieldsNil

`func (o *Product) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *Product) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
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

### SetCustomFieldsNil

`func (o *Product) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *Product) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


