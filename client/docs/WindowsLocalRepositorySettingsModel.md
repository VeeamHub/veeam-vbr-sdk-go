# WindowsLocalRepositorySettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path to the folder where backup files are stored. | [optional] 
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**EnableReadWriteLimit** | Pointer to **bool** | If *true*, reading and writing speed is limited. | [optional] 
**ReadWriteRate** | Pointer to **int32** | Maximum rate that restricts the total speed of reading and writing data to the backup repository disk. | [optional] 
**AdvancedSettings** | Pointer to [**RepositoryAdvancedSettingsModel**](RepositoryAdvancedSettingsModel.md) |  | [optional] 

## Methods

### NewWindowsLocalRepositorySettingsModel

`func NewWindowsLocalRepositorySettingsModel() *WindowsLocalRepositorySettingsModel`

NewWindowsLocalRepositorySettingsModel instantiates a new WindowsLocalRepositorySettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsLocalRepositorySettingsModelWithDefaults

`func NewWindowsLocalRepositorySettingsModelWithDefaults() *WindowsLocalRepositorySettingsModel`

NewWindowsLocalRepositorySettingsModelWithDefaults instantiates a new WindowsLocalRepositorySettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *WindowsLocalRepositorySettingsModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WindowsLocalRepositorySettingsModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WindowsLocalRepositorySettingsModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WindowsLocalRepositorySettingsModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetEnableTaskLimit

`func (o *WindowsLocalRepositorySettingsModel) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *WindowsLocalRepositorySettingsModel) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *WindowsLocalRepositorySettingsModel) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *WindowsLocalRepositorySettingsModel) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *WindowsLocalRepositorySettingsModel) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *WindowsLocalRepositorySettingsModel) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *WindowsLocalRepositorySettingsModel) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *WindowsLocalRepositorySettingsModel) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetEnableReadWriteLimit

`func (o *WindowsLocalRepositorySettingsModel) GetEnableReadWriteLimit() bool`

GetEnableReadWriteLimit returns the EnableReadWriteLimit field if non-nil, zero value otherwise.

### GetEnableReadWriteLimitOk

`func (o *WindowsLocalRepositorySettingsModel) GetEnableReadWriteLimitOk() (*bool, bool)`

GetEnableReadWriteLimitOk returns a tuple with the EnableReadWriteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReadWriteLimit

`func (o *WindowsLocalRepositorySettingsModel) SetEnableReadWriteLimit(v bool)`

SetEnableReadWriteLimit sets EnableReadWriteLimit field to given value.

### HasEnableReadWriteLimit

`func (o *WindowsLocalRepositorySettingsModel) HasEnableReadWriteLimit() bool`

HasEnableReadWriteLimit returns a boolean if a field has been set.

### GetReadWriteRate

`func (o *WindowsLocalRepositorySettingsModel) GetReadWriteRate() int32`

GetReadWriteRate returns the ReadWriteRate field if non-nil, zero value otherwise.

### GetReadWriteRateOk

`func (o *WindowsLocalRepositorySettingsModel) GetReadWriteRateOk() (*int32, bool)`

GetReadWriteRateOk returns a tuple with the ReadWriteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteRate

`func (o *WindowsLocalRepositorySettingsModel) SetReadWriteRate(v int32)`

SetReadWriteRate sets ReadWriteRate field to given value.

### HasReadWriteRate

`func (o *WindowsLocalRepositorySettingsModel) HasReadWriteRate() bool`

HasReadWriteRate returns a boolean if a field has been set.

### GetAdvancedSettings

`func (o *WindowsLocalRepositorySettingsModel) GetAdvancedSettings() RepositoryAdvancedSettingsModel`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *WindowsLocalRepositorySettingsModel) GetAdvancedSettingsOk() (*RepositoryAdvancedSettingsModel, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *WindowsLocalRepositorySettingsModel) SetAdvancedSettings(v RepositoryAdvancedSettingsModel)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *WindowsLocalRepositorySettingsModel) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


