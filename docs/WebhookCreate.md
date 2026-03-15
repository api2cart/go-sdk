# WebhookCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | **string** | Specify the entity that you want to enable webhooks for (e.g product, order, customer, category) | 
**Action** | **string** | Specify what action (event) will trigger the webhook (e.g add, delete, or update) | 
**Callback** | Pointer to **string** | Callback url that returns shipping rates. It should be able to accept POST requests with json data. | [optional] 
**Label** | Pointer to **string** | The name you give to the webhook | [optional] 
**Fields** | Pointer to **string** | Fields the webhook should send | [optional] [default to "force_all"]
**ResponseFields** | Pointer to **string** | Set this parameter in order to choose which entity fields you want to retrieve | [optional] 
**Active** | Pointer to **bool** | Webhook status | [optional] [default to true]
**LangId** | Pointer to **string** | Language id | [optional] 
**StoreId** | Pointer to **string** | Defines store id where the webhook should be assigned | [optional] 
**FilteringConditions** | Pointer to [**ParamDefinitionFilteringConditionsFilterCondition**](ParamDefinitionFilteringConditionsFilterCondition.md) | &lt;p&gt;Defines the logic for filtering webhooks based on the entity&#39;s data. If provided, the webhook will only be sent if the entity matches the specified conditions. Pass &lt;strong&gt;null&lt;/strong&gt; to disable filtering and receive all webhooks (for &lt;strong&gt;webhook.update&lt;/strong&gt; method).&lt;/p&gt;&lt;p&gt;The filter accepts a recursive JSON object. The maximum nesting level is &lt;strong&gt;5&lt;/strong&gt;. Exceeding this will result in a validation error. &lt;p&gt;&lt;strong&gt;Strict Rule:&lt;/strong&gt; Each level of the object must contain &lt;strong&gt;exactly one&lt;/strong&gt; of the following keys:&lt;/p&gt;&lt;ul&gt;&lt;li&gt;&lt;strong&gt;Logical Groups:&lt;/strong&gt;&lt;code&gt;and&lt;/code&gt;, &lt;code&gt;or&lt;/code&gt;, &lt;code&gt;not&lt;/code&gt; (used to combine multiple conditions).&lt;/li&gt;&lt;li&gt;&lt;strong&gt;Condition:&lt;/strong&gt;&lt;code&gt;field&lt;/code&gt;, &lt;code&gt;operator&lt;/code&gt; and &lt;code&gt;value&lt;/code&gt;&lt;/li&gt;&lt;/ul&gt;&lt;br/&gt;&lt;p&gt;&lt;strong&gt;Logical Operators:&lt;/strong&gt;&lt;ul&gt;&lt;li&gt;&lt;code&gt;and&lt;/code&gt;: Accepts an array of objects. All conditions must be true.&lt;/li&gt;&lt;li&gt;&lt;code&gt;or&lt;/code&gt;: Accepts an array of objects. At least one condition must be true.&lt;/li&gt;&lt;li&gt;&lt;code&gt;not&lt;/code&gt;: Accepts a single object. The condition inside must be false.&lt;/li&gt;&lt;/ul&gt;&lt;/p&gt;&lt;br/&gt;&lt;p&gt;&lt;strong&gt;Condition Object:&lt;/strong&gt;&lt;ul&gt;&lt;li&gt;&lt;code&gt;field&lt;/code&gt;: The dot-notation path to the attribute (e.g., id, customer.email, items.price).&lt;/li&gt;&lt;li&gt;&lt;code&gt;operator&lt;/code&gt;: The comparison method (see list below).&lt;/li&gt;&lt;li&gt;&lt;code&gt;value&lt;/code&gt;: The value to compare against.&lt;/li&gt;&lt;/ul&gt;&lt;/p&gt;&lt;br/&gt;&lt;p&gt;The list of available fields for filtering is returned by the &lt;strong&gt;webhook.events&lt;/strong&gt; method in the &lt;strong&gt;filterable_fields&lt;/strong&gt; property. The available fields depend on the webhook entity. If the list for a specific webhook is empty, it means that filtering is not supported for this entity or action.&lt;/p&gt;&lt;p&gt;The system validates operators against the field type defined in the entity schema.&lt;/p&gt;&lt;p&gt;&lt;ul&gt;&lt;li&gt;&lt;strong&gt;string:&lt;/strong&gt;&lt;code&gt;eq&lt;/code&gt;, &lt;code&gt;neq&lt;/code&gt;, &lt;code&gt;in&lt;/code&gt;, &lt;code&gt;not_in&lt;/code&gt;, &lt;code&gt;like&lt;/code&gt;, &lt;code&gt;not_like&lt;/code&gt;.&lt;/li&gt;&lt;li&gt;&lt;strong&gt;integer:&lt;/strong&gt;&lt;code&gt;eq&lt;/code&gt;, &lt;code&gt;neq&lt;/code&gt;, &lt;code&gt;gt&lt;/code&gt;, &lt;code&gt;gte&lt;/code&gt;, &lt;code&gt;lt&lt;/code&gt;, &lt;code&gt;lte&lt;/code&gt;.&lt;/li&gt;&lt;li&gt;&lt;strong&gt;number:&lt;/strong&gt;&lt;code&gt;eq&lt;/code&gt;, &lt;code&gt;neq&lt;/code&gt;, &lt;code&gt;gt&lt;/code&gt;, &lt;code&gt;gte&lt;/code&gt;, &lt;code&gt;lt&lt;/code&gt;, &lt;code&gt;lte&lt;/code&gt;.&lt;/li&gt;&lt;li&gt;&lt;strong&gt;boolean:&lt;/strong&gt;&lt;code&gt;eq&lt;/code&gt;, &lt;code&gt;neq&lt;/code&gt;.&lt;/li&gt;&lt;li&gt;&lt;strong&gt;date:&lt;/strong&gt;&lt;code&gt;eq&lt;/code&gt;, &lt;code&gt;neq&lt;/code&gt;, &lt;code&gt;gt&lt;/code&gt;, &lt;code&gt;gte&lt;/code&gt;, &lt;code&gt;lt&lt;/code&gt;, &lt;code&gt;lte&lt;/code&gt;.&lt;/li&gt;&lt;li&gt;&lt;strong&gt;array:&lt;/strong&gt;&lt;code&gt;eq&lt;/code&gt;, &lt;code&gt;neq&lt;/code&gt;, &lt;code&gt;in&lt;/code&gt;, &lt;code&gt;not_in&lt;/code&gt;.&lt;/li&gt;&lt;/ul&gt;&lt;/p&gt;&lt;br/&gt;&lt;p&gt;&lt;strong&gt;Operators:&lt;/strong&gt;&lt;ul&gt;&lt;li&gt;&lt;code&gt;eq&lt;/code&gt;: Exact match. The field value must match the specified value exactly (case sensitive). For array-type fields, compares two arrays regardless of the element order.&lt;/li&gt;&lt;li&gt;&lt;code&gt;neq&lt;/code&gt;: Not equal. The field value must be different from the specified value.&lt;/li&gt;&lt;li&gt;&lt;code&gt;like&lt;/code&gt;: Partial match (SQL-style). Case-insensitive. Supported wildcards:&lt;ul&gt;&lt;li&gt;&lt;strong&gt;%&lt;/strong&gt; - matches any sequence of characters (0 or more).&lt;/li&gt;&lt;li&gt;&lt;strong&gt;_&lt;/strong&gt; - matches exactly one character.&lt;/li&gt;&lt;/ul&gt; If the value contains the &#39;%&#39; or &#39;_&#39; characters, they must be escaped using a backslash (&#39;\\&#39;). For example, if the string is &#39;demo_&#39;, the value should be specified as &#39;demo\\_&#39;.&lt;/li&gt;&lt;li&gt;&lt;code&gt;not_like&lt;/code&gt;: Inverted Partial Match. Ensures the pattern does not match.&lt;/li&gt;&lt;li&gt;&lt;code&gt;in&lt;/code&gt;: In List. The field value must match one of the values in the provided array. For array-type fields, the condition is considered valid if at least one element from the provided list exists in the array.&lt;/li&gt;&lt;li&gt;&lt;code&gt;not_in&lt;/code&gt;: Not In List. The field value must not be present in the provided array. For array-type fields, the condition is considered valid if none of the elements in the array exist in the provided list.&lt;/li&gt;&lt;li&gt;&lt;code&gt;lt&lt;/code&gt;: Less (Strictly Less). The value is strictly less than value.&lt;/li&gt;&lt;li&gt;&lt;code&gt;lte&lt;/code&gt;: Less than or equal to.&lt;/li&gt;&lt;li&gt;&lt;code&gt;gt&lt;/code&gt;: Strictly Greater. The value is strictly greater than value.&lt;/li&gt;&lt;li&gt;&lt;code&gt;gte&lt;/code&gt;: Greater than or equal to.&lt;/li&gt;&lt;/ul&gt;&lt;/p&gt;&lt;br/&gt;&lt;p&gt;&lt;strong&gt;Please note: the filter is universal for all platforms; however, for some platforms, certain fields may never be populated. Before defining a filter, make sure that the field is populated for the specific platform. Otherwise, a webhook configured with such a filter may never be triggered.&lt;/strong&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**IdempotencyKey** | Pointer to **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | [optional] 

## Methods

### NewWebhookCreate

`func NewWebhookCreate(entity string, action string, ) *WebhookCreate`

NewWebhookCreate instantiates a new WebhookCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookCreateWithDefaults

`func NewWebhookCreateWithDefaults() *WebhookCreate`

NewWebhookCreateWithDefaults instantiates a new WebhookCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *WebhookCreate) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *WebhookCreate) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *WebhookCreate) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetAction

`func (o *WebhookCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WebhookCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WebhookCreate) SetAction(v string)`

SetAction sets Action field to given value.


### GetCallback

`func (o *WebhookCreate) GetCallback() string`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *WebhookCreate) GetCallbackOk() (*string, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *WebhookCreate) SetCallback(v string)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *WebhookCreate) HasCallback() bool`

HasCallback returns a boolean if a field has been set.

### GetLabel

`func (o *WebhookCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WebhookCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WebhookCreate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WebhookCreate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetFields

`func (o *WebhookCreate) GetFields() string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WebhookCreate) GetFieldsOk() (*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WebhookCreate) SetFields(v string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WebhookCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetResponseFields

`func (o *WebhookCreate) GetResponseFields() string`

GetResponseFields returns the ResponseFields field if non-nil, zero value otherwise.

### GetResponseFieldsOk

`func (o *WebhookCreate) GetResponseFieldsOk() (*string, bool)`

GetResponseFieldsOk returns a tuple with the ResponseFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFields

`func (o *WebhookCreate) SetResponseFields(v string)`

SetResponseFields sets ResponseFields field to given value.

### HasResponseFields

`func (o *WebhookCreate) HasResponseFields() bool`

HasResponseFields returns a boolean if a field has been set.

### GetActive

`func (o *WebhookCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLangId

`func (o *WebhookCreate) GetLangId() string`

GetLangId returns the LangId field if non-nil, zero value otherwise.

### GetLangIdOk

`func (o *WebhookCreate) GetLangIdOk() (*string, bool)`

GetLangIdOk returns a tuple with the LangId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangId

`func (o *WebhookCreate) SetLangId(v string)`

SetLangId sets LangId field to given value.

### HasLangId

`func (o *WebhookCreate) HasLangId() bool`

HasLangId returns a boolean if a field has been set.

### GetStoreId

`func (o *WebhookCreate) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *WebhookCreate) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *WebhookCreate) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *WebhookCreate) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetFilteringConditions

`func (o *WebhookCreate) GetFilteringConditions() ParamDefinitionFilteringConditionsFilterCondition`

GetFilteringConditions returns the FilteringConditions field if non-nil, zero value otherwise.

### GetFilteringConditionsOk

`func (o *WebhookCreate) GetFilteringConditionsOk() (*ParamDefinitionFilteringConditionsFilterCondition, bool)`

GetFilteringConditionsOk returns a tuple with the FilteringConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteringConditions

`func (o *WebhookCreate) SetFilteringConditions(v ParamDefinitionFilteringConditionsFilterCondition)`

SetFilteringConditions sets FilteringConditions field to given value.

### HasFilteringConditions

`func (o *WebhookCreate) HasFilteringConditions() bool`

HasFilteringConditions returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *WebhookCreate) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *WebhookCreate) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *WebhookCreate) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *WebhookCreate) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


