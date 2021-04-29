# PerformanceTierAdvancedSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerVmBackup** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication creates a separate backup file for every VM in the job. | [optional] 
**FullWhenExtentOffline** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication creates a full backup file instead of an incremental backup file in case the required extent is offline. | [optional] 

## Methods

### NewPerformanceTierAdvancedSettingsModel

`func NewPerformanceTierAdvancedSettingsModel() *PerformanceTierAdvancedSettingsModel`

NewPerformanceTierAdvancedSettingsModel instantiates a new PerformanceTierAdvancedSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceTierAdvancedSettingsModelWithDefaults

`func NewPerformanceTierAdvancedSettingsModelWithDefaults() *PerformanceTierAdvancedSettingsModel`

NewPerformanceTierAdvancedSettingsModelWithDefaults instantiates a new PerformanceTierAdvancedSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerVmBackup

`func (o *PerformanceTierAdvancedSettingsModel) GetPerVmBackup() bool`

GetPerVmBackup returns the PerVmBackup field if non-nil, zero value otherwise.

### GetPerVmBackupOk

`func (o *PerformanceTierAdvancedSettingsModel) GetPerVmBackupOk() (*bool, bool)`

GetPerVmBackupOk returns a tuple with the PerVmBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerVmBackup

`func (o *PerformanceTierAdvancedSettingsModel) SetPerVmBackup(v bool)`

SetPerVmBackup sets PerVmBackup field to given value.

### HasPerVmBackup

`func (o *PerformanceTierAdvancedSettingsModel) HasPerVmBackup() bool`

HasPerVmBackup returns a boolean if a field has been set.

### GetFullWhenExtentOffline

`func (o *PerformanceTierAdvancedSettingsModel) GetFullWhenExtentOffline() bool`

GetFullWhenExtentOffline returns the FullWhenExtentOffline field if non-nil, zero value otherwise.

### GetFullWhenExtentOfflineOk

`func (o *PerformanceTierAdvancedSettingsModel) GetFullWhenExtentOfflineOk() (*bool, bool)`

GetFullWhenExtentOfflineOk returns a tuple with the FullWhenExtentOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullWhenExtentOffline

`func (o *PerformanceTierAdvancedSettingsModel) SetFullWhenExtentOffline(v bool)`

SetFullWhenExtentOffline sets FullWhenExtentOffline field to given value.

### HasFullWhenExtentOffline

`func (o *PerformanceTierAdvancedSettingsModel) HasFullWhenExtentOffline() bool`

HasFullWhenExtentOffline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


