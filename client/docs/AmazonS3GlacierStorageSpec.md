# AmazonS3GlacierStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**AmazonS3StorageAccountModel**](AmazonS3StorageAccountModel.md) |  | 
**Bucket** | [**AmazonS3GlacierStorageBucketModel**](AmazonS3GlacierStorageBucketModel.md) |  | 
**ProxyAppliance** | Pointer to [**AmazonS3StorageProxyApplianceModel**](AmazonS3StorageProxyApplianceModel.md) |  | [optional] 

## Methods

### NewAmazonS3GlacierStorageSpec

`func NewAmazonS3GlacierStorageSpec(account AmazonS3StorageAccountModel, bucket AmazonS3GlacierStorageBucketModel, ) *AmazonS3GlacierStorageSpec`

NewAmazonS3GlacierStorageSpec instantiates a new AmazonS3GlacierStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3GlacierStorageSpecWithDefaults

`func NewAmazonS3GlacierStorageSpecWithDefaults() *AmazonS3GlacierStorageSpec`

NewAmazonS3GlacierStorageSpecWithDefaults instantiates a new AmazonS3GlacierStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AmazonS3GlacierStorageSpec) GetAccount() AmazonS3StorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AmazonS3GlacierStorageSpec) GetAccountOk() (*AmazonS3StorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AmazonS3GlacierStorageSpec) SetAccount(v AmazonS3StorageAccountModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *AmazonS3GlacierStorageSpec) GetBucket() AmazonS3GlacierStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AmazonS3GlacierStorageSpec) GetBucketOk() (*AmazonS3GlacierStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AmazonS3GlacierStorageSpec) SetBucket(v AmazonS3GlacierStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetProxyAppliance

`func (o *AmazonS3GlacierStorageSpec) GetProxyAppliance() AmazonS3StorageProxyApplianceModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *AmazonS3GlacierStorageSpec) GetProxyApplianceOk() (*AmazonS3StorageProxyApplianceModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *AmazonS3GlacierStorageSpec) SetProxyAppliance(v AmazonS3StorageProxyApplianceModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *AmazonS3GlacierStorageSpec) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


