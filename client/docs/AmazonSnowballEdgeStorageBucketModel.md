# AmazonSnowballEdgeStorageBucketModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | Name of the Amazon S3 bucket. | [optional] 
**FolderName** | Pointer to **string** | Name of the folder that the object storage repository is mapped to. | [optional] 
**StorageConsumptionLimit** | Pointer to [**ObjectStorageConsumptionLimitModel**](ObjectStorageConsumptionLimitModel.md) |  | [optional] 

## Methods

### NewAmazonSnowballEdgeStorageBucketModel

`func NewAmazonSnowballEdgeStorageBucketModel() *AmazonSnowballEdgeStorageBucketModel`

NewAmazonSnowballEdgeStorageBucketModel instantiates a new AmazonSnowballEdgeStorageBucketModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSnowballEdgeStorageBucketModelWithDefaults

`func NewAmazonSnowballEdgeStorageBucketModelWithDefaults() *AmazonSnowballEdgeStorageBucketModel`

NewAmazonSnowballEdgeStorageBucketModelWithDefaults instantiates a new AmazonSnowballEdgeStorageBucketModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *AmazonSnowballEdgeStorageBucketModel) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *AmazonSnowballEdgeStorageBucketModel) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *AmazonSnowballEdgeStorageBucketModel) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *AmazonSnowballEdgeStorageBucketModel) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetFolderName

`func (o *AmazonSnowballEdgeStorageBucketModel) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *AmazonSnowballEdgeStorageBucketModel) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *AmazonSnowballEdgeStorageBucketModel) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *AmazonSnowballEdgeStorageBucketModel) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### GetStorageConsumptionLimit

`func (o *AmazonSnowballEdgeStorageBucketModel) GetStorageConsumptionLimit() ObjectStorageConsumptionLimitModel`

GetStorageConsumptionLimit returns the StorageConsumptionLimit field if non-nil, zero value otherwise.

### GetStorageConsumptionLimitOk

`func (o *AmazonSnowballEdgeStorageBucketModel) GetStorageConsumptionLimitOk() (*ObjectStorageConsumptionLimitModel, bool)`

GetStorageConsumptionLimitOk returns a tuple with the StorageConsumptionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumptionLimit

`func (o *AmazonSnowballEdgeStorageBucketModel) SetStorageConsumptionLimit(v ObjectStorageConsumptionLimitModel)`

SetStorageConsumptionLimit sets StorageConsumptionLimit field to given value.

### HasStorageConsumptionLimit

`func (o *AmazonSnowballEdgeStorageBucketModel) HasStorageConsumptionLimit() bool`

HasStorageConsumptionLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


