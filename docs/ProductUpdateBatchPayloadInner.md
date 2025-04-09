# ProductUpdateBatchPayloadInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**SpecialPrice** | Pointer to **float32** |  | [optional] 
**SpriceCreate** | Pointer to **string** |  | [optional] 
**SpriceExpire** | Pointer to **string** |  | [optional] 
**CostPrice** | Pointer to **float32** |  | [optional] 
**OldPrice** | Pointer to **float32** |  | [optional] 
**FixedCostShippingPrice** | Pointer to **float32** |  | [optional] 
**AdvancedPrices** | Pointer to [**[]ProductUpdateBatchPayloadInnerAdvancedPricesInner**](ProductUpdateBatchPayloadInnerAdvancedPricesInner.md) | If an empty array is passed, all entries will be deleted when the &#39;nested_items_update_behaviour&#39; parameter is set to &#39;replace&#39;. | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**IncreaseQuantity** | Pointer to **float32** |  | [optional] 
**ReduceQuantity** | Pointer to **float32** |  | [optional] 
**ReserveQuantity** | Pointer to **float32** |  | [optional] 
**StoreId** | Pointer to **string** |  | [optional] 
**LangId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to **string** |  | [optional] 
**Visible** | Pointer to **string** |  | [optional] 
**AvailableForView** | Pointer to **bool** |  | [optional] 
**AvailableForSale** | Pointer to **bool** |  | [optional] 
**AvailFrom** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] 
**Length** | Pointer to **float32** |  | [optional] 
**Width** | Pointer to **float32** |  | [optional] 
**Height** | Pointer to **float32** |  | [optional] 
**DimensionsUnit** | Pointer to **string** |  | [optional] 
**WeightUnit** | Pointer to **string** |  | [optional] 
**ManageStock** | Pointer to **bool** |  | [optional] 
**InStock** | Pointer to **bool** |  | [optional] 
**BackorderStatus** | Pointer to **string** |  | [optional] 
**IsFreeShipping** | Pointer to **bool** |  | [optional] 
**IsVirtual** | Pointer to **bool** |  | [optional] 
**Taxable** | Pointer to **bool** |  | [optional] 
**Downloadable** | Pointer to **bool** |  | [optional] 
**WarehouseId** | Pointer to **string** |  | [optional] 
**TaxClassId** | Pointer to **string** |  | [optional] 
**CategoriesIds** | Pointer to **[]string** |  | [optional] 
**MetaTitle** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**MetaKeywords** | Pointer to **[]string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**SeoUrl** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**ManufacturerId** | Pointer to **string** |  | [optional] 
**Mpn** | Pointer to **string** |  | [optional] 
**Gtin** | Pointer to **string** |  | [optional] 
**Upc** | Pointer to **string** |  | [optional] 
**Isbn** | Pointer to **string** |  | [optional] 
**Ean** | Pointer to **string** |  | [optional] 
**Barcode** | Pointer to **string** |  | [optional] 
**Images** | Pointer to [**[]ProductUpdateBatchPayloadInnerImagesInner**](ProductUpdateBatchPayloadInnerImagesInner.md) | Property &#39;nested_items_update_behaviour&#39; does not apply. Specified items will be added to existing product images | [optional] 
**RelatedProductsIds** | Pointer to **[]string** | If an empty array is passed, all entries will be deleted when the &#39;nested_items_update_behaviour&#39; parameter is set to &#39;replace&#39;. | [optional] 
**UpSellProductsIds** | Pointer to **[]string** | If an empty array is passed, all entries will be deleted when the &#39;nested_items_update_behaviour&#39; parameter is set to &#39;replace&#39;. | [optional] 
**CrossSellProductsIds** | Pointer to **[]string** | If an empty array is passed, all entries will be deleted when the &#39;nested_items_update_behaviour&#39; parameter is set to &#39;replace&#39;. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**SearchKeywords** | Pointer to **[]string** |  | [optional] 
**HarmonizedSystemCode** | Pointer to **string** |  | [optional] 

## Methods

### NewProductUpdateBatchPayloadInner

`func NewProductUpdateBatchPayloadInner(id string, ) *ProductUpdateBatchPayloadInner`

NewProductUpdateBatchPayloadInner instantiates a new ProductUpdateBatchPayloadInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductUpdateBatchPayloadInnerWithDefaults

`func NewProductUpdateBatchPayloadInnerWithDefaults() *ProductUpdateBatchPayloadInner`

NewProductUpdateBatchPayloadInnerWithDefaults instantiates a new ProductUpdateBatchPayloadInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductUpdateBatchPayloadInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductUpdateBatchPayloadInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductUpdateBatchPayloadInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProductUpdateBatchPayloadInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductUpdateBatchPayloadInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductUpdateBatchPayloadInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductUpdateBatchPayloadInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProductUpdateBatchPayloadInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductUpdateBatchPayloadInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductUpdateBatchPayloadInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductUpdateBatchPayloadInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortDescription

`func (o *ProductUpdateBatchPayloadInner) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ProductUpdateBatchPayloadInner) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ProductUpdateBatchPayloadInner) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ProductUpdateBatchPayloadInner) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetSku

`func (o *ProductUpdateBatchPayloadInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductUpdateBatchPayloadInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductUpdateBatchPayloadInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductUpdateBatchPayloadInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetModel

`func (o *ProductUpdateBatchPayloadInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ProductUpdateBatchPayloadInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ProductUpdateBatchPayloadInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ProductUpdateBatchPayloadInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPrice

`func (o *ProductUpdateBatchPayloadInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductUpdateBatchPayloadInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductUpdateBatchPayloadInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductUpdateBatchPayloadInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSpecialPrice

`func (o *ProductUpdateBatchPayloadInner) GetSpecialPrice() float32`

GetSpecialPrice returns the SpecialPrice field if non-nil, zero value otherwise.

### GetSpecialPriceOk

`func (o *ProductUpdateBatchPayloadInner) GetSpecialPriceOk() (*float32, bool)`

GetSpecialPriceOk returns a tuple with the SpecialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPrice

`func (o *ProductUpdateBatchPayloadInner) SetSpecialPrice(v float32)`

SetSpecialPrice sets SpecialPrice field to given value.

### HasSpecialPrice

`func (o *ProductUpdateBatchPayloadInner) HasSpecialPrice() bool`

HasSpecialPrice returns a boolean if a field has been set.

### GetSpriceCreate

`func (o *ProductUpdateBatchPayloadInner) GetSpriceCreate() string`

GetSpriceCreate returns the SpriceCreate field if non-nil, zero value otherwise.

### GetSpriceCreateOk

`func (o *ProductUpdateBatchPayloadInner) GetSpriceCreateOk() (*string, bool)`

GetSpriceCreateOk returns a tuple with the SpriceCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceCreate

`func (o *ProductUpdateBatchPayloadInner) SetSpriceCreate(v string)`

SetSpriceCreate sets SpriceCreate field to given value.

### HasSpriceCreate

`func (o *ProductUpdateBatchPayloadInner) HasSpriceCreate() bool`

HasSpriceCreate returns a boolean if a field has been set.

### GetSpriceExpire

`func (o *ProductUpdateBatchPayloadInner) GetSpriceExpire() string`

GetSpriceExpire returns the SpriceExpire field if non-nil, zero value otherwise.

### GetSpriceExpireOk

`func (o *ProductUpdateBatchPayloadInner) GetSpriceExpireOk() (*string, bool)`

GetSpriceExpireOk returns a tuple with the SpriceExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceExpire

`func (o *ProductUpdateBatchPayloadInner) SetSpriceExpire(v string)`

SetSpriceExpire sets SpriceExpire field to given value.

### HasSpriceExpire

`func (o *ProductUpdateBatchPayloadInner) HasSpriceExpire() bool`

HasSpriceExpire returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductUpdateBatchPayloadInner) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductUpdateBatchPayloadInner) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductUpdateBatchPayloadInner) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductUpdateBatchPayloadInner) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetOldPrice

`func (o *ProductUpdateBatchPayloadInner) GetOldPrice() float32`

GetOldPrice returns the OldPrice field if non-nil, zero value otherwise.

### GetOldPriceOk

`func (o *ProductUpdateBatchPayloadInner) GetOldPriceOk() (*float32, bool)`

GetOldPriceOk returns a tuple with the OldPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPrice

`func (o *ProductUpdateBatchPayloadInner) SetOldPrice(v float32)`

SetOldPrice sets OldPrice field to given value.

### HasOldPrice

`func (o *ProductUpdateBatchPayloadInner) HasOldPrice() bool`

HasOldPrice returns a boolean if a field has been set.

### GetFixedCostShippingPrice

`func (o *ProductUpdateBatchPayloadInner) GetFixedCostShippingPrice() float32`

GetFixedCostShippingPrice returns the FixedCostShippingPrice field if non-nil, zero value otherwise.

### GetFixedCostShippingPriceOk

`func (o *ProductUpdateBatchPayloadInner) GetFixedCostShippingPriceOk() (*float32, bool)`

GetFixedCostShippingPriceOk returns a tuple with the FixedCostShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCostShippingPrice

`func (o *ProductUpdateBatchPayloadInner) SetFixedCostShippingPrice(v float32)`

SetFixedCostShippingPrice sets FixedCostShippingPrice field to given value.

### HasFixedCostShippingPrice

`func (o *ProductUpdateBatchPayloadInner) HasFixedCostShippingPrice() bool`

HasFixedCostShippingPrice returns a boolean if a field has been set.

### GetAdvancedPrices

`func (o *ProductUpdateBatchPayloadInner) GetAdvancedPrices() []ProductUpdateBatchPayloadInnerAdvancedPricesInner`

GetAdvancedPrices returns the AdvancedPrices field if non-nil, zero value otherwise.

### GetAdvancedPricesOk

`func (o *ProductUpdateBatchPayloadInner) GetAdvancedPricesOk() (*[]ProductUpdateBatchPayloadInnerAdvancedPricesInner, bool)`

GetAdvancedPricesOk returns a tuple with the AdvancedPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPrices

`func (o *ProductUpdateBatchPayloadInner) SetAdvancedPrices(v []ProductUpdateBatchPayloadInnerAdvancedPricesInner)`

SetAdvancedPrices sets AdvancedPrices field to given value.

### HasAdvancedPrices

`func (o *ProductUpdateBatchPayloadInner) HasAdvancedPrices() bool`

HasAdvancedPrices returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductUpdateBatchPayloadInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductUpdateBatchPayloadInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductUpdateBatchPayloadInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductUpdateBatchPayloadInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetIncreaseQuantity

`func (o *ProductUpdateBatchPayloadInner) GetIncreaseQuantity() float32`

GetIncreaseQuantity returns the IncreaseQuantity field if non-nil, zero value otherwise.

### GetIncreaseQuantityOk

`func (o *ProductUpdateBatchPayloadInner) GetIncreaseQuantityOk() (*float32, bool)`

GetIncreaseQuantityOk returns a tuple with the IncreaseQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseQuantity

`func (o *ProductUpdateBatchPayloadInner) SetIncreaseQuantity(v float32)`

SetIncreaseQuantity sets IncreaseQuantity field to given value.

### HasIncreaseQuantity

`func (o *ProductUpdateBatchPayloadInner) HasIncreaseQuantity() bool`

HasIncreaseQuantity returns a boolean if a field has been set.

### GetReduceQuantity

`func (o *ProductUpdateBatchPayloadInner) GetReduceQuantity() float32`

GetReduceQuantity returns the ReduceQuantity field if non-nil, zero value otherwise.

### GetReduceQuantityOk

`func (o *ProductUpdateBatchPayloadInner) GetReduceQuantityOk() (*float32, bool)`

GetReduceQuantityOk returns a tuple with the ReduceQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceQuantity

`func (o *ProductUpdateBatchPayloadInner) SetReduceQuantity(v float32)`

SetReduceQuantity sets ReduceQuantity field to given value.

### HasReduceQuantity

`func (o *ProductUpdateBatchPayloadInner) HasReduceQuantity() bool`

HasReduceQuantity returns a boolean if a field has been set.

### GetReserveQuantity

`func (o *ProductUpdateBatchPayloadInner) GetReserveQuantity() float32`

GetReserveQuantity returns the ReserveQuantity field if non-nil, zero value otherwise.

### GetReserveQuantityOk

`func (o *ProductUpdateBatchPayloadInner) GetReserveQuantityOk() (*float32, bool)`

GetReserveQuantityOk returns a tuple with the ReserveQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveQuantity

`func (o *ProductUpdateBatchPayloadInner) SetReserveQuantity(v float32)`

SetReserveQuantity sets ReserveQuantity field to given value.

### HasReserveQuantity

`func (o *ProductUpdateBatchPayloadInner) HasReserveQuantity() bool`

HasReserveQuantity returns a boolean if a field has been set.

### GetStoreId

`func (o *ProductUpdateBatchPayloadInner) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductUpdateBatchPayloadInner) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductUpdateBatchPayloadInner) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductUpdateBatchPayloadInner) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetLangId

`func (o *ProductUpdateBatchPayloadInner) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *ProductUpdateBatchPayloadInner) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *ProductUpdateBatchPayloadInner) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *ProductUpdateBatchPayloadInner) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetStatus

`func (o *ProductUpdateBatchPayloadInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductUpdateBatchPayloadInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductUpdateBatchPayloadInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductUpdateBatchPayloadInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ProductUpdateBatchPayloadInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductUpdateBatchPayloadInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductUpdateBatchPayloadInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductUpdateBatchPayloadInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCondition

`func (o *ProductUpdateBatchPayloadInner) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ProductUpdateBatchPayloadInner) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ProductUpdateBatchPayloadInner) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ProductUpdateBatchPayloadInner) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetVisible

`func (o *ProductUpdateBatchPayloadInner) GetVisible() string`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ProductUpdateBatchPayloadInner) GetVisibleOk() (*string, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ProductUpdateBatchPayloadInner) SetVisible(v string)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ProductUpdateBatchPayloadInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetAvailableForView

`func (o *ProductUpdateBatchPayloadInner) GetAvailableForView() bool`

GetAvailableForView returns the AvailableForView field if non-nil, zero value otherwise.

### GetAvailableForViewOk

`func (o *ProductUpdateBatchPayloadInner) GetAvailableForViewOk() (*bool, bool)`

GetAvailableForViewOk returns a tuple with the AvailableForView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForView

`func (o *ProductUpdateBatchPayloadInner) SetAvailableForView(v bool)`

SetAvailableForView sets AvailableForView field to given value.

### HasAvailableForView

`func (o *ProductUpdateBatchPayloadInner) HasAvailableForView() bool`

HasAvailableForView returns a boolean if a field has been set.

### GetAvailableForSale

`func (o *ProductUpdateBatchPayloadInner) GetAvailableForSale() bool`

GetAvailableForSale returns the AvailableForSale field if non-nil, zero value otherwise.

### GetAvailableForSaleOk

`func (o *ProductUpdateBatchPayloadInner) GetAvailableForSaleOk() (*bool, bool)`

GetAvailableForSaleOk returns a tuple with the AvailableForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForSale

`func (o *ProductUpdateBatchPayloadInner) SetAvailableForSale(v bool)`

SetAvailableForSale sets AvailableForSale field to given value.

### HasAvailableForSale

`func (o *ProductUpdateBatchPayloadInner) HasAvailableForSale() bool`

HasAvailableForSale returns a boolean if a field has been set.

### GetAvailFrom

`func (o *ProductUpdateBatchPayloadInner) GetAvailFrom() string`

GetAvailFrom returns the AvailFrom field if non-nil, zero value otherwise.

### GetAvailFromOk

`func (o *ProductUpdateBatchPayloadInner) GetAvailFromOk() (*string, bool)`

GetAvailFromOk returns a tuple with the AvailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailFrom

`func (o *ProductUpdateBatchPayloadInner) SetAvailFrom(v string)`

SetAvailFrom sets AvailFrom field to given value.

### HasAvailFrom

`func (o *ProductUpdateBatchPayloadInner) HasAvailFrom() bool`

HasAvailFrom returns a boolean if a field has been set.

### GetWeight

`func (o *ProductUpdateBatchPayloadInner) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductUpdateBatchPayloadInner) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductUpdateBatchPayloadInner) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductUpdateBatchPayloadInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetLength

`func (o *ProductUpdateBatchPayloadInner) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ProductUpdateBatchPayloadInner) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ProductUpdateBatchPayloadInner) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ProductUpdateBatchPayloadInner) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWidth

`func (o *ProductUpdateBatchPayloadInner) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ProductUpdateBatchPayloadInner) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ProductUpdateBatchPayloadInner) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ProductUpdateBatchPayloadInner) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ProductUpdateBatchPayloadInner) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ProductUpdateBatchPayloadInner) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ProductUpdateBatchPayloadInner) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ProductUpdateBatchPayloadInner) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetDimensionsUnit

`func (o *ProductUpdateBatchPayloadInner) GetDimensionsUnit() string`

GetDimensionsUnit returns the DimensionsUnit field if non-nil, zero value otherwise.

### GetDimensionsUnitOk

`func (o *ProductUpdateBatchPayloadInner) GetDimensionsUnitOk() (*string, bool)`

GetDimensionsUnitOk returns a tuple with the DimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsUnit

`func (o *ProductUpdateBatchPayloadInner) SetDimensionsUnit(v string)`

SetDimensionsUnit sets DimensionsUnit field to given value.

### HasDimensionsUnit

`func (o *ProductUpdateBatchPayloadInner) HasDimensionsUnit() bool`

HasDimensionsUnit returns a boolean if a field has been set.

### GetWeightUnit

`func (o *ProductUpdateBatchPayloadInner) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ProductUpdateBatchPayloadInner) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ProductUpdateBatchPayloadInner) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ProductUpdateBatchPayloadInner) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetManageStock

`func (o *ProductUpdateBatchPayloadInner) GetManageStock() bool`

GetManageStock returns the ManageStock field if non-nil, zero value otherwise.

### GetManageStockOk

`func (o *ProductUpdateBatchPayloadInner) GetManageStockOk() (*bool, bool)`

GetManageStockOk returns a tuple with the ManageStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStock

`func (o *ProductUpdateBatchPayloadInner) SetManageStock(v bool)`

SetManageStock sets ManageStock field to given value.

### HasManageStock

`func (o *ProductUpdateBatchPayloadInner) HasManageStock() bool`

HasManageStock returns a boolean if a field has been set.

### GetInStock

`func (o *ProductUpdateBatchPayloadInner) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *ProductUpdateBatchPayloadInner) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *ProductUpdateBatchPayloadInner) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *ProductUpdateBatchPayloadInner) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### GetBackorderStatus

`func (o *ProductUpdateBatchPayloadInner) GetBackorderStatus() string`

GetBackorderStatus returns the BackorderStatus field if non-nil, zero value otherwise.

### GetBackorderStatusOk

`func (o *ProductUpdateBatchPayloadInner) GetBackorderStatusOk() (*string, bool)`

GetBackorderStatusOk returns a tuple with the BackorderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderStatus

`func (o *ProductUpdateBatchPayloadInner) SetBackorderStatus(v string)`

SetBackorderStatus sets BackorderStatus field to given value.

### HasBackorderStatus

`func (o *ProductUpdateBatchPayloadInner) HasBackorderStatus() bool`

HasBackorderStatus returns a boolean if a field has been set.

### GetIsFreeShipping

`func (o *ProductUpdateBatchPayloadInner) GetIsFreeShipping() bool`

GetIsFreeShipping returns the IsFreeShipping field if non-nil, zero value otherwise.

### GetIsFreeShippingOk

`func (o *ProductUpdateBatchPayloadInner) GetIsFreeShippingOk() (*bool, bool)`

GetIsFreeShippingOk returns a tuple with the IsFreeShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeShipping

`func (o *ProductUpdateBatchPayloadInner) SetIsFreeShipping(v bool)`

SetIsFreeShipping sets IsFreeShipping field to given value.

### HasIsFreeShipping

`func (o *ProductUpdateBatchPayloadInner) HasIsFreeShipping() bool`

HasIsFreeShipping returns a boolean if a field has been set.

### GetIsVirtual

`func (o *ProductUpdateBatchPayloadInner) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *ProductUpdateBatchPayloadInner) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *ProductUpdateBatchPayloadInner) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *ProductUpdateBatchPayloadInner) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetTaxable

`func (o *ProductUpdateBatchPayloadInner) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ProductUpdateBatchPayloadInner) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ProductUpdateBatchPayloadInner) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ProductUpdateBatchPayloadInner) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetDownloadable

`func (o *ProductUpdateBatchPayloadInner) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *ProductUpdateBatchPayloadInner) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *ProductUpdateBatchPayloadInner) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *ProductUpdateBatchPayloadInner) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetWarehouseId

`func (o *ProductUpdateBatchPayloadInner) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ProductUpdateBatchPayloadInner) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ProductUpdateBatchPayloadInner) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ProductUpdateBatchPayloadInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetTaxClassId

`func (o *ProductUpdateBatchPayloadInner) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *ProductUpdateBatchPayloadInner) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *ProductUpdateBatchPayloadInner) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *ProductUpdateBatchPayloadInner) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### GetCategoriesIds

`func (o *ProductUpdateBatchPayloadInner) GetCategoriesIds() []string`

GetCategoriesIds returns the CategoriesIds field if non-nil, zero value otherwise.

### GetCategoriesIdsOk

`func (o *ProductUpdateBatchPayloadInner) GetCategoriesIdsOk() (*[]string, bool)`

GetCategoriesIdsOk returns a tuple with the CategoriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesIds

`func (o *ProductUpdateBatchPayloadInner) SetCategoriesIds(v []string)`

SetCategoriesIds sets CategoriesIds field to given value.

### HasCategoriesIds

`func (o *ProductUpdateBatchPayloadInner) HasCategoriesIds() bool`

HasCategoriesIds returns a boolean if a field has been set.

### GetMetaTitle

`func (o *ProductUpdateBatchPayloadInner) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *ProductUpdateBatchPayloadInner) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *ProductUpdateBatchPayloadInner) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *ProductUpdateBatchPayloadInner) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetMetaDescription

`func (o *ProductUpdateBatchPayloadInner) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *ProductUpdateBatchPayloadInner) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *ProductUpdateBatchPayloadInner) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *ProductUpdateBatchPayloadInner) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *ProductUpdateBatchPayloadInner) GetMetaKeywords() []string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *ProductUpdateBatchPayloadInner) GetMetaKeywordsOk() (*[]string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *ProductUpdateBatchPayloadInner) SetMetaKeywords(v []string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *ProductUpdateBatchPayloadInner) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetUrl

`func (o *ProductUpdateBatchPayloadInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProductUpdateBatchPayloadInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProductUpdateBatchPayloadInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProductUpdateBatchPayloadInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSeoUrl

`func (o *ProductUpdateBatchPayloadInner) GetSeoUrl() string`

GetSeoUrl returns the SeoUrl field if non-nil, zero value otherwise.

### GetSeoUrlOk

`func (o *ProductUpdateBatchPayloadInner) GetSeoUrlOk() (*string, bool)`

GetSeoUrlOk returns a tuple with the SeoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoUrl

`func (o *ProductUpdateBatchPayloadInner) SetSeoUrl(v string)`

SetSeoUrl sets SeoUrl field to given value.

### HasSeoUrl

`func (o *ProductUpdateBatchPayloadInner) HasSeoUrl() bool`

HasSeoUrl returns a boolean if a field has been set.

### GetManufacturer

`func (o *ProductUpdateBatchPayloadInner) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ProductUpdateBatchPayloadInner) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ProductUpdateBatchPayloadInner) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ProductUpdateBatchPayloadInner) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetManufacturerId

`func (o *ProductUpdateBatchPayloadInner) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *ProductUpdateBatchPayloadInner) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *ProductUpdateBatchPayloadInner) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *ProductUpdateBatchPayloadInner) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetMpn

`func (o *ProductUpdateBatchPayloadInner) GetMpn() string`

GetMpn returns the Mpn field if non-nil, zero value otherwise.

### GetMpnOk

`func (o *ProductUpdateBatchPayloadInner) GetMpnOk() (*string, bool)`

GetMpnOk returns a tuple with the Mpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpn

`func (o *ProductUpdateBatchPayloadInner) SetMpn(v string)`

SetMpn sets Mpn field to given value.

### HasMpn

`func (o *ProductUpdateBatchPayloadInner) HasMpn() bool`

HasMpn returns a boolean if a field has been set.

### GetGtin

`func (o *ProductUpdateBatchPayloadInner) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *ProductUpdateBatchPayloadInner) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *ProductUpdateBatchPayloadInner) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *ProductUpdateBatchPayloadInner) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### GetUpc

`func (o *ProductUpdateBatchPayloadInner) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ProductUpdateBatchPayloadInner) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ProductUpdateBatchPayloadInner) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ProductUpdateBatchPayloadInner) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetIsbn

`func (o *ProductUpdateBatchPayloadInner) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *ProductUpdateBatchPayloadInner) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *ProductUpdateBatchPayloadInner) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *ProductUpdateBatchPayloadInner) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### GetEan

`func (o *ProductUpdateBatchPayloadInner) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *ProductUpdateBatchPayloadInner) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *ProductUpdateBatchPayloadInner) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *ProductUpdateBatchPayloadInner) HasEan() bool`

HasEan returns a boolean if a field has been set.

### GetBarcode

`func (o *ProductUpdateBatchPayloadInner) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductUpdateBatchPayloadInner) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductUpdateBatchPayloadInner) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ProductUpdateBatchPayloadInner) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetImages

`func (o *ProductUpdateBatchPayloadInner) GetImages() []ProductUpdateBatchPayloadInnerImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ProductUpdateBatchPayloadInner) GetImagesOk() (*[]ProductUpdateBatchPayloadInnerImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ProductUpdateBatchPayloadInner) SetImages(v []ProductUpdateBatchPayloadInnerImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *ProductUpdateBatchPayloadInner) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetRelatedProductsIds

`func (o *ProductUpdateBatchPayloadInner) GetRelatedProductsIds() []string`

GetRelatedProductsIds returns the RelatedProductsIds field if non-nil, zero value otherwise.

### GetRelatedProductsIdsOk

`func (o *ProductUpdateBatchPayloadInner) GetRelatedProductsIdsOk() (*[]string, bool)`

GetRelatedProductsIdsOk returns a tuple with the RelatedProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedProductsIds

`func (o *ProductUpdateBatchPayloadInner) SetRelatedProductsIds(v []string)`

SetRelatedProductsIds sets RelatedProductsIds field to given value.

### HasRelatedProductsIds

`func (o *ProductUpdateBatchPayloadInner) HasRelatedProductsIds() bool`

HasRelatedProductsIds returns a boolean if a field has been set.

### GetUpSellProductsIds

`func (o *ProductUpdateBatchPayloadInner) GetUpSellProductsIds() []string`

GetUpSellProductsIds returns the UpSellProductsIds field if non-nil, zero value otherwise.

### GetUpSellProductsIdsOk

`func (o *ProductUpdateBatchPayloadInner) GetUpSellProductsIdsOk() (*[]string, bool)`

GetUpSellProductsIdsOk returns a tuple with the UpSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSellProductsIds

`func (o *ProductUpdateBatchPayloadInner) SetUpSellProductsIds(v []string)`

SetUpSellProductsIds sets UpSellProductsIds field to given value.

### HasUpSellProductsIds

`func (o *ProductUpdateBatchPayloadInner) HasUpSellProductsIds() bool`

HasUpSellProductsIds returns a boolean if a field has been set.

### GetCrossSellProductsIds

`func (o *ProductUpdateBatchPayloadInner) GetCrossSellProductsIds() []string`

GetCrossSellProductsIds returns the CrossSellProductsIds field if non-nil, zero value otherwise.

### GetCrossSellProductsIdsOk

`func (o *ProductUpdateBatchPayloadInner) GetCrossSellProductsIdsOk() (*[]string, bool)`

GetCrossSellProductsIdsOk returns a tuple with the CrossSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSellProductsIds

`func (o *ProductUpdateBatchPayloadInner) SetCrossSellProductsIds(v []string)`

SetCrossSellProductsIds sets CrossSellProductsIds field to given value.

### HasCrossSellProductsIds

`func (o *ProductUpdateBatchPayloadInner) HasCrossSellProductsIds() bool`

HasCrossSellProductsIds returns a boolean if a field has been set.

### GetTags

`func (o *ProductUpdateBatchPayloadInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductUpdateBatchPayloadInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductUpdateBatchPayloadInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProductUpdateBatchPayloadInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSearchKeywords

`func (o *ProductUpdateBatchPayloadInner) GetSearchKeywords() []string`

GetSearchKeywords returns the SearchKeywords field if non-nil, zero value otherwise.

### GetSearchKeywordsOk

`func (o *ProductUpdateBatchPayloadInner) GetSearchKeywordsOk() (*[]string, bool)`

GetSearchKeywordsOk returns a tuple with the SearchKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKeywords

`func (o *ProductUpdateBatchPayloadInner) SetSearchKeywords(v []string)`

SetSearchKeywords sets SearchKeywords field to given value.

### HasSearchKeywords

`func (o *ProductUpdateBatchPayloadInner) HasSearchKeywords() bool`

HasSearchKeywords returns a boolean if a field has been set.

### GetHarmonizedSystemCode

`func (o *ProductUpdateBatchPayloadInner) GetHarmonizedSystemCode() string`

GetHarmonizedSystemCode returns the HarmonizedSystemCode field if non-nil, zero value otherwise.

### GetHarmonizedSystemCodeOk

`func (o *ProductUpdateBatchPayloadInner) GetHarmonizedSystemCodeOk() (*string, bool)`

GetHarmonizedSystemCodeOk returns a tuple with the HarmonizedSystemCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarmonizedSystemCode

`func (o *ProductUpdateBatchPayloadInner) SetHarmonizedSystemCode(v string)`

SetHarmonizedSystemCode sets HarmonizedSystemCode field to given value.

### HasHarmonizedSystemCode

`func (o *ProductUpdateBatchPayloadInner) HasHarmonizedSystemCode() bool`

HasHarmonizedSystemCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


