# SmbStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the backup repository. | 
**Description** | **string** | Description of the backup repository. | 
**Tag** | **string** | VMware vSphere tag assigned to the backup repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**Share** | [**SmbRepositoryShareSettingsSpec**](SmbRepositoryShareSettingsSpec.md) |  | 
**Repository** | [**NetworkRepositorySettingsModel**](NetworkRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsImportSpec**](MountServerSettingsImportSpec.md) |  | 

## Methods

### NewSmbStorageImportSpec

`func NewSmbStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, share SmbRepositoryShareSettingsSpec, repository NetworkRepositorySettingsModel, mountServer MountServerSettingsImportSpec, ) *SmbStorageImportSpec`

NewSmbStorageImportSpec instantiates a new SmbStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbStorageImportSpecWithDefaults

`func NewSmbStorageImportSpecWithDefaults() *SmbStorageImportSpec`

NewSmbStorageImportSpecWithDefaults instantiates a new SmbStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SmbStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmbStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmbStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SmbStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SmbStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SmbStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *SmbStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SmbStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SmbStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *SmbStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SmbStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SmbStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetShare

`func (o *SmbStorageImportSpec) GetShare() SmbRepositoryShareSettingsSpec`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *SmbStorageImportSpec) GetShareOk() (*SmbRepositoryShareSettingsSpec, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *SmbStorageImportSpec) SetShare(v SmbRepositoryShareSettingsSpec)`

SetShare sets Share field to given value.


### GetRepository

`func (o *SmbStorageImportSpec) GetRepository() NetworkRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *SmbStorageImportSpec) GetRepositoryOk() (*NetworkRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *SmbStorageImportSpec) SetRepository(v NetworkRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *SmbStorageImportSpec) GetMountServer() MountServerSettingsImportSpec`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *SmbStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *SmbStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


