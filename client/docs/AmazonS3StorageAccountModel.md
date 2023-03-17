# AmazonS3StorageAccountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsId** | **string** | ID of the cloud credentials record. | 
**RegionType** | [**EAmazonRegionType**](EAmazonRegionType.md) |  | 
**ConnectionSettings** | Pointer to [**ObjectStorageConnectionModel**](ObjectStorageConnectionModel.md) |  | [optional] 

## Methods

### NewAmazonS3StorageAccountModel

`func NewAmazonS3StorageAccountModel(credentialsId string, regionType EAmazonRegionType, ) *AmazonS3StorageAccountModel`

NewAmazonS3StorageAccountModel instantiates a new AmazonS3StorageAccountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3StorageAccountModelWithDefaults

`func NewAmazonS3StorageAccountModelWithDefaults() *AmazonS3StorageAccountModel`

NewAmazonS3StorageAccountModelWithDefaults instantiates a new AmazonS3StorageAccountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsId

`func (o *AmazonS3StorageAccountModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *AmazonS3StorageAccountModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *AmazonS3StorageAccountModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetRegionType

`func (o *AmazonS3StorageAccountModel) GetRegionType() EAmazonRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AmazonS3StorageAccountModel) GetRegionTypeOk() (*EAmazonRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AmazonS3StorageAccountModel) SetRegionType(v EAmazonRegionType)`

SetRegionType sets RegionType field to given value.


### GetConnectionSettings

`func (o *AmazonS3StorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *AmazonS3StorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *AmazonS3StorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel)`

SetConnectionSettings sets ConnectionSettings field to given value.

### HasConnectionSettings

`func (o *AmazonS3StorageAccountModel) HasConnectionSettings() bool`

HasConnectionSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


