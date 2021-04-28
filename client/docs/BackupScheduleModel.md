# BackupScheduleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunAutomatically** | **bool** | If *true*, job scheduling is enabled. | [default to false]
**Daily** | Pointer to [**ScheduleDailyModel**](ScheduleDailyModel.md) |  | [optional] 
**Monthly** | Pointer to [**ScheduleMonthlyModel**](ScheduleMonthlyModel.md) |  | [optional] 
**Periodically** | Pointer to [**SchedulePeriodicallyModel**](SchedulePeriodicallyModel.md) |  | [optional] 
**Continuously** | Pointer to [**ScheduleBackupWindowModel**](ScheduleBackupWindowModel.md) |  | [optional] 
**AfterThisJob** | Pointer to [**ScheduleAfterThisJobModel**](ScheduleAfterThisJobModel.md) |  | [optional] 
**Retry** | Pointer to [**ScheduleRetryModel**](ScheduleRetryModel.md) |  | [optional] 
**BackupWindow** | Pointer to [**ScheduleBackupWindowModel**](ScheduleBackupWindowModel.md) |  | [optional] 

## Methods

### NewBackupScheduleModel

`func NewBackupScheduleModel(runAutomatically bool, ) *BackupScheduleModel`

NewBackupScheduleModel instantiates a new BackupScheduleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleModelWithDefaults

`func NewBackupScheduleModelWithDefaults() *BackupScheduleModel`

NewBackupScheduleModelWithDefaults instantiates a new BackupScheduleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunAutomatically

`func (o *BackupScheduleModel) GetRunAutomatically() bool`

GetRunAutomatically returns the RunAutomatically field if non-nil, zero value otherwise.

### GetRunAutomaticallyOk

`func (o *BackupScheduleModel) GetRunAutomaticallyOk() (*bool, bool)`

GetRunAutomaticallyOk returns a tuple with the RunAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAutomatically

`func (o *BackupScheduleModel) SetRunAutomatically(v bool)`

SetRunAutomatically sets RunAutomatically field to given value.


### GetDaily

`func (o *BackupScheduleModel) GetDaily() ScheduleDailyModel`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *BackupScheduleModel) GetDailyOk() (*ScheduleDailyModel, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *BackupScheduleModel) SetDaily(v ScheduleDailyModel)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *BackupScheduleModel) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetMonthly

`func (o *BackupScheduleModel) GetMonthly() ScheduleMonthlyModel`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *BackupScheduleModel) GetMonthlyOk() (*ScheduleMonthlyModel, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *BackupScheduleModel) SetMonthly(v ScheduleMonthlyModel)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *BackupScheduleModel) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.

### GetPeriodically

`func (o *BackupScheduleModel) GetPeriodically() SchedulePeriodicallyModel`

GetPeriodically returns the Periodically field if non-nil, zero value otherwise.

### GetPeriodicallyOk

`func (o *BackupScheduleModel) GetPeriodicallyOk() (*SchedulePeriodicallyModel, bool)`

GetPeriodicallyOk returns a tuple with the Periodically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodically

`func (o *BackupScheduleModel) SetPeriodically(v SchedulePeriodicallyModel)`

SetPeriodically sets Periodically field to given value.

### HasPeriodically

`func (o *BackupScheduleModel) HasPeriodically() bool`

HasPeriodically returns a boolean if a field has been set.

### GetContinuously

`func (o *BackupScheduleModel) GetContinuously() ScheduleBackupWindowModel`

GetContinuously returns the Continuously field if non-nil, zero value otherwise.

### GetContinuouslyOk

`func (o *BackupScheduleModel) GetContinuouslyOk() (*ScheduleBackupWindowModel, bool)`

GetContinuouslyOk returns a tuple with the Continuously field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuously

`func (o *BackupScheduleModel) SetContinuously(v ScheduleBackupWindowModel)`

SetContinuously sets Continuously field to given value.

### HasContinuously

`func (o *BackupScheduleModel) HasContinuously() bool`

HasContinuously returns a boolean if a field has been set.

### GetAfterThisJob

`func (o *BackupScheduleModel) GetAfterThisJob() ScheduleAfterThisJobModel`

GetAfterThisJob returns the AfterThisJob field if non-nil, zero value otherwise.

### GetAfterThisJobOk

`func (o *BackupScheduleModel) GetAfterThisJobOk() (*ScheduleAfterThisJobModel, bool)`

GetAfterThisJobOk returns a tuple with the AfterThisJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterThisJob

`func (o *BackupScheduleModel) SetAfterThisJob(v ScheduleAfterThisJobModel)`

SetAfterThisJob sets AfterThisJob field to given value.

### HasAfterThisJob

`func (o *BackupScheduleModel) HasAfterThisJob() bool`

HasAfterThisJob returns a boolean if a field has been set.

### GetRetry

`func (o *BackupScheduleModel) GetRetry() ScheduleRetryModel`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *BackupScheduleModel) GetRetryOk() (*ScheduleRetryModel, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *BackupScheduleModel) SetRetry(v ScheduleRetryModel)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *BackupScheduleModel) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetBackupWindow

`func (o *BackupScheduleModel) GetBackupWindow() ScheduleBackupWindowModel`

GetBackupWindow returns the BackupWindow field if non-nil, zero value otherwise.

### GetBackupWindowOk

`func (o *BackupScheduleModel) GetBackupWindowOk() (*ScheduleBackupWindowModel, bool)`

GetBackupWindowOk returns a tuple with the BackupWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupWindow

`func (o *BackupScheduleModel) SetBackupWindow(v ScheduleBackupWindowModel)`

SetBackupWindow sets BackupWindow field to given value.

### HasBackupWindow

`func (o *BackupScheduleModel) HasBackupWindow() bool`

HasBackupWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


