# BackupIndexingSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmObject** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**WindowsIndexing** | Pointer to [**BackupObjectIndexingModel**](BackupObjectIndexingModel.md) |  | [optional] 
**LinuxIndexing** | Pointer to [**BackupObjectIndexingModel**](BackupObjectIndexingModel.md) |  | [optional] 

## Methods

### NewBackupIndexingSettingsModel

`func NewBackupIndexingSettingsModel(vmObject VmwareObjectModel, ) *BackupIndexingSettingsModel`

NewBackupIndexingSettingsModel instantiates a new BackupIndexingSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupIndexingSettingsModelWithDefaults

`func NewBackupIndexingSettingsModelWithDefaults() *BackupIndexingSettingsModel`

NewBackupIndexingSettingsModelWithDefaults instantiates a new BackupIndexingSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmObject

`func (o *BackupIndexingSettingsModel) GetVmObject() VmwareObjectModel`

GetVmObject returns the VmObject field if non-nil, zero value otherwise.

### GetVmObjectOk

`func (o *BackupIndexingSettingsModel) GetVmObjectOk() (*VmwareObjectModel, bool)`

GetVmObjectOk returns a tuple with the VmObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmObject

`func (o *BackupIndexingSettingsModel) SetVmObject(v VmwareObjectModel)`

SetVmObject sets VmObject field to given value.


### GetWindowsIndexing

`func (o *BackupIndexingSettingsModel) GetWindowsIndexing() BackupObjectIndexingModel`

GetWindowsIndexing returns the WindowsIndexing field if non-nil, zero value otherwise.

### GetWindowsIndexingOk

`func (o *BackupIndexingSettingsModel) GetWindowsIndexingOk() (*BackupObjectIndexingModel, bool)`

GetWindowsIndexingOk returns a tuple with the WindowsIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsIndexing

`func (o *BackupIndexingSettingsModel) SetWindowsIndexing(v BackupObjectIndexingModel)`

SetWindowsIndexing sets WindowsIndexing field to given value.

### HasWindowsIndexing

`func (o *BackupIndexingSettingsModel) HasWindowsIndexing() bool`

HasWindowsIndexing returns a boolean if a field has been set.

### GetLinuxIndexing

`func (o *BackupIndexingSettingsModel) GetLinuxIndexing() BackupObjectIndexingModel`

GetLinuxIndexing returns the LinuxIndexing field if non-nil, zero value otherwise.

### GetLinuxIndexingOk

`func (o *BackupIndexingSettingsModel) GetLinuxIndexingOk() (*BackupObjectIndexingModel, bool)`

GetLinuxIndexingOk returns a tuple with the LinuxIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxIndexing

`func (o *BackupIndexingSettingsModel) SetLinuxIndexing(v BackupObjectIndexingModel)`

SetLinuxIndexing sets LinuxIndexing field to given value.

### HasLinuxIndexing

`func (o *BackupIndexingSettingsModel) HasLinuxIndexing() bool`

HasLinuxIndexing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


