# GoogleCloudStorageAccountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsId** | **string** | ID of a cloud credentials record used to access the storage. | 
**RegionId** | **string** | ID of a region where the storage bucket is located. | 
**ConnectionSettings** | [**ObjectStorageConnectionModel**](ObjectStorageConnectionModel.md) |  | 

## Methods

### NewGoogleCloudStorageAccountModel

`func NewGoogleCloudStorageAccountModel(credentialsId string, regionId string, connectionSettings ObjectStorageConnectionModel, ) *GoogleCloudStorageAccountModel`

NewGoogleCloudStorageAccountModel instantiates a new GoogleCloudStorageAccountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudStorageAccountModelWithDefaults

`func NewGoogleCloudStorageAccountModelWithDefaults() *GoogleCloudStorageAccountModel`

NewGoogleCloudStorageAccountModelWithDefaults instantiates a new GoogleCloudStorageAccountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsId

`func (o *GoogleCloudStorageAccountModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *GoogleCloudStorageAccountModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *GoogleCloudStorageAccountModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetRegionId

`func (o *GoogleCloudStorageAccountModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GoogleCloudStorageAccountModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GoogleCloudStorageAccountModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetConnectionSettings

`func (o *GoogleCloudStorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *GoogleCloudStorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *GoogleCloudStorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel)`

SetConnectionSettings sets ConnectionSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


