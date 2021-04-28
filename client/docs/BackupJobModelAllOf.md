# BackupJobModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsHighPriority** | **bool** | If *true*, the job has a high priority in getting backup infrastructure resources. | 
**VirtualMachines** | [**BackupJobVirtualMachinesModel**](BackupJobVirtualMachinesModel.md) |  | 
**Storage** | [**BackupJobStorageModel**](BackupJobStorageModel.md) |  | 
**GuestProcessing** | [**BackupJobGuestProcessingModel**](BackupJobGuestProcessingModel.md) |  | 
**Schedule** | [**BackupScheduleModel**](BackupScheduleModel.md) |  | 

## Methods

### NewBackupJobModelAllOf

`func NewBackupJobModelAllOf(isHighPriority bool, virtualMachines BackupJobVirtualMachinesModel, storage BackupJobStorageModel, guestProcessing BackupJobGuestProcessingModel, schedule BackupScheduleModel, ) *BackupJobModelAllOf`

NewBackupJobModelAllOf instantiates a new BackupJobModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobModelAllOfWithDefaults

`func NewBackupJobModelAllOfWithDefaults() *BackupJobModelAllOf`

NewBackupJobModelAllOfWithDefaults instantiates a new BackupJobModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsHighPriority

`func (o *BackupJobModelAllOf) GetIsHighPriority() bool`

GetIsHighPriority returns the IsHighPriority field if non-nil, zero value otherwise.

### GetIsHighPriorityOk

`func (o *BackupJobModelAllOf) GetIsHighPriorityOk() (*bool, bool)`

GetIsHighPriorityOk returns a tuple with the IsHighPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHighPriority

`func (o *BackupJobModelAllOf) SetIsHighPriority(v bool)`

SetIsHighPriority sets IsHighPriority field to given value.


### GetVirtualMachines

`func (o *BackupJobModelAllOf) GetVirtualMachines() BackupJobVirtualMachinesModel`

GetVirtualMachines returns the VirtualMachines field if non-nil, zero value otherwise.

### GetVirtualMachinesOk

`func (o *BackupJobModelAllOf) GetVirtualMachinesOk() (*BackupJobVirtualMachinesModel, bool)`

GetVirtualMachinesOk returns a tuple with the VirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachines

`func (o *BackupJobModelAllOf) SetVirtualMachines(v BackupJobVirtualMachinesModel)`

SetVirtualMachines sets VirtualMachines field to given value.


### GetStorage

`func (o *BackupJobModelAllOf) GetStorage() BackupJobStorageModel`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *BackupJobModelAllOf) GetStorageOk() (*BackupJobStorageModel, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *BackupJobModelAllOf) SetStorage(v BackupJobStorageModel)`

SetStorage sets Storage field to given value.


### GetGuestProcessing

`func (o *BackupJobModelAllOf) GetGuestProcessing() BackupJobGuestProcessingModel`

GetGuestProcessing returns the GuestProcessing field if non-nil, zero value otherwise.

### GetGuestProcessingOk

`func (o *BackupJobModelAllOf) GetGuestProcessingOk() (*BackupJobGuestProcessingModel, bool)`

GetGuestProcessingOk returns a tuple with the GuestProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestProcessing

`func (o *BackupJobModelAllOf) SetGuestProcessing(v BackupJobGuestProcessingModel)`

SetGuestProcessing sets GuestProcessing field to given value.


### GetSchedule

`func (o *BackupJobModelAllOf) GetSchedule() BackupScheduleModel`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *BackupJobModelAllOf) GetScheduleOk() (*BackupScheduleModel, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *BackupJobModelAllOf) SetSchedule(v BackupScheduleModel)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


