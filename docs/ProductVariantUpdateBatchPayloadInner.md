# ProductVariantUpdateBatchPayloadInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ProductId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**Upc** | Pointer to **string** |  | [optional] 
**Mpn** | Pointer to **string** |  | [optional] 
**Gtin** | Pointer to **string** |  | [optional] 
**Isbn** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**SpecialPrice** | Pointer to **float32** |  | [optional] 
**CostPrice** | Pointer to **float32** |  | [optional] 
**RetailPrice** | Pointer to **float32** |  | [optional] 
**AdvancedPrices** | Pointer to [**[]ProductUpdateBatchPayloadInnerAdvancedPricesInner**](ProductUpdateBatchPayloadInnerAdvancedPricesInner.md) | If an empty array is passed, all entries will be deleted when the &#39;nested_items_update_behaviour&#39; parameter is set to &#39;replace&#39;. | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**ReserveQuantity** | Pointer to **float32** |  | [optional] 
**IncreaseQuantity** | Pointer to **float32** |  | [optional] 
**ReduceQuantity** | Pointer to **float32** |  | [optional] 
**WarehouseId** | Pointer to **string** |  | [optional] 
**ManufacturerId** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] 
**Height** | Pointer to **float32** |  | [optional] 
**Length** | Pointer to **float32** |  | [optional] 
**Width** | Pointer to **float32** |  | [optional] 
**StoreId** | Pointer to **string** |  | [optional] 
**LangId** | Pointer to **string** |  | [optional] 
**TaxClassId** | Pointer to **string** |  | [optional] 
**BackorderStatus** | Pointer to **string** |  | [optional] 
**Visible** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**InStock** | Pointer to **bool** |  | [optional] 
**IsVirtual** | Pointer to **bool** |  | [optional] 
**Downloadable** | Pointer to **bool** |  | [optional] 
**ManageStock** | Pointer to **bool** |  | [optional] 
**IsFreeShipping** | Pointer to **bool** |  | [optional] 
**SeoUrl** | Pointer to **string** |  | [optional] 
**MetaTitle** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**MetaKeywords** | Pointer to **[]string** |  | [optional] 
**CategoriesIds** | Pointer to **[]string** | If an empty array is passed, all entries will be deleted when the &#39;nested_items_update_behaviour&#39; parameter is set to &#39;replace&#39;. | [optional] 
**StoresIds** | Pointer to **[]string** |  | [optional] 
**Images** | Pointer to [**[]ProductAddBatchPayloadInnerImagesInner**](ProductAddBatchPayloadInnerImagesInner.md) | The passed items will completely replace the original items | [optional] 
**ProductImagesIds** | Pointer to **[]string** |  | [optional] 
**RelatedProductsIds** | Pointer to **[]string** | If an empty array is passed, all entries will be deleted when the &#39;nested_items_update_behaviour&#39; parameter is set to &#39;replace&#39;. | [optional] 
**UpSellProductsIds** | Pointer to **[]string** | If an empty array is passed, all entries will be deleted when the &#39;nested_items_update_behaviour&#39; parameter is set to &#39;replace&#39;. | [optional] 
**CrossSellProductsIds** | Pointer to **[]string** | If an empty array is passed, all entries will be deleted when the &#39;nested_items_update_behaviour&#39; parameter is set to &#39;replace&#39;. | [optional] 

## Methods

### NewProductVariantUpdateBatchPayloadInner

`func NewProductVariantUpdateBatchPayloadInner(id string, productId string, ) *ProductVariantUpdateBatchPayloadInner`

NewProductVariantUpdateBatchPayloadInner instantiates a new ProductVariantUpdateBatchPayloadInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariantUpdateBatchPayloadInnerWithDefaults

`func NewProductVariantUpdateBatchPayloadInnerWithDefaults() *ProductVariantUpdateBatchPayloadInner`

NewProductVariantUpdateBatchPayloadInnerWithDefaults instantiates a new ProductVariantUpdateBatchPayloadInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductVariantUpdateBatchPayloadInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductVariantUpdateBatchPayloadInner) SetId(v string)`

SetId sets Id field to given value.


### GetProductId

`func (o *ProductVariantUpdateBatchPayloadInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductVariantUpdateBatchPayloadInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetName

`func (o *ProductVariantUpdateBatchPayloadInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductVariantUpdateBatchPayloadInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductVariantUpdateBatchPayloadInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProductVariantUpdateBatchPayloadInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductVariantUpdateBatchPayloadInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductVariantUpdateBatchPayloadInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortDescription

`func (o *ProductVariantUpdateBatchPayloadInner) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ProductVariantUpdateBatchPayloadInner) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ProductVariantUpdateBatchPayloadInner) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetSku

`func (o *ProductVariantUpdateBatchPayloadInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductVariantUpdateBatchPayloadInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductVariantUpdateBatchPayloadInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUpc

`func (o *ProductVariantUpdateBatchPayloadInner) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ProductVariantUpdateBatchPayloadInner) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ProductVariantUpdateBatchPayloadInner) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetMpn

`func (o *ProductVariantUpdateBatchPayloadInner) GetMpn() string`

GetMpn returns the Mpn field if non-nil, zero value otherwise.

### GetMpnOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetMpnOk() (*string, bool)`

GetMpnOk returns a tuple with the Mpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpn

`func (o *ProductVariantUpdateBatchPayloadInner) SetMpn(v string)`

SetMpn sets Mpn field to given value.

### HasMpn

`func (o *ProductVariantUpdateBatchPayloadInner) HasMpn() bool`

HasMpn returns a boolean if a field has been set.

### GetGtin

`func (o *ProductVariantUpdateBatchPayloadInner) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *ProductVariantUpdateBatchPayloadInner) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *ProductVariantUpdateBatchPayloadInner) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### GetIsbn

`func (o *ProductVariantUpdateBatchPayloadInner) GetIsbn() string`

GetIsbn returns the Isbn field if non-nil, zero value otherwise.

### GetIsbnOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetIsbnOk() (*string, bool)`

GetIsbnOk returns a tuple with the Isbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn

`func (o *ProductVariantUpdateBatchPayloadInner) SetIsbn(v string)`

SetIsbn sets Isbn field to given value.

### HasIsbn

`func (o *ProductVariantUpdateBatchPayloadInner) HasIsbn() bool`

HasIsbn returns a boolean if a field has been set.

### GetStatus

`func (o *ProductVariantUpdateBatchPayloadInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductVariantUpdateBatchPayloadInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductVariantUpdateBatchPayloadInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrice

`func (o *ProductVariantUpdateBatchPayloadInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductVariantUpdateBatchPayloadInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductVariantUpdateBatchPayloadInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSpecialPrice

`func (o *ProductVariantUpdateBatchPayloadInner) GetSpecialPrice() float32`

GetSpecialPrice returns the SpecialPrice field if non-nil, zero value otherwise.

### GetSpecialPriceOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetSpecialPriceOk() (*float32, bool)`

GetSpecialPriceOk returns a tuple with the SpecialPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPrice

`func (o *ProductVariantUpdateBatchPayloadInner) SetSpecialPrice(v float32)`

SetSpecialPrice sets SpecialPrice field to given value.

### HasSpecialPrice

`func (o *ProductVariantUpdateBatchPayloadInner) HasSpecialPrice() bool`

HasSpecialPrice returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductVariantUpdateBatchPayloadInner) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductVariantUpdateBatchPayloadInner) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductVariantUpdateBatchPayloadInner) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetRetailPrice

`func (o *ProductVariantUpdateBatchPayloadInner) GetRetailPrice() float32`

GetRetailPrice returns the RetailPrice field if non-nil, zero value otherwise.

### GetRetailPriceOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetRetailPriceOk() (*float32, bool)`

GetRetailPriceOk returns a tuple with the RetailPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailPrice

`func (o *ProductVariantUpdateBatchPayloadInner) SetRetailPrice(v float32)`

SetRetailPrice sets RetailPrice field to given value.

### HasRetailPrice

`func (o *ProductVariantUpdateBatchPayloadInner) HasRetailPrice() bool`

HasRetailPrice returns a boolean if a field has been set.

### GetAdvancedPrices

`func (o *ProductVariantUpdateBatchPayloadInner) GetAdvancedPrices() []ProductUpdateBatchPayloadInnerAdvancedPricesInner`

GetAdvancedPrices returns the AdvancedPrices field if non-nil, zero value otherwise.

### GetAdvancedPricesOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetAdvancedPricesOk() (*[]ProductUpdateBatchPayloadInnerAdvancedPricesInner, bool)`

GetAdvancedPricesOk returns a tuple with the AdvancedPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedPrices

`func (o *ProductVariantUpdateBatchPayloadInner) SetAdvancedPrices(v []ProductUpdateBatchPayloadInnerAdvancedPricesInner)`

SetAdvancedPrices sets AdvancedPrices field to given value.

### HasAdvancedPrices

`func (o *ProductVariantUpdateBatchPayloadInner) HasAdvancedPrices() bool`

HasAdvancedPrices returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReserveQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) GetReserveQuantity() float32`

GetReserveQuantity returns the ReserveQuantity field if non-nil, zero value otherwise.

### GetReserveQuantityOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetReserveQuantityOk() (*float32, bool)`

GetReserveQuantityOk returns a tuple with the ReserveQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) SetReserveQuantity(v float32)`

SetReserveQuantity sets ReserveQuantity field to given value.

### HasReserveQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) HasReserveQuantity() bool`

HasReserveQuantity returns a boolean if a field has been set.

### GetIncreaseQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) GetIncreaseQuantity() float32`

GetIncreaseQuantity returns the IncreaseQuantity field if non-nil, zero value otherwise.

### GetIncreaseQuantityOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetIncreaseQuantityOk() (*float32, bool)`

GetIncreaseQuantityOk returns a tuple with the IncreaseQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncreaseQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) SetIncreaseQuantity(v float32)`

SetIncreaseQuantity sets IncreaseQuantity field to given value.

### HasIncreaseQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) HasIncreaseQuantity() bool`

HasIncreaseQuantity returns a boolean if a field has been set.

### GetReduceQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) GetReduceQuantity() float32`

GetReduceQuantity returns the ReduceQuantity field if non-nil, zero value otherwise.

### GetReduceQuantityOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetReduceQuantityOk() (*float32, bool)`

GetReduceQuantityOk returns a tuple with the ReduceQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) SetReduceQuantity(v float32)`

SetReduceQuantity sets ReduceQuantity field to given value.

### HasReduceQuantity

`func (o *ProductVariantUpdateBatchPayloadInner) HasReduceQuantity() bool`

HasReduceQuantity returns a boolean if a field has been set.

### GetWarehouseId

`func (o *ProductVariantUpdateBatchPayloadInner) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ProductVariantUpdateBatchPayloadInner) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ProductVariantUpdateBatchPayloadInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetManufacturerId

`func (o *ProductVariantUpdateBatchPayloadInner) GetManufacturerId() string`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetManufacturerIdOk() (*string, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *ProductVariantUpdateBatchPayloadInner) SetManufacturerId(v string)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *ProductVariantUpdateBatchPayloadInner) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetWeight

`func (o *ProductVariantUpdateBatchPayloadInner) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductVariantUpdateBatchPayloadInner) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductVariantUpdateBatchPayloadInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetHeight

`func (o *ProductVariantUpdateBatchPayloadInner) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ProductVariantUpdateBatchPayloadInner) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ProductVariantUpdateBatchPayloadInner) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetLength

`func (o *ProductVariantUpdateBatchPayloadInner) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ProductVariantUpdateBatchPayloadInner) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ProductVariantUpdateBatchPayloadInner) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetWidth

`func (o *ProductVariantUpdateBatchPayloadInner) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ProductVariantUpdateBatchPayloadInner) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ProductVariantUpdateBatchPayloadInner) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetStoreId

`func (o *ProductVariantUpdateBatchPayloadInner) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductVariantUpdateBatchPayloadInner) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductVariantUpdateBatchPayloadInner) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetLangId

`func (o *ProductVariantUpdateBatchPayloadInner) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *ProductVariantUpdateBatchPayloadInner) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *ProductVariantUpdateBatchPayloadInner) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetTaxClassId

`func (o *ProductVariantUpdateBatchPayloadInner) GetTaxClassId() string`

GetTaxClassId returns the TaxClassId field if non-nil, zero value otherwise.

### GetTaxClassIdOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetTaxClassIdOk() (*string, bool)`

GetTaxClassIdOk returns a tuple with the TaxClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxClassId

`func (o *ProductVariantUpdateBatchPayloadInner) SetTaxClassId(v string)`

SetTaxClassId sets TaxClassId field to given value.

### HasTaxClassId

`func (o *ProductVariantUpdateBatchPayloadInner) HasTaxClassId() bool`

HasTaxClassId returns a boolean if a field has been set.

### GetBackorderStatus

`func (o *ProductVariantUpdateBatchPayloadInner) GetBackorderStatus() string`

GetBackorderStatus returns the BackorderStatus field if non-nil, zero value otherwise.

### GetBackorderStatusOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetBackorderStatusOk() (*string, bool)`

GetBackorderStatusOk returns a tuple with the BackorderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderStatus

`func (o *ProductVariantUpdateBatchPayloadInner) SetBackorderStatus(v string)`

SetBackorderStatus sets BackorderStatus field to given value.

### HasBackorderStatus

`func (o *ProductVariantUpdateBatchPayloadInner) HasBackorderStatus() bool`

HasBackorderStatus returns a boolean if a field has been set.

### GetVisible

`func (o *ProductVariantUpdateBatchPayloadInner) GetVisible() string`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetVisibleOk() (*string, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ProductVariantUpdateBatchPayloadInner) SetVisible(v string)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ProductVariantUpdateBatchPayloadInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetIsDefault

`func (o *ProductVariantUpdateBatchPayloadInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ProductVariantUpdateBatchPayloadInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ProductVariantUpdateBatchPayloadInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetInStock

`func (o *ProductVariantUpdateBatchPayloadInner) GetInStock() bool`

GetInStock returns the InStock field if non-nil, zero value otherwise.

### GetInStockOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetInStockOk() (*bool, bool)`

GetInStockOk returns a tuple with the InStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStock

`func (o *ProductVariantUpdateBatchPayloadInner) SetInStock(v bool)`

SetInStock sets InStock field to given value.

### HasInStock

`func (o *ProductVariantUpdateBatchPayloadInner) HasInStock() bool`

HasInStock returns a boolean if a field has been set.

### GetIsVirtual

`func (o *ProductVariantUpdateBatchPayloadInner) GetIsVirtual() bool`

GetIsVirtual returns the IsVirtual field if non-nil, zero value otherwise.

### GetIsVirtualOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetIsVirtualOk() (*bool, bool)`

GetIsVirtualOk returns a tuple with the IsVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtual

`func (o *ProductVariantUpdateBatchPayloadInner) SetIsVirtual(v bool)`

SetIsVirtual sets IsVirtual field to given value.

### HasIsVirtual

`func (o *ProductVariantUpdateBatchPayloadInner) HasIsVirtual() bool`

HasIsVirtual returns a boolean if a field has been set.

### GetDownloadable

`func (o *ProductVariantUpdateBatchPayloadInner) GetDownloadable() bool`

GetDownloadable returns the Downloadable field if non-nil, zero value otherwise.

### GetDownloadableOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetDownloadableOk() (*bool, bool)`

GetDownloadableOk returns a tuple with the Downloadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadable

`func (o *ProductVariantUpdateBatchPayloadInner) SetDownloadable(v bool)`

SetDownloadable sets Downloadable field to given value.

### HasDownloadable

`func (o *ProductVariantUpdateBatchPayloadInner) HasDownloadable() bool`

HasDownloadable returns a boolean if a field has been set.

### GetManageStock

`func (o *ProductVariantUpdateBatchPayloadInner) GetManageStock() bool`

GetManageStock returns the ManageStock field if non-nil, zero value otherwise.

### GetManageStockOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetManageStockOk() (*bool, bool)`

GetManageStockOk returns a tuple with the ManageStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStock

`func (o *ProductVariantUpdateBatchPayloadInner) SetManageStock(v bool)`

SetManageStock sets ManageStock field to given value.

### HasManageStock

`func (o *ProductVariantUpdateBatchPayloadInner) HasManageStock() bool`

HasManageStock returns a boolean if a field has been set.

### GetIsFreeShipping

`func (o *ProductVariantUpdateBatchPayloadInner) GetIsFreeShipping() bool`

GetIsFreeShipping returns the IsFreeShipping field if non-nil, zero value otherwise.

### GetIsFreeShippingOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetIsFreeShippingOk() (*bool, bool)`

GetIsFreeShippingOk returns a tuple with the IsFreeShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFreeShipping

`func (o *ProductVariantUpdateBatchPayloadInner) SetIsFreeShipping(v bool)`

SetIsFreeShipping sets IsFreeShipping field to given value.

### HasIsFreeShipping

`func (o *ProductVariantUpdateBatchPayloadInner) HasIsFreeShipping() bool`

HasIsFreeShipping returns a boolean if a field has been set.

### GetSeoUrl

`func (o *ProductVariantUpdateBatchPayloadInner) GetSeoUrl() string`

GetSeoUrl returns the SeoUrl field if non-nil, zero value otherwise.

### GetSeoUrlOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetSeoUrlOk() (*string, bool)`

GetSeoUrlOk returns a tuple with the SeoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoUrl

`func (o *ProductVariantUpdateBatchPayloadInner) SetSeoUrl(v string)`

SetSeoUrl sets SeoUrl field to given value.

### HasSeoUrl

`func (o *ProductVariantUpdateBatchPayloadInner) HasSeoUrl() bool`

HasSeoUrl returns a boolean if a field has been set.

### GetMetaTitle

`func (o *ProductVariantUpdateBatchPayloadInner) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *ProductVariantUpdateBatchPayloadInner) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *ProductVariantUpdateBatchPayloadInner) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetMetaDescription

`func (o *ProductVariantUpdateBatchPayloadInner) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *ProductVariantUpdateBatchPayloadInner) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *ProductVariantUpdateBatchPayloadInner) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *ProductVariantUpdateBatchPayloadInner) GetMetaKeywords() []string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetMetaKeywordsOk() (*[]string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *ProductVariantUpdateBatchPayloadInner) SetMetaKeywords(v []string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *ProductVariantUpdateBatchPayloadInner) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetCategoriesIds

`func (o *ProductVariantUpdateBatchPayloadInner) GetCategoriesIds() []string`

GetCategoriesIds returns the CategoriesIds field if non-nil, zero value otherwise.

### GetCategoriesIdsOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetCategoriesIdsOk() (*[]string, bool)`

GetCategoriesIdsOk returns a tuple with the CategoriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesIds

`func (o *ProductVariantUpdateBatchPayloadInner) SetCategoriesIds(v []string)`

SetCategoriesIds sets CategoriesIds field to given value.

### HasCategoriesIds

`func (o *ProductVariantUpdateBatchPayloadInner) HasCategoriesIds() bool`

HasCategoriesIds returns a boolean if a field has been set.

### GetStoresIds

`func (o *ProductVariantUpdateBatchPayloadInner) GetStoresIds() []string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetStoresIdsOk() (*[]string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *ProductVariantUpdateBatchPayloadInner) SetStoresIds(v []string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *ProductVariantUpdateBatchPayloadInner) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetImages

`func (o *ProductVariantUpdateBatchPayloadInner) GetImages() []ProductAddBatchPayloadInnerImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetImagesOk() (*[]ProductAddBatchPayloadInnerImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ProductVariantUpdateBatchPayloadInner) SetImages(v []ProductAddBatchPayloadInnerImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *ProductVariantUpdateBatchPayloadInner) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetProductImagesIds

`func (o *ProductVariantUpdateBatchPayloadInner) GetProductImagesIds() []string`

GetProductImagesIds returns the ProductImagesIds field if non-nil, zero value otherwise.

### GetProductImagesIdsOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetProductImagesIdsOk() (*[]string, bool)`

GetProductImagesIdsOk returns a tuple with the ProductImagesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImagesIds

`func (o *ProductVariantUpdateBatchPayloadInner) SetProductImagesIds(v []string)`

SetProductImagesIds sets ProductImagesIds field to given value.

### HasProductImagesIds

`func (o *ProductVariantUpdateBatchPayloadInner) HasProductImagesIds() bool`

HasProductImagesIds returns a boolean if a field has been set.

### GetRelatedProductsIds

`func (o *ProductVariantUpdateBatchPayloadInner) GetRelatedProductsIds() []string`

GetRelatedProductsIds returns the RelatedProductsIds field if non-nil, zero value otherwise.

### GetRelatedProductsIdsOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetRelatedProductsIdsOk() (*[]string, bool)`

GetRelatedProductsIdsOk returns a tuple with the RelatedProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedProductsIds

`func (o *ProductVariantUpdateBatchPayloadInner) SetRelatedProductsIds(v []string)`

SetRelatedProductsIds sets RelatedProductsIds field to given value.

### HasRelatedProductsIds

`func (o *ProductVariantUpdateBatchPayloadInner) HasRelatedProductsIds() bool`

HasRelatedProductsIds returns a boolean if a field has been set.

### GetUpSellProductsIds

`func (o *ProductVariantUpdateBatchPayloadInner) GetUpSellProductsIds() []string`

GetUpSellProductsIds returns the UpSellProductsIds field if non-nil, zero value otherwise.

### GetUpSellProductsIdsOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetUpSellProductsIdsOk() (*[]string, bool)`

GetUpSellProductsIdsOk returns a tuple with the UpSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSellProductsIds

`func (o *ProductVariantUpdateBatchPayloadInner) SetUpSellProductsIds(v []string)`

SetUpSellProductsIds sets UpSellProductsIds field to given value.

### HasUpSellProductsIds

`func (o *ProductVariantUpdateBatchPayloadInner) HasUpSellProductsIds() bool`

HasUpSellProductsIds returns a boolean if a field has been set.

### GetCrossSellProductsIds

`func (o *ProductVariantUpdateBatchPayloadInner) GetCrossSellProductsIds() []string`

GetCrossSellProductsIds returns the CrossSellProductsIds field if non-nil, zero value otherwise.

### GetCrossSellProductsIdsOk

`func (o *ProductVariantUpdateBatchPayloadInner) GetCrossSellProductsIdsOk() (*[]string, bool)`

GetCrossSellProductsIdsOk returns a tuple with the CrossSellProductsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSellProductsIds

`func (o *ProductVariantUpdateBatchPayloadInner) SetCrossSellProductsIds(v []string)`

SetCrossSellProductsIds sets CrossSellProductsIds field to given value.

### HasCrossSellProductsIds

`func (o *ProductVariantUpdateBatchPayloadInner) HasCrossSellProductsIds() bool`

HasCrossSellProductsIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


