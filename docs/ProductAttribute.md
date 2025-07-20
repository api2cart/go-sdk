# ProductAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeId** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LangId** | Pointer to **NullableString** |  | [optional] 
**StoreId** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 
**Visible** | Pointer to **NullableBool** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Position** | Pointer to **NullableInt32** |  | [optional] 
**AttributeGroupId** | Pointer to **NullableString** |  | [optional] 
**ProductId** | Pointer to **NullableString** |  | [optional] 
**VariantId** | Pointer to **NullableString** |  | [optional] 
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

### SetLangIdNil

`func (o *ProductAttribute) SetLangIdNil(b bool)`

 SetLangIdNil sets the value for LangId to be an explicit nil

### UnsetLangId
`func (o *ProductAttribute) UnsetLangId()`

UnsetLangId ensures that no value is present for LangId, not even an explicit nil
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

### SetStoreIdNil

`func (o *ProductAttribute) SetStoreIdNil(b bool)`

 SetStoreIdNil sets the value for StoreId to be an explicit nil

### UnsetStoreId
`func (o *ProductAttribute) UnsetStoreId()`

UnsetStoreId ensures that no value is present for StoreId, not even an explicit nil
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

### SetValueNil

`func (o *ProductAttribute) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProductAttribute) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
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

### SetRequiredNil

`func (o *ProductAttribute) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ProductAttribute) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
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

### SetVisibleNil

`func (o *ProductAttribute) SetVisibleNil(b bool)`

 SetVisibleNil sets the value for Visible to be an explicit nil

### UnsetVisible
`func (o *ProductAttribute) UnsetVisible()`

UnsetVisible ensures that no value is present for Visible, not even an explicit nil
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

### SetTypeNil

`func (o *ProductAttribute) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ProductAttribute) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
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

### SetPositionNil

`func (o *ProductAttribute) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *ProductAttribute) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
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

### SetAttributeGroupIdNil

`func (o *ProductAttribute) SetAttributeGroupIdNil(b bool)`

 SetAttributeGroupIdNil sets the value for AttributeGroupId to be an explicit nil

### UnsetAttributeGroupId
`func (o *ProductAttribute) UnsetAttributeGroupId()`

UnsetAttributeGroupId ensures that no value is present for AttributeGroupId, not even an explicit nil
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

### SetProductIdNil

`func (o *ProductAttribute) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *ProductAttribute) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
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

### SetVariantIdNil

`func (o *ProductAttribute) SetVariantIdNil(b bool)`

 SetVariantIdNil sets the value for VariantId to be an explicit nil

### UnsetVariantId
`func (o *ProductAttribute) UnsetVariantId()`

UnsetVariantId ensures that no value is present for VariantId, not even an explicit nil
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

### SetAdditionalFieldsNil

`func (o *ProductAttribute) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *ProductAttribute) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil
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

### SetCustomFieldsNil

`func (o *ProductAttribute) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *ProductAttribute) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


