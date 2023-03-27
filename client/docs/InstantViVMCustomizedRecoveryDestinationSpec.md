# InstantViVMCustomizedRecoveryDestinationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoredVmName** | **string** | Restored VM name. | 
**DestinationHost** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**Folder** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**ResourcePool** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**BiosUuidPolicy** | [**EInstantViVmRecoveryBiosUuidPolicyType**](EInstantViVmRecoveryBiosUuidPolicyType.md) |  | 

## Methods

### NewInstantViVMCustomizedRecoveryDestinationSpec

`func NewInstantViVMCustomizedRecoveryDestinationSpec(restoredVmName string, destinationHost VmwareObjectModel, folder VmwareObjectModel, resourcePool VmwareObjectModel, biosUuidPolicy EInstantViVmRecoveryBiosUuidPolicyType, ) *InstantViVMCustomizedRecoveryDestinationSpec`

NewInstantViVMCustomizedRecoveryDestinationSpec instantiates a new InstantViVMCustomizedRecoveryDestinationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantViVMCustomizedRecoveryDestinationSpecWithDefaults

`func NewInstantViVMCustomizedRecoveryDestinationSpecWithDefaults() *InstantViVMCustomizedRecoveryDestinationSpec`

NewInstantViVMCustomizedRecoveryDestinationSpecWithDefaults instantiates a new InstantViVMCustomizedRecoveryDestinationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoredVmName

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) GetRestoredVmName() string`

GetRestoredVmName returns the RestoredVmName field if non-nil, zero value otherwise.

### GetRestoredVmNameOk

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) GetRestoredVmNameOk() (*string, bool)`

GetRestoredVmNameOk returns a tuple with the RestoredVmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredVmName

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) SetRestoredVmName(v string)`

SetRestoredVmName sets RestoredVmName field to given value.


### GetDestinationHost

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) GetDestinationHost() VmwareObjectModel`

GetDestinationHost returns the DestinationHost field if non-nil, zero value otherwise.

### GetDestinationHostOk

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) GetDestinationHostOk() (*VmwareObjectModel, bool)`

GetDestinationHostOk returns a tuple with the DestinationHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHost

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) SetDestinationHost(v VmwareObjectModel)`

SetDestinationHost sets DestinationHost field to given value.


### GetFolder

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) GetFolder() VmwareObjectModel`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) GetFolderOk() (*VmwareObjectModel, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) SetFolder(v VmwareObjectModel)`

SetFolder sets Folder field to given value.


### GetResourcePool

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) GetResourcePool() VmwareObjectModel`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) GetResourcePoolOk() (*VmwareObjectModel, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) SetResourcePool(v VmwareObjectModel)`

SetResourcePool sets ResourcePool field to given value.


### GetBiosUuidPolicy

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) GetBiosUuidPolicy() EInstantViVmRecoveryBiosUuidPolicyType`

GetBiosUuidPolicy returns the BiosUuidPolicy field if non-nil, zero value otherwise.

### GetBiosUuidPolicyOk

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) GetBiosUuidPolicyOk() (*EInstantViVmRecoveryBiosUuidPolicyType, bool)`

GetBiosUuidPolicyOk returns a tuple with the BiosUuidPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuidPolicy

`func (o *InstantViVMCustomizedRecoveryDestinationSpec) SetBiosUuidPolicy(v EInstantViVmRecoveryBiosUuidPolicyType)`

SetBiosUuidPolicy sets BiosUuidPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


