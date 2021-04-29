# BackupJobExclusions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vms** | Pointer to [**[]VmwareObjectSizeModel**](VmwareObjectSizeModel.md) | Array of VMs excluded from the backup. | [optional] 
**Disks** | Pointer to [**[]VmwareObjectDiskModel**](VmwareObjectDiskModel.md) | Array of VM disks excluded from the backup. | [optional] 
**Templates** | Pointer to [**BackupJobExclusionsTemplates**](BackupJobExclusionsTemplates.md) |  | [optional] 

## Methods

### NewBackupJobExclusions

`func NewBackupJobExclusions() *BackupJobExclusions`

NewBackupJobExclusions instantiates a new BackupJobExclusions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobExclusionsWithDefaults

`func NewBackupJobExclusionsWithDefaults() *BackupJobExclusions`

NewBackupJobExclusionsWithDefaults instantiates a new BackupJobExclusions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVms

`func (o *BackupJobExclusions) GetVms() []VmwareObjectSizeModel`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *BackupJobExclusions) GetVmsOk() (*[]VmwareObjectSizeModel, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *BackupJobExclusions) SetVms(v []VmwareObjectSizeModel)`

SetVms sets Vms field to given value.

### HasVms

`func (o *BackupJobExclusions) HasVms() bool`

HasVms returns a boolean if a field has been set.

### GetDisks

`func (o *BackupJobExclusions) GetDisks() []VmwareObjectDiskModel`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *BackupJobExclusions) GetDisksOk() (*[]VmwareObjectDiskModel, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *BackupJobExclusions) SetDisks(v []VmwareObjectDiskModel)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *BackupJobExclusions) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetTemplates

`func (o *BackupJobExclusions) GetTemplates() BackupJobExclusionsTemplates`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *BackupJobExclusions) GetTemplatesOk() (*BackupJobExclusionsTemplates, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *BackupJobExclusions) SetTemplates(v BackupJobExclusionsTemplates)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *BackupJobExclusions) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


