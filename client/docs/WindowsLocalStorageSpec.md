# WindowsLocalStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **string** | ID of the server that is used as a backup repository. | 
**Repository** | [**WindowsLocalRepositorySettingsModel**](WindowsLocalRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 

## Methods

### NewWindowsLocalStorageSpec

`func NewWindowsLocalStorageSpec(hostId string, repository WindowsLocalRepositorySettingsModel, mountServer MountServerSettingsModel, ) *WindowsLocalStorageSpec`

NewWindowsLocalStorageSpec instantiates a new WindowsLocalStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsLocalStorageSpecWithDefaults

`func NewWindowsLocalStorageSpecWithDefaults() *WindowsLocalStorageSpec`

NewWindowsLocalStorageSpecWithDefaults instantiates a new WindowsLocalStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *WindowsLocalStorageSpec) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *WindowsLocalStorageSpec) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *WindowsLocalStorageSpec) SetHostId(v string)`

SetHostId sets HostId field to given value.


### GetRepository

`func (o *WindowsLocalStorageSpec) GetRepository() WindowsLocalRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *WindowsLocalStorageSpec) GetRepositoryOk() (*WindowsLocalRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *WindowsLocalStorageSpec) SetRepository(v WindowsLocalRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *WindowsLocalStorageSpec) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *WindowsLocalStorageSpec) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *WindowsLocalStorageSpec) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


