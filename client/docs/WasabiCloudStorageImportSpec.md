# WasabiCloudStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the object storage repository. | 
**Description** | **string** | Description of the object storage repository. | 
**Tag** | **string** | Tag that identifies the object storage repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**WasabiCloudStorageAccountImportModel**](WasabiCloudStorageAccountImportModel.md) |  | 
**Bucket** | [**WasabiCloudStorageBucketModel**](WasabiCloudStorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsImportSpec**](MountServerSettingsImportSpec.md) |  | 
**ProxyAppliance** | Pointer to [**S3CompatibleProxyModel**](S3CompatibleProxyModel.md) |  | [optional] 

## Methods

### NewWasabiCloudStorageImportSpec

`func NewWasabiCloudStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account WasabiCloudStorageAccountImportModel, bucket WasabiCloudStorageBucketModel, mountServer MountServerSettingsImportSpec, ) *WasabiCloudStorageImportSpec`

NewWasabiCloudStorageImportSpec instantiates a new WasabiCloudStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWasabiCloudStorageImportSpecWithDefaults

`func NewWasabiCloudStorageImportSpecWithDefaults() *WasabiCloudStorageImportSpec`

NewWasabiCloudStorageImportSpecWithDefaults instantiates a new WasabiCloudStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WasabiCloudStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WasabiCloudStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WasabiCloudStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WasabiCloudStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WasabiCloudStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WasabiCloudStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *WasabiCloudStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *WasabiCloudStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *WasabiCloudStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *WasabiCloudStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WasabiCloudStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WasabiCloudStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetEnableTaskLimit

`func (o *WasabiCloudStorageImportSpec) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *WasabiCloudStorageImportSpec) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *WasabiCloudStorageImportSpec) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *WasabiCloudStorageImportSpec) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *WasabiCloudStorageImportSpec) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *WasabiCloudStorageImportSpec) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *WasabiCloudStorageImportSpec) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *WasabiCloudStorageImportSpec) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *WasabiCloudStorageImportSpec) GetAccount() WasabiCloudStorageAccountImportModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WasabiCloudStorageImportSpec) GetAccountOk() (*WasabiCloudStorageAccountImportModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WasabiCloudStorageImportSpec) SetAccount(v WasabiCloudStorageAccountImportModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *WasabiCloudStorageImportSpec) GetBucket() WasabiCloudStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *WasabiCloudStorageImportSpec) GetBucketOk() (*WasabiCloudStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *WasabiCloudStorageImportSpec) SetBucket(v WasabiCloudStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *WasabiCloudStorageImportSpec) GetMountServer() MountServerSettingsImportSpec`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *WasabiCloudStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *WasabiCloudStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec)`

SetMountServer sets MountServer field to given value.


### GetProxyAppliance

`func (o *WasabiCloudStorageImportSpec) GetProxyAppliance() S3CompatibleProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *WasabiCloudStorageImportSpec) GetProxyApplianceOk() (*S3CompatibleProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *WasabiCloudStorageImportSpec) SetProxyAppliance(v S3CompatibleProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *WasabiCloudStorageImportSpec) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


