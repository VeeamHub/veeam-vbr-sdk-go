# ScheduleDailyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, daily schedule is enabled. | [default to true]
**LocalTime** | Pointer to **string** | Local time when the job must start. | [optional] 
**DailyKind** | Pointer to [**EDailyKinds**](EDailyKinds.md) |  | [optional] 
**Days** | Pointer to [**[]EDayOfWeek**](EDayOfWeek.md) | Days of the week when the job must start. | [optional] 

## Methods

### NewScheduleDailyModel

`func NewScheduleDailyModel(isEnabled bool, ) *ScheduleDailyModel`

NewScheduleDailyModel instantiates a new ScheduleDailyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleDailyModelWithDefaults

`func NewScheduleDailyModelWithDefaults() *ScheduleDailyModel`

NewScheduleDailyModelWithDefaults instantiates a new ScheduleDailyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ScheduleDailyModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ScheduleDailyModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ScheduleDailyModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetLocalTime

`func (o *ScheduleDailyModel) GetLocalTime() string`

GetLocalTime returns the LocalTime field if non-nil, zero value otherwise.

### GetLocalTimeOk

`func (o *ScheduleDailyModel) GetLocalTimeOk() (*string, bool)`

GetLocalTimeOk returns a tuple with the LocalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTime

`func (o *ScheduleDailyModel) SetLocalTime(v string)`

SetLocalTime sets LocalTime field to given value.

### HasLocalTime

`func (o *ScheduleDailyModel) HasLocalTime() bool`

HasLocalTime returns a boolean if a field has been set.

### GetDailyKind

`func (o *ScheduleDailyModel) GetDailyKind() EDailyKinds`

GetDailyKind returns the DailyKind field if non-nil, zero value otherwise.

### GetDailyKindOk

`func (o *ScheduleDailyModel) GetDailyKindOk() (*EDailyKinds, bool)`

GetDailyKindOk returns a tuple with the DailyKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyKind

`func (o *ScheduleDailyModel) SetDailyKind(v EDailyKinds)`

SetDailyKind sets DailyKind field to given value.

### HasDailyKind

`func (o *ScheduleDailyModel) HasDailyKind() bool`

HasDailyKind returns a boolean if a field has been set.

### GetDays

`func (o *ScheduleDailyModel) GetDays() []EDayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ScheduleDailyModel) GetDaysOk() (*[]EDayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ScheduleDailyModel) SetDays(v []EDayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *ScheduleDailyModel) HasDays() bool`

HasDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


