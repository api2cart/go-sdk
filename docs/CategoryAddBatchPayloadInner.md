# CategoryAddBatchPayloadInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Avail** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MetaTitle** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**MetaKeywords** | Pointer to **[]string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**SeoUrl** | Pointer to **string** |  | [optional] 
**StoreId** | Pointer to **string** |  | [optional] 
**Images** | Pointer to [**[]CategoryAddBatchPayloadInnerImagesInner**](CategoryAddBatchPayloadInnerImagesInner.md) |  | [optional] 
**StoresIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCategoryAddBatchPayloadInner

`func NewCategoryAddBatchPayloadInner(name string, ) *CategoryAddBatchPayloadInner`

NewCategoryAddBatchPayloadInner instantiates a new CategoryAddBatchPayloadInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryAddBatchPayloadInnerWithDefaults

`func NewCategoryAddBatchPayloadInnerWithDefaults() *CategoryAddBatchPayloadInner`

NewCategoryAddBatchPayloadInnerWithDefaults instantiates a new CategoryAddBatchPayloadInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CategoryAddBatchPayloadInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryAddBatchPayloadInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryAddBatchPayloadInner) SetName(v string)`

SetName sets Name field to given value.


### GetAvail

`func (o *CategoryAddBatchPayloadInner) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *CategoryAddBatchPayloadInner) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *CategoryAddBatchPayloadInner) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *CategoryAddBatchPayloadInner) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetDescription

`func (o *CategoryAddBatchPayloadInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CategoryAddBatchPayloadInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CategoryAddBatchPayloadInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CategoryAddBatchPayloadInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetaTitle

`func (o *CategoryAddBatchPayloadInner) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *CategoryAddBatchPayloadInner) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *CategoryAddBatchPayloadInner) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *CategoryAddBatchPayloadInner) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetMetaDescription

`func (o *CategoryAddBatchPayloadInner) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *CategoryAddBatchPayloadInner) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *CategoryAddBatchPayloadInner) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *CategoryAddBatchPayloadInner) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *CategoryAddBatchPayloadInner) GetMetaKeywords() []string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *CategoryAddBatchPayloadInner) GetMetaKeywordsOk() (*[]string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *CategoryAddBatchPayloadInner) SetMetaKeywords(v []string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *CategoryAddBatchPayloadInner) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetParentId

`func (o *CategoryAddBatchPayloadInner) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CategoryAddBatchPayloadInner) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CategoryAddBatchPayloadInner) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CategoryAddBatchPayloadInner) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetSortOrder

`func (o *CategoryAddBatchPayloadInner) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *CategoryAddBatchPayloadInner) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *CategoryAddBatchPayloadInner) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *CategoryAddBatchPayloadInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSeoUrl

`func (o *CategoryAddBatchPayloadInner) GetSeoUrl() string`

GetSeoUrl returns the SeoUrl field if non-nil, zero value otherwise.

### GetSeoUrlOk

`func (o *CategoryAddBatchPayloadInner) GetSeoUrlOk() (*string, bool)`

GetSeoUrlOk returns a tuple with the SeoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoUrl

`func (o *CategoryAddBatchPayloadInner) SetSeoUrl(v string)`

SetSeoUrl sets SeoUrl field to given value.

### HasSeoUrl

`func (o *CategoryAddBatchPayloadInner) HasSeoUrl() bool`

HasSeoUrl returns a boolean if a field has been set.

### GetStoreId

`func (o *CategoryAddBatchPayloadInner) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CategoryAddBatchPayloadInner) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CategoryAddBatchPayloadInner) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *CategoryAddBatchPayloadInner) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetImages

`func (o *CategoryAddBatchPayloadInner) GetImages() []CategoryAddBatchPayloadInnerImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CategoryAddBatchPayloadInner) GetImagesOk() (*[]CategoryAddBatchPayloadInnerImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CategoryAddBatchPayloadInner) SetImages(v []CategoryAddBatchPayloadInnerImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *CategoryAddBatchPayloadInner) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetStoresIds

`func (o *CategoryAddBatchPayloadInner) GetStoresIds() []string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *CategoryAddBatchPayloadInner) GetStoresIdsOk() (*[]string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *CategoryAddBatchPayloadInner) SetStoresIds(v []string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *CategoryAddBatchPayloadInner) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


