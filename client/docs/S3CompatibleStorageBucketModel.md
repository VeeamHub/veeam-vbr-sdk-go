# S3CompatibleStorageBucketModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | Bucket name. | 
**FolderName** | **string** | Name of the folder that the object storage repository is mapped to. | 
**StorageConsumptionLimit** | Pointer to [**ObjectStorageConsumptionLimitModel**](ObjectStorageConsumptionLimitModel.md) |  | [optional] 
**Immutability** | Pointer to [**ObjectStorageImmutabilityModel**](ObjectStorageImmutabilityModel.md) |  | [optional] 

## Methods

### NewS3CompatibleStorageBucketModel

`func NewS3CompatibleStorageBucketModel(bucketName string, folderName string, ) *S3CompatibleStorageBucketModel`

NewS3CompatibleStorageBucketModel instantiates a new S3CompatibleStorageBucketModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CompatibleStorageBucketModelWithDefaults

`func NewS3CompatibleStorageBucketModelWithDefaults() *S3CompatibleStorageBucketModel`

NewS3CompatibleStorageBucketModelWithDefaults instantiates a new S3CompatibleStorageBucketModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *S3CompatibleStorageBucketModel) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3CompatibleStorageBucketModel) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3CompatibleStorageBucketModel) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetFolderName

`func (o *S3CompatibleStorageBucketModel) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *S3CompatibleStorageBucketModel) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *S3CompatibleStorageBucketModel) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.


### GetStorageConsumptionLimit

`func (o *S3CompatibleStorageBucketModel) GetStorageConsumptionLimit() ObjectStorageConsumptionLimitModel`

GetStorageConsumptionLimit returns the StorageConsumptionLimit field if non-nil, zero value otherwise.

### GetStorageConsumptionLimitOk

`func (o *S3CompatibleStorageBucketModel) GetStorageConsumptionLimitOk() (*ObjectStorageConsumptionLimitModel, bool)`

GetStorageConsumptionLimitOk returns a tuple with the StorageConsumptionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumptionLimit

`func (o *S3CompatibleStorageBucketModel) SetStorageConsumptionLimit(v ObjectStorageConsumptionLimitModel)`

SetStorageConsumptionLimit sets StorageConsumptionLimit field to given value.

### HasStorageConsumptionLimit

`func (o *S3CompatibleStorageBucketModel) HasStorageConsumptionLimit() bool`

HasStorageConsumptionLimit returns a boolean if a field has been set.

### GetImmutability

`func (o *S3CompatibleStorageBucketModel) GetImmutability() ObjectStorageImmutabilityModel`

GetImmutability returns the Immutability field if non-nil, zero value otherwise.

### GetImmutabilityOk

`func (o *S3CompatibleStorageBucketModel) GetImmutabilityOk() (*ObjectStorageImmutabilityModel, bool)`

GetImmutabilityOk returns a tuple with the Immutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutability

`func (o *S3CompatibleStorageBucketModel) SetImmutability(v ObjectStorageImmutabilityModel)`

SetImmutability sets Immutability field to given value.

### HasImmutability

`func (o *S3CompatibleStorageBucketModel) HasImmutability() bool`

HasImmutability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


