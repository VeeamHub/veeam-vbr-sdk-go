# ConfigBackupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, configuration backup is enabled. | 
**BackupRepositoryId** | **string** | ID of the backup repository on which the configuration backup is stored. | 
**RestorePointsToKeep** | **int32** | Number of restore points to keep in the backup repository. | 
**Notifications** | [**ConfigBackupNotificationsModel**](ConfigBackupNotificationsModel.md) |  | 
**Schedule** | [**ConfigBackupScheduleModel**](ConfigBackupScheduleModel.md) |  | 
**LastSuccessfulBackup** | [**ConfigBackupLastSuccessfulModel**](ConfigBackupLastSuccessfulModel.md) |  | 
**Encryption** | [**ConfigBackupEncryptionModel**](ConfigBackupEncryptionModel.md) |  | 

## Methods

### NewConfigBackupModel

`func NewConfigBackupModel(isEnabled bool, backupRepositoryId string, restorePointsToKeep int32, notifications ConfigBackupNotificationsModel, schedule ConfigBackupScheduleModel, lastSuccessfulBackup ConfigBackupLastSuccessfulModel, encryption ConfigBackupEncryptionModel, ) *ConfigBackupModel`

NewConfigBackupModel instantiates a new ConfigBackupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigBackupModelWithDefaults

`func NewConfigBackupModelWithDefaults() *ConfigBackupModel`

NewConfigBackupModelWithDefaults instantiates a new ConfigBackupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ConfigBackupModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ConfigBackupModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ConfigBackupModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetBackupRepositoryId

`func (o *ConfigBackupModel) GetBackupRepositoryId() string`

GetBackupRepositoryId returns the BackupRepositoryId field if non-nil, zero value otherwise.

### GetBackupRepositoryIdOk

`func (o *ConfigBackupModel) GetBackupRepositoryIdOk() (*string, bool)`

GetBackupRepositoryIdOk returns a tuple with the BackupRepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRepositoryId

`func (o *ConfigBackupModel) SetBackupRepositoryId(v string)`

SetBackupRepositoryId sets BackupRepositoryId field to given value.


### GetRestorePointsToKeep

`func (o *ConfigBackupModel) GetRestorePointsToKeep() int32`

GetRestorePointsToKeep returns the RestorePointsToKeep field if non-nil, zero value otherwise.

### GetRestorePointsToKeepOk

`func (o *ConfigBackupModel) GetRestorePointsToKeepOk() (*int32, bool)`

GetRestorePointsToKeepOk returns a tuple with the RestorePointsToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorePointsToKeep

`func (o *ConfigBackupModel) SetRestorePointsToKeep(v int32)`

SetRestorePointsToKeep sets RestorePointsToKeep field to given value.


### GetNotifications

`func (o *ConfigBackupModel) GetNotifications() ConfigBackupNotificationsModel`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ConfigBackupModel) GetNotificationsOk() (*ConfigBackupNotificationsModel, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ConfigBackupModel) SetNotifications(v ConfigBackupNotificationsModel)`

SetNotifications sets Notifications field to given value.


### GetSchedule

`func (o *ConfigBackupModel) GetSchedule() ConfigBackupScheduleModel`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ConfigBackupModel) GetScheduleOk() (*ConfigBackupScheduleModel, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ConfigBackupModel) SetSchedule(v ConfigBackupScheduleModel)`

SetSchedule sets Schedule field to given value.


### GetLastSuccessfulBackup

`func (o *ConfigBackupModel) GetLastSuccessfulBackup() ConfigBackupLastSuccessfulModel`

GetLastSuccessfulBackup returns the LastSuccessfulBackup field if non-nil, zero value otherwise.

### GetLastSuccessfulBackupOk

`func (o *ConfigBackupModel) GetLastSuccessfulBackupOk() (*ConfigBackupLastSuccessfulModel, bool)`

GetLastSuccessfulBackupOk returns a tuple with the LastSuccessfulBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulBackup

`func (o *ConfigBackupModel) SetLastSuccessfulBackup(v ConfigBackupLastSuccessfulModel)`

SetLastSuccessfulBackup sets LastSuccessfulBackup field to given value.


### GetEncryption

`func (o *ConfigBackupModel) GetEncryption() ConfigBackupEncryptionModel`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *ConfigBackupModel) GetEncryptionOk() (*ConfigBackupEncryptionModel, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *ConfigBackupModel) SetEncryption(v ConfigBackupEncryptionModel)`

SetEncryption sets Encryption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


