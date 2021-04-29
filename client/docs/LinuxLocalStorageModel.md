# LinuxLocalStorageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **string** | ID of the server that is used as a backup repository. | 
**Repository** | [**LinuxLocalRepositorySettingsModel**](LinuxLocalRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 

## Methods

### NewLinuxLocalStorageModel

`func NewLinuxLocalStorageModel(hostId string, repository LinuxLocalRepositorySettingsModel, mountServer MountServerSettingsModel, ) *LinuxLocalStorageModel`

NewLinuxLocalStorageModel instantiates a new LinuxLocalStorageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxLocalStorageModelWithDefaults

`func NewLinuxLocalStorageModelWithDefaults() *LinuxLocalStorageModel`

NewLinuxLocalStorageModelWithDefaults instantiates a new LinuxLocalStorageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *LinuxLocalStorageModel) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *LinuxLocalStorageModel) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *LinuxLocalStorageModel) SetHostId(v string)`

SetHostId sets HostId field to given value.


### GetRepository

`func (o *LinuxLocalStorageModel) GetRepository() LinuxLocalRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *LinuxLocalStorageModel) GetRepositoryOk() (*LinuxLocalRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *LinuxLocalStorageModel) SetRepository(v LinuxLocalRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *LinuxLocalStorageModel) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *LinuxLocalStorageModel) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *LinuxLocalStorageModel) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


