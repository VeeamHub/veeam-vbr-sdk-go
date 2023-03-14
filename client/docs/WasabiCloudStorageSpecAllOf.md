# WasabiCloudStorageSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**WasabiCloudStorageAccountModel**](WasabiCloudStorageAccountModel.md) |  | 
**Bucket** | [**WasabiCloudStorageBucketModel**](WasabiCloudStorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 
**ProxyAppliance** | Pointer to [**S3CompatibleProxyModel**](S3CompatibleProxyModel.md) |  | [optional] 

## Methods

### NewWasabiCloudStorageSpecAllOf

`func NewWasabiCloudStorageSpecAllOf(account WasabiCloudStorageAccountModel, bucket WasabiCloudStorageBucketModel, mountServer MountServerSettingsModel, ) *WasabiCloudStorageSpecAllOf`

NewWasabiCloudStorageSpecAllOf instantiates a new WasabiCloudStorageSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWasabiCloudStorageSpecAllOfWithDefaults

`func NewWasabiCloudStorageSpecAllOfWithDefaults() *WasabiCloudStorageSpecAllOf`

NewWasabiCloudStorageSpecAllOfWithDefaults instantiates a new WasabiCloudStorageSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *WasabiCloudStorageSpecAllOf) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *WasabiCloudStorageSpecAllOf) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *WasabiCloudStorageSpecAllOf) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *WasabiCloudStorageSpecAllOf) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *WasabiCloudStorageSpecAllOf) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *WasabiCloudStorageSpecAllOf) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *WasabiCloudStorageSpecAllOf) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *WasabiCloudStorageSpecAllOf) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *WasabiCloudStorageSpecAllOf) GetAccount() WasabiCloudStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WasabiCloudStorageSpecAllOf) GetAccountOk() (*WasabiCloudStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WasabiCloudStorageSpecAllOf) SetAccount(v WasabiCloudStorageAccountModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *WasabiCloudStorageSpecAllOf) GetBucket() WasabiCloudStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *WasabiCloudStorageSpecAllOf) GetBucketOk() (*WasabiCloudStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *WasabiCloudStorageSpecAllOf) SetBucket(v WasabiCloudStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *WasabiCloudStorageSpecAllOf) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *WasabiCloudStorageSpecAllOf) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *WasabiCloudStorageSpecAllOf) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.


### GetProxyAppliance

`func (o *WasabiCloudStorageSpecAllOf) GetProxyAppliance() S3CompatibleProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *WasabiCloudStorageSpecAllOf) GetProxyApplianceOk() (*S3CompatibleProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *WasabiCloudStorageSpecAllOf) SetProxyAppliance(v S3CompatibleProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *WasabiCloudStorageSpecAllOf) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


