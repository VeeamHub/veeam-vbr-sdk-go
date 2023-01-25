# AmazonSnowballEdgeStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**AmazonSnowballEdgeStorageAccountModel**](AmazonSnowballEdgeStorageAccountModel.md) |  | 
**Bucket** | [**AmazonSnowballEdgeStorageBucketModel**](AmazonSnowballEdgeStorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 

## Methods

### NewAmazonSnowballEdgeStorageSpec

`func NewAmazonSnowballEdgeStorageSpec(account AmazonSnowballEdgeStorageAccountModel, bucket AmazonSnowballEdgeStorageBucketModel, mountServer MountServerSettingsModel, ) *AmazonSnowballEdgeStorageSpec`

NewAmazonSnowballEdgeStorageSpec instantiates a new AmazonSnowballEdgeStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSnowballEdgeStorageSpecWithDefaults

`func NewAmazonSnowballEdgeStorageSpecWithDefaults() *AmazonSnowballEdgeStorageSpec`

NewAmazonSnowballEdgeStorageSpecWithDefaults instantiates a new AmazonSnowballEdgeStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *AmazonSnowballEdgeStorageSpec) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *AmazonSnowballEdgeStorageSpec) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *AmazonSnowballEdgeStorageSpec) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *AmazonSnowballEdgeStorageSpec) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *AmazonSnowballEdgeStorageSpec) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *AmazonSnowballEdgeStorageSpec) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *AmazonSnowballEdgeStorageSpec) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *AmazonSnowballEdgeStorageSpec) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *AmazonSnowballEdgeStorageSpec) GetAccount() AmazonSnowballEdgeStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AmazonSnowballEdgeStorageSpec) GetAccountOk() (*AmazonSnowballEdgeStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AmazonSnowballEdgeStorageSpec) SetAccount(v AmazonSnowballEdgeStorageAccountModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *AmazonSnowballEdgeStorageSpec) GetBucket() AmazonSnowballEdgeStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AmazonSnowballEdgeStorageSpec) GetBucketOk() (*AmazonSnowballEdgeStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AmazonSnowballEdgeStorageSpec) SetBucket(v AmazonSnowballEdgeStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *AmazonSnowballEdgeStorageSpec) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *AmazonSnowballEdgeStorageSpec) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *AmazonSnowballEdgeStorageSpec) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


