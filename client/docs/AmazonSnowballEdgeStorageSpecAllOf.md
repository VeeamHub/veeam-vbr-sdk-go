# AmazonSnowballEdgeStorageSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**AmazonSnowballEdgeStorageAccountModel**](AmazonSnowballEdgeStorageAccountModel.md) |  | 
**Bucket** | [**AmazonSnowballEdgeStorageBucketModel**](AmazonSnowballEdgeStorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 

## Methods

### NewAmazonSnowballEdgeStorageSpecAllOf

`func NewAmazonSnowballEdgeStorageSpecAllOf(account AmazonSnowballEdgeStorageAccountModel, bucket AmazonSnowballEdgeStorageBucketModel, mountServer MountServerSettingsModel, ) *AmazonSnowballEdgeStorageSpecAllOf`

NewAmazonSnowballEdgeStorageSpecAllOf instantiates a new AmazonSnowballEdgeStorageSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSnowballEdgeStorageSpecAllOfWithDefaults

`func NewAmazonSnowballEdgeStorageSpecAllOfWithDefaults() *AmazonSnowballEdgeStorageSpecAllOf`

NewAmazonSnowballEdgeStorageSpecAllOfWithDefaults instantiates a new AmazonSnowballEdgeStorageSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *AmazonSnowballEdgeStorageSpecAllOf) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *AmazonSnowballEdgeStorageSpecAllOf) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *AmazonSnowballEdgeStorageSpecAllOf) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *AmazonSnowballEdgeStorageSpecAllOf) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *AmazonSnowballEdgeStorageSpecAllOf) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *AmazonSnowballEdgeStorageSpecAllOf) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *AmazonSnowballEdgeStorageSpecAllOf) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *AmazonSnowballEdgeStorageSpecAllOf) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *AmazonSnowballEdgeStorageSpecAllOf) GetAccount() AmazonSnowballEdgeStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AmazonSnowballEdgeStorageSpecAllOf) GetAccountOk() (*AmazonSnowballEdgeStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AmazonSnowballEdgeStorageSpecAllOf) SetAccount(v AmazonSnowballEdgeStorageAccountModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *AmazonSnowballEdgeStorageSpecAllOf) GetBucket() AmazonSnowballEdgeStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AmazonSnowballEdgeStorageSpecAllOf) GetBucketOk() (*AmazonSnowballEdgeStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AmazonSnowballEdgeStorageSpecAllOf) SetBucket(v AmazonSnowballEdgeStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *AmazonSnowballEdgeStorageSpecAllOf) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *AmazonSnowballEdgeStorageSpecAllOf) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *AmazonSnowballEdgeStorageSpecAllOf) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


