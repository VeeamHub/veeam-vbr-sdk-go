# S3CompatibleStorageSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**S3CompatibleStorageAccountModel**](S3CompatibleStorageAccountModel.md) |  | 
**Bucket** | [**S3CompatibleStorageBucketModel**](S3CompatibleStorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 
**ProxyAppliance** | Pointer to [**S3CompatibleProxyModel**](S3CompatibleProxyModel.md) |  | [optional] 

## Methods

### NewS3CompatibleStorageSpecAllOf

`func NewS3CompatibleStorageSpecAllOf(account S3CompatibleStorageAccountModel, bucket S3CompatibleStorageBucketModel, mountServer MountServerSettingsModel, ) *S3CompatibleStorageSpecAllOf`

NewS3CompatibleStorageSpecAllOf instantiates a new S3CompatibleStorageSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CompatibleStorageSpecAllOfWithDefaults

`func NewS3CompatibleStorageSpecAllOfWithDefaults() *S3CompatibleStorageSpecAllOf`

NewS3CompatibleStorageSpecAllOfWithDefaults instantiates a new S3CompatibleStorageSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *S3CompatibleStorageSpecAllOf) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *S3CompatibleStorageSpecAllOf) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *S3CompatibleStorageSpecAllOf) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *S3CompatibleStorageSpecAllOf) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *S3CompatibleStorageSpecAllOf) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *S3CompatibleStorageSpecAllOf) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *S3CompatibleStorageSpecAllOf) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *S3CompatibleStorageSpecAllOf) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *S3CompatibleStorageSpecAllOf) GetAccount() S3CompatibleStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *S3CompatibleStorageSpecAllOf) GetAccountOk() (*S3CompatibleStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *S3CompatibleStorageSpecAllOf) SetAccount(v S3CompatibleStorageAccountModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *S3CompatibleStorageSpecAllOf) GetBucket() S3CompatibleStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *S3CompatibleStorageSpecAllOf) GetBucketOk() (*S3CompatibleStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *S3CompatibleStorageSpecAllOf) SetBucket(v S3CompatibleStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *S3CompatibleStorageSpecAllOf) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *S3CompatibleStorageSpecAllOf) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *S3CompatibleStorageSpecAllOf) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.


### GetProxyAppliance

`func (o *S3CompatibleStorageSpecAllOf) GetProxyAppliance() S3CompatibleProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *S3CompatibleStorageSpecAllOf) GetProxyApplianceOk() (*S3CompatibleProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *S3CompatibleStorageSpecAllOf) SetProxyAppliance(v S3CompatibleProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *S3CompatibleStorageSpecAllOf) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


