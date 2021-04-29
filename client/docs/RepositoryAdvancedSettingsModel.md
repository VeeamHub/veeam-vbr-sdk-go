# RepositoryAdvancedSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlignDataBlocks** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication aligns VM data saved to a backup file at a 4 KB block boundary. | [optional] 
**DecompressBeforeStoring** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication decompresses backup data blocks before storing to improve deduplication ratios. | [optional] 
**RotatedDrives** | Pointer to **bool** | If *true*, the repository drive is rotated. | [optional] 
**PerVmBackup** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication creates a separate backup file for every VM in the job. | [optional] 

## Methods

### NewRepositoryAdvancedSettingsModel

`func NewRepositoryAdvancedSettingsModel() *RepositoryAdvancedSettingsModel`

NewRepositoryAdvancedSettingsModel instantiates a new RepositoryAdvancedSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryAdvancedSettingsModelWithDefaults

`func NewRepositoryAdvancedSettingsModelWithDefaults() *RepositoryAdvancedSettingsModel`

NewRepositoryAdvancedSettingsModelWithDefaults instantiates a new RepositoryAdvancedSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlignDataBlocks

`func (o *RepositoryAdvancedSettingsModel) GetAlignDataBlocks() bool`

GetAlignDataBlocks returns the AlignDataBlocks field if non-nil, zero value otherwise.

### GetAlignDataBlocksOk

`func (o *RepositoryAdvancedSettingsModel) GetAlignDataBlocksOk() (*bool, bool)`

GetAlignDataBlocksOk returns a tuple with the AlignDataBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignDataBlocks

`func (o *RepositoryAdvancedSettingsModel) SetAlignDataBlocks(v bool)`

SetAlignDataBlocks sets AlignDataBlocks field to given value.

### HasAlignDataBlocks

`func (o *RepositoryAdvancedSettingsModel) HasAlignDataBlocks() bool`

HasAlignDataBlocks returns a boolean if a field has been set.

### GetDecompressBeforeStoring

`func (o *RepositoryAdvancedSettingsModel) GetDecompressBeforeStoring() bool`

GetDecompressBeforeStoring returns the DecompressBeforeStoring field if non-nil, zero value otherwise.

### GetDecompressBeforeStoringOk

`func (o *RepositoryAdvancedSettingsModel) GetDecompressBeforeStoringOk() (*bool, bool)`

GetDecompressBeforeStoringOk returns a tuple with the DecompressBeforeStoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecompressBeforeStoring

`func (o *RepositoryAdvancedSettingsModel) SetDecompressBeforeStoring(v bool)`

SetDecompressBeforeStoring sets DecompressBeforeStoring field to given value.

### HasDecompressBeforeStoring

`func (o *RepositoryAdvancedSettingsModel) HasDecompressBeforeStoring() bool`

HasDecompressBeforeStoring returns a boolean if a field has been set.

### GetRotatedDrives

`func (o *RepositoryAdvancedSettingsModel) GetRotatedDrives() bool`

GetRotatedDrives returns the RotatedDrives field if non-nil, zero value otherwise.

### GetRotatedDrivesOk

`func (o *RepositoryAdvancedSettingsModel) GetRotatedDrivesOk() (*bool, bool)`

GetRotatedDrivesOk returns a tuple with the RotatedDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedDrives

`func (o *RepositoryAdvancedSettingsModel) SetRotatedDrives(v bool)`

SetRotatedDrives sets RotatedDrives field to given value.

### HasRotatedDrives

`func (o *RepositoryAdvancedSettingsModel) HasRotatedDrives() bool`

HasRotatedDrives returns a boolean if a field has been set.

### GetPerVmBackup

`func (o *RepositoryAdvancedSettingsModel) GetPerVmBackup() bool`

GetPerVmBackup returns the PerVmBackup field if non-nil, zero value otherwise.

### GetPerVmBackupOk

`func (o *RepositoryAdvancedSettingsModel) GetPerVmBackupOk() (*bool, bool)`

GetPerVmBackupOk returns a tuple with the PerVmBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerVmBackup

`func (o *RepositoryAdvancedSettingsModel) SetPerVmBackup(v bool)`

SetPerVmBackup sets PerVmBackup field to given value.

### HasPerVmBackup

`func (o *RepositoryAdvancedSettingsModel) HasPerVmBackup() bool`

HasPerVmBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


