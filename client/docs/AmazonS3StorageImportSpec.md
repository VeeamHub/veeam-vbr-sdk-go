# AmazonS3StorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the object storage repository. | 
**Description** | **string** | Description of the object storage repository. | 
**Tag** | **string** | Tag that identifies the object storage repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**AmazonS3StorageAccountImportModel**](AmazonS3StorageAccountImportModel.md) |  | 
**Bucket** | [**AmazonS3StorageBucketModel**](AmazonS3StorageBucketModel.md) |  | 
**MountServer** | [**MountServerSettingsImportSpec**](MountServerSettingsImportSpec.md) |  | 
**ProxyAppliance** | Pointer to [**AmazonS3StorageProxyApplianceModel**](AmazonS3StorageProxyApplianceModel.md) |  | [optional] 

## Methods

### NewAmazonS3StorageImportSpec

`func NewAmazonS3StorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account AmazonS3StorageAccountImportModel, bucket AmazonS3StorageBucketModel, mountServer MountServerSettingsImportSpec, ) *AmazonS3StorageImportSpec`

NewAmazonS3StorageImportSpec instantiates a new AmazonS3StorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3StorageImportSpecWithDefaults

`func NewAmazonS3StorageImportSpecWithDefaults() *AmazonS3StorageImportSpec`

NewAmazonS3StorageImportSpecWithDefaults instantiates a new AmazonS3StorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AmazonS3StorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmazonS3StorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmazonS3StorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AmazonS3StorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonS3StorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonS3StorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *AmazonS3StorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AmazonS3StorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AmazonS3StorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *AmazonS3StorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AmazonS3StorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AmazonS3StorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetEnableTaskLimit

`func (o *AmazonS3StorageImportSpec) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *AmazonS3StorageImportSpec) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *AmazonS3StorageImportSpec) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *AmazonS3StorageImportSpec) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *AmazonS3StorageImportSpec) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *AmazonS3StorageImportSpec) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *AmazonS3StorageImportSpec) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *AmazonS3StorageImportSpec) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *AmazonS3StorageImportSpec) GetAccount() AmazonS3StorageAccountImportModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AmazonS3StorageImportSpec) GetAccountOk() (*AmazonS3StorageAccountImportModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AmazonS3StorageImportSpec) SetAccount(v AmazonS3StorageAccountImportModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *AmazonS3StorageImportSpec) GetBucket() AmazonS3StorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AmazonS3StorageImportSpec) GetBucketOk() (*AmazonS3StorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AmazonS3StorageImportSpec) SetBucket(v AmazonS3StorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetMountServer

`func (o *AmazonS3StorageImportSpec) GetMountServer() MountServerSettingsImportSpec`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *AmazonS3StorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *AmazonS3StorageImportSpec) SetMountServer(v MountServerSettingsImportSpec)`

SetMountServer sets MountServer field to given value.


### GetProxyAppliance

`func (o *AmazonS3StorageImportSpec) GetProxyAppliance() AmazonS3StorageProxyApplianceModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *AmazonS3StorageImportSpec) GetProxyApplianceOk() (*AmazonS3StorageProxyApplianceModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *AmazonS3StorageImportSpec) SetProxyAppliance(v AmazonS3StorageProxyApplianceModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *AmazonS3StorageImportSpec) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


