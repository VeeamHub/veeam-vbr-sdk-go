# NfsStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the backup repository. | 
**Description** | **string** | Description of the backup repository. | 
**Tag** | **string** | VMware vSphere tag assigned to the backup repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**Share** | [**NfsRepositoryShareSettingsSpec**](NfsRepositoryShareSettingsSpec.md) |  | 
**Repository** | [**NetworkRepositorySettingsModel**](NetworkRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsImportSpec**](MountServerSettingsImportSpec.md) |  | 

## Methods

### NewNfsStorageImportSpec

`func NewNfsStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, share NfsRepositoryShareSettingsSpec, repository NetworkRepositorySettingsModel, mountServer MountServerSettingsImportSpec, ) *NfsStorageImportSpec`

NewNfsStorageImportSpec instantiates a new NfsStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsStorageImportSpecWithDefaults

`func NewNfsStorageImportSpecWithDefaults() *NfsStorageImportSpec`

NewNfsStorageImportSpecWithDefaults instantiates a new NfsStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NfsStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NfsStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NfsStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NfsStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NfsStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NfsStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *NfsStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *NfsStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *NfsStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *NfsStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NfsStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NfsStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetShare

`func (o *NfsStorageImportSpec) GetShare() NfsRepositoryShareSettingsSpec`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *NfsStorageImportSpec) GetShareOk() (*NfsRepositoryShareSettingsSpec, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *NfsStorageImportSpec) SetShare(v NfsRepositoryShareSettingsSpec)`

SetShare sets Share field to given value.


### GetRepository

`func (o *NfsStorageImportSpec) GetRepository() NetworkRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *NfsStorageImportSpec) GetRepositoryOk() (*NetworkRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *NfsStorageImportSpec) SetRepository(v NetworkRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *NfsStorageImportSpec) GetMountServer() MountServerSettingsImportSpec`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *NfsStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *NfsStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


