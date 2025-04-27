# ProductVariantImageAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** | Defines product id where the variant image has to be added | [optional] 
**ProductVariantId** | **string** | Defines product&#39;s variants specified by variant id | 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**ImageName** | **string** | Defines image&#39;s name | 
**Type** | **string** | Defines image&#39;s types that are specified by comma-separated list | [default to "base"]
**Url** | Pointer to **string** | Defines URL of the image that has to be added | [optional] 
**Content** | Pointer to **string** | Content(body) encoded in base64 of image file | [optional] 
**Label** | Pointer to **string** | Defines alternative text that has to be attached to the picture | [optional] 
**Mime** | Pointer to **string** | Mime type of image http://en.wikipedia.org/wiki/Internet_media_type. | [optional] 
**Position** | Pointer to **int32** | Defines image’s position in the list | [optional] [default to 0]
**OptionId** | Pointer to **string** | Defines option id of the product variant for which the image will be added | [optional] 

## Methods

### NewProductVariantImageAdd

`func NewProductVariantImageAdd(productVariantId string, imageName string, type_ string, ) *ProductVariantImageAdd`

NewProductVariantImageAdd instantiates a new ProductVariantImageAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariantImageAddWithDefaults

`func NewProductVariantImageAddWithDefaults() *ProductVariantImageAdd`

NewProductVariantImageAddWithDefaults instantiates a new ProductVariantImageAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProductVariantImageAdd) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductVariantImageAdd) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductVariantImageAdd) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductVariantImageAdd) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductVariantId

`func (o *ProductVariantImageAdd) GetProductVariantId() string`

GetProductVariantId returns the ProductVariantId field if non-nil, zero value otherwise.

### GetProductVariantIdOk

`func (o *ProductVariantImageAdd) GetProductVariantIdOk() (*string, bool)`

GetProductVariantIdOk returns a tuple with the ProductVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariantId

`func (o *ProductVariantImageAdd) SetProductVariantId(v string)`

SetProductVariantId sets ProductVariantId field to given value.


### GetStoreId

`func (o *ProductVariantImageAdd) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductVariantImageAdd) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductVariantImageAdd) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductVariantImageAdd) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetImageName

`func (o *ProductVariantImageAdd) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ProductVariantImageAdd) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ProductVariantImageAdd) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetType

`func (o *ProductVariantImageAdd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductVariantImageAdd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductVariantImageAdd) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *ProductVariantImageAdd) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProductVariantImageAdd) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProductVariantImageAdd) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProductVariantImageAdd) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetContent

`func (o *ProductVariantImageAdd) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ProductVariantImageAdd) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ProductVariantImageAdd) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ProductVariantImageAdd) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetLabel

`func (o *ProductVariantImageAdd) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ProductVariantImageAdd) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ProductVariantImageAdd) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ProductVariantImageAdd) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMime

`func (o *ProductVariantImageAdd) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *ProductVariantImageAdd) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *ProductVariantImageAdd) SetMime(v string)`

SetMime sets Mime field to given value.

### HasMime

`func (o *ProductVariantImageAdd) HasMime() bool`

HasMime returns a boolean if a field has been set.

### GetPosition

`func (o *ProductVariantImageAdd) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProductVariantImageAdd) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProductVariantImageAdd) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ProductVariantImageAdd) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetOptionId

`func (o *ProductVariantImageAdd) GetOptionId() string`

GetOptionId returns the OptionId field if non-nil, zero value otherwise.

### GetOptionIdOk

`func (o *ProductVariantImageAdd) GetOptionIdOk() (*string, bool)`

GetOptionIdOk returns a tuple with the OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionId

`func (o *ProductVariantImageAdd) SetOptionId(v string)`

SetOptionId sets OptionId field to given value.

### HasOptionId

`func (o *ProductVariantImageAdd) HasOptionId() bool`

HasOptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


