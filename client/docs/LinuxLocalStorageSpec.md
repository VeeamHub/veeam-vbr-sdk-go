# LinuxLocalStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **string** | ID of the server that is used as a backup repository. | 
**Repository** | [**LinuxLocalRepositorySettingsModel**](LinuxLocalRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 

## Methods

### NewLinuxLocalStorageSpec

`func NewLinuxLocalStorageSpec(hostId string, repository LinuxLocalRepositorySettingsModel, mountServer MountServerSettingsModel, ) *LinuxLocalStorageSpec`

NewLinuxLocalStorageSpec instantiates a new LinuxLocalStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxLocalStorageSpecWithDefaults

`func NewLinuxLocalStorageSpecWithDefaults() *LinuxLocalStorageSpec`

NewLinuxLocalStorageSpecWithDefaults instantiates a new LinuxLocalStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *LinuxLocalStorageSpec) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *LinuxLocalStorageSpec) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *LinuxLocalStorageSpec) SetHostId(v string)`

SetHostId sets HostId field to given value.


### GetRepository

`func (o *LinuxLocalStorageSpec) GetRepository() LinuxLocalRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *LinuxLocalStorageSpec) GetRepositoryOk() (*LinuxLocalRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *LinuxLocalStorageSpec) SetRepository(v LinuxLocalRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *LinuxLocalStorageSpec) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *LinuxLocalStorageSpec) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *LinuxLocalStorageSpec) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

