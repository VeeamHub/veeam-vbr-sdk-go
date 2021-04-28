# BackupJobStorageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRepositoryId** | **string** | ID of the backup repository. | 
**BackupProxies** | [**BackupProxiesSettingsModel**](BackupProxiesSettingsModel.md) |  | 
**RetentionPolicy** | [**BackupJobRetentionPolicySettingsModel**](BackupJobRetentionPolicySettingsModel.md) |  | 
**GfsPolicy** | Pointer to [**GFSPolicySettingsModel**](GFSPolicySettingsModel.md) |  | [optional] 
**AdvancedSettings** | Pointer to [**BackupJobAdvancedSettingsModel**](BackupJobAdvancedSettingsModel.md) |  | [optional] 

## Methods

### NewBackupJobStorageModel

`func NewBackupJobStorageModel(backupRepositoryId string, backupProxies BackupProxiesSettingsModel, retentionPolicy BackupJobRetentionPolicySettingsModel, ) *BackupJobStorageModel`

NewBackupJobStorageModel instantiates a new BackupJobStorageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobStorageModelWithDefaults

`func NewBackupJobStorageModelWithDefaults() *BackupJobStorageModel`

NewBackupJobStorageModelWithDefaults instantiates a new BackupJobStorageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRepositoryId

`func (o *BackupJobStorageModel) GetBackupRepositoryId() string`

GetBackupRepositoryId returns the BackupRepositoryId field if non-nil, zero value otherwise.

### GetBackupRepositoryIdOk

`func (o *BackupJobStorageModel) GetBackupRepositoryIdOk() (*string, bool)`

GetBackupRepositoryIdOk returns a tuple with the BackupRepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRepositoryId

`func (o *BackupJobStorageModel) SetBackupRepositoryId(v string)`

SetBackupRepositoryId sets BackupRepositoryId field to given value.


### GetBackupProxies

`func (o *BackupJobStorageModel) GetBackupProxies() BackupProxiesSettingsModel`

GetBackupProxies returns the BackupProxies field if non-nil, zero value otherwise.

### GetBackupProxiesOk

`func (o *BackupJobStorageModel) GetBackupProxiesOk() (*BackupProxiesSettingsModel, bool)`

GetBackupProxiesOk returns a tuple with the BackupProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProxies

`func (o *BackupJobStorageModel) SetBackupProxies(v BackupProxiesSettingsModel)`

SetBackupProxies sets BackupProxies field to given value.


### GetRetentionPolicy

`func (o *BackupJobStorageModel) GetRetentionPolicy() BackupJobRetentionPolicySettingsModel`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *BackupJobStorageModel) GetRetentionPolicyOk() (*BackupJobRetentionPolicySettingsModel, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *BackupJobStorageModel) SetRetentionPolicy(v BackupJobRetentionPolicySettingsModel)`

SetRetentionPolicy sets RetentionPolicy field to given value.


### GetGfsPolicy

`func (o *BackupJobStorageModel) GetGfsPolicy() GFSPolicySettingsModel`

GetGfsPolicy returns the GfsPolicy field if non-nil, zero value otherwise.

### GetGfsPolicyOk

`func (o *BackupJobStorageModel) GetGfsPolicyOk() (*GFSPolicySettingsModel, bool)`

GetGfsPolicyOk returns a tuple with the GfsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGfsPolicy

`func (o *BackupJobStorageModel) SetGfsPolicy(v GFSPolicySettingsModel)`

SetGfsPolicy sets GfsPolicy field to given value.

### HasGfsPolicy

`func (o *BackupJobStorageModel) HasGfsPolicy() bool`

HasGfsPolicy returns a boolean if a field has been set.

### GetAdvancedSettings

`func (o *BackupJobStorageModel) GetAdvancedSettings() BackupJobAdvancedSettingsModel`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *BackupJobStorageModel) GetAdvancedSettingsOk() (*BackupJobAdvancedSettingsModel, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *BackupJobStorageModel) SetAdvancedSettings(v BackupJobAdvancedSettingsModel)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *BackupJobStorageModel) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


