# BackupJobVirtualMachinesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Includes** | [**[]VmwareObjectSizeModel**](VmwareObjectSizeModel.md) | Array of VM and VM containers processed by the job. | 
**Excludes** | Pointer to [**BackupJobExclusions**](BackupJobExclusions.md) |  | [optional] 

## Methods

### NewBackupJobVirtualMachinesModel

`func NewBackupJobVirtualMachinesModel(includes []VmwareObjectSizeModel, ) *BackupJobVirtualMachinesModel`

NewBackupJobVirtualMachinesModel instantiates a new BackupJobVirtualMachinesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobVirtualMachinesModelWithDefaults

`func NewBackupJobVirtualMachinesModelWithDefaults() *BackupJobVirtualMachinesModel`

NewBackupJobVirtualMachinesModelWithDefaults instantiates a new BackupJobVirtualMachinesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludes

`func (o *BackupJobVirtualMachinesModel) GetIncludes() []VmwareObjectSizeModel`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *BackupJobVirtualMachinesModel) GetIncludesOk() (*[]VmwareObjectSizeModel, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *BackupJobVirtualMachinesModel) SetIncludes(v []VmwareObjectSizeModel)`

SetIncludes sets Includes field to given value.


### GetExcludes

`func (o *BackupJobVirtualMachinesModel) GetExcludes() BackupJobExclusions`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *BackupJobVirtualMachinesModel) GetExcludesOk() (*BackupJobExclusions, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *BackupJobVirtualMachinesModel) SetExcludes(v BackupJobExclusions)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *BackupJobVirtualMachinesModel) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


