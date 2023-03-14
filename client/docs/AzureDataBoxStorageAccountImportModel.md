# AzureDataBoxStorageAccountImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceEndpoint** | **string** | Service endpoint address of the Azure Data Box device. | 
**Credentials** | [**CloudCredentialsImportModel**](CloudCredentialsImportModel.md) |  | 
**ConnectionSettings** | [**ObjectStorageConnectionImportSpec**](ObjectStorageConnectionImportSpec.md) |  | 

## Methods

### NewAzureDataBoxStorageAccountImportModel

`func NewAzureDataBoxStorageAccountImportModel(serviceEndpoint string, credentials CloudCredentialsImportModel, connectionSettings ObjectStorageConnectionImportSpec, ) *AzureDataBoxStorageAccountImportModel`

NewAzureDataBoxStorageAccountImportModel instantiates a new AzureDataBoxStorageAccountImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataBoxStorageAccountImportModelWithDefaults

`func NewAzureDataBoxStorageAccountImportModelWithDefaults() *AzureDataBoxStorageAccountImportModel`

NewAzureDataBoxStorageAccountImportModelWithDefaults instantiates a new AzureDataBoxStorageAccountImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceEndpoint

`func (o *AzureDataBoxStorageAccountImportModel) GetServiceEndpoint() string`

GetServiceEndpoint returns the ServiceEndpoint field if non-nil, zero value otherwise.

### GetServiceEndpointOk

`func (o *AzureDataBoxStorageAccountImportModel) GetServiceEndpointOk() (*string, bool)`

GetServiceEndpointOk returns a tuple with the ServiceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpoint

`func (o *AzureDataBoxStorageAccountImportModel) SetServiceEndpoint(v string)`

SetServiceEndpoint sets ServiceEndpoint field to given value.


### GetCredentials

`func (o *AzureDataBoxStorageAccountImportModel) GetCredentials() CloudCredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AzureDataBoxStorageAccountImportModel) GetCredentialsOk() (*CloudCredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AzureDataBoxStorageAccountImportModel) SetCredentials(v CloudCredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetConnectionSettings

`func (o *AzureDataBoxStorageAccountImportModel) GetConnectionSettings() ObjectStorageConnectionImportSpec`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *AzureDataBoxStorageAccountImportModel) GetConnectionSettingsOk() (*ObjectStorageConnectionImportSpec, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *AzureDataBoxStorageAccountImportModel) SetConnectionSettings(v ObjectStorageConnectionImportSpec)`

SetConnectionSettings sets ConnectionSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


