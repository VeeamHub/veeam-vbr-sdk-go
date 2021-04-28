# BackupJobVirtualMachinesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Includes** | [**[]VmwareObjectModel**](VmwareObjectModel.md) |  | 
**Excludes** | Pointer to [**BackupJobExclusionsSpec**](BackupJobExclusionsSpec.md) |  | [optional] 

## Methods

### NewBackupJobVirtualMachinesSpec

`func NewBackupJobVirtualMachinesSpec(includes []VmwareObjectModel, ) *BackupJobVirtualMachinesSpec`

NewBackupJobVirtualMachinesSpec instantiates a new BackupJobVirtualMachinesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobVirtualMachinesSpecWithDefaults

`func NewBackupJobVirtualMachinesSpecWithDefaults() *BackupJobVirtualMachinesSpec`

NewBackupJobVirtualMachinesSpecWithDefaults instantiates a new BackupJobVirtualMachinesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludes

`func (o *BackupJobVirtualMachinesSpec) GetIncludes() []VmwareObjectModel`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *BackupJobVirtualMachinesSpec) GetIncludesOk() (*[]VmwareObjectModel, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *BackupJobVirtualMachinesSpec) SetIncludes(v []VmwareObjectModel)`

SetIncludes sets Includes field to given value.


### GetExcludes

`func (o *BackupJobVirtualMachinesSpec) GetExcludes() BackupJobExclusionsSpec`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *BackupJobVirtualMachinesSpec) GetExcludesOk() (*BackupJobExclusionsSpec, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *BackupJobVirtualMachinesSpec) SetExcludes(v BackupJobExclusionsSpec)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *BackupJobVirtualMachinesSpec) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


