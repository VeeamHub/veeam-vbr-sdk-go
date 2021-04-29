# BackupHealthCheckSettingsModels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, the health check is enabled. | 
**Weekly** | Pointer to [**AdvancedStorageScheduleWeeklyModel**](AdvancedStorageScheduleWeeklyModel.md) |  | [optional] 
**Monthly** | Pointer to [**AdvancedStorageScheduleMonthlyModel**](AdvancedStorageScheduleMonthlyModel.md) |  | [optional] 

## Methods

### NewBackupHealthCheckSettingsModels

`func NewBackupHealthCheckSettingsModels(isEnabled bool, ) *BackupHealthCheckSettingsModels`

NewBackupHealthCheckSettingsModels instantiates a new BackupHealthCheckSettingsModels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupHealthCheckSettingsModelsWithDefaults

`func NewBackupHealthCheckSettingsModelsWithDefaults() *BackupHealthCheckSettingsModels`

NewBackupHealthCheckSettingsModelsWithDefaults instantiates a new BackupHealthCheckSettingsModels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *BackupHealthCheckSettingsModels) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *BackupHealthCheckSettingsModels) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *BackupHealthCheckSettingsModels) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetWeekly

`func (o *BackupHealthCheckSettingsModels) GetWeekly() AdvancedStorageScheduleWeeklyModel`

GetWeekly returns the Weekly field if non-nil, zero value otherwise.

### GetWeeklyOk

`func (o *BackupHealthCheckSettingsModels) GetWeeklyOk() (*AdvancedStorageScheduleWeeklyModel, bool)`

GetWeeklyOk returns a tuple with the Weekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekly

`func (o *BackupHealthCheckSettingsModels) SetWeekly(v AdvancedStorageScheduleWeeklyModel)`

SetWeekly sets Weekly field to given value.

### HasWeekly

`func (o *BackupHealthCheckSettingsModels) HasWeekly() bool`

HasWeekly returns a boolean if a field has been set.

### GetMonthly

`func (o *BackupHealthCheckSettingsModels) GetMonthly() AdvancedStorageScheduleMonthlyModel`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *BackupHealthCheckSettingsModels) GetMonthlyOk() (*AdvancedStorageScheduleMonthlyModel, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *BackupHealthCheckSettingsModels) SetMonthly(v AdvancedStorageScheduleMonthlyModel)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *BackupHealthCheckSettingsModels) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


