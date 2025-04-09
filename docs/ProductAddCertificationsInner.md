# ProductAddCertificationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Certification ID | 
**Images** | Pointer to [**[]ProductAddCertificationsInnerImagesInner**](ProductAddCertificationsInnerImagesInner.md) | Certification images | [optional] 
**Files** | Pointer to [**[]ProductAddCertificationsInnerFilesInner**](ProductAddCertificationsInnerFilesInner.md) | Certification files | [optional] 

## Methods

### NewProductAddCertificationsInner

`func NewProductAddCertificationsInner(id string, ) *ProductAddCertificationsInner`

NewProductAddCertificationsInner instantiates a new ProductAddCertificationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAddCertificationsInnerWithDefaults

`func NewProductAddCertificationsInnerWithDefaults() *ProductAddCertificationsInner`

NewProductAddCertificationsInnerWithDefaults instantiates a new ProductAddCertificationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductAddCertificationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductAddCertificationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductAddCertificationsInner) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *ProductAddCertificationsInner) GetImages() []ProductAddCertificationsInnerImagesInner`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ProductAddCertificationsInner) GetImagesOk() (*[]ProductAddCertificationsInnerImagesInner, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ProductAddCertificationsInner) SetImages(v []ProductAddCertificationsInnerImagesInner)`

SetImages sets Images field to given value.

### HasImages

`func (o *ProductAddCertificationsInner) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetFiles

`func (o *ProductAddCertificationsInner) GetFiles() []ProductAddCertificationsInnerFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProductAddCertificationsInner) GetFilesOk() (*[]ProductAddCertificationsInnerFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProductAddCertificationsInner) SetFiles(v []ProductAddCertificationsInnerFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ProductAddCertificationsInner) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


