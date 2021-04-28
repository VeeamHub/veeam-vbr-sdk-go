# FullBackupMaintenanceDefragmentAndCompactModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, Veeam Backup &amp; Replication compacts full backup files. | 
**Weekly** | Pointer to [**AdvancedStorageScheduleWeeklyModel**](AdvancedStorageScheduleWeeklyModel.md) |  | [optional] 
**Monthly** | Pointer to [**AdvancedStorageScheduleMonthlyModel**](AdvancedStorageScheduleMonthlyModel.md) |  | [optional] 

## Methods

### NewFullBackupMaintenanceDefragmentAndCompactModel

`func NewFullBackupMaintenanceDefragmentAndCompactModel(isEnabled bool, ) *FullBackupMaintenanceDefragmentAndCompactModel`

NewFullBackupMaintenanceDefragmentAndCompactModel instantiates a new FullBackupMaintenanceDefragmentAndCompactModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullBackupMaintenanceDefragmentAndCompactModelWithDefaults

`func NewFullBackupMaintenanceDefragmentAndCompactModelWithDefaults() *FullBackupMaintenanceDefragmentAndCompactModel`

NewFullBackupMaintenanceDefragmentAndCompactModelWithDefaults instantiates a new FullBackupMaintenanceDefragmentAndCompactModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *FullBackupMaintenanceDefragmentAndCompactModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FullBackupMaintenanceDefragmentAndCompactModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FullBackupMaintenanceDefragmentAndCompactModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetWeekly

`func (o *FullBackupMaintenanceDefragmentAndCompactModel) GetWeekly() AdvancedStorageScheduleWeeklyModel`

GetWeekly returns the Weekly field if non-nil, zero value otherwise.

### GetWeeklyOk

`func (o *FullBackupMaintenanceDefragmentAndCompactModel) GetWeeklyOk() (*AdvancedStorageScheduleWeeklyModel, bool)`

GetWeeklyOk returns a tuple with the Weekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekly

`func (o *FullBackupMaintenanceDefragmentAndCompactModel) SetWeekly(v AdvancedStorageScheduleWeeklyModel)`

SetWeekly sets Weekly field to given value.

### HasWeekly

`func (o *FullBackupMaintenanceDefragmentAndCompactModel) HasWeekly() bool`

HasWeekly returns a boolean if a field has been set.

### GetMonthly

`func (o *FullBackupMaintenanceDefragmentAndCompactModel) GetMonthly() AdvancedStorageScheduleMonthlyModel`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *FullBackupMaintenanceDefragmentAndCompactModel) GetMonthlyOk() (*AdvancedStorageScheduleMonthlyModel, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *FullBackupMaintenanceDefragmentAndCompactModel) SetMonthly(v AdvancedStorageScheduleMonthlyModel)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *FullBackupMaintenanceDefragmentAndCompactModel) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


