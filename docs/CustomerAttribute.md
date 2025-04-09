# CustomerAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeId** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Values** | Pointer to [**[]CustomerAttributeValue**](CustomerAttributeValue.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCustomerAttribute

`func NewCustomerAttribute() *CustomerAttribute`

NewCustomerAttribute instantiates a new CustomerAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAttributeWithDefaults

`func NewCustomerAttributeWithDefaults() *CustomerAttribute`

NewCustomerAttributeWithDefaults instantiates a new CustomerAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeId

`func (o *CustomerAttribute) GetAttributeId() string`

GetAttributeId returns the AttributeId field if non-nil, zero value otherwise.

### GetAttributeIdOk

`func (o *CustomerAttribute) GetAttributeIdOk() (*string, bool)`

GetAttributeIdOk returns a tuple with the AttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeId

`func (o *CustomerAttribute) SetAttributeId(v string)`

SetAttributeId sets AttributeId field to given value.

### HasAttributeId

`func (o *CustomerAttribute) HasAttributeId() bool`

HasAttributeId returns a boolean if a field has been set.

### GetCode

`func (o *CustomerAttribute) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustomerAttribute) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustomerAttribute) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CustomerAttribute) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *CustomerAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CustomerAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerAttribute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomerAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValues

`func (o *CustomerAttribute) GetValues() []CustomerAttributeValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CustomerAttribute) GetValuesOk() (*[]CustomerAttributeValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CustomerAttribute) SetValues(v []CustomerAttributeValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *CustomerAttribute) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *CustomerAttribute) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *CustomerAttribute) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *CustomerAttribute) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *CustomerAttribute) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *CustomerAttribute) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CustomerAttribute) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CustomerAttribute) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CustomerAttribute) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


