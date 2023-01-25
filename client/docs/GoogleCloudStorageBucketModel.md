# GoogleCloudStorageBucketModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | Name of a Google Cloud bucket. | 
**FolderName** | **string** | Name of the folder that the object storage repository is mapped to. | 
**StorageConsumptionLimit** | Pointer to [**ObjectStorageConsumptionLimitModel**](ObjectStorageConsumptionLimitModel.md) |  | [optional] 
**NearlineStorageEnabled** | Pointer to **bool** | If *true*, the nearline storage class is used. | [optional] 

## Methods

### NewGoogleCloudStorageBucketModel

`func NewGoogleCloudStorageBucketModel(bucketName string, folderName string, ) *GoogleCloudStorageBucketModel`

NewGoogleCloudStorageBucketModel instantiates a new GoogleCloudStorageBucketModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudStorageBucketModelWithDefaults

`func NewGoogleCloudStorageBucketModelWithDefaults() *GoogleCloudStorageBucketModel`

NewGoogleCloudStorageBucketModelWithDefaults instantiates a new GoogleCloudStorageBucketModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *GoogleCloudStorageBucketModel) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *GoogleCloudStorageBucketModel) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *GoogleCloudStorageBucketModel) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetFolderName

`func (o *GoogleCloudStorageBucketModel) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *GoogleCloudStorageBucketModel) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *GoogleCloudStorageBucketModel) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.


### GetStorageConsumptionLimit

`func (o *GoogleCloudStorageBucketModel) GetStorageConsumptionLimit() ObjectStorageConsumptionLimitModel`

GetStorageConsumptionLimit returns the StorageConsumptionLimit field if non-nil, zero value otherwise.

### GetStorageConsumptionLimitOk

`func (o *GoogleCloudStorageBucketModel) GetStorageConsumptionLimitOk() (*ObjectStorageConsumptionLimitModel, bool)`

GetStorageConsumptionLimitOk returns a tuple with the StorageConsumptionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumptionLimit

`func (o *GoogleCloudStorageBucketModel) SetStorageConsumptionLimit(v ObjectStorageConsumptionLimitModel)`

SetStorageConsumptionLimit sets StorageConsumptionLimit field to given value.

### HasStorageConsumptionLimit

`func (o *GoogleCloudStorageBucketModel) HasStorageConsumptionLimit() bool`

HasStorageConsumptionLimit returns a boolean if a field has been set.

### GetNearlineStorageEnabled

`func (o *GoogleCloudStorageBucketModel) GetNearlineStorageEnabled() bool`

GetNearlineStorageEnabled returns the NearlineStorageEnabled field if non-nil, zero value otherwise.

### GetNearlineStorageEnabledOk

`func (o *GoogleCloudStorageBucketModel) GetNearlineStorageEnabledOk() (*bool, bool)`

GetNearlineStorageEnabledOk returns a tuple with the NearlineStorageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearlineStorageEnabled

`func (o *GoogleCloudStorageBucketModel) SetNearlineStorageEnabled(v bool)`

SetNearlineStorageEnabled sets NearlineStorageEnabled field to given value.

### HasNearlineStorageEnabled

`func (o *GoogleCloudStorageBucketModel) HasNearlineStorageEnabled() bool`

HasNearlineStorageEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


