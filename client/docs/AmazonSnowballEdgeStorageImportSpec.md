# AmazonSnowballEdgeStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the object storage repository. | 
**Description** | **string** | Description of the object storage repository. | 
**Tag** | **string** | Tag that identifies the object storage repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**AmazonSnowballEdgeStorageAccountImportModel**](AmazonSnowballEdgeStorageAccountImportModel.md) |  | 
**Bucket** | [**AmazonSnowballEdgeStorageBucketModel**](AmazonSnowballEdgeStorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsImportSpec**](MountServerSettingsImportSpec.md) |  | 

## Methods

### NewAmazonSnowballEdgeStorageImportSpec

`func NewAmazonSnowballEdgeStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account AmazonSnowballEdgeStorageAccountImportModel, bucket AmazonSnowballEdgeStorageBucketModel, mountServer MountServerSettingsImportSpec, ) *AmazonSnowballEdgeStorageImportSpec`

NewAmazonSnowballEdgeStorageImportSpec instantiates a new AmazonSnowballEdgeStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSnowballEdgeStorageImportSpecWithDefaults

`func NewAmazonSnowballEdgeStorageImportSpecWithDefaults() *AmazonSnowballEdgeStorageImportSpec`

NewAmazonSnowballEdgeStorageImportSpecWithDefaults instantiates a new AmazonSnowballEdgeStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AmazonSnowballEdgeStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmazonSnowballEdgeStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmazonSnowballEdgeStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AmazonSnowballEdgeStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonSnowballEdgeStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonSnowballEdgeStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *AmazonSnowballEdgeStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AmazonSnowballEdgeStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AmazonSnowballEdgeStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *AmazonSnowballEdgeStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AmazonSnowballEdgeStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AmazonSnowballEdgeStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetEnableTaskLimit

`func (o *AmazonSnowballEdgeStorageImportSpec) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *AmazonSnowballEdgeStorageImportSpec) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *AmazonSnowballEdgeStorageImportSpec) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *AmazonSnowballEdgeStorageImportSpec) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *AmazonSnowballEdgeStorageImportSpec) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *AmazonSnowballEdgeStorageImportSpec) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *AmazonSnowballEdgeStorageImportSpec) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *AmazonSnowballEdgeStorageImportSpec) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *AmazonSnowballEdgeStorageImportSpec) GetAccount() AmazonSnowballEdgeStorageAccountImportModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AmazonSnowballEdgeStorageImportSpec) GetAccountOk() (*AmazonSnowballEdgeStorageAccountImportModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AmazonSnowballEdgeStorageImportSpec) SetAccount(v AmazonSnowballEdgeStorageAccountImportModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *AmazonSnowballEdgeStorageImportSpec) GetBucket() AmazonSnowballEdgeStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AmazonSnowballEdgeStorageImportSpec) GetBucketOk() (*AmazonSnowballEdgeStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AmazonSnowballEdgeStorageImportSpec) SetBucket(v AmazonSnowballEdgeStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *AmazonSnowballEdgeStorageImportSpec) GetMountServer() MountServerSettingsImportSpec`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *AmazonSnowballEdgeStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *AmazonSnowballEdgeStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


