# IBMCloudStorageBucketModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | Name of an IBM Cloud bucket. | 
**FolderName** | **string** | Name of the folder that the object storage repository is mapped to. | 
**StorageConsumptionLimit** | Pointer to [**ObjectStorageConsumptionLimitModel**](ObjectStorageConsumptionLimitModel.md) |  | [optional] 
**Immutability** | Pointer to [**ObjectStorageImmutabilityModel**](ObjectStorageImmutabilityModel.md) |  | [optional] 

## Methods

### NewIBMCloudStorageBucketModel

`func NewIBMCloudStorageBucketModel(bucketName string, folderName string, ) *IBMCloudStorageBucketModel`

NewIBMCloudStorageBucketModel instantiates a new IBMCloudStorageBucketModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIBMCloudStorageBucketModelWithDefaults

`func NewIBMCloudStorageBucketModelWithDefaults() *IBMCloudStorageBucketModel`

NewIBMCloudStorageBucketModelWithDefaults instantiates a new IBMCloudStorageBucketModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *IBMCloudStorageBucketModel) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *IBMCloudStorageBucketModel) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *IBMCloudStorageBucketModel) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetFolderName

`func (o *IBMCloudStorageBucketModel) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *IBMCloudStorageBucketModel) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *IBMCloudStorageBucketModel) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.


### GetStorageConsumptionLimit

`func (o *IBMCloudStorageBucketModel) GetStorageConsumptionLimit() ObjectStorageConsumptionLimitModel`

GetStorageConsumptionLimit returns the StorageConsumptionLimit field if non-nil, zero value otherwise.

### GetStorageConsumptionLimitOk

`func (o *IBMCloudStorageBucketModel) GetStorageConsumptionLimitOk() (*ObjectStorageConsumptionLimitModel, bool)`

GetStorageConsumptionLimitOk returns a tuple with the StorageConsumptionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConsumptionLimit

`func (o *IBMCloudStorageBucketModel) SetStorageConsumptionLimit(v ObjectStorageConsumptionLimitModel)`

SetStorageConsumptionLimit sets StorageConsumptionLimit field to given value.

### HasStorageConsumptionLimit

`func (o *IBMCloudStorageBucketModel) HasStorageConsumptionLimit() bool`

HasStorageConsumptionLimit returns a boolean if a field has been set.

### GetImmutability

`func (o *IBMCloudStorageBucketModel) GetImmutability() ObjectStorageImmutabilityModel`

GetImmutability returns the Immutability field if non-nil, zero value otherwise.

### GetImmutabilityOk

`func (o *IBMCloudStorageBucketModel) GetImmutabilityOk() (*ObjectStorageImmutabilityModel, bool)`

GetImmutabilityOk returns a tuple with the Immutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutability

`func (o *IBMCloudStorageBucketModel) SetImmutability(v ObjectStorageImmutabilityModel)`

SetImmutability sets Immutability field to given value.

### HasImmutability

`func (o *IBMCloudStorageBucketModel) HasImmutability() bool`

HasImmutability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


