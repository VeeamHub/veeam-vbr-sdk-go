# AmazonS3StorageBucketModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | **string** | ID of a region where the Amazon S3 bucket is located. | 
**BucketName** | **string** | Name of an Amazon S3 bucket. | 
**FolderName** | **string** | Name of the folder that the object storage repository is mapped to. | 
**StorageConsumptionLimit** | Pointer to [**ObjectStorageConsumptionLimitModel**](ObjectStorageConsumptionLimitModel.md) |  | [optional] 
**Immutability** | Pointer to [**ObjectStorageImmutabilityModel**](ObjectStorageImmutabilityModel.md) |  | [optional] 
**InfrequentAccessStorage** | Pointer to [**AmazonS3IAStorageModel**](AmazonS3IAStorageModel.md) |  | [optional] 

## Methods

### NewAmazonS3StorageBucketModel

`func NewAmazonS3StorageBucketModel(regionId string, bucketName string, folderName string, ) *AmazonS3StorageBucketModel`

NewAmazonS3StorageBucketModel instantiates a new AmazonS3StorageBucketModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3StorageBucketModelWithDefaults

`func NewAmazonS3StorageBucketModelWithDefaults() *AmazonS3StorageBucketModel`

NewAmazonS3StorageBucketModelWithDefaults instantiates a new AmazonS3StorageBucketModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *AmazonS3StorageBucketModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AmazonS3StorageBucketModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AmazonS3StorageBucketModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetBucketName

`func (o *AmazonS3StorageBucketModel) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *AmazonS3StorageBucketModel) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *AmazonS3StorageBucketModel) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetFolderName

`func (o *AmazonS3StorageBucketModel) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *AmazonS3StorageBucketModel) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *AmazonS3StorageBucketModel) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.


### GetStorageConsumptionLimit

`func (o *AmazonS3StorageBucketModel) GetStorageConsumptionLimit() ObjectStorageConsumptionLimitModel`

GetStorageConsumptionLimit returns the StorageConsumptionLimit field if non-nil, zero value otherwise.

### GetStorageConsumptionLimitOk

`func (o *AmazonS3StorageBucketModel) GetStorageConsumptionLimitOk() (*ObjectStorageConsumptionLimitModel, bool)`

GetStorageConsumptionLimitOk returns a tuple with the StorageConsumptionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumptionLimit

`func (o *AmazonS3StorageBucketModel) SetStorageConsumptionLimit(v ObjectStorageConsumptionLimitModel)`

SetStorageConsumptionLimit sets StorageConsumptionLimit field to given value.

### HasStorageConsumptionLimit

`func (o *AmazonS3StorageBucketModel) HasStorageConsumptionLimit() bool`

HasStorageConsumptionLimit returns a boolean if a field has been set.

### GetImmutability

`func (o *AmazonS3StorageBucketModel) GetImmutability() ObjectStorageImmutabilityModel`

GetImmutability returns the Immutability field if non-nil, zero value otherwise.

### GetImmutabilityOk

`func (o *AmazonS3StorageBucketModel) GetImmutabilityOk() (*ObjectStorageImmutabilityModel, bool)`

GetImmutabilityOk returns a tuple with the Immutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutability

`func (o *AmazonS3StorageBucketModel) SetImmutability(v ObjectStorageImmutabilityModel)`

SetImmutability sets Immutability field to given value.

### HasImmutability

`func (o *AmazonS3StorageBucketModel) HasImmutability() bool`

HasImmutability returns a boolean if a field has been set.

### GetInfrequentAccessStorage

`func (o *AmazonS3StorageBucketModel) GetInfrequentAccessStorage() AmazonS3IAStorageModel`

GetInfrequentAccessStorage returns the InfrequentAccessStorage field if non-nil, zero value otherwise.

### GetInfrequentAccessStorageOk

`func (o *AmazonS3StorageBucketModel) GetInfrequentAccessStorageOk() (*AmazonS3IAStorageModel, bool)`

GetInfrequentAccessStorageOk returns a tuple with the InfrequentAccessStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrequentAccessStorage

`func (o *AmazonS3StorageBucketModel) SetInfrequentAccessStorage(v AmazonS3IAStorageModel)`

SetInfrequentAccessStorage sets InfrequentAccessStorage field to given value.

### HasInfrequentAccessStorage

`func (o *AmazonS3StorageBucketModel) HasInfrequentAccessStorage() bool`

HasInfrequentAccessStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


