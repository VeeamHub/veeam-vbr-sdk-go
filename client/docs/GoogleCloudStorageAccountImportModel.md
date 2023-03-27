# GoogleCloudStorageAccountImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**CloudCredentialsImportModel**](CloudCredentialsImportModel.md) |  | 
**RegionId** | **string** | ID of a region where the storage bucket is located. | 
**ConnectionSettings** | [**ObjectStorageConnectionImportSpec**](ObjectStorageConnectionImportSpec.md) |  | 

## Methods

### NewGoogleCloudStorageAccountImportModel

`func NewGoogleCloudStorageAccountImportModel(credentials CloudCredentialsImportModel, regionId string, connectionSettings ObjectStorageConnectionImportSpec, ) *GoogleCloudStorageAccountImportModel`

NewGoogleCloudStorageAccountImportModel instantiates a new GoogleCloudStorageAccountImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudStorageAccountImportModelWithDefaults

`func NewGoogleCloudStorageAccountImportModelWithDefaults() *GoogleCloudStorageAccountImportModel`

NewGoogleCloudStorageAccountImportModelWithDefaults instantiates a new GoogleCloudStorageAccountImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *GoogleCloudStorageAccountImportModel) GetCredentials() CloudCredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GoogleCloudStorageAccountImportModel) GetCredentialsOk() (*CloudCredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GoogleCloudStorageAccountImportModel) SetCredentials(v CloudCredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetRegionId

`func (o *GoogleCloudStorageAccountImportModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GoogleCloudStorageAccountImportModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GoogleCloudStorageAccountImportModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetConnectionSettings

`func (o *GoogleCloudStorageAccountImportModel) GetConnectionSettings() ObjectStorageConnectionImportSpec`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *GoogleCloudStorageAccountImportModel) GetConnectionSettingsOk() (*ObjectStorageConnectionImportSpec, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *GoogleCloudStorageAccountImportModel) SetConnectionSettings(v ObjectStorageConnectionImportSpec)`

SetConnectionSettings sets ConnectionSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


