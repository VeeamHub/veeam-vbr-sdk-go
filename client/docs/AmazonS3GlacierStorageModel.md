# AmazonS3GlacierStorageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**AmazonS3StorageAccountModel**](AmazonS3StorageAccountModel.md) |  | 
**Bucket** | [**AmazonS3GlacierStorageBucketModel**](AmazonS3GlacierStorageBucketModel.md) |  | 
**ProxyAppliance** | Pointer to [**AmazonS3StorageProxyApplianceModel**](AmazonS3StorageProxyApplianceModel.md) |  | [optional] 

## Methods

### NewAmazonS3GlacierStorageModel

`func NewAmazonS3GlacierStorageModel(account AmazonS3StorageAccountModel, bucket AmazonS3GlacierStorageBucketModel, ) *AmazonS3GlacierStorageModel`

NewAmazonS3GlacierStorageModel instantiates a new AmazonS3GlacierStorageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3GlacierStorageModelWithDefaults

`func NewAmazonS3GlacierStorageModelWithDefaults() *AmazonS3GlacierStorageModel`

NewAmazonS3GlacierStorageModelWithDefaults instantiates a new AmazonS3GlacierStorageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AmazonS3GlacierStorageModel) GetAccount() AmazonS3StorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AmazonS3GlacierStorageModel) GetAccountOk() (*AmazonS3StorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AmazonS3GlacierStorageModel) SetAccount(v AmazonS3StorageAccountModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *AmazonS3GlacierStorageModel) GetBucket() AmazonS3GlacierStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AmazonS3GlacierStorageModel) GetBucketOk() (*AmazonS3GlacierStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AmazonS3GlacierStorageModel) SetBucket(v AmazonS3GlacierStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetProxyAppliance

`func (o *AmazonS3GlacierStorageModel) GetProxyAppliance() AmazonS3StorageProxyApplianceModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *AmazonS3GlacierStorageModel) GetProxyApplianceOk() (*AmazonS3StorageProxyApplianceModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *AmazonS3GlacierStorageModel) SetProxyAppliance(v AmazonS3StorageProxyApplianceModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *AmazonS3GlacierStorageModel) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


