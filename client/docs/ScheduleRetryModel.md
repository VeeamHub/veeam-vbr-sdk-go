# ScheduleRetryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | If *true*, retry options are enabled. | [optional] [default to false]
**RetryCount** | Pointer to **int32** | Number of retries set for the job. Must be greater than zero. | [optional] [default to 3]
**AwaitMinutes** | Pointer to **int32** | Time interval between job retries in minutes. Must be greater than zero. | [optional] [default to 10]

## Methods

### NewScheduleRetryModel

`func NewScheduleRetryModel() *ScheduleRetryModel`

NewScheduleRetryModel instantiates a new ScheduleRetryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleRetryModelWithDefaults

`func NewScheduleRetryModelWithDefaults() *ScheduleRetryModel`

NewScheduleRetryModelWithDefaults instantiates a new ScheduleRetryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ScheduleRetryModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ScheduleRetryModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ScheduleRetryModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ScheduleRetryModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetRetryCount

`func (o *ScheduleRetryModel) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ScheduleRetryModel) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ScheduleRetryModel) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ScheduleRetryModel) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetAwaitMinutes

`func (o *ScheduleRetryModel) GetAwaitMinutes() int32`

GetAwaitMinutes returns the AwaitMinutes field if non-nil, zero value otherwise.

### GetAwaitMinutesOk

`func (o *ScheduleRetryModel) GetAwaitMinutesOk() (*int32, bool)`

GetAwaitMinutesOk returns a tuple with the AwaitMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitMinutes

`func (o *ScheduleRetryModel) SetAwaitMinutes(v int32)`

SetAwaitMinutes sets AwaitMinutes field to given value.

### HasAwaitMinutes

`func (o *ScheduleRetryModel) HasAwaitMinutes() bool`

HasAwaitMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


