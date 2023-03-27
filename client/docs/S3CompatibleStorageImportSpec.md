# S3CompatibleStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the object storage repository. | 
**Description** | **string** | Description of the object storage repository. | 
**Tag** | **string** | Tag that identifies the object storage repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**S3CompatibleStorageAccountImportModel**](S3CompatibleStorageAccountImportModel.md) |  | 
**Bucket** | [**S3CompatibleStorageBucketModel**](S3CompatibleStorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsImportSpec**](MountServerSettingsImportSpec.md) |  | 
**ProxyAppliance** | Pointer to [**S3CompatibleProxyModel**](S3CompatibleProxyModel.md) |  | [optional] 

## Methods

### NewS3CompatibleStorageImportSpec

`func NewS3CompatibleStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account S3CompatibleStorageAccountImportModel, bucket S3CompatibleStorageBucketModel, mountServer MountServerSettingsImportSpec, ) *S3CompatibleStorageImportSpec`

NewS3CompatibleStorageImportSpec instantiates a new S3CompatibleStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CompatibleStorageImportSpecWithDefaults

`func NewS3CompatibleStorageImportSpecWithDefaults() *S3CompatibleStorageImportSpec`

NewS3CompatibleStorageImportSpecWithDefaults instantiates a new S3CompatibleStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *S3CompatibleStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3CompatibleStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3CompatibleStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *S3CompatibleStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *S3CompatibleStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *S3CompatibleStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *S3CompatibleStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *S3CompatibleStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *S3CompatibleStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *S3CompatibleStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *S3CompatibleStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *S3CompatibleStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetEnableTaskLimit

`func (o *S3CompatibleStorageImportSpec) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *S3CompatibleStorageImportSpec) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *S3CompatibleStorageImportSpec) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *S3CompatibleStorageImportSpec) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *S3CompatibleStorageImportSpec) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *S3CompatibleStorageImportSpec) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *S3CompatibleStorageImportSpec) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *S3CompatibleStorageImportSpec) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *S3CompatibleStorageImportSpec) GetAccount() S3CompatibleStorageAccountImportModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *S3CompatibleStorageImportSpec) GetAccountOk() (*S3CompatibleStorageAccountImportModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *S3CompatibleStorageImportSpec) SetAccount(v S3CompatibleStorageAccountImportModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *S3CompatibleStorageImportSpec) GetBucket() S3CompatibleStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *S3CompatibleStorageImportSpec) GetBucketOk() (*S3CompatibleStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *S3CompatibleStorageImportSpec) SetBucket(v S3CompatibleStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *S3CompatibleStorageImportSpec) GetMountServer() MountServerSettingsImportSpec`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *S3CompatibleStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *S3CompatibleStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec)`

SetMountServer sets MountServer field to given value.


### GetProxyAppliance

`func (o *S3CompatibleStorageImportSpec) GetProxyAppliance() S3CompatibleProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *S3CompatibleStorageImportSpec) GetProxyApplianceOk() (*S3CompatibleProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *S3CompatibleStorageImportSpec) SetProxyAppliance(v S3CompatibleProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *S3CompatibleStorageImportSpec) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


