# RepositorySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the backup repository. | 
**Description** | **string** | Description of the backup repository. | 
**Tag** | Pointer to **string** | Tag that identifies the backup repository. | [optional] 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**HostId** | **string** | ID of the server that is used as a backup repository. | 
**Repository** | [**LinuxHardenedRepositorySettingsModel**](LinuxHardenedRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 
**Share** | [**SmbRepositoryShareSettingsModel**](SmbRepositoryShareSettingsModel.md) |  | 
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**WasabiCloudStorageAccountModel**](WasabiCloudStorageAccountModel.md) |  | 
**Container** | [**AzureArchiveStorageContainerModel**](AzureArchiveStorageContainerModel.md) |  | 
**ProxyAppliance** | [**S3CompatibleProxyModel**](S3CompatibleProxyModel.md) |  | 
**Bucket** | [**WasabiCloudStorageBucketModel**](WasabiCloudStorageBucketModel.md) |  | 

## Methods

### NewRepositorySpec

`func NewRepositorySpec(name string, description string, type_ ERepositoryType, hostId string, repository LinuxHardenedRepositorySettingsModel, mountServer MountServerSettingsModel, share SmbRepositoryShareSettingsModel, account WasabiCloudStorageAccountModel, container AzureArchiveStorageContainerModel, proxyAppliance S3CompatibleProxyModel, bucket WasabiCloudStorageBucketModel, ) *RepositorySpec`

NewRepositorySpec instantiates a new RepositorySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositorySpecWithDefaults

`func NewRepositorySpecWithDefaults() *RepositorySpec`

NewRepositorySpecWithDefaults instantiates a new RepositorySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RepositorySpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositorySpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositorySpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RepositorySpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepositorySpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepositorySpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *RepositorySpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RepositorySpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RepositorySpec) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RepositorySpec) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *RepositorySpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositorySpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositorySpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetHostId

`func (o *RepositorySpec) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *RepositorySpec) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *RepositorySpec) SetHostId(v string)`

SetHostId sets HostId field to given value.


### GetRepository

`func (o *RepositorySpec) GetRepository() LinuxHardenedRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *RepositorySpec) GetRepositoryOk() (*LinuxHardenedRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *RepositorySpec) SetRepository(v LinuxHardenedRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *RepositorySpec) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *RepositorySpec) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *RepositorySpec) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.


### GetShare

`func (o *RepositorySpec) GetShare() SmbRepositoryShareSettingsModel`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *RepositorySpec) GetShareOk() (*SmbRepositoryShareSettingsModel, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *RepositorySpec) SetShare(v SmbRepositoryShareSettingsModel)`

SetShare sets Share field to given value.


### GetEnableTaskLimit

`func (o *RepositorySpec) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *RepositorySpec) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *RepositorySpec) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *RepositorySpec) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *RepositorySpec) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *RepositorySpec) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *RepositorySpec) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *RepositorySpec) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *RepositorySpec) GetAccount() WasabiCloudStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RepositorySpec) GetAccountOk() (*WasabiCloudStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RepositorySpec) SetAccount(v WasabiCloudStorageAccountModel)`

SetAccount sets Account field to given value.


### GetContainer

`func (o *RepositorySpec) GetContainer() AzureArchiveStorageContainerModel`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *RepositorySpec) GetContainerOk() (*AzureArchiveStorageContainerModel, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *RepositorySpec) SetContainer(v AzureArchiveStorageContainerModel)`

SetContainer sets Container field to given value.


### GetProxyAppliance

`func (o *RepositorySpec) GetProxyAppliance() S3CompatibleProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *RepositorySpec) GetProxyApplianceOk() (*S3CompatibleProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *RepositorySpec) SetProxyAppliance(v S3CompatibleProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.


### GetBucket

`func (o *RepositorySpec) GetBucket() WasabiCloudStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *RepositorySpec) GetBucketOk() (*WasabiCloudStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *RepositorySpec) SetBucket(v WasabiCloudStorageBucketModel)`

SetBucket sets Bucket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


