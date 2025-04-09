# ProductAddSpecificsInnerBookingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** |  | 
**Type** | **string** |  | 
**SessionDuration** | Pointer to **int32** |  | [optional] 
**SessionGap** | Pointer to **int32** |  | [optional] 
**SessionsCount** | **int32** |  | 
**TimeStrictValue** | **float32** |  | 
**TimeStrictType** | **string** |  | [default to "days"]
**Availabilities** | [**[]ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner**](ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner.md) |  | 
**Overrides** | Pointer to [**[]ProductAddSpecificsInnerBookingDetailsOverridesInner**](ProductAddSpecificsInnerBookingDetailsOverridesInner.md) |  | [optional] 

## Methods

### NewProductAddSpecificsInnerBookingDetails

`func NewProductAddSpecificsInnerBookingDetails(location string, type_ string, sessionsCount int32, timeStrictValue float32, timeStrictType string, availabilities []ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner, ) *ProductAddSpecificsInnerBookingDetails`

NewProductAddSpecificsInnerBookingDetails instantiates a new ProductAddSpecificsInnerBookingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAddSpecificsInnerBookingDetailsWithDefaults

`func NewProductAddSpecificsInnerBookingDetailsWithDefaults() *ProductAddSpecificsInnerBookingDetails`

NewProductAddSpecificsInnerBookingDetailsWithDefaults instantiates a new ProductAddSpecificsInnerBookingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ProductAddSpecificsInnerBookingDetails) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ProductAddSpecificsInnerBookingDetails) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ProductAddSpecificsInnerBookingDetails) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetType

`func (o *ProductAddSpecificsInnerBookingDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductAddSpecificsInnerBookingDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductAddSpecificsInnerBookingDetails) SetType(v string)`

SetType sets Type field to given value.


### GetSessionDuration

`func (o *ProductAddSpecificsInnerBookingDetails) GetSessionDuration() int32`

GetSessionDuration returns the SessionDuration field if non-nil, zero value otherwise.

### GetSessionDurationOk

`func (o *ProductAddSpecificsInnerBookingDetails) GetSessionDurationOk() (*int32, bool)`

GetSessionDurationOk returns a tuple with the SessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDuration

`func (o *ProductAddSpecificsInnerBookingDetails) SetSessionDuration(v int32)`

SetSessionDuration sets SessionDuration field to given value.

### HasSessionDuration

`func (o *ProductAddSpecificsInnerBookingDetails) HasSessionDuration() bool`

HasSessionDuration returns a boolean if a field has been set.

### GetSessionGap

`func (o *ProductAddSpecificsInnerBookingDetails) GetSessionGap() int32`

GetSessionGap returns the SessionGap field if non-nil, zero value otherwise.

### GetSessionGapOk

`func (o *ProductAddSpecificsInnerBookingDetails) GetSessionGapOk() (*int32, bool)`

GetSessionGapOk returns a tuple with the SessionGap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionGap

`func (o *ProductAddSpecificsInnerBookingDetails) SetSessionGap(v int32)`

SetSessionGap sets SessionGap field to given value.

### HasSessionGap

`func (o *ProductAddSpecificsInnerBookingDetails) HasSessionGap() bool`

HasSessionGap returns a boolean if a field has been set.

### GetSessionsCount

`func (o *ProductAddSpecificsInnerBookingDetails) GetSessionsCount() int32`

GetSessionsCount returns the SessionsCount field if non-nil, zero value otherwise.

### GetSessionsCountOk

`func (o *ProductAddSpecificsInnerBookingDetails) GetSessionsCountOk() (*int32, bool)`

GetSessionsCountOk returns a tuple with the SessionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsCount

`func (o *ProductAddSpecificsInnerBookingDetails) SetSessionsCount(v int32)`

SetSessionsCount sets SessionsCount field to given value.


### GetTimeStrictValue

`func (o *ProductAddSpecificsInnerBookingDetails) GetTimeStrictValue() float32`

GetTimeStrictValue returns the TimeStrictValue field if non-nil, zero value otherwise.

### GetTimeStrictValueOk

`func (o *ProductAddSpecificsInnerBookingDetails) GetTimeStrictValueOk() (*float32, bool)`

GetTimeStrictValueOk returns a tuple with the TimeStrictValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStrictValue

`func (o *ProductAddSpecificsInnerBookingDetails) SetTimeStrictValue(v float32)`

SetTimeStrictValue sets TimeStrictValue field to given value.


### GetTimeStrictType

`func (o *ProductAddSpecificsInnerBookingDetails) GetTimeStrictType() string`

GetTimeStrictType returns the TimeStrictType field if non-nil, zero value otherwise.

### GetTimeStrictTypeOk

`func (o *ProductAddSpecificsInnerBookingDetails) GetTimeStrictTypeOk() (*string, bool)`

GetTimeStrictTypeOk returns a tuple with the TimeStrictType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStrictType

`func (o *ProductAddSpecificsInnerBookingDetails) SetTimeStrictType(v string)`

SetTimeStrictType sets TimeStrictType field to given value.


### GetAvailabilities

`func (o *ProductAddSpecificsInnerBookingDetails) GetAvailabilities() []ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner`

GetAvailabilities returns the Availabilities field if non-nil, zero value otherwise.

### GetAvailabilitiesOk

`func (o *ProductAddSpecificsInnerBookingDetails) GetAvailabilitiesOk() (*[]ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner, bool)`

GetAvailabilitiesOk returns a tuple with the Availabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilities

`func (o *ProductAddSpecificsInnerBookingDetails) SetAvailabilities(v []ProductAddSpecificsInnerBookingDetailsAvailabilitiesInner)`

SetAvailabilities sets Availabilities field to given value.


### GetOverrides

`func (o *ProductAddSpecificsInnerBookingDetails) GetOverrides() []ProductAddSpecificsInnerBookingDetailsOverridesInner`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *ProductAddSpecificsInnerBookingDetails) GetOverridesOk() (*[]ProductAddSpecificsInnerBookingDetailsOverridesInner, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *ProductAddSpecificsInnerBookingDetails) SetOverrides(v []ProductAddSpecificsInnerBookingDetailsOverridesInner)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *ProductAddSpecificsInnerBookingDetails) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


