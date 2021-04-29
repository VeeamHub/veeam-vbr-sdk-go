# ScheduleAfterThisJobModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, job chaining is enabled. | 
**JobName** | Pointer to **string** | Name of the preceding job. | [optional] 

## Methods

### NewScheduleAfterThisJobModel

`func NewScheduleAfterThisJobModel(isEnabled bool, ) *ScheduleAfterThisJobModel`

NewScheduleAfterThisJobModel instantiates a new ScheduleAfterThisJobModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleAfterThisJobModelWithDefaults

`func NewScheduleAfterThisJobModelWithDefaults() *ScheduleAfterThisJobModel`

NewScheduleAfterThisJobModelWithDefaults instantiates a new ScheduleAfterThisJobModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ScheduleAfterThisJobModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ScheduleAfterThisJobModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ScheduleAfterThisJobModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetJobName

`func (o *ScheduleAfterThisJobModel) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *ScheduleAfterThisJobModel) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *ScheduleAfterThisJobModel) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *ScheduleAfterThisJobModel) HasJobName() bool`

HasJobName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


