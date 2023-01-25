# LinuxHardenedRepositorySettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path to the folder where backup files are stored. | [optional] 
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**EnableReadWriteLimit** | Pointer to **bool** | If *true*, reading and writing speed is limited. | [optional] 
**ReadWriteRate** | Pointer to **int32** | Maximum rate that restricts the total speed of reading and writing data to the backup repository disk. | [optional] 
**UseFastCloningOnXFSVolumes** | Pointer to **bool** | If *true*, fast cloning on XFS volumes is used. | [optional] 
**MakeRecentBackupsImmutableDays** | Pointer to **int32** | Number of days to keep immutable backups. | [optional] 
**AdvancedSettings** | Pointer to [**RepositoryAdvancedSettingsModel**](RepositoryAdvancedSettingsModel.md) |  | [optional] 

## Methods

### NewLinuxHardenedRepositorySettingsModel

`func NewLinuxHardenedRepositorySettingsModel() *LinuxHardenedRepositorySettingsModel`

NewLinuxHardenedRepositorySettingsModel instantiates a new LinuxHardenedRepositorySettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxHardenedRepositorySettingsModelWithDefaults

`func NewLinuxHardenedRepositorySettingsModelWithDefaults() *LinuxHardenedRepositorySettingsModel`

NewLinuxHardenedRepositorySettingsModelWithDefaults instantiates a new LinuxHardenedRepositorySettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *LinuxHardenedRepositorySettingsModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LinuxHardenedRepositorySettingsModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LinuxHardenedRepositorySettingsModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LinuxHardenedRepositorySettingsModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetEnableTaskLimit

`func (o *LinuxHardenedRepositorySettingsModel) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *LinuxHardenedRepositorySettingsModel) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *LinuxHardenedRepositorySettingsModel) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *LinuxHardenedRepositorySettingsModel) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *LinuxHardenedRepositorySettingsModel) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *LinuxHardenedRepositorySettingsModel) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *LinuxHardenedRepositorySettingsModel) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *LinuxHardenedRepositorySettingsModel) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetEnableReadWriteLimit

`func (o *LinuxHardenedRepositorySettingsModel) GetEnableReadWriteLimit() bool`

GetEnableReadWriteLimit returns the EnableReadWriteLimit field if non-nil, zero value otherwise.

### GetEnableReadWriteLimitOk

`func (o *LinuxHardenedRepositorySettingsModel) GetEnableReadWriteLimitOk() (*bool, bool)`

GetEnableReadWriteLimitOk returns a tuple with the EnableReadWriteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReadWriteLimit

`func (o *LinuxHardenedRepositorySettingsModel) SetEnableReadWriteLimit(v bool)`

SetEnableReadWriteLimit sets EnableReadWriteLimit field to given value.

### HasEnableReadWriteLimit

`func (o *LinuxHardenedRepositorySettingsModel) HasEnableReadWriteLimit() bool`

HasEnableReadWriteLimit returns a boolean if a field has been set.

### GetReadWriteRate

`func (o *LinuxHardenedRepositorySettingsModel) GetReadWriteRate() int32`

GetReadWriteRate returns the ReadWriteRate field if non-nil, zero value otherwise.

### GetReadWriteRateOk

`func (o *LinuxHardenedRepositorySettingsModel) GetReadWriteRateOk() (*int32, bool)`

GetReadWriteRateOk returns a tuple with the ReadWriteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteRate

`func (o *LinuxHardenedRepositorySettingsModel) SetReadWriteRate(v int32)`

SetReadWriteRate sets ReadWriteRate field to given value.

### HasReadWriteRate

`func (o *LinuxHardenedRepositorySettingsModel) HasReadWriteRate() bool`

HasReadWriteRate returns a boolean if a field has been set.

### GetUseFastCloningOnXFSVolumes

`func (o *LinuxHardenedRepositorySettingsModel) GetUseFastCloningOnXFSVolumes() bool`

GetUseFastCloningOnXFSVolumes returns the UseFastCloningOnXFSVolumes field if non-nil, zero value otherwise.

### GetUseFastCloningOnXFSVolumesOk

`func (o *LinuxHardenedRepositorySettingsModel) GetUseFastCloningOnXFSVolumesOk() (*bool, bool)`

GetUseFastCloningOnXFSVolumesOk returns a tuple with the UseFastCloningOnXFSVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFastCloningOnXFSVolumes

`func (o *LinuxHardenedRepositorySettingsModel) SetUseFastCloningOnXFSVolumes(v bool)`

SetUseFastCloningOnXFSVolumes sets UseFastCloningOnXFSVolumes field to given value.

### HasUseFastCloningOnXFSVolumes

`func (o *LinuxHardenedRepositorySettingsModel) HasUseFastCloningOnXFSVolumes() bool`

HasUseFastCloningOnXFSVolumes returns a boolean if a field has been set.

### GetMakeRecentBackupsImmutableDays

`func (o *LinuxHardenedRepositorySettingsModel) GetMakeRecentBackupsImmutableDays() int32`

GetMakeRecentBackupsImmutableDays returns the MakeRecentBackupsImmutableDays field if non-nil, zero value otherwise.

### GetMakeRecentBackupsImmutableDaysOk

`func (o *LinuxHardenedRepositorySettingsModel) GetMakeRecentBackupsImmutableDaysOk() (*int32, bool)`

GetMakeRecentBackupsImmutableDaysOk returns a tuple with the MakeRecentBackupsImmutableDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeRecentBackupsImmutableDays

`func (o *LinuxHardenedRepositorySettingsModel) SetMakeRecentBackupsImmutableDays(v int32)`

SetMakeRecentBackupsImmutableDays sets MakeRecentBackupsImmutableDays field to given value.

### HasMakeRecentBackupsImmutableDays

`func (o *LinuxHardenedRepositorySettingsModel) HasMakeRecentBackupsImmutableDays() bool`

HasMakeRecentBackupsImmutableDays returns a boolean if a field has been set.

### GetAdvancedSettings

`func (o *LinuxHardenedRepositorySettingsModel) GetAdvancedSettings() RepositoryAdvancedSettingsModel`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *LinuxHardenedRepositorySettingsModel) GetAdvancedSettingsOk() (*RepositoryAdvancedSettingsModel, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *LinuxHardenedRepositorySettingsModel) SetAdvancedSettings(v RepositoryAdvancedSettingsModel)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *LinuxHardenedRepositorySettingsModel) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


