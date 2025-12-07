# ProductVariantAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** | Defines product&#39;s id where the variant has to be added | [optional] 
**Attributes** | Pointer to [**[]ProductVariantAddAttributesInner**](ProductVariantAddAttributesInner.md) | Defines variant&#39;s attributes list | [optional] 
**Name** | Pointer to **string** | Defines variant&#39;s name that has to be added | [optional] 
**Model** | **string** | Specifies variant&#39;s model that has to be added | 
**Description** | Pointer to **string** | Specifies variant&#39;s description | [optional] 
**ShortDescription** | Pointer to **string** | Defines short description | [optional] 
**AvailableForView** | Pointer to **bool** | Specifies the set of visible/invisible product&#39;s variants for users | [optional] [default to true]
**AvailableForSale** | Pointer to **bool** | Specifies the set of visible/invisible product&#39;s variants for sale | [optional] [default to true]
**Status** | Pointer to **string** | Defines status | [optional] 
**IsVirtual** | Pointer to **bool** | Defines whether the product is virtual | [optional] [default to false]
**IsDefault** | Pointer to **bool** | Defines as a default variant | [optional] 
**StoreId** | Pointer to **string** | Add variants specified by store id | [optional] 
**StoresIds** | Pointer to **string** | Assign variant to the stores that is specified by comma-separated stores&#39; id | [optional] 
**LangId** | Pointer to **string** | Language id | [optional] 
**Price** | Pointer to **float32** | Defines new product&#39;s variant price | [optional] 
**OldPrice** | Pointer to **float32** | Defines product&#39;s old price | [optional] 
**CostPrice** | Pointer to **float32** | Defines new product&#39;s cost price | [optional] 
**SpecialPrice** | Pointer to **float32** | Specifies variant&#39;s model that has to be added | [optional] 
**SpriceCreate** | Pointer to **string** | Defines the date of special price creation | [optional] 
**SpriceModified** | Pointer to **string** | Defines the date of special price modification | [optional] 
**SpriceExpire** | Pointer to **string** | Defines the term of special price offer duration | [optional] 
**TierPrices** | Pointer to [**[]ProductAddTierPricesInner**](ProductAddTierPricesInner.md) | Defines product&#39;s tier prices | [optional] 
**Quantity** | Pointer to **float32** | Defines product variant&#39;s quantity that has to be added | [optional] [default to 0]
**WarehouseId** | Pointer to **string** | This parameter is used for selecting a warehouse where you need to set/modify a product quantity. | [optional] 
**InStock** | Pointer to **bool** | Set stock status | [optional] 
**BackorderStatus** | Pointer to **string** | Set backorder status | [optional] 
**ManageStock** | Pointer to **bool** | Defines inventory tracking for product variant | [optional] 
**LowStockThreshold** | Pointer to **float32** | Specify the quantity threshold below which the product is considered low in stock | [optional] 
**Weight** | Pointer to **float32** | Weight | [optional] [default to 0]
**Width** | Pointer to **float32** | Defines product&#39;s width | [optional] 
**Height** | Pointer to **float32** | Defines product&#39;s height | [optional] 
**Length** | Pointer to **float32** | Defines product&#39;s length | [optional] 
**WeightUnit** | Pointer to **string** | Weight Unit | [optional] 
**Sku** | Pointer to **string** | Defines variant&#39;s sku that has to be added | [optional] 
**Barcode** | Pointer to **string** | A barcode is a unique code composed of numbers used as a product identifier. | [optional] 
**Gtin** | Pointer to **string** | Global Trade Item Number. An GTIN is an identifier for trade items. | [optional] 
**Upc** | Pointer to **string** | Universal Product Code. A UPC (UPC-A) is a commonly used identifer for many different products. | [optional] 
**Ean** | Pointer to **string** | European Article Number. An EAN is a unique 8 or 13-digit identifier that many industries (such as book publishers) use to identify products. | [optional] 
**Mpn** | Pointer to **string** | Manufacturer Part Number. A MPN is an identifier of a particular part design or material used. | [optional] 
**Isbn** | Pointer to **string** | International Standard Book Number. An ISBN is a unique identifier for books. | [optional] 
**SeoUrl** | Pointer to **string** | Defines unique URL for SEO | [optional] 
**Manufacturer** | Pointer to **string** | Specifies the product variant&#39;s manufacturer | [optional] 
**CreatedAt** | Pointer to **string** | Defines the date of entity creation | [optional] 
**MetaTitle** | Pointer to **string** | Defines unique meta title for each entity | [optional] 
**MetaKeywords** | Pointer to **string** | Defines unique meta keywords for each entity | [optional] 
**MetaDescription** | Pointer to **string** | Defines unique meta description of a entity | [optional] 
**Url** | Pointer to **string** | Defines unique product variant&#39;s URL | [optional] 
**TaxClassId** | Pointer to **string** | Defines tax classes where entity has to be added | [optional] 
**Taxable** | Pointer to **bool** | Specifies whether a tax is charged | [optional] [default to true]
**FixedCostShippingPrice** | Pointer to **float32** | Specifies fixed cost shipping price | [optional] 
**IsFreeShipping** | Pointer to **bool** | Specifies variant&#39;s free shipping flag that has to be added | [optional] 
**CountryOfOrigin** | Pointer to **string** | The country where the inventory item was made | [optional] 
**HarmonizedSystemCode** | Pointer to **string** | Harmonized System Code. An HSC is a 6-digit identifier that allows participating countries to classify traded goods on a common basis for customs purposes | [optional] 
**ProcessingProfileId** | Pointer to **int32** | The numeric ID of the processing profile (readiness state) for physical products in Etsy. You can find possible values in the \&quot;cart.info\&quot; API method response, in the field processing_profiles[]-&gt;readiness_state_id. | [optional] 
**MarketplaceItemProperties** | Pointer to **string** | String containing the JSON representation of the supplied data | [optional] 
**ClearCache** | Pointer to **bool** | Is cache clear required | [optional] [default to true]

## Methods

### NewProductVariantAdd

`func NewProductVariantAdd(model string, ) *ProductVariantAdd`

NewProductVariantAdd instantiates a new ProductVariantAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariantAddWithDefaults

`func NewProductVariantAddWithDefaults() *ProductVariantAdd`

NewProductVariantAddWithDefaults instantiates a new ProductVariantAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProductVariantAdd) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductVariantAdd) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductVariantAdd) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductVariantAdd) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetAttributes

`func (o *ProductVariantAdd) GetAttributes() []ProductVariantAddAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProductVariantAdd) GetAttributesOk() (*[]ProductVariantAddAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProductVariantAdd) SetAttributes(v []ProductVariantAddAttributesInner)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ProductVariantAdd) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetName

`func (o *ProductVariantAdd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductVariantAdd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductVariantAdd) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductVariantAdd) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *ProductVariantAdd) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ProductVariantAdd) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ProductVariantAdd) SetModel(v string)`

SetModel sets Model field to given value.


### GetDescription

`func (o *ProductVariantAdd) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductVariantAdd) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductVariantAdd) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductVariantAdd) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortDescription

`func (o *ProductVariantAdd) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ProductVariantAdd) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ProductVariantAdd) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ProductVariantAdd) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetAvailableForView

`func (o *ProductVariantAdd) GetAvailableForView() bool`

GetAvailableForView returns the AvailableForView field if non-nil, zero value otherwise.

### GetAvailableForViewOk

`func (o *ProductVariantAdd) GetAvailableForViewOk() (*bool, bool)`

GetAvailableForViewOk returns a tuple with the AvailableForView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForView

`func (o *ProductVariantAdd) SetAvailableForView(v bool)`

SetAvailableForView sets AvailableForView field to given value.

### HasAvailableForView

`func (o *ProductVariantAdd) HasAvailableForView() bool`

HasAvailableForView returns a boolean if a field has been set.

### GetAvailableForSale

`func (o *ProductVariantAdd) GetAvailableForSale() bool`

GetAvailableForSale returns the AvailableForSale field if non-nil, zero value otherwise.

### GetAvailableForSaleOk

`func (o *ProductVariantAdd) GetAvailableForSaleOk() (*bool, bool)`

GetAvailableForSaleOk returns a tuple with the AvailableForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForSale

`func (o *ProductVariantAdd) SetAvailableForSale(v bool)`

SetAvailableForSale sets AvailableForSale field to given value.

### HasAvailableForSale

`func (o *ProductVariantAdd) HasAvailableForSale() bool`

HasAvailableForSale returns a boolean if a field has been set.

### GetStatus

`func (o *ProductVariantAdd) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductVariantAdd) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductVariantAdd) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductVariantAdd) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsVirtual

`func (o *ProductVariantAdd) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *ProductVariantAdd) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *ProductVariantAdd) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *ProductVariantAdd) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetIsDefault

`func (o *ProductVariantAdd) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ProductVariantAdd) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ProductVariantAdd) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ProductVariantAdd) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetStoreId

`func (o *ProductVariantAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductVariantAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductVariantAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductVariantAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetStoresIds

`func (o *ProductVariantAdd) GetStoresIds() string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *ProductVariantAdd) GetStoresIdsOk() (*string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *ProductVariantAdd) SetStoresIds(v string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *ProductVariantAdd) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetLangId

`func (o *ProductVariantAdd) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *ProductVariantAdd) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *ProductVariantAdd) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *ProductVariantAdd) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetPrice

`func (o *ProductVariantAdd) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductVariantAdd) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductVariantAdd) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductVariantAdd) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOldPrice

`func (o *ProductVariantAdd) GetOldPrice() float32`

GetOldPrice returns the OldPrice field if non-nil, zero value otherwise.

### GetOldPriceOk

`func (o *ProductVariantAdd) GetOldPriceOk() (*float32, bool)`

GetOldPriceOk returns a tuple with the OldPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPrice

`func (o *ProductVariantAdd) SetOldPrice(v float32)`

SetOldPrice sets OldPrice field to given value.

### HasOldPrice

`func (o *ProductVariantAdd) HasOldPrice() bool`

HasOldPrice returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductVariantAdd) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductVariantAdd) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductVariantAdd) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductVariantAdd) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetSpecialPrice

`func (o *ProductVariantAdd) GetSpecialPrice() float32`

GetSpecialPrice returns the SpecialPrice field if non-nil, zero value otherwise.

### GetSpecialPriceOk

`func (o *ProductVariantAdd) GetSpecialPriceOk() (*float32, bool)`

GetSpecialPriceOk returns a tuple with the SpecialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPrice

`func (o *ProductVariantAdd) SetSpecialPrice(v float32)`

SetSpecialPrice sets SpecialPrice field to given value.

### HasSpecialPrice

`func (o *ProductVariantAdd) HasSpecialPrice() bool`

HasSpecialPrice returns a boolean if a field has been set.

### GetSpriceCreate

`func (o *ProductVariantAdd) GetSpriceCreate() string`

GetSpriceCreate returns the SpriceCreate field if non-nil, zero value otherwise.

### GetSpriceCreateOk

`func (o *ProductVariantAdd) GetSpriceCreateOk() (*string, bool)`

GetSpriceCreateOk returns a tuple with the SpriceCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceCreate

`func (o *ProductVariantAdd) SetSpriceCreate(v string)`

SetSpriceCreate sets SpriceCreate field to given value.

### HasSpriceCreate

`func (o *ProductVariantAdd) HasSpriceCreate() bool`

HasSpriceCreate returns a boolean if a field has been set.

### GetSpriceModified

`func (o *ProductVariantAdd) GetSpriceModified() string`

GetSpriceModified returns the SpriceModified field if non-nil, zero value otherwise.

### GetSpriceModifiedOk

`func (o *ProductVariantAdd) GetSpriceModifiedOk() (*string, bool)`

GetSpriceModifiedOk returns a tuple with the SpriceModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceModified

`func (o *ProductVariantAdd) SetSpriceModified(v string)`

SetSpriceModified sets SpriceModified field to given value.

### HasSpriceModified

`func (o *ProductVariantAdd) HasSpriceModified() bool`

HasSpriceModified returns a boolean if a field has been set.

### GetSpriceExpire

`func (o *ProductVariantAdd) GetSpriceExpire() string`

GetSpriceExpire returns the SpriceExpire field if non-nil, zero value otherwise.

### GetSpriceExpireOk

`func (o *ProductVariantAdd) GetSpriceExpireOk() (*string, bool)`

GetSpriceExpireOk returns a tuple with the SpriceExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceExpire

`func (o *ProductVariantAdd) SetSpriceExpire(v string)`

SetSpriceExpire sets SpriceExpire field to given value.

### HasSpriceExpire

`func (o *ProductVariantAdd) HasSpriceExpire() bool`

HasSpriceExpire returns a boolean if a field has been set.

### GetTierPrices

`func (o *ProductVariantAdd) GetTierPrices() []ProductAddTierPricesInner`

GetTierPrices returns the TierPrices field if non-nil, zero value otherwise.

### GetTierPricesOk

`func (o *ProductVariantAdd) GetTierPricesOk() (*[]ProductAddTierPricesInner, bool)`

GetTierPricesOk returns a tuple with the TierPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierPrices

`func (o *ProductVariantAdd) SetTierPrices(v []ProductAddTierPricesInner)`

SetTierPrices sets TierPrices field to given value.

### HasTierPrices

`func (o *ProductVariantAdd) HasTierPrices() bool`

HasTierPrices returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductVariantAdd) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductVariantAdd) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductVariantAdd) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductVariantAdd) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetWarehouseId

`func (o *ProductVariantAdd) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ProductVariantAdd) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ProductVariantAdd) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ProductVariantAdd) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetInStock

`func (o *ProductVariantAdd) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *ProductVariantAdd) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *ProductVariantAdd) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *ProductVariantAdd) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### GetBackorderStatus

`func (o *ProductVariantAdd) GetBackorderStatus() string`

GetBackorderStatus returns the BackorderStatus field if non-nil, zero value otherwise.

### GetBackorderStatusOk

`func (o *ProductVariantAdd) GetBackorderStatusOk() (*string, bool)`

GetBackorderStatusOk returns a tuple with the BackorderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderStatus

`func (o *ProductVariantAdd) SetBackorderStatus(v string)`

SetBackorderStatus sets BackorderStatus field to given value.

### HasBackorderStatus

`func (o *ProductVariantAdd) HasBackorderStatus() bool`

HasBackorderStatus returns a boolean if a field has been set.

### GetManageStock

`func (o *ProductVariantAdd) GetManageStock() bool`

GetManageStock returns the ManageStock field if non-nil, zero value otherwise.

### GetManageStockOk

`func (o *ProductVariantAdd) GetManageStockOk() (*bool, bool)`

GetManageStockOk returns a tuple with the ManageStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStock

`func (o *ProductVariantAdd) SetManageStock(v bool)`

SetManageStock sets ManageStock field to given value.

### HasManageStock

`func (o *ProductVariantAdd) HasManageStock() bool`

HasManageStock returns a boolean if a field has been set.

### GetLowStockThreshold

`func (o *ProductVariantAdd) GetLowStockThreshold() float32`

GetLowStockThreshold returns the LowStockThreshold field if non-nil, zero value otherwise.

### GetLowStockThresholdOk

`func (o *ProductVariantAdd) GetLowStockThresholdOk() (*float32, bool)`

GetLowStockThresholdOk returns a tuple with the LowStockThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowStockThreshold

`func (o *ProductVariantAdd) SetLowStockThreshold(v float32)`

SetLowStockThreshold sets LowStockThreshold field to given value.

### HasLowStockThreshold

`func (o *ProductVariantAdd) HasLowStockThreshold() bool`

HasLowStockThreshold returns a boolean if a field has been set.

### GetWeight

`func (o *ProductVariantAdd) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductVariantAdd) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductVariantAdd) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductVariantAdd) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWidth

`func (o *ProductVariantAdd) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ProductVariantAdd) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ProductVariantAdd) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ProductVariantAdd) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ProductVariantAdd) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ProductVariantAdd) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ProductVariantAdd) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ProductVariantAdd) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetLength

`func (o *ProductVariantAdd) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ProductVariantAdd) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ProductVariantAdd) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ProductVariantAdd) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWeightUnit

`func (o *ProductVariantAdd) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ProductVariantAdd) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ProductVariantAdd) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ProductVariantAdd) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetSku

`func (o *ProductVariantAdd) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductVariantAdd) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductVariantAdd) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductVariantAdd) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetBarcode

`func (o *ProductVariantAdd) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductVariantAdd) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductVariantAdd) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ProductVariantAdd) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetGtin

`func (o *ProductVariantAdd) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *ProductVariantAdd) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *ProductVariantAdd) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *ProductVariantAdd) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### GetUpc

`func (o *ProductVariantAdd) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ProductVariantAdd) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ProductVariantAdd) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ProductVariantAdd) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetEan

`func (o *ProductVariantAdd) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *ProductVariantAdd) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *ProductVariantAdd) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *ProductVariantAdd) HasEan() bool`

HasEan returns a boolean if a field has been set.

### GetMpn

`func (o *ProductVariantAdd) GetMpn() string`

GetMpn returns the Mpn field if non-nil, zero value otherwise.

### GetMpnOk

`func (o *ProductVariantAdd) GetMpnOk() (*string, bool)`

GetMpnOk returns a tuple with the Mpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpn

`func (o *ProductVariantAdd) SetMpn(v string)`

SetMpn sets Mpn field to given value.

### HasMpn

`func (o *ProductVariantAdd) HasMpn() bool`

HasMpn returns a boolean if a field has been set.

### GetIsbn

`func (o *ProductVariantAdd) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *ProductVariantAdd) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *ProductVariantAdd) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *ProductVariantAdd) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### GetSeoUrl

`func (o *ProductVariantAdd) GetSeoUrl() string`

GetSeoUrl returns the SeoUrl field if non-nil, zero value otherwise.

### GetSeoUrlOk

`func (o *ProductVariantAdd) GetSeoUrlOk() (*string, bool)`

GetSeoUrlOk returns a tuple with the SeoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoUrl

`func (o *ProductVariantAdd) SetSeoUrl(v string)`

SetSeoUrl sets SeoUrl field to given value.

### HasSeoUrl

`func (o *ProductVariantAdd) HasSeoUrl() bool`

HasSeoUrl returns a boolean if a field has been set.

### GetManufacturer

`func (o *ProductVariantAdd) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ProductVariantAdd) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ProductVariantAdd) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ProductVariantAdd) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductVariantAdd) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductVariantAdd) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductVariantAdd) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProductVariantAdd) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMetaTitle

`func (o *ProductVariantAdd) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *ProductVariantAdd) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *ProductVariantAdd) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *ProductVariantAdd) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *ProductVariantAdd) GetMetaKeywords() string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *ProductVariantAdd) GetMetaKeywordsOk() (*string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *ProductVariantAdd) SetMetaKeywords(v string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *ProductVariantAdd) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetMetaDescription

`func (o *ProductVariantAdd) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *ProductVariantAdd) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *ProductVariantAdd) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *ProductVariantAdd) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetUrl

`func (o *ProductVariantAdd) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProductVariantAdd) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProductVariantAdd) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProductVariantAdd) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTaxClassId

`func (o *ProductVariantAdd) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *ProductVariantAdd) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *ProductVariantAdd) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *ProductVariantAdd) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### GetTaxable

`func (o *ProductVariantAdd) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ProductVariantAdd) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ProductVariantAdd) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ProductVariantAdd) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetFixedCostShippingPrice

`func (o *ProductVariantAdd) GetFixedCostShippingPrice() float32`

GetFixedCostShippingPrice returns the FixedCostShippingPrice field if non-nil, zero value otherwise.

### GetFixedCostShippingPriceOk

`func (o *ProductVariantAdd) GetFixedCostShippingPriceOk() (*float32, bool)`

GetFixedCostShippingPriceOk returns a tuple with the FixedCostShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCostShippingPrice

`func (o *ProductVariantAdd) SetFixedCostShippingPrice(v float32)`

SetFixedCostShippingPrice sets FixedCostShippingPrice field to given value.

### HasFixedCostShippingPrice

`func (o *ProductVariantAdd) HasFixedCostShippingPrice() bool`

HasFixedCostShippingPrice returns a boolean if a field has been set.

### GetIsFreeShipping

`func (o *ProductVariantAdd) GetIsFreeShipping() bool`

GetIsFreeShipping returns the IsFreeShipping field if non-nil, zero value otherwise.

### GetIsFreeShippingOk

`func (o *ProductVariantAdd) GetIsFreeShippingOk() (*bool, bool)`

GetIsFreeShippingOk returns a tuple with the IsFreeShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeShipping

`func (o *ProductVariantAdd) SetIsFreeShipping(v bool)`

SetIsFreeShipping sets IsFreeShipping field to given value.

### HasIsFreeShipping

`func (o *ProductVariantAdd) HasIsFreeShipping() bool`

HasIsFreeShipping returns a boolean if a field has been set.

### GetCountryOfOrigin

`func (o *ProductVariantAdd) GetCountryOfOrigin() string`

GetCountryOfOrigin returns the CountryOfOrigin field if non-nil, zero value otherwise.

### GetCountryOfOriginOk

`func (o *ProductVariantAdd) GetCountryOfOriginOk() (*string, bool)`

GetCountryOfOriginOk returns a tuple with the CountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfOrigin

`func (o *ProductVariantAdd) SetCountryOfOrigin(v string)`

SetCountryOfOrigin sets CountryOfOrigin field to given value.

### HasCountryOfOrigin

`func (o *ProductVariantAdd) HasCountryOfOrigin() bool`

HasCountryOfOrigin returns a boolean if a field has been set.

### GetHarmonizedSystemCode

`func (o *ProductVariantAdd) GetHarmonizedSystemCode() string`

GetHarmonizedSystemCode returns the HarmonizedSystemCode field if non-nil, zero value otherwise.

### GetHarmonizedSystemCodeOk

`func (o *ProductVariantAdd) GetHarmonizedSystemCodeOk() (*string, bool)`

GetHarmonizedSystemCodeOk returns a tuple with the HarmonizedSystemCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarmonizedSystemCode

`func (o *ProductVariantAdd) SetHarmonizedSystemCode(v string)`

SetHarmonizedSystemCode sets HarmonizedSystemCode field to given value.

### HasHarmonizedSystemCode

`func (o *ProductVariantAdd) HasHarmonizedSystemCode() bool`

HasHarmonizedSystemCode returns a boolean if a field has been set.

### GetProcessingProfileId

`func (o *ProductVariantAdd) GetProcessingProfileId() int32`

GetProcessingProfileId returns the ProcessingProfileId field if non-nil, zero value otherwise.

### GetProcessingProfileIdOk

`func (o *ProductVariantAdd) GetProcessingProfileIdOk() (*int32, bool)`

GetProcessingProfileIdOk returns a tuple with the ProcessingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingProfileId

`func (o *ProductVariantAdd) SetProcessingProfileId(v int32)`

SetProcessingProfileId sets ProcessingProfileId field to given value.

### HasProcessingProfileId

`func (o *ProductVariantAdd) HasProcessingProfileId() bool`

HasProcessingProfileId returns a boolean if a field has been set.

### GetMarketplaceItemProperties

`func (o *ProductVariantAdd) GetMarketplaceItemProperties() string`

GetMarketplaceItemProperties returns the MarketplaceItemProperties field if non-nil, zero value otherwise.

### GetMarketplaceItemPropertiesOk

`func (o *ProductVariantAdd) GetMarketplaceItemPropertiesOk() (*string, bool)`

GetMarketplaceItemPropertiesOk returns a tuple with the MarketplaceItemProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceItemProperties

`func (o *ProductVariantAdd) SetMarketplaceItemProperties(v string)`

SetMarketplaceItemProperties sets MarketplaceItemProperties field to given value.

### HasMarketplaceItemProperties

`func (o *ProductVariantAdd) HasMarketplaceItemProperties() bool`

HasMarketplaceItemProperties returns a boolean if a field has been set.

### GetClearCache

`func (o *ProductVariantAdd) GetClearCache() bool`

GetClearCache returns the ClearCache field if non-nil, zero value otherwise.

### GetClearCacheOk

`func (o *ProductVariantAdd) GetClearCacheOk() (*bool, bool)`

GetClearCacheOk returns a tuple with the ClearCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCache

`func (o *ProductVariantAdd) SetClearCache(v bool)`

SetClearCache sets ClearCache field to given value.

### HasClearCache

`func (o *ProductVariantAdd) HasClearCache() bool`

HasClearCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


