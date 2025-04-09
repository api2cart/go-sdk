# ProductVariantAddBatchPayloadInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** |  | 
**Combination** | [**[]ProductVariantAddBatchPayloadInnerCombinationInner**](ProductVariantAddBatchPayloadInnerCombinationInner.md) | A unique combination that contains an array of options and their values, which form a variation. | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Sku** | **string** |  | 
**Model** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**OldPrice** | Pointer to **float32** |  | [optional] 
**CostPrice** | Pointer to **float32** |  | [optional] 
**SpecialPrice** | Pointer to **float32** |  | [optional] 
**SpriceCreate** | Pointer to **string** |  | [optional] 
**SpriceExpire** | Pointer to **string** |  | [optional] 
**AdvancedPrices** | Pointer to [**[]ProductUpdateBatchPayloadInnerAdvancedPricesInner**](ProductUpdateBatchPayloadInnerAdvancedPricesInner.md) |  | [optional] 
**MetaTitle** | Pointer to **float32** |  | [optional] 
**MetaDescription** | Pointer to **float32** |  | [optional] 
**MetaKeywords** | Pointer to **[]string** |  | [optional] 
**CategoriesIds** | Pointer to **[]string** |  | [optional] 
**StoresIds** | Pointer to **[]string** |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] 
**Width** | Pointer to **float32** |  | [optional] 
**Height** | Pointer to **float32** |  | [optional] 
**Length** | Pointer to **float32** |  | [optional] 
**WeightUnit** | Pointer to **string** |  | [optional] 
**WarehouseId** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**ManageStock** | Pointer to **bool** |  | [optional] 
**InStock** | Pointer to **bool** |  | [optional] 
**StoreId** | Pointer to **string** |  | [optional] 
**LangId** | Pointer to **string** |  | [optional] 
**TaxClassId** | Pointer to **string** |  | [optional] 
**BackorderStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Visible** | Pointer to **string** |  | [optional] 
**IsVirtual** | Pointer to **bool** |  | [optional] 
**Downloadable** | Pointer to **bool** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Upc** | Pointer to **string** |  | [optional] 
**Isbn** | Pointer to **string** |  | [optional] 
**Mpn** | Pointer to **string** |  | [optional] 
**Ean** | Pointer to **string** |  | [optional] 
**Barcode** | Pointer to **string** |  | [optional] 
**AvailableForSale** | Pointer to **bool** |  | [optional] 
**IsFreeShipping** | Pointer to **bool** |  | [optional] 
**Taxable** | Pointer to **bool** |  | [optional] 
**SeoUrl** | Pointer to **string** |  | [optional] 
**ManufacturerId** | Pointer to **string** |  | [optional] 
**HarmonizedSystemCode** | Pointer to **string** |  | [optional] 
**MarketplaceItemProperties** | Pointer to **map[string]interface{}** |  | [optional] 
**Images** | Pointer to [**[]ProductAddBatchPayloadInnerImagesInner**](ProductAddBatchPayloadInnerImagesInner.md) |  | [optional] 
**ProductImagesIds** | Pointer to **[]string** |  | [optional] 
**RelatedProductsIds** | Pointer to **[]string** |  | [optional] 
**UpSellProductsIds** | Pointer to **[]string** |  | [optional] 
**CrossSellProductsIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProductVariantAddBatchPayloadInner

`func NewProductVariantAddBatchPayloadInner(productId string, combination []ProductVariantAddBatchPayloadInnerCombinationInner, sku string, ) *ProductVariantAddBatchPayloadInner`

NewProductVariantAddBatchPayloadInner instantiates a new ProductVariantAddBatchPayloadInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariantAddBatchPayloadInnerWithDefaults

`func NewProductVariantAddBatchPayloadInnerWithDefaults() *ProductVariantAddBatchPayloadInner`

NewProductVariantAddBatchPayloadInnerWithDefaults instantiates a new ProductVariantAddBatchPayloadInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProductVariantAddBatchPayloadInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductVariantAddBatchPayloadInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductVariantAddBatchPayloadInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetCombination

`func (o *ProductVariantAddBatchPayloadInner) GetCombination() []ProductVariantAddBatchPayloadInnerCombinationInner`

GetCombination returns the Combination field if non-nil, zero value otherwise.

### GetCombinationOk

`func (o *ProductVariantAddBatchPayloadInner) GetCombinationOk() (*[]ProductVariantAddBatchPayloadInnerCombinationInner, bool)`

GetCombinationOk returns a tuple with the Combination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombination

`func (o *ProductVariantAddBatchPayloadInner) SetCombination(v []ProductVariantAddBatchPayloadInnerCombinationInner)`

SetCombination sets Combination field to given value.


### GetName

`func (o *ProductVariantAddBatchPayloadInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductVariantAddBatchPayloadInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductVariantAddBatchPayloadInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductVariantAddBatchPayloadInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProductVariantAddBatchPayloadInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductVariantAddBatchPayloadInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductVariantAddBatchPayloadInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductVariantAddBatchPayloadInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortDescription

`func (o *ProductVariantAddBatchPayloadInner) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ProductVariantAddBatchPayloadInner) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ProductVariantAddBatchPayloadInner) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ProductVariantAddBatchPayloadInner) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetSku

`func (o *ProductVariantAddBatchPayloadInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductVariantAddBatchPayloadInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductVariantAddBatchPayloadInner) SetSku(v string)`

SetSku sets Sku field to given value.


### GetModel

`func (o *ProductVariantAddBatchPayloadInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ProductVariantAddBatchPayloadInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ProductVariantAddBatchPayloadInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ProductVariantAddBatchPayloadInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPrice

`func (o *ProductVariantAddBatchPayloadInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductVariantAddBatchPayloadInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductVariantAddBatchPayloadInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductVariantAddBatchPayloadInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOldPrice

`func (o *ProductVariantAddBatchPayloadInner) GetOldPrice() float32`

GetOldPrice returns the OldPrice field if non-nil, zero value otherwise.

### GetOldPriceOk

`func (o *ProductVariantAddBatchPayloadInner) GetOldPriceOk() (*float32, bool)`

GetOldPriceOk returns a tuple with the OldPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPrice

`func (o *ProductVariantAddBatchPayloadInner) SetOldPrice(v float32)`

SetOldPrice sets OldPrice field to given value.

### HasOldPrice

`func (o *ProductVariantAddBatchPayloadInner) HasOldPrice() bool`

HasOldPrice returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductVariantAddBatchPayloadInner) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductVariantAddBatchPayloadInner) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductVariantAddBatchPayloadInner) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductVariantAddBatchPayloadInner) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetSpecialPrice

`func (o *ProductVariantAddBatchPayloadInner) GetSpecialPrice() float32`

GetSpecialPrice returns the SpecialPrice field if non-nil, zero value otherwise.

### GetSpecialPriceOk

`func (o *ProductVariantAddBatchPayloadInner) GetSpecialPriceOk() (*float32, bool)`

GetSpecialPriceOk returns a tuple with the SpecialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPrice

`func (o *ProductVariantAddBatchPayloadInner) SetSpecialPrice(v float32)`

SetSpecialPrice sets SpecialPrice field to given value.

### HasSpecialPrice

`func (o *ProductVariantAddBatchPayloadInner) HasSpecialPrice() bool`

HasSpecialPrice returns a boolean if a field has been set.

### GetSpriceCreate

`func (o *ProductVariantAddBatchPayloadInner) GetSpriceCreate() string`

GetSpriceCreate returns the SpriceCreate field if non-nil, zero value otherwise.

### GetSpriceCreateOk

`func (o *ProductVariantAddBatchPayloadInner) GetSpriceCreateOk() (*string, bool)`

GetSpriceCreateOk returns a tuple with the SpriceCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceCreate

`func (o *ProductVariantAddBatchPayloadInner) SetSpriceCreate(v string)`

SetSpriceCreate sets SpriceCreate field to given value.

### HasSpriceCreate

`func (o *ProductVariantAddBatchPayloadInner) HasSpriceCreate() bool`

HasSpriceCreate returns a boolean if a field has been set.

### GetSpriceExpire

`func (o *ProductVariantAddBatchPayloadInner) GetSpriceExpire() string`

GetSpriceExpire returns the SpriceExpire field if non-nil, zero value otherwise.

### GetSpriceExpireOk

`func (o *ProductVariantAddBatchPayloadInner) GetSpriceExpireOk() (*string, bool)`

GetSpriceExpireOk returns a tuple with the SpriceExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceExpire

`func (o *ProductVariantAddBatchPayloadInner) SetSpriceExpire(v string)`

SetSpriceExpire sets SpriceExpire field to given value.

### HasSpriceExpire

`func (o *ProductVariantAddBatchPayloadInner) HasSpriceExpire() bool`

HasSpriceExpire returns a boolean if a field has been set.

### GetAdvancedPrices

`func (o *ProductVariantAddBatchPayloadInner) GetAdvancedPrices() []ProductUpdateBatchPayloadInnerAdvancedPricesInner`

GetAdvancedPrices returns the AdvancedPrices field if non-nil, zero value otherwise.

### GetAdvancedPricesOk

`func (o *ProductVariantAddBatchPayloadInner) GetAdvancedPricesOk() (*[]ProductUpdateBatchPayloadInnerAdvancedPricesInner, bool)`

GetAdvancedPricesOk returns a tuple with the AdvancedPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPrices

`func (o *ProductVariantAddBatchPayloadInner) SetAdvancedPrices(v []ProductUpdateBatchPayloadInnerAdvancedPricesInner)`

SetAdvancedPrices sets AdvancedPrices field to given value.

### HasAdvancedPrices

`func (o *ProductVariantAddBatchPayloadInner) HasAdvancedPrices() bool`

HasAdvancedPrices returns a boolean if a field has been set.

### GetMetaTitle

`func (o *ProductVariantAddBatchPayloadInner) GetMetaTitle() float32`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *ProductVariantAddBatchPayloadInner) GetMetaTitleOk() (*float32, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *ProductVariantAddBatchPayloadInner) SetMetaTitle(v float32)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *ProductVariantAddBatchPayloadInner) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetMetaDescription

`func (o *ProductVariantAddBatchPayloadInner) GetMetaDescription() float32`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *ProductVariantAddBatchPayloadInner) GetMetaDescriptionOk() (*float32, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *ProductVariantAddBatchPayloadInner) SetMetaDescription(v float32)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *ProductVariantAddBatchPayloadInner) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *ProductVariantAddBatchPayloadInner) GetMetaKeywords() []string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *ProductVariantAddBatchPayloadInner) GetMetaKeywordsOk() (*[]string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *ProductVariantAddBatchPayloadInner) SetMetaKeywords(v []string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *ProductVariantAddBatchPayloadInner) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetCategoriesIds

`func (o *ProductVariantAddBatchPayloadInner) GetCategoriesIds() []string`

GetCategoriesIds returns the CategoriesIds field if non-nil, zero value otherwise.

### GetCategoriesIdsOk

`func (o *ProductVariantAddBatchPayloadInner) GetCategoriesIdsOk() (*[]string, bool)`

GetCategoriesIdsOk returns a tuple with the CategoriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesIds

`func (o *ProductVariantAddBatchPayloadInner) SetCategoriesIds(v []string)`

SetCategoriesIds sets CategoriesIds field to given value.

### HasCategoriesIds

`func (o *ProductVariantAddBatchPayloadInner) HasCategoriesIds() bool`

HasCategoriesIds returns a boolean if a field has been set.

### GetStoresIds

`func (o *ProductVariantAddBatchPayloadInner) GetStoresIds() []string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *ProductVariantAddBatchPayloadInner) GetStoresIdsOk() (*[]string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *ProductVariantAddBatchPayloadInner) SetStoresIds(v []string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *ProductVariantAddBatchPayloadInner) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetWeight

`func (o *ProductVariantAddBatchPayloadInner) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductVariantAddBatchPayloadInner) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductVariantAddBatchPayloadInner) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductVariantAddBatchPayloadInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWidth

`func (o *ProductVariantAddBatchPayloadInner) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ProductVariantAddBatchPayloadInner) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ProductVariantAddBatchPayloadInner) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ProductVariantAddBatchPayloadInner) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ProductVariantAddBatchPayloadInner) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ProductVariantAddBatchPayloadInner) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ProductVariantAddBatchPayloadInner) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ProductVariantAddBatchPayloadInner) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetLength

`func (o *ProductVariantAddBatchPayloadInner) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ProductVariantAddBatchPayloadInner) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ProductVariantAddBatchPayloadInner) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ProductVariantAddBatchPayloadInner) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWeightUnit

`func (o *ProductVariantAddBatchPayloadInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ProductVariantAddBatchPayloadInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ProductVariantAddBatchPayloadInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ProductVariantAddBatchPayloadInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetWarehouseId

`func (o *ProductVariantAddBatchPayloadInner) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ProductVariantAddBatchPayloadInner) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ProductVariantAddBatchPayloadInner) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ProductVariantAddBatchPayloadInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductVariantAddBatchPayloadInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductVariantAddBatchPayloadInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductVariantAddBatchPayloadInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductVariantAddBatchPayloadInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetManageStock

`func (o *ProductVariantAddBatchPayloadInner) GetManageStock() bool`

GetManageStock returns the ManageStock field if non-nil, zero value otherwise.

### GetManageStockOk

`func (o *ProductVariantAddBatchPayloadInner) GetManageStockOk() (*bool, bool)`

GetManageStockOk returns a tuple with the ManageStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStock

`func (o *ProductVariantAddBatchPayloadInner) SetManageStock(v bool)`

SetManageStock sets ManageStock field to given value.

### HasManageStock

`func (o *ProductVariantAddBatchPayloadInner) HasManageStock() bool`

HasManageStock returns a boolean if a field has been set.

### GetInStock

`func (o *ProductVariantAddBatchPayloadInner) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *ProductVariantAddBatchPayloadInner) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *ProductVariantAddBatchPayloadInner) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *ProductVariantAddBatchPayloadInner) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### GetStoreId

`func (o *ProductVariantAddBatchPayloadInner) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductVariantAddBatchPayloadInner) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductVariantAddBatchPayloadInner) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductVariantAddBatchPayloadInner) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetLangId

`func (o *ProductVariantAddBatchPayloadInner) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *ProductVariantAddBatchPayloadInner) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *ProductVariantAddBatchPayloadInner) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *ProductVariantAddBatchPayloadInner) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetTaxClassId

`func (o *ProductVariantAddBatchPayloadInner) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *ProductVariantAddBatchPayloadInner) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *ProductVariantAddBatchPayloadInner) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *ProductVariantAddBatchPayloadInner) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### GetBackorderStatus

`func (o *ProductVariantAddBatchPayloadInner) GetBackorderStatus() string`

GetBackorderStatus returns the BackorderStatus field if non-nil, zero value otherwise.

### GetBackorderStatusOk

`func (o *ProductVariantAddBatchPayloadInner) GetBackorderStatusOk() (*string, bool)`

GetBackorderStatusOk returns a tuple with the BackorderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderStatus

`func (o *ProductVariantAddBatchPayloadInner) SetBackorderStatus(v string)`

SetBackorderStatus sets BackorderStatus field to given value.

### HasBackorderStatus

`func (o *ProductVariantAddBatchPayloadInner) HasBackorderStatus() bool`

HasBackorderStatus returns a boolean if a field has been set.

### GetStatus

`func (o *ProductVariantAddBatchPayloadInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductVariantAddBatchPayloadInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductVariantAddBatchPayloadInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductVariantAddBatchPayloadInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisible

`func (o *ProductVariantAddBatchPayloadInner) GetVisible() string`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ProductVariantAddBatchPayloadInner) GetVisibleOk() (*string, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ProductVariantAddBatchPayloadInner) SetVisible(v string)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ProductVariantAddBatchPayloadInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetIsVirtual

`func (o *ProductVariantAddBatchPayloadInner) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *ProductVariantAddBatchPayloadInner) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *ProductVariantAddBatchPayloadInner) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *ProductVariantAddBatchPayloadInner) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetDownloadable

`func (o *ProductVariantAddBatchPayloadInner) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *ProductVariantAddBatchPayloadInner) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *ProductVariantAddBatchPayloadInner) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *ProductVariantAddBatchPayloadInner) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetIsDefault

`func (o *ProductVariantAddBatchPayloadInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ProductVariantAddBatchPayloadInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ProductVariantAddBatchPayloadInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ProductVariantAddBatchPayloadInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetUpc

`func (o *ProductVariantAddBatchPayloadInner) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ProductVariantAddBatchPayloadInner) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ProductVariantAddBatchPayloadInner) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ProductVariantAddBatchPayloadInner) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetIsbn

`func (o *ProductVariantAddBatchPayloadInner) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *ProductVariantAddBatchPayloadInner) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *ProductVariantAddBatchPayloadInner) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *ProductVariantAddBatchPayloadInner) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### GetMpn

`func (o *ProductVariantAddBatchPayloadInner) GetMpn() string`

GetMpn returns the Mpn field if non-nil, zero value otherwise.

### GetMpnOk

`func (o *ProductVariantAddBatchPayloadInner) GetMpnOk() (*string, bool)`

GetMpnOk returns a tuple with the Mpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpn

`func (o *ProductVariantAddBatchPayloadInner) SetMpn(v string)`

SetMpn sets Mpn field to given value.

### HasMpn

`func (o *ProductVariantAddBatchPayloadInner) HasMpn() bool`

HasMpn returns a boolean if a field has been set.

### GetEan

`func (o *ProductVariantAddBatchPayloadInner) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *ProductVariantAddBatchPayloadInner) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *ProductVariantAddBatchPayloadInner) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *ProductVariantAddBatchPayloadInner) HasEan() bool`

HasEan returns a boolean if a field has been set.

### GetBarcode

`func (o *ProductVariantAddBatchPayloadInner) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductVariantAddBatchPayloadInner) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductVariantAddBatchPayloadInner) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ProductVariantAddBatchPayloadInner) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetAvailableForSale

`func (o *ProductVariantAddBatchPayloadInner) GetAvailableForSale() bool`

GetAvailableForSale returns the AvailableForSale field if non-nil, zero value otherwise.

### GetAvailableForSaleOk

`func (o *ProductVariantAddBatchPayloadInner) GetAvailableForSaleOk() (*bool, bool)`

GetAvailableForSaleOk returns a tuple with the AvailableForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForSale

`func (o *ProductVariantAddBatchPayloadInner) SetAvailableForSale(v bool)`

SetAvailableForSale sets AvailableForSale field to given value.

### HasAvailableForSale

`func (o *ProductVariantAddBatchPayloadInner) HasAvailableForSale() bool`

HasAvailableForSale returns a boolean if a field has been set.

### GetIsFreeShipping

`func (o *ProductVariantAddBatchPayloadInner) GetIsFreeShipping() bool`

GetIsFreeShipping returns the IsFreeShipping field if non-nil, zero value otherwise.

### GetIsFreeShippingOk

`func (o *ProductVariantAddBatchPayloadInner) GetIsFreeShippingOk() (*bool, bool)`

GetIsFreeShippingOk returns a tuple with the IsFreeShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeShipping

`func (o *ProductVariantAddBatchPayloadInner) SetIsFreeShipping(v bool)`

SetIsFreeShipping sets IsFreeShipping field to given value.

### HasIsFreeShipping

`func (o *ProductVariantAddBatchPayloadInner) HasIsFreeShipping() bool`

HasIsFreeShipping returns a boolean if a field has been set.

### GetTaxable

`func (o *ProductVariantAddBatchPayloadInner) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ProductVariantAddBatchPayloadInner) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ProductVariantAddBatchPayloadInner) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ProductVariantAddBatchPayloadInner) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetSeoUrl

`func (o *ProductVariantAddBatchPayloadInner) GetSeoUrl() string`

GetSeoUrl returns the SeoUrl field if non-nil, zero value otherwise.

### GetSeoUrlOk

`func (o *ProductVariantAddBatchPayloadInner) GetSeoUrlOk() (*string, bool)`

GetSeoUrlOk returns a tuple with the SeoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoUrl

`func (o *ProductVariantAddBatchPayloadInner) SetSeoUrl(v string)`

SetSeoUrl sets SeoUrl field to given value.

### HasSeoUrl

`func (o *ProductVariantAddBatchPayloadInner) HasSeoUrl() bool`

HasSeoUrl returns a boolean if a field has been set.

### GetManufacturerId

`func (o *ProductVariantAddBatchPayloadInner) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *ProductVariantAddBatchPayloadInner) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *ProductVariantAddBatchPayloadInner) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *ProductVariantAddBatchPayloadInner) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetHarmonizedSystemCode

`func (o *ProductVariantAddBatchPayloadInner) GetHarmonizedSystemCode() string`

GetHarmonizedSystemCode returns the HarmonizedSystemCode field if non-nil, zero value otherwise.

### GetHarmonizedSystemCodeOk

`func (o *ProductVariantAddBatchPayloadInner) GetHarmonizedSystemCodeOk() (*string, bool)`

GetHarmonizedSystemCodeOk returns a tuple with the HarmonizedSystemCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarmonizedSystemCode

`func (o *ProductVariantAddBatchPayloadInner) SetHarmonizedSystemCode(v string)`

SetHarmonizedSystemCode sets HarmonizedSystemCode field to given value.

### HasHarmonizedSystemCode

`func (o *ProductVariantAddBatchPayloadInner) HasHarmonizedSystemCode() bool`

HasHarmonizedSystemCode returns a boolean if a field has been set.

### GetMarketplaceItemProperties

`func (o *ProductVariantAddBatchPayloadInner) GetMarketplaceItemProperties() map[string]interface{}`

GetMarketplaceItemProperties returns the MarketplaceItemProperties field if non-nil, zero value otherwise.

### GetMarketplaceItemPropertiesOk

`func (o *ProductVariantAddBatchPayloadInner) GetMarketplaceItemPropertiesOk() (*map[string]interface{}, bool)`

GetMarketplaceItemPropertiesOk returns a tuple with the MarketplaceItemProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceItemProperties

`func (o *ProductVariantAddBatchPayloadInner) SetMarketplaceItemProperties(v map[string]interface{})`

SetMarketplaceItemProperties sets MarketplaceItemProperties field to given value.

### HasMarketplaceItemProperties

`func (o *ProductVariantAddBatchPayloadInner) HasMarketplaceItemProperties() bool`

HasMarketplaceItemProperties returns a boolean if a field has been set.

### GetImages

`func (o *ProductVariantAddBatchPayloadInner) GetImages() []ProductAddBatchPayloadInnerImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ProductVariantAddBatchPayloadInner) GetImagesOk() (*[]ProductAddBatchPayloadInnerImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ProductVariantAddBatchPayloadInner) SetImages(v []ProductAddBatchPayloadInnerImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *ProductVariantAddBatchPayloadInner) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetProductImagesIds

`func (o *ProductVariantAddBatchPayloadInner) GetProductImagesIds() []string`

GetProductImagesIds returns the ProductImagesIds field if non-nil, zero value otherwise.

### GetProductImagesIdsOk

`func (o *ProductVariantAddBatchPayloadInner) GetProductImagesIdsOk() (*[]string, bool)`

GetProductImagesIdsOk returns a tuple with the ProductImagesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImagesIds

`func (o *ProductVariantAddBatchPayloadInner) SetProductImagesIds(v []string)`

SetProductImagesIds sets ProductImagesIds field to given value.

### HasProductImagesIds

`func (o *ProductVariantAddBatchPayloadInner) HasProductImagesIds() bool`

HasProductImagesIds returns a boolean if a field has been set.

### GetRelatedProductsIds

`func (o *ProductVariantAddBatchPayloadInner) GetRelatedProductsIds() []string`

GetRelatedProductsIds returns the RelatedProductsIds field if non-nil, zero value otherwise.

### GetRelatedProductsIdsOk

`func (o *ProductVariantAddBatchPayloadInner) GetRelatedProductsIdsOk() (*[]string, bool)`

GetRelatedProductsIdsOk returns a tuple with the RelatedProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedProductsIds

`func (o *ProductVariantAddBatchPayloadInner) SetRelatedProductsIds(v []string)`

SetRelatedProductsIds sets RelatedProductsIds field to given value.

### HasRelatedProductsIds

`func (o *ProductVariantAddBatchPayloadInner) HasRelatedProductsIds() bool`

HasRelatedProductsIds returns a boolean if a field has been set.

### GetUpSellProductsIds

`func (o *ProductVariantAddBatchPayloadInner) GetUpSellProductsIds() []string`

GetUpSellProductsIds returns the UpSellProductsIds field if non-nil, zero value otherwise.

### GetUpSellProductsIdsOk

`func (o *ProductVariantAddBatchPayloadInner) GetUpSellProductsIdsOk() (*[]string, bool)`

GetUpSellProductsIdsOk returns a tuple with the UpSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSellProductsIds

`func (o *ProductVariantAddBatchPayloadInner) SetUpSellProductsIds(v []string)`

SetUpSellProductsIds sets UpSellProductsIds field to given value.

### HasUpSellProductsIds

`func (o *ProductVariantAddBatchPayloadInner) HasUpSellProductsIds() bool`

HasUpSellProductsIds returns a boolean if a field has been set.

### GetCrossSellProductsIds

`func (o *ProductVariantAddBatchPayloadInner) GetCrossSellProductsIds() []string`

GetCrossSellProductsIds returns the CrossSellProductsIds field if non-nil, zero value otherwise.

### GetCrossSellProductsIdsOk

`func (o *ProductVariantAddBatchPayloadInner) GetCrossSellProductsIdsOk() (*[]string, bool)`

GetCrossSellProductsIdsOk returns a tuple with the CrossSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSellProductsIds

`func (o *ProductVariantAddBatchPayloadInner) SetCrossSellProductsIds(v []string)`

SetCrossSellProductsIds sets CrossSellProductsIds field to given value.

### HasCrossSellProductsIds

`func (o *ProductVariantAddBatchPayloadInner) HasCrossSellProductsIds() bool`

HasCrossSellProductsIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


