/*
API2Cart OpenAPI

API2Cart

API version: 1.1
Contact: contact@api2cart.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MarketplaceProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketplaceProduct{}

// MarketplaceProduct struct for MarketplaceProduct
type MarketplaceProduct struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	UAsin *string `json:"u_asin,omitempty"`
	UEan *string `json:"u_ean,omitempty"`
	UGtin *string `json:"u_gtin,omitempty"`
	UIsbn *string `json:"u_isbn,omitempty"`
	UMpn *string `json:"u_mpn,omitempty"`
	UUpc *string `json:"u_upc,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Url *string `json:"url,omitempty"`
	Price *float32 `json:"price,omitempty"`
	Images []Image `json:"images,omitempty"`
	ProductOptions []ProductOption `json:"product_options,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	Brand *string `json:"brand,omitempty"`
	Weight *float32 `json:"weight,omitempty"`
	WeightUnit *string `json:"weight_unit,omitempty"`
	DimensionsUnit *string `json:"dimensions_unit,omitempty"`
	Width *float32 `json:"width,omitempty"`
	Height *float32 `json:"height,omitempty"`
	Length *float32 `json:"length,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewMarketplaceProduct instantiates a new MarketplaceProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceProduct() *MarketplaceProduct {
	this := MarketplaceProduct{}
	return &this
}

// NewMarketplaceProductWithDefaults instantiates a new MarketplaceProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceProductWithDefaults() *MarketplaceProduct {
	this := MarketplaceProduct{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MarketplaceProduct) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MarketplaceProduct) SetType(v string) {
	o.Type = &v
}

// GetUAsin returns the UAsin field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetUAsin() string {
	if o == nil || IsNil(o.UAsin) {
		var ret string
		return ret
	}
	return *o.UAsin
}

// GetUAsinOk returns a tuple with the UAsin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetUAsinOk() (*string, bool) {
	if o == nil || IsNil(o.UAsin) {
		return nil, false
	}
	return o.UAsin, true
}

// HasUAsin returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasUAsin() bool {
	if o != nil && !IsNil(o.UAsin) {
		return true
	}

	return false
}

// SetUAsin gets a reference to the given string and assigns it to the UAsin field.
func (o *MarketplaceProduct) SetUAsin(v string) {
	o.UAsin = &v
}

// GetUEan returns the UEan field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetUEan() string {
	if o == nil || IsNil(o.UEan) {
		var ret string
		return ret
	}
	return *o.UEan
}

// GetUEanOk returns a tuple with the UEan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetUEanOk() (*string, bool) {
	if o == nil || IsNil(o.UEan) {
		return nil, false
	}
	return o.UEan, true
}

// HasUEan returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasUEan() bool {
	if o != nil && !IsNil(o.UEan) {
		return true
	}

	return false
}

// SetUEan gets a reference to the given string and assigns it to the UEan field.
func (o *MarketplaceProduct) SetUEan(v string) {
	o.UEan = &v
}

// GetUGtin returns the UGtin field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetUGtin() string {
	if o == nil || IsNil(o.UGtin) {
		var ret string
		return ret
	}
	return *o.UGtin
}

// GetUGtinOk returns a tuple with the UGtin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetUGtinOk() (*string, bool) {
	if o == nil || IsNil(o.UGtin) {
		return nil, false
	}
	return o.UGtin, true
}

// HasUGtin returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasUGtin() bool {
	if o != nil && !IsNil(o.UGtin) {
		return true
	}

	return false
}

// SetUGtin gets a reference to the given string and assigns it to the UGtin field.
func (o *MarketplaceProduct) SetUGtin(v string) {
	o.UGtin = &v
}

// GetUIsbn returns the UIsbn field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetUIsbn() string {
	if o == nil || IsNil(o.UIsbn) {
		var ret string
		return ret
	}
	return *o.UIsbn
}

// GetUIsbnOk returns a tuple with the UIsbn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetUIsbnOk() (*string, bool) {
	if o == nil || IsNil(o.UIsbn) {
		return nil, false
	}
	return o.UIsbn, true
}

// HasUIsbn returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasUIsbn() bool {
	if o != nil && !IsNil(o.UIsbn) {
		return true
	}

	return false
}

// SetUIsbn gets a reference to the given string and assigns it to the UIsbn field.
func (o *MarketplaceProduct) SetUIsbn(v string) {
	o.UIsbn = &v
}

// GetUMpn returns the UMpn field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetUMpn() string {
	if o == nil || IsNil(o.UMpn) {
		var ret string
		return ret
	}
	return *o.UMpn
}

// GetUMpnOk returns a tuple with the UMpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetUMpnOk() (*string, bool) {
	if o == nil || IsNil(o.UMpn) {
		return nil, false
	}
	return o.UMpn, true
}

// HasUMpn returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasUMpn() bool {
	if o != nil && !IsNil(o.UMpn) {
		return true
	}

	return false
}

// SetUMpn gets a reference to the given string and assigns it to the UMpn field.
func (o *MarketplaceProduct) SetUMpn(v string) {
	o.UMpn = &v
}

// GetUUpc returns the UUpc field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetUUpc() string {
	if o == nil || IsNil(o.UUpc) {
		var ret string
		return ret
	}
	return *o.UUpc
}

// GetUUpcOk returns a tuple with the UUpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetUUpcOk() (*string, bool) {
	if o == nil || IsNil(o.UUpc) {
		return nil, false
	}
	return o.UUpc, true
}

// HasUUpc returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasUUpc() bool {
	if o != nil && !IsNil(o.UUpc) {
		return true
	}

	return false
}

// SetUUpc gets a reference to the given string and assigns it to the UUpc field.
func (o *MarketplaceProduct) SetUUpc(v string) {
	o.UUpc = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MarketplaceProduct) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MarketplaceProduct) SetDescription(v string) {
	o.Description = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MarketplaceProduct) SetUrl(v string) {
	o.Url = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *MarketplaceProduct) SetPrice(v float32) {
	o.Price = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetImages() []Image {
	if o == nil || IsNil(o.Images) {
		var ret []Image
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetImagesOk() ([]Image, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []Image and assigns it to the Images field.
func (o *MarketplaceProduct) SetImages(v []Image) {
	o.Images = v
}

// GetProductOptions returns the ProductOptions field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetProductOptions() []ProductOption {
	if o == nil || IsNil(o.ProductOptions) {
		var ret []ProductOption
		return ret
	}
	return o.ProductOptions
}

// GetProductOptionsOk returns a tuple with the ProductOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetProductOptionsOk() ([]ProductOption, bool) {
	if o == nil || IsNil(o.ProductOptions) {
		return nil, false
	}
	return o.ProductOptions, true
}

// HasProductOptions returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasProductOptions() bool {
	if o != nil && !IsNil(o.ProductOptions) {
		return true
	}

	return false
}

// SetProductOptions gets a reference to the given []ProductOption and assigns it to the ProductOptions field.
func (o *MarketplaceProduct) SetProductOptions(v []ProductOption) {
	o.ProductOptions = v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetManufacturer() string {
	if o == nil || IsNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetManufacturerOk() (*string, bool) {
	if o == nil || IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasManufacturer() bool {
	if o != nil && !IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *MarketplaceProduct) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *MarketplaceProduct) SetBrand(v string) {
	o.Brand = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetWeight() float32 {
	if o == nil || IsNil(o.Weight) {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *MarketplaceProduct) SetWeight(v float32) {
	o.Weight = &v
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetWeightUnit() string {
	if o == nil || IsNil(o.WeightUnit) {
		var ret string
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetWeightUnitOk() (*string, bool) {
	if o == nil || IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasWeightUnit() bool {
	if o != nil && !IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given string and assigns it to the WeightUnit field.
func (o *MarketplaceProduct) SetWeightUnit(v string) {
	o.WeightUnit = &v
}

// GetDimensionsUnit returns the DimensionsUnit field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetDimensionsUnit() string {
	if o == nil || IsNil(o.DimensionsUnit) {
		var ret string
		return ret
	}
	return *o.DimensionsUnit
}

// GetDimensionsUnitOk returns a tuple with the DimensionsUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetDimensionsUnitOk() (*string, bool) {
	if o == nil || IsNil(o.DimensionsUnit) {
		return nil, false
	}
	return o.DimensionsUnit, true
}

// HasDimensionsUnit returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasDimensionsUnit() bool {
	if o != nil && !IsNil(o.DimensionsUnit) {
		return true
	}

	return false
}

// SetDimensionsUnit gets a reference to the given string and assigns it to the DimensionsUnit field.
func (o *MarketplaceProduct) SetDimensionsUnit(v string) {
	o.DimensionsUnit = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *MarketplaceProduct) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *MarketplaceProduct) SetHeight(v float32) {
	o.Height = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetLength() float32 {
	if o == nil || IsNil(o.Length) {
		var ret float32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetLengthOk() (*float32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given float32 and assigns it to the Length field.
func (o *MarketplaceProduct) SetLength(v float32) {
	o.Length = &v
}

// GetAdditionalFields returns the AdditionalFields field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetAdditionalFields() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalFields
}

// GetAdditionalFieldsOk returns a tuple with the AdditionalFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetAdditionalFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalFields) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalFields, true
}

// HasAdditionalFields returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasAdditionalFields() bool {
	if o != nil && !IsNil(o.AdditionalFields) {
		return true
	}

	return false
}

// SetAdditionalFields gets a reference to the given map[string]interface{} and assigns it to the AdditionalFields field.
func (o *MarketplaceProduct) SetAdditionalFields(v map[string]interface{}) {
	o.AdditionalFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *MarketplaceProduct) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceProduct) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *MarketplaceProduct) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *MarketplaceProduct) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o MarketplaceProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketplaceProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UAsin) {
		toSerialize["u_asin"] = o.UAsin
	}
	if !IsNil(o.UEan) {
		toSerialize["u_ean"] = o.UEan
	}
	if !IsNil(o.UGtin) {
		toSerialize["u_gtin"] = o.UGtin
	}
	if !IsNil(o.UIsbn) {
		toSerialize["u_isbn"] = o.UIsbn
	}
	if !IsNil(o.UMpn) {
		toSerialize["u_mpn"] = o.UMpn
	}
	if !IsNil(o.UUpc) {
		toSerialize["u_upc"] = o.UUpc
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.ProductOptions) {
		toSerialize["product_options"] = o.ProductOptions
	}
	if !IsNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.WeightUnit) {
		toSerialize["weight_unit"] = o.WeightUnit
	}
	if !IsNil(o.DimensionsUnit) {
		toSerialize["dimensions_unit"] = o.DimensionsUnit
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.AdditionalFields) {
		toSerialize["additional_fields"] = o.AdditionalFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableMarketplaceProduct struct {
	value *MarketplaceProduct
	isSet bool
}

func (v NullableMarketplaceProduct) Get() *MarketplaceProduct {
	return v.value
}

func (v *NullableMarketplaceProduct) Set(val *MarketplaceProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceProduct(val *MarketplaceProduct) *NullableMarketplaceProduct {
	return &NullableMarketplaceProduct{value: val, isSet: true}
}

func (v NullableMarketplaceProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


