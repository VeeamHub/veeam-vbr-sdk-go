# ViVMQuickMigrationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationHost** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**ResourcePool** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 
**Folder** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 
**Datastore** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 
**SourceProxyIds** | Pointer to **[]string** | Array of source backup proxies. | [optional] 
**TargetProxyIds** | Pointer to **[]string** | Array of target backup proxies. | [optional] 
**VeeamQMEnabled** | Pointer to **bool** | If *true*, the Veeam Quick Migration mechanism is used. Otherwise, Veeam Backup &amp; Replication will use VMware vMotion for migration. | [optional] 
**DeleteSourceVmsFiles** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication will delete source VM files upon successful migration. | [optional] 

## Methods

### NewViVMQuickMigrationSpec

`func NewViVMQuickMigrationSpec(destinationHost VmwareObjectModel, ) *ViVMQuickMigrationSpec`

NewViVMQuickMigrationSpec instantiates a new ViVMQuickMigrationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViVMQuickMigrationSpecWithDefaults

`func NewViVMQuickMigrationSpecWithDefaults() *ViVMQuickMigrationSpec`

NewViVMQuickMigrationSpecWithDefaults instantiates a new ViVMQuickMigrationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationHost

`func (o *ViVMQuickMigrationSpec) GetDestinationHost() VmwareObjectModel`

GetDestinationHost returns the DestinationHost field if non-nil, zero value otherwise.

### GetDestinationHostOk

`func (o *ViVMQuickMigrationSpec) GetDestinationHostOk() (*VmwareObjectModel, bool)`

GetDestinationHostOk returns a tuple with the DestinationHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHost

`func (o *ViVMQuickMigrationSpec) SetDestinationHost(v VmwareObjectModel)`

SetDestinationHost sets DestinationHost field to given value.


### GetResourcePool

`func (o *ViVMQuickMigrationSpec) GetResourcePool() VmwareObjectModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *ViVMQuickMigrationSpec) GetResourcePoolOk() (*VmwareObjectModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *ViVMQuickMigrationSpec) SetResourcePool(v VmwareObjectModel)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *ViVMQuickMigrationSpec) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetFolder

`func (o *ViVMQuickMigrationSpec) GetFolder() VmwareObjectModel`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *ViVMQuickMigrationSpec) GetFolderOk() (*VmwareObjectModel, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *ViVMQuickMigrationSpec) SetFolder(v VmwareObjectModel)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *ViVMQuickMigrationSpec) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetDatastore

`func (o *ViVMQuickMigrationSpec) GetDatastore() VmwareObjectModel`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ViVMQuickMigrationSpec) GetDatastoreOk() (*VmwareObjectModel, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ViVMQuickMigrationSpec) SetDatastore(v VmwareObjectModel)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *ViVMQuickMigrationSpec) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetSourceProxyIds

`func (o *ViVMQuickMigrationSpec) GetSourceProxyIds() []string`

GetSourceProxyIds returns the SourceProxyIds field if non-nil, zero value otherwise.

### GetSourceProxyIdsOk

`func (o *ViVMQuickMigrationSpec) GetSourceProxyIdsOk() (*[]string, bool)`

GetSourceProxyIdsOk returns a tuple with the SourceProxyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProxyIds

`func (o *ViVMQuickMigrationSpec) SetSourceProxyIds(v []string)`

SetSourceProxyIds sets SourceProxyIds field to given value.

### HasSourceProxyIds

`func (o *ViVMQuickMigrationSpec) HasSourceProxyIds() bool`

HasSourceProxyIds returns a boolean if a field has been set.

### GetTargetProxyIds

`func (o *ViVMQuickMigrationSpec) GetTargetProxyIds() []string`

GetTargetProxyIds returns the TargetProxyIds field if non-nil, zero value otherwise.

### GetTargetProxyIdsOk

`func (o *ViVMQuickMigrationSpec) GetTargetProxyIdsOk() (*[]string, bool)`

GetTargetProxyIdsOk returns a tuple with the TargetProxyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProxyIds

`func (o *ViVMQuickMigrationSpec) SetTargetProxyIds(v []string)`

SetTargetProxyIds sets TargetProxyIds field to given value.

### HasTargetProxyIds

`func (o *ViVMQuickMigrationSpec) HasTargetProxyIds() bool`

HasTargetProxyIds returns a boolean if a field has been set.

### GetVeeamQMEnabled

`func (o *ViVMQuickMigrationSpec) GetVeeamQMEnabled() bool`

GetVeeamQMEnabled returns the VeeamQMEnabled field if non-nil, zero value otherwise.

### GetVeeamQMEnabledOk

`func (o *ViVMQuickMigrationSpec) GetVeeamQMEnabledOk() (*bool, bool)`

GetVeeamQMEnabledOk returns a tuple with the VeeamQMEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVeeamQMEnabled

`func (o *ViVMQuickMigrationSpec) SetVeeamQMEnabled(v bool)`

SetVeeamQMEnabled sets VeeamQMEnabled field to given value.

### HasVeeamQMEnabled

`func (o *ViVMQuickMigrationSpec) HasVeeamQMEnabled() bool`

HasVeeamQMEnabled returns a boolean if a field has been set.

### GetDeleteSourceVmsFiles

`func (o *ViVMQuickMigrationSpec) GetDeleteSourceVmsFiles() bool`

GetDeleteSourceVmsFiles returns the DeleteSourceVmsFiles field if non-nil, zero value otherwise.

### GetDeleteSourceVmsFilesOk

`func (o *ViVMQuickMigrationSpec) GetDeleteSourceVmsFilesOk() (*bool, bool)`

GetDeleteSourceVmsFilesOk returns a tuple with the DeleteSourceVmsFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSourceVmsFiles

`func (o *ViVMQuickMigrationSpec) SetDeleteSourceVmsFiles(v bool)`

SetDeleteSourceVmsFiles sets DeleteSourceVmsFiles field to given value.

### HasDeleteSourceVmsFiles

`func (o *ViVMQuickMigrationSpec) HasDeleteSourceVmsFiles() bool`

HasDeleteSourceVmsFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


