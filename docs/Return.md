# Return

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**StoreId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ReturnStatus**](ReturnStatus.md) |  | [optional] 
**OrderProducts** | Pointer to [**[]ReturnOrderProduct**](ReturnOrderProduct.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**StaffNote** | Pointer to **string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewReturn

`func NewReturn() *Return`

NewReturn instantiates a new Return object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnWithDefaults

`func NewReturnWithDefaults() *Return`

NewReturnWithDefaults instantiates a new Return object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Return) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Return) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Return) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Return) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Return) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Return) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Return) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Return) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrderId

`func (o *Return) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Return) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Return) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Return) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCustomerId

`func (o *Return) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Return) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Return) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Return) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetStoreId

`func (o *Return) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *Return) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *Return) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *Return) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Return) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Return) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Return) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Return) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Return) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Return) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Return) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Return) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Return) GetStatus() ReturnStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Return) GetStatusOk() (*ReturnStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Return) SetStatus(v ReturnStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Return) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrderProducts

`func (o *Return) GetOrderProducts() []ReturnOrderProduct`

GetOrderProducts returns the OrderProducts field if non-nil, zero value otherwise.

### GetOrderProductsOk

`func (o *Return) GetOrderProductsOk() (*[]ReturnOrderProduct, bool)`

GetOrderProductsOk returns a tuple with the OrderProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderProducts

`func (o *Return) SetOrderProducts(v []ReturnOrderProduct)`

SetOrderProducts sets OrderProducts field to given value.

### HasOrderProducts

`func (o *Return) HasOrderProducts() bool`

HasOrderProducts returns a boolean if a field has been set.

### GetComment

`func (o *Return) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Return) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Return) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Return) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetStaffNote

`func (o *Return) GetStaffNote() string`

GetStaffNote returns the StaffNote field if non-nil, zero value otherwise.

### GetStaffNoteOk

`func (o *Return) GetStaffNoteOk() (*string, bool)`

GetStaffNoteOk returns a tuple with the StaffNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaffNote

`func (o *Return) SetStaffNote(v string)`

SetStaffNote sets StaffNote field to given value.

### HasStaffNote

`func (o *Return) HasStaffNote() bool`

HasStaffNote returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Return) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Return) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Return) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Return) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *Return) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Return) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Return) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Return) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


