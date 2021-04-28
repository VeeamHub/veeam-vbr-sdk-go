# GuestOsCredentialsPerMachineImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowsCreds** | Pointer to [**CredentialsImportModel**](CredentialsImportModel.md) |  | [optional] 
**LinuxCreds** | Pointer to [**CredentialsImportModel**](CredentialsImportModel.md) |  | [optional] 
**VmObject** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 

## Methods

### NewGuestOsCredentialsPerMachineImportModel

`func NewGuestOsCredentialsPerMachineImportModel(vmObject VmwareObjectModel, ) *GuestOsCredentialsPerMachineImportModel`

NewGuestOsCredentialsPerMachineImportModel instantiates a new GuestOsCredentialsPerMachineImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestOsCredentialsPerMachineImportModelWithDefaults

`func NewGuestOsCredentialsPerMachineImportModelWithDefaults() *GuestOsCredentialsPerMachineImportModel`

NewGuestOsCredentialsPerMachineImportModelWithDefaults instantiates a new GuestOsCredentialsPerMachineImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindowsCreds

`func (o *GuestOsCredentialsPerMachineImportModel) GetWindowsCreds() CredentialsImportModel`

GetWindowsCreds returns the WindowsCreds field if non-nil, zero value otherwise.

### GetWindowsCredsOk

`func (o *GuestOsCredentialsPerMachineImportModel) GetWindowsCredsOk() (*CredentialsImportModel, bool)`

GetWindowsCredsOk returns a tuple with the WindowsCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsCreds

`func (o *GuestOsCredentialsPerMachineImportModel) SetWindowsCreds(v CredentialsImportModel)`

SetWindowsCreds sets WindowsCreds field to given value.

### HasWindowsCreds

`func (o *GuestOsCredentialsPerMachineImportModel) HasWindowsCreds() bool`

HasWindowsCreds returns a boolean if a field has been set.

### GetLinuxCreds

`func (o *GuestOsCredentialsPerMachineImportModel) GetLinuxCreds() CredentialsImportModel`

GetLinuxCreds returns the LinuxCreds field if non-nil, zero value otherwise.

### GetLinuxCredsOk

`func (o *GuestOsCredentialsPerMachineImportModel) GetLinuxCredsOk() (*CredentialsImportModel, bool)`

GetLinuxCredsOk returns a tuple with the LinuxCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxCreds

`func (o *GuestOsCredentialsPerMachineImportModel) SetLinuxCreds(v CredentialsImportModel)`

SetLinuxCreds sets LinuxCreds field to given value.

### HasLinuxCreds

`func (o *GuestOsCredentialsPerMachineImportModel) HasLinuxCreds() bool`

HasLinuxCreds returns a boolean if a field has been set.

### GetVmObject

`func (o *GuestOsCredentialsPerMachineImportModel) GetVmObject() VmwareObjectModel`

GetVmObject returns the VmObject field if non-nil, zero value otherwise.

### GetVmObjectOk

`func (o *GuestOsCredentialsPerMachineImportModel) GetVmObjectOk() (*VmwareObjectModel, bool)`

GetVmObjectOk returns a tuple with the VmObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmObject

`func (o *GuestOsCredentialsPerMachineImportModel) SetVmObject(v VmwareObjectModel)`

SetVmObject sets VmObject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


