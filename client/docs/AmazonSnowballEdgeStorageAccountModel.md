# AmazonSnowballEdgeStorageAccountModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePoint** | **string** | Service point address and port number of the AWS Snowball Edge device. | 
**RegionId** | **string** | For AWS Snowball Edge, the region is &#x60;snow&#x60;. | 
**CredentialsId** | **string** | ID of the cloud credentials record. | 
**ConnectionSettings** | Pointer to [**ObjectStorageConnectionModel**](ObjectStorageConnectionModel.md) |  | [optional] 

## Methods

### NewAmazonSnowballEdgeStorageAccountModel

`func NewAmazonSnowballEdgeStorageAccountModel(servicePoint string, regionId string, credentialsId string, ) *AmazonSnowballEdgeStorageAccountModel`

NewAmazonSnowballEdgeStorageAccountModel instantiates a new AmazonSnowballEdgeStorageAccountModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSnowballEdgeStorageAccountModelWithDefaults

`func NewAmazonSnowballEdgeStorageAccountModelWithDefaults() *AmazonSnowballEdgeStorageAccountModel`

NewAmazonSnowballEdgeStorageAccountModelWithDefaults instantiates a new AmazonSnowballEdgeStorageAccountModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePoint

`func (o *AmazonSnowballEdgeStorageAccountModel) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *AmazonSnowballEdgeStorageAccountModel) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *AmazonSnowballEdgeStorageAccountModel) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.


### GetRegionId

`func (o *AmazonSnowballEdgeStorageAccountModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AmazonSnowballEdgeStorageAccountModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AmazonSnowballEdgeStorageAccountModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetCredentialsId

`func (o *AmazonSnowballEdgeStorageAccountModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *AmazonSnowballEdgeStorageAccountModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *AmazonSnowballEdgeStorageAccountModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetConnectionSettings

`func (o *AmazonSnowballEdgeStorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *AmazonSnowballEdgeStorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *AmazonSnowballEdgeStorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel)`

SetConnectionSettings sets ConnectionSettings field to given value.

### HasConnectionSettings

`func (o *AmazonSnowballEdgeStorageAccountModel) HasConnectionSettings() bool`

HasConnectionSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


