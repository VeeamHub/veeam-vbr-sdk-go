# S3CompatibleStorageAccountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePoint** | **string** | Endpoint address and port number of the storage. | 
**RegionId** | **string** | ID of a region where the storage is located. | 
**CredentialsId** | **string** | ID of a cloud credentials record used to access the storage. | 
**ConnectionSettings** | Pointer to [**ObjectStorageConnectionModel**](ObjectStorageConnectionModel.md) |  | [optional] 

## Methods

### NewS3CompatibleStorageAccountModel

`func NewS3CompatibleStorageAccountModel(servicePoint string, regionId string, credentialsId string, ) *S3CompatibleStorageAccountModel`

NewS3CompatibleStorageAccountModel instantiates a new S3CompatibleStorageAccountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CompatibleStorageAccountModelWithDefaults

`func NewS3CompatibleStorageAccountModelWithDefaults() *S3CompatibleStorageAccountModel`

NewS3CompatibleStorageAccountModelWithDefaults instantiates a new S3CompatibleStorageAccountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePoint

`func (o *S3CompatibleStorageAccountModel) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *S3CompatibleStorageAccountModel) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *S3CompatibleStorageAccountModel) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.


### GetRegionId

`func (o *S3CompatibleStorageAccountModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *S3CompatibleStorageAccountModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *S3CompatibleStorageAccountModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetCredentialsId

`func (o *S3CompatibleStorageAccountModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *S3CompatibleStorageAccountModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *S3CompatibleStorageAccountModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetConnectionSettings

`func (o *S3CompatibleStorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *S3CompatibleStorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *S3CompatibleStorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel)`

SetConnectionSettings sets ConnectionSettings field to given value.

### HasConnectionSettings

`func (o *S3CompatibleStorageAccountModel) HasConnectionSettings() bool`

HasConnectionSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


