# BackupJobStorageImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRepository** | [**BackupRepositoryImportModel**](BackupRepositoryImportModel.md) |  | 
**BackupProxies** | [**BackupJobImportProxiesModel**](BackupJobImportProxiesModel.md) |  | 
**RetentionPolicy** | [**BackupJobRetentionPolicySettingsModel**](BackupJobRetentionPolicySettingsModel.md) |  | 
**GfsPolicy** | Pointer to [**GFSPolicySettingsModel**](GFSPolicySettingsModel.md) |  | [optional] 
**AdvancedSettings** | Pointer to [**BackupJobAdvancedSettingsModel**](BackupJobAdvancedSettingsModel.md) |  | [optional] 

## Methods

### NewBackupJobStorageImportModel

`func NewBackupJobStorageImportModel(backupRepository BackupRepositoryImportModel, backupProxies BackupJobImportProxiesModel, retentionPolicy BackupJobRetentionPolicySettingsModel, ) *BackupJobStorageImportModel`

NewBackupJobStorageImportModel instantiates a new BackupJobStorageImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobStorageImportModelWithDefaults

`func NewBackupJobStorageImportModelWithDefaults() *BackupJobStorageImportModel`

NewBackupJobStorageImportModelWithDefaults instantiates a new BackupJobStorageImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRepository

`func (o *BackupJobStorageImportModel) GetBackupRepository() BackupRepositoryImportModel`

GetBackupRepository returns the BackupRepository field if non-nil, zero value otherwise.

### GetBackupRepositoryOk

`func (o *BackupJobStorageImportModel) GetBackupRepositoryOk() (*BackupRepositoryImportModel, bool)`

GetBackupRepositoryOk returns a tuple with the BackupRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRepository

`func (o *BackupJobStorageImportModel) SetBackupRepository(v BackupRepositoryImportModel)`

SetBackupRepository sets BackupRepository field to given value.


### GetBackupProxies

`func (o *BackupJobStorageImportModel) GetBackupProxies() BackupJobImportProxiesModel`

GetBackupProxies returns the BackupProxies field if non-nil, zero value otherwise.

### GetBackupProxiesOk

`func (o *BackupJobStorageImportModel) GetBackupProxiesOk() (*BackupJobImportProxiesModel, bool)`

GetBackupProxiesOk returns a tuple with the BackupProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProxies

`func (o *BackupJobStorageImportModel) SetBackupProxies(v BackupJobImportProxiesModel)`

SetBackupProxies sets BackupProxies field to given value.


### GetRetentionPolicy

`func (o *BackupJobStorageImportModel) GetRetentionPolicy() BackupJobRetentionPolicySettingsModel`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *BackupJobStorageImportModel) GetRetentionPolicyOk() (*BackupJobRetentionPolicySettingsModel, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *BackupJobStorageImportModel) SetRetentionPolicy(v BackupJobRetentionPolicySettingsModel)`

SetRetentionPolicy sets RetentionPolicy field to given value.


### GetGfsPolicy

`func (o *BackupJobStorageImportModel) GetGfsPolicy() GFSPolicySettingsModel`

GetGfsPolicy returns the GfsPolicy field if non-nil, zero value otherwise.

### GetGfsPolicyOk

`func (o *BackupJobStorageImportModel) GetGfsPolicyOk() (*GFSPolicySettingsModel, bool)`

GetGfsPolicyOk returns a tuple with the GfsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGfsPolicy

`func (o *BackupJobStorageImportModel) SetGfsPolicy(v GFSPolicySettingsModel)`

SetGfsPolicy sets GfsPolicy field to given value.

### HasGfsPolicy

`func (o *BackupJobStorageImportModel) HasGfsPolicy() bool`

HasGfsPolicy returns a boolean if a field has been set.

### GetAdvancedSettings

`func (o *BackupJobStorageImportModel) GetAdvancedSettings() BackupJobAdvancedSettingsModel`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *BackupJobStorageImportModel) GetAdvancedSettingsOk() (*BackupJobAdvancedSettingsModel, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *BackupJobStorageImportModel) SetAdvancedSettings(v BackupJobAdvancedSettingsModel)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *BackupJobStorageImportModel) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


