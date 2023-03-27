# LinuxHardenedStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the backup repository. | 
**Description** | **string** | Description of the backup repository. | 
**Tag** | **string** | Tag that identifies the backup repository. | 
**HostName** | **string** | ID of the server that is used as a backup repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**Repository** | [**LinuxHardenedRepositorySettingsModel**](LinuxHardenedRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsImportSpec**](MountServerSettingsImportSpec.md) |  | 

## Methods

### NewLinuxHardenedStorageImportSpec

`func NewLinuxHardenedStorageImportSpec(name string, description string, tag string, hostName string, type_ ERepositoryType, repository LinuxHardenedRepositorySettingsModel, mountServer MountServerSettingsImportSpec, ) *LinuxHardenedStorageImportSpec`

NewLinuxHardenedStorageImportSpec instantiates a new LinuxHardenedStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxHardenedStorageImportSpecWithDefaults

`func NewLinuxHardenedStorageImportSpecWithDefaults() *LinuxHardenedStorageImportSpec`

NewLinuxHardenedStorageImportSpecWithDefaults instantiates a new LinuxHardenedStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LinuxHardenedStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinuxHardenedStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinuxHardenedStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LinuxHardenedStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LinuxHardenedStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LinuxHardenedStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *LinuxHardenedStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *LinuxHardenedStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *LinuxHardenedStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetHostName

`func (o *LinuxHardenedStorageImportSpec) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *LinuxHardenedStorageImportSpec) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *LinuxHardenedStorageImportSpec) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetType

`func (o *LinuxHardenedStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinuxHardenedStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinuxHardenedStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetRepository

`func (o *LinuxHardenedStorageImportSpec) GetRepository() LinuxHardenedRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *LinuxHardenedStorageImportSpec) GetRepositoryOk() (*LinuxHardenedRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *LinuxHardenedStorageImportSpec) SetRepository(v LinuxHardenedRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *LinuxHardenedStorageImportSpec) GetMountServer() MountServerSettingsImportSpec`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *LinuxHardenedStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *LinuxHardenedStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


