# ScheduleBackupWindowModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, periodic schedule is enabled. | [default to false]
**BackupWindow** | Pointer to [**BackupWindowSettingModel**](BackupWindowSettingModel.md) |  | [optional] 

## Methods

### NewScheduleBackupWindowModel

`func NewScheduleBackupWindowModel(isEnabled bool, ) *ScheduleBackupWindowModel`

NewScheduleBackupWindowModel instantiates a new ScheduleBackupWindowModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleBackupWindowModelWithDefaults

`func NewScheduleBackupWindowModelWithDefaults() *ScheduleBackupWindowModel`

NewScheduleBackupWindowModelWithDefaults instantiates a new ScheduleBackupWindowModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ScheduleBackupWindowModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ScheduleBackupWindowModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ScheduleBackupWindowModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetBackupWindow

`func (o *ScheduleBackupWindowModel) GetBackupWindow() BackupWindowSettingModel`

GetBackupWindow returns the BackupWindow field if non-nil, zero value otherwise.

### GetBackupWindowOk

`func (o *ScheduleBackupWindowModel) GetBackupWindowOk() (*BackupWindowSettingModel, bool)`

GetBackupWindowOk returns a tuple with the BackupWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupWindow

`func (o *ScheduleBackupWindowModel) SetBackupWindow(v BackupWindowSettingModel)`

SetBackupWindow sets BackupWindow field to given value.

### HasBackupWindow

`func (o *ScheduleBackupWindowModel) HasBackupWindow() bool`

HasBackupWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


