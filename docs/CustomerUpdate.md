# CustomerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Entity id | [optional] 
**GroupId** | Pointer to **string** | Customer group_id | [optional] 
**GroupIds** | Pointer to **string** | Groups that will be assigned to a customer | [optional] 
**Group** | Pointer to **string** | Defines the group where the customer | [optional] 
**Email** | Pointer to **string** | Defines customer&#39;s email | [optional] 
**Phone** | Pointer to **string** | Defines customer&#39;s phone number | [optional] 
**FirstName** | Pointer to **string** | Defines customer&#39;s first name | [optional] 
**LastName** | Pointer to **string** | Defines customer&#39;s last name | [optional] 
**BirthDay** | Pointer to **string** | Defines customer&#39;s birthday | [optional] 
**NewsLetterSubscription** | Pointer to **bool** | Defines whether the newsletter subscription is available for the user | [optional] 
**Consents** | Pointer to [**[]CustomerAddConsentsInner**](CustomerAddConsentsInner.md) | Defines consents to notifications | [optional] 
**Tags** | Pointer to **string** | Customer tags | [optional] 
**Gender** | Pointer to **string** | Defines customer&#39;s gender | [optional] 
**Note** | Pointer to **string** | The customer note. | [optional] 
**Status** | Pointer to **string** | Defines customer&#39;s status | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 
**Address** | Pointer to [**[]CustomerUpdateAddressInner**](CustomerUpdateAddressInner.md) |  | [optional] 

## Methods

### NewCustomerUpdate

`func NewCustomerUpdate() *CustomerUpdate`

NewCustomerUpdate instantiates a new CustomerUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerUpdateWithDefaults

`func NewCustomerUpdateWithDefaults() *CustomerUpdate`

NewCustomerUpdateWithDefaults instantiates a new CustomerUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *CustomerUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CustomerUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CustomerUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CustomerUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupIds

`func (o *CustomerUpdate) GetGroupIds() string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *CustomerUpdate) GetGroupIdsOk() (*string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *CustomerUpdate) SetGroupIds(v string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *CustomerUpdate) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetGroup

`func (o *CustomerUpdate) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CustomerUpdate) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CustomerUpdate) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CustomerUpdate) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerUpdate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerUpdate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerUpdate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerUpdate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerUpdate) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerUpdate) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerUpdate) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerUpdate) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetFirstName

`func (o *CustomerUpdate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerUpdate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerUpdate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CustomerUpdate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CustomerUpdate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerUpdate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerUpdate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CustomerUpdate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetBirthDay

`func (o *CustomerUpdate) GetBirthDay() string`

GetBirthDay returns the BirthDay field if non-nil, zero value otherwise.

### GetBirthDayOk

`func (o *CustomerUpdate) GetBirthDayOk() (*string, bool)`

GetBirthDayOk returns a tuple with the BirthDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDay

`func (o *CustomerUpdate) SetBirthDay(v string)`

SetBirthDay sets BirthDay field to given value.

### HasBirthDay

`func (o *CustomerUpdate) HasBirthDay() bool`

HasBirthDay returns a boolean if a field has been set.

### GetNewsLetterSubscription

`func (o *CustomerUpdate) GetNewsLetterSubscription() bool`

GetNewsLetterSubscription returns the NewsLetterSubscription field if non-nil, zero value otherwise.

### GetNewsLetterSubscriptionOk

`func (o *CustomerUpdate) GetNewsLetterSubscriptionOk() (*bool, bool)`

GetNewsLetterSubscriptionOk returns a tuple with the NewsLetterSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsLetterSubscription

`func (o *CustomerUpdate) SetNewsLetterSubscription(v bool)`

SetNewsLetterSubscription sets NewsLetterSubscription field to given value.

### HasNewsLetterSubscription

`func (o *CustomerUpdate) HasNewsLetterSubscription() bool`

HasNewsLetterSubscription returns a boolean if a field has been set.

### GetConsents

`func (o *CustomerUpdate) GetConsents() []CustomerAddConsentsInner`

GetConsents returns the Consents field if non-nil, zero value otherwise.

### GetConsentsOk

`func (o *CustomerUpdate) GetConsentsOk() (*[]CustomerAddConsentsInner, bool)`

GetConsentsOk returns a tuple with the Consents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsents

`func (o *CustomerUpdate) SetConsents(v []CustomerAddConsentsInner)`

SetConsents sets Consents field to given value.

### HasConsents

`func (o *CustomerUpdate) HasConsents() bool`

HasConsents returns a boolean if a field has been set.

### GetTags

`func (o *CustomerUpdate) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CustomerUpdate) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CustomerUpdate) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CustomerUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetGender

`func (o *CustomerUpdate) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CustomerUpdate) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CustomerUpdate) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CustomerUpdate) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetNote

`func (o *CustomerUpdate) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CustomerUpdate) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CustomerUpdate) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CustomerUpdate) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStoreId

`func (o *CustomerUpdate) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CustomerUpdate) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CustomerUpdate) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *CustomerUpdate) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *CustomerUpdate) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *CustomerUpdate) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *CustomerUpdate) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *CustomerUpdate) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetAddress

`func (o *CustomerUpdate) GetAddress() []CustomerUpdateAddressInner`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerUpdate) GetAddressOk() (*[]CustomerUpdateAddressInner, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerUpdate) SetAddress(v []CustomerUpdateAddressInner)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustomerUpdate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


