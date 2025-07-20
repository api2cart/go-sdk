# MarketplaceProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**UAsin** | Pointer to **NullableString** |  | [optional] 
**UEan** | Pointer to **NullableString** |  | [optional] 
**UGtin** | Pointer to **NullableString** |  | [optional] 
**UIsbn** | Pointer to **NullableString** |  | [optional] 
**UMpn** | Pointer to **NullableString** |  | [optional] 
**UUpc** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) |  | [optional] 
**ProductOptions** | Pointer to [**[]ProductOption**](ProductOption.md) |  | [optional] 
**Manufacturer** | Pointer to **NullableString** |  | [optional] 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **NullableFloat32** |  | [optional] 
**WeightUnit** | Pointer to **NullableString** |  | [optional] 
**DimensionsUnit** | Pointer to **NullableString** |  | [optional] 
**Width** | Pointer to **NullableFloat32** |  | [optional] 
**Height** | Pointer to **NullableFloat32** |  | [optional] 
**Length** | Pointer to **NullableFloat32** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMarketplaceProduct

`func NewMarketplaceProduct() *MarketplaceProduct`

NewMarketplaceProduct instantiates a new MarketplaceProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceProductWithDefaults

`func NewMarketplaceProductWithDefaults() *MarketplaceProduct`

NewMarketplaceProductWithDefaults instantiates a new MarketplaceProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketplaceProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MarketplaceProduct) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketplaceProduct) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketplaceProduct) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MarketplaceProduct) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MarketplaceProduct) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MarketplaceProduct) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUAsin

`func (o *MarketplaceProduct) GetUAsin() string`

GetUAsin returns the UAsin field if non-nil, zero value otherwise.

### GetUAsinOk

`func (o *MarketplaceProduct) GetUAsinOk() (*string, bool)`

GetUAsinOk returns a tuple with the UAsin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUAsin

`func (o *MarketplaceProduct) SetUAsin(v string)`

SetUAsin sets UAsin field to given value.

### HasUAsin

`func (o *MarketplaceProduct) HasUAsin() bool`

HasUAsin returns a boolean if a field has been set.

### SetUAsinNil

`func (o *MarketplaceProduct) SetUAsinNil(b bool)`

 SetUAsinNil sets the value for UAsin to be an explicit nil

### UnsetUAsin
`func (o *MarketplaceProduct) UnsetUAsin()`

UnsetUAsin ensures that no value is present for UAsin, not even an explicit nil
### GetUEan

`func (o *MarketplaceProduct) GetUEan() string`

GetUEan returns the UEan field if non-nil, zero value otherwise.

### GetUEanOk

`func (o *MarketplaceProduct) GetUEanOk() (*string, bool)`

GetUEanOk returns a tuple with the UEan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUEan

`func (o *MarketplaceProduct) SetUEan(v string)`

SetUEan sets UEan field to given value.

### HasUEan

`func (o *MarketplaceProduct) HasUEan() bool`

HasUEan returns a boolean if a field has been set.

### SetUEanNil

`func (o *MarketplaceProduct) SetUEanNil(b bool)`

 SetUEanNil sets the value for UEan to be an explicit nil

### UnsetUEan
`func (o *MarketplaceProduct) UnsetUEan()`

UnsetUEan ensures that no value is present for UEan, not even an explicit nil
### GetUGtin

`func (o *MarketplaceProduct) GetUGtin() string`

GetUGtin returns the UGtin field if non-nil, zero value otherwise.

### GetUGtinOk

`func (o *MarketplaceProduct) GetUGtinOk() (*string, bool)`

GetUGtinOk returns a tuple with the UGtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUGtin

`func (o *MarketplaceProduct) SetUGtin(v string)`

SetUGtin sets UGtin field to given value.

### HasUGtin

`func (o *MarketplaceProduct) HasUGtin() bool`

HasUGtin returns a boolean if a field has been set.

### SetUGtinNil

`func (o *MarketplaceProduct) SetUGtinNil(b bool)`

 SetUGtinNil sets the value for UGtin to be an explicit nil

### UnsetUGtin
`func (o *MarketplaceProduct) UnsetUGtin()`

UnsetUGtin ensures that no value is present for UGtin, not even an explicit nil
### GetUIsbn

`func (o *MarketplaceProduct) GetUIsbn() string`

GetUIsbn returns the UIsbn field if non-nil, zero value otherwise.

### GetUIsbnOk

`func (o *MarketplaceProduct) GetUIsbnOk() (*string, bool)`

GetUIsbnOk returns a tuple with the UIsbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUIsbn

`func (o *MarketplaceProduct) SetUIsbn(v string)`

SetUIsbn sets UIsbn field to given value.

### HasUIsbn

`func (o *MarketplaceProduct) HasUIsbn() bool`

HasUIsbn returns a boolean if a field has been set.

### SetUIsbnNil

`func (o *MarketplaceProduct) SetUIsbnNil(b bool)`

 SetUIsbnNil sets the value for UIsbn to be an explicit nil

### UnsetUIsbn
`func (o *MarketplaceProduct) UnsetUIsbn()`

UnsetUIsbn ensures that no value is present for UIsbn, not even an explicit nil
### GetUMpn

`func (o *MarketplaceProduct) GetUMpn() string`

GetUMpn returns the UMpn field if non-nil, zero value otherwise.

### GetUMpnOk

`func (o *MarketplaceProduct) GetUMpnOk() (*string, bool)`

GetUMpnOk returns a tuple with the UMpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUMpn

`func (o *MarketplaceProduct) SetUMpn(v string)`

SetUMpn sets UMpn field to given value.

### HasUMpn

`func (o *MarketplaceProduct) HasUMpn() bool`

HasUMpn returns a boolean if a field has been set.

### SetUMpnNil

`func (o *MarketplaceProduct) SetUMpnNil(b bool)`

 SetUMpnNil sets the value for UMpn to be an explicit nil

### UnsetUMpn
`func (o *MarketplaceProduct) UnsetUMpn()`

UnsetUMpn ensures that no value is present for UMpn, not even an explicit nil
### GetUUpc

`func (o *MarketplaceProduct) GetUUpc() string`

GetUUpc returns the UUpc field if non-nil, zero value otherwise.

### GetUUpcOk

`func (o *MarketplaceProduct) GetUUpcOk() (*string, bool)`

GetUUpcOk returns a tuple with the UUpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUpc

`func (o *MarketplaceProduct) SetUUpc(v string)`

SetUUpc sets UUpc field to given value.

### HasUUpc

`func (o *MarketplaceProduct) HasUUpc() bool`

HasUUpc returns a boolean if a field has been set.

### SetUUpcNil

`func (o *MarketplaceProduct) SetUUpcNil(b bool)`

 SetUUpcNil sets the value for UUpc to be an explicit nil

### UnsetUUpc
`func (o *MarketplaceProduct) UnsetUUpc()`

UnsetUUpc ensures that no value is present for UUpc, not even an explicit nil
### GetName

`func (o *MarketplaceProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketplaceProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketplaceProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketplaceProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MarketplaceProduct) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MarketplaceProduct) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *MarketplaceProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketplaceProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketplaceProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketplaceProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MarketplaceProduct) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MarketplaceProduct) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *MarketplaceProduct) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MarketplaceProduct) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MarketplaceProduct) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MarketplaceProduct) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MarketplaceProduct) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MarketplaceProduct) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetPrice

`func (o *MarketplaceProduct) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarketplaceProduct) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarketplaceProduct) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarketplaceProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *MarketplaceProduct) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *MarketplaceProduct) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetImages

`func (o *MarketplaceProduct) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *MarketplaceProduct) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *MarketplaceProduct) SetImages(v []Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *MarketplaceProduct) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetProductOptions

`func (o *MarketplaceProduct) GetProductOptions() []ProductOption`

GetProductOptions returns the ProductOptions field if non-nil, zero value otherwise.

### GetProductOptionsOk

`func (o *MarketplaceProduct) GetProductOptionsOk() (*[]ProductOption, bool)`

GetProductOptionsOk returns a tuple with the ProductOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductOptions

`func (o *MarketplaceProduct) SetProductOptions(v []ProductOption)`

SetProductOptions sets ProductOptions field to given value.

### HasProductOptions

`func (o *MarketplaceProduct) HasProductOptions() bool`

HasProductOptions returns a boolean if a field has been set.

### GetManufacturer

`func (o *MarketplaceProduct) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *MarketplaceProduct) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *MarketplaceProduct) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *MarketplaceProduct) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *MarketplaceProduct) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *MarketplaceProduct) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetBrand

`func (o *MarketplaceProduct) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *MarketplaceProduct) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *MarketplaceProduct) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *MarketplaceProduct) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *MarketplaceProduct) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *MarketplaceProduct) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetWeight

`func (o *MarketplaceProduct) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *MarketplaceProduct) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *MarketplaceProduct) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *MarketplaceProduct) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *MarketplaceProduct) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *MarketplaceProduct) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetWeightUnit

`func (o *MarketplaceProduct) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *MarketplaceProduct) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *MarketplaceProduct) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *MarketplaceProduct) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *MarketplaceProduct) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *MarketplaceProduct) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetDimensionsUnit

`func (o *MarketplaceProduct) GetDimensionsUnit() string`

GetDimensionsUnit returns the DimensionsUnit field if non-nil, zero value otherwise.

### GetDimensionsUnitOk

`func (o *MarketplaceProduct) GetDimensionsUnitOk() (*string, bool)`

GetDimensionsUnitOk returns a tuple with the DimensionsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsUnit

`func (o *MarketplaceProduct) SetDimensionsUnit(v string)`

SetDimensionsUnit sets DimensionsUnit field to given value.

### HasDimensionsUnit

`func (o *MarketplaceProduct) HasDimensionsUnit() bool`

HasDimensionsUnit returns a boolean if a field has been set.

### SetDimensionsUnitNil

`func (o *MarketplaceProduct) SetDimensionsUnitNil(b bool)`

 SetDimensionsUnitNil sets the value for DimensionsUnit to be an explicit nil

### UnsetDimensionsUnit
`func (o *MarketplaceProduct) UnsetDimensionsUnit()`

UnsetDimensionsUnit ensures that no value is present for DimensionsUnit, not even an explicit nil
### GetWidth

`func (o *MarketplaceProduct) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MarketplaceProduct) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *MarketplaceProduct) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *MarketplaceProduct) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *MarketplaceProduct) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *MarketplaceProduct) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *MarketplaceProduct) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *MarketplaceProduct) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *MarketplaceProduct) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *MarketplaceProduct) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *MarketplaceProduct) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *MarketplaceProduct) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetLength

`func (o *MarketplaceProduct) GetLength() float32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *MarketplaceProduct) GetLengthOk() (*float32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *MarketplaceProduct) SetLength(v float32)`

SetLength sets Length field to given value.

### HasLength

`func (o *MarketplaceProduct) HasLength() bool`

HasLength returns a boolean if a field has been set.

### SetLengthNil

`func (o *MarketplaceProduct) SetLengthNil(b bool)`

 SetLengthNil sets the value for Length to be an explicit nil

### UnsetLength
`func (o *MarketplaceProduct) UnsetLength()`

UnsetLength ensures that no value is present for Length, not even an explicit nil
### GetAdditionalFields

`func (o *MarketplaceProduct) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *MarketplaceProduct) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *MarketplaceProduct) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *MarketplaceProduct) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *MarketplaceProduct) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *MarketplaceProduct) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *MarketplaceProduct) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MarketplaceProduct) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MarketplaceProduct) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MarketplaceProduct) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *MarketplaceProduct) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *MarketplaceProduct) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


