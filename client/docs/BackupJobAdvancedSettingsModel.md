# BackupJobAdvancedSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupModeType** | [**EBackupModeType**](EBackupModeType.md) |  | 
**SynthenticFulls** | Pointer to [**SyntheticFullSettingsModel**](SyntheticFullSettingsModel.md) |  | [optional] 
**ActiveFulls** | Pointer to [**ActiveFullSettingsModel**](ActiveFullSettingsModel.md) |  | [optional] 
**BackupHealth** | Pointer to [**BackupHealthCheckSettingsModels**](BackupHealthCheckSettingsModels.md) |  | [optional] 
**FullBackupMaintenance** | Pointer to [**FullBackupMaintenanceModel**](FullBackupMaintenanceModel.md) |  | [optional] 
**StorageData** | Pointer to [**BackupStorageSettingModel**](BackupStorageSettingModel.md) |  | [optional] 
**Notifications** | Pointer to [**NotificationSettingsModel**](NotificationSettingsModel.md) |  | [optional] 
**VSphere** | Pointer to [**BackupJobAdvancedSettingsVSphereModel**](BackupJobAdvancedSettingsVSphereModel.md) |  | [optional] 
**StorageIntegration** | Pointer to [**PrimaryStorageIntegrationSettingsModel**](PrimaryStorageIntegrationSettingsModel.md) |  | [optional] 
**Scripts** | Pointer to [**JobScriptsSettingsModel**](JobScriptsSettingsModel.md) |  | [optional] 

## Methods

### NewBackupJobAdvancedSettingsModel

`func NewBackupJobAdvancedSettingsModel(backupModeType EBackupModeType, ) *BackupJobAdvancedSettingsModel`

NewBackupJobAdvancedSettingsModel instantiates a new BackupJobAdvancedSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobAdvancedSettingsModelWithDefaults

`func NewBackupJobAdvancedSettingsModelWithDefaults() *BackupJobAdvancedSettingsModel`

NewBackupJobAdvancedSettingsModelWithDefaults instantiates a new BackupJobAdvancedSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupModeType

`func (o *BackupJobAdvancedSettingsModel) GetBackupModeType() EBackupModeType`

GetBackupModeType returns the BackupModeType field if non-nil, zero value otherwise.

### GetBackupModeTypeOk

`func (o *BackupJobAdvancedSettingsModel) GetBackupModeTypeOk() (*EBackupModeType, bool)`

GetBackupModeTypeOk returns a tuple with the BackupModeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupModeType

`func (o *BackupJobAdvancedSettingsModel) SetBackupModeType(v EBackupModeType)`

SetBackupModeType sets BackupModeType field to given value.


### GetSynthenticFulls

`func (o *BackupJobAdvancedSettingsModel) GetSynthenticFulls() SyntheticFullSettingsModel`

GetSynthenticFulls returns the SynthenticFulls field if non-nil, zero value otherwise.

### GetSynthenticFullsOk

`func (o *BackupJobAdvancedSettingsModel) GetSynthenticFullsOk() (*SyntheticFullSettingsModel, bool)`

GetSynthenticFullsOk returns a tuple with the SynthenticFulls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthenticFulls

`func (o *BackupJobAdvancedSettingsModel) SetSynthenticFulls(v SyntheticFullSettingsModel)`

SetSynthenticFulls sets SynthenticFulls field to given value.

### HasSynthenticFulls

`func (o *BackupJobAdvancedSettingsModel) HasSynthenticFulls() bool`

HasSynthenticFulls returns a boolean if a field has been set.

### GetActiveFulls

`func (o *BackupJobAdvancedSettingsModel) GetActiveFulls() ActiveFullSettingsModel`

GetActiveFulls returns the ActiveFulls field if non-nil, zero value otherwise.

### GetActiveFullsOk

`func (o *BackupJobAdvancedSettingsModel) GetActiveFullsOk() (*ActiveFullSettingsModel, bool)`

GetActiveFullsOk returns a tuple with the ActiveFulls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFulls

`func (o *BackupJobAdvancedSettingsModel) SetActiveFulls(v ActiveFullSettingsModel)`

SetActiveFulls sets ActiveFulls field to given value.

### HasActiveFulls

`func (o *BackupJobAdvancedSettingsModel) HasActiveFulls() bool`

HasActiveFulls returns a boolean if a field has been set.

### GetBackupHealth

`func (o *BackupJobAdvancedSettingsModel) GetBackupHealth() BackupHealthCheckSettingsModels`

GetBackupHealth returns the BackupHealth field if non-nil, zero value otherwise.

### GetBackupHealthOk

`func (o *BackupJobAdvancedSettingsModel) GetBackupHealthOk() (*BackupHealthCheckSettingsModels, bool)`

GetBackupHealthOk returns a tuple with the BackupHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupHealth

`func (o *BackupJobAdvancedSettingsModel) SetBackupHealth(v BackupHealthCheckSettingsModels)`

SetBackupHealth sets BackupHealth field to given value.

### HasBackupHealth

`func (o *BackupJobAdvancedSettingsModel) HasBackupHealth() bool`

HasBackupHealth returns a boolean if a field has been set.

### GetFullBackupMaintenance

`func (o *BackupJobAdvancedSettingsModel) GetFullBackupMaintenance() FullBackupMaintenanceModel`

GetFullBackupMaintenance returns the FullBackupMaintenance field if non-nil, zero value otherwise.

### GetFullBackupMaintenanceOk

`func (o *BackupJobAdvancedSettingsModel) GetFullBackupMaintenanceOk() (*FullBackupMaintenanceModel, bool)`

GetFullBackupMaintenanceOk returns a tuple with the FullBackupMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupMaintenance

`func (o *BackupJobAdvancedSettingsModel) SetFullBackupMaintenance(v FullBackupMaintenanceModel)`

SetFullBackupMaintenance sets FullBackupMaintenance field to given value.

### HasFullBackupMaintenance

`func (o *BackupJobAdvancedSettingsModel) HasFullBackupMaintenance() bool`

HasFullBackupMaintenance returns a boolean if a field has been set.

### GetStorageData

`func (o *BackupJobAdvancedSettingsModel) GetStorageData() BackupStorageSettingModel`

GetStorageData returns the StorageData field if non-nil, zero value otherwise.

### GetStorageDataOk

`func (o *BackupJobAdvancedSettingsModel) GetStorageDataOk() (*BackupStorageSettingModel, bool)`

GetStorageDataOk returns a tuple with the StorageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageData

`func (o *BackupJobAdvancedSettingsModel) SetStorageData(v BackupStorageSettingModel)`

SetStorageData sets StorageData field to given value.

### HasStorageData

`func (o *BackupJobAdvancedSettingsModel) HasStorageData() bool`

HasStorageData returns a boolean if a field has been set.

### GetNotifications

`func (o *BackupJobAdvancedSettingsModel) GetNotifications() NotificationSettingsModel`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *BackupJobAdvancedSettingsModel) GetNotificationsOk() (*NotificationSettingsModel, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *BackupJobAdvancedSettingsModel) SetNotifications(v NotificationSettingsModel)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *BackupJobAdvancedSettingsModel) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetVSphere

`func (o *BackupJobAdvancedSettingsModel) GetVSphere() BackupJobAdvancedSettingsVSphereModel`

GetVSphere returns the VSphere field if non-nil, zero value otherwise.

### GetVSphereOk

`func (o *BackupJobAdvancedSettingsModel) GetVSphereOk() (*BackupJobAdvancedSettingsVSphereModel, bool)`

GetVSphereOk returns a tuple with the VSphere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSphere

`func (o *BackupJobAdvancedSettingsModel) SetVSphere(v BackupJobAdvancedSettingsVSphereModel)`

SetVSphere sets VSphere field to given value.

### HasVSphere

`func (o *BackupJobAdvancedSettingsModel) HasVSphere() bool`

HasVSphere returns a boolean if a field has been set.

### GetStorageIntegration

`func (o *BackupJobAdvancedSettingsModel) GetStorageIntegration() PrimaryStorageIntegrationSettingsModel`

GetStorageIntegration returns the StorageIntegration field if non-nil, zero value otherwise.

### GetStorageIntegrationOk

`func (o *BackupJobAdvancedSettingsModel) GetStorageIntegrationOk() (*PrimaryStorageIntegrationSettingsModel, bool)`

GetStorageIntegrationOk returns a tuple with the StorageIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageIntegration

`func (o *BackupJobAdvancedSettingsModel) SetStorageIntegration(v PrimaryStorageIntegrationSettingsModel)`

SetStorageIntegration sets StorageIntegration field to given value.

### HasStorageIntegration

`func (o *BackupJobAdvancedSettingsModel) HasStorageIntegration() bool`

HasStorageIntegration returns a boolean if a field has been set.

### GetScripts

`func (o *BackupJobAdvancedSettingsModel) GetScripts() JobScriptsSettingsModel`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *BackupJobAdvancedSettingsModel) GetScriptsOk() (*JobScriptsSettingsModel, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *BackupJobAdvancedSettingsModel) SetScripts(v JobScriptsSettingsModel)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *BackupJobAdvancedSettingsModel) HasScripts() bool`

HasScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


