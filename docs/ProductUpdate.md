# ProductUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Defines product id that has to be updated | [optional] 
**Model** | Pointer to **string** | Defines product model that has to be updated | [optional] 
**Sku** | Pointer to **string** | Defines new product&#39;s sku | [optional] 
**Name** | Pointer to **string** | Defines product&#39;s name that has to be updated | [optional] 
**Description** | Pointer to **string** | Defines new product&#39;s description | [optional] 
**ShortDescription** | Pointer to **string** | Defines short description | [optional] 
**Price** | Pointer to **float32** | Defines new product&#39;s price | [optional] 
**OldPrice** | Pointer to **float32** | Defines product&#39;s old price | [optional] 
**SpecialPrice** | Pointer to **float32** | Defines new product&#39;s special price | [optional] 
**SpriceCreate** | Pointer to **string** | Defines the date of special price creation | [optional] 
**SpriceExpire** | Pointer to **string** | Defines the term of special price offer duration | [optional] 
**CostPrice** | Pointer to **float32** | Defines new product&#39;s cost price | [optional] 
**FixedCostShippingPrice** | Pointer to **float32** | Specifies product&#39;s fixed cost shipping price | [optional] 
**RetailPrice** | Pointer to **float32** | Defines new product&#39;s retail price | [optional] 
**TierPrices** | Pointer to [**[]ProductAddTierPricesInner**](ProductAddTierPricesInner.md) | Defines product&#39;s tier prices | [optional] 
**ReservePrice** | Pointer to **float32** | Defines reserve price value | [optional] 
**BuyitnowPrice** | Pointer to **float32** | Defines buy it now value | [optional] 
**Taxable** | Pointer to **bool** | Specifies whether a tax is charged | [optional] 
**TaxClassId** | Pointer to **string** | Defines tax classes where entity has to be added | [optional] 
**Type** | Pointer to **string** | Defines product&#39;s type | [optional] 
**Status** | Pointer to **string** | Defines product&#39;s status | [optional] 
**Condition** | Pointer to **string** | The human-readable label for the condition (e.g., \&quot;New\&quot;). | [optional] 
**Visible** | Pointer to **string** | Set visibility status | [optional] 
**InStock** | Pointer to **bool** | Set stock status | [optional] 
**Avail** | Pointer to **bool** | Defines category&#39;s visibility status | [optional] [default to true]
**AvailFrom** | Pointer to **string** | Allows to schedule a time in the future that the item becomes available. The value should be greater than the current date and time. | [optional] 
**ProductClass** | Pointer to **string** | A categorization for the product | [optional] 
**BrandName** | Pointer to **string** | Retrieves brands specified by brand name | [optional] 
**AvailableForView** | Pointer to **bool** | Specifies the set of visible/invisible products for users | [optional] 
**StoresIds** | Pointer to **string** | Assign product to the stores that is specified by comma-separated stores&#39; id | [optional] 
**StoreId** | Pointer to **string** | Defines store id where the product should be found | [optional] 
**LangId** | Pointer to **string** | Language id | [optional] 
**Quantity** | Pointer to **float32** | Defines new product&#39;s quantity | [optional] 
**ReserveQuantity** | Pointer to **float32** | This parameter allows to reserve/unreserve product quantity. | [optional] 
**ManageStock** | Pointer to **bool** | Defines inventory tracking for product | [optional] 
**BackorderStatus** | Pointer to **string** | Set backorder status | [optional] 
**IncreaseQuantity** | Pointer to **float32** | Defines the incremental changes in product quantity | [optional] 
**ReduceQuantity** | Pointer to **float32** | Defines the decrement changes in product quantity | [optional] 
**LowStockThreshold** | Pointer to **float32** | Specify the quantity threshold below which the product is considered low in stock | [optional] 
**WarehouseId** | Pointer to **string** | This parameter is used for selecting a warehouse where you need to set/modify a product quantity. | [optional] 
**Weight** | Pointer to **float32** | Weight | [optional] 
**WeightUnit** | Pointer to **string** | Weight Unit | [optional] 
**Height** | Pointer to **float32** | Defines product&#39;s height | [optional] 
**Length** | Pointer to **float32** | Defines product&#39;s length | [optional] 
**Width** | Pointer to **float32** | Defines product&#39;s width | [optional] 
**DimensionsUnit** | Pointer to **string** | Weight Unit | [optional] 
**IsVirtual** | Pointer to **bool** | Defines whether the product is virtual | [optional] [default to false]
**IsFreeShipping** | Pointer to **bool** | Specifies product free shipping flag that has to be updated | [optional] 
**Gtin** | Pointer to **string** | Global Trade Item Number. An GTIN is an identifier for trade items. | [optional] 
**Upc** | Pointer to **string** | Universal Product Code. A UPC (UPC-A) is a commonly used identifer for many different products. | [optional] 
**Mpn** | Pointer to **string** | Manufacturer Part Number. A MPN is an identifier of a particular part design or material used. | [optional] 
**Ean** | Pointer to **string** | European Article Number. An EAN is a unique 8 or 13-digit identifier that many industries (such as book publishers) use to identify products. | [optional] 
**Isbn** | Pointer to **string** | International Standard Book Number. An ISBN is a unique identifier for books. | [optional] 
**Barcode** | Pointer to **string** | A barcode is a unique code composed of numbers used as a product identifier. | [optional] 
**Manufacturer** | Pointer to **string** | Defines product&#39;s manufacturer | [optional] 
**ManufacturerId** | Pointer to **string** | Defines product&#39;s manufacturer by manufacturer_id | [optional] 
**CategoriesIds** | Pointer to **string** | Defines product add that is specified by comma-separated categories id | [optional] 
**RelatedProductsIds** | Pointer to **string** | Defines product related products ids that has to be updated | [optional] 
**UpSellProductsIds** | Pointer to **string** | Defines product up-sell products ids that has to be updated | [optional] 
**CrossSellProductsIds** | Pointer to **string** | Defines product cross-sells products ids that has to be updated | [optional] 
**MetaTitle** | Pointer to **string** | Defines unique meta title for each entity | [optional] 
**MetaKeywords** | Pointer to **string** | Defines unique meta keywords for each entity | [optional] 
**MetaDescription** | Pointer to **string** | Defines unique meta description of a entity | [optional] 
**SeoUrl** | Pointer to **string** | Defines unique URL for SEO | [optional] 
**SearchKeywords** | Pointer to **string** | Defines unique search keywords | [optional] 
**Tags** | Pointer to **string** | Product tags | [optional] 
**DeliveryCode** | Pointer to **string** | The delivery promise that applies to offer | [optional] 
**PackageDetails** | Pointer to [**ProductAddPackageDetails**](ProductAddPackageDetails.md) |  | [optional] 
**CountryOfOrigin** | Pointer to **string** | The country where the inventory item was made | [optional] 
**HarmonizedSystemCode** | Pointer to **string** | Harmonized System Code. An HSC is a 6-digit identifier that allows participating countries to classify traded goods on a common basis for customs purposes | [optional] 
**ShippingTemplateId** | Pointer to **int32** | The numeric ID of the shipping template associated with the products in Etsy. You can find possible values in the \&quot;cart.info\&quot; API method response, in the field shipping_zones[]-&gt;id. | [optional] [default to 0]
**ProcessingProfileId** | Pointer to **int32** | The numeric ID of the processing profile (readiness state) for physical products in Etsy. You can find possible values in the \&quot;cart.info\&quot; API method response, in the field processing_profiles[]-&gt;readiness_state_id. | [optional] 
**WhenMade** | Pointer to **string** | An enumerated string for the era in which the maker made the product. | [optional] [default to "made_to_order"]
**IsSupply** | Pointer to **bool** | If true, it indicates the product as a supply, otherwise it indicates that it is a finished product. | [optional] [default to true]
**Downloadable** | Pointer to **bool** | Defines whether the product is downloadable | [optional] [default to false]
**Materials** | Pointer to **[]string** | A list of material strings for materials used in the product. | [optional] 
**AutoRenew** | Pointer to **bool** | When true, automatically renews a listing upon its expiration. | [optional] [default to false]
**OnSale** | Pointer to **bool** | Set whether the product on sale | [optional] [default to false]
**ProductionPartnerIds** | Pointer to **string** | Defines product production partner ids that has to be updated | [optional] 
**ManufacturerInfo** | Pointer to [**ProductAddManufacturerInfo**](ProductAddManufacturerInfo.md) |  | [optional] 
**ReportRequestId** | Pointer to **string** | Report request id | [optional] 
**DisableReportCache** | Pointer to **bool** | Disable report cache for current request | [optional] [default to false]
**Reindex** | Pointer to **bool** | Is reindex required | [optional] [default to true]
**ClearCache** | Pointer to **bool** | Is cache clear required | [optional] [default to true]
**CheckProcessStatus** | Pointer to **bool** | Disable or enable check process status. Please note that the response will be slower due to additional requests to the store. | [optional] [default to false]
**Specifics** | Pointer to [**[]ProductAddSpecificsInner**](ProductAddSpecificsInner.md) | An array of Item Specific Name/Value pairs used by the seller to provide descriptive details of an item in a structured manner.         The list of possible specifications can be obtained using the category.info method (additional_fields-&gt;product_specifics).         &lt;b&gt;The structure of the parameter is different for specific platforms.&lt;/b&gt; | [optional] 
**ShopSectionId** | Pointer to **int32** | Add Shop Section Id | [optional] 
**PersonalizationDetails** | Pointer to [**ProductAddPersonalizationDetails**](ProductAddPersonalizationDetails.md) |  | [optional] 
**ExternalProductLink** | Pointer to **string** | External product link | [optional] 
**MarketplaceItemProperties** | Pointer to **string** | String containing the JSON representation of the supplied data | [optional] 
**MinOrderQuantity** | Pointer to **float32** | The minimum quantity an order must contain, to be eligible to purchase this product. | [optional] 

## Methods

### NewProductUpdate

`func NewProductUpdate() *ProductUpdate`

NewProductUpdate instantiates a new ProductUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductUpdateWithDefaults

`func NewProductUpdateWithDefaults() *ProductUpdate`

NewProductUpdateWithDefaults instantiates a new ProductUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *ProductUpdate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ProductUpdate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ProductUpdate) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ProductUpdate) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSku

`func (o *ProductUpdate) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductUpdate) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductUpdate) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductUpdate) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetName

`func (o *ProductUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProductUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortDescription

`func (o *ProductUpdate) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ProductUpdate) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ProductUpdate) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ProductUpdate) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetPrice

`func (o *ProductUpdate) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductUpdate) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductUpdate) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductUpdate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOldPrice

`func (o *ProductUpdate) GetOldPrice() float32`

GetOldPrice returns the OldPrice field if non-nil, zero value otherwise.

### GetOldPriceOk

`func (o *ProductUpdate) GetOldPriceOk() (*float32, bool)`

GetOldPriceOk returns a tuple with the OldPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPrice

`func (o *ProductUpdate) SetOldPrice(v float32)`

SetOldPrice sets OldPrice field to given value.

### HasOldPrice

`func (o *ProductUpdate) HasOldPrice() bool`

HasOldPrice returns a boolean if a field has been set.

### GetSpecialPrice

`func (o *ProductUpdate) GetSpecialPrice() float32`

GetSpecialPrice returns the SpecialPrice field if non-nil, zero value otherwise.

### GetSpecialPriceOk

`func (o *ProductUpdate) GetSpecialPriceOk() (*float32, bool)`

GetSpecialPriceOk returns a tuple with the SpecialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPrice

`func (o *ProductUpdate) SetSpecialPrice(v float32)`

SetSpecialPrice sets SpecialPrice field to given value.

### HasSpecialPrice

`func (o *ProductUpdate) HasSpecialPrice() bool`

HasSpecialPrice returns a boolean if a field has been set.

### GetSpriceCreate

`func (o *ProductUpdate) GetSpriceCreate() string`

GetSpriceCreate returns the SpriceCreate field if non-nil, zero value otherwise.

### GetSpriceCreateOk

`func (o *ProductUpdate) GetSpriceCreateOk() (*string, bool)`

GetSpriceCreateOk returns a tuple with the SpriceCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceCreate

`func (o *ProductUpdate) SetSpriceCreate(v string)`

SetSpriceCreate sets SpriceCreate field to given value.

### HasSpriceCreate

`func (o *ProductUpdate) HasSpriceCreate() bool`

HasSpriceCreate returns a boolean if a field has been set.

### GetSpriceExpire

`func (o *ProductUpdate) GetSpriceExpire() string`

GetSpriceExpire returns the SpriceExpire field if non-nil, zero value otherwise.

### GetSpriceExpireOk

`func (o *ProductUpdate) GetSpriceExpireOk() (*string, bool)`

GetSpriceExpireOk returns a tuple with the SpriceExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceExpire

`func (o *ProductUpdate) SetSpriceExpire(v string)`

SetSpriceExpire sets SpriceExpire field to given value.

### HasSpriceExpire

`func (o *ProductUpdate) HasSpriceExpire() bool`

HasSpriceExpire returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductUpdate) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductUpdate) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductUpdate) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductUpdate) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetFixedCostShippingPrice

`func (o *ProductUpdate) GetFixedCostShippingPrice() float32`

GetFixedCostShippingPrice returns the FixedCostShippingPrice field if non-nil, zero value otherwise.

### GetFixedCostShippingPriceOk

`func (o *ProductUpdate) GetFixedCostShippingPriceOk() (*float32, bool)`

GetFixedCostShippingPriceOk returns a tuple with the FixedCostShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCostShippingPrice

`func (o *ProductUpdate) SetFixedCostShippingPrice(v float32)`

SetFixedCostShippingPrice sets FixedCostShippingPrice field to given value.

### HasFixedCostShippingPrice

`func (o *ProductUpdate) HasFixedCostShippingPrice() bool`

HasFixedCostShippingPrice returns a boolean if a field has been set.

### GetRetailPrice

`func (o *ProductUpdate) GetRetailPrice() float32`

GetRetailPrice returns the RetailPrice field if non-nil, zero value otherwise.

### GetRetailPriceOk

`func (o *ProductUpdate) GetRetailPriceOk() (*float32, bool)`

GetRetailPriceOk returns a tuple with the RetailPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailPrice

`func (o *ProductUpdate) SetRetailPrice(v float32)`

SetRetailPrice sets RetailPrice field to given value.

### HasRetailPrice

`func (o *ProductUpdate) HasRetailPrice() bool`

HasRetailPrice returns a boolean if a field has been set.

### GetTierPrices

`func (o *ProductUpdate) GetTierPrices() []ProductAddTierPricesInner`

GetTierPrices returns the TierPrices field if non-nil, zero value otherwise.

### GetTierPricesOk

`func (o *ProductUpdate) GetTierPricesOk() (*[]ProductAddTierPricesInner, bool)`

GetTierPricesOk returns a tuple with the TierPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierPrices

`func (o *ProductUpdate) SetTierPrices(v []ProductAddTierPricesInner)`

SetTierPrices sets TierPrices field to given value.

### HasTierPrices

`func (o *ProductUpdate) HasTierPrices() bool`

HasTierPrices returns a boolean if a field has been set.

### GetReservePrice

`func (o *ProductUpdate) GetReservePrice() float32`

GetReservePrice returns the ReservePrice field if non-nil, zero value otherwise.

### GetReservePriceOk

`func (o *ProductUpdate) GetReservePriceOk() (*float32, bool)`

GetReservePriceOk returns a tuple with the ReservePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservePrice

`func (o *ProductUpdate) SetReservePrice(v float32)`

SetReservePrice sets ReservePrice field to given value.

### HasReservePrice

`func (o *ProductUpdate) HasReservePrice() bool`

HasReservePrice returns a boolean if a field has been set.

### GetBuyitnowPrice

`func (o *ProductUpdate) GetBuyitnowPrice() float32`

GetBuyitnowPrice returns the BuyitnowPrice field if non-nil, zero value otherwise.

### GetBuyitnowPriceOk

`func (o *ProductUpdate) GetBuyitnowPriceOk() (*float32, bool)`

GetBuyitnowPriceOk returns a tuple with the BuyitnowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyitnowPrice

`func (o *ProductUpdate) SetBuyitnowPrice(v float32)`

SetBuyitnowPrice sets BuyitnowPrice field to given value.

### HasBuyitnowPrice

`func (o *ProductUpdate) HasBuyitnowPrice() bool`

HasBuyitnowPrice returns a boolean if a field has been set.

### GetTaxable

`func (o *ProductUpdate) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ProductUpdate) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ProductUpdate) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ProductUpdate) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetTaxClassId

`func (o *ProductUpdate) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *ProductUpdate) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *ProductUpdate) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *ProductUpdate) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### GetType

`func (o *ProductUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *ProductUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCondition

`func (o *ProductUpdate) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ProductUpdate) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ProductUpdate) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ProductUpdate) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetVisible

`func (o *ProductUpdate) GetVisible() string`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ProductUpdate) GetVisibleOk() (*string, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ProductUpdate) SetVisible(v string)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ProductUpdate) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetInStock

`func (o *ProductUpdate) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *ProductUpdate) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *ProductUpdate) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *ProductUpdate) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### GetAvail

`func (o *ProductUpdate) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *ProductUpdate) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *ProductUpdate) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *ProductUpdate) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetAvailFrom

`func (o *ProductUpdate) GetAvailFrom() string`

GetAvailFrom returns the AvailFrom field if non-nil, zero value otherwise.

### GetAvailFromOk

`func (o *ProductUpdate) GetAvailFromOk() (*string, bool)`

GetAvailFromOk returns a tuple with the AvailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailFrom

`func (o *ProductUpdate) SetAvailFrom(v string)`

SetAvailFrom sets AvailFrom field to given value.

### HasAvailFrom

`func (o *ProductUpdate) HasAvailFrom() bool`

HasAvailFrom returns a boolean if a field has been set.

### GetProductClass

`func (o *ProductUpdate) GetProductClass() string`

GetProductClass returns the ProductClass field if non-nil, zero value otherwise.

### GetProductClassOk

`func (o *ProductUpdate) GetProductClassOk() (*string, bool)`

GetProductClassOk returns a tuple with the ProductClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductClass

`func (o *ProductUpdate) SetProductClass(v string)`

SetProductClass sets ProductClass field to given value.

### HasProductClass

`func (o *ProductUpdate) HasProductClass() bool`

HasProductClass returns a boolean if a field has been set.

### GetBrandName

`func (o *ProductUpdate) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *ProductUpdate) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *ProductUpdate) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *ProductUpdate) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.

### GetAvailableForView

`func (o *ProductUpdate) GetAvailableForView() bool`

GetAvailableForView returns the AvailableForView field if non-nil, zero value otherwise.

### GetAvailableForViewOk

`func (o *ProductUpdate) GetAvailableForViewOk() (*bool, bool)`

GetAvailableForViewOk returns a tuple with the AvailableForView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForView

`func (o *ProductUpdate) SetAvailableForView(v bool)`

SetAvailableForView sets AvailableForView field to given value.

### HasAvailableForView

`func (o *ProductUpdate) HasAvailableForView() bool`

HasAvailableForView returns a boolean if a field has been set.

### GetStoresIds

`func (o *ProductUpdate) GetStoresIds() string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *ProductUpdate) GetStoresIdsOk() (*string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *ProductUpdate) SetStoresIds(v string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *ProductUpdate) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetStoreId

`func (o *ProductUpdate) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductUpdate) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductUpdate) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductUpdate) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetLangId

`func (o *ProductUpdate) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *ProductUpdate) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *ProductUpdate) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *ProductUpdate) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductUpdate) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductUpdate) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductUpdate) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductUpdate) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReserveQuantity

`func (o *ProductUpdate) GetReserveQuantity() float32`

GetReserveQuantity returns the ReserveQuantity field if non-nil, zero value otherwise.

### GetReserveQuantityOk

`func (o *ProductUpdate) GetReserveQuantityOk() (*float32, bool)`

GetReserveQuantityOk returns a tuple with the ReserveQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveQuantity

`func (o *ProductUpdate) SetReserveQuantity(v float32)`

SetReserveQuantity sets ReserveQuantity field to given value.

### HasReserveQuantity

`func (o *ProductUpdate) HasReserveQuantity() bool`

HasReserveQuantity returns a boolean if a field has been set.

### GetManageStock

`func (o *ProductUpdate) GetManageStock() bool`

GetManageStock returns the ManageStock field if non-nil, zero value otherwise.

### GetManageStockOk

`func (o *ProductUpdate) GetManageStockOk() (*bool, bool)`

GetManageStockOk returns a tuple with the ManageStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStock

`func (o *ProductUpdate) SetManageStock(v bool)`

SetManageStock sets ManageStock field to given value.

### HasManageStock

`func (o *ProductUpdate) HasManageStock() bool`

HasManageStock returns a boolean if a field has been set.

### GetBackorderStatus

`func (o *ProductUpdate) GetBackorderStatus() string`

GetBackorderStatus returns the BackorderStatus field if non-nil, zero value otherwise.

### GetBackorderStatusOk

`func (o *ProductUpdate) GetBackorderStatusOk() (*string, bool)`

GetBackorderStatusOk returns a tuple with the BackorderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderStatus

`func (o *ProductUpdate) SetBackorderStatus(v string)`

SetBackorderStatus sets BackorderStatus field to given value.

### HasBackorderStatus

`func (o *ProductUpdate) HasBackorderStatus() bool`

HasBackorderStatus returns a boolean if a field has been set.

### GetIncreaseQuantity

`func (o *ProductUpdate) GetIncreaseQuantity() float32`

GetIncreaseQuantity returns the IncreaseQuantity field if non-nil, zero value otherwise.

### GetIncreaseQuantityOk

`func (o *ProductUpdate) GetIncreaseQuantityOk() (*float32, bool)`

GetIncreaseQuantityOk returns a tuple with the IncreaseQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseQuantity

`func (o *ProductUpdate) SetIncreaseQuantity(v float32)`

SetIncreaseQuantity sets IncreaseQuantity field to given value.

### HasIncreaseQuantity

`func (o *ProductUpdate) HasIncreaseQuantity() bool`

HasIncreaseQuantity returns a boolean if a field has been set.

### GetReduceQuantity

`func (o *ProductUpdate) GetReduceQuantity() float32`

GetReduceQuantity returns the ReduceQuantity field if non-nil, zero value otherwise.

### GetReduceQuantityOk

`func (o *ProductUpdate) GetReduceQuantityOk() (*float32, bool)`

GetReduceQuantityOk returns a tuple with the ReduceQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceQuantity

`func (o *ProductUpdate) SetReduceQuantity(v float32)`

SetReduceQuantity sets ReduceQuantity field to given value.

### HasReduceQuantity

`func (o *ProductUpdate) HasReduceQuantity() bool`

HasReduceQuantity returns a boolean if a field has been set.

### GetLowStockThreshold

`func (o *ProductUpdate) GetLowStockThreshold() float32`

GetLowStockThreshold returns the LowStockThreshold field if non-nil, zero value otherwise.

### GetLowStockThresholdOk

`func (o *ProductUpdate) GetLowStockThresholdOk() (*float32, bool)`

GetLowStockThresholdOk returns a tuple with the LowStockThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowStockThreshold

`func (o *ProductUpdate) SetLowStockThreshold(v float32)`

SetLowStockThreshold sets LowStockThreshold field to given value.

### HasLowStockThreshold

`func (o *ProductUpdate) HasLowStockThreshold() bool`

HasLowStockThreshold returns a boolean if a field has been set.

### GetWarehouseId

`func (o *ProductUpdate) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ProductUpdate) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ProductUpdate) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ProductUpdate) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetWeight

`func (o *ProductUpdate) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductUpdate) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductUpdate) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductUpdate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *ProductUpdate) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ProductUpdate) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ProductUpdate) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ProductUpdate) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetHeight

`func (o *ProductUpdate) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ProductUpdate) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ProductUpdate) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ProductUpdate) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetLength

`func (o *ProductUpdate) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ProductUpdate) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ProductUpdate) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ProductUpdate) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWidth

`func (o *ProductUpdate) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ProductUpdate) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ProductUpdate) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ProductUpdate) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetDimensionsUnit

`func (o *ProductUpdate) GetDimensionsUnit() string`

GetDimensionsUnit returns the DimensionsUnit field if non-nil, zero value otherwise.

### GetDimensionsUnitOk

`func (o *ProductUpdate) GetDimensionsUnitOk() (*string, bool)`

GetDimensionsUnitOk returns a tuple with the DimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsUnit

`func (o *ProductUpdate) SetDimensionsUnit(v string)`

SetDimensionsUnit sets DimensionsUnit field to given value.

### HasDimensionsUnit

`func (o *ProductUpdate) HasDimensionsUnit() bool`

HasDimensionsUnit returns a boolean if a field has been set.

### GetIsVirtual

`func (o *ProductUpdate) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *ProductUpdate) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *ProductUpdate) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *ProductUpdate) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetIsFreeShipping

`func (o *ProductUpdate) GetIsFreeShipping() bool`

GetIsFreeShipping returns the IsFreeShipping field if non-nil, zero value otherwise.

### GetIsFreeShippingOk

`func (o *ProductUpdate) GetIsFreeShippingOk() (*bool, bool)`

GetIsFreeShippingOk returns a tuple with the IsFreeShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeShipping

`func (o *ProductUpdate) SetIsFreeShipping(v bool)`

SetIsFreeShipping sets IsFreeShipping field to given value.

### HasIsFreeShipping

`func (o *ProductUpdate) HasIsFreeShipping() bool`

HasIsFreeShipping returns a boolean if a field has been set.

### GetGtin

`func (o *ProductUpdate) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *ProductUpdate) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *ProductUpdate) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *ProductUpdate) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### GetUpc

`func (o *ProductUpdate) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ProductUpdate) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ProductUpdate) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ProductUpdate) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetMpn

`func (o *ProductUpdate) GetMpn() string`

GetMpn returns the Mpn field if non-nil, zero value otherwise.

### GetMpnOk

`func (o *ProductUpdate) GetMpnOk() (*string, bool)`

GetMpnOk returns a tuple with the Mpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpn

`func (o *ProductUpdate) SetMpn(v string)`

SetMpn sets Mpn field to given value.

### HasMpn

`func (o *ProductUpdate) HasMpn() bool`

HasMpn returns a boolean if a field has been set.

### GetEan

`func (o *ProductUpdate) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *ProductUpdate) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *ProductUpdate) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *ProductUpdate) HasEan() bool`

HasEan returns a boolean if a field has been set.

### GetIsbn

`func (o *ProductUpdate) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *ProductUpdate) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *ProductUpdate) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *ProductUpdate) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### GetBarcode

`func (o *ProductUpdate) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductUpdate) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductUpdate) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ProductUpdate) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetManufacturer

`func (o *ProductUpdate) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ProductUpdate) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ProductUpdate) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ProductUpdate) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetManufacturerId

`func (o *ProductUpdate) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *ProductUpdate) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *ProductUpdate) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *ProductUpdate) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetCategoriesIds

`func (o *ProductUpdate) GetCategoriesIds() string`

GetCategoriesIds returns the CategoriesIds field if non-nil, zero value otherwise.

### GetCategoriesIdsOk

`func (o *ProductUpdate) GetCategoriesIdsOk() (*string, bool)`

GetCategoriesIdsOk returns a tuple with the CategoriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesIds

`func (o *ProductUpdate) SetCategoriesIds(v string)`

SetCategoriesIds sets CategoriesIds field to given value.

### HasCategoriesIds

`func (o *ProductUpdate) HasCategoriesIds() bool`

HasCategoriesIds returns a boolean if a field has been set.

### GetRelatedProductsIds

`func (o *ProductUpdate) GetRelatedProductsIds() string`

GetRelatedProductsIds returns the RelatedProductsIds field if non-nil, zero value otherwise.

### GetRelatedProductsIdsOk

`func (o *ProductUpdate) GetRelatedProductsIdsOk() (*string, bool)`

GetRelatedProductsIdsOk returns a tuple with the RelatedProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedProductsIds

`func (o *ProductUpdate) SetRelatedProductsIds(v string)`

SetRelatedProductsIds sets RelatedProductsIds field to given value.

### HasRelatedProductsIds

`func (o *ProductUpdate) HasRelatedProductsIds() bool`

HasRelatedProductsIds returns a boolean if a field has been set.

### GetUpSellProductsIds

`func (o *ProductUpdate) GetUpSellProductsIds() string`

GetUpSellProductsIds returns the UpSellProductsIds field if non-nil, zero value otherwise.

### GetUpSellProductsIdsOk

`func (o *ProductUpdate) GetUpSellProductsIdsOk() (*string, bool)`

GetUpSellProductsIdsOk returns a tuple with the UpSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSellProductsIds

`func (o *ProductUpdate) SetUpSellProductsIds(v string)`

SetUpSellProductsIds sets UpSellProductsIds field to given value.

### HasUpSellProductsIds

`func (o *ProductUpdate) HasUpSellProductsIds() bool`

HasUpSellProductsIds returns a boolean if a field has been set.

### GetCrossSellProductsIds

`func (o *ProductUpdate) GetCrossSellProductsIds() string`

GetCrossSellProductsIds returns the CrossSellProductsIds field if non-nil, zero value otherwise.

### GetCrossSellProductsIdsOk

`func (o *ProductUpdate) GetCrossSellProductsIdsOk() (*string, bool)`

GetCrossSellProductsIdsOk returns a tuple with the CrossSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSellProductsIds

`func (o *ProductUpdate) SetCrossSellProductsIds(v string)`

SetCrossSellProductsIds sets CrossSellProductsIds field to given value.

### HasCrossSellProductsIds

`func (o *ProductUpdate) HasCrossSellProductsIds() bool`

HasCrossSellProductsIds returns a boolean if a field has been set.

### GetMetaTitle

`func (o *ProductUpdate) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *ProductUpdate) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *ProductUpdate) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *ProductUpdate) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *ProductUpdate) GetMetaKeywords() string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *ProductUpdate) GetMetaKeywordsOk() (*string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *ProductUpdate) SetMetaKeywords(v string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *ProductUpdate) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetMetaDescription

`func (o *ProductUpdate) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *ProductUpdate) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *ProductUpdate) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *ProductUpdate) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetSeoUrl

`func (o *ProductUpdate) GetSeoUrl() string`

GetSeoUrl returns the SeoUrl field if non-nil, zero value otherwise.

### GetSeoUrlOk

`func (o *ProductUpdate) GetSeoUrlOk() (*string, bool)`

GetSeoUrlOk returns a tuple with the SeoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoUrl

`func (o *ProductUpdate) SetSeoUrl(v string)`

SetSeoUrl sets SeoUrl field to given value.

### HasSeoUrl

`func (o *ProductUpdate) HasSeoUrl() bool`

HasSeoUrl returns a boolean if a field has been set.

### GetSearchKeywords

`func (o *ProductUpdate) GetSearchKeywords() string`

GetSearchKeywords returns the SearchKeywords field if non-nil, zero value otherwise.

### GetSearchKeywordsOk

`func (o *ProductUpdate) GetSearchKeywordsOk() (*string, bool)`

GetSearchKeywordsOk returns a tuple with the SearchKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKeywords

`func (o *ProductUpdate) SetSearchKeywords(v string)`

SetSearchKeywords sets SearchKeywords field to given value.

### HasSearchKeywords

`func (o *ProductUpdate) HasSearchKeywords() bool`

HasSearchKeywords returns a boolean if a field has been set.

### GetTags

`func (o *ProductUpdate) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductUpdate) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductUpdate) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProductUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDeliveryCode

`func (o *ProductUpdate) GetDeliveryCode() string`

GetDeliveryCode returns the DeliveryCode field if non-nil, zero value otherwise.

### GetDeliveryCodeOk

`func (o *ProductUpdate) GetDeliveryCodeOk() (*string, bool)`

GetDeliveryCodeOk returns a tuple with the DeliveryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCode

`func (o *ProductUpdate) SetDeliveryCode(v string)`

SetDeliveryCode sets DeliveryCode field to given value.

### HasDeliveryCode

`func (o *ProductUpdate) HasDeliveryCode() bool`

HasDeliveryCode returns a boolean if a field has been set.

### GetPackageDetails

`func (o *ProductUpdate) GetPackageDetails() ProductAddPackageDetails`

GetPackageDetails returns the PackageDetails field if non-nil, zero value otherwise.

### GetPackageDetailsOk

`func (o *ProductUpdate) GetPackageDetailsOk() (*ProductAddPackageDetails, bool)`

GetPackageDetailsOk returns a tuple with the PackageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDetails

`func (o *ProductUpdate) SetPackageDetails(v ProductAddPackageDetails)`

SetPackageDetails sets PackageDetails field to given value.

### HasPackageDetails

`func (o *ProductUpdate) HasPackageDetails() bool`

HasPackageDetails returns a boolean if a field has been set.

### GetCountryOfOrigin

`func (o *ProductUpdate) GetCountryOfOrigin() string`

GetCountryOfOrigin returns the CountryOfOrigin field if non-nil, zero value otherwise.

### GetCountryOfOriginOk

`func (o *ProductUpdate) GetCountryOfOriginOk() (*string, bool)`

GetCountryOfOriginOk returns a tuple with the CountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfOrigin

`func (o *ProductUpdate) SetCountryOfOrigin(v string)`

SetCountryOfOrigin sets CountryOfOrigin field to given value.

### HasCountryOfOrigin

`func (o *ProductUpdate) HasCountryOfOrigin() bool`

HasCountryOfOrigin returns a boolean if a field has been set.

### GetHarmonizedSystemCode

`func (o *ProductUpdate) GetHarmonizedSystemCode() string`

GetHarmonizedSystemCode returns the HarmonizedSystemCode field if non-nil, zero value otherwise.

### GetHarmonizedSystemCodeOk

`func (o *ProductUpdate) GetHarmonizedSystemCodeOk() (*string, bool)`

GetHarmonizedSystemCodeOk returns a tuple with the HarmonizedSystemCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarmonizedSystemCode

`func (o *ProductUpdate) SetHarmonizedSystemCode(v string)`

SetHarmonizedSystemCode sets HarmonizedSystemCode field to given value.

### HasHarmonizedSystemCode

`func (o *ProductUpdate) HasHarmonizedSystemCode() bool`

HasHarmonizedSystemCode returns a boolean if a field has been set.

### GetShippingTemplateId

`func (o *ProductUpdate) GetShippingTemplateId() int32`

GetShippingTemplateId returns the ShippingTemplateId field if non-nil, zero value otherwise.

### GetShippingTemplateIdOk

`func (o *ProductUpdate) GetShippingTemplateIdOk() (*int32, bool)`

GetShippingTemplateIdOk returns a tuple with the ShippingTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTemplateId

`func (o *ProductUpdate) SetShippingTemplateId(v int32)`

SetShippingTemplateId sets ShippingTemplateId field to given value.

### HasShippingTemplateId

`func (o *ProductUpdate) HasShippingTemplateId() bool`

HasShippingTemplateId returns a boolean if a field has been set.

### GetProcessingProfileId

`func (o *ProductUpdate) GetProcessingProfileId() int32`

GetProcessingProfileId returns the ProcessingProfileId field if non-nil, zero value otherwise.

### GetProcessingProfileIdOk

`func (o *ProductUpdate) GetProcessingProfileIdOk() (*int32, bool)`

GetProcessingProfileIdOk returns a tuple with the ProcessingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingProfileId

`func (o *ProductUpdate) SetProcessingProfileId(v int32)`

SetProcessingProfileId sets ProcessingProfileId field to given value.

### HasProcessingProfileId

`func (o *ProductUpdate) HasProcessingProfileId() bool`

HasProcessingProfileId returns a boolean if a field has been set.

### GetWhenMade

`func (o *ProductUpdate) GetWhenMade() string`

GetWhenMade returns the WhenMade field if non-nil, zero value otherwise.

### GetWhenMadeOk

`func (o *ProductUpdate) GetWhenMadeOk() (*string, bool)`

GetWhenMadeOk returns a tuple with the WhenMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenMade

`func (o *ProductUpdate) SetWhenMade(v string)`

SetWhenMade sets WhenMade field to given value.

### HasWhenMade

`func (o *ProductUpdate) HasWhenMade() bool`

HasWhenMade returns a boolean if a field has been set.

### GetIsSupply

`func (o *ProductUpdate) GetIsSupply() bool`

GetIsSupply returns the IsSupply field if non-nil, zero value otherwise.

### GetIsSupplyOk

`func (o *ProductUpdate) GetIsSupplyOk() (*bool, bool)`

GetIsSupplyOk returns a tuple with the IsSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupply

`func (o *ProductUpdate) SetIsSupply(v bool)`

SetIsSupply sets IsSupply field to given value.

### HasIsSupply

`func (o *ProductUpdate) HasIsSupply() bool`

HasIsSupply returns a boolean if a field has been set.

### GetDownloadable

`func (o *ProductUpdate) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *ProductUpdate) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *ProductUpdate) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *ProductUpdate) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetMaterials

`func (o *ProductUpdate) GetMaterials() []string`

GetMaterials returns the Materials field if non-nil, zero value otherwise.

### GetMaterialsOk

`func (o *ProductUpdate) GetMaterialsOk() (*[]string, bool)`

GetMaterialsOk returns a tuple with the Materials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterials

`func (o *ProductUpdate) SetMaterials(v []string)`

SetMaterials sets Materials field to given value.

### HasMaterials

`func (o *ProductUpdate) HasMaterials() bool`

HasMaterials returns a boolean if a field has been set.

### GetAutoRenew

`func (o *ProductUpdate) GetAutoRenew() bool`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *ProductUpdate) GetAutoRenewOk() (*bool, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *ProductUpdate) SetAutoRenew(v bool)`

SetAutoRenew sets AutoRenew field to given value.

### HasAutoRenew

`func (o *ProductUpdate) HasAutoRenew() bool`

HasAutoRenew returns a boolean if a field has been set.

### GetOnSale

`func (o *ProductUpdate) GetOnSale() bool`

GetOnSale returns the OnSale field if non-nil, zero value otherwise.

### GetOnSaleOk

`func (o *ProductUpdate) GetOnSaleOk() (*bool, bool)`

GetOnSaleOk returns a tuple with the OnSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSale

`func (o *ProductUpdate) SetOnSale(v bool)`

SetOnSale sets OnSale field to given value.

### HasOnSale

`func (o *ProductUpdate) HasOnSale() bool`

HasOnSale returns a boolean if a field has been set.

### GetProductionPartnerIds

`func (o *ProductUpdate) GetProductionPartnerIds() string`

GetProductionPartnerIds returns the ProductionPartnerIds field if non-nil, zero value otherwise.

### GetProductionPartnerIdsOk

`func (o *ProductUpdate) GetProductionPartnerIdsOk() (*string, bool)`

GetProductionPartnerIdsOk returns a tuple with the ProductionPartnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionPartnerIds

`func (o *ProductUpdate) SetProductionPartnerIds(v string)`

SetProductionPartnerIds sets ProductionPartnerIds field to given value.

### HasProductionPartnerIds

`func (o *ProductUpdate) HasProductionPartnerIds() bool`

HasProductionPartnerIds returns a boolean if a field has been set.

### GetManufacturerInfo

`func (o *ProductUpdate) GetManufacturerInfo() ProductAddManufacturerInfo`

GetManufacturerInfo returns the ManufacturerInfo field if non-nil, zero value otherwise.

### GetManufacturerInfoOk

`func (o *ProductUpdate) GetManufacturerInfoOk() (*ProductAddManufacturerInfo, bool)`

GetManufacturerInfoOk returns a tuple with the ManufacturerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerInfo

`func (o *ProductUpdate) SetManufacturerInfo(v ProductAddManufacturerInfo)`

SetManufacturerInfo sets ManufacturerInfo field to given value.

### HasManufacturerInfo

`func (o *ProductUpdate) HasManufacturerInfo() bool`

HasManufacturerInfo returns a boolean if a field has been set.

### GetReportRequestId

`func (o *ProductUpdate) GetReportRequestId() string`

GetReportRequestId returns the ReportRequestId field if non-nil, zero value otherwise.

### GetReportRequestIdOk

`func (o *ProductUpdate) GetReportRequestIdOk() (*string, bool)`

GetReportRequestIdOk returns a tuple with the ReportRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportRequestId

`func (o *ProductUpdate) SetReportRequestId(v string)`

SetReportRequestId sets ReportRequestId field to given value.

### HasReportRequestId

`func (o *ProductUpdate) HasReportRequestId() bool`

HasReportRequestId returns a boolean if a field has been set.

### GetDisableReportCache

`func (o *ProductUpdate) GetDisableReportCache() bool`

GetDisableReportCache returns the DisableReportCache field if non-nil, zero value otherwise.

### GetDisableReportCacheOk

`func (o *ProductUpdate) GetDisableReportCacheOk() (*bool, bool)`

GetDisableReportCacheOk returns a tuple with the DisableReportCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReportCache

`func (o *ProductUpdate) SetDisableReportCache(v bool)`

SetDisableReportCache sets DisableReportCache field to given value.

### HasDisableReportCache

`func (o *ProductUpdate) HasDisableReportCache() bool`

HasDisableReportCache returns a boolean if a field has been set.

### GetReindex

`func (o *ProductUpdate) GetReindex() bool`

GetReindex returns the Reindex field if non-nil, zero value otherwise.

### GetReindexOk

`func (o *ProductUpdate) GetReindexOk() (*bool, bool)`

GetReindexOk returns a tuple with the Reindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReindex

`func (o *ProductUpdate) SetReindex(v bool)`

SetReindex sets Reindex field to given value.

### HasReindex

`func (o *ProductUpdate) HasReindex() bool`

HasReindex returns a boolean if a field has been set.

### GetClearCache

`func (o *ProductUpdate) GetClearCache() bool`

GetClearCache returns the ClearCache field if non-nil, zero value otherwise.

### GetClearCacheOk

`func (o *ProductUpdate) GetClearCacheOk() (*bool, bool)`

GetClearCacheOk returns a tuple with the ClearCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCache

`func (o *ProductUpdate) SetClearCache(v bool)`

SetClearCache sets ClearCache field to given value.

### HasClearCache

`func (o *ProductUpdate) HasClearCache() bool`

HasClearCache returns a boolean if a field has been set.

### GetCheckProcessStatus

`func (o *ProductUpdate) GetCheckProcessStatus() bool`

GetCheckProcessStatus returns the CheckProcessStatus field if non-nil, zero value otherwise.

### GetCheckProcessStatusOk

`func (o *ProductUpdate) GetCheckProcessStatusOk() (*bool, bool)`

GetCheckProcessStatusOk returns a tuple with the CheckProcessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckProcessStatus

`func (o *ProductUpdate) SetCheckProcessStatus(v bool)`

SetCheckProcessStatus sets CheckProcessStatus field to given value.

### HasCheckProcessStatus

`func (o *ProductUpdate) HasCheckProcessStatus() bool`

HasCheckProcessStatus returns a boolean if a field has been set.

### GetSpecifics

`func (o *ProductUpdate) GetSpecifics() []ProductAddSpecificsInner`

GetSpecifics returns the Specifics field if non-nil, zero value otherwise.

### GetSpecificsOk

`func (o *ProductUpdate) GetSpecificsOk() (*[]ProductAddSpecificsInner, bool)`

GetSpecificsOk returns a tuple with the Specifics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifics

`func (o *ProductUpdate) SetSpecifics(v []ProductAddSpecificsInner)`

SetSpecifics sets Specifics field to given value.

### HasSpecifics

`func (o *ProductUpdate) HasSpecifics() bool`

HasSpecifics returns a boolean if a field has been set.

### GetShopSectionId

`func (o *ProductUpdate) GetShopSectionId() int32`

GetShopSectionId returns the ShopSectionId field if non-nil, zero value otherwise.

### GetShopSectionIdOk

`func (o *ProductUpdate) GetShopSectionIdOk() (*int32, bool)`

GetShopSectionIdOk returns a tuple with the ShopSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopSectionId

`func (o *ProductUpdate) SetShopSectionId(v int32)`

SetShopSectionId sets ShopSectionId field to given value.

### HasShopSectionId

`func (o *ProductUpdate) HasShopSectionId() bool`

HasShopSectionId returns a boolean if a field has been set.

### GetPersonalizationDetails

`func (o *ProductUpdate) GetPersonalizationDetails() ProductAddPersonalizationDetails`

GetPersonalizationDetails returns the PersonalizationDetails field if non-nil, zero value otherwise.

### GetPersonalizationDetailsOk

`func (o *ProductUpdate) GetPersonalizationDetailsOk() (*ProductAddPersonalizationDetails, bool)`

GetPersonalizationDetailsOk returns a tuple with the PersonalizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalizationDetails

`func (o *ProductUpdate) SetPersonalizationDetails(v ProductAddPersonalizationDetails)`

SetPersonalizationDetails sets PersonalizationDetails field to given value.

### HasPersonalizationDetails

`func (o *ProductUpdate) HasPersonalizationDetails() bool`

HasPersonalizationDetails returns a boolean if a field has been set.

### GetExternalProductLink

`func (o *ProductUpdate) GetExternalProductLink() string`

GetExternalProductLink returns the ExternalProductLink field if non-nil, zero value otherwise.

### GetExternalProductLinkOk

`func (o *ProductUpdate) GetExternalProductLinkOk() (*string, bool)`

GetExternalProductLinkOk returns a tuple with the ExternalProductLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProductLink

`func (o *ProductUpdate) SetExternalProductLink(v string)`

SetExternalProductLink sets ExternalProductLink field to given value.

### HasExternalProductLink

`func (o *ProductUpdate) HasExternalProductLink() bool`

HasExternalProductLink returns a boolean if a field has been set.

### GetMarketplaceItemProperties

`func (o *ProductUpdate) GetMarketplaceItemProperties() string`

GetMarketplaceItemProperties returns the MarketplaceItemProperties field if non-nil, zero value otherwise.

### GetMarketplaceItemPropertiesOk

`func (o *ProductUpdate) GetMarketplaceItemPropertiesOk() (*string, bool)`

GetMarketplaceItemPropertiesOk returns a tuple with the MarketplaceItemProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceItemProperties

`func (o *ProductUpdate) SetMarketplaceItemProperties(v string)`

SetMarketplaceItemProperties sets MarketplaceItemProperties field to given value.

### HasMarketplaceItemProperties

`func (o *ProductUpdate) HasMarketplaceItemProperties() bool`

HasMarketplaceItemProperties returns a boolean if a field has been set.

### GetMinOrderQuantity

`func (o *ProductUpdate) GetMinOrderQuantity() float32`

GetMinOrderQuantity returns the MinOrderQuantity field if non-nil, zero value otherwise.

### GetMinOrderQuantityOk

`func (o *ProductUpdate) GetMinOrderQuantityOk() (*float32, bool)`

GetMinOrderQuantityOk returns a tuple with the MinOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrderQuantity

`func (o *ProductUpdate) SetMinOrderQuantity(v float32)`

SetMinOrderQuantity sets MinOrderQuantity field to given value.

### HasMinOrderQuantity

`func (o *ProductUpdate) HasMinOrderQuantity() bool`

HasMinOrderQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


