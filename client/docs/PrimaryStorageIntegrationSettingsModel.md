# PrimaryStorageIntegrationSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, the primary storage integration is enabled. In this case, storage snapshots (instead of VM snapshots) are used for VM data processing. | 
**LimitProcessedVm** | Pointer to **bool** | If *true*, the number of processed VMs per storage snapshot is limited. | [optional] 
**LimitProcessedVmCount** | Pointer to **int32** | Number of processed VMs per storage snapshot. | [optional] 
**FailoverToStandardBackup** | Pointer to **bool** | If *true*, failover to the regular VM processing mode is enabled. In this case, if backup from storage snapshot fails, VM snapshots are used. | [optional] 

## Methods

### NewPrimaryStorageIntegrationSettingsModel

`func NewPrimaryStorageIntegrationSettingsModel(isEnabled bool, ) *PrimaryStorageIntegrationSettingsModel`

NewPrimaryStorageIntegrationSettingsModel instantiates a new PrimaryStorageIntegrationSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrimaryStorageIntegrationSettingsModelWithDefaults

`func NewPrimaryStorageIntegrationSettingsModelWithDefaults() *PrimaryStorageIntegrationSettingsModel`

NewPrimaryStorageIntegrationSettingsModelWithDefaults instantiates a new PrimaryStorageIntegrationSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *PrimaryStorageIntegrationSettingsModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PrimaryStorageIntegrationSettingsModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PrimaryStorageIntegrationSettingsModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetLimitProcessedVm

`func (o *PrimaryStorageIntegrationSettingsModel) GetLimitProcessedVm() bool`

GetLimitProcessedVm returns the LimitProcessedVm field if non-nil, zero value otherwise.

### GetLimitProcessedVmOk

`func (o *PrimaryStorageIntegrationSettingsModel) GetLimitProcessedVmOk() (*bool, bool)`

GetLimitProcessedVmOk returns a tuple with the LimitProcessedVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitProcessedVm

`func (o *PrimaryStorageIntegrationSettingsModel) SetLimitProcessedVm(v bool)`

SetLimitProcessedVm sets LimitProcessedVm field to given value.

### HasLimitProcessedVm

`func (o *PrimaryStorageIntegrationSettingsModel) HasLimitProcessedVm() bool`

HasLimitProcessedVm returns a boolean if a field has been set.

### GetLimitProcessedVmCount

`func (o *PrimaryStorageIntegrationSettingsModel) GetLimitProcessedVmCount() int32`

GetLimitProcessedVmCount returns the LimitProcessedVmCount field if non-nil, zero value otherwise.

### GetLimitProcessedVmCountOk

`func (o *PrimaryStorageIntegrationSettingsModel) GetLimitProcessedVmCountOk() (*int32, bool)`

GetLimitProcessedVmCountOk returns a tuple with the LimitProcessedVmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitProcessedVmCount

`func (o *PrimaryStorageIntegrationSettingsModel) SetLimitProcessedVmCount(v int32)`

SetLimitProcessedVmCount sets LimitProcessedVmCount field to given value.

### HasLimitProcessedVmCount

`func (o *PrimaryStorageIntegrationSettingsModel) HasLimitProcessedVmCount() bool`

HasLimitProcessedVmCount returns a boolean if a field has been set.

### GetFailoverToStandardBackup

`func (o *PrimaryStorageIntegrationSettingsModel) GetFailoverToStandardBackup() bool`

GetFailoverToStandardBackup returns the FailoverToStandardBackup field if non-nil, zero value otherwise.

### GetFailoverToStandardBackupOk

`func (o *PrimaryStorageIntegrationSettingsModel) GetFailoverToStandardBackupOk() (*bool, bool)`

GetFailoverToStandardBackupOk returns a tuple with the FailoverToStandardBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverToStandardBackup

`func (o *PrimaryStorageIntegrationSettingsModel) SetFailoverToStandardBackup(v bool)`

SetFailoverToStandardBackup sets FailoverToStandardBackup field to given value.

### HasFailoverToStandardBackup

`func (o *PrimaryStorageIntegrationSettingsModel) HasFailoverToStandardBackup() bool`

HasFailoverToStandardBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


