# GuestOsCredentialsPerMachineModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowsCredsId** | Pointer to **string** | Credentials ID for a Microsoft Windows VM. | [optional] 
**LinuxCredsId** | Pointer to **string** | Credentials ID for a Linux VM. | [optional] 
**VmObject** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 

## Methods

### NewGuestOsCredentialsPerMachineModel

`func NewGuestOsCredentialsPerMachineModel(vmObject VmwareObjectModel, ) *GuestOsCredentialsPerMachineModel`

NewGuestOsCredentialsPerMachineModel instantiates a new GuestOsCredentialsPerMachineModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestOsCredentialsPerMachineModelWithDefaults

`func NewGuestOsCredentialsPerMachineModelWithDefaults() *GuestOsCredentialsPerMachineModel`

NewGuestOsCredentialsPerMachineModelWithDefaults instantiates a new GuestOsCredentialsPerMachineModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindowsCredsId

`func (o *GuestOsCredentialsPerMachineModel) GetWindowsCredsId() string`

GetWindowsCredsId returns the WindowsCredsId field if non-nil, zero value otherwise.

### GetWindowsCredsIdOk

`func (o *GuestOsCredentialsPerMachineModel) GetWindowsCredsIdOk() (*string, bool)`

GetWindowsCredsIdOk returns a tuple with the WindowsCredsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsCredsId

`func (o *GuestOsCredentialsPerMachineModel) SetWindowsCredsId(v string)`

SetWindowsCredsId sets WindowsCredsId field to given value.

### HasWindowsCredsId

`func (o *GuestOsCredentialsPerMachineModel) HasWindowsCredsId() bool`

HasWindowsCredsId returns a boolean if a field has been set.

### GetLinuxCredsId

`func (o *GuestOsCredentialsPerMachineModel) GetLinuxCredsId() string`

GetLinuxCredsId returns the LinuxCredsId field if non-nil, zero value otherwise.

### GetLinuxCredsIdOk

`func (o *GuestOsCredentialsPerMachineModel) GetLinuxCredsIdOk() (*string, bool)`

GetLinuxCredsIdOk returns a tuple with the LinuxCredsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxCredsId

`func (o *GuestOsCredentialsPerMachineModel) SetLinuxCredsId(v string)`

SetLinuxCredsId sets LinuxCredsId field to given value.

### HasLinuxCredsId

`func (o *GuestOsCredentialsPerMachineModel) HasLinuxCredsId() bool`

HasLinuxCredsId returns a boolean if a field has been set.

### GetVmObject

`func (o *GuestOsCredentialsPerMachineModel) GetVmObject() VmwareObjectModel`

GetVmObject returns the VmObject field if non-nil, zero value otherwise.

### GetVmObjectOk

`func (o *GuestOsCredentialsPerMachineModel) GetVmObjectOk() (*VmwareObjectModel, bool)`

GetVmObjectOk returns a tuple with the VmObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmObject

`func (o *GuestOsCredentialsPerMachineModel) SetVmObject(v VmwareObjectModel)`

SetVmObject sets VmObject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


