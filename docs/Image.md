# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**HttpPath** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**MimeType** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt32** |  | [optional] 
**CreateAt** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**ModifiedAt** | Pointer to [**NullableA2CDateTime**](A2CDateTime.md) |  | [optional] 
**Alt** | Pointer to **NullableString** |  | [optional] 
**Avail** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewImage

`func NewImage() *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Image) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Image) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Image) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Image) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetHttpPath

`func (o *Image) GetHttpPath() string`

GetHttpPath returns the HttpPath field if non-nil, zero value otherwise.

### GetHttpPathOk

`func (o *Image) GetHttpPathOk() (*string, bool)`

GetHttpPathOk returns a tuple with the HttpPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPath

`func (o *Image) SetHttpPath(v string)`

SetHttpPath sets HttpPath field to given value.

### HasHttpPath

`func (o *Image) HasHttpPath() bool`

HasHttpPath returns a boolean if a field has been set.

### GetFileName

`func (o *Image) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Image) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Image) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Image) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *Image) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *Image) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetMimeType

`func (o *Image) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Image) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Image) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *Image) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeTypeNil

`func (o *Image) SetMimeTypeNil(b bool)`

 SetMimeTypeNil sets the value for MimeType to be an explicit nil

### UnsetMimeType
`func (o *Image) UnsetMimeType()`

UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
### GetSize

`func (o *Image) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Image) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Image) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Image) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *Image) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Image) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetCreateAt

`func (o *Image) GetCreateAt() A2CDateTime`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *Image) GetCreateAtOk() (*A2CDateTime, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *Image) SetCreateAt(v A2CDateTime)`

SetCreateAt sets CreateAt field to given value.

### HasCreateAt

`func (o *Image) HasCreateAt() bool`

HasCreateAt returns a boolean if a field has been set.

### SetCreateAtNil

`func (o *Image) SetCreateAtNil(b bool)`

 SetCreateAtNil sets the value for CreateAt to be an explicit nil

### UnsetCreateAt
`func (o *Image) UnsetCreateAt()`

UnsetCreateAt ensures that no value is present for CreateAt, not even an explicit nil
### GetModifiedAt

`func (o *Image) GetModifiedAt() A2CDateTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Image) GetModifiedAtOk() (*A2CDateTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Image) SetModifiedAt(v A2CDateTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Image) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### SetModifiedAtNil

`func (o *Image) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *Image) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetAlt

`func (o *Image) GetAlt() string`

GetAlt returns the Alt field if non-nil, zero value otherwise.

### GetAltOk

`func (o *Image) GetAltOk() (*string, bool)`

GetAltOk returns a tuple with the Alt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlt

`func (o *Image) SetAlt(v string)`

SetAlt sets Alt field to given value.

### HasAlt

`func (o *Image) HasAlt() bool`

HasAlt returns a boolean if a field has been set.

### SetAltNil

`func (o *Image) SetAltNil(b bool)`

 SetAltNil sets the value for Alt to be an explicit nil

### UnsetAlt
`func (o *Image) UnsetAlt()`

UnsetAlt ensures that no value is present for Alt, not even an explicit nil
### GetAvail

`func (o *Image) GetAvail() bool`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *Image) GetAvailOk() (*bool, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *Image) SetAvail(v bool)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *Image) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetSortOrder

`func (o *Image) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Image) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Image) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Image) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *Image) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *Image) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetType

`func (o *Image) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Image) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Image) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Image) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Image) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Image) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Image) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Image) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *Image) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *Image) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
### GetCustomFields

`func (o *Image) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Image) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Image) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Image) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *Image) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *Image) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


