# Subscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Subscribed** | Pointer to **bool** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**StoresIds** | Pointer to **[]string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**ModifiedTime** | Pointer to **string** |  | [optional] 
**LangId** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSubscriber

`func NewSubscriber() *Subscriber`

NewSubscriber instantiates a new Subscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberWithDefaults

`func NewSubscriberWithDefaults() *Subscriber`

NewSubscriberWithDefaults instantiates a new Subscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscriber) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscriber) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscriber) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscriber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomerId

`func (o *Subscriber) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Subscriber) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Subscriber) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Subscriber) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEmail

`func (o *Subscriber) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Subscriber) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Subscriber) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Subscriber) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSubscribed

`func (o *Subscriber) GetSubscribed() bool`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *Subscriber) GetSubscribedOk() (*bool, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *Subscriber) SetSubscribed(v bool)`

SetSubscribed sets Subscribed field to given value.

### HasSubscribed

`func (o *Subscriber) HasSubscribed() bool`

HasSubscribed returns a boolean if a field has been set.

### GetFirstName

`func (o *Subscriber) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Subscriber) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Subscriber) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Subscriber) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Subscriber) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Subscriber) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Subscriber) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Subscriber) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetStoresIds

`func (o *Subscriber) GetStoresIds() []string`

GetStoresIds returns the StoresIds field if non-nil, zero value otherwise.

### GetStoresIdsOk

`func (o *Subscriber) GetStoresIdsOk() (*[]string, bool)`

GetStoresIdsOk returns a tuple with the StoresIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoresIds

`func (o *Subscriber) SetStoresIds(v []string)`

SetStoresIds sets StoresIds field to given value.

### HasStoresIds

`func (o *Subscriber) HasStoresIds() bool`

HasStoresIds returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Subscriber) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Subscriber) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Subscriber) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Subscriber) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetModifiedTime

`func (o *Subscriber) GetModifiedTime() string`

GetModifiedTime returns the ModifiedTime field if non-nil, zero value otherwise.

### GetModifiedTimeOk

`func (o *Subscriber) GetModifiedTimeOk() (*string, bool)`

GetModifiedTimeOk returns a tuple with the ModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedTime

`func (o *Subscriber) SetModifiedTime(v string)`

SetModifiedTime sets ModifiedTime field to given value.

### HasModifiedTime

`func (o *Subscriber) HasModifiedTime() bool`

HasModifiedTime returns a boolean if a field has been set.

### GetLangId

`func (o *Subscriber) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *Subscriber) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *Subscriber) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *Subscriber) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetGender

`func (o *Subscriber) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Subscriber) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Subscriber) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Subscriber) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *Subscriber) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *Subscriber) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *Subscriber) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *Subscriber) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *Subscriber) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Subscriber) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Subscriber) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Subscriber) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


