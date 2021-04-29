# BackupJobExclusionsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vms** | Pointer to [**[]VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 
**Disks** | Pointer to [**[]VmwareObjectDiskModel**](VmwareObjectDiskModel.md) |  | [optional] 
**Templates** | Pointer to [**BackupJobExclusionsTemplates**](BackupJobExclusionsTemplates.md) |  | [optional] 

## Methods

### NewBackupJobExclusionsSpec

`func NewBackupJobExclusionsSpec() *BackupJobExclusionsSpec`

NewBackupJobExclusionsSpec instantiates a new BackupJobExclusionsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobExclusionsSpecWithDefaults

`func NewBackupJobExclusionsSpecWithDefaults() *BackupJobExclusionsSpec`

NewBackupJobExclusionsSpecWithDefaults instantiates a new BackupJobExclusionsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVms

`func (o *BackupJobExclusionsSpec) GetVms() []VmwareObjectModel`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *BackupJobExclusionsSpec) GetVmsOk() (*[]VmwareObjectModel, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *BackupJobExclusionsSpec) SetVms(v []VmwareObjectModel)`

SetVms sets Vms field to given value.

### HasVms

`func (o *BackupJobExclusionsSpec) HasVms() bool`

HasVms returns a boolean if a field has been set.

### GetDisks

`func (o *BackupJobExclusionsSpec) GetDisks() []VmwareObjectDiskModel`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *BackupJobExclusionsSpec) GetDisksOk() (*[]VmwareObjectDiskModel, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *BackupJobExclusionsSpec) SetDisks(v []VmwareObjectDiskModel)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *BackupJobExclusionsSpec) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetTemplates

`func (o *BackupJobExclusionsSpec) GetTemplates() BackupJobExclusionsTemplates`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *BackupJobExclusionsSpec) GetTemplatesOk() (*BackupJobExclusionsTemplates, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *BackupJobExclusionsSpec) SetTemplates(v BackupJobExclusionsTemplates)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *BackupJobExclusionsSpec) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


