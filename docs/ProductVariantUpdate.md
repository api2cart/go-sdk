# ProductVariantUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Defines variant update specified by variant id | [optional] 
**ProductId** | Pointer to **string** | Defines product&#39;s id where the variant has to be updated | [optional] 
**StoreId** | Pointer to **string** | Defines store id where the variant should be found | [optional] 
**LangId** | Pointer to **string** | Language id | [optional] 
**Options** | Pointer to [**[]ProductVariantUpdateOptionsInner**](ProductVariantUpdateOptionsInner.md) | Defines variant&#39;s options list | [optional] 
**Name** | Pointer to **string** | Defines variant&#39;s name that has to be updated | [optional] 
**Description** | Pointer to **string** | Specifies variant&#39;s description | [optional] 
**ShortDescription** | Pointer to **string** | Defines short description | [optional] 
**Model** | Pointer to **string** | Specifies variant&#39;s model that has to be added | [optional] 
**Sku** | Pointer to **string** | Defines new product&#39;s variant sku | [optional] 
**Visible** | Pointer to **string** | Set visibility status | [optional] 
**Status** | Pointer to **string** | Defines product variant&#39;s status | [optional] 
**BackorderStatus** | Pointer to **string** | Set backorder status | [optional] 
**AvailableForSale** | Pointer to **bool** | Specifies the set of visible/invisible product&#39;s variants for sale | [optional] [default to true]
**Avail** | Pointer to **bool** | Defines category&#39;s visibility status | [optional] [default to true]
**IsDefault** | Pointer to **bool** | Defines as a default variant | [optional] 
**IsFreeShipping** | Pointer to **bool** | Specifies variant&#39;s free shipping flag that has to be added | [optional] 
**Taxable** | Pointer to **bool** | Specifies whether a tax is charged | [optional] [default to true]
**TaxClassId** | Pointer to **string** | Defines tax classes where entity has to be added | [optional] 
**IsVirtual** | Pointer to **bool** | Defines whether the product is virtual | [optional] [default to false]
**ManageStock** | Pointer to **bool** | Defines inventory tracking for product variant | [optional] 
**InStock** | Pointer to **bool** | Set stock status | [optional] 
**WarehouseId** | Pointer to **string** | This parameter is used for selecting a warehouse where you need to set/modify a product quantity. | [optional] 
**ReserveQuantity** | Pointer to **float32** | This parameter allows to reserve/unreserve product variants quantity. | [optional] 
**Quantity** | Pointer to **float32** | Defines new products&#39; variants quantity | [optional] 
**IncreaseQuantity** | Pointer to **float32** | Defines the incremental changes in product quantity | [optional] [default to 0]
**ReduceQuantity** | Pointer to **float32** | Defines the decrement changes in product quantity | [optional] [default to 0]
**Price** | Pointer to **float32** | Defines new product&#39;s variant price | [optional] 
**SpecialPrice** | Pointer to **float32** | Defines new product&#39;s variant special price | [optional] 
**RetailPrice** | Pointer to **float32** | Defines new product&#39;s retail price | [optional] 
**OldPrice** | Pointer to **float32** | Defines product&#39;s old price | [optional] 
**CostPrice** | Pointer to **float32** | Defines new product&#39;s cost price | [optional] 
**FixedCostShippingPrice** | Pointer to **float32** | Specifies fixed cost shipping price | [optional] 
**SpriceCreate** | Pointer to **string** | Defines the date of special price creation | [optional] 
**SpriceExpire** | Pointer to **string** | Defines the term of special price offer duration | [optional] 
**Weight** | Pointer to **float32** | Weight | [optional] [default to 0]
**Barcode** | Pointer to **string** | A barcode is a unique code composed of numbers used as a product identifier. | [optional] 
**Width** | Pointer to **float32** | Defines product&#39;s width | [optional] 
**WeightUnit** | Pointer to **string** | Weight Unit | [optional] 
**Height** | Pointer to **float32** | Defines product&#39;s height | [optional] 
**Length** | Pointer to **float32** | Defines product&#39;s length | [optional] 
**Gtin** | Pointer to **string** | Global Trade Item Number. An GTIN is an identifier for trade items. | [optional] 
**Upc** | Pointer to **string** | Universal Product Code. A UPC (UPC-A) is a commonly used identifer for many different products. | [optional] 
**Mpn** | Pointer to **string** | Manufacturer Part Number. A MPN is an identifier of a particular part design or material used. | [optional] 
**Ean** | Pointer to **string** | European Article Number. An EAN is a unique 8 or 13-digit identifier that many industries (such as book publishers) use to identify products. | [optional] 
**Isbn** | Pointer to **string** | International Standard Book Number. An ISBN is a unique identifier for books. | [optional] 
**HarmonizedSystemCode** | Pointer to **string** | Harmonized System Code. An HSC is a 6-digit identifier that allows participating countries to classify traded goods on a common basis for customs purposes | [optional] 
**CountryOfOrigin** | Pointer to **string** | The country where the inventory item was made | [optional] 
**MetaTitle** | Pointer to **string** | Defines unique meta title for each entity | [optional] 
**MetaDescription** | Pointer to **string** | Defines unique meta description of a entity | [optional] 
**MetaKeywords** | Pointer to **string** | Defines unique meta keywords for each entity | [optional] 
**Reindex** | Pointer to **bool** | Is reindex required | [optional] [default to true]
**ClearCache** | Pointer to **bool** | Is cache clear required | [optional] [default to true]

## Methods

### NewProductVariantUpdate

`func NewProductVariantUpdate() *ProductVariantUpdate`

NewProductVariantUpdate instantiates a new ProductVariantUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariantUpdateWithDefaults

`func NewProductVariantUpdateWithDefaults() *ProductVariantUpdate`

NewProductVariantUpdateWithDefaults instantiates a new ProductVariantUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductVariantUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductVariantUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductVariantUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductVariantUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *ProductVariantUpdate) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductVariantUpdate) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductVariantUpdate) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductVariantUpdate) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetStoreId

`func (o *ProductVariantUpdate) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductVariantUpdate) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductVariantUpdate) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductVariantUpdate) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetLangId

`func (o *ProductVariantUpdate) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *ProductVariantUpdate) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *ProductVariantUpdate) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *ProductVariantUpdate) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetOptions

`func (o *ProductVariantUpdate) GetOptions() []ProductVariantUpdateOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ProductVariantUpdate) GetOptionsOk() (*[]ProductVariantUpdateOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ProductVariantUpdate) SetOptions(v []ProductVariantUpdateOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ProductVariantUpdate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetName

`func (o *ProductVariantUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductVariantUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductVariantUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductVariantUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProductVariantUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductVariantUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductVariantUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductVariantUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortDescription

`func (o *ProductVariantUpdate) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ProductVariantUpdate) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ProductVariantUpdate) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ProductVariantUpdate) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetModel

`func (o *ProductVariantUpdate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ProductVariantUpdate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ProductVariantUpdate) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ProductVariantUpdate) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSku

`func (o *ProductVariantUpdate) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductVariantUpdate) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductVariantUpdate) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductVariantUpdate) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVisible

`func (o *ProductVariantUpdate) GetVisible() string`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ProductVariantUpdate) GetVisibleOk() (*string, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ProductVariantUpdate) SetVisible(v string)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ProductVariantUpdate) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetStatus

`func (o *ProductVariantUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductVariantUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductVariantUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductVariantUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBackorderStatus

`func (o *ProductVariantUpdate) GetBackorderStatus() string`

GetBackorderStatus returns the BackorderStatus field if non-nil, zero value otherwise.

### GetBackorderStatusOk

`func (o *ProductVariantUpdate) GetBackorderStatusOk() (*string, bool)`

GetBackorderStatusOk returns a tuple with the BackorderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderStatus

`func (o *ProductVariantUpdate) SetBackorderStatus(v string)`

SetBackorderStatus sets BackorderStatus field to given value.

### HasBackorderStatus

`func (o *ProductVariantUpdate) HasBackorderStatus() bool`

HasBackorderStatus returns a boolean if a field has been set.

### GetAvailableForSale

`func (o *ProductVariantUpdate) GetAvailableForSale() bool`

GetAvailableForSale returns the AvailableForSale field if non-nil, zero value otherwise.

### GetAvailableForSaleOk

`func (o *ProductVariantUpdate) GetAvailableForSaleOk() (*bool, bool)`

GetAvailableForSaleOk returns a tuple with the AvailableForSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableForSale

`func (o *ProductVariantUpdate) SetAvailableForSale(v bool)`

SetAvailableForSale sets AvailableForSale field to given value.

### HasAvailableForSale

`func (o *ProductVariantUpdate) HasAvailableForSale() bool`

HasAvailableForSale returns a boolean if a field has been set.

### GetAvail

`func (o *ProductVariantUpdate) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *ProductVariantUpdate) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *ProductVariantUpdate) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *ProductVariantUpdate) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetIsDefault

`func (o *ProductVariantUpdate) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ProductVariantUpdate) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ProductVariantUpdate) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ProductVariantUpdate) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsFreeShipping

`func (o *ProductVariantUpdate) GetIsFreeShipping() bool`

GetIsFreeShipping returns the IsFreeShipping field if non-nil, zero value otherwise.

### GetIsFreeShippingOk

`func (o *ProductVariantUpdate) GetIsFreeShippingOk() (*bool, bool)`

GetIsFreeShippingOk returns a tuple with the IsFreeShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeShipping

`func (o *ProductVariantUpdate) SetIsFreeShipping(v bool)`

SetIsFreeShipping sets IsFreeShipping field to given value.

### HasIsFreeShipping

`func (o *ProductVariantUpdate) HasIsFreeShipping() bool`

HasIsFreeShipping returns a boolean if a field has been set.

### GetTaxable

`func (o *ProductVariantUpdate) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *ProductVariantUpdate) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *ProductVariantUpdate) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *ProductVariantUpdate) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetTaxClassId

`func (o *ProductVariantUpdate) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *ProductVariantUpdate) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *ProductVariantUpdate) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *ProductVariantUpdate) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### GetIsVirtual

`func (o *ProductVariantUpdate) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *ProductVariantUpdate) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *ProductVariantUpdate) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *ProductVariantUpdate) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetManageStock

`func (o *ProductVariantUpdate) GetManageStock() bool`

GetManageStock returns the ManageStock field if non-nil, zero value otherwise.

### GetManageStockOk

`func (o *ProductVariantUpdate) GetManageStockOk() (*bool, bool)`

GetManageStockOk returns a tuple with the ManageStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStock

`func (o *ProductVariantUpdate) SetManageStock(v bool)`

SetManageStock sets ManageStock field to given value.

### HasManageStock

`func (o *ProductVariantUpdate) HasManageStock() bool`

HasManageStock returns a boolean if a field has been set.

### GetInStock

`func (o *ProductVariantUpdate) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *ProductVariantUpdate) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *ProductVariantUpdate) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *ProductVariantUpdate) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### GetWarehouseId

`func (o *ProductVariantUpdate) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ProductVariantUpdate) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ProductVariantUpdate) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ProductVariantUpdate) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetReserveQuantity

`func (o *ProductVariantUpdate) GetReserveQuantity() float32`

GetReserveQuantity returns the ReserveQuantity field if non-nil, zero value otherwise.

### GetReserveQuantityOk

`func (o *ProductVariantUpdate) GetReserveQuantityOk() (*float32, bool)`

GetReserveQuantityOk returns a tuple with the ReserveQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveQuantity

`func (o *ProductVariantUpdate) SetReserveQuantity(v float32)`

SetReserveQuantity sets ReserveQuantity field to given value.

### HasReserveQuantity

`func (o *ProductVariantUpdate) HasReserveQuantity() bool`

HasReserveQuantity returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductVariantUpdate) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductVariantUpdate) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductVariantUpdate) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductVariantUpdate) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetIncreaseQuantity

`func (o *ProductVariantUpdate) GetIncreaseQuantity() float32`

GetIncreaseQuantity returns the IncreaseQuantity field if non-nil, zero value otherwise.

### GetIncreaseQuantityOk

`func (o *ProductVariantUpdate) GetIncreaseQuantityOk() (*float32, bool)`

GetIncreaseQuantityOk returns a tuple with the IncreaseQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseQuantity

`func (o *ProductVariantUpdate) SetIncreaseQuantity(v float32)`

SetIncreaseQuantity sets IncreaseQuantity field to given value.

### HasIncreaseQuantity

`func (o *ProductVariantUpdate) HasIncreaseQuantity() bool`

HasIncreaseQuantity returns a boolean if a field has been set.

### GetReduceQuantity

`func (o *ProductVariantUpdate) GetReduceQuantity() float32`

GetReduceQuantity returns the ReduceQuantity field if non-nil, zero value otherwise.

### GetReduceQuantityOk

`func (o *ProductVariantUpdate) GetReduceQuantityOk() (*float32, bool)`

GetReduceQuantityOk returns a tuple with the ReduceQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceQuantity

`func (o *ProductVariantUpdate) SetReduceQuantity(v float32)`

SetReduceQuantity sets ReduceQuantity field to given value.

### HasReduceQuantity

`func (o *ProductVariantUpdate) HasReduceQuantity() bool`

HasReduceQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *ProductVariantUpdate) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductVariantUpdate) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductVariantUpdate) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductVariantUpdate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSpecialPrice

`func (o *ProductVariantUpdate) GetSpecialPrice() float32`

GetSpecialPrice returns the SpecialPrice field if non-nil, zero value otherwise.

### GetSpecialPriceOk

`func (o *ProductVariantUpdate) GetSpecialPriceOk() (*float32, bool)`

GetSpecialPriceOk returns a tuple with the SpecialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPrice

`func (o *ProductVariantUpdate) SetSpecialPrice(v float32)`

SetSpecialPrice sets SpecialPrice field to given value.

### HasSpecialPrice

`func (o *ProductVariantUpdate) HasSpecialPrice() bool`

HasSpecialPrice returns a boolean if a field has been set.

### GetRetailPrice

`func (o *ProductVariantUpdate) GetRetailPrice() float32`

GetRetailPrice returns the RetailPrice field if non-nil, zero value otherwise.

### GetRetailPriceOk

`func (o *ProductVariantUpdate) GetRetailPriceOk() (*float32, bool)`

GetRetailPriceOk returns a tuple with the RetailPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailPrice

`func (o *ProductVariantUpdate) SetRetailPrice(v float32)`

SetRetailPrice sets RetailPrice field to given value.

### HasRetailPrice

`func (o *ProductVariantUpdate) HasRetailPrice() bool`

HasRetailPrice returns a boolean if a field has been set.

### GetOldPrice

`func (o *ProductVariantUpdate) GetOldPrice() float32`

GetOldPrice returns the OldPrice field if non-nil, zero value otherwise.

### GetOldPriceOk

`func (o *ProductVariantUpdate) GetOldPriceOk() (*float32, bool)`

GetOldPriceOk returns a tuple with the OldPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPrice

`func (o *ProductVariantUpdate) SetOldPrice(v float32)`

SetOldPrice sets OldPrice field to given value.

### HasOldPrice

`func (o *ProductVariantUpdate) HasOldPrice() bool`

HasOldPrice returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductVariantUpdate) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductVariantUpdate) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductVariantUpdate) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductVariantUpdate) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetFixedCostShippingPrice

`func (o *ProductVariantUpdate) GetFixedCostShippingPrice() float32`

GetFixedCostShippingPrice returns the FixedCostShippingPrice field if non-nil, zero value otherwise.

### GetFixedCostShippingPriceOk

`func (o *ProductVariantUpdate) GetFixedCostShippingPriceOk() (*float32, bool)`

GetFixedCostShippingPriceOk returns a tuple with the FixedCostShippingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCostShippingPrice

`func (o *ProductVariantUpdate) SetFixedCostShippingPrice(v float32)`

SetFixedCostShippingPrice sets FixedCostShippingPrice field to given value.

### HasFixedCostShippingPrice

`func (o *ProductVariantUpdate) HasFixedCostShippingPrice() bool`

HasFixedCostShippingPrice returns a boolean if a field has been set.

### GetSpriceCreate

`func (o *ProductVariantUpdate) GetSpriceCreate() string`

GetSpriceCreate returns the SpriceCreate field if non-nil, zero value otherwise.

### GetSpriceCreateOk

`func (o *ProductVariantUpdate) GetSpriceCreateOk() (*string, bool)`

GetSpriceCreateOk returns a tuple with the SpriceCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceCreate

`func (o *ProductVariantUpdate) SetSpriceCreate(v string)`

SetSpriceCreate sets SpriceCreate field to given value.

### HasSpriceCreate

`func (o *ProductVariantUpdate) HasSpriceCreate() bool`

HasSpriceCreate returns a boolean if a field has been set.

### GetSpriceExpire

`func (o *ProductVariantUpdate) GetSpriceExpire() string`

GetSpriceExpire returns the SpriceExpire field if non-nil, zero value otherwise.

### GetSpriceExpireOk

`func (o *ProductVariantUpdate) GetSpriceExpireOk() (*string, bool)`

GetSpriceExpireOk returns a tuple with the SpriceExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpriceExpire

`func (o *ProductVariantUpdate) SetSpriceExpire(v string)`

SetSpriceExpire sets SpriceExpire field to given value.

### HasSpriceExpire

`func (o *ProductVariantUpdate) HasSpriceExpire() bool`

HasSpriceExpire returns a boolean if a field has been set.

### GetWeight

`func (o *ProductVariantUpdate) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductVariantUpdate) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductVariantUpdate) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductVariantUpdate) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetBarcode

`func (o *ProductVariantUpdate) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductVariantUpdate) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductVariantUpdate) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *ProductVariantUpdate) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetWidth

`func (o *ProductVariantUpdate) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ProductVariantUpdate) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ProductVariantUpdate) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ProductVariantUpdate) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetWeightUnit

`func (o *ProductVariantUpdate) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ProductVariantUpdate) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ProductVariantUpdate) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ProductVariantUpdate) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetHeight

`func (o *ProductVariantUpdate) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ProductVariantUpdate) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ProductVariantUpdate) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ProductVariantUpdate) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetLength

`func (o *ProductVariantUpdate) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ProductVariantUpdate) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ProductVariantUpdate) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ProductVariantUpdate) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetGtin

`func (o *ProductVariantUpdate) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *ProductVariantUpdate) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *ProductVariantUpdate) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *ProductVariantUpdate) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### GetUpc

`func (o *ProductVariantUpdate) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ProductVariantUpdate) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ProductVariantUpdate) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ProductVariantUpdate) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetMpn

`func (o *ProductVariantUpdate) GetMpn() string`

GetMpn returns the Mpn field if non-nil, zero value otherwise.

### GetMpnOk

`func (o *ProductVariantUpdate) GetMpnOk() (*string, bool)`

GetMpnOk returns a tuple with the Mpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpn

`func (o *ProductVariantUpdate) SetMpn(v string)`

SetMpn sets Mpn field to given value.

### HasMpn

`func (o *ProductVariantUpdate) HasMpn() bool`

HasMpn returns a boolean if a field has been set.

### GetEan

`func (o *ProductVariantUpdate) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *ProductVariantUpdate) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *ProductVariantUpdate) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *ProductVariantUpdate) HasEan() bool`

HasEan returns a boolean if a field has been set.

### GetIsbn

`func (o *ProductVariantUpdate) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *ProductVariantUpdate) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *ProductVariantUpdate) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *ProductVariantUpdate) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### GetHarmonizedSystemCode

`func (o *ProductVariantUpdate) GetHarmonizedSystemCode() string`

GetHarmonizedSystemCode returns the HarmonizedSystemCode field if non-nil, zero value otherwise.

### GetHarmonizedSystemCodeOk

`func (o *ProductVariantUpdate) GetHarmonizedSystemCodeOk() (*string, bool)`

GetHarmonizedSystemCodeOk returns a tuple with the HarmonizedSystemCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarmonizedSystemCode

`func (o *ProductVariantUpdate) SetHarmonizedSystemCode(v string)`

SetHarmonizedSystemCode sets HarmonizedSystemCode field to given value.

### HasHarmonizedSystemCode

`func (o *ProductVariantUpdate) HasHarmonizedSystemCode() bool`

HasHarmonizedSystemCode returns a boolean if a field has been set.

### GetCountryOfOrigin

`func (o *ProductVariantUpdate) GetCountryOfOrigin() string`

GetCountryOfOrigin returns the CountryOfOrigin field if non-nil, zero value otherwise.

### GetCountryOfOriginOk

`func (o *ProductVariantUpdate) GetCountryOfOriginOk() (*string, bool)`

GetCountryOfOriginOk returns a tuple with the CountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfOrigin

`func (o *ProductVariantUpdate) SetCountryOfOrigin(v string)`

SetCountryOfOrigin sets CountryOfOrigin field to given value.

### HasCountryOfOrigin

`func (o *ProductVariantUpdate) HasCountryOfOrigin() bool`

HasCountryOfOrigin returns a boolean if a field has been set.

### GetMetaTitle

`func (o *ProductVariantUpdate) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *ProductVariantUpdate) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *ProductVariantUpdate) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *ProductVariantUpdate) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetMetaDescription

`func (o *ProductVariantUpdate) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *ProductVariantUpdate) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *ProductVariantUpdate) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *ProductVariantUpdate) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *ProductVariantUpdate) GetMetaKeywords() string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *ProductVariantUpdate) GetMetaKeywordsOk() (*string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *ProductVariantUpdate) SetMetaKeywords(v string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *ProductVariantUpdate) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetReindex

`func (o *ProductVariantUpdate) GetReindex() bool`

GetReindex returns the Reindex field if non-nil, zero value otherwise.

### GetReindexOk

`func (o *ProductVariantUpdate) GetReindexOk() (*bool, bool)`

GetReindexOk returns a tuple with the Reindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReindex

`func (o *ProductVariantUpdate) SetReindex(v bool)`

SetReindex sets Reindex field to given value.

### HasReindex

`func (o *ProductVariantUpdate) HasReindex() bool`

HasReindex returns a boolean if a field has been set.

### GetClearCache

`func (o *ProductVariantUpdate) GetClearCache() bool`

GetClearCache returns the ClearCache field if non-nil, zero value otherwise.

### GetClearCacheOk

`func (o *ProductVariantUpdate) GetClearCacheOk() (*bool, bool)`

GetClearCacheOk returns a tuple with the ClearCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearCache

`func (o *ProductVariantUpdate) SetClearCache(v bool)`

SetClearCache sets ClearCache field to given value.

### HasClearCache

`func (o *ProductVariantUpdate) HasClearCache() bool`

HasClearCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


