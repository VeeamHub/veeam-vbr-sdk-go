# SchedulePeriodicallyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, periodic schedule is enabled. | [default to false]
**PeriodicallyKind** | Pointer to [**EPeriodicallyKinds**](EPeriodicallyKinds.md) |  | [optional] 
**Frequency** | Pointer to **int32** | Number of time units that defines the time interval. | [optional] 
**BackupWindow** | Pointer to [**BackupWindowSettingModel**](BackupWindowSettingModel.md) |  | [optional] 
**StartTimeWithinAnHour** | Pointer to **int32** |  | [optional] 

## Methods

### NewSchedulePeriodicallyModel

`func NewSchedulePeriodicallyModel(isEnabled bool, ) *SchedulePeriodicallyModel`

NewSchedulePeriodicallyModel instantiates a new SchedulePeriodicallyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePeriodicallyModelWithDefaults

`func NewSchedulePeriodicallyModelWithDefaults() *SchedulePeriodicallyModel`

NewSchedulePeriodicallyModelWithDefaults instantiates a new SchedulePeriodicallyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *SchedulePeriodicallyModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SchedulePeriodicallyModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SchedulePeriodicallyModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetPeriodicallyKind

`func (o *SchedulePeriodicallyModel) GetPeriodicallyKind() EPeriodicallyKinds`

GetPeriodicallyKind returns the PeriodicallyKind field if non-nil, zero value otherwise.

### GetPeriodicallyKindOk

`func (o *SchedulePeriodicallyModel) GetPeriodicallyKindOk() (*EPeriodicallyKinds, bool)`

GetPeriodicallyKindOk returns a tuple with the PeriodicallyKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicallyKind

`func (o *SchedulePeriodicallyModel) SetPeriodicallyKind(v EPeriodicallyKinds)`

SetPeriodicallyKind sets PeriodicallyKind field to given value.

### HasPeriodicallyKind

`func (o *SchedulePeriodicallyModel) HasPeriodicallyKind() bool`

HasPeriodicallyKind returns a boolean if a field has been set.

### GetFrequency

`func (o *SchedulePeriodicallyModel) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *SchedulePeriodicallyModel) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *SchedulePeriodicallyModel) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *SchedulePeriodicallyModel) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetBackupWindow

`func (o *SchedulePeriodicallyModel) GetBackupWindow() BackupWindowSettingModel`

GetBackupWindow returns the BackupWindow field if non-nil, zero value otherwise.

### GetBackupWindowOk

`func (o *SchedulePeriodicallyModel) GetBackupWindowOk() (*BackupWindowSettingModel, bool)`

GetBackupWindowOk returns a tuple with the BackupWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupWindow

`func (o *SchedulePeriodicallyModel) SetBackupWindow(v BackupWindowSettingModel)`

SetBackupWindow sets BackupWindow field to given value.

### HasBackupWindow

`func (o *SchedulePeriodicallyModel) HasBackupWindow() bool`

HasBackupWindow returns a boolean if a field has been set.

### GetStartTimeWithinAnHour

`func (o *SchedulePeriodicallyModel) GetStartTimeWithinAnHour() int32`

GetStartTimeWithinAnHour returns the StartTimeWithinAnHour field if non-nil, zero value otherwise.

### GetStartTimeWithinAnHourOk

`func (o *SchedulePeriodicallyModel) GetStartTimeWithinAnHourOk() (*int32, bool)`

GetStartTimeWithinAnHourOk returns a tuple with the StartTimeWithinAnHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeWithinAnHour

`func (o *SchedulePeriodicallyModel) SetStartTimeWithinAnHour(v int32)`

SetStartTimeWithinAnHour sets StartTimeWithinAnHour field to given value.

### HasStartTimeWithinAnHour

`func (o *SchedulePeriodicallyModel) HasStartTimeWithinAnHour() bool`

HasStartTimeWithinAnHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


