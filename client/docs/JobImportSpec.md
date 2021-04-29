# JobImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the job. | 
**Description** | **string** | Description of the job. | 
**IsHighPriority** | **bool** | If *true*, the job has a high priority in getting backup infrastructure resources. | 
**Type** | [**EJobType**](EJobType.md) |  | 
**VirtualMachines** | [**BackupJobVirtualMachinesSpec**](BackupJobVirtualMachinesSpec.md) |  | 
**Storage** | [**BackupJobStorageImportModel**](BackupJobStorageImportModel.md) |  | 
**GuestProcessing** | [**BackupJobGuestProcessingImportModel**](BackupJobGuestProcessingImportModel.md) |  | 
**Schedule** | [**BackupScheduleModel**](BackupScheduleModel.md) |  | 

## Methods

### NewJobImportSpec

`func NewJobImportSpec(name string, description string, isHighPriority bool, type_ EJobType, virtualMachines BackupJobVirtualMachinesSpec, storage BackupJobStorageImportModel, guestProcessing BackupJobGuestProcessingImportModel, schedule BackupScheduleModel, ) *JobImportSpec`

NewJobImportSpec instantiates a new JobImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobImportSpecWithDefaults

`func NewJobImportSpecWithDefaults() *JobImportSpec`

NewJobImportSpecWithDefaults instantiates a new JobImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JobImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *JobImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsHighPriority

`func (o *JobImportSpec) GetIsHighPriority() bool`

GetIsHighPriority returns the IsHighPriority field if non-nil, zero value otherwise.

### GetIsHighPriorityOk

`func (o *JobImportSpec) GetIsHighPriorityOk() (*bool, bool)`

GetIsHighPriorityOk returns a tuple with the IsHighPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHighPriority

`func (o *JobImportSpec) SetIsHighPriority(v bool)`

SetIsHighPriority sets IsHighPriority field to given value.


### GetType

`func (o *JobImportSpec) GetType() EJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobImportSpec) GetTypeOk() (*EJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobImportSpec) SetType(v EJobType)`

SetType sets Type field to given value.


### GetVirtualMachines

`func (o *JobImportSpec) GetVirtualMachines() BackupJobVirtualMachinesSpec`

GetVirtualMachines returns the VirtualMachines field if non-nil, zero value otherwise.

### GetVirtualMachinesOk

`func (o *JobImportSpec) GetVirtualMachinesOk() (*BackupJobVirtualMachinesSpec, bool)`

GetVirtualMachinesOk returns a tuple with the VirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachines

`func (o *JobImportSpec) SetVirtualMachines(v BackupJobVirtualMachinesSpec)`

SetVirtualMachines sets VirtualMachines field to given value.


### GetStorage

`func (o *JobImportSpec) GetStorage() BackupJobStorageImportModel`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *JobImportSpec) GetStorageOk() (*BackupJobStorageImportModel, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *JobImportSpec) SetStorage(v BackupJobStorageImportModel)`

SetStorage sets Storage field to given value.


### GetGuestProcessing

`func (o *JobImportSpec) GetGuestProcessing() BackupJobGuestProcessingImportModel`

GetGuestProcessing returns the GuestProcessing field if non-nil, zero value otherwise.

### GetGuestProcessingOk

`func (o *JobImportSpec) GetGuestProcessingOk() (*BackupJobGuestProcessingImportModel, bool)`

GetGuestProcessingOk returns a tuple with the GuestProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestProcessing

`func (o *JobImportSpec) SetGuestProcessing(v BackupJobGuestProcessingImportModel)`

SetGuestProcessing sets GuestProcessing field to given value.


### GetSchedule

`func (o *JobImportSpec) GetSchedule() BackupScheduleModel`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *JobImportSpec) GetScheduleOk() (*BackupScheduleModel, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *JobImportSpec) SetSchedule(v BackupScheduleModel)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


