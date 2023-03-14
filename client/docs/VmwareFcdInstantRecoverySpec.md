# VmwareFcdInstantRecoverySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectRestorePointId** | **string** | ID of the restore point. | 
**DestinationCluster** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**DisksMapping** | [**[]VmwareFcdInstantRecoveryDiskSpec**](VmwareFcdInstantRecoveryDiskSpec.md) | Array of disks for restore. | 
**WriteCache** | Pointer to [**VmwareFcdWriteCacheSpec**](VmwareFcdWriteCacheSpec.md) |  | [optional] 

## Methods

### NewVmwareFcdInstantRecoverySpec

`func NewVmwareFcdInstantRecoverySpec(objectRestorePointId string, destinationCluster VmwareObjectModel, disksMapping []VmwareFcdInstantRecoveryDiskSpec, ) *VmwareFcdInstantRecoverySpec`

NewVmwareFcdInstantRecoverySpec instantiates a new VmwareFcdInstantRecoverySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareFcdInstantRecoverySpecWithDefaults

`func NewVmwareFcdInstantRecoverySpecWithDefaults() *VmwareFcdInstantRecoverySpec`

NewVmwareFcdInstantRecoverySpecWithDefaults instantiates a new VmwareFcdInstantRecoverySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectRestorePointId

`func (o *VmwareFcdInstantRecoverySpec) GetObjectRestorePointId() string`

GetObjectRestorePointId returns the ObjectRestorePointId field if non-nil, zero value otherwise.

### GetObjectRestorePointIdOk

`func (o *VmwareFcdInstantRecoverySpec) GetObjectRestorePointIdOk() (*string, bool)`

GetObjectRestorePointIdOk returns a tuple with the ObjectRestorePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRestorePointId

`func (o *VmwareFcdInstantRecoverySpec) SetObjectRestorePointId(v string)`

SetObjectRestorePointId sets ObjectRestorePointId field to given value.


### GetDestinationCluster

`func (o *VmwareFcdInstantRecoverySpec) GetDestinationCluster() VmwareObjectModel`

GetDestinationCluster returns the DestinationCluster field if non-nil, zero value otherwise.

### GetDestinationClusterOk

`func (o *VmwareFcdInstantRecoverySpec) GetDestinationClusterOk() (*VmwareObjectModel, bool)`

GetDestinationClusterOk returns a tuple with the DestinationCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCluster

`func (o *VmwareFcdInstantRecoverySpec) SetDestinationCluster(v VmwareObjectModel)`

SetDestinationCluster sets DestinationCluster field to given value.


### GetDisksMapping

`func (o *VmwareFcdInstantRecoverySpec) GetDisksMapping() []VmwareFcdInstantRecoveryDiskSpec`

GetDisksMapping returns the DisksMapping field if non-nil, zero value otherwise.

### GetDisksMappingOk

`func (o *VmwareFcdInstantRecoverySpec) GetDisksMappingOk() (*[]VmwareFcdInstantRecoveryDiskSpec, bool)`

GetDisksMappingOk returns a tuple with the DisksMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisksMapping

`func (o *VmwareFcdInstantRecoverySpec) SetDisksMapping(v []VmwareFcdInstantRecoveryDiskSpec)`

SetDisksMapping sets DisksMapping field to given value.


### GetWriteCache

`func (o *VmwareFcdInstantRecoverySpec) GetWriteCache() VmwareFcdWriteCacheSpec`

GetWriteCache returns the WriteCache field if non-nil, zero value otherwise.

### GetWriteCacheOk

`func (o *VmwareFcdInstantRecoverySpec) GetWriteCacheOk() (*VmwareFcdWriteCacheSpec, bool)`

GetWriteCacheOk returns a tuple with the WriteCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCache

`func (o *VmwareFcdInstantRecoverySpec) SetWriteCache(v VmwareFcdWriteCacheSpec)`

SetWriteCache sets WriteCache field to given value.

### HasWriteCache

`func (o *VmwareFcdInstantRecoverySpec) HasWriteCache() bool`

HasWriteCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


