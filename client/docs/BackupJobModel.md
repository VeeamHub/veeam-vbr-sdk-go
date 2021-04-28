# BackupJobModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsHighPriority** | **bool** | If *true*, the job has a high priority in getting backup infrastructure resources. | 
**VirtualMachines** | [**BackupJobVirtualMachinesModel**](BackupJobVirtualMachinesModel.md) |  | 
**Storage** | [**BackupJobStorageModel**](BackupJobStorageModel.md) |  | 
**GuestProcessing** | [**BackupJobGuestProcessingModel**](BackupJobGuestProcessingModel.md) |  | 
**Schedule** | [**BackupScheduleModel**](BackupScheduleModel.md) |  | 

## Methods

### NewBackupJobModel

`func NewBackupJobModel(isHighPriority bool, virtualMachines BackupJobVirtualMachinesModel, storage BackupJobStorageModel, guestProcessing BackupJobGuestProcessingModel, schedule BackupScheduleModel, ) *BackupJobModel`

NewBackupJobModel instantiates a new BackupJobModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobModelWithDefaults

`func NewBackupJobModelWithDefaults() *BackupJobModel`

NewBackupJobModelWithDefaults instantiates a new BackupJobModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsHighPriority

`func (o *BackupJobModel) GetIsHighPriority() bool`

GetIsHighPriority returns the IsHighPriority field if non-nil, zero value otherwise.

### GetIsHighPriorityOk

`func (o *BackupJobModel) GetIsHighPriorityOk() (*bool, bool)`

GetIsHighPriorityOk returns a tuple with the IsHighPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHighPriority

`func (o *BackupJobModel) SetIsHighPriority(v bool)`

SetIsHighPriority sets IsHighPriority field to given value.


### GetVirtualMachines

`func (o *BackupJobModel) GetVirtualMachines() BackupJobVirtualMachinesModel`

GetVirtualMachines returns the VirtualMachines field if non-nil, zero value otherwise.

### GetVirtualMachinesOk

`func (o *BackupJobModel) GetVirtualMachinesOk() (*BackupJobVirtualMachinesModel, bool)`

GetVirtualMachinesOk returns a tuple with the VirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachines

`func (o *BackupJobModel) SetVirtualMachines(v BackupJobVirtualMachinesModel)`

SetVirtualMachines sets VirtualMachines field to given value.


### GetStorage

`func (o *BackupJobModel) GetStorage() BackupJobStorageModel`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *BackupJobModel) GetStorageOk() (*BackupJobStorageModel, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *BackupJobModel) SetStorage(v BackupJobStorageModel)`

SetStorage sets Storage field to given value.


### GetGuestProcessing

`func (o *BackupJobModel) GetGuestProcessing() BackupJobGuestProcessingModel`

GetGuestProcessing returns the GuestProcessing field if non-nil, zero value otherwise.

### GetGuestProcessingOk

`func (o *BackupJobModel) GetGuestProcessingOk() (*BackupJobGuestProcessingModel, bool)`

GetGuestProcessingOk returns a tuple with the GuestProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestProcessing

`func (o *BackupJobModel) SetGuestProcessing(v BackupJobGuestProcessingModel)`

SetGuestProcessing sets GuestProcessing field to given value.


### GetSchedule

`func (o *BackupJobModel) GetSchedule() BackupScheduleModel`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *BackupJobModel) GetScheduleOk() (*BackupScheduleModel, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *BackupJobModel) SetSchedule(v BackupScheduleModel)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


