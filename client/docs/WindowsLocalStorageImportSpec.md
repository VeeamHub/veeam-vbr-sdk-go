# WindowsLocalStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the backup repository. | 
**Description** | **string** | Description of the backup repository. | 
**Tag** | **string** | VMware vSphere tag assigned to the backup repository. | 
**HostName** | **string** | ID of the server that is used as a backup repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**Repository** | [**WindowsLocalRepositorySettingsModel**](WindowsLocalRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsImportSpec**](MountServerSettingsImportSpec.md) |  | 

## Methods

### NewWindowsLocalStorageImportSpec

`func NewWindowsLocalStorageImportSpec(name string, description string, tag string, hostName string, type_ ERepositoryType, repository WindowsLocalRepositorySettingsModel, mountServer MountServerSettingsImportSpec, ) *WindowsLocalStorageImportSpec`

NewWindowsLocalStorageImportSpec instantiates a new WindowsLocalStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsLocalStorageImportSpecWithDefaults

`func NewWindowsLocalStorageImportSpecWithDefaults() *WindowsLocalStorageImportSpec`

NewWindowsLocalStorageImportSpecWithDefaults instantiates a new WindowsLocalStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WindowsLocalStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WindowsLocalStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WindowsLocalStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WindowsLocalStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WindowsLocalStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WindowsLocalStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *WindowsLocalStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *WindowsLocalStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *WindowsLocalStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetHostName

`func (o *WindowsLocalStorageImportSpec) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *WindowsLocalStorageImportSpec) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *WindowsLocalStorageImportSpec) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetType

`func (o *WindowsLocalStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WindowsLocalStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WindowsLocalStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetRepository

`func (o *WindowsLocalStorageImportSpec) GetRepository() WindowsLocalRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *WindowsLocalStorageImportSpec) GetRepositoryOk() (*WindowsLocalRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *WindowsLocalStorageImportSpec) SetRepository(v WindowsLocalRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *WindowsLocalStorageImportSpec) GetMountServer() MountServerSettingsImportSpec`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *WindowsLocalStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *WindowsLocalStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


