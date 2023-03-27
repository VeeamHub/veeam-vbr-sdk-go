# AmazonSnowballEdgeStorageAccountImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePoint** | **string** | Service point address and port number of the AWS Snowball Edge device. | 
**RegionId** | **string** | For AWS Snowball Edge, the region is *snow*. | 
**Credentials** | [**CloudCredentialsImportModel**](CloudCredentialsImportModel.md) |  | 
**ConnectionSettings** | [**ObjectStorageConnectionImportSpec**](ObjectStorageConnectionImportSpec.md) |  | 

## Methods

### NewAmazonSnowballEdgeStorageAccountImportModel

`func NewAmazonSnowballEdgeStorageAccountImportModel(servicePoint string, regionId string, credentials CloudCredentialsImportModel, connectionSettings ObjectStorageConnectionImportSpec, ) *AmazonSnowballEdgeStorageAccountImportModel`

NewAmazonSnowballEdgeStorageAccountImportModel instantiates a new AmazonSnowballEdgeStorageAccountImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSnowballEdgeStorageAccountImportModelWithDefaults

`func NewAmazonSnowballEdgeStorageAccountImportModelWithDefaults() *AmazonSnowballEdgeStorageAccountImportModel`

NewAmazonSnowballEdgeStorageAccountImportModelWithDefaults instantiates a new AmazonSnowballEdgeStorageAccountImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePoint

`func (o *AmazonSnowballEdgeStorageAccountImportModel) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *AmazonSnowballEdgeStorageAccountImportModel) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *AmazonSnowballEdgeStorageAccountImportModel) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.


### GetRegionId

`func (o *AmazonSnowballEdgeStorageAccountImportModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AmazonSnowballEdgeStorageAccountImportModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AmazonSnowballEdgeStorageAccountImportModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetCredentials

`func (o *AmazonSnowballEdgeStorageAccountImportModel) GetCredentials() CloudCredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AmazonSnowballEdgeStorageAccountImportModel) GetCredentialsOk() (*CloudCredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AmazonSnowballEdgeStorageAccountImportModel) SetCredentials(v CloudCredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetConnectionSettings

`func (o *AmazonSnowballEdgeStorageAccountImportModel) GetConnectionSettings() ObjectStorageConnectionImportSpec`

GetConnectionSettings returns the ConnectionSettings field if non-nil, zero value otherwise.

### GetConnectionSettingsOk

`func (o *AmazonSnowballEdgeStorageAccountImportModel) GetConnectionSettingsOk() (*ObjectStorageConnectionImportSpec, bool)`

GetConnectionSettingsOk returns a tuple with the ConnectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSettings

`func (o *AmazonSnowballEdgeStorageAccountImportModel) SetConnectionSettings(v ObjectStorageConnectionImportSpec)`

SetConnectionSettings sets ConnectionSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


