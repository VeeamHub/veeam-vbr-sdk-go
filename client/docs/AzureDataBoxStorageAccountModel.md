# AzureDataBoxStorageAccountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceEndpoint** | **string** | Service endpoint address of the Azure Data Box device. | 
**CredentialsId** | **string** | ID of the cloud credentials record. | 
**ConnectionSettings** | Pointer to [**ObjectStorageConnectionModel**](ObjectStorageConnectionModel.md) |  | [optional] 

## Methods

### NewAzureDataBoxStorageAccountModel

`func NewAzureDataBoxStorageAccountModel(serviceEndpoint string, credentialsId string, ) *AzureDataBoxStorageAccountModel`

NewAzureDataBoxStorageAccountModel instantiates a new AzureDataBoxStorageAccountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataBoxStorageAccountModelWithDefaults

`func NewAzureDataBoxStorageAccountModelWithDefaults() *AzureDataBoxStorageAccountModel`

NewAzureDataBoxStorageAccountModelWithDefaults instantiates a new AzureDataBoxStorageAccountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceEndpoint

`func (o *AzureDataBoxStorageAccountModel) GetServiceEndpoint() string`

GetServiceEndpoint returns the ServiceEndpoint field if non-nil, zero value otherwise.

### GetServiceEndpointOk

`func (o *AzureDataBoxStorageAccountModel) GetServiceEndpointOk() (*string, bool)`

GetServiceEndpointOk returns a tuple with the ServiceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpoint

`func (o *AzureDataBoxStorageAccountModel) SetServiceEndpoint(v string)`

SetServiceEndpoint sets ServiceEndpoint field to given value.


### GetCredentialsId

`func (o *AzureDataBoxStorageAccountModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *AzureDataBoxStorageAccountModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *AzureDataBoxStorageAccountModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetConnectionSettings

`func (o *AzureDataBoxStorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *AzureDataBoxStorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *AzureDataBoxStorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel)`

SetConnectionSettings sets ConnectionSettings field to given value.

### HasConnectionSettings

`func (o *AzureDataBoxStorageAccountModel) HasConnectionSettings() bool`

HasConnectionSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


