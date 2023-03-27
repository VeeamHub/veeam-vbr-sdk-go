# S3CompatibleBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server used to connect to the S3 compatible storage. | [optional] 
**ConnectionPoint** | Pointer to **string** | Service point address and port number of the S3 compatible storage. | [optional] 
**Regions** | Pointer to [**[]S3CompatibleRegionBrowserModel**](S3CompatibleRegionBrowserModel.md) | Array of regions. | [optional] 

## Methods

### NewS3CompatibleBrowserModel

`func NewS3CompatibleBrowserModel() *S3CompatibleBrowserModel`

NewS3CompatibleBrowserModel instantiates a new S3CompatibleBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CompatibleBrowserModelWithDefaults

`func NewS3CompatibleBrowserModelWithDefaults() *S3CompatibleBrowserModel`

NewS3CompatibleBrowserModelWithDefaults instantiates a new S3CompatibleBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *S3CompatibleBrowserModel) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *S3CompatibleBrowserModel) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *S3CompatibleBrowserModel) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *S3CompatibleBrowserModel) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetConnectionPoint

`func (o *S3CompatibleBrowserModel) GetConnectionPoint() string`

GetConnectionPoint returns the ConnectionPoint field if non-nil, zero value otherwise.

### GetConnectionPointOk

`func (o *S3CompatibleBrowserModel) GetConnectionPointOk() (*string, bool)`

GetConnectionPointOk returns a tuple with the ConnectionPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPoint

`func (o *S3CompatibleBrowserModel) SetConnectionPoint(v string)`

SetConnectionPoint sets ConnectionPoint field to given value.

### HasConnectionPoint

`func (o *S3CompatibleBrowserModel) HasConnectionPoint() bool`

HasConnectionPoint returns a boolean if a field has been set.

### GetRegions

`func (o *S3CompatibleBrowserModel) GetRegions() []S3CompatibleRegionBrowserModel`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *S3CompatibleBrowserModel) GetRegionsOk() (*[]S3CompatibleRegionBrowserModel, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *S3CompatibleBrowserModel) SetRegions(v []S3CompatibleRegionBrowserModel)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *S3CompatibleBrowserModel) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


