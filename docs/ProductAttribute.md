# ProductAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeId** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LangId** | Pointer to **string** |  | [optional] 
**StoreId** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**AttributeGroupId** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**VariantId** | Pointer to **string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProductAttribute

`func NewProductAttribute() *ProductAttribute`

NewProductAttribute instantiates a new ProductAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAttributeWithDefaults

`func NewProductAttributeWithDefaults() *ProductAttribute`

NewProductAttributeWithDefaults instantiates a new ProductAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeId

`func (o *ProductAttribute) GetAttributeId() string`

GetAttributeId returns the AttributeId field if non-nil, zero value otherwise.

### GetAttributeIdOk

`func (o *ProductAttribute) GetAttributeIdOk() (*string, bool)`

GetAttributeIdOk returns a tuple with the AttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeId

`func (o *ProductAttribute) SetAttributeId(v string)`

SetAttributeId sets AttributeId field to given value.

### HasAttributeId

`func (o *ProductAttribute) HasAttributeId() bool`

HasAttributeId returns a boolean if a field has been set.

### GetCode

`func (o *ProductAttribute) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProductAttribute) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProductAttribute) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProductAttribute) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ProductAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLangId

`func (o *ProductAttribute) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *ProductAttribute) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *ProductAttribute) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *ProductAttribute) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetStoreId

`func (o *ProductAttribute) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *ProductAttribute) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *ProductAttribute) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *ProductAttribute) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetValue

`func (o *ProductAttribute) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProductAttribute) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProductAttribute) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProductAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRequired

`func (o *ProductAttribute) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ProductAttribute) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ProductAttribute) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ProductAttribute) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetVisible

`func (o *ProductAttribute) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ProductAttribute) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ProductAttribute) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ProductAttribute) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetType

`func (o *ProductAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductAttribute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPosition

`func (o *ProductAttribute) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProductAttribute) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProductAttribute) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ProductAttribute) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetAttributeGroupId

`func (o *ProductAttribute) GetAttributeGroupId() string`

GetAttributeGroupId returns the AttributeGroupId field if non-nil, zero value otherwise.

### GetAttributeGroupIdOk

`func (o *ProductAttribute) GetAttributeGroupIdOk() (*string, bool)`

GetAttributeGroupIdOk returns a tuple with the AttributeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeGroupId

`func (o *ProductAttribute) SetAttributeGroupId(v string)`

SetAttributeGroupId sets AttributeGroupId field to given value.

### HasAttributeGroupId

`func (o *ProductAttribute) HasAttributeGroupId() bool`

HasAttributeGroupId returns a boolean if a field has been set.

### GetProductId

`func (o *ProductAttribute) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductAttribute) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductAttribute) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductAttribute) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetVariantId

`func (o *ProductAttribute) GetVariantId() string`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *ProductAttribute) GetVariantIdOk() (*string, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *ProductAttribute) SetVariantId(v string)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *ProductAttribute) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ProductAttribute) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ProductAttribute) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ProductAttribute) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ProductAttribute) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ProductAttribute) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProductAttribute) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProductAttribute) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProductAttribute) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


