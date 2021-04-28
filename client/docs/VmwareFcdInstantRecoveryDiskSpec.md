# VmwareFcdInstantRecoveryDiskSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NameInBackup** | **string** | Disk name displayed in the backup. | 
**MountedDiskName** | **string** | Name of the VMDK file that will be stored in the datastore. | 
**RegisteredFcdName** | **string** | Name under which the disk will be registered as an FCD in the vCenter Managed Object Browser. | 

## Methods

### NewVmwareFcdInstantRecoveryDiskSpec

`func NewVmwareFcdInstantRecoveryDiskSpec(nameInBackup string, mountedDiskName string, registeredFcdName string, ) *VmwareFcdInstantRecoveryDiskSpec`

NewVmwareFcdInstantRecoveryDiskSpec instantiates a new VmwareFcdInstantRecoveryDiskSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareFcdInstantRecoveryDiskSpecWithDefaults

`func NewVmwareFcdInstantRecoveryDiskSpecWithDefaults() *VmwareFcdInstantRecoveryDiskSpec`

NewVmwareFcdInstantRecoveryDiskSpecWithDefaults instantiates a new VmwareFcdInstantRecoveryDiskSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameInBackup

`func (o *VmwareFcdInstantRecoveryDiskSpec) GetNameInBackup() string`

GetNameInBackup returns the NameInBackup field if non-nil, zero value otherwise.

### GetNameInBackupOk

`func (o *VmwareFcdInstantRecoveryDiskSpec) GetNameInBackupOk() (*string, bool)`

GetNameInBackupOk returns a tuple with the NameInBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameInBackup

`func (o *VmwareFcdInstantRecoveryDiskSpec) SetNameInBackup(v string)`

SetNameInBackup sets NameInBackup field to given value.


### GetMountedDiskName

`func (o *VmwareFcdInstantRecoveryDiskSpec) GetMountedDiskName() string`

GetMountedDiskName returns the MountedDiskName field if non-nil, zero value otherwise.

### GetMountedDiskNameOk

`func (o *VmwareFcdInstantRecoveryDiskSpec) GetMountedDiskNameOk() (*string, bool)`

GetMountedDiskNameOk returns a tuple with the MountedDiskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedDiskName

`func (o *VmwareFcdInstantRecoveryDiskSpec) SetMountedDiskName(v string)`

SetMountedDiskName sets MountedDiskName field to given value.


### GetRegisteredFcdName

`func (o *VmwareFcdInstantRecoveryDiskSpec) GetRegisteredFcdName() string`

GetRegisteredFcdName returns the RegisteredFcdName field if non-nil, zero value otherwise.

### GetRegisteredFcdNameOk

`func (o *VmwareFcdInstantRecoveryDiskSpec) GetRegisteredFcdNameOk() (*string, bool)`

GetRegisteredFcdNameOk returns a tuple with the RegisteredFcdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredFcdName

`func (o *VmwareFcdInstantRecoveryDiskSpec) SetRegisteredFcdName(v string)`

SetRegisteredFcdName sets RegisteredFcdName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


