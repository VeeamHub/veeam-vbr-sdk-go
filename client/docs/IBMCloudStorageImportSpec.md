# IBMCloudStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the object storage repository. | 
**Description** | **string** | Description of the object storage repository. | 
**Tag** | **string** | Tag that identifies the object storage repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**IBMCloudStorageAccountImportModel**](IBMCloudStorageAccountImportModel.md) |  | 
**Bucket** | [**IBMCloudStorageBucketModel**](IBMCloudStorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsImportSpec**](MountServerSettingsImportSpec.md) |  | 
**ProxyAppliance** | Pointer to [**S3CompatibleProxyModel**](S3CompatibleProxyModel.md) |  | [optional] 

## Methods

### NewIBMCloudStorageImportSpec

`func NewIBMCloudStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account IBMCloudStorageAccountImportModel, bucket IBMCloudStorageBucketModel, mountServer MountServerSettingsImportSpec, ) *IBMCloudStorageImportSpec`

NewIBMCloudStorageImportSpec instantiates a new IBMCloudStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIBMCloudStorageImportSpecWithDefaults

`func NewIBMCloudStorageImportSpecWithDefaults() *IBMCloudStorageImportSpec`

NewIBMCloudStorageImportSpecWithDefaults instantiates a new IBMCloudStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IBMCloudStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IBMCloudStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IBMCloudStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IBMCloudStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IBMCloudStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IBMCloudStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *IBMCloudStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IBMCloudStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IBMCloudStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *IBMCloudStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IBMCloudStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IBMCloudStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetEnableTaskLimit

`func (o *IBMCloudStorageImportSpec) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *IBMCloudStorageImportSpec) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *IBMCloudStorageImportSpec) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *IBMCloudStorageImportSpec) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *IBMCloudStorageImportSpec) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *IBMCloudStorageImportSpec) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *IBMCloudStorageImportSpec) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *IBMCloudStorageImportSpec) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *IBMCloudStorageImportSpec) GetAccount() IBMCloudStorageAccountImportModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IBMCloudStorageImportSpec) GetAccountOk() (*IBMCloudStorageAccountImportModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IBMCloudStorageImportSpec) SetAccount(v IBMCloudStorageAccountImportModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *IBMCloudStorageImportSpec) GetBucket() IBMCloudStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IBMCloudStorageImportSpec) GetBucketOk() (*IBMCloudStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IBMCloudStorageImportSpec) SetBucket(v IBMCloudStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *IBMCloudStorageImportSpec) GetMountServer() MountServerSettingsImportSpec`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *IBMCloudStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *IBMCloudStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec)`

SetMountServer sets MountServer field to given value.


### GetProxyAppliance

`func (o *IBMCloudStorageImportSpec) GetProxyAppliance() S3CompatibleProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *IBMCloudStorageImportSpec) GetProxyApplianceOk() (*S3CompatibleProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *IBMCloudStorageImportSpec) SetProxyAppliance(v S3CompatibleProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *IBMCloudStorageImportSpec) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


