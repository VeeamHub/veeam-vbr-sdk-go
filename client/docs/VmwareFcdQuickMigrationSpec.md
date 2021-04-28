# VmwareFcdQuickMigrationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountedDiskNames** | Pointer to **[]string** | Array of disks that will be migrated to the &#x60;targetDatastore&#x60; associated with the &#x60;storagePolicy&#x60;. | [optional] 
**TargetDatastore** | [**VmwareObjectModel**](VmwareObjectModel.md) | Target datastore. For details on how to get a datastore model, see [Get VMware vSphere Server Objects](#operation/GetVmwareHostObject). | 
**StoragePolicy** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) | Storage policy that will be applied to the migrated disks. For details on how to get a storage policy model, see [Get VMware vSphere Server Objects](#operation/GetVmwareHostObject). | [optional] 

## Methods

### NewVmwareFcdQuickMigrationSpec

`func NewVmwareFcdQuickMigrationSpec(targetDatastore VmwareObjectModel, ) *VmwareFcdQuickMigrationSpec`

NewVmwareFcdQuickMigrationSpec instantiates a new VmwareFcdQuickMigrationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareFcdQuickMigrationSpecWithDefaults

`func NewVmwareFcdQuickMigrationSpecWithDefaults() *VmwareFcdQuickMigrationSpec`

NewVmwareFcdQuickMigrationSpecWithDefaults instantiates a new VmwareFcdQuickMigrationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountedDiskNames

`func (o *VmwareFcdQuickMigrationSpec) GetMountedDiskNames() []string`

GetMountedDiskNames returns the MountedDiskNames field if non-nil, zero value otherwise.

### GetMountedDiskNamesOk

`func (o *VmwareFcdQuickMigrationSpec) GetMountedDiskNamesOk() (*[]string, bool)`

GetMountedDiskNamesOk returns a tuple with the MountedDiskNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedDiskNames

`func (o *VmwareFcdQuickMigrationSpec) SetMountedDiskNames(v []string)`

SetMountedDiskNames sets MountedDiskNames field to given value.

### HasMountedDiskNames

`func (o *VmwareFcdQuickMigrationSpec) HasMountedDiskNames() bool`

HasMountedDiskNames returns a boolean if a field has been set.

### GetTargetDatastore

`func (o *VmwareFcdQuickMigrationSpec) GetTargetDatastore() VmwareObjectModel`

GetTargetDatastore returns the TargetDatastore field if non-nil, zero value otherwise.

### GetTargetDatastoreOk

`func (o *VmwareFcdQuickMigrationSpec) GetTargetDatastoreOk() (*VmwareObjectModel, bool)`

GetTargetDatastoreOk returns a tuple with the TargetDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDatastore

`func (o *VmwareFcdQuickMigrationSpec) SetTargetDatastore(v VmwareObjectModel)`

SetTargetDatastore sets TargetDatastore field to given value.


### GetStoragePolicy

`func (o *VmwareFcdQuickMigrationSpec) GetStoragePolicy() VmwareObjectModel`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *VmwareFcdQuickMigrationSpec) GetStoragePolicyOk() (*VmwareObjectModel, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *VmwareFcdQuickMigrationSpec) SetStoragePolicy(v VmwareObjectModel)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *VmwareFcdQuickMigrationSpec) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


