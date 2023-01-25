# EntireViVMRestoreSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectRestorePointId** | **string** | ID of the restore point. | 
**Type** | [**EEntireVMRestoreModeType**](EEntireVMRestoreModeType.md) |  | 
**RestoreProxies** | Pointer to [**RestoreProxySpec**](RestoreProxySpec.md) |  | [optional] 
**SecureRestore** | Pointer to [**SecureRestoreSpec**](SecureRestoreSpec.md) |  | [optional] 
**PowerUp** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication starts the restored VM on the target host. | [optional] 
**Reason** | Pointer to **string** | Reason for restoring the VM. | [optional] 
**QuickRollback** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication performs incremental restore. | [optional] 
**DestinationHost** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 
**ResourcePool** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 
**Datastore** | Pointer to [**RestoreTargetDatastoreSpec**](RestoreTargetDatastoreSpec.md) |  | [optional] 
**Folder** | Pointer to [**RestoreTargetFolderSpec**](RestoreTargetFolderSpec.md) |  | [optional] 
**Network** | Pointer to [**RestoreTargetNetworkSpec**](RestoreTargetNetworkSpec.md) |  | [optional] 

## Methods

### NewEntireViVMRestoreSpec

`func NewEntireViVMRestoreSpec(objectRestorePointId string, type_ EEntireVMRestoreModeType, ) *EntireViVMRestoreSpec`

NewEntireViVMRestoreSpec instantiates a new EntireViVMRestoreSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntireViVMRestoreSpecWithDefaults

`func NewEntireViVMRestoreSpecWithDefaults() *EntireViVMRestoreSpec`

NewEntireViVMRestoreSpecWithDefaults instantiates a new EntireViVMRestoreSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectRestorePointId

`func (o *EntireViVMRestoreSpec) GetObjectRestorePointId() string`

GetObjectRestorePointId returns the ObjectRestorePointId field if non-nil, zero value otherwise.

### GetObjectRestorePointIdOk

`func (o *EntireViVMRestoreSpec) GetObjectRestorePointIdOk() (*string, bool)`

GetObjectRestorePointIdOk returns a tuple with the ObjectRestorePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRestorePointId

`func (o *EntireViVMRestoreSpec) SetObjectRestorePointId(v string)`

SetObjectRestorePointId sets ObjectRestorePointId field to given value.


### GetType

`func (o *EntireViVMRestoreSpec) GetType() EEntireVMRestoreModeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntireViVMRestoreSpec) GetTypeOk() (*EEntireVMRestoreModeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntireViVMRestoreSpec) SetType(v EEntireVMRestoreModeType)`

SetType sets Type field to given value.


### GetRestoreProxies

`func (o *EntireViVMRestoreSpec) GetRestoreProxies() RestoreProxySpec`

GetRestoreProxies returns the RestoreProxies field if non-nil, zero value otherwise.

### GetRestoreProxiesOk

`func (o *EntireViVMRestoreSpec) GetRestoreProxiesOk() (*RestoreProxySpec, bool)`

GetRestoreProxiesOk returns a tuple with the RestoreProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreProxies

`func (o *EntireViVMRestoreSpec) SetRestoreProxies(v RestoreProxySpec)`

SetRestoreProxies sets RestoreProxies field to given value.

### HasRestoreProxies

`func (o *EntireViVMRestoreSpec) HasRestoreProxies() bool`

HasRestoreProxies returns a boolean if a field has been set.

### GetSecureRestore

`func (o *EntireViVMRestoreSpec) GetSecureRestore() SecureRestoreSpec`

GetSecureRestore returns the SecureRestore field if non-nil, zero value otherwise.

### GetSecureRestoreOk

`func (o *EntireViVMRestoreSpec) GetSecureRestoreOk() (*SecureRestoreSpec, bool)`

GetSecureRestoreOk returns a tuple with the SecureRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureRestore

`func (o *EntireViVMRestoreSpec) SetSecureRestore(v SecureRestoreSpec)`

SetSecureRestore sets SecureRestore field to given value.

### HasSecureRestore

`func (o *EntireViVMRestoreSpec) HasSecureRestore() bool`

HasSecureRestore returns a boolean if a field has been set.

### GetPowerUp

`func (o *EntireViVMRestoreSpec) GetPowerUp() bool`

GetPowerUp returns the PowerUp field if non-nil, zero value otherwise.

### GetPowerUpOk

`func (o *EntireViVMRestoreSpec) GetPowerUpOk() (*bool, bool)`

GetPowerUpOk returns a tuple with the PowerUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerUp

`func (o *EntireViVMRestoreSpec) SetPowerUp(v bool)`

SetPowerUp sets PowerUp field to given value.

### HasPowerUp

`func (o *EntireViVMRestoreSpec) HasPowerUp() bool`

HasPowerUp returns a boolean if a field has been set.

### GetReason

`func (o *EntireViVMRestoreSpec) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EntireViVMRestoreSpec) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EntireViVMRestoreSpec) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EntireViVMRestoreSpec) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetQuickRollback

`func (o *EntireViVMRestoreSpec) GetQuickRollback() bool`

GetQuickRollback returns the QuickRollback field if non-nil, zero value otherwise.

### GetQuickRollbackOk

`func (o *EntireViVMRestoreSpec) GetQuickRollbackOk() (*bool, bool)`

GetQuickRollbackOk returns a tuple with the QuickRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickRollback

`func (o *EntireViVMRestoreSpec) SetQuickRollback(v bool)`

SetQuickRollback sets QuickRollback field to given value.

### HasQuickRollback

`func (o *EntireViVMRestoreSpec) HasQuickRollback() bool`

HasQuickRollback returns a boolean if a field has been set.

### GetDestinationHost

`func (o *EntireViVMRestoreSpec) GetDestinationHost() VmwareObjectModel`

GetDestinationHost returns the DestinationHost field if non-nil, zero value otherwise.

### GetDestinationHostOk

`func (o *EntireViVMRestoreSpec) GetDestinationHostOk() (*VmwareObjectModel, bool)`

GetDestinationHostOk returns a tuple with the DestinationHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHost

`func (o *EntireViVMRestoreSpec) SetDestinationHost(v VmwareObjectModel)`

SetDestinationHost sets DestinationHost field to given value.

### HasDestinationHost

`func (o *EntireViVMRestoreSpec) HasDestinationHost() bool`

HasDestinationHost returns a boolean if a field has been set.

### GetResourcePool

`func (o *EntireViVMRestoreSpec) GetResourcePool() VmwareObjectModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *EntireViVMRestoreSpec) GetResourcePoolOk() (*VmwareObjectModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *EntireViVMRestoreSpec) SetResourcePool(v VmwareObjectModel)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *EntireViVMRestoreSpec) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetDatastore

`func (o *EntireViVMRestoreSpec) GetDatastore() RestoreTargetDatastoreSpec`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *EntireViVMRestoreSpec) GetDatastoreOk() (*RestoreTargetDatastoreSpec, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *EntireViVMRestoreSpec) SetDatastore(v RestoreTargetDatastoreSpec)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *EntireViVMRestoreSpec) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetFolder

`func (o *EntireViVMRestoreSpec) GetFolder() RestoreTargetFolderSpec`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *EntireViVMRestoreSpec) GetFolderOk() (*RestoreTargetFolderSpec, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *EntireViVMRestoreSpec) SetFolder(v RestoreTargetFolderSpec)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *EntireViVMRestoreSpec) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetNetwork

`func (o *EntireViVMRestoreSpec) GetNetwork() RestoreTargetNetworkSpec`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *EntireViVMRestoreSpec) GetNetworkOk() (*RestoreTargetNetworkSpec, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *EntireViVMRestoreSpec) SetNetwork(v RestoreTargetNetworkSpec)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *EntireViVMRestoreSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


