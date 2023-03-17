# WasabiCloudStorageAccountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | **string** | ID of a region where the storage is located. | 
**CredentialsId** | **string** | ID of a cloud credentials record used to access the storage. | 
**ConnectionSettings** | [**ObjectStorageConnectionModel**](ObjectStorageConnectionModel.md) |  | 

## Methods

### NewWasabiCloudStorageAccountModel

`func NewWasabiCloudStorageAccountModel(regionId string, credentialsId string, connectionSettings ObjectStorageConnectionModel, ) *WasabiCloudStorageAccountModel`

NewWasabiCloudStorageAccountModel instantiates a new WasabiCloudStorageAccountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWasabiCloudStorageAccountModelWithDefaults

`func NewWasabiCloudStorageAccountModelWithDefaults() *WasabiCloudStorageAccountModel`

NewWasabiCloudStorageAccountModelWithDefaults instantiates a new WasabiCloudStorageAccountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *WasabiCloudStorageAccountModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *WasabiCloudStorageAccountModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *WasabiCloudStorageAccountModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetCredentialsId

`func (o *WasabiCloudStorageAccountModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *WasabiCloudStorageAccountModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *WasabiCloudStorageAccountModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetConnectionSettings

`func (o *WasabiCloudStorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *WasabiCloudStorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *WasabiCloudStorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel)`

SetConnectionSettings sets ConnectionSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


