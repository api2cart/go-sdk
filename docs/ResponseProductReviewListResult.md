# ResponseProductReviewListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Reviews** | Pointer to [**[]ProductReview**](ProductReview.md) |  | [optional] 
**AdditionalFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResponseProductReviewListResult

`func NewResponseProductReviewListResult() *ResponseProductReviewListResult`

NewResponseProductReviewListResult instantiates a new ResponseProductReviewListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseProductReviewListResultWithDefaults

`func NewResponseProductReviewListResultWithDefaults() *ResponseProductReviewListResult`

NewResponseProductReviewListResultWithDefaults instantiates a new ResponseProductReviewListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ResponseProductReviewListResult) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResponseProductReviewListResult) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResponseProductReviewListResult) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ResponseProductReviewListResult) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetReviews

`func (o *ResponseProductReviewListResult) GetReviews() []ProductReview`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *ResponseProductReviewListResult) GetReviewsOk() (*[]ProductReview, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *ResponseProductReviewListResult) SetReviews(v []ProductReview)`

SetReviews sets Reviews field to given value.

### HasReviews

`func (o *ResponseProductReviewListResult) HasReviews() bool`

HasReviews returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *ResponseProductReviewListResult) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *ResponseProductReviewListResult) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *ResponseProductReviewListResult) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *ResponseProductReviewListResult) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### GetCustomFields

`func (o *ResponseProductReviewListResult) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ResponseProductReviewListResult) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ResponseProductReviewListResult) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ResponseProductReviewListResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


