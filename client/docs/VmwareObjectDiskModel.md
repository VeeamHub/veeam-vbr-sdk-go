# VmwareObjectDiskModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmObject** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**DisksToProcess** | [**EVmwareDisksTypeToProcess**](EVmwareDisksTypeToProcess.md) |  | 
**Disks** | **[]string** | Array of disks. | 
**RemoveFromVMConfiguration** | Pointer to **bool** | If *true*, the disk is removed from VM configuration. | [optional] 

## Methods

### NewVmwareObjectDiskModel

`func NewVmwareObjectDiskModel(vmObject VmwareObjectModel, disksToProcess EVmwareDisksTypeToProcess, disks []string, ) *VmwareObjectDiskModel`

NewVmwareObjectDiskModel instantiates a new VmwareObjectDiskModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectDiskModelWithDefaults

`func NewVmwareObjectDiskModelWithDefaults() *VmwareObjectDiskModel`

NewVmwareObjectDiskModelWithDefaults instantiates a new VmwareObjectDiskModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmObject

`func (o *VmwareObjectDiskModel) GetVmObject() VmwareObjectModel`

GetVmObject returns the VmObject field if non-nil, zero value otherwise.

### GetVmObjectOk

`func (o *VmwareObjectDiskModel) GetVmObjectOk() (*VmwareObjectModel, bool)`

GetVmObjectOk returns a tuple with the VmObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmObject

`func (o *VmwareObjectDiskModel) SetVmObject(v VmwareObjectModel)`

SetVmObject sets VmObject field to given value.


### GetDisksToProcess

`func (o *VmwareObjectDiskModel) GetDisksToProcess() EVmwareDisksTypeToProcess`

GetDisksToProcess returns the DisksToProcess field if non-nil, zero value otherwise.

### GetDisksToProcessOk

`func (o *VmwareObjectDiskModel) GetDisksToProcessOk() (*EVmwareDisksTypeToProcess, bool)`

GetDisksToProcessOk returns a tuple with the DisksToProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisksToProcess

`func (o *VmwareObjectDiskModel) SetDisksToProcess(v EVmwareDisksTypeToProcess)`

SetDisksToProcess sets DisksToProcess field to given value.


### GetDisks

`func (o *VmwareObjectDiskModel) GetDisks() []string`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VmwareObjectDiskModel) GetDisksOk() (*[]string, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VmwareObjectDiskModel) SetDisks(v []string)`

SetDisks sets Disks field to given value.


### GetRemoveFromVMConfiguration

`func (o *VmwareObjectDiskModel) GetRemoveFromVMConfiguration() bool`

GetRemoveFromVMConfiguration returns the RemoveFromVMConfiguration field if non-nil, zero value otherwise.

### GetRemoveFromVMConfigurationOk

`func (o *VmwareObjectDiskModel) GetRemoveFromVMConfigurationOk() (*bool, bool)`

GetRemoveFromVMConfigurationOk returns a tuple with the RemoveFromVMConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFromVMConfiguration

`func (o *VmwareObjectDiskModel) SetRemoveFromVMConfiguration(v bool)`

SetRemoveFromVMConfiguration sets RemoveFromVMConfiguration field to given value.

### HasRemoveFromVMConfiguration

`func (o *VmwareObjectDiskModel) HasRemoveFromVMConfiguration() bool`

HasRemoveFromVMConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


