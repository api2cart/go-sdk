# Brand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**ModifiedTime** | Pointer to **string** |  | [optional] 
**FullDescription** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**StoresIds** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**MetaTitle** | Pointer to **string** |  | [optional] 
**MetaKeywords** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBrand

`func NewBrand() *Brand`

NewBrand instantiates a new Brand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandWithDefaults

`func NewBrandWithDefaults() *Brand`

NewBrandWithDefaults instantiates a new Brand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Brand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Brand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Brand) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Brand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Brand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Brand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Brand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Brand) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Brand) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Brand) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Brand) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Brand) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Brand) GetModifiedTime() string`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Brand) GetModifiedTimeOk() (*string, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Brand) SetModifiedTime(v string)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Brand) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetFullDescription

`func (o *Brand) GetFullDescription() string`

GetFullDescription returns the FullDescription field if non-nil, zero value otherwise.

### GetFullDescriptionOk

`func (o *Brand) GetFullDescriptionOk() (*string, bool)`

GetFullDescriptionOk returns a tuple with the FullDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDescription

`func (o *Brand) SetFullDescription(v string)`

SetFullDescription sets FullDescription field to given value.

### HasFullDescription

`func (o *Brand) HasFullDescription() bool`

HasFullDescription returns a boolean if a field has been set.

### GetShortDescription

`func (o *Brand) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *Brand) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *Brand) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *Brand) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetStoresIds

`func (o *Brand) GetStoresIds() []string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *Brand) GetStoresIdsOk() (*[]string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *Brand) SetStoresIds(v []string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *Brand) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetActive

`func (o *Brand) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Brand) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Brand) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Brand) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUrl

`func (o *Brand) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Brand) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Brand) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Brand) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMetaTitle

`func (o *Brand) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *Brand) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *Brand) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *Brand) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *Brand) GetMetaKeywords() string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *Brand) GetMetaKeywordsOk() (*string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *Brand) SetMetaKeywords(v string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *Brand) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetMetaDescription

`func (o *Brand) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *Brand) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *Brand) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *Brand) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetImages

`func (o *Brand) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Brand) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Brand) SetImages(v []Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *Brand) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Brand) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Brand) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Brand) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Brand) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *Brand) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Brand) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Brand) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Brand) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


