# AmazonS3GlacierStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the object storage repository. | 
**Description** | **string** | Description of the object storage repository. | 
**Tag** | **string** | Tag that identifies the object storage repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**Account** | [**AmazonS3StorageAccountImportModel**](AmazonS3StorageAccountImportModel.md) |  | 
**Bucket** | [**AmazonS3GlacierStorageBucketModel**](AmazonS3GlacierStorageBucketModel.md) |  | 
**ProxyAppliance** | Pointer to [**AmazonS3StorageProxyApplianceModel**](AmazonS3StorageProxyApplianceModel.md) |  | [optional] 

## Methods

### NewAmazonS3GlacierStorageImportSpec

`func NewAmazonS3GlacierStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account AmazonS3StorageAccountImportModel, bucket AmazonS3GlacierStorageBucketModel, ) *AmazonS3GlacierStorageImportSpec`

NewAmazonS3GlacierStorageImportSpec instantiates a new AmazonS3GlacierStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3GlacierStorageImportSpecWithDefaults

`func NewAmazonS3GlacierStorageImportSpecWithDefaults() *AmazonS3GlacierStorageImportSpec`

NewAmazonS3GlacierStorageImportSpecWithDefaults instantiates a new AmazonS3GlacierStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AmazonS3GlacierStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmazonS3GlacierStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmazonS3GlacierStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AmazonS3GlacierStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AmazonS3GlacierStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AmazonS3GlacierStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *AmazonS3GlacierStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AmazonS3GlacierStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AmazonS3GlacierStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *AmazonS3GlacierStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AmazonS3GlacierStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AmazonS3GlacierStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetAccount

`func (o *AmazonS3GlacierStorageImportSpec) GetAccount() AmazonS3StorageAccountImportModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AmazonS3GlacierStorageImportSpec) GetAccountOk() (*AmazonS3StorageAccountImportModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AmazonS3GlacierStorageImportSpec) SetAccount(v AmazonS3StorageAccountImportModel)`

SetAccount sets Account field to given value.


### GetBucket

`func (o *AmazonS3GlacierStorageImportSpec) GetBucket() AmazonS3GlacierStorageBucketModel`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AmazonS3GlacierStorageImportSpec) GetBucketOk() (*AmazonS3GlacierStorageBucketModel, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AmazonS3GlacierStorageImportSpec) SetBucket(v AmazonS3GlacierStorageBucketModel)`

SetBucket sets Bucket field to given value.


### GetProxyAppliance

`func (o *AmazonS3GlacierStorageImportSpec) GetProxyAppliance() AmazonS3StorageProxyApplianceModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *AmazonS3GlacierStorageImportSpec) GetProxyApplianceOk() (*AmazonS3StorageProxyApplianceModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *AmazonS3GlacierStorageImportSpec) SetProxyAppliance(v AmazonS3StorageProxyApplianceModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *AmazonS3GlacierStorageImportSpec) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


