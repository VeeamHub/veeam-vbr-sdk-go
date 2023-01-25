# AzureArchiveStorageAccountImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**CloudCredentialsImportModel**](CloudCredentialsImportModel.md) |  | 
**RegionType** | [**EAzureRegionType**](EAzureRegionType.md) |  | 
**ConnectionSettings** | Pointer to [**ObjectStorageConnectionImportSpec**](ObjectStorageConnectionImportSpec.md) |  | [optional] 

## Methods

### NewAzureArchiveStorageAccountImportModel

`func NewAzureArchiveStorageAccountImportModel(credentials CloudCredentialsImportModel, regionType EAzureRegionType, ) *AzureArchiveStorageAccountImportModel`

NewAzureArchiveStorageAccountImportModel instantiates a new AzureArchiveStorageAccountImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureArchiveStorageAccountImportModelWithDefaults

`func NewAzureArchiveStorageAccountImportModelWithDefaults() *AzureArchiveStorageAccountImportModel`

NewAzureArchiveStorageAccountImportModelWithDefaults instantiates a new AzureArchiveStorageAccountImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *AzureArchiveStorageAccountImportModel) GetCredentials() CloudCredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AzureArchiveStorageAccountImportModel) GetCredentialsOk() (*CloudCredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AzureArchiveStorageAccountImportModel) SetCredentials(v CloudCredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetRegionType

`func (o *AzureArchiveStorageAccountImportModel) GetRegionType() EAzureRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AzureArchiveStorageAccountImportModel) GetRegionTypeOk() (*EAzureRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AzureArchiveStorageAccountImportModel) SetRegionType(v EAzureRegionType)`

SetRegionType sets RegionType field to given value.


### GetConnectionSettings

`func (o *AzureArchiveStorageAccountImportModel) GetConnectionSettings() ObjectStorageConnectionImportSpec`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *AzureArchiveStorageAccountImportModel) GetConnectionSettingsOk() (*ObjectStorageConnectionImportSpec, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *AzureArchiveStorageAccountImportModel) SetConnectionSettings(v ObjectStorageConnectionImportSpec)`

SetConnectionSettings sets ConnectionSettings field to given value.

### HasConnectionSettings

`func (o *AzureArchiveStorageAccountImportModel) HasConnectionSettings() bool`

HasConnectionSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


