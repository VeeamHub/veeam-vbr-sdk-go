# S3CompatibleStorageAccountImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePoint** | **string** | Endpoint address of the storage. | 
**RegionId** | **string** | ID of a region where the storage is located. | 
**Credentials** | [**CloudCredentialsImportModel**](CloudCredentialsImportModel.md) |  | 
**ConnectionSettings** | [**ObjectStorageConnectionImportSpec**](ObjectStorageConnectionImportSpec.md) |  | 

## Methods

### NewS3CompatibleStorageAccountImportModel

`func NewS3CompatibleStorageAccountImportModel(servicePoint string, regionId string, credentials CloudCredentialsImportModel, connectionSettings ObjectStorageConnectionImportSpec, ) *S3CompatibleStorageAccountImportModel`

NewS3CompatibleStorageAccountImportModel instantiates a new S3CompatibleStorageAccountImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CompatibleStorageAccountImportModelWithDefaults

`func NewS3CompatibleStorageAccountImportModelWithDefaults() *S3CompatibleStorageAccountImportModel`

NewS3CompatibleStorageAccountImportModelWithDefaults instantiates a new S3CompatibleStorageAccountImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePoint

`func (o *S3CompatibleStorageAccountImportModel) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *S3CompatibleStorageAccountImportModel) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *S3CompatibleStorageAccountImportModel) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.


### GetRegionId

`func (o *S3CompatibleStorageAccountImportModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *S3CompatibleStorageAccountImportModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *S3CompatibleStorageAccountImportModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetCredentials

`func (o *S3CompatibleStorageAccountImportModel) GetCredentials() CloudCredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *S3CompatibleStorageAccountImportModel) GetCredentialsOk() (*CloudCredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *S3CompatibleStorageAccountImportModel) SetCredentials(v CloudCredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetConnectionSettings

`func (o *S3CompatibleStorageAccountImportModel) GetConnectionSettings() ObjectStorageConnectionImportSpec`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *S3CompatibleStorageAccountImportModel) GetConnectionSettingsOk() (*ObjectStorageConnectionImportSpec, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *S3CompatibleStorageAccountImportModel) SetConnectionSettings(v ObjectStorageConnectionImportSpec)`

SetConnectionSettings sets ConnectionSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


