# CartConfigUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbTablesPrefix** | Pointer to **string** | This parameter is deprecated for this method. Please, use this parameter in method account.config.update | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | This parameter sets the list of params to the shopping cart. | [optional] 
**StoreId** | Pointer to **string** | Store Id | [optional] 
**UserAgent** | Pointer to **string** | This parameter allows you to set your custom user agent, which will be used in requests to the store. Please use it cautiously, as the store&#39;s firewall may block specific values. | [optional] 

## Methods

### NewCartConfigUpdate

`func NewCartConfigUpdate() *CartConfigUpdate`

NewCartConfigUpdate instantiates a new CartConfigUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartConfigUpdateWithDefaults

`func NewCartConfigUpdateWithDefaults() *CartConfigUpdate`

NewCartConfigUpdateWithDefaults instantiates a new CartConfigUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbTablesPrefix

`func (o *CartConfigUpdate) GetDbTablesPrefix() string`

GetDbTablesPrefix returns the DbTablesPrefix field if non-nil, zero value otherwise.

### GetDbTablesPrefixOk

`func (o *CartConfigUpdate) GetDbTablesPrefixOk() (*string, bool)`

GetDbTablesPrefixOk returns a tuple with the DbTablesPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbTablesPrefix

`func (o *CartConfigUpdate) SetDbTablesPrefix(v string)`

SetDbTablesPrefix sets DbTablesPrefix field to given value.

### HasDbTablesPrefix

`func (o *CartConfigUpdate) HasDbTablesPrefix() bool`

HasDbTablesPrefix returns a boolean if a field has been set.

### GetCustomFields

`func (o *CartConfigUpdate) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CartConfigUpdate) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CartConfigUpdate) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CartConfigUpdate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetStoreId

`func (o *CartConfigUpdate) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CartConfigUpdate) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CartConfigUpdate) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *CartConfigUpdate) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetUserAgent

`func (o *CartConfigUpdate) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *CartConfigUpdate) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *CartConfigUpdate) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *CartConfigUpdate) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


