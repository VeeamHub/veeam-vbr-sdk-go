# AmazonS3StorageAccountImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**CloudCredentialsImportModel**](CloudCredentialsImportModel.md) |  | 
**RegionType** | [**EAmazonRegionType**](EAmazonRegionType.md) |  | 
**ConnectionSettings** | [**ObjectStorageConnectionImportSpec**](ObjectStorageConnectionImportSpec.md) |  | 

## Methods

### NewAmazonS3StorageAccountImportModel

`func NewAmazonS3StorageAccountImportModel(credentials CloudCredentialsImportModel, regionType EAmazonRegionType, connectionSettings ObjectStorageConnectionImportSpec, ) *AmazonS3StorageAccountImportModel`

NewAmazonS3StorageAccountImportModel instantiates a new AmazonS3StorageAccountImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3StorageAccountImportModelWithDefaults

`func NewAmazonS3StorageAccountImportModelWithDefaults() *AmazonS3StorageAccountImportModel`

NewAmazonS3StorageAccountImportModelWithDefaults instantiates a new AmazonS3StorageAccountImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *AmazonS3StorageAccountImportModel) GetCredentials() CloudCredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AmazonS3StorageAccountImportModel) GetCredentialsOk() (*CloudCredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AmazonS3StorageAccountImportModel) SetCredentials(v CloudCredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetRegionType

`func (o *AmazonS3StorageAccountImportModel) GetRegionType() EAmazonRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AmazonS3StorageAccountImportModel) GetRegionTypeOk() (*EAmazonRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AmazonS3StorageAccountImportModel) SetRegionType(v EAmazonRegionType)`

SetRegionType sets RegionType field to given value.


### GetConnectionSettings

`func (o *AmazonS3StorageAccountImportModel) GetConnectionSettings() ObjectStorageConnectionImportSpec`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *AmazonS3StorageAccountImportModel) GetConnectionSettingsOk() (*ObjectStorageConnectionImportSpec, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *AmazonS3StorageAccountImportModel) SetConnectionSettings(v ObjectStorageConnectionImportSpec)`

SetConnectionSettings sets ConnectionSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


