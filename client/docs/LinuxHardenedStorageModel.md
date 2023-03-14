# LinuxHardenedStorageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **string** | ID of the server that is used as a backup repository. | 
**Repository** | [**LinuxHardenedRepositorySettingsModel**](LinuxHardenedRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 

## Methods

### NewLinuxHardenedStorageModel

`func NewLinuxHardenedStorageModel(hostId string, repository LinuxHardenedRepositorySettingsModel, mountServer MountServerSettingsModel, ) *LinuxHardenedStorageModel`

NewLinuxHardenedStorageModel instantiates a new LinuxHardenedStorageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxHardenedStorageModelWithDefaults

`func NewLinuxHardenedStorageModelWithDefaults() *LinuxHardenedStorageModel`

NewLinuxHardenedStorageModelWithDefaults instantiates a new LinuxHardenedStorageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *LinuxHardenedStorageModel) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *LinuxHardenedStorageModel) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *LinuxHardenedStorageModel) SetHostId(v string)`

SetHostId sets HostId field to given value.


### GetRepository

`func (o *LinuxHardenedStorageModel) GetRepository() LinuxHardenedRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *LinuxHardenedStorageModel) GetRepositoryOk() (*LinuxHardenedRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *LinuxHardenedStorageModel) SetRepository(v LinuxHardenedRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *LinuxHardenedStorageModel) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *LinuxHardenedStorageModel) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *LinuxHardenedStorageModel) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


