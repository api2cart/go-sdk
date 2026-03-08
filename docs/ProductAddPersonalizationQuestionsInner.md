# ProductAddPersonalizationQuestionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionText** | **string** |  | 
**Instructions** | Pointer to **string** |  | [optional] 
**QuestionType** | **string** |  | 
**Required** | **bool** |  | 
**MaxAllowedCharacters** | Pointer to **NullableInt32** |  | [optional] 
**MaxAllowedFiles** | Pointer to **NullableInt32** |  | [optional] 
**Options** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProductAddPersonalizationQuestionsInner

`func NewProductAddPersonalizationQuestionsInner(questionText string, questionType string, required bool, ) *ProductAddPersonalizationQuestionsInner`

NewProductAddPersonalizationQuestionsInner instantiates a new ProductAddPersonalizationQuestionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAddPersonalizationQuestionsInnerWithDefaults

`func NewProductAddPersonalizationQuestionsInnerWithDefaults() *ProductAddPersonalizationQuestionsInner`

NewProductAddPersonalizationQuestionsInnerWithDefaults instantiates a new ProductAddPersonalizationQuestionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionText

`func (o *ProductAddPersonalizationQuestionsInner) GetQuestionText() string`

GetQuestionText returns the QuestionText field if non-nil, zero value otherwise.

### GetQuestionTextOk

`func (o *ProductAddPersonalizationQuestionsInner) GetQuestionTextOk() (*string, bool)`

GetQuestionTextOk returns a tuple with the QuestionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionText

`func (o *ProductAddPersonalizationQuestionsInner) SetQuestionText(v string)`

SetQuestionText sets QuestionText field to given value.


### GetInstructions

`func (o *ProductAddPersonalizationQuestionsInner) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ProductAddPersonalizationQuestionsInner) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ProductAddPersonalizationQuestionsInner) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ProductAddPersonalizationQuestionsInner) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetQuestionType

`func (o *ProductAddPersonalizationQuestionsInner) GetQuestionType() string`

GetQuestionType returns the QuestionType field if non-nil, zero value otherwise.

### GetQuestionTypeOk

`func (o *ProductAddPersonalizationQuestionsInner) GetQuestionTypeOk() (*string, bool)`

GetQuestionTypeOk returns a tuple with the QuestionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionType

`func (o *ProductAddPersonalizationQuestionsInner) SetQuestionType(v string)`

SetQuestionType sets QuestionType field to given value.


### GetRequired

`func (o *ProductAddPersonalizationQuestionsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ProductAddPersonalizationQuestionsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ProductAddPersonalizationQuestionsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetMaxAllowedCharacters

`func (o *ProductAddPersonalizationQuestionsInner) GetMaxAllowedCharacters() int32`

GetMaxAllowedCharacters returns the MaxAllowedCharacters field if non-nil, zero value otherwise.

### GetMaxAllowedCharactersOk

`func (o *ProductAddPersonalizationQuestionsInner) GetMaxAllowedCharactersOk() (*int32, bool)`

GetMaxAllowedCharactersOk returns a tuple with the MaxAllowedCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedCharacters

`func (o *ProductAddPersonalizationQuestionsInner) SetMaxAllowedCharacters(v int32)`

SetMaxAllowedCharacters sets MaxAllowedCharacters field to given value.

### HasMaxAllowedCharacters

`func (o *ProductAddPersonalizationQuestionsInner) HasMaxAllowedCharacters() bool`

HasMaxAllowedCharacters returns a boolean if a field has been set.

### SetMaxAllowedCharactersNil

`func (o *ProductAddPersonalizationQuestionsInner) SetMaxAllowedCharactersNil(b bool)`

 SetMaxAllowedCharactersNil sets the value for MaxAllowedCharacters to be an explicit nil

### UnsetMaxAllowedCharacters
`func (o *ProductAddPersonalizationQuestionsInner) UnsetMaxAllowedCharacters()`

UnsetMaxAllowedCharacters ensures that no value is present for MaxAllowedCharacters, not even an explicit nil
### GetMaxAllowedFiles

`func (o *ProductAddPersonalizationQuestionsInner) GetMaxAllowedFiles() int32`

GetMaxAllowedFiles returns the MaxAllowedFiles field if non-nil, zero value otherwise.

### GetMaxAllowedFilesOk

`func (o *ProductAddPersonalizationQuestionsInner) GetMaxAllowedFilesOk() (*int32, bool)`

GetMaxAllowedFilesOk returns a tuple with the MaxAllowedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedFiles

`func (o *ProductAddPersonalizationQuestionsInner) SetMaxAllowedFiles(v int32)`

SetMaxAllowedFiles sets MaxAllowedFiles field to given value.

### HasMaxAllowedFiles

`func (o *ProductAddPersonalizationQuestionsInner) HasMaxAllowedFiles() bool`

HasMaxAllowedFiles returns a boolean if a field has been set.

### SetMaxAllowedFilesNil

`func (o *ProductAddPersonalizationQuestionsInner) SetMaxAllowedFilesNil(b bool)`

 SetMaxAllowedFilesNil sets the value for MaxAllowedFiles to be an explicit nil

### UnsetMaxAllowedFiles
`func (o *ProductAddPersonalizationQuestionsInner) UnsetMaxAllowedFiles()`

UnsetMaxAllowedFiles ensures that no value is present for MaxAllowedFiles, not even an explicit nil
### GetOptions

`func (o *ProductAddPersonalizationQuestionsInner) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ProductAddPersonalizationQuestionsInner) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ProductAddPersonalizationQuestionsInner) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ProductAddPersonalizationQuestionsInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ProductAddPersonalizationQuestionsInner) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ProductAddPersonalizationQuestionsInner) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


