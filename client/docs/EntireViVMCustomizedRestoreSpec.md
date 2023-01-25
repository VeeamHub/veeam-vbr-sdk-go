# EntireViVMCustomizedRestoreSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationHost** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 
**ResourcePool** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 
**Datastore** | Pointer to [**RestoreTargetDatastoreSpec**](RestoreTargetDatastoreSpec.md) |  | [optional] 
**Folder** | Pointer to [**RestoreTargetFolderSpec**](RestoreTargetFolderSpec.md) |  | [optional] 
**Network** | Pointer to [**RestoreTargetNetworkSpec**](RestoreTargetNetworkSpec.md) |  | [optional] 

## Methods

### NewEntireViVMCustomizedRestoreSpec

`func NewEntireViVMCustomizedRestoreSpec() *EntireViVMCustomizedRestoreSpec`

NewEntireViVMCustomizedRestoreSpec instantiates a new EntireViVMCustomizedRestoreSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntireViVMCustomizedRestoreSpecWithDefaults

`func NewEntireViVMCustomizedRestoreSpecWithDefaults() *EntireViVMCustomizedRestoreSpec`

NewEntireViVMCustomizedRestoreSpecWithDefaults instantiates a new EntireViVMCustomizedRestoreSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationHost

`func (o *EntireViVMCustomizedRestoreSpec) GetDestinationHost() VmwareObjectModel`

GetDestinationHost returns the DestinationHost field if non-nil, zero value otherwise.

### GetDestinationHostOk

`func (o *EntireViVMCustomizedRestoreSpec) GetDestinationHostOk() (*VmwareObjectModel, bool)`

GetDestinationHostOk returns a tuple with the DestinationHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHost

`func (o *EntireViVMCustomizedRestoreSpec) SetDestinationHost(v VmwareObjectModel)`

SetDestinationHost sets DestinationHost field to given value.

### HasDestinationHost

`func (o *EntireViVMCustomizedRestoreSpec) HasDestinationHost() bool`

HasDestinationHost returns a boolean if a field has been set.

### GetResourcePool

`func (o *EntireViVMCustomizedRestoreSpec) GetResourcePool() VmwareObjectModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *EntireViVMCustomizedRestoreSpec) GetResourcePoolOk() (*VmwareObjectModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *EntireViVMCustomizedRestoreSpec) SetResourcePool(v VmwareObjectModel)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *EntireViVMCustomizedRestoreSpec) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetDatastore

`func (o *EntireViVMCustomizedRestoreSpec) GetDatastore() RestoreTargetDatastoreSpec`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *EntireViVMCustomizedRestoreSpec) GetDatastoreOk() (*RestoreTargetDatastoreSpec, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *EntireViVMCustomizedRestoreSpec) SetDatastore(v RestoreTargetDatastoreSpec)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *EntireViVMCustomizedRestoreSpec) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetFolder

`func (o *EntireViVMCustomizedRestoreSpec) GetFolder() RestoreTargetFolderSpec`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *EntireViVMCustomizedRestoreSpec) GetFolderOk() (*RestoreTargetFolderSpec, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *EntireViVMCustomizedRestoreSpec) SetFolder(v RestoreTargetFolderSpec)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *EntireViVMCustomizedRestoreSpec) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetNetwork

`func (o *EntireViVMCustomizedRestoreSpec) GetNetwork() RestoreTargetNetworkSpec`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *EntireViVMCustomizedRestoreSpec) GetNetworkOk() (*RestoreTargetNetworkSpec, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *EntireViVMCustomizedRestoreSpec) SetNetwork(v RestoreTargetNetworkSpec)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *EntireViVMCustomizedRestoreSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


