# LinuxLocalRepositorySettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path to the folder where backup files are stored. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**ReadWriteRate** | Pointer to **int32** | Maximum rate that restricts the total speed of reading and writing data to the backup repository disk. | [optional] 
**UseFastCloningOnXFSVolumes** | Pointer to **bool** | [For Linux repository] If *true*, fast cloning on XFS volumes is used.  | [optional] 
**UseImmutableBackups** | Pointer to **bool** | If *true*, the Object Lock feature is used to protect recent backups. | [optional] 
**MakeRecentBackupsImmutableDays** | Pointer to **int32** | Number of days to keep immutable backups. | [optional] 
**AdvancedSettings** | Pointer to [**RepositoryAdvancedSettingsModel**](RepositoryAdvancedSettingsModel.md) |  | [optional] 

## Methods

### NewLinuxLocalRepositorySettingsModel

`func NewLinuxLocalRepositorySettingsModel() *LinuxLocalRepositorySettingsModel`

NewLinuxLocalRepositorySettingsModel instantiates a new LinuxLocalRepositorySettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxLocalRepositorySettingsModelWithDefaults

`func NewLinuxLocalRepositorySettingsModelWithDefaults() *LinuxLocalRepositorySettingsModel`

NewLinuxLocalRepositorySettingsModelWithDefaults instantiates a new LinuxLocalRepositorySettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *LinuxLocalRepositorySettingsModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LinuxLocalRepositorySettingsModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LinuxLocalRepositorySettingsModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LinuxLocalRepositorySettingsModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *LinuxLocalRepositorySettingsModel) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *LinuxLocalRepositorySettingsModel) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *LinuxLocalRepositorySettingsModel) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *LinuxLocalRepositorySettingsModel) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetReadWriteRate

`func (o *LinuxLocalRepositorySettingsModel) GetReadWriteRate() int32`

GetReadWriteRate returns the ReadWriteRate field if non-nil, zero value otherwise.

### GetReadWriteRateOk

`func (o *LinuxLocalRepositorySettingsModel) GetReadWriteRateOk() (*int32, bool)`

GetReadWriteRateOk returns a tuple with the ReadWriteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteRate

`func (o *LinuxLocalRepositorySettingsModel) SetReadWriteRate(v int32)`

SetReadWriteRate sets ReadWriteRate field to given value.

### HasReadWriteRate

`func (o *LinuxLocalRepositorySettingsModel) HasReadWriteRate() bool`

HasReadWriteRate returns a boolean if a field has been set.

### GetUseFastCloningOnXFSVolumes

`func (o *LinuxLocalRepositorySettingsModel) GetUseFastCloningOnXFSVolumes() bool`

GetUseFastCloningOnXFSVolumes returns the UseFastCloningOnXFSVolumes field if non-nil, zero value otherwise.

### GetUseFastCloningOnXFSVolumesOk

`func (o *LinuxLocalRepositorySettingsModel) GetUseFastCloningOnXFSVolumesOk() (*bool, bool)`

GetUseFastCloningOnXFSVolumesOk returns a tuple with the UseFastCloningOnXFSVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFastCloningOnXFSVolumes

`func (o *LinuxLocalRepositorySettingsModel) SetUseFastCloningOnXFSVolumes(v bool)`

SetUseFastCloningOnXFSVolumes sets UseFastCloningOnXFSVolumes field to given value.

### HasUseFastCloningOnXFSVolumes

`func (o *LinuxLocalRepositorySettingsModel) HasUseFastCloningOnXFSVolumes() bool`

HasUseFastCloningOnXFSVolumes returns a boolean if a field has been set.

### GetUseImmutableBackups

`func (o *LinuxLocalRepositorySettingsModel) GetUseImmutableBackups() bool`

GetUseImmutableBackups returns the UseImmutableBackups field if non-nil, zero value otherwise.

### GetUseImmutableBackupsOk

`func (o *LinuxLocalRepositorySettingsModel) GetUseImmutableBackupsOk() (*bool, bool)`

GetUseImmutableBackupsOk returns a tuple with the UseImmutableBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseImmutableBackups

`func (o *LinuxLocalRepositorySettingsModel) SetUseImmutableBackups(v bool)`

SetUseImmutableBackups sets UseImmutableBackups field to given value.

### HasUseImmutableBackups

`func (o *LinuxLocalRepositorySettingsModel) HasUseImmutableBackups() bool`

HasUseImmutableBackups returns a boolean if a field has been set.

### GetMakeRecentBackupsImmutableDays

`func (o *LinuxLocalRepositorySettingsModel) GetMakeRecentBackupsImmutableDays() int32`

GetMakeRecentBackupsImmutableDays returns the MakeRecentBackupsImmutableDays field if non-nil, zero value otherwise.

### GetMakeRecentBackupsImmutableDaysOk

`func (o *LinuxLocalRepositorySettingsModel) GetMakeRecentBackupsImmutableDaysOk() (*int32, bool)`

GetMakeRecentBackupsImmutableDaysOk returns a tuple with the MakeRecentBackupsImmutableDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeRecentBackupsImmutableDays

`func (o *LinuxLocalRepositorySettingsModel) SetMakeRecentBackupsImmutableDays(v int32)`

SetMakeRecentBackupsImmutableDays sets MakeRecentBackupsImmutableDays field to given value.

### HasMakeRecentBackupsImmutableDays

`func (o *LinuxLocalRepositorySettingsModel) HasMakeRecentBackupsImmutableDays() bool`

HasMakeRecentBackupsImmutableDays returns a boolean if a field has been set.

### GetAdvancedSettings

`func (o *LinuxLocalRepositorySettingsModel) GetAdvancedSettings() RepositoryAdvancedSettingsModel`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *LinuxLocalRepositorySettingsModel) GetAdvancedSettingsOk() (*RepositoryAdvancedSettingsModel, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *LinuxLocalRepositorySettingsModel) SetAdvancedSettings(v RepositoryAdvancedSettingsModel)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *LinuxLocalRepositorySettingsModel) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


