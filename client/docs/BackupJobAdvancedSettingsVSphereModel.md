# BackupJobAdvancedSettingsVSphereModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableVMWareToolsQuiescence** | Pointer to **bool** | If *true*, VMware Tools quiescence is enabled for freezing the VM file system and application data. | [optional] 
**ChangedBlockTracking** | Pointer to [**VSphereChangedBlockTrackingSettingsModel**](VSphereChangedBlockTrackingSettingsModel.md) |  | [optional] 

## Methods

### NewBackupJobAdvancedSettingsVSphereModel

`func NewBackupJobAdvancedSettingsVSphereModel() *BackupJobAdvancedSettingsVSphereModel`

NewBackupJobAdvancedSettingsVSphereModel instantiates a new BackupJobAdvancedSettingsVSphereModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobAdvancedSettingsVSphereModelWithDefaults

`func NewBackupJobAdvancedSettingsVSphereModelWithDefaults() *BackupJobAdvancedSettingsVSphereModel`

NewBackupJobAdvancedSettingsVSphereModelWithDefaults instantiates a new BackupJobAdvancedSettingsVSphereModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableVMWareToolsQuiescence

`func (o *BackupJobAdvancedSettingsVSphereModel) GetEnableVMWareToolsQuiescence() bool`

GetEnableVMWareToolsQuiescence returns the EnableVMWareToolsQuiescence field if non-nil, zero value otherwise.

### GetEnableVMWareToolsQuiescenceOk

`func (o *BackupJobAdvancedSettingsVSphereModel) GetEnableVMWareToolsQuiescenceOk() (*bool, bool)`

GetEnableVMWareToolsQuiescenceOk returns a tuple with the EnableVMWareToolsQuiescence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVMWareToolsQuiescence

`func (o *BackupJobAdvancedSettingsVSphereModel) SetEnableVMWareToolsQuiescence(v bool)`

SetEnableVMWareToolsQuiescence sets EnableVMWareToolsQuiescence field to given value.

### HasEnableVMWareToolsQuiescence

`func (o *BackupJobAdvancedSettingsVSphereModel) HasEnableVMWareToolsQuiescence() bool`

HasEnableVMWareToolsQuiescence returns a boolean if a field has been set.

### GetChangedBlockTracking

`func (o *BackupJobAdvancedSettingsVSphereModel) GetChangedBlockTracking() VSphereChangedBlockTrackingSettingsModel`

GetChangedBlockTracking returns the ChangedBlockTracking field if non-nil, zero value otherwise.

### GetChangedBlockTrackingOk

`func (o *BackupJobAdvancedSettingsVSphereModel) GetChangedBlockTrackingOk() (*VSphereChangedBlockTrackingSettingsModel, bool)`

GetChangedBlockTrackingOk returns a tuple with the ChangedBlockTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBlockTracking

`func (o *BackupJobAdvancedSettingsVSphereModel) SetChangedBlockTracking(v VSphereChangedBlockTrackingSettingsModel)`

SetChangedBlockTracking sets ChangedBlockTracking field to given value.

### HasChangedBlockTracking

`func (o *BackupJobAdvancedSettingsVSphereModel) HasChangedBlockTracking() bool`

HasChangedBlockTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


