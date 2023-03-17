# LinuxLocalStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the backup repository. | 
**Description** | **string** | Description of the backup repository. | 
**Tag** | **string** | Tag that identifies the backup repository. | 
**HostName** | **string** | ID of the server that is used as a backup repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**Repository** | [**LinuxLocalRepositorySettingsModel**](LinuxLocalRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsImportSpec**](MountServerSettingsImportSpec.md) |  | 

## Methods

### NewLinuxLocalStorageImportSpec

`func NewLinuxLocalStorageImportSpec(name string, description string, tag string, hostName string, type_ ERepositoryType, repository LinuxLocalRepositorySettingsModel, mountServer MountServerSettingsImportSpec, ) *LinuxLocalStorageImportSpec`

NewLinuxLocalStorageImportSpec instantiates a new LinuxLocalStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxLocalStorageImportSpecWithDefaults

`func NewLinuxLocalStorageImportSpecWithDefaults() *LinuxLocalStorageImportSpec`

NewLinuxLocalStorageImportSpecWithDefaults instantiates a new LinuxLocalStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LinuxLocalStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinuxLocalStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinuxLocalStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LinuxLocalStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LinuxLocalStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LinuxLocalStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *LinuxLocalStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *LinuxLocalStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *LinuxLocalStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetHostName

`func (o *LinuxLocalStorageImportSpec) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *LinuxLocalStorageImportSpec) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *LinuxLocalStorageImportSpec) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetType

`func (o *LinuxLocalStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinuxLocalStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinuxLocalStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetRepository

`func (o *LinuxLocalStorageImportSpec) GetRepository() LinuxLocalRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *LinuxLocalStorageImportSpec) GetRepositoryOk() (*LinuxLocalRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *LinuxLocalStorageImportSpec) SetRepository(v LinuxLocalRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *LinuxLocalStorageImportSpec) GetMountServer() MountServerSettingsImportSpec`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *LinuxLocalStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *LinuxLocalStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


