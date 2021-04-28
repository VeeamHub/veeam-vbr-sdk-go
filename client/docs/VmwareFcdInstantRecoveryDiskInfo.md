# VmwareFcdInstantRecoveryDiskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NameInBackup** | **string** | Disk name displayed in the backup. | 
**MountedDiskName** | **string** | Name of the VMDK file that is stored in the datastore. | 
**RegisteredFcdName** | **string** | Name under which the disk is registered as an FCD in the vCenter Managed Object Browser. | 
**ObjectId** | **string** | FCD ID. | 

## Methods

### NewVmwareFcdInstantRecoveryDiskInfo

`func NewVmwareFcdInstantRecoveryDiskInfo(nameInBackup string, mountedDiskName string, registeredFcdName string, objectId string, ) *VmwareFcdInstantRecoveryDiskInfo`

NewVmwareFcdInstantRecoveryDiskInfo instantiates a new VmwareFcdInstantRecoveryDiskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareFcdInstantRecoveryDiskInfoWithDefaults

`func NewVmwareFcdInstantRecoveryDiskInfoWithDefaults() *VmwareFcdInstantRecoveryDiskInfo`

NewVmwareFcdInstantRecoveryDiskInfoWithDefaults instantiates a new VmwareFcdInstantRecoveryDiskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameInBackup

`func (o *VmwareFcdInstantRecoveryDiskInfo) GetNameInBackup() string`

GetNameInBackup returns the NameInBackup field if non-nil, zero value otherwise.

### GetNameInBackupOk

`func (o *VmwareFcdInstantRecoveryDiskInfo) GetNameInBackupOk() (*string, bool)`

GetNameInBackupOk returns a tuple with the NameInBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameInBackup

`func (o *VmwareFcdInstantRecoveryDiskInfo) SetNameInBackup(v string)`

SetNameInBackup sets NameInBackup field to given value.


### GetMountedDiskName

`func (o *VmwareFcdInstantRecoveryDiskInfo) GetMountedDiskName() string`

GetMountedDiskName returns the MountedDiskName field if non-nil, zero value otherwise.

### GetMountedDiskNameOk

`func (o *VmwareFcdInstantRecoveryDiskInfo) GetMountedDiskNameOk() (*string, bool)`

GetMountedDiskNameOk returns a tuple with the MountedDiskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedDiskName

`func (o *VmwareFcdInstantRecoveryDiskInfo) SetMountedDiskName(v string)`

SetMountedDiskName sets MountedDiskName field to given value.


### GetRegisteredFcdName

`func (o *VmwareFcdInstantRecoveryDiskInfo) GetRegisteredFcdName() string`

GetRegisteredFcdName returns the RegisteredFcdName field if non-nil, zero value otherwise.

### GetRegisteredFcdNameOk

`func (o *VmwareFcdInstantRecoveryDiskInfo) GetRegisteredFcdNameOk() (*string, bool)`

GetRegisteredFcdNameOk returns a tuple with the RegisteredFcdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredFcdName

`func (o *VmwareFcdInstantRecoveryDiskInfo) SetRegisteredFcdName(v string)`

SetRegisteredFcdName sets RegisteredFcdName field to given value.


### GetObjectId

`func (o *VmwareFcdInstantRecoveryDiskInfo) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *VmwareFcdInstantRecoveryDiskInfo) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *VmwareFcdInstantRecoveryDiskInfo) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


