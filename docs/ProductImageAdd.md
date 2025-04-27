# ProductImageAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines image&#39;s types that are specified by comma-separated list | 
**ImageName** | **string** | Defines image&#39;s name | 
**ProductId** | Pointer to **string** | Defines product id where the image should be added | [optional] 
**ProductVariantId** | Pointer to **string** | Defines product&#39;s variants specified by variant id | [optional] 
**VariantIds** | Pointer to **string** | Defines product&#39;s variants ids | [optional] 
**OptionValueIds** | Pointer to **string** | Defines product&#39;s option values ids | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**LangId** | Pointer to **string** | Add product image on specified language id | [optional] 
**Url** | Pointer to **string** | Defines URL of the image that has to be added | [optional] 
**Content** | Pointer to **string** | Content(body) encoded in base64 of image file | [optional] 
**Label** | Pointer to **string** | Defines alternative text that has to be attached to the picture | [optional] 
**Mime** | Pointer to **string** | Mime type of image http://en.wikipedia.org/wiki/Internet_media_type. | [optional] 
**Position** | Pointer to **int32** | Defines image’s position in the list | [optional] [default to 0]
**UseLatestApiVersion** | Pointer to **bool** | Use the latest platform API version | [optional] [default to false]

## Methods

### NewProductImageAdd

`func NewProductImageAdd(type_ string, imageName string, ) *ProductImageAdd`

NewProductImageAdd instantiates a new ProductImageAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductImageAddWithDefaults

`func NewProductImageAddWithDefaults() *ProductImageAdd`

NewProductImageAddWithDefaults instantiates a new ProductImageAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProductImageAdd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductImageAdd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductImageAdd) SetType(v string)`

SetType sets Type field to given value.


### GetImageName

`func (o *ProductImageAdd) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ProductImageAdd) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ProductImageAdd) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetProductId

`func (o *ProductImageAdd) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductImageAdd) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductImageAdd) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductImageAdd) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductVariantId

`func (o *ProductImageAdd) GetProductVariantId() string`

GetProductVariantId returns the ProductVariantId field if non-nil, zero value otherwise.

### GetProductVariantIdOk

`func (o *ProductImageAdd) GetProductVariantIdOk() (*string, bool)`

GetProductVariantIdOk returns a tuple with the ProductVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariantId

`func (o *ProductImageAdd) SetProductVariantId(v string)`

SetProductVariantId sets ProductVariantId field to given value.

### HasProductVariantId

`func (o *ProductImageAdd) HasProductVariantId() bool`

HasProductVariantId returns a boolean if a field has been set.

### GetVariantIds

`func (o *ProductImageAdd) GetVariantIds() string`

GetVariantIds returns the VariantIds field if non-nil, zero value otherwise.

### GetVariantIdsOk

`func (o *ProductImageAdd) GetVariantIdsOk() (*string, bool)`

GetVariantIdsOk returns a tuple with the VariantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantIds

`func (o *ProductImageAdd) SetVariantIds(v string)`

SetVariantIds sets VariantIds field to given value.

### HasVariantIds

`func (o *ProductImageAdd) HasVariantIds() bool`

HasVariantIds returns a boolean if a field has been set.

### GetOptionValueIds

`func (o *ProductImageAdd) GetOptionValueIds() string`

GetOptionValueIds returns the OptionValueIds field if non-nil, zero value otherwise.

### GetOptionValueIdsOk

`func (o *ProductImageAdd) GetOptionValueIdsOk() (*string, bool)`

GetOptionValueIdsOk returns a tuple with the OptionValueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionValueIds

`func (o *ProductImageAdd) SetOptionValueIds(v string)`

SetOptionValueIds sets OptionValueIds field to given value.

### HasOptionValueIds

`func (o *ProductImageAdd) HasOptionValueIds() bool`

HasOptionValueIds returns a boolean if a field has been set.

### GetStoreId

`func (o *ProductImageAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductImageAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductImageAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductImageAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetLangId

`func (o *ProductImageAdd) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *ProductImageAdd) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *ProductImageAdd) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *ProductImageAdd) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetUrl

`func (o *ProductImageAdd) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProductImageAdd) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProductImageAdd) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProductImageAdd) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetContent

`func (o *ProductImageAdd) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ProductImageAdd) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ProductImageAdd) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ProductImageAdd) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetLabel

`func (o *ProductImageAdd) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ProductImageAdd) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ProductImageAdd) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ProductImageAdd) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMime

`func (o *ProductImageAdd) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *ProductImageAdd) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *ProductImageAdd) SetMime(v string)`

SetMime sets Mime field to given value.

### HasMime

`func (o *ProductImageAdd) HasMime() bool`

HasMime returns a boolean if a field has been set.

### GetPosition

`func (o *ProductImageAdd) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProductImageAdd) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProductImageAdd) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ProductImageAdd) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetUseLatestApiVersion

`func (o *ProductImageAdd) GetUseLatestApiVersion() bool`

GetUseLatestApiVersion returns the UseLatestApiVersion field if non-nil, zero value otherwise.

### GetUseLatestApiVersionOk

`func (o *ProductImageAdd) GetUseLatestApiVersionOk() (*bool, bool)`

GetUseLatestApiVersionOk returns a tuple with the UseLatestApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLatestApiVersion

`func (o *ProductImageAdd) SetUseLatestApiVersion(v bool)`

SetUseLatestApiVersion sets UseLatestApiVersion field to given value.

### HasUseLatestApiVersion

`func (o *ProductImageAdd) HasUseLatestApiVersion() bool`

HasUseLatestApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


