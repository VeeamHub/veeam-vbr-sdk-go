# BackupStorageSettingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableInlineDataDedup** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication deduplicates VM data before storing it in the backup repository. | [optional] 
**ExcludeSwapFileBlocks** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication excludes swap file blocks from processing. | [optional] 
**ExcludeDeletedFileBlocks** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication does not copy deleted file blocks. | [optional] 
**CompressionLevel** | Pointer to [**ECompressionLevel**](ECompressionLevel.md) |  | [optional] 
**StorageOptimization** | Pointer to [**EStorageOptimization**](EStorageOptimization.md) |  | [optional] 
**Encryption** | Pointer to [**BackupStorageSettingsEncryptionModel**](BackupStorageSettingsEncryptionModel.md) |  | [optional] 

## Methods

### NewBackupStorageSettingModel

`func NewBackupStorageSettingModel() *BackupStorageSettingModel`

NewBackupStorageSettingModel instantiates a new BackupStorageSettingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupStorageSettingModelWithDefaults

`func NewBackupStorageSettingModelWithDefaults() *BackupStorageSettingModel`

NewBackupStorageSettingModelWithDefaults instantiates a new BackupStorageSettingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableInlineDataDedup

`func (o *BackupStorageSettingModel) GetEnableInlineDataDedup() bool`

GetEnableInlineDataDedup returns the EnableInlineDataDedup field if non-nil, zero value otherwise.

### GetEnableInlineDataDedupOk

`func (o *BackupStorageSettingModel) GetEnableInlineDataDedupOk() (*bool, bool)`

GetEnableInlineDataDedupOk returns a tuple with the EnableInlineDataDedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInlineDataDedup

`func (o *BackupStorageSettingModel) SetEnableInlineDataDedup(v bool)`

SetEnableInlineDataDedup sets EnableInlineDataDedup field to given value.

### HasEnableInlineDataDedup

`func (o *BackupStorageSettingModel) HasEnableInlineDataDedup() bool`

HasEnableInlineDataDedup returns a boolean if a field has been set.

### GetExcludeSwapFileBlocks

`func (o *BackupStorageSettingModel) GetExcludeSwapFileBlocks() bool`

GetExcludeSwapFileBlocks returns the ExcludeSwapFileBlocks field if non-nil, zero value otherwise.

### GetExcludeSwapFileBlocksOk

`func (o *BackupStorageSettingModel) GetExcludeSwapFileBlocksOk() (*bool, bool)`

GetExcludeSwapFileBlocksOk returns a tuple with the ExcludeSwapFileBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSwapFileBlocks

`func (o *BackupStorageSettingModel) SetExcludeSwapFileBlocks(v bool)`

SetExcludeSwapFileBlocks sets ExcludeSwapFileBlocks field to given value.

### HasExcludeSwapFileBlocks

`func (o *BackupStorageSettingModel) HasExcludeSwapFileBlocks() bool`

HasExcludeSwapFileBlocks returns a boolean if a field has been set.

### GetExcludeDeletedFileBlocks

`func (o *BackupStorageSettingModel) GetExcludeDeletedFileBlocks() bool`

GetExcludeDeletedFileBlocks returns the ExcludeDeletedFileBlocks field if non-nil, zero value otherwise.

### GetExcludeDeletedFileBlocksOk

`func (o *BackupStorageSettingModel) GetExcludeDeletedFileBlocksOk() (*bool, bool)`

GetExcludeDeletedFileBlocksOk returns a tuple with the ExcludeDeletedFileBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDeletedFileBlocks

`func (o *BackupStorageSettingModel) SetExcludeDeletedFileBlocks(v bool)`

SetExcludeDeletedFileBlocks sets ExcludeDeletedFileBlocks field to given value.

### HasExcludeDeletedFileBlocks

`func (o *BackupStorageSettingModel) HasExcludeDeletedFileBlocks() bool`

HasExcludeDeletedFileBlocks returns a boolean if a field has been set.

### GetCompressionLevel

`func (o *BackupStorageSettingModel) GetCompressionLevel() ECompressionLevel`

GetCompressionLevel returns the CompressionLevel field if non-nil, zero value otherwise.

### GetCompressionLevelOk

`func (o *BackupStorageSettingModel) GetCompressionLevelOk() (*ECompressionLevel, bool)`

GetCompressionLevelOk returns a tuple with the CompressionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionLevel

`func (o *BackupStorageSettingModel) SetCompressionLevel(v ECompressionLevel)`

SetCompressionLevel sets CompressionLevel field to given value.

### HasCompressionLevel

`func (o *BackupStorageSettingModel) HasCompressionLevel() bool`

HasCompressionLevel returns a boolean if a field has been set.

### GetStorageOptimization

`func (o *BackupStorageSettingModel) GetStorageOptimization() EStorageOptimization`

GetStorageOptimization returns the StorageOptimization field if non-nil, zero value otherwise.

### GetStorageOptimizationOk

`func (o *BackupStorageSettingModel) GetStorageOptimizationOk() (*EStorageOptimization, bool)`

GetStorageOptimizationOk returns a tuple with the StorageOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOptimization

`func (o *BackupStorageSettingModel) SetStorageOptimization(v EStorageOptimization)`

SetStorageOptimization sets StorageOptimization field to given value.

### HasStorageOptimization

`func (o *BackupStorageSettingModel) HasStorageOptimization() bool`

HasStorageOptimization returns a boolean if a field has been set.

### GetEncryption

`func (o *BackupStorageSettingModel) GetEncryption() BackupStorageSettingsEncryptionModel`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *BackupStorageSettingModel) GetEncryptionOk() (*BackupStorageSettingsEncryptionModel, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *BackupStorageSettingModel) SetEncryption(v BackupStorageSettingsEncryptionModel)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *BackupStorageSettingModel) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


