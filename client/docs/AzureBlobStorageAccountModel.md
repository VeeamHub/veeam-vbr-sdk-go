# AzureBlobStorageAccountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsId** | **string** | ID of the cloud credentials record. | 
**RegionType** | [**EAzureRegionType**](EAzureRegionType.md) |  | 
**ConnectionSettings** | [**ObjectStorageConnectionModel**](ObjectStorageConnectionModel.md) |  | 

## Methods

### NewAzureBlobStorageAccountModel

`func NewAzureBlobStorageAccountModel(credentialsId string, regionType EAzureRegionType, connectionSettings ObjectStorageConnectionModel, ) *AzureBlobStorageAccountModel`

NewAzureBlobStorageAccountModel instantiates a new AzureBlobStorageAccountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobStorageAccountModelWithDefaults

`func NewAzureBlobStorageAccountModelWithDefaults() *AzureBlobStorageAccountModel`

NewAzureBlobStorageAccountModelWithDefaults instantiates a new AzureBlobStorageAccountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsId

`func (o *AzureBlobStorageAccountModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *AzureBlobStorageAccountModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *AzureBlobStorageAccountModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetRegionType

`func (o *AzureBlobStorageAccountModel) GetRegionType() EAzureRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AzureBlobStorageAccountModel) GetRegionTypeOk() (*EAzureRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AzureBlobStorageAccountModel) SetRegionType(v EAzureRegionType)`

SetRegionType sets RegionType field to given value.


### GetConnectionSettings

`func (o *AzureBlobStorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *AzureBlobStorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *AzureBlobStorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel)`

SetConnectionSettings sets ConnectionSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


