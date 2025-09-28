# ProductAddBatchPayloadInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Asin** | Pointer to **string** |  | [optional] 
**Upc** | Pointer to **string** |  | [optional] 
**Ean** | Pointer to **string** |  | [optional] 
**Gtin** | Pointer to **string** |  | [optional] 
**Mpn** | Pointer to **string** |  | [optional] 
**Isbn** | Pointer to **string** |  | [optional] 
**Barcode** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**OldPrice** | Pointer to **float32** |  | [optional] 
**CostPrice** | Pointer to **float32** |  | [optional] 
**SpecialPrice** | Pointer to **float32** |  | [optional] 
**SpriceCreate** | Pointer to **string** |  | [optional] 
**SpriceExpire** | Pointer to **string** |  | [optional] 
**AvailFrom** | Pointer to **string** |  | [optional] 
**AdvancedPrices** | Pointer to [**[]ProductAddBatchPayloadInnerAdvancedPricesInner**](ProductAddBatchPayloadInnerAdvancedPricesInner.md) |  | [optional] 
**FixedCostShippingPrice** | Pointer to **float32** |  | [optional] 
**BuyitnowPrice** | Pointer to **float32** |  | [optional] 
**ReservePrice** | Pointer to **float32** |  | [optional] 
**BestOffer** | Pointer to **float32** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**ManageStock** | Pointer to **bool** |  | [optional] 
**ProductType** | Pointer to **string** |  | [optional] 
**MarketplaceItemProperties** | Pointer to **map[string]interface{}** |  | [optional] 
**Specifics** | Pointer to **map[string]interface{}** |  | [optional] 
**IsFreeShipping** | Pointer to **bool** |  | [optional] 
**Taxable** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to **string** |  | [optional] 
**ConditionDescription** | Pointer to **string** |  | [optional] 
**Visible** | Pointer to **string** |  | [optional] 
**AvailableForView** | Pointer to **bool** |  | [optional] 
**AvailableForSale** | Pointer to **bool** |  | [optional] 
**IsVirtual** | Pointer to **bool** |  | [optional] 
**InStock** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ListingType** | Pointer to **string** |  | [optional] 
**ListingDuration** | Pointer to **string** |  | [optional] 
**Downloadable** | Pointer to **bool** |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] 
**Length** | Pointer to **float32** |  | [optional] 
**Width** | Pointer to **float32** |  | [optional] 
**Height** | Pointer to **float32** |  | [optional] 
**WeightUnit** | Pointer to **string** |  | [optional] 
**DimensionsUnit** | Pointer to **string** |  | [optional] 
**StoreId** | Pointer to **string** |  | [optional] 
**LangId** | Pointer to **string** |  | [optional] 
**CategoryId** | Pointer to **string** |  | [optional] 
**WarehouseId** | Pointer to **string** |  | [optional] 
**CategoriesIds** | Pointer to **[]string** |  | [optional] 
**RelatedProductsIds** | Pointer to **[]string** |  | [optional] 
**UpSellProductsIds** | Pointer to **[]string** |  | [optional] 
**CrossSellProductsIds** | Pointer to **[]string** |  | [optional] 
**StoresIds** | Pointer to **[]string** |  | [optional] 
**TaxClassId** | Pointer to **string** |  | [optional] 
**SalesTax** | Pointer to [**ProductAddBatchPayloadInnerSalesTax**](ProductAddBatchPayloadInnerSalesTax.md) |  | [optional] 
**MetaTitle** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**MetaKeywords** | Pointer to **[]string** |  | [optional] 
**SearchKeywords** | Pointer to **[]string** |  | [optional] 
**HarmonizedSystemCode** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**SeoUrl** | Pointer to **string** |  | [optional] 
**ExternalProductLink** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**ManufacturerId** | Pointer to **string** |  | [optional] 
**BackorderStatus** | Pointer to **string** |  | [optional] 
**Images** | Pointer to [**[]ProductAddBatchPayloadInnerImagesInner**](ProductAddBatchPayloadInnerImagesInner.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Files** | Pointer to [**[]ProductAddFilesInner**](ProductAddFilesInner.md) |  | [optional] 

## Methods

### NewProductAddBatchPayloadInner

`func NewProductAddBatchPayloadInner() *ProductAddBatchPayloadInner`

NewProductAddBatchPayloadInner instantiates a new ProductAddBatchPayloadInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAddBatchPayloadInnerWithDefaults

`func NewProductAddBatchPayloadInnerWithDefaults() *ProductAddBatchPayloadInner`

NewProductAddBatchPayloadInnerWithDefaults instantiates a new ProductAddBatchPayloadInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductAddBatchPayloadInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductAddBatchPayloadInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductAddBatchPayloadInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductAddBatchPayloadInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProductAddBatchPayloadInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductAddBatchPayloadInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductAddBatchPayloadInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductAddBatchPayloadInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortDescription

`func (o *ProductAddBatchPayloadInner) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ProductAddBatchPayloadInner) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ProductAddBatchPayloadInner) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ProductAddBatchPayloadInner) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetSku

`func (o *ProductAddBatchPayloadInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductAddBatchPayloadInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductAddBatchPayloadInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductAddBatchPayloadInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetModel

`func (o *ProductAddBatchPayloadInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ProductAddBatchPayloadInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ProductAddBatchPayloadInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ProductAddBatchPayloadInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetAsin

`func (o *ProductAddBatchPayloadInner) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *ProductAddBatchPayloadInner) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *ProductAddBatchPayloadInner) SetAsin(v string)`

SetAsin sets Asin field to given value.

### HasAsin

`func (o *ProductAddBatchPayloadInner) HasAsin() bool`

HasAsin returns a boolean if a field has been set.

### GetUpc

`func (o *ProductAddBatchPayloadInner) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ProductAddBatchPayloadInner) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ProductAddBatchPayloadInner) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ProductAddBatchPayloadInner) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetEan

`func (o *ProductAddBatchPayloadInner) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *ProductAddBatchPayloadInner) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *ProductAddBatchPayloadInner) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *ProductAddBatchPayloadInner) HasEan() bool`

HasEan returns a boolean if a field has been set.

### GetGtin

`func (o *ProductAddBatchPayloadInner) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *ProductAddBatchPayloadInner) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *ProductAddBatchPayloadInner) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *ProductAddBatchPayloadInner) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### GetMpn

`func (o *ProductAddBatchPayloadInner) GetMpn() string`

GetMpn returns the Mpn field if non-nil, zero value otherwise.

### GetMpnOk

`func (o *ProductAddBatchPayloadInner) GetMpnOk() (*string, bool)`

GetMpnOk returns a tuple with the Mpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpn

`func (o *ProductAddBatchPayloadInner) SetMpn(v string)`

SetMpn sets Mpn field to given value.

### HasMpn

`func (o *ProductAddBatchPayloadInner) HasMpn() bool`

HasMpn returns a boolean if a field has been set.

### GetIsbn

`func (o *ProductAddBatchPayloadInner) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *ProductAddBatchPayloadInner) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *ProductAddBatchPayloadInner) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *ProductAddBatchPayloadInner) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### GetBarcode

`func (o *ProductAddBatchPayloadInner) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductAddBatchPayloadInner) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductAddBatchPayloadInner) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ProductAddBatchPayloadInner) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetPrice

`func (o *ProductAddBatchPayloadInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductAddBatchPayloadInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductAddBatchPayloadInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductAddBatchPayloadInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOldPrice

`func (o *ProductAddBatchPayloadInner) GetOldPrice() float32`

GetOldPrice returns the OldPrice field if non-nil, zero value otherwise.

### GetOldPriceOk

`func (o *ProductAddBatchPayloadInner) GetOldPriceOk() (*float32, bool)`

GetOldPriceOk returns a tuple with the OldPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPrice

`func (o *ProductAddBatchPayloadInner) SetOldPrice(v float32)`

SetOldPrice sets OldPrice field to given value.

### HasOldPrice

`func (o *ProductAddBatchPayloadInner) HasOldPrice() bool`

HasOldPrice returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductAddBatchPayloadInner) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductAddBatchPayloadInner) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductAddBatchPayloadInner) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductAddBatchPayloadInner) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetSpecialPrice

`func (o *ProductAddBatchPayloadInner) GetSpecialPrice() float32`

GetSpecialPrice returns the SpecialPrice field if non-nil, zero value otherwise.

### GetSpecialPriceOk

`func (o *ProductAddBatchPayloadInner) GetSpecialPriceOk() (*float32, bool)`

GetSpecialPriceOk returns a tuple with the SpecialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPrice

`func (o *ProductAddBatchPayloadInner) SetSpecialPrice(v float32)`

SetSpecialPrice sets SpecialPrice field to given value.

### HasSpecialPrice

`func (o *ProductAddBatchPayloadInner) HasSpecialPrice() bool`

HasSpecialPrice returns a boolean if a field has been set.

### GetSpriceCreate

`func (o *ProductAddBatchPayloadInner) GetSpriceCreate() string`

GetSpriceCreate returns the SpriceCreate field if non-nil, zero value otherwise.

### GetSpriceCreateOk

`func (o *ProductAddBatchPayloadInner) GetSpriceCreateOk() (*string, bool)`

GetSpriceCreateOk returns a tuple with the SpriceCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceCreate

`func (o *ProductAddBatchPayloadInner) SetSpriceCreate(v string)`

SetSpriceCreate sets SpriceCreate field to given value.

### HasSpriceCreate

`func (o *ProductAddBatchPayloadInner) HasSpriceCreate() bool`

HasSpriceCreate returns a boolean if a field has been set.

### GetSpriceExpire

`func (o *ProductAddBatchPayloadInner) GetSpriceExpire() string`

GetSpriceExpire returns the SpriceExpire field if non-nil, zero value otherwise.

### GetSpriceExpireOk

`func (o *ProductAddBatchPayloadInner) GetSpriceExpireOk() (*string, bool)`

GetSpriceExpireOk returns a tuple with the SpriceExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceExpire

`func (o *ProductAddBatchPayloadInner) SetSpriceExpire(v string)`

SetSpriceExpire sets SpriceExpire field to given value.

### HasSpriceExpire

`func (o *ProductAddBatchPayloadInner) HasSpriceExpire() bool`

HasSpriceExpire returns a boolean if a field has been set.

### GetAvailFrom

`func (o *ProductAddBatchPayloadInner) GetAvailFrom() string`

GetAvailFrom returns the AvailFrom field if non-nil, zero value otherwise.

### GetAvailFromOk

`func (o *ProductAddBatchPayloadInner) GetAvailFromOk() (*string, bool)`

GetAvailFromOk returns a tuple with the AvailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailFrom

`func (o *ProductAddBatchPayloadInner) SetAvailFrom(v string)`

SetAvailFrom sets AvailFrom field to given value.

### HasAvailFrom

`func (o *ProductAddBatchPayloadInner) HasAvailFrom() bool`

HasAvailFrom returns a boolean if a field has been set.

### GetAdvancedPrices

`func (o *ProductAddBatchPayloadInner) GetAdvancedPrices() []ProductAddBatchPayloadInnerAdvancedPricesInner`

GetAdvancedPrices returns the AdvancedPrices field if non-nil, zero value otherwise.

### GetAdvancedPricesOk

`func (o *ProductAddBatchPayloadInner) GetAdvancedPricesOk() (*[]ProductAddBatchPayloadInnerAdvancedPricesInner, bool)`

GetAdvancedPricesOk returns a tuple with the AdvancedPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPrices

`func (o *ProductAddBatchPayloadInner) SetAdvancedPrices(v []ProductAddBatchPayloadInnerAdvancedPricesInner)`

SetAdvancedPrices sets AdvancedPrices field to given value.

### HasAdvancedPrices

`func (o *ProductAddBatchPayloadInner) HasAdvancedPrices() bool`

HasAdvancedPrices returns a boolean if a field has been set.

### GetFixedCostShippingPrice

`func (o *ProductAddBatchPayloadInner) GetFixedCostShippingPrice() float32`

GetFixedCostShippingPrice returns the FixedCostShippingPrice field if non-nil, zero value otherwise.

### GetFixedCostShippingPriceOk

`func (o *ProductAddBatchPayloadInner) GetFixedCostShippingPriceOk() (*float32, bool)`

GetFixedCostShippingPriceOk returns a tuple with the FixedCostShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCostShippingPrice

`func (o *ProductAddBatchPayloadInner) SetFixedCostShippingPrice(v float32)`

SetFixedCostShippingPrice sets FixedCostShippingPrice field to given value.

### HasFixedCostShippingPrice

`func (o *ProductAddBatchPayloadInner) HasFixedCostShippingPrice() bool`

HasFixedCostShippingPrice returns a boolean if a field has been set.

### GetBuyitnowPrice

`func (o *ProductAddBatchPayloadInner) GetBuyitnowPrice() float32`

GetBuyitnowPrice returns the BuyitnowPrice field if non-nil, zero value otherwise.

### GetBuyitnowPriceOk

`func (o *ProductAddBatchPayloadInner) GetBuyitnowPriceOk() (*float32, bool)`

GetBuyitnowPriceOk returns a tuple with the BuyitnowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyitnowPrice

`func (o *ProductAddBatchPayloadInner) SetBuyitnowPrice(v float32)`

SetBuyitnowPrice sets BuyitnowPrice field to given value.

### HasBuyitnowPrice

`func (o *ProductAddBatchPayloadInner) HasBuyitnowPrice() bool`

HasBuyitnowPrice returns a boolean if a field has been set.

### GetReservePrice

`func (o *ProductAddBatchPayloadInner) GetReservePrice() float32`

GetReservePrice returns the ReservePrice field if non-nil, zero value otherwise.

### GetReservePriceOk

`func (o *ProductAddBatchPayloadInner) GetReservePriceOk() (*float32, bool)`

GetReservePriceOk returns a tuple with the ReservePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservePrice

`func (o *ProductAddBatchPayloadInner) SetReservePrice(v float32)`

SetReservePrice sets ReservePrice field to given value.

### HasReservePrice

`func (o *ProductAddBatchPayloadInner) HasReservePrice() bool`

HasReservePrice returns a boolean if a field has been set.

### GetBestOffer

`func (o *ProductAddBatchPayloadInner) GetBestOffer() float32`

GetBestOffer returns the BestOffer field if non-nil, zero value otherwise.

### GetBestOfferOk

`func (o *ProductAddBatchPayloadInner) GetBestOfferOk() (*float32, bool)`

GetBestOfferOk returns a tuple with the BestOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOffer

`func (o *ProductAddBatchPayloadInner) SetBestOffer(v float32)`

SetBestOffer sets BestOffer field to given value.

### HasBestOffer

`func (o *ProductAddBatchPayloadInner) HasBestOffer() bool`

HasBestOffer returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductAddBatchPayloadInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductAddBatchPayloadInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductAddBatchPayloadInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductAddBatchPayloadInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetManageStock

`func (o *ProductAddBatchPayloadInner) GetManageStock() bool`

GetManageStock returns the ManageStock field if non-nil, zero value otherwise.

### GetManageStockOk

`func (o *ProductAddBatchPayloadInner) GetManageStockOk() (*bool, bool)`

GetManageStockOk returns a tuple with the ManageStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStock

`func (o *ProductAddBatchPayloadInner) SetManageStock(v bool)`

SetManageStock sets ManageStock field to given value.

### HasManageStock

`func (o *ProductAddBatchPayloadInner) HasManageStock() bool`

HasManageStock returns a boolean if a field has been set.

### GetProductType

`func (o *ProductAddBatchPayloadInner) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductAddBatchPayloadInner) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductAddBatchPayloadInner) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *ProductAddBatchPayloadInner) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetMarketplaceItemProperties

`func (o *ProductAddBatchPayloadInner) GetMarketplaceItemProperties() map[string]interface{}`

GetMarketplaceItemProperties returns the MarketplaceItemProperties field if non-nil, zero value otherwise.

### GetMarketplaceItemPropertiesOk

`func (o *ProductAddBatchPayloadInner) GetMarketplaceItemPropertiesOk() (*map[string]interface{}, bool)`

GetMarketplaceItemPropertiesOk returns a tuple with the MarketplaceItemProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceItemProperties

`func (o *ProductAddBatchPayloadInner) SetMarketplaceItemProperties(v map[string]interface{})`

SetMarketplaceItemProperties sets MarketplaceItemProperties field to given value.

### HasMarketplaceItemProperties

`func (o *ProductAddBatchPayloadInner) HasMarketplaceItemProperties() bool`

HasMarketplaceItemProperties returns a boolean if a field has been set.

### GetSpecifics

`func (o *ProductAddBatchPayloadInner) GetSpecifics() map[string]interface{}`

GetSpecifics returns the Specifics field if non-nil, zero value otherwise.

### GetSpecificsOk

`func (o *ProductAddBatchPayloadInner) GetSpecificsOk() (*map[string]interface{}, bool)`

GetSpecificsOk returns a tuple with the Specifics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifics

`func (o *ProductAddBatchPayloadInner) SetSpecifics(v map[string]interface{})`

SetSpecifics sets Specifics field to given value.

### HasSpecifics

`func (o *ProductAddBatchPayloadInner) HasSpecifics() bool`

HasSpecifics returns a boolean if a field has been set.

### GetIsFreeShipping

`func (o *ProductAddBatchPayloadInner) GetIsFreeShipping() bool`

GetIsFreeShipping returns the IsFreeShipping field if non-nil, zero value otherwise.

### GetIsFreeShippingOk

`func (o *ProductAddBatchPayloadInner) GetIsFreeShippingOk() (*bool, bool)`

GetIsFreeShippingOk returns a tuple with the IsFreeShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeShipping

`func (o *ProductAddBatchPayloadInner) SetIsFreeShipping(v bool)`

SetIsFreeShipping sets IsFreeShipping field to given value.

### HasIsFreeShipping

`func (o *ProductAddBatchPayloadInner) HasIsFreeShipping() bool`

HasIsFreeShipping returns a boolean if a field has been set.

### GetTaxable

`func (o *ProductAddBatchPayloadInner) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ProductAddBatchPayloadInner) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ProductAddBatchPayloadInner) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ProductAddBatchPayloadInner) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetStatus

`func (o *ProductAddBatchPayloadInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductAddBatchPayloadInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductAddBatchPayloadInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductAddBatchPayloadInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCondition

`func (o *ProductAddBatchPayloadInner) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ProductAddBatchPayloadInner) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ProductAddBatchPayloadInner) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ProductAddBatchPayloadInner) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetConditionDescription

`func (o *ProductAddBatchPayloadInner) GetConditionDescription() string`

GetConditionDescription returns the ConditionDescription field if non-nil, zero value otherwise.

### GetConditionDescriptionOk

`func (o *ProductAddBatchPayloadInner) GetConditionDescriptionOk() (*string, bool)`

GetConditionDescriptionOk returns a tuple with the ConditionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionDescription

`func (o *ProductAddBatchPayloadInner) SetConditionDescription(v string)`

SetConditionDescription sets ConditionDescription field to given value.

### HasConditionDescription

`func (o *ProductAddBatchPayloadInner) HasConditionDescription() bool`

HasConditionDescription returns a boolean if a field has been set.

### GetVisible

`func (o *ProductAddBatchPayloadInner) GetVisible() string`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ProductAddBatchPayloadInner) GetVisibleOk() (*string, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ProductAddBatchPayloadInner) SetVisible(v string)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ProductAddBatchPayloadInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetAvailableForView

`func (o *ProductAddBatchPayloadInner) GetAvailableForView() bool`

GetAvailableForView returns the AvailableForView field if non-nil, zero value otherwise.

### GetAvailableForViewOk

`func (o *ProductAddBatchPayloadInner) GetAvailableForViewOk() (*bool, bool)`

GetAvailableForViewOk returns a tuple with the AvailableForView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForView

`func (o *ProductAddBatchPayloadInner) SetAvailableForView(v bool)`

SetAvailableForView sets AvailableForView field to given value.

### HasAvailableForView

`func (o *ProductAddBatchPayloadInner) HasAvailableForView() bool`

HasAvailableForView returns a boolean if a field has been set.

### GetAvailableForSale

`func (o *ProductAddBatchPayloadInner) GetAvailableForSale() bool`

GetAvailableForSale returns the AvailableForSale field if non-nil, zero value otherwise.

### GetAvailableForSaleOk

`func (o *ProductAddBatchPayloadInner) GetAvailableForSaleOk() (*bool, bool)`

GetAvailableForSaleOk returns a tuple with the AvailableForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForSale

`func (o *ProductAddBatchPayloadInner) SetAvailableForSale(v bool)`

SetAvailableForSale sets AvailableForSale field to given value.

### HasAvailableForSale

`func (o *ProductAddBatchPayloadInner) HasAvailableForSale() bool`

HasAvailableForSale returns a boolean if a field has been set.

### GetIsVirtual

`func (o *ProductAddBatchPayloadInner) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *ProductAddBatchPayloadInner) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *ProductAddBatchPayloadInner) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *ProductAddBatchPayloadInner) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetInStock

`func (o *ProductAddBatchPayloadInner) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *ProductAddBatchPayloadInner) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *ProductAddBatchPayloadInner) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *ProductAddBatchPayloadInner) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### GetType

`func (o *ProductAddBatchPayloadInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductAddBatchPayloadInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductAddBatchPayloadInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductAddBatchPayloadInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetListingType

`func (o *ProductAddBatchPayloadInner) GetListingType() string`

GetListingType returns the ListingType field if non-nil, zero value otherwise.

### GetListingTypeOk

`func (o *ProductAddBatchPayloadInner) GetListingTypeOk() (*string, bool)`

GetListingTypeOk returns a tuple with the ListingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingType

`func (o *ProductAddBatchPayloadInner) SetListingType(v string)`

SetListingType sets ListingType field to given value.

### HasListingType

`func (o *ProductAddBatchPayloadInner) HasListingType() bool`

HasListingType returns a boolean if a field has been set.

### GetListingDuration

`func (o *ProductAddBatchPayloadInner) GetListingDuration() string`

GetListingDuration returns the ListingDuration field if non-nil, zero value otherwise.

### GetListingDurationOk

`func (o *ProductAddBatchPayloadInner) GetListingDurationOk() (*string, bool)`

GetListingDurationOk returns a tuple with the ListingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingDuration

`func (o *ProductAddBatchPayloadInner) SetListingDuration(v string)`

SetListingDuration sets ListingDuration field to given value.

### HasListingDuration

`func (o *ProductAddBatchPayloadInner) HasListingDuration() bool`

HasListingDuration returns a boolean if a field has been set.

### GetDownloadable

`func (o *ProductAddBatchPayloadInner) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *ProductAddBatchPayloadInner) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *ProductAddBatchPayloadInner) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *ProductAddBatchPayloadInner) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetWeight

`func (o *ProductAddBatchPayloadInner) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductAddBatchPayloadInner) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductAddBatchPayloadInner) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductAddBatchPayloadInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetLength

`func (o *ProductAddBatchPayloadInner) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ProductAddBatchPayloadInner) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ProductAddBatchPayloadInner) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ProductAddBatchPayloadInner) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWidth

`func (o *ProductAddBatchPayloadInner) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ProductAddBatchPayloadInner) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ProductAddBatchPayloadInner) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ProductAddBatchPayloadInner) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ProductAddBatchPayloadInner) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ProductAddBatchPayloadInner) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ProductAddBatchPayloadInner) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ProductAddBatchPayloadInner) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *ProductAddBatchPayloadInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ProductAddBatchPayloadInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ProductAddBatchPayloadInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ProductAddBatchPayloadInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetDimensionsUnit

`func (o *ProductAddBatchPayloadInner) GetDimensionsUnit() string`

GetDimensionsUnit returns the DimensionsUnit field if non-nil, zero value otherwise.

### GetDimensionsUnitOk

`func (o *ProductAddBatchPayloadInner) GetDimensionsUnitOk() (*string, bool)`

GetDimensionsUnitOk returns a tuple with the DimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsUnit

`func (o *ProductAddBatchPayloadInner) SetDimensionsUnit(v string)`

SetDimensionsUnit sets DimensionsUnit field to given value.

### HasDimensionsUnit

`func (o *ProductAddBatchPayloadInner) HasDimensionsUnit() bool`

HasDimensionsUnit returns a boolean if a field has been set.

### GetStoreId

`func (o *ProductAddBatchPayloadInner) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductAddBatchPayloadInner) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductAddBatchPayloadInner) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductAddBatchPayloadInner) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetLangId

`func (o *ProductAddBatchPayloadInner) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *ProductAddBatchPayloadInner) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *ProductAddBatchPayloadInner) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *ProductAddBatchPayloadInner) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetCategoryId

`func (o *ProductAddBatchPayloadInner) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ProductAddBatchPayloadInner) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ProductAddBatchPayloadInner) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ProductAddBatchPayloadInner) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *ProductAddBatchPayloadInner) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ProductAddBatchPayloadInner) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ProductAddBatchPayloadInner) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ProductAddBatchPayloadInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetCategoriesIds

`func (o *ProductAddBatchPayloadInner) GetCategoriesIds() []string`

GetCategoriesIds returns the CategoriesIds field if non-nil, zero value otherwise.

### GetCategoriesIdsOk

`func (o *ProductAddBatchPayloadInner) GetCategoriesIdsOk() (*[]string, bool)`

GetCategoriesIdsOk returns a tuple with the CategoriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesIds

`func (o *ProductAddBatchPayloadInner) SetCategoriesIds(v []string)`

SetCategoriesIds sets CategoriesIds field to given value.

### HasCategoriesIds

`func (o *ProductAddBatchPayloadInner) HasCategoriesIds() bool`

HasCategoriesIds returns a boolean if a field has been set.

### GetRelatedProductsIds

`func (o *ProductAddBatchPayloadInner) GetRelatedProductsIds() []string`

GetRelatedProductsIds returns the RelatedProductsIds field if non-nil, zero value otherwise.

### GetRelatedProductsIdsOk

`func (o *ProductAddBatchPayloadInner) GetRelatedProductsIdsOk() (*[]string, bool)`

GetRelatedProductsIdsOk returns a tuple with the RelatedProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedProductsIds

`func (o *ProductAddBatchPayloadInner) SetRelatedProductsIds(v []string)`

SetRelatedProductsIds sets RelatedProductsIds field to given value.

### HasRelatedProductsIds

`func (o *ProductAddBatchPayloadInner) HasRelatedProductsIds() bool`

HasRelatedProductsIds returns a boolean if a field has been set.

### GetUpSellProductsIds

`func (o *ProductAddBatchPayloadInner) GetUpSellProductsIds() []string`

GetUpSellProductsIds returns the UpSellProductsIds field if non-nil, zero value otherwise.

### GetUpSellProductsIdsOk

`func (o *ProductAddBatchPayloadInner) GetUpSellProductsIdsOk() (*[]string, bool)`

GetUpSellProductsIdsOk returns a tuple with the UpSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSellProductsIds

`func (o *ProductAddBatchPayloadInner) SetUpSellProductsIds(v []string)`

SetUpSellProductsIds sets UpSellProductsIds field to given value.

### HasUpSellProductsIds

`func (o *ProductAddBatchPayloadInner) HasUpSellProductsIds() bool`

HasUpSellProductsIds returns a boolean if a field has been set.

### GetCrossSellProductsIds

`func (o *ProductAddBatchPayloadInner) GetCrossSellProductsIds() []string`

GetCrossSellProductsIds returns the CrossSellProductsIds field if non-nil, zero value otherwise.

### GetCrossSellProductsIdsOk

`func (o *ProductAddBatchPayloadInner) GetCrossSellProductsIdsOk() (*[]string, bool)`

GetCrossSellProductsIdsOk returns a tuple with the CrossSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSellProductsIds

`func (o *ProductAddBatchPayloadInner) SetCrossSellProductsIds(v []string)`

SetCrossSellProductsIds sets CrossSellProductsIds field to given value.

### HasCrossSellProductsIds

`func (o *ProductAddBatchPayloadInner) HasCrossSellProductsIds() bool`

HasCrossSellProductsIds returns a boolean if a field has been set.

### GetStoresIds

`func (o *ProductAddBatchPayloadInner) GetStoresIds() []string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *ProductAddBatchPayloadInner) GetStoresIdsOk() (*[]string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *ProductAddBatchPayloadInner) SetStoresIds(v []string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *ProductAddBatchPayloadInner) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetTaxClassId

`func (o *ProductAddBatchPayloadInner) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *ProductAddBatchPayloadInner) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *ProductAddBatchPayloadInner) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *ProductAddBatchPayloadInner) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### GetSalesTax

`func (o *ProductAddBatchPayloadInner) GetSalesTax() ProductAddBatchPayloadInnerSalesTax`

GetSalesTax returns the SalesTax field if non-nil, zero value otherwise.

### GetSalesTaxOk

`func (o *ProductAddBatchPayloadInner) GetSalesTaxOk() (*ProductAddBatchPayloadInnerSalesTax, bool)`

GetSalesTaxOk returns a tuple with the SalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTax

`func (o *ProductAddBatchPayloadInner) SetSalesTax(v ProductAddBatchPayloadInnerSalesTax)`

SetSalesTax sets SalesTax field to given value.

### HasSalesTax

`func (o *ProductAddBatchPayloadInner) HasSalesTax() bool`

HasSalesTax returns a boolean if a field has been set.

### GetMetaTitle

`func (o *ProductAddBatchPayloadInner) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *ProductAddBatchPayloadInner) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *ProductAddBatchPayloadInner) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *ProductAddBatchPayloadInner) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetMetaDescription

`func (o *ProductAddBatchPayloadInner) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *ProductAddBatchPayloadInner) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *ProductAddBatchPayloadInner) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *ProductAddBatchPayloadInner) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *ProductAddBatchPayloadInner) GetMetaKeywords() []string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *ProductAddBatchPayloadInner) GetMetaKeywordsOk() (*[]string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *ProductAddBatchPayloadInner) SetMetaKeywords(v []string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *ProductAddBatchPayloadInner) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetSearchKeywords

`func (o *ProductAddBatchPayloadInner) GetSearchKeywords() []string`

GetSearchKeywords returns the SearchKeywords field if non-nil, zero value otherwise.

### GetSearchKeywordsOk

`func (o *ProductAddBatchPayloadInner) GetSearchKeywordsOk() (*[]string, bool)`

GetSearchKeywordsOk returns a tuple with the SearchKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKeywords

`func (o *ProductAddBatchPayloadInner) SetSearchKeywords(v []string)`

SetSearchKeywords sets SearchKeywords field to given value.

### HasSearchKeywords

`func (o *ProductAddBatchPayloadInner) HasSearchKeywords() bool`

HasSearchKeywords returns a boolean if a field has been set.

### GetHarmonizedSystemCode

`func (o *ProductAddBatchPayloadInner) GetHarmonizedSystemCode() string`

GetHarmonizedSystemCode returns the HarmonizedSystemCode field if non-nil, zero value otherwise.

### GetHarmonizedSystemCodeOk

`func (o *ProductAddBatchPayloadInner) GetHarmonizedSystemCodeOk() (*string, bool)`

GetHarmonizedSystemCodeOk returns a tuple with the HarmonizedSystemCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarmonizedSystemCode

`func (o *ProductAddBatchPayloadInner) SetHarmonizedSystemCode(v string)`

SetHarmonizedSystemCode sets HarmonizedSystemCode field to given value.

### HasHarmonizedSystemCode

`func (o *ProductAddBatchPayloadInner) HasHarmonizedSystemCode() bool`

HasHarmonizedSystemCode returns a boolean if a field has been set.

### GetUrl

`func (o *ProductAddBatchPayloadInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProductAddBatchPayloadInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProductAddBatchPayloadInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProductAddBatchPayloadInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSeoUrl

`func (o *ProductAddBatchPayloadInner) GetSeoUrl() string`

GetSeoUrl returns the SeoUrl field if non-nil, zero value otherwise.

### GetSeoUrlOk

`func (o *ProductAddBatchPayloadInner) GetSeoUrlOk() (*string, bool)`

GetSeoUrlOk returns a tuple with the SeoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoUrl

`func (o *ProductAddBatchPayloadInner) SetSeoUrl(v string)`

SetSeoUrl sets SeoUrl field to given value.

### HasSeoUrl

`func (o *ProductAddBatchPayloadInner) HasSeoUrl() bool`

HasSeoUrl returns a boolean if a field has been set.

### GetExternalProductLink

`func (o *ProductAddBatchPayloadInner) GetExternalProductLink() string`

GetExternalProductLink returns the ExternalProductLink field if non-nil, zero value otherwise.

### GetExternalProductLinkOk

`func (o *ProductAddBatchPayloadInner) GetExternalProductLinkOk() (*string, bool)`

GetExternalProductLinkOk returns a tuple with the ExternalProductLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProductLink

`func (o *ProductAddBatchPayloadInner) SetExternalProductLink(v string)`

SetExternalProductLink sets ExternalProductLink field to given value.

### HasExternalProductLink

`func (o *ProductAddBatchPayloadInner) HasExternalProductLink() bool`

HasExternalProductLink returns a boolean if a field has been set.

### GetManufacturer

`func (o *ProductAddBatchPayloadInner) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ProductAddBatchPayloadInner) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ProductAddBatchPayloadInner) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ProductAddBatchPayloadInner) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetManufacturerId

`func (o *ProductAddBatchPayloadInner) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *ProductAddBatchPayloadInner) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *ProductAddBatchPayloadInner) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *ProductAddBatchPayloadInner) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetBackorderStatus

`func (o *ProductAddBatchPayloadInner) GetBackorderStatus() string`

GetBackorderStatus returns the BackorderStatus field if non-nil, zero value otherwise.

### GetBackorderStatusOk

`func (o *ProductAddBatchPayloadInner) GetBackorderStatusOk() (*string, bool)`

GetBackorderStatusOk returns a tuple with the BackorderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderStatus

`func (o *ProductAddBatchPayloadInner) SetBackorderStatus(v string)`

SetBackorderStatus sets BackorderStatus field to given value.

### HasBackorderStatus

`func (o *ProductAddBatchPayloadInner) HasBackorderStatus() bool`

HasBackorderStatus returns a boolean if a field has been set.

### GetImages

`func (o *ProductAddBatchPayloadInner) GetImages() []ProductAddBatchPayloadInnerImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ProductAddBatchPayloadInner) GetImagesOk() (*[]ProductAddBatchPayloadInnerImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ProductAddBatchPayloadInner) SetImages(v []ProductAddBatchPayloadInnerImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *ProductAddBatchPayloadInner) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetTags

`func (o *ProductAddBatchPayloadInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductAddBatchPayloadInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductAddBatchPayloadInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProductAddBatchPayloadInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFiles

`func (o *ProductAddBatchPayloadInner) GetFiles() []ProductAddFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProductAddBatchPayloadInner) GetFilesOk() (*[]ProductAddFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProductAddBatchPayloadInner) SetFiles(v []ProductAddFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ProductAddBatchPayloadInner) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


