# AmazonS3GlacierStorageBucketModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | **string** | ID of a region where the Amazon S3 bucket is located. | 
**BucketName** | **string** | Name of an Amazon S3 Glacier bucket. | 
**FolderName** | **string** | Name of the folder that the object storage repository is mapped to. | 
**ImmutabilityEnabled** | Pointer to **bool** | If *true*, storage immutability is enabled. | [optional] 
**UseDeepArchive** | Pointer to **bool** | If *true*, Glacier Deep Archive is used for backups with the retention policy over 180 days. | [optional] 

## Methods

### NewAmazonS3GlacierStorageBucketModel

`func NewAmazonS3GlacierStorageBucketModel(regionId string, bucketName string, folderName string, ) *AmazonS3GlacierStorageBucketModel`

NewAmazonS3GlacierStorageBucketModel instantiates a new AmazonS3GlacierStorageBucketModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3GlacierStorageBucketModelWithDefaults

`func NewAmazonS3GlacierStorageBucketModelWithDefaults() *AmazonS3GlacierStorageBucketModel`

NewAmazonS3GlacierStorageBucketModelWithDefaults instantiates a new AmazonS3GlacierStorageBucketModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *AmazonS3GlacierStorageBucketModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AmazonS3GlacierStorageBucketModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AmazonS3GlacierStorageBucketModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetBucketName

`func (o *AmazonS3GlacierStorageBucketModel) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *AmazonS3GlacierStorageBucketModel) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *AmazonS3GlacierStorageBucketModel) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetFolderName

`func (o *AmazonS3GlacierStorageBucketModel) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *AmazonS3GlacierStorageBucketModel) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *AmazonS3GlacierStorageBucketModel) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.


### GetImmutabilityEnabled

`func (o *AmazonS3GlacierStorageBucketModel) GetImmutabilityEnabled() bool`

GetImmutabilityEnabled returns the ImmutabilityEnabled field if non-nil, zero value otherwise.

### GetImmutabilityEnabledOk

`func (o *AmazonS3GlacierStorageBucketModel) GetImmutabilityEnabledOk() (*bool, bool)`

GetImmutabilityEnabledOk returns a tuple with the ImmutabilityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutabilityEnabled

`func (o *AmazonS3GlacierStorageBucketModel) SetImmutabilityEnabled(v bool)`

SetImmutabilityEnabled sets ImmutabilityEnabled field to given value.

### HasImmutabilityEnabled

`func (o *AmazonS3GlacierStorageBucketModel) HasImmutabilityEnabled() bool`

HasImmutabilityEnabled returns a boolean if a field has been set.

### GetUseDeepArchive

`func (o *AmazonS3GlacierStorageBucketModel) GetUseDeepArchive() bool`

GetUseDeepArchive returns the UseDeepArchive field if non-nil, zero value otherwise.

### GetUseDeepArchiveOk

`func (o *AmazonS3GlacierStorageBucketModel) GetUseDeepArchiveOk() (*bool, bool)`

GetUseDeepArchiveOk returns a tuple with the UseDeepArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDeepArchive

`func (o *AmazonS3GlacierStorageBucketModel) SetUseDeepArchive(v bool)`

SetUseDeepArchive sets UseDeepArchive field to given value.

### HasUseDeepArchive

`func (o *AmazonS3GlacierStorageBucketModel) HasUseDeepArchive() bool`

HasUseDeepArchive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


