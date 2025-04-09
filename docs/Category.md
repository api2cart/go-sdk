# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedAt** | Pointer to [**A2CDateTime**](A2CDateTime.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**StoresIds** | Pointer to **[]string** |  | [optional] 
**Keywords** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**MetaTitle** | Pointer to **string** |  | [optional] 
**Avail** | Pointer to **bool** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**SeoUrl** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCategory

`func NewCategory() *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Category) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Category) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Category) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Category) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *Category) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Category) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Category) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Category) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Category) GetCreatedAt() A2CDateTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Category) GetCreatedAtOk() (*A2CDateTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Category) SetCreatedAt(v A2CDateTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Category) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Category) GetModifiedAt() A2CDateTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Category) GetModifiedAtOk() (*A2CDateTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Category) SetModifiedAt(v A2CDateTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Category) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Category) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Category) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Category) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Category) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortDescription

`func (o *Category) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *Category) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *Category) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *Category) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetDescription

`func (o *Category) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Category) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Category) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Category) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStoresIds

`func (o *Category) GetStoresIds() []string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *Category) GetStoresIdsOk() (*[]string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *Category) SetStoresIds(v []string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *Category) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetKeywords

`func (o *Category) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *Category) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *Category) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *Category) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetMetaDescription

`func (o *Category) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *Category) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *Category) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *Category) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetMetaTitle

`func (o *Category) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *Category) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *Category) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *Category) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetAvail

`func (o *Category) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *Category) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *Category) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *Category) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetPath

`func (o *Category) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Category) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Category) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Category) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSeoUrl

`func (o *Category) GetSeoUrl() string`

GetSeoUrl returns the SeoUrl field if non-nil, zero value otherwise.

### GetSeoUrlOk

`func (o *Category) GetSeoUrlOk() (*string, bool)`

GetSeoUrlOk returns a tuple with the SeoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoUrl

`func (o *Category) SetSeoUrl(v string)`

SetSeoUrl sets SeoUrl field to given value.

### HasSeoUrl

`func (o *Category) HasSeoUrl() bool`

HasSeoUrl returns a boolean if a field has been set.

### GetSortOrder

`func (o *Category) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Category) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Category) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Category) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetImages

`func (o *Category) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Category) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Category) SetImages(v []Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *Category) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Category) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Category) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Category) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Category) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *Category) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Category) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Category) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Category) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


