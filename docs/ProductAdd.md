# ProductAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Defines product&#39;s name that has to be added | 
**Model** | **string** | Defines product&#39;s model that has to be added | 
**Sku** | Pointer to **string** | Defines product&#39;s sku that has to be added | [optional] 
**Description** | **string** | Defines product&#39;s description that has to be added | 
**Price** | **float32** | Defines product&#39;s price that has to be added | 
**OldPrice** | Pointer to **float32** | Defines product&#39;s old price | [optional] 
**SpecialPrice** | Pointer to **float32** | Defines product&#39;s model that has to be added | [optional] 
**CostPrice** | Pointer to **float32** | Defines new product&#39;s cost price | [optional] 
**FixedCostShippingPrice** | Pointer to **float32** | Specifies product&#39;s fixed cost shipping price | [optional] 
**SpriceCreate** | Pointer to **string** | Defines the date of special price creation | [optional] 
**SpriceModified** | Pointer to **string** | Defines the date of special price modification | [optional] 
**SpriceExpire** | Pointer to **string** | Defines the term of special price offer duration | [optional] 
**TierPrices** | Pointer to [**[]ProductAddTierPricesInner**](ProductAddTierPricesInner.md) | Defines product&#39;s tier prices | [optional] 
**GroupPrices** | Pointer to [**[]ProductAddGroupPricesInner**](ProductAddGroupPricesInner.md) | Defines product&#39;s group prices | [optional] 
**AvailableForView** | Pointer to **bool** | Specifies the set of visible/invisible products for users | [optional] [default to true]
**AvailableForSale** | Pointer to **bool** | Specifies the set of visible/invisible products for sale | [optional] [default to true]
**Weight** | Pointer to **float32** | Weight | [optional] [default to 0]
**Width** | Pointer to **float32** | Defines product&#39;s width | [optional] 
**Height** | Pointer to **float32** | Defines product&#39;s height | [optional] 
**Length** | Pointer to **float32** | Defines product&#39;s length | [optional] 
**WeightUnit** | Pointer to **string** | Weight Unit | [optional] 
**DimensionsUnit** | Pointer to **string** | Weight Unit | [optional] 
**ShortDescription** | Pointer to **string** | Defines short description | [optional] 
**WarehouseId** | Pointer to **string** | This parameter is used for selecting a warehouse where you need to set/modify a product quantity. | [optional] 
**BackorderStatus** | Pointer to **string** | Set backorder status | [optional] 
**Quantity** | Pointer to **float32** | Defines product&#39;s quantity that has to be added | [optional] [default to 0]
**Downloadable** | Pointer to **bool** | Defines whether the product is downloadable | [optional] [default to false]
**WholesalePrice** | Pointer to **float32** | Defines product&#39;s sale price | [optional] 
**CreatedAt** | Pointer to **string** | Defines the date of entity creation | [optional] 
**Manufacturer** | Pointer to **string** | Defines product&#39;s manufacturer | [optional] 
**ManufacturerId** | Pointer to **string** | Defines product&#39;s manufacturer by manufacturer_id | [optional] 
**CategoriesIds** | Pointer to **string** | Defines product add that is specified by comma-separated categories id | [optional] 
**RelatedProductsIds** | Pointer to **string** | Defines product&#39;s related products ids that has to be added | [optional] 
**UpSellProductsIds** | Pointer to **string** | Defines product&#39;s up-sell products ids that has to be added | [optional] 
**CrossSellProductsIds** | Pointer to **string** | Defines product&#39;s cross-sell products ids that has to be added | [optional] 
**TaxClassId** | Pointer to **string** | Defines tax classes where entity has to be added | [optional] 
**Type** | Pointer to **string** | Defines product&#39;s type | [optional] [default to "simple"]
**MetaTitle** | Pointer to **string** | Defines unique meta title for each entity | [optional] 
**MetaKeywords** | Pointer to **string** | Defines unique meta keywords for each entity | [optional] 
**MetaDescription** | Pointer to **string** | Defines unique meta description of a entity | [optional] 
**Url** | Pointer to **string** | Defines unique product&#39;s URL | [optional] 
**LangId** | Pointer to **string** | Language id | [optional] 
**StoresIds** | Pointer to **string** | Assign product to the stores that is specified by comma-separated stores&#39; id | [optional] 
**CategoryId** | Pointer to **string** | Defines product add that is specified by category id | [optional] 
**ViewedCount** | Pointer to **int32** | Specifies the number of product&#39;s reviews | [optional] [default to 0]
**OrderedCount** | Pointer to **int32** | Defines how many times the product was ordered | [optional] [default to 0]
**AttributeSetName** | Pointer to **string** | Defines product’s attribute set name in Magento | [optional] [default to "Default"]
**AttributeName** | Pointer to **string** | Defines product’s attribute name separated with a comma in Magento | [optional] 
**ShippingTemplateId** | Pointer to **int32** | The numeric ID of the shipping template associated with the products in Etsy. You can find possible values in the \&quot;cart.info\&quot; API method response, in the field shipping_zones[]-&gt;id. | [optional] [default to 0]
**ProductionPartnerIds** | Pointer to **string** | Defines product&#39;s production partner ids that has to be added | [optional] 
**Condition** | Pointer to **string** | The human-readable label for the condition (e.g., \&quot;New\&quot;). | [optional] 
**ListingDuration** | Pointer to **string** | Describes the number of days the seller wants the listing to be active. Look at cart.info method response for allowed values. | [optional] 
**ListingType** | Pointer to **string** | Indicates the selling format of the marketplace listing. | [optional] [default to "FixedPrice"]
**PaymentMethods** | Pointer to **[]string** | Identifies the payment method (such as PayPal) that the seller will accept when the buyer pays for the item. Look at cart.info method response for allowed values.&lt;hr&gt;&lt;div style&#x3D;\&quot;font-style:normal\&quot;&gt;Param structure:&lt;div style&#x3D;\&quot;margin-left: 2%;\&quot;&gt;&lt;code style&#x3D;\&quot;padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;\&quot;&gt;payment_methods[0] &#x3D; string&lt;/br&gt;payment_methods[1] &#x3D; string&lt;/br&gt;&lt;/code&gt;&lt;/div&gt;&lt;/div&gt; | [optional] 
**ReturnAccepted** | Pointer to **bool** | Indicates whether the seller allows the buyer to return the item. | [optional] 
**ShippingDetails** | Pointer to [**[]ProductAddShippingDetailsInner**](ProductAddShippingDetailsInner.md) | The shipping details, including flat and calculated shipping costs and shipping insurance costs. Look at cart.info method response for allowed values.&lt;hr&gt;&lt;div style&#x3D;\&quot;font-style:normal\&quot;&gt;Param structure:&lt;div style&#x3D;\&quot;margin-left: 2%;\&quot;&gt;&lt;code style&#x3D;\&quot;padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;\&quot;&gt;shipping_details[0][&lt;b&gt;shipping_type&lt;/b&gt;] &#x3D; string &lt;/br&gt;shipping_details[0][&lt;b&gt;shipping_service&lt;/b&gt;] &#x3D; string&lt;/br&gt;shipping_details[0][&lt;b&gt;shipping_cost&lt;/b&gt;] &#x3D; decimal&lt;/br&gt;shipping_details[1][&lt;b&gt;shipping_type&lt;/b&gt;] &#x3D; string &lt;/br&gt;shipping_details[1][&lt;b&gt;shipping_service&lt;/b&gt;] &#x3D; string&lt;/br&gt;shipping_details[1][&lt;b&gt;shipping_cost&lt;/b&gt;] &#x3D; decimal&lt;/br&gt;&lt;/code&gt;&lt;/div&gt;&lt;/div&gt; | [optional] 
**PaypalEmail** | Pointer to **string** | Valid PayPal email address for the PayPal account that the seller will use if they offer PayPal as a payment method for the listing. | [optional] 
**SellerProfiles** | Pointer to [**ProductAddSellerProfiles**](ProductAddSellerProfiles.md) |  | [optional] 
**PackageDetails** | Pointer to [**ProductAddPackageDetails**](ProductAddPackageDetails.md) |  | [optional] 
**BestOffer** | Pointer to [**ProductAddBestOffer**](ProductAddBestOffer.md) |  | [optional] 
**SalesTax** | Pointer to [**ProductAddSalesTax**](ProductAddSalesTax.md) |  | [optional] 
**Barcode** | Pointer to **string** | A barcode is a unique code composed of numbers used as a product identifier. | [optional] 
**Upc** | Pointer to **string** | Universal Product Code. A UPC (UPC-A) is a commonly used identifer for many different products. | [optional] 
**Ean** | Pointer to **string** | European Article Number. An EAN is a unique 8 or 13-digit identifier that many industries (such as book publishers) use to identify products. | [optional] 
**Isbn** | Pointer to **string** | International Standard Book Number. An ISBN is a unique identifier for books. | [optional] 
**Specifics** | Pointer to [**[]ProductAddSpecificsInner**](ProductAddSpecificsInner.md) | An array of Item Specific Name/Value pairs used by the seller to provide descriptive details of an item in a structured manner.         The list of possible specifications can be obtained using the category.info method (additional_fields-&gt;product_specifics).         &lt;b&gt;The structure of the parameter is different for specific platforms.&lt;/b&gt; | [optional] 
**ImageUrl** | Pointer to **string** | Image Url | [optional] 
**ImageName** | Pointer to **string** | Defines image&#39;s name | [optional] 
**ReservePrice** | Pointer to **float32** | Defines reserve price value | [optional] 
**BuyitnowPrice** | Pointer to **float32** | Defines buy it now value | [optional] 
**ConditionDescription** | Pointer to **string** | Detailed description of the product condition. | [optional] 
**AuctionConfidentialityLevel** | Pointer to **string** | This allows buyers to remain anonymous when the bid or buy an item. | [optional] 
**AvailFrom** | Pointer to **string** | Allows to schedule a time in the future that the item becomes available. The value should be greater than the current date and time. | [optional] 
**Tags** | Pointer to **string** | Product tags | [optional] 
**ClearCache** | Pointer to **bool** | Is cache clear required | [optional] [default to true]
**Asin** | Pointer to **string** | Amazon Standard Identification Number. | [optional] 
**Gtin** | Pointer to **string** | Global Trade Item Number. An GTIN is an identifier for trade items. | [optional] 
**Mpn** | Pointer to **string** | Manufacturer Part Number. A MPN is an identifier of a particular part design or material used. | [optional] 
**Taxable** | Pointer to **bool** | Specifies whether a tax is charged | [optional] [default to true]
**Visible** | Pointer to **string** | Set visibility status | [optional] 
**Status** | Pointer to **string** | Defines product&#39;s status | [optional] 
**SeoUrl** | Pointer to **string** | Defines unique URL for SEO | [optional] 
**ProductClass** | Pointer to **string** | A categorization for the product | [optional] 
**ProductType** | Pointer to **string** | A categorization for the product | [optional] 
**MarketplaceItemProperties** | Pointer to **string** | String containing the JSON representation of the supplied data | [optional] 
**ManageStock** | Pointer to **bool** | Defines inventory tracking for product | [optional] 
**HarmonizedSystemCode** | Pointer to **string** | Harmonized System Code. An HSC is a 6-digit identifier that allows participating countries to classify traded goods on a common basis for customs purposes | [optional] 
**CountryOfOrigin** | Pointer to **string** | The country where the inventory item was made | [optional] 
**Files** | Pointer to [**[]ProductAddFilesInner**](ProductAddFilesInner.md) | File Url | [optional] 
**SearchKeywords** | Pointer to **string** | Defines unique search keywords | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**BrandName** | Pointer to **string** | Defines product brand name | [optional] 
**IsVirtual** | Pointer to **bool** | Defines whether the product is virtual | [optional] [default to false]
**IsFreeShipping** | Pointer to **bool** | Specifies product&#39;s free shipping flag that has to be added | [optional] 
**InStock** | Pointer to **bool** | Set stock status | [optional] 
**DeliveryCode** | Pointer to **string** | The delivery promise that applies to offer | [optional] 
**ProductReference** | Pointer to **string** | Groups all variations, that you want to combine into one product. | [optional] 
**DeliveryType** | Pointer to **string** | Defines the type of the delivery. | [optional] 
**DeliveryTime** | Pointer to **int32** | Defines delivery time in days. | [optional] 
**SizeChart** | Pointer to [**ProductAddSizeChart**](ProductAddSizeChart.md) |  | [optional] 
**Certifications** | Pointer to [**[]ProductAddCertificationsInner**](ProductAddCertificationsInner.md) | An array of product certifications. The list of possible certifications can be obtained using the \&quot;&lt;i&gt;category.info&lt;/i&gt;\&quot; method (&lt;i&gt;additional_fields-&gt;rules-&gt;product_certifications&lt;/i&gt;). | [optional] 
**DeliveryOptionIds** | Pointer to **string** | Defines delivery options for product by ids. | [optional] 
**ManufacturerInfo** | Pointer to [**ProductAddManufacturerInfo**](ProductAddManufacturerInfo.md) |  | [optional] 
**WhenMade** | Pointer to **string** | An enumerated string for the era in which the maker made the product. | [optional] [default to "made_to_order"]
**IsSupply** | Pointer to **bool** | If true, it indicates the product as a supply, otherwise it indicates that it is a finished product. | [optional] [default to true]
**Materials** | Pointer to **[]string** | A list of material strings for materials used in the product. | [optional] 
**AutoRenew** | Pointer to **bool** | When true, automatically renews a listing upon its expiration. | [optional] [default to false]
**AllowDisplayCondition** | Pointer to **bool** | Flag used to determine whether the product condition is shown to the customer on the product page. | [optional] 
**MinOrderQuantity** | Pointer to **float32** | The minimum quantity an order must contain, to be eligible to purchase this product. | [optional] 
**MaxOrderQuantity** | Pointer to **float32** | The maximum quantity an order can contain when purchasing the product. | [optional] 

## Methods

### NewProductAdd

`func NewProductAdd(name string, model string, description string, price float32, ) *ProductAdd`

NewProductAdd instantiates a new ProductAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAddWithDefaults

`func NewProductAddWithDefaults() *ProductAdd`

NewProductAddWithDefaults instantiates a new ProductAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductAdd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductAdd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductAdd) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *ProductAdd) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ProductAdd) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ProductAdd) SetModel(v string)`

SetModel sets Model field to given value.


### GetSku

`func (o *ProductAdd) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductAdd) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductAdd) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductAdd) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetDescription

`func (o *ProductAdd) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductAdd) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductAdd) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPrice

`func (o *ProductAdd) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductAdd) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductAdd) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetOldPrice

`func (o *ProductAdd) GetOldPrice() float32`

GetOldPrice returns the OldPrice field if non-nil, zero value otherwise.

### GetOldPriceOk

`func (o *ProductAdd) GetOldPriceOk() (*float32, bool)`

GetOldPriceOk returns a tuple with the OldPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPrice

`func (o *ProductAdd) SetOldPrice(v float32)`

SetOldPrice sets OldPrice field to given value.

### HasOldPrice

`func (o *ProductAdd) HasOldPrice() bool`

HasOldPrice returns a boolean if a field has been set.

### GetSpecialPrice

`func (o *ProductAdd) GetSpecialPrice() float32`

GetSpecialPrice returns the SpecialPrice field if non-nil, zero value otherwise.

### GetSpecialPriceOk

`func (o *ProductAdd) GetSpecialPriceOk() (*float32, bool)`

GetSpecialPriceOk returns a tuple with the SpecialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPrice

`func (o *ProductAdd) SetSpecialPrice(v float32)`

SetSpecialPrice sets SpecialPrice field to given value.

### HasSpecialPrice

`func (o *ProductAdd) HasSpecialPrice() bool`

HasSpecialPrice returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductAdd) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductAdd) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductAdd) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductAdd) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetFixedCostShippingPrice

`func (o *ProductAdd) GetFixedCostShippingPrice() float32`

GetFixedCostShippingPrice returns the FixedCostShippingPrice field if non-nil, zero value otherwise.

### GetFixedCostShippingPriceOk

`func (o *ProductAdd) GetFixedCostShippingPriceOk() (*float32, bool)`

GetFixedCostShippingPriceOk returns a tuple with the FixedCostShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCostShippingPrice

`func (o *ProductAdd) SetFixedCostShippingPrice(v float32)`

SetFixedCostShippingPrice sets FixedCostShippingPrice field to given value.

### HasFixedCostShippingPrice

`func (o *ProductAdd) HasFixedCostShippingPrice() bool`

HasFixedCostShippingPrice returns a boolean if a field has been set.

### GetSpriceCreate

`func (o *ProductAdd) GetSpriceCreate() string`

GetSpriceCreate returns the SpriceCreate field if non-nil, zero value otherwise.

### GetSpriceCreateOk

`func (o *ProductAdd) GetSpriceCreateOk() (*string, bool)`

GetSpriceCreateOk returns a tuple with the SpriceCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceCreate

`func (o *ProductAdd) SetSpriceCreate(v string)`

SetSpriceCreate sets SpriceCreate field to given value.

### HasSpriceCreate

`func (o *ProductAdd) HasSpriceCreate() bool`

HasSpriceCreate returns a boolean if a field has been set.

### GetSpriceModified

`func (o *ProductAdd) GetSpriceModified() string`

GetSpriceModified returns the SpriceModified field if non-nil, zero value otherwise.

### GetSpriceModifiedOk

`func (o *ProductAdd) GetSpriceModifiedOk() (*string, bool)`

GetSpriceModifiedOk returns a tuple with the SpriceModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceModified

`func (o *ProductAdd) SetSpriceModified(v string)`

SetSpriceModified sets SpriceModified field to given value.

### HasSpriceModified

`func (o *ProductAdd) HasSpriceModified() bool`

HasSpriceModified returns a boolean if a field has been set.

### GetSpriceExpire

`func (o *ProductAdd) GetSpriceExpire() string`

GetSpriceExpire returns the SpriceExpire field if non-nil, zero value otherwise.

### GetSpriceExpireOk

`func (o *ProductAdd) GetSpriceExpireOk() (*string, bool)`

GetSpriceExpireOk returns a tuple with the SpriceExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceExpire

`func (o *ProductAdd) SetSpriceExpire(v string)`

SetSpriceExpire sets SpriceExpire field to given value.

### HasSpriceExpire

`func (o *ProductAdd) HasSpriceExpire() bool`

HasSpriceExpire returns a boolean if a field has been set.

### GetTierPrices

`func (o *ProductAdd) GetTierPrices() []ProductAddTierPricesInner`

GetTierPrices returns the TierPrices field if non-nil, zero value otherwise.

### GetTierPricesOk

`func (o *ProductAdd) GetTierPricesOk() (*[]ProductAddTierPricesInner, bool)`

GetTierPricesOk returns a tuple with the TierPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierPrices

`func (o *ProductAdd) SetTierPrices(v []ProductAddTierPricesInner)`

SetTierPrices sets TierPrices field to given value.

### HasTierPrices

`func (o *ProductAdd) HasTierPrices() bool`

HasTierPrices returns a boolean if a field has been set.

### GetGroupPrices

`func (o *ProductAdd) GetGroupPrices() []ProductAddGroupPricesInner`

GetGroupPrices returns the GroupPrices field if non-nil, zero value otherwise.

### GetGroupPricesOk

`func (o *ProductAdd) GetGroupPricesOk() (*[]ProductAddGroupPricesInner, bool)`

GetGroupPricesOk returns a tuple with the GroupPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPrices

`func (o *ProductAdd) SetGroupPrices(v []ProductAddGroupPricesInner)`

SetGroupPrices sets GroupPrices field to given value.

### HasGroupPrices

`func (o *ProductAdd) HasGroupPrices() bool`

HasGroupPrices returns a boolean if a field has been set.

### GetAvailableForView

`func (o *ProductAdd) GetAvailableForView() bool`

GetAvailableForView returns the AvailableForView field if non-nil, zero value otherwise.

### GetAvailableForViewOk

`func (o *ProductAdd) GetAvailableForViewOk() (*bool, bool)`

GetAvailableForViewOk returns a tuple with the AvailableForView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForView

`func (o *ProductAdd) SetAvailableForView(v bool)`

SetAvailableForView sets AvailableForView field to given value.

### HasAvailableForView

`func (o *ProductAdd) HasAvailableForView() bool`

HasAvailableForView returns a boolean if a field has been set.

### GetAvailableForSale

`func (o *ProductAdd) GetAvailableForSale() bool`

GetAvailableForSale returns the AvailableForSale field if non-nil, zero value otherwise.

### GetAvailableForSaleOk

`func (o *ProductAdd) GetAvailableForSaleOk() (*bool, bool)`

GetAvailableForSaleOk returns a tuple with the AvailableForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForSale

`func (o *ProductAdd) SetAvailableForSale(v bool)`

SetAvailableForSale sets AvailableForSale field to given value.

### HasAvailableForSale

`func (o *ProductAdd) HasAvailableForSale() bool`

HasAvailableForSale returns a boolean if a field has been set.

### GetWeight

`func (o *ProductAdd) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductAdd) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductAdd) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductAdd) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWidth

`func (o *ProductAdd) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ProductAdd) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ProductAdd) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ProductAdd) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ProductAdd) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ProductAdd) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ProductAdd) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ProductAdd) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetLength

`func (o *ProductAdd) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ProductAdd) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ProductAdd) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ProductAdd) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWeightUnit

`func (o *ProductAdd) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ProductAdd) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ProductAdd) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ProductAdd) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetDimensionsUnit

`func (o *ProductAdd) GetDimensionsUnit() string`

GetDimensionsUnit returns the DimensionsUnit field if non-nil, zero value otherwise.

### GetDimensionsUnitOk

`func (o *ProductAdd) GetDimensionsUnitOk() (*string, bool)`

GetDimensionsUnitOk returns a tuple with the DimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsUnit

`func (o *ProductAdd) SetDimensionsUnit(v string)`

SetDimensionsUnit sets DimensionsUnit field to given value.

### HasDimensionsUnit

`func (o *ProductAdd) HasDimensionsUnit() bool`

HasDimensionsUnit returns a boolean if a field has been set.

### GetShortDescription

`func (o *ProductAdd) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ProductAdd) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ProductAdd) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ProductAdd) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetWarehouseId

`func (o *ProductAdd) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ProductAdd) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ProductAdd) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ProductAdd) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetBackorderStatus

`func (o *ProductAdd) GetBackorderStatus() string`

GetBackorderStatus returns the BackorderStatus field if non-nil, zero value otherwise.

### GetBackorderStatusOk

`func (o *ProductAdd) GetBackorderStatusOk() (*string, bool)`

GetBackorderStatusOk returns a tuple with the BackorderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderStatus

`func (o *ProductAdd) SetBackorderStatus(v string)`

SetBackorderStatus sets BackorderStatus field to given value.

### HasBackorderStatus

`func (o *ProductAdd) HasBackorderStatus() bool`

HasBackorderStatus returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductAdd) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductAdd) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductAdd) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductAdd) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDownloadable

`func (o *ProductAdd) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *ProductAdd) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *ProductAdd) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *ProductAdd) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetWholesalePrice

`func (o *ProductAdd) GetWholesalePrice() float32`

GetWholesalePrice returns the WholesalePrice field if non-nil, zero value otherwise.

### GetWholesalePriceOk

`func (o *ProductAdd) GetWholesalePriceOk() (*float32, bool)`

GetWholesalePriceOk returns a tuple with the WholesalePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholesalePrice

`func (o *ProductAdd) SetWholesalePrice(v float32)`

SetWholesalePrice sets WholesalePrice field to given value.

### HasWholesalePrice

`func (o *ProductAdd) HasWholesalePrice() bool`

HasWholesalePrice returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductAdd) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductAdd) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductAdd) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProductAdd) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetManufacturer

`func (o *ProductAdd) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ProductAdd) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ProductAdd) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ProductAdd) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetManufacturerId

`func (o *ProductAdd) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *ProductAdd) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *ProductAdd) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *ProductAdd) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetCategoriesIds

`func (o *ProductAdd) GetCategoriesIds() string`

GetCategoriesIds returns the CategoriesIds field if non-nil, zero value otherwise.

### GetCategoriesIdsOk

`func (o *ProductAdd) GetCategoriesIdsOk() (*string, bool)`

GetCategoriesIdsOk returns a tuple with the CategoriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesIds

`func (o *ProductAdd) SetCategoriesIds(v string)`

SetCategoriesIds sets CategoriesIds field to given value.

### HasCategoriesIds

`func (o *ProductAdd) HasCategoriesIds() bool`

HasCategoriesIds returns a boolean if a field has been set.

### GetRelatedProductsIds

`func (o *ProductAdd) GetRelatedProductsIds() string`

GetRelatedProductsIds returns the RelatedProductsIds field if non-nil, zero value otherwise.

### GetRelatedProductsIdsOk

`func (o *ProductAdd) GetRelatedProductsIdsOk() (*string, bool)`

GetRelatedProductsIdsOk returns a tuple with the RelatedProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedProductsIds

`func (o *ProductAdd) SetRelatedProductsIds(v string)`

SetRelatedProductsIds sets RelatedProductsIds field to given value.

### HasRelatedProductsIds

`func (o *ProductAdd) HasRelatedProductsIds() bool`

HasRelatedProductsIds returns a boolean if a field has been set.

### GetUpSellProductsIds

`func (o *ProductAdd) GetUpSellProductsIds() string`

GetUpSellProductsIds returns the UpSellProductsIds field if non-nil, zero value otherwise.

### GetUpSellProductsIdsOk

`func (o *ProductAdd) GetUpSellProductsIdsOk() (*string, bool)`

GetUpSellProductsIdsOk returns a tuple with the UpSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSellProductsIds

`func (o *ProductAdd) SetUpSellProductsIds(v string)`

SetUpSellProductsIds sets UpSellProductsIds field to given value.

### HasUpSellProductsIds

`func (o *ProductAdd) HasUpSellProductsIds() bool`

HasUpSellProductsIds returns a boolean if a field has been set.

### GetCrossSellProductsIds

`func (o *ProductAdd) GetCrossSellProductsIds() string`

GetCrossSellProductsIds returns the CrossSellProductsIds field if non-nil, zero value otherwise.

### GetCrossSellProductsIdsOk

`func (o *ProductAdd) GetCrossSellProductsIdsOk() (*string, bool)`

GetCrossSellProductsIdsOk returns a tuple with the CrossSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSellProductsIds

`func (o *ProductAdd) SetCrossSellProductsIds(v string)`

SetCrossSellProductsIds sets CrossSellProductsIds field to given value.

### HasCrossSellProductsIds

`func (o *ProductAdd) HasCrossSellProductsIds() bool`

HasCrossSellProductsIds returns a boolean if a field has been set.

### GetTaxClassId

`func (o *ProductAdd) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *ProductAdd) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *ProductAdd) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *ProductAdd) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### GetType

`func (o *ProductAdd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductAdd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductAdd) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductAdd) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetaTitle

`func (o *ProductAdd) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *ProductAdd) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *ProductAdd) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *ProductAdd) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *ProductAdd) GetMetaKeywords() string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *ProductAdd) GetMetaKeywordsOk() (*string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *ProductAdd) SetMetaKeywords(v string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *ProductAdd) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetMetaDescription

`func (o *ProductAdd) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *ProductAdd) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *ProductAdd) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *ProductAdd) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetUrl

`func (o *ProductAdd) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProductAdd) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProductAdd) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProductAdd) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLangId

`func (o *ProductAdd) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *ProductAdd) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *ProductAdd) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *ProductAdd) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetStoresIds

`func (o *ProductAdd) GetStoresIds() string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *ProductAdd) GetStoresIdsOk() (*string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *ProductAdd) SetStoresIds(v string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *ProductAdd) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetCategoryId

`func (o *ProductAdd) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ProductAdd) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ProductAdd) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ProductAdd) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetViewedCount

`func (o *ProductAdd) GetViewedCount() int32`

GetViewedCount returns the ViewedCount field if non-nil, zero value otherwise.

### GetViewedCountOk

`func (o *ProductAdd) GetViewedCountOk() (*int32, bool)`

GetViewedCountOk returns a tuple with the ViewedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedCount

`func (o *ProductAdd) SetViewedCount(v int32)`

SetViewedCount sets ViewedCount field to given value.

### HasViewedCount

`func (o *ProductAdd) HasViewedCount() bool`

HasViewedCount returns a boolean if a field has been set.

### GetOrderedCount

`func (o *ProductAdd) GetOrderedCount() int32`

GetOrderedCount returns the OrderedCount field if non-nil, zero value otherwise.

### GetOrderedCountOk

`func (o *ProductAdd) GetOrderedCountOk() (*int32, bool)`

GetOrderedCountOk returns a tuple with the OrderedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedCount

`func (o *ProductAdd) SetOrderedCount(v int32)`

SetOrderedCount sets OrderedCount field to given value.

### HasOrderedCount

`func (o *ProductAdd) HasOrderedCount() bool`

HasOrderedCount returns a boolean if a field has been set.

### GetAttributeSetName

`func (o *ProductAdd) GetAttributeSetName() string`

GetAttributeSetName returns the AttributeSetName field if non-nil, zero value otherwise.

### GetAttributeSetNameOk

`func (o *ProductAdd) GetAttributeSetNameOk() (*string, bool)`

GetAttributeSetNameOk returns a tuple with the AttributeSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSetName

`func (o *ProductAdd) SetAttributeSetName(v string)`

SetAttributeSetName sets AttributeSetName field to given value.

### HasAttributeSetName

`func (o *ProductAdd) HasAttributeSetName() bool`

HasAttributeSetName returns a boolean if a field has been set.

### GetAttributeName

`func (o *ProductAdd) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *ProductAdd) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *ProductAdd) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *ProductAdd) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetShippingTemplateId

`func (o *ProductAdd) GetShippingTemplateId() int32`

GetShippingTemplateId returns the ShippingTemplateId field if non-nil, zero value otherwise.

### GetShippingTemplateIdOk

`func (o *ProductAdd) GetShippingTemplateIdOk() (*int32, bool)`

GetShippingTemplateIdOk returns a tuple with the ShippingTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTemplateId

`func (o *ProductAdd) SetShippingTemplateId(v int32)`

SetShippingTemplateId sets ShippingTemplateId field to given value.

### HasShippingTemplateId

`func (o *ProductAdd) HasShippingTemplateId() bool`

HasShippingTemplateId returns a boolean if a field has been set.

### GetProductionPartnerIds

`func (o *ProductAdd) GetProductionPartnerIds() string`

GetProductionPartnerIds returns the ProductionPartnerIds field if non-nil, zero value otherwise.

### GetProductionPartnerIdsOk

`func (o *ProductAdd) GetProductionPartnerIdsOk() (*string, bool)`

GetProductionPartnerIdsOk returns a tuple with the ProductionPartnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionPartnerIds

`func (o *ProductAdd) SetProductionPartnerIds(v string)`

SetProductionPartnerIds sets ProductionPartnerIds field to given value.

### HasProductionPartnerIds

`func (o *ProductAdd) HasProductionPartnerIds() bool`

HasProductionPartnerIds returns a boolean if a field has been set.

### GetCondition

`func (o *ProductAdd) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ProductAdd) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ProductAdd) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ProductAdd) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetListingDuration

`func (o *ProductAdd) GetListingDuration() string`

GetListingDuration returns the ListingDuration field if non-nil, zero value otherwise.

### GetListingDurationOk

`func (o *ProductAdd) GetListingDurationOk() (*string, bool)`

GetListingDurationOk returns a tuple with the ListingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingDuration

`func (o *ProductAdd) SetListingDuration(v string)`

SetListingDuration sets ListingDuration field to given value.

### HasListingDuration

`func (o *ProductAdd) HasListingDuration() bool`

HasListingDuration returns a boolean if a field has been set.

### GetListingType

`func (o *ProductAdd) GetListingType() string`

GetListingType returns the ListingType field if non-nil, zero value otherwise.

### GetListingTypeOk

`func (o *ProductAdd) GetListingTypeOk() (*string, bool)`

GetListingTypeOk returns a tuple with the ListingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingType

`func (o *ProductAdd) SetListingType(v string)`

SetListingType sets ListingType field to given value.

### HasListingType

`func (o *ProductAdd) HasListingType() bool`

HasListingType returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *ProductAdd) GetPaymentMethods() []string`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *ProductAdd) GetPaymentMethodsOk() (*[]string, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *ProductAdd) SetPaymentMethods(v []string)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *ProductAdd) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetReturnAccepted

`func (o *ProductAdd) GetReturnAccepted() bool`

GetReturnAccepted returns the ReturnAccepted field if non-nil, zero value otherwise.

### GetReturnAcceptedOk

`func (o *ProductAdd) GetReturnAcceptedOk() (*bool, bool)`

GetReturnAcceptedOk returns a tuple with the ReturnAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAccepted

`func (o *ProductAdd) SetReturnAccepted(v bool)`

SetReturnAccepted sets ReturnAccepted field to given value.

### HasReturnAccepted

`func (o *ProductAdd) HasReturnAccepted() bool`

HasReturnAccepted returns a boolean if a field has been set.

### GetShippingDetails

`func (o *ProductAdd) GetShippingDetails() []ProductAddShippingDetailsInner`

GetShippingDetails returns the ShippingDetails field if non-nil, zero value otherwise.

### GetShippingDetailsOk

`func (o *ProductAdd) GetShippingDetailsOk() (*[]ProductAddShippingDetailsInner, bool)`

GetShippingDetailsOk returns a tuple with the ShippingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDetails

`func (o *ProductAdd) SetShippingDetails(v []ProductAddShippingDetailsInner)`

SetShippingDetails sets ShippingDetails field to given value.

### HasShippingDetails

`func (o *ProductAdd) HasShippingDetails() bool`

HasShippingDetails returns a boolean if a field has been set.

### GetPaypalEmail

`func (o *ProductAdd) GetPaypalEmail() string`

GetPaypalEmail returns the PaypalEmail field if non-nil, zero value otherwise.

### GetPaypalEmailOk

`func (o *ProductAdd) GetPaypalEmailOk() (*string, bool)`

GetPaypalEmailOk returns a tuple with the PaypalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalEmail

`func (o *ProductAdd) SetPaypalEmail(v string)`

SetPaypalEmail sets PaypalEmail field to given value.

### HasPaypalEmail

`func (o *ProductAdd) HasPaypalEmail() bool`

HasPaypalEmail returns a boolean if a field has been set.

### GetSellerProfiles

`func (o *ProductAdd) GetSellerProfiles() ProductAddSellerProfiles`

GetSellerProfiles returns the SellerProfiles field if non-nil, zero value otherwise.

### GetSellerProfilesOk

`func (o *ProductAdd) GetSellerProfilesOk() (*ProductAddSellerProfiles, bool)`

GetSellerProfilesOk returns a tuple with the SellerProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerProfiles

`func (o *ProductAdd) SetSellerProfiles(v ProductAddSellerProfiles)`

SetSellerProfiles sets SellerProfiles field to given value.

### HasSellerProfiles

`func (o *ProductAdd) HasSellerProfiles() bool`

HasSellerProfiles returns a boolean if a field has been set.

### GetPackageDetails

`func (o *ProductAdd) GetPackageDetails() ProductAddPackageDetails`

GetPackageDetails returns the PackageDetails field if non-nil, zero value otherwise.

### GetPackageDetailsOk

`func (o *ProductAdd) GetPackageDetailsOk() (*ProductAddPackageDetails, bool)`

GetPackageDetailsOk returns a tuple with the PackageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDetails

`func (o *ProductAdd) SetPackageDetails(v ProductAddPackageDetails)`

SetPackageDetails sets PackageDetails field to given value.

### HasPackageDetails

`func (o *ProductAdd) HasPackageDetails() bool`

HasPackageDetails returns a boolean if a field has been set.

### GetBestOffer

`func (o *ProductAdd) GetBestOffer() ProductAddBestOffer`

GetBestOffer returns the BestOffer field if non-nil, zero value otherwise.

### GetBestOfferOk

`func (o *ProductAdd) GetBestOfferOk() (*ProductAddBestOffer, bool)`

GetBestOfferOk returns a tuple with the BestOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOffer

`func (o *ProductAdd) SetBestOffer(v ProductAddBestOffer)`

SetBestOffer sets BestOffer field to given value.

### HasBestOffer

`func (o *ProductAdd) HasBestOffer() bool`

HasBestOffer returns a boolean if a field has been set.

### GetSalesTax

`func (o *ProductAdd) GetSalesTax() ProductAddSalesTax`

GetSalesTax returns the SalesTax field if non-nil, zero value otherwise.

### GetSalesTaxOk

`func (o *ProductAdd) GetSalesTaxOk() (*ProductAddSalesTax, bool)`

GetSalesTaxOk returns a tuple with the SalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTax

`func (o *ProductAdd) SetSalesTax(v ProductAddSalesTax)`

SetSalesTax sets SalesTax field to given value.

### HasSalesTax

`func (o *ProductAdd) HasSalesTax() bool`

HasSalesTax returns a boolean if a field has been set.

### GetBarcode

`func (o *ProductAdd) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductAdd) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductAdd) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ProductAdd) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetUpc

`func (o *ProductAdd) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ProductAdd) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ProductAdd) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ProductAdd) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetEan

`func (o *ProductAdd) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *ProductAdd) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *ProductAdd) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *ProductAdd) HasEan() bool`

HasEan returns a boolean if a field has been set.

### GetIsbn

`func (o *ProductAdd) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *ProductAdd) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *ProductAdd) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *ProductAdd) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### GetSpecifics

`func (o *ProductAdd) GetSpecifics() []ProductAddSpecificsInner`

GetSpecifics returns the Specifics field if non-nil, zero value otherwise.

### GetSpecificsOk

`func (o *ProductAdd) GetSpecificsOk() (*[]ProductAddSpecificsInner, bool)`

GetSpecificsOk returns a tuple with the Specifics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifics

`func (o *ProductAdd) SetSpecifics(v []ProductAddSpecificsInner)`

SetSpecifics sets Specifics field to given value.

### HasSpecifics

`func (o *ProductAdd) HasSpecifics() bool`

HasSpecifics returns a boolean if a field has been set.

### GetImageUrl

`func (o *ProductAdd) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ProductAdd) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ProductAdd) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ProductAdd) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetImageName

`func (o *ProductAdd) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ProductAdd) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ProductAdd) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ProductAdd) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetReservePrice

`func (o *ProductAdd) GetReservePrice() float32`

GetReservePrice returns the ReservePrice field if non-nil, zero value otherwise.

### GetReservePriceOk

`func (o *ProductAdd) GetReservePriceOk() (*float32, bool)`

GetReservePriceOk returns a tuple with the ReservePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservePrice

`func (o *ProductAdd) SetReservePrice(v float32)`

SetReservePrice sets ReservePrice field to given value.

### HasReservePrice

`func (o *ProductAdd) HasReservePrice() bool`

HasReservePrice returns a boolean if a field has been set.

### GetBuyitnowPrice

`func (o *ProductAdd) GetBuyitnowPrice() float32`

GetBuyitnowPrice returns the BuyitnowPrice field if non-nil, zero value otherwise.

### GetBuyitnowPriceOk

`func (o *ProductAdd) GetBuyitnowPriceOk() (*float32, bool)`

GetBuyitnowPriceOk returns a tuple with the BuyitnowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyitnowPrice

`func (o *ProductAdd) SetBuyitnowPrice(v float32)`

SetBuyitnowPrice sets BuyitnowPrice field to given value.

### HasBuyitnowPrice

`func (o *ProductAdd) HasBuyitnowPrice() bool`

HasBuyitnowPrice returns a boolean if a field has been set.

### GetConditionDescription

`func (o *ProductAdd) GetConditionDescription() string`

GetConditionDescription returns the ConditionDescription field if non-nil, zero value otherwise.

### GetConditionDescriptionOk

`func (o *ProductAdd) GetConditionDescriptionOk() (*string, bool)`

GetConditionDescriptionOk returns a tuple with the ConditionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionDescription

`func (o *ProductAdd) SetConditionDescription(v string)`

SetConditionDescription sets ConditionDescription field to given value.

### HasConditionDescription

`func (o *ProductAdd) HasConditionDescription() bool`

HasConditionDescription returns a boolean if a field has been set.

### GetAuctionConfidentialityLevel

`func (o *ProductAdd) GetAuctionConfidentialityLevel() string`

GetAuctionConfidentialityLevel returns the AuctionConfidentialityLevel field if non-nil, zero value otherwise.

### GetAuctionConfidentialityLevelOk

`func (o *ProductAdd) GetAuctionConfidentialityLevelOk() (*string, bool)`

GetAuctionConfidentialityLevelOk returns a tuple with the AuctionConfidentialityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionConfidentialityLevel

`func (o *ProductAdd) SetAuctionConfidentialityLevel(v string)`

SetAuctionConfidentialityLevel sets AuctionConfidentialityLevel field to given value.

### HasAuctionConfidentialityLevel

`func (o *ProductAdd) HasAuctionConfidentialityLevel() bool`

HasAuctionConfidentialityLevel returns a boolean if a field has been set.

### GetAvailFrom

`func (o *ProductAdd) GetAvailFrom() string`

GetAvailFrom returns the AvailFrom field if non-nil, zero value otherwise.

### GetAvailFromOk

`func (o *ProductAdd) GetAvailFromOk() (*string, bool)`

GetAvailFromOk returns a tuple with the AvailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailFrom

`func (o *ProductAdd) SetAvailFrom(v string)`

SetAvailFrom sets AvailFrom field to given value.

### HasAvailFrom

`func (o *ProductAdd) HasAvailFrom() bool`

HasAvailFrom returns a boolean if a field has been set.

### GetTags

`func (o *ProductAdd) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductAdd) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductAdd) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProductAdd) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetClearCache

`func (o *ProductAdd) GetClearCache() bool`

GetClearCache returns the ClearCache field if non-nil, zero value otherwise.

### GetClearCacheOk

`func (o *ProductAdd) GetClearCacheOk() (*bool, bool)`

GetClearCacheOk returns a tuple with the ClearCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCache

`func (o *ProductAdd) SetClearCache(v bool)`

SetClearCache sets ClearCache field to given value.

### HasClearCache

`func (o *ProductAdd) HasClearCache() bool`

HasClearCache returns a boolean if a field has been set.

### GetAsin

`func (o *ProductAdd) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *ProductAdd) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *ProductAdd) SetAsin(v string)`

SetAsin sets Asin field to given value.

### HasAsin

`func (o *ProductAdd) HasAsin() bool`

HasAsin returns a boolean if a field has been set.

### GetGtin

`func (o *ProductAdd) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *ProductAdd) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *ProductAdd) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *ProductAdd) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### GetMpn

`func (o *ProductAdd) GetMpn() string`

GetMpn returns the Mpn field if non-nil, zero value otherwise.

### GetMpnOk

`func (o *ProductAdd) GetMpnOk() (*string, bool)`

GetMpnOk returns a tuple with the Mpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpn

`func (o *ProductAdd) SetMpn(v string)`

SetMpn sets Mpn field to given value.

### HasMpn

`func (o *ProductAdd) HasMpn() bool`

HasMpn returns a boolean if a field has been set.

### GetTaxable

`func (o *ProductAdd) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ProductAdd) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ProductAdd) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ProductAdd) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetVisible

`func (o *ProductAdd) GetVisible() string`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ProductAdd) GetVisibleOk() (*string, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ProductAdd) SetVisible(v string)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ProductAdd) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetStatus

`func (o *ProductAdd) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductAdd) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductAdd) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductAdd) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeoUrl

`func (o *ProductAdd) GetSeoUrl() string`

GetSeoUrl returns the SeoUrl field if non-nil, zero value otherwise.

### GetSeoUrlOk

`func (o *ProductAdd) GetSeoUrlOk() (*string, bool)`

GetSeoUrlOk returns a tuple with the SeoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoUrl

`func (o *ProductAdd) SetSeoUrl(v string)`

SetSeoUrl sets SeoUrl field to given value.

### HasSeoUrl

`func (o *ProductAdd) HasSeoUrl() bool`

HasSeoUrl returns a boolean if a field has been set.

### GetProductClass

`func (o *ProductAdd) GetProductClass() string`

GetProductClass returns the ProductClass field if non-nil, zero value otherwise.

### GetProductClassOk

`func (o *ProductAdd) GetProductClassOk() (*string, bool)`

GetProductClassOk returns a tuple with the ProductClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductClass

`func (o *ProductAdd) SetProductClass(v string)`

SetProductClass sets ProductClass field to given value.

### HasProductClass

`func (o *ProductAdd) HasProductClass() bool`

HasProductClass returns a boolean if a field has been set.

### GetProductType

`func (o *ProductAdd) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductAdd) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductAdd) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *ProductAdd) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetMarketplaceItemProperties

`func (o *ProductAdd) GetMarketplaceItemProperties() string`

GetMarketplaceItemProperties returns the MarketplaceItemProperties field if non-nil, zero value otherwise.

### GetMarketplaceItemPropertiesOk

`func (o *ProductAdd) GetMarketplaceItemPropertiesOk() (*string, bool)`

GetMarketplaceItemPropertiesOk returns a tuple with the MarketplaceItemProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceItemProperties

`func (o *ProductAdd) SetMarketplaceItemProperties(v string)`

SetMarketplaceItemProperties sets MarketplaceItemProperties field to given value.

### HasMarketplaceItemProperties

`func (o *ProductAdd) HasMarketplaceItemProperties() bool`

HasMarketplaceItemProperties returns a boolean if a field has been set.

### GetManageStock

`func (o *ProductAdd) GetManageStock() bool`

GetManageStock returns the ManageStock field if non-nil, zero value otherwise.

### GetManageStockOk

`func (o *ProductAdd) GetManageStockOk() (*bool, bool)`

GetManageStockOk returns a tuple with the ManageStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStock

`func (o *ProductAdd) SetManageStock(v bool)`

SetManageStock sets ManageStock field to given value.

### HasManageStock

`func (o *ProductAdd) HasManageStock() bool`

HasManageStock returns a boolean if a field has been set.

### GetHarmonizedSystemCode

`func (o *ProductAdd) GetHarmonizedSystemCode() string`

GetHarmonizedSystemCode returns the HarmonizedSystemCode field if non-nil, zero value otherwise.

### GetHarmonizedSystemCodeOk

`func (o *ProductAdd) GetHarmonizedSystemCodeOk() (*string, bool)`

GetHarmonizedSystemCodeOk returns a tuple with the HarmonizedSystemCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarmonizedSystemCode

`func (o *ProductAdd) SetHarmonizedSystemCode(v string)`

SetHarmonizedSystemCode sets HarmonizedSystemCode field to given value.

### HasHarmonizedSystemCode

`func (o *ProductAdd) HasHarmonizedSystemCode() bool`

HasHarmonizedSystemCode returns a boolean if a field has been set.

### GetCountryOfOrigin

`func (o *ProductAdd) GetCountryOfOrigin() string`

GetCountryOfOrigin returns the CountryOfOrigin field if non-nil, zero value otherwise.

### GetCountryOfOriginOk

`func (o *ProductAdd) GetCountryOfOriginOk() (*string, bool)`

GetCountryOfOriginOk returns a tuple with the CountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfOrigin

`func (o *ProductAdd) SetCountryOfOrigin(v string)`

SetCountryOfOrigin sets CountryOfOrigin field to given value.

### HasCountryOfOrigin

`func (o *ProductAdd) HasCountryOfOrigin() bool`

HasCountryOfOrigin returns a boolean if a field has been set.

### GetFiles

`func (o *ProductAdd) GetFiles() []ProductAddFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProductAdd) GetFilesOk() (*[]ProductAddFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProductAdd) SetFiles(v []ProductAddFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ProductAdd) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetSearchKeywords

`func (o *ProductAdd) GetSearchKeywords() string`

GetSearchKeywords returns the SearchKeywords field if non-nil, zero value otherwise.

### GetSearchKeywordsOk

`func (o *ProductAdd) GetSearchKeywordsOk() (*string, bool)`

GetSearchKeywordsOk returns a tuple with the SearchKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKeywords

`func (o *ProductAdd) SetSearchKeywords(v string)`

SetSearchKeywords sets SearchKeywords field to given value.

### HasSearchKeywords

`func (o *ProductAdd) HasSearchKeywords() bool`

HasSearchKeywords returns a boolean if a field has been set.

### GetStoreId

`func (o *ProductAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetBrandName

`func (o *ProductAdd) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *ProductAdd) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *ProductAdd) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *ProductAdd) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.

### GetIsVirtual

`func (o *ProductAdd) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *ProductAdd) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *ProductAdd) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *ProductAdd) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetIsFreeShipping

`func (o *ProductAdd) GetIsFreeShipping() bool`

GetIsFreeShipping returns the IsFreeShipping field if non-nil, zero value otherwise.

### GetIsFreeShippingOk

`func (o *ProductAdd) GetIsFreeShippingOk() (*bool, bool)`

GetIsFreeShippingOk returns a tuple with the IsFreeShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeShipping

`func (o *ProductAdd) SetIsFreeShipping(v bool)`

SetIsFreeShipping sets IsFreeShipping field to given value.

### HasIsFreeShipping

`func (o *ProductAdd) HasIsFreeShipping() bool`

HasIsFreeShipping returns a boolean if a field has been set.

### GetInStock

`func (o *ProductAdd) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *ProductAdd) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *ProductAdd) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *ProductAdd) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### GetDeliveryCode

`func (o *ProductAdd) GetDeliveryCode() string`

GetDeliveryCode returns the DeliveryCode field if non-nil, zero value otherwise.

### GetDeliveryCodeOk

`func (o *ProductAdd) GetDeliveryCodeOk() (*string, bool)`

GetDeliveryCodeOk returns a tuple with the DeliveryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCode

`func (o *ProductAdd) SetDeliveryCode(v string)`

SetDeliveryCode sets DeliveryCode field to given value.

### HasDeliveryCode

`func (o *ProductAdd) HasDeliveryCode() bool`

HasDeliveryCode returns a boolean if a field has been set.

### GetProductReference

`func (o *ProductAdd) GetProductReference() string`

GetProductReference returns the ProductReference field if non-nil, zero value otherwise.

### GetProductReferenceOk

`func (o *ProductAdd) GetProductReferenceOk() (*string, bool)`

GetProductReferenceOk returns a tuple with the ProductReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductReference

`func (o *ProductAdd) SetProductReference(v string)`

SetProductReference sets ProductReference field to given value.

### HasProductReference

`func (o *ProductAdd) HasProductReference() bool`

HasProductReference returns a boolean if a field has been set.

### GetDeliveryType

`func (o *ProductAdd) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *ProductAdd) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *ProductAdd) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *ProductAdd) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetDeliveryTime

`func (o *ProductAdd) GetDeliveryTime() int32`

GetDeliveryTime returns the DeliveryTime field if non-nil, zero value otherwise.

### GetDeliveryTimeOk

`func (o *ProductAdd) GetDeliveryTimeOk() (*int32, bool)`

GetDeliveryTimeOk returns a tuple with the DeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTime

`func (o *ProductAdd) SetDeliveryTime(v int32)`

SetDeliveryTime sets DeliveryTime field to given value.

### HasDeliveryTime

`func (o *ProductAdd) HasDeliveryTime() bool`

HasDeliveryTime returns a boolean if a field has been set.

### GetSizeChart

`func (o *ProductAdd) GetSizeChart() ProductAddSizeChart`

GetSizeChart returns the SizeChart field if non-nil, zero value otherwise.

### GetSizeChartOk

`func (o *ProductAdd) GetSizeChartOk() (*ProductAddSizeChart, bool)`

GetSizeChartOk returns a tuple with the SizeChart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeChart

`func (o *ProductAdd) SetSizeChart(v ProductAddSizeChart)`

SetSizeChart sets SizeChart field to given value.

### HasSizeChart

`func (o *ProductAdd) HasSizeChart() bool`

HasSizeChart returns a boolean if a field has been set.

### GetCertifications

`func (o *ProductAdd) GetCertifications() []ProductAddCertificationsInner`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *ProductAdd) GetCertificationsOk() (*[]ProductAddCertificationsInner, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *ProductAdd) SetCertifications(v []ProductAddCertificationsInner)`

SetCertifications sets Certifications field to given value.

### HasCertifications

`func (o *ProductAdd) HasCertifications() bool`

HasCertifications returns a boolean if a field has been set.

### GetDeliveryOptionIds

`func (o *ProductAdd) GetDeliveryOptionIds() string`

GetDeliveryOptionIds returns the DeliveryOptionIds field if non-nil, zero value otherwise.

### GetDeliveryOptionIdsOk

`func (o *ProductAdd) GetDeliveryOptionIdsOk() (*string, bool)`

GetDeliveryOptionIdsOk returns a tuple with the DeliveryOptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryOptionIds

`func (o *ProductAdd) SetDeliveryOptionIds(v string)`

SetDeliveryOptionIds sets DeliveryOptionIds field to given value.

### HasDeliveryOptionIds

`func (o *ProductAdd) HasDeliveryOptionIds() bool`

HasDeliveryOptionIds returns a boolean if a field has been set.

### GetManufacturerInfo

`func (o *ProductAdd) GetManufacturerInfo() ProductAddManufacturerInfo`

GetManufacturerInfo returns the ManufacturerInfo field if non-nil, zero value otherwise.

### GetManufacturerInfoOk

`func (o *ProductAdd) GetManufacturerInfoOk() (*ProductAddManufacturerInfo, bool)`

GetManufacturerInfoOk returns a tuple with the ManufacturerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerInfo

`func (o *ProductAdd) SetManufacturerInfo(v ProductAddManufacturerInfo)`

SetManufacturerInfo sets ManufacturerInfo field to given value.

### HasManufacturerInfo

`func (o *ProductAdd) HasManufacturerInfo() bool`

HasManufacturerInfo returns a boolean if a field has been set.

### GetWhenMade

`func (o *ProductAdd) GetWhenMade() string`

GetWhenMade returns the WhenMade field if non-nil, zero value otherwise.

### GetWhenMadeOk

`func (o *ProductAdd) GetWhenMadeOk() (*string, bool)`

GetWhenMadeOk returns a tuple with the WhenMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenMade

`func (o *ProductAdd) SetWhenMade(v string)`

SetWhenMade sets WhenMade field to given value.

### HasWhenMade

`func (o *ProductAdd) HasWhenMade() bool`

HasWhenMade returns a boolean if a field has been set.

### GetIsSupply

`func (o *ProductAdd) GetIsSupply() bool`

GetIsSupply returns the IsSupply field if non-nil, zero value otherwise.

### GetIsSupplyOk

`func (o *ProductAdd) GetIsSupplyOk() (*bool, bool)`

GetIsSupplyOk returns a tuple with the IsSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupply

`func (o *ProductAdd) SetIsSupply(v bool)`

SetIsSupply sets IsSupply field to given value.

### HasIsSupply

`func (o *ProductAdd) HasIsSupply() bool`

HasIsSupply returns a boolean if a field has been set.

### GetMaterials

`func (o *ProductAdd) GetMaterials() []string`

GetMaterials returns the Materials field if non-nil, zero value otherwise.

### GetMaterialsOk

`func (o *ProductAdd) GetMaterialsOk() (*[]string, bool)`

GetMaterialsOk returns a tuple with the Materials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterials

`func (o *ProductAdd) SetMaterials(v []string)`

SetMaterials sets Materials field to given value.

### HasMaterials

`func (o *ProductAdd) HasMaterials() bool`

HasMaterials returns a boolean if a field has been set.

### GetAutoRenew

`func (o *ProductAdd) GetAutoRenew() bool`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *ProductAdd) GetAutoRenewOk() (*bool, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *ProductAdd) SetAutoRenew(v bool)`

SetAutoRenew sets AutoRenew field to given value.

### HasAutoRenew

`func (o *ProductAdd) HasAutoRenew() bool`

HasAutoRenew returns a boolean if a field has been set.

### GetAllowDisplayCondition

`func (o *ProductAdd) GetAllowDisplayCondition() bool`

GetAllowDisplayCondition returns the AllowDisplayCondition field if non-nil, zero value otherwise.

### GetAllowDisplayConditionOk

`func (o *ProductAdd) GetAllowDisplayConditionOk() (*bool, bool)`

GetAllowDisplayConditionOk returns a tuple with the AllowDisplayCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDisplayCondition

`func (o *ProductAdd) SetAllowDisplayCondition(v bool)`

SetAllowDisplayCondition sets AllowDisplayCondition field to given value.

### HasAllowDisplayCondition

`func (o *ProductAdd) HasAllowDisplayCondition() bool`

HasAllowDisplayCondition returns a boolean if a field has been set.

### GetMinOrderQuantity

`func (o *ProductAdd) GetMinOrderQuantity() float32`

GetMinOrderQuantity returns the MinOrderQuantity field if non-nil, zero value otherwise.

### GetMinOrderQuantityOk

`func (o *ProductAdd) GetMinOrderQuantityOk() (*float32, bool)`

GetMinOrderQuantityOk returns a tuple with the MinOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrderQuantity

`func (o *ProductAdd) SetMinOrderQuantity(v float32)`

SetMinOrderQuantity sets MinOrderQuantity field to given value.

### HasMinOrderQuantity

`func (o *ProductAdd) HasMinOrderQuantity() bool`

HasMinOrderQuantity returns a boolean if a field has been set.

### GetMaxOrderQuantity

`func (o *ProductAdd) GetMaxOrderQuantity() float32`

GetMaxOrderQuantity returns the MaxOrderQuantity field if non-nil, zero value otherwise.

### GetMaxOrderQuantityOk

`func (o *ProductAdd) GetMaxOrderQuantityOk() (*float32, bool)`

GetMaxOrderQuantityOk returns a tuple with the MaxOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrderQuantity

`func (o *ProductAdd) SetMaxOrderQuantity(v float32)`

SetMaxOrderQuantity sets MaxOrderQuantity field to given value.

### HasMaxOrderQuantity

`func (o *ProductAdd) HasMaxOrderQuantity() bool`

HasMaxOrderQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


